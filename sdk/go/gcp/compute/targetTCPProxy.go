// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a TargetTcpProxy resource, which is used by one or more
// global forwarding rule to route incoming TCP requests to a Backend
// service.
//
// To get more information about TargetTcpProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetTcpProxies)
// * How-to Guides
//     * [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
//
// ## Example Usage
// ### Target Tcp Proxy Basic
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
// 		defaultHealthCheck, err := compute.NewHealthCheck(ctx, "defaultHealthCheck", &compute.HealthCheckArgs{
// 			TimeoutSec:       pulumi.Int(1),
// 			CheckIntervalSec: pulumi.Int(1),
// 			TcpHealthCheck: &compute.HealthCheckTcpHealthCheckArgs{
// 				Port: pulumi.Int(443),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultBackendService, err := compute.NewBackendService(ctx, "defaultBackendService", &compute.BackendServiceArgs{
// 			Protocol:   pulumi.String("TCP"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String(pulumi.String{
// 				defaultHealthCheck.ID(),
// 			}),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewTargetTCPProxy(ctx, "defaultTargetTCPProxy", &compute.TargetTCPProxyArgs{
// 			BackendService: defaultBackendService.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// TargetTcpProxy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default projects/{{project}}/global/targetTcpProxies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{name}}
// ```
type TargetTCPProxy struct {
	pulumi.CustomResourceState

	// A reference to the BackendService resource.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrOutput `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewTargetTCPProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetTCPProxy(ctx *pulumi.Context,
	name string, args *TargetTCPProxyArgs, opts ...pulumi.ResourceOption) (*TargetTCPProxy, error) {
	if args == nil || args.BackendService == nil {
		return nil, errors.New("missing required argument 'BackendService'")
	}
	if args == nil {
		args = &TargetTCPProxyArgs{}
	}
	var resource TargetTCPProxy
	err := ctx.RegisterResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetTCPProxy gets an existing TargetTCPProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetTCPProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetTCPProxyState, opts ...pulumi.ResourceOption) (*TargetTCPProxy, error) {
	var resource TargetTCPProxy
	err := ctx.ReadResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetTCPProxy resources.
type targetTCPProxyState struct {
	// A reference to the BackendService resource.
	BackendService *string `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type TargetTCPProxyState struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (TargetTCPProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetTCPProxyState)(nil)).Elem()
}

type targetTCPProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService string `pulumi:"backendService"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader *string `pulumi:"proxyHeader"`
}

// The set of arguments for constructing a TargetTCPProxy resource.
type TargetTCPProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrInput
}

func (TargetTCPProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetTCPProxyArgs)(nil)).Elem()
}

type TargetTCPProxyInput interface {
	pulumi.Input

	ToTargetTCPProxyOutput() TargetTCPProxyOutput
	ToTargetTCPProxyOutputWithContext(ctx context.Context) TargetTCPProxyOutput
}

func (TargetTCPProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetTCPProxy)(nil)).Elem()
}

func (i TargetTCPProxy) ToTargetTCPProxyOutput() TargetTCPProxyOutput {
	return i.ToTargetTCPProxyOutputWithContext(context.Background())
}

func (i TargetTCPProxy) ToTargetTCPProxyOutputWithContext(ctx context.Context) TargetTCPProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetTCPProxyOutput)
}

type TargetTCPProxyOutput struct {
	*pulumi.OutputState
}

func (TargetTCPProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetTCPProxyOutput)(nil)).Elem()
}

func (o TargetTCPProxyOutput) ToTargetTCPProxyOutput() TargetTCPProxyOutput {
	return o
}

func (o TargetTCPProxyOutput) ToTargetTCPProxyOutputWithContext(ctx context.Context) TargetTCPProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TargetTCPProxyOutput{})
}
