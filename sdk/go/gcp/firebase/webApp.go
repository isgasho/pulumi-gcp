// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Google Cloud Firebase web application instance
//
// To get more information about WebApp, see:
//
// * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
// * How-to Guides
//     * [Official Documentation](https://firebase.google.com/)
//
// ## Example Usage
//
// ## Import
//
// WebApp can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:firebase/webApp:WebApp default {{name}}
// ```
type WebApp struct {
	pulumi.CustomResourceState

	// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The user-assigned display name of the App.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebApp registers a new resource with the given unique name, arguments, and options.
func NewWebApp(ctx *pulumi.Context,
	name string, args *WebAppArgs, opts ...pulumi.ResourceOption) (*WebApp, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &WebAppArgs{}
	}
	var resource WebApp
	err := ctx.RegisterResource("gcp:firebase/webApp:WebApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebApp gets an existing WebApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppState, opts ...pulumi.ResourceOption) (*WebApp, error) {
	var resource WebApp
	err := ctx.ReadResource("gcp:firebase/webApp:WebApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebApp resources.
type webAppState struct {
	// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId *string `pulumi:"appId"`
	// The user-assigned display name of the App.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebAppState struct {
	// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
	// token, as the data format is not specified.
	AppId pulumi.StringPtrInput
	// The user-assigned display name of the App.
	DisplayName pulumi.StringPtrInput
	// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppState)(nil)).Elem()
}

type webAppArgs struct {
	// The user-assigned display name of the App.
	DisplayName string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebApp resource.
type WebAppArgs struct {
	// The user-assigned display name of the App.
	DisplayName pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppArgs)(nil)).Elem()
}

type WebAppInput interface {
	pulumi.Input

	ToWebAppOutput() WebAppOutput
	ToWebAppOutputWithContext(ctx context.Context) WebAppOutput
}

func (WebApp) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApp)(nil)).Elem()
}

func (i WebApp) ToWebAppOutput() WebAppOutput {
	return i.ToWebAppOutputWithContext(context.Background())
}

func (i WebApp) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppOutput)
}

type WebAppOutput struct {
	*pulumi.OutputState
}

func (WebAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppOutput)(nil)).Elem()
}

func (o WebAppOutput) ToWebAppOutput() WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppOutput{})
}
