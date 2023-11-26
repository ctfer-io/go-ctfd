package main

import (
	"github.com/ctfer-io/go-ctfd/deploy/components"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config := config.New(ctx, "go-ctfd")
		ns := config.Get("namespace") // default value "" is fine

		ctfd, err := components.NewCTFd(ctx, &components.CTFdArgs{
			Namespace: pulumi.String(ns),
		})
		if err != nil {
			return err
		}

		ctx.Export("port", ctfd.Port)

		return nil
	})
}
