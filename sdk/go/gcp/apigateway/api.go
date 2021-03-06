// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Api can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:apigateway/api:Api default projects/{{project}}/locations/global/apis/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:apigateway/api:Api default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:apigateway/api:Api default {{name}}
// ```
type Api struct {
	pulumi.CustomResourceState

	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A user-visible name for the API.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringOutput `pulumi:"managedService"`
	// The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil {
		args = &ApiArgs{}
	}
	var resource Api
	err := ctx.RegisterResource("gcp:apigateway/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("gcp:apigateway/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId *string `pulumi:"apiId"`
	// Creation timestamp in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService *string `pulumi:"managedService"`
	// The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApiState struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringPtrInput
	// The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId string `pulumi:"apiId"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService *string `pulumi:"managedService"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId pulumi.StringInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (Api) ElementType() reflect.Type {
	return reflect.TypeOf((*Api)(nil)).Elem()
}

func (i Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct {
	*pulumi.OutputState
}

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOutput)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiOutput{})
}
