// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
 * 
 * * `gcp.dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
 * * `gcp.dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
 * * `gcp.dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
 * 
 * > **Note:** `gcp.dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `gcp.dataproc.ClusterIAMBinding` and `gcp.dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `gcp.dataproc.ClusterIAMPolicy` replaces the entire policy.
 * 
 * > **Note:** `gcp.dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `gcp.dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_pubsub\_subscription\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.dataproc.ClusterIAMBinding("editor", {
 *     cluster: "your-dataproc-cluster",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 * 
 * ## google\_pubsub\_subscription\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const editor = new gcp.dataproc.ClusterIAMMember("editor", {
 *     cluster: "your-dataproc-cluster",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster_iam.html.markdown.
 */
export class ClusterIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing ClusterIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterIAMMemberState, opts?: pulumi.CustomResourceOptions): ClusterIAMMember {
        return new ClusterIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/clusterIAMMember:ClusterIAMMember';

    /**
     * Returns true if the given object is an instance of ClusterIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterIAMMember.__pulumiType;
    }

    /**
     * The name or relative resource id of the cluster to manage IAM policies for.
     */
    public readonly cluster!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<outputs.dataproc.ClusterIAMMemberCondition | undefined>;
    /**
     * (Computed) The etag of the clusters's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The project in which the cluster belongs. If it
     * is not provided, the provider will use a default.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the cluster belongs. If it
     * is not provided, the provider will use a default.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ClusterIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterIAMMemberArgs | ClusterIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterIAMMemberState | undefined;
            inputs["cluster"] = state ? state.cluster : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ClusterIAMMemberArgs | undefined;
            if (!args || args.cluster === undefined) {
                throw new Error("Missing required property 'cluster'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterIAMMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterIAMMember resources.
 */
export interface ClusterIAMMemberState {
    /**
     * The name or relative resource id of the cluster to manage IAM policies for.
     */
    readonly cluster?: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.dataproc.ClusterIAMMemberCondition>;
    /**
     * (Computed) The etag of the clusters's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The project in which the cluster belongs. If it
     * is not provided, the provider will use a default.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster belongs. If it
     * is not provided, the provider will use a default.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterIAMMember resource.
 */
export interface ClusterIAMMemberArgs {
    /**
     * The name or relative resource id of the cluster to manage IAM policies for.
     */
    readonly cluster: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.dataproc.ClusterIAMMemberCondition>;
    readonly member: pulumi.Input<string>;
    /**
     * The project in which the cluster belongs. If it
     * is not provided, the provider will use a default.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster belongs. If it
     * is not provided, the provider will use a default.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
