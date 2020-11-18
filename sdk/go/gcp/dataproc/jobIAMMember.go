// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
//
// * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
// * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
// * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
//
// > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.
//
// > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_dataproc\_job\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataproc.NewJobIAMPolicy(ctx, "editor", &dataproc.JobIAMPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Region:     pulumi.String("your-region"),
// 			JobId:      pulumi.String("your-dataproc-job"),
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
// ## google\_dataproc\_job\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewJobIAMBinding(ctx, "editor", &dataproc.JobIAMBindingArgs{
// 			JobId: pulumi.String("your-dataproc-job"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_dataproc\_job\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/dataproc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataproc.NewJobIAMMember(ctx, "editor", &dataproc.JobIAMMemberArgs{
// 			JobId:  pulumi.String("your-dataproc-job"),
// 			Member: pulumi.String("user:jane@example.com"),
// 			Role:   pulumi.String("roles/editor"),
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
// Job IAM resources can be imported using the project, region, job id, role and/or member.
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMMember:JobIAMMember editor "projects/{project}/regions/{region}/jobs/{job_id}"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMMember:JobIAMMember editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:dataproc/jobIAMMember:JobIAMMember editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type JobIAMMember struct {
	pulumi.CustomResourceState

	Condition JobIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the jobs's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	JobId  pulumi.StringOutput `pulumi:"jobId"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewJobIAMMember registers a new resource with the given unique name, arguments, and options.
func NewJobIAMMember(ctx *pulumi.Context,
	name string, args *JobIAMMemberArgs, opts ...pulumi.ResourceOption) (*JobIAMMember, error) {
	if args == nil || args.JobId == nil {
		return nil, errors.New("missing required argument 'JobId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &JobIAMMemberArgs{}
	}
	var resource JobIAMMember
	err := ctx.RegisterResource("gcp:dataproc/jobIAMMember:JobIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobIAMMember gets an existing JobIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobIAMMemberState, opts ...pulumi.ResourceOption) (*JobIAMMember, error) {
	var resource JobIAMMember
	err := ctx.ReadResource("gcp:dataproc/jobIAMMember:JobIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobIAMMember resources.
type jobIAMMemberState struct {
	Condition *JobIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the jobs's IAM policy.
	Etag   *string `pulumi:"etag"`
	JobId  *string `pulumi:"jobId"`
	Member *string `pulumi:"member"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type JobIAMMemberState struct {
	Condition JobIAMMemberConditionPtrInput
	// (Computed) The etag of the jobs's IAM policy.
	Etag   pulumi.StringPtrInput
	JobId  pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (JobIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMMemberState)(nil)).Elem()
}

type jobIAMMemberArgs struct {
	Condition *JobIAMMemberCondition `pulumi:"condition"`
	JobId     string                 `pulumi:"jobId"`
	Member    string                 `pulumi:"member"`
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project *string `pulumi:"project"`
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a JobIAMMember resource.
type JobIAMMemberArgs struct {
	Condition JobIAMMemberConditionPtrInput
	JobId     pulumi.StringInput
	Member    pulumi.StringInput
	// The project in which the job belongs. If it
	// is not provided, the provider will use a default.
	Project pulumi.StringPtrInput
	// The region in which the job belongs. If it
	// is not provided, the provider will use a default.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (JobIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobIAMMemberArgs)(nil)).Elem()
}

type JobIAMMemberInput interface {
	pulumi.Input

	ToJobIAMMemberOutput() JobIAMMemberOutput
	ToJobIAMMemberOutputWithContext(ctx context.Context) JobIAMMemberOutput
}

func (JobIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*JobIAMMember)(nil)).Elem()
}

func (i JobIAMMember) ToJobIAMMemberOutput() JobIAMMemberOutput {
	return i.ToJobIAMMemberOutputWithContext(context.Background())
}

func (i JobIAMMember) ToJobIAMMemberOutputWithContext(ctx context.Context) JobIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobIAMMemberOutput)
}

type JobIAMMemberOutput struct {
	*pulumi.OutputState
}

func (JobIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobIAMMemberOutput)(nil)).Elem()
}

func (o JobIAMMemberOutput) ToJobIAMMemberOutput() JobIAMMemberOutput {
	return o
}

func (o JobIAMMemberOutput) ToJobIAMMemberOutputWithContext(ctx context.Context) JobIAMMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobIAMMemberOutput{})
}
