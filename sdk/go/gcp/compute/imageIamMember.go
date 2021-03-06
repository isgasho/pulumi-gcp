// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Image. Each of these resources serves a different use case:
//
// * `compute.ImageIamPolicy`: Authoritative. Sets the IAM policy for the image and replaces any existing policy already attached.
// * `compute.ImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the image are preserved.
// * `compute.ImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the image are preserved.
//
// > **Note:** `compute.ImageIamPolicy` **cannot** be used in conjunction with `compute.ImageIamBinding` and `compute.ImageIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.ImageIamBinding` resources **can be** used in conjunction with `compute.ImageIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_image\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/compute.imageUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewImageIamPolicy(ctx, "policy", &compute.ImageIamPolicyArgs{
// 			Project:    pulumi.Any(google_compute_image.Example.Project),
// 			Image:      pulumi.Any(google_compute_image.Example.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/compute.imageUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewImageIamPolicy(ctx, "policy", &compute.ImageIamPolicyArgs{
// 			Project:    pulumi.Any(google_compute_image.Example.Project),
// 			Image:      pulumi.Any(google_compute_image.Example.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_image\_iam\_binding
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
// 		_, err := compute.NewImageIamBinding(ctx, "binding", &compute.ImageIamBindingArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
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
// 		_, err := compute.NewImageIamBinding(ctx, "binding", &compute.ImageIamBindingArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &compute.ImageIamBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_compute\_image\_iam\_member
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
// 		_, err := compute.NewImageIamMember(ctx, "member", &compute.ImageIamMemberArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
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
// 		_, err := compute.NewImageIamMember(ctx, "member", &compute.ImageIamMemberArgs{
// 			Project: pulumi.Any(google_compute_image.Example.Project),
// 			Image:   pulumi.Any(google_compute_image.Example.Name),
// 			Role:    pulumi.String("roles/compute.imageUser"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Condition: &compute.ImageIamMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/images/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine image IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/imageIamMember:ImageIamMember editor "projects/{{project}}/global/images/{{image}} roles/compute.imageUser user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/imageIamMember:ImageIamMember editor "projects/{{project}}/global/images/{{image}} roles/compute.imageUser"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/imageIamMember:ImageIamMember editor projects/{{project}}/global/images/{{image}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ImageIamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition ImageIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Image  pulumi.StringOutput `pulumi:"image"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewImageIamMember registers a new resource with the given unique name, arguments, and options.
func NewImageIamMember(ctx *pulumi.Context,
	name string, args *ImageIamMemberArgs, opts ...pulumi.ResourceOption) (*ImageIamMember, error) {
	if args == nil || args.Image == nil {
		return nil, errors.New("missing required argument 'Image'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ImageIamMemberArgs{}
	}
	var resource ImageIamMember
	err := ctx.RegisterResource("gcp:compute/imageIamMember:ImageIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageIamMember gets an existing ImageIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageIamMemberState, opts ...pulumi.ResourceOption) (*ImageIamMember, error) {
	var resource ImageIamMember
	err := ctx.ReadResource("gcp:compute/imageIamMember:ImageIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageIamMember resources.
type imageIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *ImageIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Image  *string `pulumi:"image"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ImageIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition ImageIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Image  pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ImageIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageIamMemberState)(nil)).Elem()
}

type imageIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *ImageIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Image  string `pulumi:"image"`
	Member string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ImageIamMember resource.
type ImageIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition ImageIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Image  pulumi.StringInput
	Member pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.ImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ImageIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageIamMemberArgs)(nil)).Elem()
}

type ImageIamMemberInput interface {
	pulumi.Input

	ToImageIamMemberOutput() ImageIamMemberOutput
	ToImageIamMemberOutputWithContext(ctx context.Context) ImageIamMemberOutput
}

func (ImageIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageIamMember)(nil)).Elem()
}

func (i ImageIamMember) ToImageIamMemberOutput() ImageIamMemberOutput {
	return i.ToImageIamMemberOutputWithContext(context.Background())
}

func (i ImageIamMember) ToImageIamMemberOutputWithContext(ctx context.Context) ImageIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamMemberOutput)
}

type ImageIamMemberOutput struct {
	*pulumi.OutputState
}

func (ImageIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageIamMemberOutput)(nil)).Elem()
}

func (o ImageIamMemberOutput) ToImageIamMemberOutput() ImageIamMemberOutput {
	return o
}

func (o ImageIamMemberOutput) ToImageIamMemberOutputWithContext(ctx context.Context) ImageIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageIamMemberOutput{})
}
