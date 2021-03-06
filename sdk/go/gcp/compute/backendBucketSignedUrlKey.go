// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A key for signing Cloud CDN signed URLs for BackendBuckets.
//
// To get more information about BackendBucketSignedUrlKey, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendBuckets)
// * How-to Guides
//     * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
//
// > **Warning:** All arguments including `keyValue` will be stored in the raw
// state as plain-text.
//
// ## Example Usage
// ### Backend Bucket Signed Url Key
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		urlSignature, err := random.NewRandomId(ctx, "urlSignature", &random.RandomIdArgs{
// 			ByteLength: pulumi.Int(16),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
// 			Location: pulumi.String("EU"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testBackend, err := compute.NewBackendBucket(ctx, "testBackend", &compute.BackendBucketArgs{
// 			Description: pulumi.String("Contains beautiful images"),
// 			BucketName:  bucket.Name,
// 			EnableCdn:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewBackendBucketSignedUrlKey(ctx, "backendKey", &compute.BackendBucketSignedUrlKeyArgs{
// 			KeyValue:      urlSignature.B64Url,
// 			BackendBucket: testBackend.Name,
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
// This resource does not support import.
type BackendBucketSignedUrlKey struct {
	pulumi.CustomResourceState

	// The backend bucket this signed URL key belongs.
	BackendBucket pulumi.StringOutput `pulumi:"backendBucket"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringOutput `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBackendBucketSignedUrlKey registers a new resource with the given unique name, arguments, and options.
func NewBackendBucketSignedUrlKey(ctx *pulumi.Context,
	name string, args *BackendBucketSignedUrlKeyArgs, opts ...pulumi.ResourceOption) (*BackendBucketSignedUrlKey, error) {
	if args == nil || args.BackendBucket == nil {
		return nil, errors.New("missing required argument 'BackendBucket'")
	}
	if args == nil || args.KeyValue == nil {
		return nil, errors.New("missing required argument 'KeyValue'")
	}
	if args == nil {
		args = &BackendBucketSignedUrlKeyArgs{}
	}
	var resource BackendBucketSignedUrlKey
	err := ctx.RegisterResource("gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucketSignedUrlKey gets an existing BackendBucketSignedUrlKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucketSignedUrlKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketSignedUrlKeyState, opts ...pulumi.ResourceOption) (*BackendBucketSignedUrlKey, error) {
	var resource BackendBucketSignedUrlKey
	err := ctx.ReadResource("gcp:compute/backendBucketSignedUrlKey:BackendBucketSignedUrlKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucketSignedUrlKey resources.
type backendBucketSignedUrlKeyState struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket *string `pulumi:"backendBucket"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue *string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type BackendBucketSignedUrlKeyState struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket pulumi.StringPtrInput
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringPtrInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketSignedUrlKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketSignedUrlKeyState)(nil)).Elem()
}

type backendBucketSignedUrlKeyArgs struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket string `pulumi:"backendBucket"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendBucketSignedUrlKey resource.
type BackendBucketSignedUrlKeyArgs struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket pulumi.StringInput
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketSignedUrlKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketSignedUrlKeyArgs)(nil)).Elem()
}

type BackendBucketSignedUrlKeyInput interface {
	pulumi.Input

	ToBackendBucketSignedUrlKeyOutput() BackendBucketSignedUrlKeyOutput
	ToBackendBucketSignedUrlKeyOutputWithContext(ctx context.Context) BackendBucketSignedUrlKeyOutput
}

func (BackendBucketSignedUrlKey) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendBucketSignedUrlKey)(nil)).Elem()
}

func (i BackendBucketSignedUrlKey) ToBackendBucketSignedUrlKeyOutput() BackendBucketSignedUrlKeyOutput {
	return i.ToBackendBucketSignedUrlKeyOutputWithContext(context.Background())
}

func (i BackendBucketSignedUrlKey) ToBackendBucketSignedUrlKeyOutputWithContext(ctx context.Context) BackendBucketSignedUrlKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketSignedUrlKeyOutput)
}

type BackendBucketSignedUrlKeyOutput struct {
	*pulumi.OutputState
}

func (BackendBucketSignedUrlKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendBucketSignedUrlKeyOutput)(nil)).Elem()
}

func (o BackendBucketSignedUrlKeyOutput) ToBackendBucketSignedUrlKeyOutput() BackendBucketSignedUrlKeyOutput {
	return o
}

func (o BackendBucketSignedUrlKeyOutput) ToBackendBucketSignedUrlKeyOutputWithContext(ctx context.Context) BackendBucketSignedUrlKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackendBucketSignedUrlKeyOutput{})
}
