// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/machineImages/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine machineimage IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/machineImageIamMember:MachineImageIamMember editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/machineImageIamMember:MachineImageIamMember editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/machineImageIamMember:MachineImageIamMember editor projects/{{project}}/global/machineImages/{{machine_image}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MachineImageIamMember struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringOutput `pulumi:"machineImage"`
	Member       pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewMachineImageIamMember registers a new resource with the given unique name, arguments, and options.
func NewMachineImageIamMember(ctx *pulumi.Context,
	name string, args *MachineImageIamMemberArgs, opts ...pulumi.ResourceOption) (*MachineImageIamMember, error) {
	if args == nil || args.MachineImage == nil {
		return nil, errors.New("missing required argument 'MachineImage'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &MachineImageIamMemberArgs{}
	}
	var resource MachineImageIamMember
	err := ctx.RegisterResource("gcp:compute/machineImageIamMember:MachineImageIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineImageIamMember gets an existing MachineImageIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineImageIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineImageIamMemberState, opts ...pulumi.ResourceOption) (*MachineImageIamMember, error) {
	var resource MachineImageIamMember
	err := ctx.ReadResource("gcp:compute/machineImageIamMember:MachineImageIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineImageIamMember resources.
type machineImageIamMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *MachineImageIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage *string `pulumi:"machineImage"`
	Member       *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type MachineImageIamMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringPtrInput
	Member       pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (MachineImageIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamMemberState)(nil)).Elem()
}

type machineImageIamMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *MachineImageIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	MachineImage string `pulumi:"machineImage"`
	Member       string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a MachineImageIamMember resource.
type MachineImageIamMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition MachineImageIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	MachineImage pulumi.StringInput
	Member       pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.MachineImageIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (MachineImageIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineImageIamMemberArgs)(nil)).Elem()
}

type MachineImageIamMemberInput interface {
	pulumi.Input

	ToMachineImageIamMemberOutput() MachineImageIamMemberOutput
	ToMachineImageIamMemberOutputWithContext(ctx context.Context) MachineImageIamMemberOutput
}

func (MachineImageIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImageIamMember)(nil)).Elem()
}

func (i MachineImageIamMember) ToMachineImageIamMemberOutput() MachineImageIamMemberOutput {
	return i.ToMachineImageIamMemberOutputWithContext(context.Background())
}

func (i MachineImageIamMember) ToMachineImageIamMemberOutputWithContext(ctx context.Context) MachineImageIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineImageIamMemberOutput)
}

type MachineImageIamMemberOutput struct {
	*pulumi.OutputState
}

func (MachineImageIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineImageIamMemberOutput)(nil)).Elem()
}

func (o MachineImageIamMemberOutput) ToMachineImageIamMemberOutput() MachineImageIamMemberOutput {
	return o
}

func (o MachineImageIamMemberOutput) ToMachineImageIamMemberOutputWithContext(ctx context.Context) MachineImageIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineImageIamMemberOutput{})
}