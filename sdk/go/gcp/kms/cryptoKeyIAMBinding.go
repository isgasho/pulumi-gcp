// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:
//
// * `kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
// * `kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
// * `kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.
//
// > **Note:** `kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `kms.CryptoKeyIAMBinding` and `kms.CryptoKeyIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		keyring, err := kms.NewKeyRing(ctx, "keyring", &kms.KeyRingArgs{
// 			Location: pulumi.String("global"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		key, err := kms.NewCryptoKey(ctx, "key", &kms.CryptoKeyArgs{
// 			KeyRing:        keyring.ID(),
// 			RotationPeriod: pulumi.String("100000s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/cloudkms.cryptoKeyEncrypter",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCryptoKeyIAMPolicy(ctx, "cryptoKey", &kms.CryptoKeyIAMPolicyArgs{
// 			CryptoKeyId: key.ID(),
// 			PolicyData:  pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 						Title:       "expires_after_2019_12_31",
// 					},
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Role: "roles/cloudkms.cryptoKeyEncrypter",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKey", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKey", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &kms.CryptoKeyIAMBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 			Condition: &kms.CryptoKeyIAMMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; first the resource in question and then the role.
//
// These bindings can be imported using the `crypto_key_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `crypto_key_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding crypto_key your-project-id/location-name/key-ring-name/key-name
// ```
type CryptoKeyIAMBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMBindingConditionPtrOutput `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringOutput `pulumi:"cryptoKeyId"`
	// (Computed) The etag of the project's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCryptoKeyIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewCryptoKeyIAMBinding(ctx *pulumi.Context,
	name string, args *CryptoKeyIAMBindingArgs, opts ...pulumi.ResourceOption) (*CryptoKeyIAMBinding, error) {
	if args == nil || args.CryptoKeyId == nil {
		return nil, errors.New("missing required argument 'CryptoKeyId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &CryptoKeyIAMBindingArgs{}
	}
	var resource CryptoKeyIAMBinding
	err := ctx.RegisterResource("gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCryptoKeyIAMBinding gets an existing CryptoKeyIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKeyIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CryptoKeyIAMBindingState, opts ...pulumi.ResourceOption) (*CryptoKeyIAMBinding, error) {
	var resource CryptoKeyIAMBinding
	err := ctx.ReadResource("gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CryptoKeyIAMBinding resources.
type cryptoKeyIAMBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CryptoKeyIAMBindingCondition `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId *string `pulumi:"cryptoKeyId"`
	// (Computed) The etag of the project's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type CryptoKeyIAMBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMBindingConditionPtrInput
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringPtrInput
	// (Computed) The etag of the project's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (CryptoKeyIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyIAMBindingState)(nil)).Elem()
}

type cryptoKeyIAMBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CryptoKeyIAMBindingCondition `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId string   `pulumi:"cryptoKeyId"`
	Members     []string `pulumi:"members"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CryptoKeyIAMBinding resource.
type CryptoKeyIAMBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMBindingConditionPtrInput
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringInput
	Members     pulumi.StringArrayInput
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (CryptoKeyIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyIAMBindingArgs)(nil)).Elem()
}

type CryptoKeyIAMBindingInput interface {
	pulumi.Input

	ToCryptoKeyIAMBindingOutput() CryptoKeyIAMBindingOutput
	ToCryptoKeyIAMBindingOutputWithContext(ctx context.Context) CryptoKeyIAMBindingOutput
}

func (CryptoKeyIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyIAMBinding)(nil)).Elem()
}

func (i CryptoKeyIAMBinding) ToCryptoKeyIAMBindingOutput() CryptoKeyIAMBindingOutput {
	return i.ToCryptoKeyIAMBindingOutputWithContext(context.Background())
}

func (i CryptoKeyIAMBinding) ToCryptoKeyIAMBindingOutputWithContext(ctx context.Context) CryptoKeyIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMBindingOutput)
}

type CryptoKeyIAMBindingOutput struct {
	*pulumi.OutputState
}

func (CryptoKeyIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyIAMBindingOutput)(nil)).Elem()
}

func (o CryptoKeyIAMBindingOutput) ToCryptoKeyIAMBindingOutput() CryptoKeyIAMBindingOutput {
	return o
}

func (o CryptoKeyIAMBindingOutput) ToCryptoKeyIAMBindingOutputWithContext(ctx context.Context) CryptoKeyIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CryptoKeyIAMBindingOutput{})
}
