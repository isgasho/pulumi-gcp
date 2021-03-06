// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows configuring a single access level condition to be appended to an access level's conditions.
// This resource is intended to be used in cases where it is not possible to compile a full list
// of conditions to include in a `accesscontextmanager.AccessLevel` resource,
// to enable them to be added separately.
//
// > **Note:** If this resource is used alongside a `accesscontextmanager.AccessLevel` resource,
// the access level resource must have a `lifecycle` block with `ignoreChanges = [basic[0].conditions]` so
// they don't fight over which service accounts should be included.
//
// To get more information about AccessLevelCondition, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
// * How-to Guides
//     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
type AccessLevelsCondition struct {
	pulumi.CustomResourceState

	// The name of the Access Level to add this condition to.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy AccessLevelsConditionDevicePolicyPtrOutput `pulumi:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks pulumi.StringArrayOutput `pulumi:"ipSubnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate pulumi.BoolPtrOutput `pulumi:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels pulumi.StringArrayOutput `pulumi:"requiredAccessLevels"`
}

// NewAccessLevelsCondition registers a new resource with the given unique name, arguments, and options.
func NewAccessLevelsCondition(ctx *pulumi.Context,
	name string, args *AccessLevelsConditionArgs, opts ...pulumi.ResourceOption) (*AccessLevelsCondition, error) {
	if args == nil || args.AccessLevel == nil {
		return nil, errors.New("missing required argument 'AccessLevel'")
	}
	if args == nil {
		args = &AccessLevelsConditionArgs{}
	}
	var resource AccessLevelsCondition
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessLevelsCondition:AccessLevelsCondition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLevelsCondition gets an existing AccessLevelsCondition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLevelsCondition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLevelsConditionState, opts ...pulumi.ResourceOption) (*AccessLevelsCondition, error) {
	var resource AccessLevelsCondition
	err := ctx.ReadResource("gcp:accesscontextmanager/accessLevelsCondition:AccessLevelsCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLevelsCondition resources.
type accessLevelsConditionState struct {
	// The name of the Access Level to add this condition to.
	AccessLevel *string `pulumi:"accessLevel"`
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy *AccessLevelsConditionDevicePolicy `pulumi:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks []string `pulumi:"ipSubnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members []string `pulumi:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `pulumi:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []string `pulumi:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []string `pulumi:"requiredAccessLevels"`
}

type AccessLevelsConditionState struct {
	// The name of the Access Level to add this condition to.
	AccessLevel pulumi.StringPtrInput
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy AccessLevelsConditionDevicePolicyPtrInput
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks pulumi.StringArrayInput
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members pulumi.StringArrayInput
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate pulumi.BoolPtrInput
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions pulumi.StringArrayInput
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels pulumi.StringArrayInput
}

func (AccessLevelsConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelsConditionState)(nil)).Elem()
}

type accessLevelsConditionArgs struct {
	// The name of the Access Level to add this condition to.
	AccessLevel string `pulumi:"accessLevel"`
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy *AccessLevelsConditionDevicePolicy `pulumi:"devicePolicy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks []string `pulumi:"ipSubnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members []string `pulumi:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `pulumi:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []string `pulumi:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []string `pulumi:"requiredAccessLevels"`
}

// The set of arguments for constructing a AccessLevelsCondition resource.
type AccessLevelsConditionArgs struct {
	// The name of the Access Level to add this condition to.
	AccessLevel pulumi.StringInput
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy AccessLevelsConditionDevicePolicyPtrInput
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IpSubnetworks pulumi.StringArrayInput
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: `user:{emailid}`, `serviceAccount:{emailid}`
	Members pulumi.StringArrayInput
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate pulumi.BoolPtrInput
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions pulumi.StringArrayInput
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels pulumi.StringArrayInput
}

func (AccessLevelsConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelsConditionArgs)(nil)).Elem()
}
