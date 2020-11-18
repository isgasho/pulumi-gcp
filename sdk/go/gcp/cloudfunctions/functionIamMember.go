// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Functions CloudFunction. Each of these resources serves a different use case:
//
// * `cloudfunctions.FunctionIamPolicy`: Authoritative. Sets the IAM policy for the cloudfunction and replaces any existing policy already attached.
// * `cloudfunctions.FunctionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cloudfunction are preserved.
// * `cloudfunctions.FunctionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cloudfunction are preserved.
//
// > **Note:** `cloudfunctions.FunctionIamPolicy` **cannot** be used in conjunction with `cloudfunctions.FunctionIamBinding` and `cloudfunctions.FunctionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudfunctions.FunctionIamBinding` resources **can be** used in conjunction with `cloudfunctions.FunctionIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloudfunctions\_function\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudfunctions"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudfunctions.NewFunctionIamPolicy(ctx, "policy", &cloudfunctions.FunctionIamPolicyArgs{
// 			Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
// 			Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
// 			CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
// 			PolicyData:    pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_cloudfunctions\_function\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudfunctions"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfunctions.NewFunctionIamBinding(ctx, "binding", &cloudfunctions.FunctionIamBindingArgs{
// 			Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
// 			Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
// 			CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
// 			Role:          pulumi.String("roles/viewer"),
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
// ## google\_cloudfunctions\_function\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudfunctions"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfunctions.NewFunctionIamMember(ctx, "member", &cloudfunctions.FunctionIamMemberArgs{
// 			Project:       pulumi.Any(google_cloudfunctions_function.Function.Project),
// 			Region:        pulumi.Any(google_cloudfunctions_function.Function.Region),
// 			CloudFunction: pulumi.Any(google_cloudfunctions_function.Function.Name),
// 			Role:          pulumi.String("roles/viewer"),
// 			Member:        pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} * {{project}}/{{region}}/{{cloud_function}} * {{region}}/{{cloud_function}} * {{cloud_function}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Functions cloudfunction IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudfunctions/functionIamMember:FunctionIamMember editor "projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudfunctions/functionIamMember:FunctionIamMember editor "projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudfunctions/functionIamMember:FunctionIamMember editor projects/{{project}}/locations/{{region}}/functions/{{cloud_function}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type FunctionIamMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringOutput                 `pulumi:"cloudFunction"`
	Condition     FunctionIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFunctionIamMember registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamMember(ctx *pulumi.Context,
	name string, args *FunctionIamMemberArgs, opts ...pulumi.ResourceOption) (*FunctionIamMember, error) {
	if args == nil || args.CloudFunction == nil {
		return nil, errors.New("missing required argument 'CloudFunction'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &FunctionIamMemberArgs{}
	}
	var resource FunctionIamMember
	err := ctx.RegisterResource("gcp:cloudfunctions/functionIamMember:FunctionIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionIamMember gets an existing FunctionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionIamMemberState, opts ...pulumi.ResourceOption) (*FunctionIamMember, error) {
	var resource FunctionIamMember
	err := ctx.ReadResource("gcp:cloudfunctions/functionIamMember:FunctionIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionIamMember resources.
type functionIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction *string                     `pulumi:"cloudFunction"`
	Condition     *FunctionIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FunctionIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringPtrInput
	Condition     FunctionIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FunctionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamMemberState)(nil)).Elem()
}

type functionIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction string                      `pulumi:"cloudFunction"`
	Condition     *FunctionIamMemberCondition `pulumi:"condition"`
	Member        string                      `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FunctionIamMember resource.
type FunctionIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction pulumi.StringInput
	Condition     FunctionIamMemberConditionPtrInput
	Member        pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FunctionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionIamMemberArgs)(nil)).Elem()
}

type FunctionIamMemberInput interface {
	pulumi.Input

	ToFunctionIamMemberOutput() FunctionIamMemberOutput
	ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput
}

func (FunctionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionIamMember)(nil)).Elem()
}

func (i FunctionIamMember) ToFunctionIamMemberOutput() FunctionIamMemberOutput {
	return i.ToFunctionIamMemberOutputWithContext(context.Background())
}

func (i FunctionIamMember) ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionIamMemberOutput)
}

type FunctionIamMemberOutput struct {
	*pulumi.OutputState
}

func (FunctionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionIamMemberOutput)(nil)).Elem()
}

func (o FunctionIamMemberOutput) ToFunctionIamMemberOutput() FunctionIamMemberOutput {
	return o
}

func (o FunctionIamMemberOutput) ToFunctionIamMemberOutputWithContext(ctx context.Context) FunctionIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FunctionIamMemberOutput{})
}
