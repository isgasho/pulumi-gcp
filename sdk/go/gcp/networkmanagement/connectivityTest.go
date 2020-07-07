// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A connectivity test are a static analysis of your resource configurations
// that enables you to evaluate connectivity to and from Google Cloud
// resources in your Virtual Private Cloud (VPC) network.
//
// To get more information about ConnectivityTest, see:
//
// * [API documentation](https://cloud.google.com/network-intelligence-center/docs/connectivity-tests/reference/networkmanagement/rest/v1/projects.locations.global.connectivityTests)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/network-intelligence-center/docs)
//
// ## Example Usage
type ConnectivityTest struct {
	pulumi.CustomResourceState

	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.  Structure is documented below.
	Destination ConnectivityTestDestinationOutput `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Unique name for the connectivity test.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects pulumi.StringArrayOutput `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.  Structure is documented below.
	Source ConnectivityTestSourceOutput `pulumi:"source"`
}

// NewConnectivityTest registers a new resource with the given unique name, arguments, and options.
func NewConnectivityTest(ctx *pulumi.Context,
	name string, args *ConnectivityTestArgs, opts ...pulumi.ResourceOption) (*ConnectivityTest, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	if args == nil {
		args = &ConnectivityTestArgs{}
	}
	var resource ConnectivityTest
	err := ctx.RegisterResource("gcp:networkmanagement/connectivityTest:ConnectivityTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectivityTest gets an existing ConnectivityTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectivityTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectivityTestState, opts ...pulumi.ResourceOption) (*ConnectivityTest, error) {
	var resource ConnectivityTest
	err := ctx.ReadResource("gcp:networkmanagement/connectivityTest:ConnectivityTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectivityTest resources.
type connectivityTestState struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.  Structure is documented below.
	Destination *ConnectivityTestDestination `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique name for the connectivity test.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.  Structure is documented below.
	Source *ConnectivityTestSource `pulumi:"source"`
}

type ConnectivityTestState struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.  Structure is documented below.
	Destination ConnectivityTestDestinationPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Unique name for the connectivity test.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrInput
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects pulumi.StringArrayInput
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.  Structure is documented below.
	Source ConnectivityTestSourcePtrInput
}

func (ConnectivityTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityTestState)(nil)).Elem()
}

type connectivityTestArgs struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.  Structure is documented below.
	Destination ConnectivityTestDestination `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique name for the connectivity test.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.  Structure is documented below.
	Source ConnectivityTestSource `pulumi:"source"`
}

// The set of arguments for constructing a ConnectivityTest resource.
type ConnectivityTestArgs struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.  Structure is documented below.
	Destination ConnectivityTestDestinationInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Unique name for the connectivity test.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrInput
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects pulumi.StringArrayInput
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.  Structure is documented below.
	Source ConnectivityTestSourceInput
}

func (ConnectivityTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityTestArgs)(nil)).Elem()
}
