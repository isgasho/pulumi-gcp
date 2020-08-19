// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a SSL policy. SSL policies give you the ability to control the
// features of SSL that your SSL proxy or HTTPS load balancer negotiates.
//
//
// To get more information about SslPolicy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslPolicies)
// * How-to Guides
//     * [Using SSL Policies](https://cloud.google.com/compute/docs/load-balancing/ssl-policies)
type SSLPolicy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayOutput `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayOutput `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion pulumi.StringPtrOutput `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile pulumi.StringPtrOutput `pulumi:"profile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewSSLPolicy registers a new resource with the given unique name, arguments, and options.
func NewSSLPolicy(ctx *pulumi.Context,
	name string, args *SSLPolicyArgs, opts ...pulumi.ResourceOption) (*SSLPolicy, error) {
	if args == nil {
		args = &SSLPolicyArgs{}
	}
	var resource SSLPolicy
	err := ctx.RegisterResource("gcp:compute/sSLPolicy:SSLPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSLPolicy gets an existing SSLPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSLPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSLPolicyState, opts ...pulumi.ResourceOption) (*SSLPolicy, error) {
	var resource SSLPolicy
	err := ctx.ReadResource("gcp:compute/sSLPolicy:SSLPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSLPolicy resources.
type sslpolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile *string `pulumi:"profile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type SSLPolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (SSLPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslpolicyState)(nil)).Elem()
}

type sslpolicyArgs struct {
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile *string `pulumi:"profile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a SSLPolicy resource.
type SSLPolicyArgs struct {
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the `CUSTOM` profile. This argument
	// *must not* be present when using any other profile.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer.
	// Default value is `TLS_1_0`.
	// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using `CUSTOM`,
	// the set of SSL features to enable must be specified in the
	// `customFeatures` field.
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// `CUSTOM` is used, the `customFeatures` attribute **must be set**.
	// Default value is `COMPATIBLE`.
	// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
	Profile pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (SSLPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslpolicyArgs)(nil)).Elem()
}
