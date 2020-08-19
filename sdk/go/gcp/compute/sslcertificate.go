// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An SslCertificate resource, used for HTTPS load balancing. This resource
// provides a mechanism to upload an SSL key and certificate to
// the load balancer to serve secure connections from the user.
//
//
// To get more information about SslCertificate, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
//
// > **Warning:** All arguments including `certificate` and `privateKey` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
type SSLCertificate struct {
	pulumi.CustomResourceState

	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The unique identifier for the resource.
	CertificateId pulumi.IntOutput `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewSSLCertificate registers a new resource with the given unique name, arguments, and options.
func NewSSLCertificate(ctx *pulumi.Context,
	name string, args *SSLCertificateArgs, opts ...pulumi.ResourceOption) (*SSLCertificate, error) {
	if args == nil || args.Certificate == nil {
		return nil, errors.New("missing required argument 'Certificate'")
	}
	if args == nil || args.PrivateKey == nil {
		return nil, errors.New("missing required argument 'PrivateKey'")
	}
	if args == nil {
		args = &SSLCertificateArgs{}
	}
	var resource SSLCertificate
	err := ctx.RegisterResource("gcp:compute/sSLCertificate:SSLCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSSLCertificate gets an existing SSLCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSSLCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SSLCertificateState, opts ...pulumi.ResourceOption) (*SSLCertificate, error) {
	var resource SSLCertificate
	err := ctx.ReadResource("gcp:compute/sSLCertificate:SSLCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SSLCertificate resources.
type sslcertificateState struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate *string `pulumi:"certificate"`
	// The unique identifier for the resource.
	CertificateId *int `pulumi:"certificateId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey *string `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type SSLCertificateState struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate pulumi.StringPtrInput
	// The unique identifier for the resource.
	CertificateId pulumi.IntPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (SSLCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslcertificateState)(nil)).Elem()
}

type sslcertificateArgs struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate string `pulumi:"certificate"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey string `pulumi:"privateKey"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a SSLCertificate resource.
type SSLCertificateArgs struct {
	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Certificate pulumi.StringInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The write-only private key in PEM format.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	PrivateKey pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (SSLCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslcertificateArgs)(nil)).Elem()
}
