// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package osconfig

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Patch deployments are configurations that individual patch jobs use to complete a patch.
// These configurations include instance filter, package repository settings, and a schedule.
//
//
// To get more information about PatchDeployment, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/osconfig/rest)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/os-patch-management)
//
// ## Example Usage
//
// ### Os Config Patch Deployment Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/osconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		patch, err := osconfig.NewPatchDeployment(ctx, "patch", &osconfig.PatchDeploymentArgs{
// 			InstanceFilter: &osconfig.PatchDeploymentInstanceFilterArgs{
// 				All: pulumi.Bool(true),
// 			},
// 			PatchDeploymentId: pulumi.String("patch-deploy-inst"),
// 			RecurringSchedule: &osconfig.PatchDeploymentRecurringScheduleArgs{
// 				TimeOfDay: &osconfig.PatchDeploymentRecurringScheduleTimeOfDayArgs{
// 					Hours: pulumi.Int(1),
// 				},
// 				TimeZone: &osconfig.PatchDeploymentRecurringScheduleTimeZoneArgs{
// 					Id: pulumi.String("America/New_York"),
// 				},
// 				Weekly: &osconfig.PatchDeploymentRecurringScheduleWeeklyArgs{
// 					DayOfWeek: pulumi.String("MONDAY"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PatchDeployment struct {
	pulumi.CustomResourceState

	// Time the patch deployment was created. Timestamp is in RFC3339 text format. A timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilterOutput `pulumi:"instanceFilter"`
	// -
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	LastExecuteTime pulumi.StringOutput `pulumi:"lastExecuteTime"`
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form:
	// projects/{project_id}/patchDeployments/{patchDeploymentId}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule PatchDeploymentOneTimeSchedulePtrOutput `pulumi:"oneTimeSchedule"`
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig PatchDeploymentPatchConfigPtrOutput `pulumi:"patchConfig"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId pulumi.StringOutput `pulumi:"patchDeploymentId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule PatchDeploymentRecurringSchedulePtrOutput `pulumi:"recurringSchedule"`
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format. A timestamp in RFC3339 UTC "Zulu"
	// format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewPatchDeployment registers a new resource with the given unique name, arguments, and options.
func NewPatchDeployment(ctx *pulumi.Context,
	name string, args *PatchDeploymentArgs, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	if args == nil || args.InstanceFilter == nil {
		return nil, errors.New("missing required argument 'InstanceFilter'")
	}
	if args == nil || args.PatchDeploymentId == nil {
		return nil, errors.New("missing required argument 'PatchDeploymentId'")
	}
	if args == nil {
		args = &PatchDeploymentArgs{}
	}
	var resource PatchDeployment
	err := ctx.RegisterResource("gcp:osconfig/patchDeployment:PatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchDeployment gets an existing PatchDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchDeploymentState, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	var resource PatchDeployment
	err := ctx.ReadResource("gcp:osconfig/patchDeployment:PatchDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchDeployment resources.
type patchDeploymentState struct {
	// Time the patch deployment was created. Timestamp is in RFC3339 text format. A timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration *string `pulumi:"duration"`
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter *PatchDeploymentInstanceFilter `pulumi:"instanceFilter"`
	// -
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	LastExecuteTime *string `pulumi:"lastExecuteTime"`
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form:
	// projects/{project_id}/patchDeployments/{patchDeploymentId}.
	Name *string `pulumi:"name"`
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule *PatchDeploymentOneTimeSchedule `pulumi:"oneTimeSchedule"`
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig *PatchDeploymentPatchConfig `pulumi:"patchConfig"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId *string `pulumi:"patchDeploymentId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule *PatchDeploymentRecurringSchedule `pulumi:"recurringSchedule"`
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format. A timestamp in RFC3339 UTC "Zulu"
	// format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type PatchDeploymentState struct {
	// Time the patch deployment was created. Timestamp is in RFC3339 text format. A timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration pulumi.StringPtrInput
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilterPtrInput
	// -
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	LastExecuteTime pulumi.StringPtrInput
	// Unique name for the patch deployment resource in a project. The patch deployment name is in the form:
	// projects/{project_id}/patchDeployments/{patchDeploymentId}.
	Name pulumi.StringPtrInput
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule PatchDeploymentOneTimeSchedulePtrInput
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig PatchDeploymentPatchConfigPtrInput
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule PatchDeploymentRecurringSchedulePtrInput
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format. A timestamp in RFC3339 UTC "Zulu"
	// format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (PatchDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchDeploymentState)(nil)).Elem()
}

type patchDeploymentArgs struct {
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration *string `pulumi:"duration"`
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilter `pulumi:"instanceFilter"`
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule *PatchDeploymentOneTimeSchedule `pulumi:"oneTimeSchedule"`
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig *PatchDeploymentPatchConfig `pulumi:"patchConfig"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId string `pulumi:"patchDeploymentId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule *PatchDeploymentRecurringSchedule `pulumi:"recurringSchedule"`
}

// The set of arguments for constructing a PatchDeployment resource.
type PatchDeploymentArgs struct {
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration pulumi.StringPtrInput
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilterInput
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule PatchDeploymentOneTimeSchedulePtrInput
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig PatchDeploymentPatchConfigPtrInput
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule PatchDeploymentRecurringSchedulePtrInput
}

func (PatchDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchDeploymentArgs)(nil)).Elem()
}
