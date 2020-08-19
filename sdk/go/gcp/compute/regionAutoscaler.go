// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an Autoscaler resource.
//
// Autoscalers allow you to automatically scale virtual machine instances in
// managed instance groups according to an autoscaling policy that you
// define.
//
//
// To get more information about RegionAutoscaler, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
// * How-to Guides
//     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
type RegionAutoscaler struct {
	pulumi.CustomResourceState

	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy RegionAutoscalerAutoscalingPolicyOutput `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// URL of the region where the instance group resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewRegionAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewRegionAutoscaler(ctx *pulumi.Context,
	name string, args *RegionAutoscalerArgs, opts ...pulumi.ResourceOption) (*RegionAutoscaler, error) {
	if args == nil || args.AutoscalingPolicy == nil {
		return nil, errors.New("missing required argument 'AutoscalingPolicy'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	if args == nil {
		args = &RegionAutoscalerArgs{}
	}
	var resource RegionAutoscaler
	err := ctx.RegisterResource("gcp:compute/regionAutoscaler:RegionAutoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionAutoscaler gets an existing RegionAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionAutoscalerState, opts ...pulumi.ResourceOption) (*RegionAutoscaler, error) {
	var resource RegionAutoscaler
	err := ctx.ReadResource("gcp:compute/regionAutoscaler:RegionAutoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionAutoscaler resources.
type regionAutoscalerState struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy *RegionAutoscalerAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// URL of the region where the instance group resides.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target *string `pulumi:"target"`
}

type RegionAutoscalerState struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy RegionAutoscalerAutoscalingPolicyPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// URL of the region where the instance group resides.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringPtrInput
}

func (RegionAutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionAutoscalerState)(nil)).Elem()
}

type regionAutoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy RegionAutoscalerAutoscalingPolicy `pulumi:"autoscalingPolicy"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// URL of the region where the instance group resides.
	Region *string `pulumi:"region"`
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a RegionAutoscaler resource.
type RegionAutoscalerArgs struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	// Structure is documented below.
	AutoscalingPolicy RegionAutoscalerAutoscalingPolicyInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The identifier (type) of the Stackdriver Monitoring metric.
	// The metric cannot have negative values.
	// The metric must have a value type of INT64 or DOUBLE.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// URL of the region where the instance group resides.
	Region pulumi.StringPtrInput
	// Fraction of backend capacity utilization (set in HTTP(s) load
	// balancing configuration) that autoscaler should maintain. Must
	// be a positive float value. If not defined, the default is 0.8.
	Target pulumi.StringInput
}

func (RegionAutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionAutoscalerArgs)(nil)).Elem()
}
