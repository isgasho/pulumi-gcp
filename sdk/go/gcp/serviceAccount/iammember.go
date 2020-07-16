// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the googleProjectIam set of resources.
//
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
//
// * `serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
//
// > **Note:** `serviceAccount.IAMPolicy` **cannot** be used in conjunction with `serviceAccount.IAMBinding` and `serviceAccount.IAMMember` or they will fight over what your policy should be.
//
// > **Note:** `serviceAccount.IAMBinding` resources **can be** used in conjunction with `serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
type IAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the service account IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServiceAccountId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountId'")
	}
	if args == nil {
		args = &IAMMemberArgs{}
	}
	var resource IAMMember
	err := ctx.RegisterResource("gcp:serviceAccount/iAMMember:IAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMMemberState, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	var resource IAMMember
	err := ctx.ReadResource("gcp:serviceAccount/iAMMember:IAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMMember resources.
type iammemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the service account IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type IAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	// (Computed) The etag of the service account IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringPtrInput
}

func (IAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberState)(nil)).Elem()
}

type iammemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	Member    string              `pulumi:"member"`
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The role that should be applied. Only one
	// `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringInput
}

func (IAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberArgs)(nil)).Elem()
}
