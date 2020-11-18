// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a organization-level logging bucket config. For more information see
// [the official logging documentation](https://cloud.google.com/logging/docs/) and
// [Storing Logs](https://cloud.google.com/logging/docs/storage).
//
// > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "123456789"
// 		_default, err := organizations.GetOrganization(ctx, &organizations.GetOrganizationArgs{
// 			Organization: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewOrganizationBucketConfig(ctx, "basic", &logging.OrganizationBucketConfigArgs{
// 			Organization:  pulumi.String(_default.Organization),
// 			Location:      pulumi.String("global"),
// 			RetentionDays: pulumi.Int(30),
// 			BucketId:      pulumi.String("_Default"),
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
// This resource can be imported using the following format
//
// ```sh
//  $ pulumi import gcp:logging/organizationBucketConfig:OrganizationBucketConfig default organizations/{{organization}}/locations/{{location}}/buckets/{{bucket_id}}
// ```
type OrganizationBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewOrganizationBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewOrganizationBucketConfig(ctx *pulumi.Context,
	name string, args *OrganizationBucketConfigArgs, opts ...pulumi.ResourceOption) (*OrganizationBucketConfig, error) {
	if args == nil || args.BucketId == nil {
		return nil, errors.New("missing required argument 'BucketId'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Organization == nil {
		return nil, errors.New("missing required argument 'Organization'")
	}
	if args == nil {
		args = &OrganizationBucketConfigArgs{}
	}
	var resource OrganizationBucketConfig
	err := ctx.RegisterResource("gcp:logging/organizationBucketConfig:OrganizationBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationBucketConfig gets an existing OrganizationBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationBucketConfigState, opts ...pulumi.ResourceOption) (*OrganizationBucketConfig, error) {
	var resource OrganizationBucketConfig
	err := ctx.ReadResource("gcp:logging/organizationBucketConfig:OrganizationBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationBucketConfig resources.
type organizationBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Organization *string `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

type OrganizationBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketConfigState)(nil)).Elem()
}

type organizationBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location string `pulumi:"location"`
	// The parent resource that contains the logging bucket.
	Organization string `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a OrganizationBucketConfig resource.
type OrganizationBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringInput
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketConfigArgs)(nil)).Elem()
}

type OrganizationBucketConfigInput interface {
	pulumi.Input

	ToOrganizationBucketConfigOutput() OrganizationBucketConfigOutput
	ToOrganizationBucketConfigOutputWithContext(ctx context.Context) OrganizationBucketConfigOutput
}

func (OrganizationBucketConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationBucketConfig)(nil)).Elem()
}

func (i OrganizationBucketConfig) ToOrganizationBucketConfigOutput() OrganizationBucketConfigOutput {
	return i.ToOrganizationBucketConfigOutputWithContext(context.Background())
}

func (i OrganizationBucketConfig) ToOrganizationBucketConfigOutputWithContext(ctx context.Context) OrganizationBucketConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBucketConfigOutput)
}

type OrganizationBucketConfigOutput struct {
	*pulumi.OutputState
}

func (OrganizationBucketConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationBucketConfigOutput)(nil)).Elem()
}

func (o OrganizationBucketConfigOutput) ToOrganizationBucketConfigOutput() OrganizationBucketConfigOutput {
	return o
}

func (o OrganizationBucketConfigOutput) ToOrganizationBucketConfigOutputWithContext(ctx context.Context) OrganizationBucketConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationBucketConfigOutput{})
}
