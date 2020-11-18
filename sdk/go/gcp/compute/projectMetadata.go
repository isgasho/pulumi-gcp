// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Authoritatively manages metadata common to all instances for a project in GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/storing-retrieving-metadata)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/projects/setCommonInstanceMetadata).
//
// > **Note:**  This resource manages all project-level metadata including project-level ssh keys.
// Keys unset in config but set on the server will be removed. If you want to manage only single
// key/value pairs within the project metadata rather than the entire set, then use
// google_compute_project_metadata_item.
//
// ## Example Usage
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
// 		_, err := compute.NewProjectMetadata(ctx, "_default", &compute.ProjectMetadataArgs{
// 			Metadata: pulumi.StringMap{
// 				"13":   pulumi.String("42"),
// 				"fizz": pulumi.String("buzz"),
// 				"foo":  pulumi.String("bar"),
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
// This resource can be imported using the project ID
//
// ```sh
//  $ pulumi import gcp:compute/projectMetadata:ProjectMetadata foo my-project-id`
// ```
type ProjectMetadata struct {
	pulumi.CustomResourceState

	// A series of key value pairs.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewProjectMetadata registers a new resource with the given unique name, arguments, and options.
func NewProjectMetadata(ctx *pulumi.Context,
	name string, args *ProjectMetadataArgs, opts ...pulumi.ResourceOption) (*ProjectMetadata, error) {
	if args == nil || args.Metadata == nil {
		return nil, errors.New("missing required argument 'Metadata'")
	}
	if args == nil {
		args = &ProjectMetadataArgs{}
	}
	var resource ProjectMetadata
	err := ctx.RegisterResource("gcp:compute/projectMetadata:ProjectMetadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMetadata gets an existing ProjectMetadata resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMetadataState, opts ...pulumi.ResourceOption) (*ProjectMetadata, error) {
	var resource ProjectMetadata
	err := ctx.ReadResource("gcp:compute/projectMetadata:ProjectMetadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMetadata resources.
type projectMetadataState struct {
	// A series of key value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ProjectMetadataState struct {
	// A series of key value pairs.
	Metadata pulumi.StringMapInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectMetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMetadataState)(nil)).Elem()
}

type projectMetadataArgs struct {
	// A series of key value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ProjectMetadata resource.
type ProjectMetadataArgs struct {
	// A series of key value pairs.
	Metadata pulumi.StringMapInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ProjectMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMetadataArgs)(nil)).Elem()
}

type ProjectMetadataInput interface {
	pulumi.Input

	ToProjectMetadataOutput() ProjectMetadataOutput
	ToProjectMetadataOutputWithContext(ctx context.Context) ProjectMetadataOutput
}

func (ProjectMetadata) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectMetadata)(nil)).Elem()
}

func (i ProjectMetadata) ToProjectMetadataOutput() ProjectMetadataOutput {
	return i.ToProjectMetadataOutputWithContext(context.Background())
}

func (i ProjectMetadata) ToProjectMetadataOutputWithContext(ctx context.Context) ProjectMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMetadataOutput)
}

type ProjectMetadataOutput struct {
	*pulumi.OutputState
}

func (ProjectMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectMetadataOutput)(nil)).Elem()
}

func (o ProjectMetadataOutput) ToProjectMetadataOutput() ProjectMetadataOutput {
	return o
}

func (o ProjectMetadataOutput) ToProjectMetadataOutputWithContext(ctx context.Context) ProjectMetadataOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectMetadataOutput{})
}
