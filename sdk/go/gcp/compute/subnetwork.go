// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A VPC network is a virtual version of the traditional physical networks
// that exist within and between physical data centers. A VPC network
// provides connectivity for your Compute Engine virtual machine (VM)
// instances, Container Engine containers, App Engine Flex services, and
// other network-related resources.
//
// Each GCP project contains one or more VPC networks. Each VPC network is a
// global entity spanning all GCP regions. This global VPC network allows VM
// instances and other resources to communicate with each other via internal,
// private IP addresses.
//
// Each VPC network is subdivided into subnets, and each subnet is contained
// within a single region. You can have more than one subnet in a region for
// a given VPC network. Each subnet has a contiguous private RFC1918 IP
// space. You create instances, containers, and the like in these subnets.
// When you create an instance, you must create it in a subnet, and the
// instance draws its internal IP address from that subnet.
//
// Virtual machine (VM) instances in a VPC network can communicate with
// instances in all other subnets of the same VPC network, regardless of
// region, using their RFC1918 private IP addresses. You can isolate portions
// of the network, even entire subnets, using firewall rules.
//
//
// To get more information about Subnetwork, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/subnetworks)
// * How-to Guides
//     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
//     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
type Subnetwork struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: This field is not useful for users, and has been removed as an output.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringOutput `pulumi:"gatewayAddress"`
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
	LogConfig SubnetworkLogConfigPtrOutput `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringOutput `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrOutput `pulumi:"privateIpGoogleAccess"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// URL of the GCP region for this subnetwork.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges SubnetworkSecondaryIpRangeArrayOutput `pulumi:"secondaryIpRanges"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewSubnetwork registers a new resource with the given unique name, arguments, and options.
func NewSubnetwork(ctx *pulumi.Context,
	name string, args *SubnetworkArgs, opts ...pulumi.ResourceOption) (*Subnetwork, error) {
	if args == nil || args.IpCidrRange == nil {
		return nil, errors.New("missing required argument 'IpCidrRange'")
	}
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	if args == nil {
		args = &SubnetworkArgs{}
	}
	var resource Subnetwork
	err := ctx.RegisterResource("gcp:compute/subnetwork:Subnetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetwork gets an existing Subnetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkState, opts ...pulumi.ResourceOption) (*Subnetwork, error) {
	var resource Subnetwork
	err := ctx.ReadResource("gcp:compute/subnetwork:Subnetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnetwork resources.
type subnetworkState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: This field is not useful for users, and has been removed as an output.
	Fingerprint *string `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress *string `pulumi:"gatewayAddress"`
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
	LogConfig *SubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network *string `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	Purpose *string `pulumi:"purpose"`
	// URL of the GCP region for this subnetwork.
	Region *string `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	Role *string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges []SubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type SubnetworkState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: This field is not useful for users, and has been removed as an output.
	Fingerprint pulumi.StringPtrInput
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringPtrInput
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange pulumi.StringPtrInput
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
	LogConfig SubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringPtrInput
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	Purpose pulumi.StringPtrInput
	// URL of the GCP region for this subnetwork.
	Region pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	Role pulumi.StringPtrInput
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges SubnetworkSecondaryIpRangeArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (SubnetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkState)(nil)).Elem()
}

type subnetworkArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description *string `pulumi:"description"`
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
	LogConfig *SubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network string `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	Purpose *string `pulumi:"purpose"`
	// URL of the GCP region for this subnetwork.
	Region *string `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	Role *string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges []SubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
}

// The set of arguments for constructing a Subnetwork resource.
type SubnetworkArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description pulumi.StringPtrInput
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange pulumi.StringInput
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
	LogConfig SubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringInput
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	Purpose pulumi.StringPtrInput
	// URL of the GCP region for this subnetwork.
	Region pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	Role pulumi.StringPtrInput
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges SubnetworkSecondaryIpRangeArrayInput
}

func (SubnetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkArgs)(nil)).Elem()
}
