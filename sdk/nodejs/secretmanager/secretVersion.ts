// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A secret version resource.
 * 
 * > **Warning:** All arguments including `payload.secret_data` will be stored in the raw
 * state as plain-text.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/secret_manager_secret_version.html.markdown.
 */
export class SecretVersion extends pulumi.CustomResource {
    /**
     * Get an existing SecretVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretVersionState, opts?: pulumi.CustomResourceOptions): SecretVersion {
        return new SecretVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:secretmanager/secretVersion:SecretVersion';

    /**
     * Returns true if the given object is an instance of SecretVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretVersion.__pulumiType;
    }

    /**
     * The time at which the Secret was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The time at which the Secret was destroyed. Only present if state is DESTROYED.
     */
    public /*out*/ readonly destroyTime!: pulumi.Output<string>;
    /**
     * The current state of the SecretVersion.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Secret Manager secret resource
     */
    public readonly secret!: pulumi.Output<string>;
    /**
     * The secret data. Must be no larger than 64KiB.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly secretData!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretVersionArgs | SecretVersionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretVersionState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["destroyTime"] = state ? state.destroyTime : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["secretData"] = state ? state.secretData : undefined;
        } else {
            const args = argsOrState as SecretVersionArgs | undefined;
            if (!args || args.secret === undefined) {
                throw new Error("Missing required property 'secret'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["secret"] = args ? args.secret : undefined;
            inputs["secretData"] = args ? args.secretData : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["destroyTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretVersion resources.
 */
export interface SecretVersionState {
    /**
     * The time at which the Secret was created.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * The time at which the Secret was destroyed. Only present if state is DESTROYED.
     */
    readonly destroyTime?: pulumi.Input<string>;
    /**
     * The current state of the SecretVersion.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The resource name of the SecretVersion. Format: 'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Secret Manager secret resource
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * The secret data. Must be no larger than 64KiB.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly secretData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretVersion resource.
 */
export interface SecretVersionArgs {
    /**
     * The current state of the SecretVersion.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Secret Manager secret resource
     */
    readonly secret: pulumi.Input<string>;
    /**
     * The secret data. Must be no larger than 64KiB.  **Note**: This property is sensitive and will not be displayed in the plan.
     */
    readonly secretData?: pulumi.Input<string>;
}
