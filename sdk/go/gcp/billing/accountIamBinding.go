// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AccountIamBinding struct {
	pulumi.CustomResourceState

	BillingAccountId pulumi.StringOutput                 `pulumi:"billingAccountId"`
	Condition        AccountIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag             pulumi.StringOutput                 `pulumi:"etag"`
	Members          pulumi.StringArrayOutput            `pulumi:"members"`
	Role             pulumi.StringOutput                 `pulumi:"role"`
}

// NewAccountIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAccountIamBinding(ctx *pulumi.Context,
	name string, args *AccountIamBindingArgs, opts ...pulumi.ResourceOption) (*AccountIamBinding, error) {
	if args == nil || args.BillingAccountId == nil {
		return nil, errors.New("missing required argument 'BillingAccountId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &AccountIamBindingArgs{}
	}
	var resource AccountIamBinding
	err := ctx.RegisterResource("gcp:billing/accountIamBinding:AccountIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamBinding gets an existing AccountIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamBindingState, opts ...pulumi.ResourceOption) (*AccountIamBinding, error) {
	var resource AccountIamBinding
	err := ctx.ReadResource("gcp:billing/accountIamBinding:AccountIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamBinding resources.
type accountIamBindingState struct {
	BillingAccountId *string                     `pulumi:"billingAccountId"`
	Condition        *AccountIamBindingCondition `pulumi:"condition"`
	Etag             *string                     `pulumi:"etag"`
	Members          []string                    `pulumi:"members"`
	Role             *string                     `pulumi:"role"`
}

type AccountIamBindingState struct {
	BillingAccountId pulumi.StringPtrInput
	Condition        AccountIamBindingConditionPtrInput
	Etag             pulumi.StringPtrInput
	Members          pulumi.StringArrayInput
	Role             pulumi.StringPtrInput
}

func (AccountIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamBindingState)(nil)).Elem()
}

type accountIamBindingArgs struct {
	BillingAccountId string                      `pulumi:"billingAccountId"`
	Condition        *AccountIamBindingCondition `pulumi:"condition"`
	Members          []string                    `pulumi:"members"`
	Role             string                      `pulumi:"role"`
}

// The set of arguments for constructing a AccountIamBinding resource.
type AccountIamBindingArgs struct {
	BillingAccountId pulumi.StringInput
	Condition        AccountIamBindingConditionPtrInput
	Members          pulumi.StringArrayInput
	Role             pulumi.StringInput
}

func (AccountIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamBindingArgs)(nil)).Elem()
}

type AccountIamBindingInput interface {
	pulumi.Input

	ToAccountIamBindingOutput() AccountIamBindingOutput
	ToAccountIamBindingOutputWithContext(ctx context.Context) AccountIamBindingOutput
}

func (AccountIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIamBinding)(nil)).Elem()
}

func (i AccountIamBinding) ToAccountIamBindingOutput() AccountIamBindingOutput {
	return i.ToAccountIamBindingOutputWithContext(context.Background())
}

func (i AccountIamBinding) ToAccountIamBindingOutputWithContext(ctx context.Context) AccountIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamBindingOutput)
}

type AccountIamBindingOutput struct {
	*pulumi.OutputState
}

func (AccountIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountIamBindingOutput)(nil)).Elem()
}

func (o AccountIamBindingOutput) ToAccountIamBindingOutput() AccountIamBindingOutput {
	return o
}

func (o AccountIamBindingOutput) ToAccountIamBindingOutputWithContext(ctx context.Context) AccountIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountIamBindingOutput{})
}
