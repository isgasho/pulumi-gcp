// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A tag template defines a tag, which can have one or more typed fields.
// The template is used to create and attach the tag to GCP resources.
//
// To get more information about TagTemplate, see:
//
// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.tagTemplates)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
//
// ## Example Usage
// ### Data Catalog Tag Template Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/datacatalog"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datacatalog.NewTagTemplate(ctx, "basicTagTemplate", &datacatalog.TagTemplateArgs{
// 			DisplayName: pulumi.String("Demo Tag Template"),
// 			Fields: datacatalog.TagTemplateFieldArray{
// 				&datacatalog.TagTemplateFieldArgs{
// 					DisplayName: pulumi.String("Source of data asset"),
// 					FieldId:     pulumi.String("source"),
// 					IsRequired:  pulumi.Bool(true),
// 					Type: &datacatalog.TagTemplateFieldTypeArgs{
// 						PrimitiveType: pulumi.String("STRING"),
// 					},
// 				},
// 				&datacatalog.TagTemplateFieldArgs{
// 					DisplayName: pulumi.String("Number of rows in the data asset"),
// 					FieldId:     pulumi.String("num_rows"),
// 					Type: &datacatalog.TagTemplateFieldTypeArgs{
// 						PrimitiveType: pulumi.String("DOUBLE"),
// 					},
// 				},
// 				&datacatalog.TagTemplateFieldArgs{
// 					DisplayName: pulumi.String("PII type"),
// 					FieldId:     pulumi.String("pii_type"),
// 					Type: &datacatalog.TagTemplateFieldTypeArgs{
// 						EnumType: &datacatalog.TagTemplateFieldTypeEnumTypeArgs{
// 							AllowedValues: datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArray{
// 								&datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs{
// 									DisplayName: pulumi.String("EMAIL"),
// 								},
// 								&datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs{
// 									DisplayName: pulumi.String("SOCIAL SECURITY NUMBER"),
// 								},
// 								&datacatalog.TagTemplateFieldTypeEnumTypeAllowedValueArgs{
// 									DisplayName: pulumi.String("NONE"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			ForceDelete:   pulumi.Bool(false),
// 			Region:        pulumi.String("us-central1"),
// 			TagTemplateId: pulumi.String("my_template"),
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
// TagTemplate can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:datacatalog/tagTemplate:TagTemplate default {{name}}
// ```
type TagTemplate struct {
	pulumi.CustomResourceState

	// The display name for this template.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
	// Structure is documented below.
	Fields TagTemplateFieldArrayOutput `pulumi:"fields"`
	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// -
	// The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Template location region.
	Region pulumi.StringOutput `pulumi:"region"`
	// The id of the tag template to create.
	TagTemplateId pulumi.StringOutput `pulumi:"tagTemplateId"`
}

// NewTagTemplate registers a new resource with the given unique name, arguments, and options.
func NewTagTemplate(ctx *pulumi.Context,
	name string, args *TagTemplateArgs, opts ...pulumi.ResourceOption) (*TagTemplate, error) {
	if args == nil || args.Fields == nil {
		return nil, errors.New("missing required argument 'Fields'")
	}
	if args == nil || args.TagTemplateId == nil {
		return nil, errors.New("missing required argument 'TagTemplateId'")
	}
	if args == nil {
		args = &TagTemplateArgs{}
	}
	var resource TagTemplate
	err := ctx.RegisterResource("gcp:datacatalog/tagTemplate:TagTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplate gets an existing TagTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateState, opts ...pulumi.ResourceOption) (*TagTemplate, error) {
	var resource TagTemplate
	err := ctx.ReadResource("gcp:datacatalog/tagTemplate:TagTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplate resources.
type tagTemplateState struct {
	// The display name for this template.
	DisplayName *string `pulumi:"displayName"`
	// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
	// Structure is documented below.
	Fields []TagTemplateField `pulumi:"fields"`
	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	ForceDelete *bool `pulumi:"forceDelete"`
	// -
	// The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Template location region.
	Region *string `pulumi:"region"`
	// The id of the tag template to create.
	TagTemplateId *string `pulumi:"tagTemplateId"`
}

type TagTemplateState struct {
	// The display name for this template.
	DisplayName pulumi.StringPtrInput
	// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
	// Structure is documented below.
	Fields TagTemplateFieldArrayInput
	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	ForceDelete pulumi.BoolPtrInput
	// -
	// The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Template location region.
	Region pulumi.StringPtrInput
	// The id of the tag template to create.
	TagTemplateId pulumi.StringPtrInput
}

func (TagTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateState)(nil)).Elem()
}

type tagTemplateArgs struct {
	// The display name for this template.
	DisplayName *string `pulumi:"displayName"`
	// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
	// Structure is documented below.
	Fields []TagTemplateField `pulumi:"fields"`
	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Template location region.
	Region *string `pulumi:"region"`
	// The id of the tag template to create.
	TagTemplateId string `pulumi:"tagTemplateId"`
}

// The set of arguments for constructing a TagTemplate resource.
type TagTemplateArgs struct {
	// The display name for this template.
	DisplayName pulumi.StringPtrInput
	// Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields.
	// Structure is documented below.
	Fields TagTemplateFieldArrayInput
	// This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.
	ForceDelete pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Template location region.
	Region pulumi.StringPtrInput
	// The id of the tag template to create.
	TagTemplateId pulumi.StringInput
}

func (TagTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateArgs)(nil)).Elem()
}

type TagTemplateInput interface {
	pulumi.Input

	ToTagTemplateOutput() TagTemplateOutput
	ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput
}

func (TagTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*TagTemplate)(nil)).Elem()
}

func (i TagTemplate) ToTagTemplateOutput() TagTemplateOutput {
	return i.ToTagTemplateOutputWithContext(context.Background())
}

func (i TagTemplate) ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateOutput)
}

type TagTemplateOutput struct {
	*pulumi.OutputState
}

func (TagTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagTemplateOutput)(nil)).Elem()
}

func (o TagTemplateOutput) ToTagTemplateOutput() TagTemplateOutput {
	return o
}

func (o TagTemplateOutput) ToTagTemplateOutputWithContext(ctx context.Context) TagTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagTemplateOutput{})
}
