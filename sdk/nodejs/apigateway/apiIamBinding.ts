// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway api IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiIamBinding:ApiIamBinding editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiIamBinding:ApiIamBinding editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiIamBinding:ApiIamBinding editor projects/{{project}}/locations/global/apis/{{api}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class ApiIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing ApiIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiIamBindingState, opts?: pulumi.CustomResourceOptions): ApiIamBinding {
        return new ApiIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigateway/apiIamBinding:ApiIamBinding';

    /**
     * Returns true if the given object is an instance of ApiIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiIamBinding.__pulumiType;
    }

    public readonly api!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<outputs.apigateway.ApiIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ApiIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiIamBindingArgs | ApiIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApiIamBindingState | undefined;
            inputs["api"] = state ? state.api : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ApiIamBindingArgs | undefined;
            if (!args || args.api === undefined) {
                throw new Error("Missing required property 'api'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["api"] = args ? args.api : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ApiIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiIamBinding resources.
 */
export interface ApiIamBindingState {
    readonly api?: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.apigateway.ApiIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiIamBinding resource.
 */
export interface ApiIamBindingArgs {
    readonly api: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.apigateway.ApiIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
