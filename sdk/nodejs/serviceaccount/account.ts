// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of a [Google Cloud Platform service account](https://cloud.google.com/compute/docs/access/service-accounts)
 *
 * > Creation of service accounts is eventually consistent, and that can lead to
 * errors when you try to apply ACLs to service accounts immediately after
 * creation.
 *
 * ## Example Usage
 *
 * This snippet creates a service account in a project.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const serviceAccount = new gcp.serviceAccount.Account("service_account", {
 *     accountId: "service_account_id",
 *     displayName: "Service Account",
 * });
 * ```
 *
 * ## Import
 *
 * Service accounts can be imported using their URI, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/account:Account my_sa projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:serviceAccount/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * The account id that is used to generate the service
     * account email address and a stable unique id. It is unique within a project,
     * must be 6-30 characters long, and match the regular expression `a-z`
     * to comply with RFC1035. Changing this forces a new service account to be created.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * A text description of the service account.
     * Must be less than or equal to 256 UTF-8 bytes.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name for the service account.
     * Can be updated without creating a new resource.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The e-mail address of the service account. This value
     * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
     * that would grant the service account privileges.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The fully-qualified name of the service account.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The unique id of the service account.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccountState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["email"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Account.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * The account id that is used to generate the service
     * account email address and a stable unique id. It is unique within a project,
     * must be 6-30 characters long, and match the regular expression `a-z`
     * to comply with RFC1035. Changing this forces a new service account to be created.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * A text description of the service account.
     * Must be less than or equal to 256 UTF-8 bytes.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name for the service account.
     * Can be updated without creating a new resource.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The e-mail address of the service account. This value
     * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
     * that would grant the service account privileges.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The unique id of the service account.
     */
    readonly uniqueId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The account id that is used to generate the service
     * account email address and a stable unique id. It is unique within a project,
     * must be 6-30 characters long, and match the regular expression `a-z`
     * to comply with RFC1035. Changing this forces a new service account to be created.
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * A text description of the service account.
     * Must be less than or equal to 256 UTF-8 bytes.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name for the service account.
     * Can be updated without creating a new resource.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The ID of the project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    readonly project?: pulumi.Input<string>;
}
