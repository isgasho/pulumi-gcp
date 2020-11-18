// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Secret Manager Secret. Each of these resources serves a different use case:
 *
 * * `gcp.secretmanager.SecretIamPolicy`: Authoritative. Sets the IAM policy for the secret and replaces any existing policy already attached.
 * * `gcp.secretmanager.SecretIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the secret are preserved.
 * * `gcp.secretmanager.SecretIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the secret are preserved.
 *
 * > **Note:** `gcp.secretmanager.SecretIamPolicy` **cannot** be used in conjunction with `gcp.secretmanager.SecretIamBinding` and `gcp.secretmanager.SecretIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.secretmanager.SecretIamBinding` resources **can be** used in conjunction with `gcp.secretmanager.SecretIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_secret\_manager\_secret\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/secretmanager.secretAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.secretmanager.SecretIamPolicy("policy", {
 *     project: google_secret_manager_secret["secret-basic"].project,
 *     secretId: google_secret_manager_secret["secret-basic"].secret_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_secret\_manager\_secret\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.secretmanager.SecretIamBinding("binding", {
 *     project: google_secret_manager_secret["secret-basic"].project,
 *     secretId: google_secret_manager_secret["secret-basic"].secret_id,
 *     role: "roles/secretmanager.secretAccessor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_secret\_manager\_secret\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.secretmanager.SecretIamMember("member", {
 *     project: google_secret_manager_secret["secret-basic"].project,
 *     secretId: google_secret_manager_secret["secret-basic"].secret_id,
 *     role: "roles/secretmanager.secretAccessor",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/secrets/{{secret_id}} * {{project}}/{{secret_id}} * {{secret_id}} Any variables not passed in the import command will be taken from the provider configuration. Secret Manager secret IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:secretmanager/secretIamBinding:SecretIamBinding editor "projects/{{project}}/secrets/{{secret_id}} roles/secretmanager.secretAccessor user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:secretmanager/secretIamBinding:SecretIamBinding editor "projects/{{project}}/secrets/{{secret_id}} roles/secretmanager.secretAccessor"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:secretmanager/secretIamBinding:SecretIamBinding editor projects/{{project}}/secrets/{{secret_id}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class SecretIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing SecretIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretIamBindingState, opts?: pulumi.CustomResourceOptions): SecretIamBinding {
        return new SecretIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:secretmanager/secretIamBinding:SecretIamBinding';

    /**
     * Returns true if the given object is an instance of SecretIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.secretmanager.SecretIamBindingCondition | undefined>;
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
     * `gcp.secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    public readonly secretId!: pulumi.Output<string>;

    /**
     * Create a SecretIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretIamBindingArgs | SecretIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["secretId"] = state ? state.secretId : undefined;
        } else {
            const args = argsOrState as SecretIamBindingArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.secretId === undefined) {
                throw new Error("Missing required property 'secretId'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["secretId"] = args ? args.secretId : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretIamBinding resources.
 */
export interface SecretIamBindingState {
    readonly condition?: pulumi.Input<inputs.secretmanager.SecretIamBindingCondition>;
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
     * `gcp.secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
    readonly secretId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretIamBinding resource.
 */
export interface SecretIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.secretmanager.SecretIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
    readonly secretId: pulumi.Input<string>;
}
