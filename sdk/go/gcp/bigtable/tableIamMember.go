// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:
//
// * `bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
// * `bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
//
// > **Note:** `bigtable.TableIamPolicy` **cannot** be used in conjunction with `bigtable.TableIamBinding` and `bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `bigtable.TableIamPolicy` replaces the entire policy.
//
// > **Note:** `bigtable.TableIamBinding` resources **can be** used in conjunction with `bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigtable\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/bigtable.user",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigtable.NewTableIamPolicy(ctx, "editor", &bigtable.TableIamPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Instance:   pulumi.String("your-bigtable-instance"),
// 			Table:      pulumi.String("your-bigtable-table"),
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
// ## google\_bigtable\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewTableIamBinding(ctx, "editor", &bigtable.TableIamBindingArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role:  pulumi.String("roles/bigtable.user"),
// 			Table: pulumi.String("your-bigtable-table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_bigtable\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewTableIamMember(ctx, "editor", &bigtable.TableIamMemberArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Role:     pulumi.String("roles/bigtable.user"),
// 			Table:    pulumi.String("your-bigtable-table"),
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
// Instance IAM resources can be imported using the project, table name, role and/or member.
//
// ```sh
//  $ pulumi import gcp:bigtable/tableIamMember:TableIamMember editor "projects/{project}/tables/{table}"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/tableIamMember:TableIamMember editor "projects/{project}/tables/{table} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/tableIamMember:TableIamMember editor "projects/{project}/tables/{table} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TableIamMember struct {
	pulumi.CustomResourceState

	Condition TableIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringOutput `pulumi:"instance"`
	Member   pulumi.StringOutput `pulumi:"member"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringOutput `pulumi:"role"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringOutput `pulumi:"table"`
}

// NewTableIamMember registers a new resource with the given unique name, arguments, and options.
func NewTableIamMember(ctx *pulumi.Context,
	name string, args *TableIamMemberArgs, opts ...pulumi.ResourceOption) (*TableIamMember, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Table == nil {
		return nil, errors.New("missing required argument 'Table'")
	}
	if args == nil {
		args = &TableIamMemberArgs{}
	}
	var resource TableIamMember
	err := ctx.RegisterResource("gcp:bigtable/tableIamMember:TableIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableIamMember gets an existing TableIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableIamMemberState, opts ...pulumi.ResourceOption) (*TableIamMember, error) {
	var resource TableIamMember
	err := ctx.ReadResource("gcp:bigtable/tableIamMember:TableIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableIamMember resources.
type tableIamMemberState struct {
	Condition *TableIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the tables's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance that owns the table.
	Instance *string `pulumi:"instance"`
	Member   *string `pulumi:"member"`
	Project  *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role *string `pulumi:"role"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table *string `pulumi:"table"`
}

type TableIamMemberState struct {
	Condition TableIamMemberConditionPtrInput
	// (Computed) The etag of the tables's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringPtrInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringPtrInput
}

func (TableIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamMemberState)(nil)).Elem()
}

type tableIamMemberArgs struct {
	Condition *TableIamMemberCondition `pulumi:"condition"`
	// The name or relative resource id of the instance that owns the table.
	Instance string  `pulumi:"instance"`
	Member   string  `pulumi:"member"`
	Project  *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role string `pulumi:"role"`
	// The name or relative resource id of the table to manage IAM policies for.
	Table string `pulumi:"table"`
}

// The set of arguments for constructing a TableIamMember resource.
type TableIamMemberArgs struct {
	Condition TableIamMemberConditionPtrInput
	// The name or relative resource id of the instance that owns the table.
	Instance pulumi.StringInput
	Member   pulumi.StringInput
	Project  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringInput
	// The name or relative resource id of the table to manage IAM policies for.
	Table pulumi.StringInput
}

func (TableIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableIamMemberArgs)(nil)).Elem()
}

type TableIamMemberInput interface {
	pulumi.Input

	ToTableIamMemberOutput() TableIamMemberOutput
	ToTableIamMemberOutputWithContext(ctx context.Context) TableIamMemberOutput
}

func (TableIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*TableIamMember)(nil)).Elem()
}

func (i TableIamMember) ToTableIamMemberOutput() TableIamMemberOutput {
	return i.ToTableIamMemberOutputWithContext(context.Background())
}

func (i TableIamMember) ToTableIamMemberOutputWithContext(ctx context.Context) TableIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableIamMemberOutput)
}

type TableIamMemberOutput struct {
	*pulumi.OutputState
}

func (TableIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableIamMemberOutput)(nil)).Elem()
}

func (o TableIamMemberOutput) ToTableIamMemberOutput() TableIamMemberOutput {
	return o
}

func (o TableIamMemberOutput) ToTableIamMemberOutputWithContext(ctx context.Context) TableIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TableIamMemberOutput{})
}