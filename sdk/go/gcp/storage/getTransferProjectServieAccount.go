// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve Storage Transfer service account for this project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := storage.GetTransferProjectServieAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("defaultAccount", _default.Email)
// 		return nil
// 	})
// }
// ```
func GetTransferProjectServieAccount(ctx *pulumi.Context, args *GetTransferProjectServieAccountArgs, opts ...pulumi.InvokeOption) (*GetTransferProjectServieAccountResult, error) {
	var rv GetTransferProjectServieAccountResult
	err := ctx.Invoke("gcp:storage/getTransferProjectServieAccount:getTransferProjectServieAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransferProjectServieAccount.
type GetTransferProjectServieAccountArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getTransferProjectServieAccount.
type GetTransferProjectServieAccountResult struct {
	// Email address of the default service account used by Storage Transfer Jobs running in this project
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Project string `pulumi:"project"`
}
