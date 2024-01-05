package components

import (
	"fmt"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type (
	CTFd struct {
		rd *Redis

		dep *appsv1.Deployment
		svc *corev1.Service

		Port pulumi.IntOutput
	}

	CTFdArgs struct {
		Namespace pulumi.StringInput
	}
)

// NewCTFd deploys a minimal CTFd configuration with just enough
// Kubernetes infrastructure to test the Go API wrapper.
//
// WARNING: Do not use this component for production purposes.
func NewCTFd(ctx *pulumi.Context, args *CTFdArgs, opts ...pulumi.ResourceOption) (*CTFd, error) {
	if args == nil {
		args = &CTFdArgs{}
	}

	ctfd := &CTFd{}

	if err := ctfd.provision(ctx, args, opts...); err != nil {
		return nil, err
	}

	ctfd.outputs(ctx)
	return ctfd, nil
}

func (ctfd *CTFd) provision(ctx *pulumi.Context, args *CTFdArgs, opts ...pulumi.ResourceOption) (err error) {
	uid := randName()

	// Dependencies
	ctfd.rd, err = NewRedis(ctx, &RedisArgs{
		Namespace: args.Namespace,
	}, opts...)
	if err != nil {
		return
	}

	// Uniquely identify the resources with labels
	labels := pulumi.ToStringMap(map[string]string{
		"app":        "ctfd",
		"repository": "github.com_ctfer-io_go-ctfd",
	})

	// => Deployment
	ctfd.dep, err = appsv1.NewDeployment(ctx, "ctfd-dep-"+uid, &appsv1.DeploymentArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name:      pulumi.String("ctfd-dep-" + uid),
			Namespace: args.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpecArgs{
			Selector: metav1.LabelSelectorArgs{
				MatchLabels: labels,
			},
			Replicas: pulumi.Int(1),
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Namespace: args.Namespace,
					Labels:    labels,
				},
				Spec: &corev1.PodSpecArgs{
					InitContainers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name: pulumi.String("wait-for-redis"),
							// TODO rebuild image or replace
							Image: pulumi.String("redis:7.0.10"),
							Args: pulumi.StringArray{
								pulumi.String("sh"), pulumi.String("-c"),
								ctfd.rd.URL.ApplyT(func(url string) string {
									return fmt.Sprintf("until redis-cli -u %s get hello; do echo \"Sleeping a bit\"; sleep 1; done; echo \"ready!\";", url)
								}).(pulumi.StringOutput),
								// pulumi.All(ctfd.rd.svc.Metadata, ctfd.rd.svc.Spec).ApplyT(func(args []any) string {
								// 	meta := args[0].(metav1.ObjectMeta)
								// 	spec := args[1].(corev1.ServiceSpec)

								// }).(pulumi.StringOutput),
							},
						},
					},
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:  pulumi.String("ctfd"),
							Image: pulumi.String("ctfd/ctfd:3.6.1"),
							Ports: corev1.ContainerPortArray{
								corev1.ContainerPortArgs{
									ContainerPort: pulumi.Int(8000),
								},
							},
							Env: corev1.EnvVarArray{
								corev1.EnvVarArgs{
									Name:  pulumi.String("REDIS_URL"),
									Value: ctfd.rd.URL,
								},
							},
							ReadinessProbe: corev1.ProbeArgs{
								HttpGet: corev1.HTTPGetActionArgs{
									Path: pulumi.String("/setup"),
									Port: pulumi.Int(8000),
								},
							},
						},
					},
				},
			},
		},
	}, opts...)
	if err != nil {
		return
	}

	ctfd.svc, err = corev1.NewService(ctx, "ctfd-svc-"+uid, &corev1.ServiceArgs{
		Metadata: metav1.ObjectMetaArgs{
			Labels:    labels,
			Name:      pulumi.String("ctfd-svc-" + uid),
			Namespace: args.Namespace,
		},
		Spec: &corev1.ServiceSpecArgs{
			Selector: labels,
			Type:     pulumi.String("NodePort"),
			Ports: corev1.ServicePortArray{
				corev1.ServicePortArgs{
					TargetPort: pulumi.Int(8000),
					Port:       pulumi.Int(8000),
					Name:       pulumi.String("web"),
				},
			},
		},
	}, opts...)
	if err != nil {
		return
	}

	return nil
}

func (ctfd *CTFd) outputs(ctx *pulumi.Context) {
	ctfd.Port = ctfd.svc.Spec.ApplyT(func(spec corev1.ServiceSpec) int {
		if spec.ClusterIP == nil {
			return 0
		}
		return *spec.Ports[0].NodePort
	}).(pulumi.IntOutput)
}
