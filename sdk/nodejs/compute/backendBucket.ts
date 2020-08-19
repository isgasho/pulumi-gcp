// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S)
 * load balancing.
 *
 * An HTTP(S) load balancer can direct traffic to specified URLs to a
 * backend bucket rather than a backend service. It can send requests for
 * static content to a Cloud Storage bucket and requests for dynamic content
 * to a virtual machine instance.
 *
 *
 * To get more information about BackendBucket, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendBuckets)
 * * How-to Guides
 *     * [Using a Cloud Storage bucket as a load balancer backend](https://cloud.google.com/compute/docs/load-balancing/http/backend-bucket)
 *
 * ## Example Usage
 *
 * ### Backend Bucket Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const imageBucket = new gcp.storage.Bucket("imageBucket", {location: "EU"});
 * const imageBackend = new gcp.compute.BackendBucket("imageBackend", {
 *     description: "Contains beautiful images",
 *     bucketName: imageBucket.name,
 *     enableCdn: true,
 * });
 * ```
 */
export class BackendBucket extends pulumi.CustomResource {
    /**
     * Get an existing BackendBucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendBucketState, opts?: pulumi.CustomResourceOptions): BackendBucket {
        return new BackendBucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/backendBucket:BackendBucket';

    /**
     * Returns true if the given object is an instance of BackendBucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendBucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendBucket.__pulumiType;
    }

    /**
     * Cloud Storage bucket name.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * Cloud CDN configuration for this Backend Bucket.
     * Structure is documented below.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.BackendBucketCdnPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional textual description of the resource; provided by the
     * client when the resource is created.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If true, enable Cloud CDN for this BackendBucket.
     */
    public readonly enableCdn!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a BackendBucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendBucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendBucketArgs | BackendBucketState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendBucketState | undefined;
            inputs["bucketName"] = state ? state.bucketName : undefined;
            inputs["cdnPolicy"] = state ? state.cdnPolicy : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableCdn"] = state ? state.enableCdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as BackendBucketArgs | undefined;
            if (!args || args.bucketName === undefined) {
                throw new Error("Missing required property 'bucketName'");
            }
            inputs["bucketName"] = args ? args.bucketName : undefined;
            inputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableCdn"] = args ? args.enableCdn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BackendBucket.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendBucket resources.
 */
export interface BackendBucketState {
    /**
     * Cloud Storage bucket name.
     */
    readonly bucketName?: pulumi.Input<string>;
    /**
     * Cloud CDN configuration for this Backend Bucket.
     * Structure is documented below.
     */
    readonly cdnPolicy?: pulumi.Input<inputs.compute.BackendBucketCdnPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional textual description of the resource; provided by the
     * client when the resource is created.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this BackendBucket.
     */
    readonly enableCdn?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendBucket resource.
 */
export interface BackendBucketArgs {
    /**
     * Cloud Storage bucket name.
     */
    readonly bucketName: pulumi.Input<string>;
    /**
     * Cloud CDN configuration for this Backend Bucket.
     * Structure is documented below.
     */
    readonly cdnPolicy?: pulumi.Input<inputs.compute.BackendBucketCdnPolicy>;
    /**
     * An optional textual description of the resource; provided by the
     * client when the resource is created.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this BackendBucket.
     */
    readonly enableCdn?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
