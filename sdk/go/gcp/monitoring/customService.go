// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Service is a discrete, autonomous, and network-accessible unit,
// designed to solve an individual concern (Wikipedia). In Cloud Monitoring,
// a Service acts as the root resource under which operational aspects of
// the service are accessible
//
// To get more information about Service, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
// * How-to Guides
//     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
//     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Monitoring Service Custom
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := monitoring.NewCustomService(ctx, "custom", &monitoring.CustomServiceArgs{
// 			DisplayName: pulumi.String("My Custom Service custom-srv"),
// 			ServiceId:   pulumi.String("custom-srv"),
// 			Telemetry: &monitoring.CustomServiceTelemetryArgs{
// 				ResourceName: pulumi.String("//product.googleapis.com/foo/foo/services/test"),
// 			},
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
// Service can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:monitoring/customService:CustomService default {{name}}
// ```
type CustomService struct {
	pulumi.CustomResourceState

	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry CustomServiceTelemetryPtrOutput `pulumi:"telemetry"`
}

// NewCustomService registers a new resource with the given unique name, arguments, and options.
func NewCustomService(ctx *pulumi.Context,
	name string, args *CustomServiceArgs, opts ...pulumi.ResourceOption) (*CustomService, error) {
	if args == nil {
		args = &CustomServiceArgs{}
	}
	var resource CustomService
	err := ctx.RegisterResource("gcp:monitoring/customService:CustomService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomService gets an existing CustomService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomServiceState, opts ...pulumi.ResourceOption) (*CustomService, error) {
	var resource CustomService
	err := ctx.ReadResource("gcp:monitoring/customService:CustomService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomService resources.
type customServiceState struct {
	// Name used for UI elements listing this Service.
	DisplayName *string `pulumi:"displayName"`
	// The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId *string `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry *CustomServiceTelemetry `pulumi:"telemetry"`
}

type CustomServiceState struct {
	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrInput
	// The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId pulumi.StringPtrInput
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry CustomServiceTelemetryPtrInput
}

func (CustomServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*customServiceState)(nil)).Elem()
}

type customServiceArgs struct {
	// Name used for UI elements listing this Service.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId *string `pulumi:"serviceId"`
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry *CustomServiceTelemetry `pulumi:"telemetry"`
}

// The set of arguments for constructing a CustomService resource.
type CustomServiceArgs struct {
	// Name used for UI elements listing this Service.
	DisplayName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceId pulumi.StringPtrInput
	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry CustomServiceTelemetryPtrInput
}

func (CustomServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customServiceArgs)(nil)).Elem()
}

type CustomServiceInput interface {
	pulumi.Input

	ToCustomServiceOutput() CustomServiceOutput
	ToCustomServiceOutputWithContext(ctx context.Context) CustomServiceOutput
}

func (CustomService) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomService)(nil)).Elem()
}

func (i CustomService) ToCustomServiceOutput() CustomServiceOutput {
	return i.ToCustomServiceOutputWithContext(context.Background())
}

func (i CustomService) ToCustomServiceOutputWithContext(ctx context.Context) CustomServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomServiceOutput)
}

type CustomServiceOutput struct {
	*pulumi.OutputState
}

func (CustomServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomServiceOutput)(nil)).Elem()
}

func (o CustomServiceOutput) ToCustomServiceOutput() CustomServiceOutput {
	return o
}

func (o CustomServiceOutput) ToCustomServiceOutputWithContext(ctx context.Context) CustomServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomServiceOutput{})
}
