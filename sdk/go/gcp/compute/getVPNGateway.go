// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get a VPN gateway within GCE from its name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.LookupVPNGateway(ctx, &compute.LookupVPNGatewayArgs{
// 			Name: "vpn-gateway-us-east1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVPNGateway(ctx *pulumi.Context, args *LookupVPNGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVPNGatewayResult, error) {
	var rv LookupVPNGatewayResult
	err := ctx.Invoke("gcp:compute/getVPNGateway:getVPNGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVPNGateway.
type LookupVPNGatewayArgs struct {
	// The name of the VPN gateway.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the project region is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVPNGateway.
type LookupVPNGatewayResult struct {
	// Description of this VPN gateway.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The network of this VPN gateway.
	Network string `pulumi:"network"`
	Project string `pulumi:"project"`
	// Region of this VPN gateway.
	Region string `pulumi:"region"`
	// The URI of the resource.
	SelfLink string `pulumi:"selfLink"`
}
