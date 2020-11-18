// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a folder-level logging exclusion. For more information see
// [the official documentation](https://cloud.google.com/logging/docs/) and
// [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
//
// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
// granted to the credentials used with this provider.
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
// 		_, err := organizations.NewFolder(ctx, "my_folder", &organizations.FolderArgs{
// 			DisplayName: pulumi.String("My folder"),
// 			Parent:      pulumi.String("organizations/123456"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewFolderExclusion(ctx, "my_exclusion", &logging.FolderExclusionArgs{
// 			Folder:      my_folder.Name,
// 			Description: pulumi.String("Exclude GCE instance debug logs"),
// 			Filter:      pulumi.String("resource.type = gce_instance AND severity <= DEBUG"),
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
// Folder-level logging exclusions can be imported using their URI, e.g.
//
// ```sh
//  $ pulumi import gcp:logging/folderExclusion:FolderExclusion my_exclusion folders/my-folder/exclusions/my-exclusion
// ```
type FolderExclusion struct {
	pulumi.CustomResourceState

	// A human-readable description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// The name of the logging exclusion.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFolderExclusion registers a new resource with the given unique name, arguments, and options.
func NewFolderExclusion(ctx *pulumi.Context,
	name string, args *FolderExclusionArgs, opts ...pulumi.ResourceOption) (*FolderExclusion, error) {
	if args == nil || args.Filter == nil {
		return nil, errors.New("missing required argument 'Filter'")
	}
	if args == nil || args.Folder == nil {
		return nil, errors.New("missing required argument 'Folder'")
	}
	if args == nil {
		args = &FolderExclusionArgs{}
	}
	var resource FolderExclusion
	err := ctx.RegisterResource("gcp:logging/folderExclusion:FolderExclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderExclusion gets an existing FolderExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderExclusionState, opts ...pulumi.ResourceOption) (*FolderExclusion, error) {
	var resource FolderExclusion
	err := ctx.ReadResource("gcp:logging/folderExclusion:FolderExclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderExclusion resources.
type folderExclusionState struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder *string `pulumi:"folder"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
}

type FolderExclusionState struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
}

func (FolderExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderExclusionState)(nil)).Elem()
}

type folderExclusionArgs struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter string `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder string `pulumi:"folder"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FolderExclusion resource.
type FolderExclusionArgs struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringInput
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder pulumi.StringInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
}

func (FolderExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderExclusionArgs)(nil)).Elem()
}

type FolderExclusionInput interface {
	pulumi.Input

	ToFolderExclusionOutput() FolderExclusionOutput
	ToFolderExclusionOutputWithContext(ctx context.Context) FolderExclusionOutput
}

func (FolderExclusion) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderExclusion)(nil)).Elem()
}

func (i FolderExclusion) ToFolderExclusionOutput() FolderExclusionOutput {
	return i.ToFolderExclusionOutputWithContext(context.Background())
}

func (i FolderExclusion) ToFolderExclusionOutputWithContext(ctx context.Context) FolderExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderExclusionOutput)
}

type FolderExclusionOutput struct {
	*pulumi.OutputState
}

func (FolderExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FolderExclusionOutput)(nil)).Elem()
}

func (o FolderExclusionOutput) ToFolderExclusionOutput() FolderExclusionOutput {
	return o
}

func (o FolderExclusionOutput) ToFolderExclusionOutputWithContext(ctx context.Context) FolderExclusionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FolderExclusionOutput{})
}
