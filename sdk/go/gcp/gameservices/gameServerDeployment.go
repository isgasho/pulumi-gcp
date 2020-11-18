// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gameservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A game server deployment resource.
//
// To get more information about GameServerDeployment, see:
//
// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/game-servers/docs)
//
// ## Example Usage
// ### Game Service Deployment Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/gameservices"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gameservices.NewGameServerDeployment(ctx, "_default", &gameservices.GameServerDeploymentArgs{
// 			DeploymentId: pulumi.String("tf-test-deployment"),
// 			Description:  pulumi.String("a deployment description"),
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
// GameServerDeployment can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerDeployment:GameServerDeployment default projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerDeployment:GameServerDeployment default {{project}}/{{location}}/{{deployment_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:gameservices/gameServerDeployment:GameServerDeployment default {{location}}/{{deployment_id}}
// ```
type GameServerDeployment struct {
	pulumi.CustomResourceState

	// A unique id for the deployment.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Human readable description of the game server deployment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the Deployment.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource id of the game server deployment, eg:
	// 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
	// 'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewGameServerDeployment registers a new resource with the given unique name, arguments, and options.
func NewGameServerDeployment(ctx *pulumi.Context,
	name string, args *GameServerDeploymentArgs, opts ...pulumi.ResourceOption) (*GameServerDeployment, error) {
	if args == nil || args.DeploymentId == nil {
		return nil, errors.New("missing required argument 'DeploymentId'")
	}
	if args == nil {
		args = &GameServerDeploymentArgs{}
	}
	var resource GameServerDeployment
	err := ctx.RegisterResource("gcp:gameservices/gameServerDeployment:GameServerDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerDeployment gets an existing GameServerDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerDeploymentState, opts ...pulumi.ResourceOption) (*GameServerDeployment, error) {
	var resource GameServerDeployment
	err := ctx.ReadResource("gcp:gameservices/gameServerDeployment:GameServerDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerDeployment resources.
type gameServerDeploymentState struct {
	// A unique id for the deployment.
	DeploymentId *string `pulumi:"deploymentId"`
	// Human readable description of the game server deployment.
	Description *string `pulumi:"description"`
	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Deployment.
	Location *string `pulumi:"location"`
	// The resource id of the game server deployment, eg:
	// 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
	// 'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type GameServerDeploymentState struct {
	// A unique id for the deployment.
	DeploymentId pulumi.StringPtrInput
	// Human readable description of the game server deployment.
	Description pulumi.StringPtrInput
	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Deployment.
	Location pulumi.StringPtrInput
	// The resource id of the game server deployment, eg:
	// 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
	// 'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GameServerDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerDeploymentState)(nil)).Elem()
}

type gameServerDeploymentArgs struct {
	// A unique id for the deployment.
	DeploymentId string `pulumi:"deploymentId"`
	// Human readable description of the game server deployment.
	Description *string `pulumi:"description"`
	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// Location of the Deployment.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a GameServerDeployment resource.
type GameServerDeploymentArgs struct {
	// A unique id for the deployment.
	DeploymentId pulumi.StringInput
	// Human readable description of the game server deployment.
	Description pulumi.StringPtrInput
	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	Labels pulumi.StringMapInput
	// Location of the Deployment.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GameServerDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerDeploymentArgs)(nil)).Elem()
}

type GameServerDeploymentInput interface {
	pulumi.Input

	ToGameServerDeploymentOutput() GameServerDeploymentOutput
	ToGameServerDeploymentOutputWithContext(ctx context.Context) GameServerDeploymentOutput
}

func (GameServerDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerDeployment)(nil)).Elem()
}

func (i GameServerDeployment) ToGameServerDeploymentOutput() GameServerDeploymentOutput {
	return i.ToGameServerDeploymentOutputWithContext(context.Background())
}

func (i GameServerDeployment) ToGameServerDeploymentOutputWithContext(ctx context.Context) GameServerDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerDeploymentOutput)
}

type GameServerDeploymentOutput struct {
	*pulumi.OutputState
}

func (GameServerDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerDeploymentOutput)(nil)).Elem()
}

func (o GameServerDeploymentOutput) ToGameServerDeploymentOutput() GameServerDeploymentOutput {
	return o
}

func (o GameServerDeploymentOutput) ToGameServerDeploymentOutputWithContext(ctx context.Context) GameServerDeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GameServerDeploymentOutput{})
}
