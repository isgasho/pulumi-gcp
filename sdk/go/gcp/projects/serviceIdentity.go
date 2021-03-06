// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Generate service identity for a service.
//
// > **Note**: Once created, this resource cannot be updated or destroyed. These
// actions are a no-op.
//
// To get more information about Service Identity, see:
//
// * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1beta1/services/generateServiceIdentity)
//
// ## Example Usage
// ### Service Identity Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		hcSa, err := projects.NewServiceIdentity(ctx, "hcSa", &projects.ServiceIdentityArgs{
// 			Project: pulumi.String(project.ProjectId),
// 			Service: pulumi.String("healthcare.googleapis.com"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMMember(ctx, "hcSaBqJobuser", &projects.IAMMemberArgs{
// 			Project: pulumi.String(project.ProjectId),
// 			Role:    pulumi.String("roles/bigquery.jobUser"),
// 			Member: hcSa.Email.ApplyT(func(email string) (string, error) {
// 				return fmt.Sprintf("%v%v", "serviceAccount:", email), nil
// 			}).(pulumi.StringOutput),
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
// This resource does not support import.
type ServiceIdentity struct {
	pulumi.CustomResourceState

	// The email address of the Google managed service account.
	Email pulumi.StringOutput `pulumi:"email"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The service to generate identity for.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewServiceIdentity registers a new resource with the given unique name, arguments, and options.
func NewServiceIdentity(ctx *pulumi.Context,
	name string, args *ServiceIdentityArgs, opts ...pulumi.ResourceOption) (*ServiceIdentity, error) {
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &ServiceIdentityArgs{}
	}
	var resource ServiceIdentity
	err := ctx.RegisterResource("gcp:projects/serviceIdentity:ServiceIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIdentity gets an existing ServiceIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIdentityState, opts ...pulumi.ResourceOption) (*ServiceIdentity, error) {
	var resource ServiceIdentity
	err := ctx.ReadResource("gcp:projects/serviceIdentity:ServiceIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIdentity resources.
type serviceIdentityState struct {
	// The email address of the Google managed service account.
	Email *string `pulumi:"email"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service to generate identity for.
	Service *string `pulumi:"service"`
}

type ServiceIdentityState struct {
	// The email address of the Google managed service account.
	Email pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service to generate identity for.
	Service pulumi.StringPtrInput
}

func (ServiceIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIdentityState)(nil)).Elem()
}

type serviceIdentityArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service to generate identity for.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a ServiceIdentity resource.
type ServiceIdentityArgs struct {
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service to generate identity for.
	Service pulumi.StringInput
}

func (ServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIdentityArgs)(nil)).Elem()
}

type ServiceIdentityInput interface {
	pulumi.Input

	ToServiceIdentityOutput() ServiceIdentityOutput
	ToServiceIdentityOutputWithContext(ctx context.Context) ServiceIdentityOutput
}

func (ServiceIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIdentity)(nil)).Elem()
}

func (i ServiceIdentity) ToServiceIdentityOutput() ServiceIdentityOutput {
	return i.ToServiceIdentityOutputWithContext(context.Background())
}

func (i ServiceIdentity) ToServiceIdentityOutputWithContext(ctx context.Context) ServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceIdentityOutput)
}

type ServiceIdentityOutput struct {
	*pulumi.OutputState
}

func (ServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceIdentityOutput)(nil)).Elem()
}

func (o ServiceIdentityOutput) ToServiceIdentityOutput() ServiceIdentityOutput {
	return o
}

func (o ServiceIdentityOutput) ToServiceIdentityOutputWithContext(ctx context.Context) ServiceIdentityOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServiceIdentityOutput{})
}
