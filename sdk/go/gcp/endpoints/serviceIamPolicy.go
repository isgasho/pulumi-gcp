// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package endpoints

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Management Service. Each of these resources serves a different use case:
// 
// * `endpoints.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `endpoints.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `endpoints.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
// 
// > **Note:** `endpoints.ServiceIamPolicy` **cannot** be used in conjunction with `endpoints.ServiceIamBinding` and `endpoints.ServiceIamMember` or they will fight over what your policy should be.
// 
// > **Note:** `endpoints.ServiceIamBinding` resources **can be** used in conjunction with `endpoints.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/endpoints_service_iam_policy.html.markdown.
type ServiceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceIamPolicy(ctx *pulumi.Context,
	name string, args *ServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ServiceIamPolicyArgs{}
	}
	var resource ServiceIamPolicy
	err := ctx.RegisterResource("gcp:endpoints/serviceIamPolicy:ServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceIamPolicy gets an existing ServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceIamPolicyState, opts ...pulumi.ResourceOption) (*ServiceIamPolicy, error) {
	var resource ServiceIamPolicy
	err := ctx.ReadResource("gcp:endpoints/serviceIamPolicy:ServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceIamPolicy resources.
type serviceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName *string `pulumi:"serviceName"`
}

type ServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringPtrInput
}

func (ServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamPolicyState)(nil)).Elem()
}

type serviceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServiceIamPolicy resource.
type ServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The name of the service. Used to find the parent resource to bind the IAM policy to
	ServiceName pulumi.StringInput
}

func (ServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceIamPolicyArgs)(nil)).Elem()
}

