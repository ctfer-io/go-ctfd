package components

import (
	"fmt"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type (
	Redis struct {
		rand *random.RandomPassword
		sec  *corev1.Secret
		sts  *appsv1.StatefulSet
		svc  *corev1.Service

		URL pulumi.StringOutput
	}

	RedisArgs struct {
		Namespace pulumi.StringInput
	}
)

func NewRedis(ctx *pulumi.Context, args *RedisArgs, opts ...pulumi.ResourceOption) (*Redis, error) {
	if args == nil {
		args = &RedisArgs{}
	}

	rd := &Redis{}

	if err := rd.provision(ctx, args, opts...); err != nil {
		return nil, err
	}

	rd.outputs()
	return rd, nil
}

func (rd *Redis) provision(ctx *pulumi.Context, args *RedisArgs, opts ...pulumi.ResourceOption) (err error) {
	uid := randName()

	// Uniquely identify the resources with labels
	labels := pulumi.ToStringMap(map[string]string{
		"app":        "redis",
		"repository": "github.com_ctfer-io_go-ctfd",
	})

	// => Credentials
	rd.rand, err = random.NewRandomPassword(ctx, "redis-pass-"+uid, &random.RandomPasswordArgs{
		Length:  pulumi.Int(64),
		Special: pulumi.BoolPtr(false),
	}, opts...)
	if err != nil {
		return err
	}

	// => Service
	rd.svc, err = corev1.NewService(ctx, "redis-svc-"+uid, &corev1.ServiceArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name:      pulumi.String("redis-svc-" + uid),
			Labels:    labels,
			Namespace: args.Namespace,
		},
		Spec: corev1.ServiceSpecArgs{
			Ports: corev1.ServicePortArray{
				corev1.ServicePortArgs{
					Port:       pulumi.Int(6379),
					TargetPort: pulumi.Int(6379),
					Name:       pulumi.String("client"),
				},
			},
			// Headless, for DNS purposes
			ClusterIP: pulumi.String("None"),
			Selector:  labels,
		},
	}, opts...)
	if err != nil {
		return err
	}

	// /!\ Register output URL /!\
	rd.URL = pulumi.All(rd.rand.Result, rd.svc.Metadata, rd.svc.Spec).ApplyT(func(args []any) string {
		rand := args[0].(string)
		meta := args[1].(metav1.ObjectMeta)
		spec := args[2].(corev1.ServiceSpec)

		return fmt.Sprintf("redis://:%s@%s:%d", rand, *meta.Name, spec.Ports[0].Port)
	}).(pulumi.StringOutput)

	// => Secret
	rd.sec, err = corev1.NewSecret(ctx, "redis-secret-"+uid, &corev1.SecretArgs{
		Metadata: metav1.ObjectMetaArgs{
			Labels:    labels,
			Name:      pulumi.String("redis-secret-" + uid),
			Namespace: args.Namespace,
		},
		Type: pulumi.String("Opaque"),
		StringData: pulumi.ToStringMapOutput(map[string]pulumi.StringOutput{
			"redis-password": rd.rand.Result,
			"redis-url":      rd.URL,
		}),
	}, opts...)
	if err != nil {
		return err
	}

	// => StatefulSet
	rd.sts, err = appsv1.NewStatefulSet(ctx, "redis-sts-"+uid, &appsv1.StatefulSetArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name:      pulumi.String("redis-sts-" + uid),
			Labels:    labels,
			Namespace: args.Namespace,
		},
		Spec: appsv1.StatefulSetSpecArgs{
			ServiceName: rd.svc.Metadata.Name().Elem(),
			Replicas:    pulumi.Int(1),
			Selector: metav1.LabelSelectorArgs{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpecArgs{
				Metadata: metav1.ObjectMetaArgs{
					Namespace: args.Namespace,
					Labels:    labels,
				},
				Spec: corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:  pulumi.String("redis"),
							Image: pulumi.String("redis:7.0.10"),
							Ports: corev1.ContainerPortArray{
								corev1.ContainerPortArgs{
									ContainerPort: pulumi.Int(6379),
									Name:          pulumi.String("client"),
								},
							},
							Args: pulumi.ToStringArray([]string{
								"--requirepass",
								"$(REDIS_PASSWORD)",
							}),
							Env: corev1.EnvVarArray{
								corev1.EnvVarArgs{
									Name: pulumi.String("REDIS_PASSWORD"),
									ValueFrom: corev1.EnvVarSourceArgs{
										SecretKeyRef: corev1.SecretKeySelectorArgs{
											Name: rd.sec.Metadata.Name(),
											Key:  pulumi.String("redis-password"),
										},
									},
								},
							},
							VolumeMounts: corev1.VolumeMountArray{
								corev1.VolumeMountArgs{
									Name:      pulumi.String("data"),
									MountPath: pulumi.String("/data"),
									ReadOnly:  pulumi.Bool(false),
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: corev1.PersistentVolumeClaimTypeArray{
				corev1.PersistentVolumeClaimTypeArgs{
					Metadata: metav1.ObjectMetaArgs{
						Name: pulumi.String("data"),
					},
					Spec: corev1.PersistentVolumeClaimSpecArgs{
						AccessModes: pulumi.ToStringArray([]string{
							"ReadWriteOnce",
						}),
						Resources: corev1.VolumeResourceRequirementsArgs{
							Requests: pulumi.ToStringMap(map[string]string{
								"storage": "1Gi",
							}),
						},
					},
				},
			},
		},
	}, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (rd *Redis) outputs() {
	// rd.URL has already been registered during provisionning
}
