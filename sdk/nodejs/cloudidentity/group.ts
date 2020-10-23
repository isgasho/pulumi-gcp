// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Cloud Identity resource representing a Group.
 *
 * > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
 * you must specify a `billingProject` and set `userProjectOverride` to true
 * in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
 * Your account must have the `serviceusage.services.use` permission on the
 * `billingProject` you defined.
 *
 * ## Example Usage
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudidentity/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * The time when the Group was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An extended description to help users determine the purpose of a Group.
     * Must not be longer than 4,096 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the Group.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * EntityKey of the Group.
     * Structure is documented below.
     */
    public readonly groupKey!: pulumi.Output<outputs.cloudidentity.GroupGroupKey>;
    /**
     * The labels that apply to the Group.
     * Must not contain more than one entry. Must contain the entry
     * 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
     * 'system/groups/external': '' if the Group is an external-identity-mapped group.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource name of the entity under which this Group resides in the
     * Cloud Identity resource hierarchy.
     * Must be of the form identitysources/{identity_source_id} for external-identity-mapped
     * groups or customers/{customer_id} for Google Groups.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * The time when the Group was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["groupKey"] = state ? state.groupKey : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if (!args || args.groupKey === undefined) {
                throw new Error("Missing required property 'groupKey'");
            }
            if (!args || args.labels === undefined) {
                throw new Error("Missing required property 'labels'");
            }
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["groupKey"] = args ? args.groupKey : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * The time when the Group was created.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * An extended description to help users determine the purpose of a Group.
     * Must not be longer than 4,096 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the Group.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * EntityKey of the Group.
     * Structure is documented below.
     */
    readonly groupKey?: pulumi.Input<inputs.cloudidentity.GroupGroupKey>;
    /**
     * The labels that apply to the Group.
     * Must not contain more than one entry. Must contain the entry
     * 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
     * 'system/groups/external': '' if the Group is an external-identity-mapped group.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource name of the entity under which this Group resides in the
     * Cloud Identity resource hierarchy.
     * Must be of the form identitysources/{identity_source_id} for external-identity-mapped
     * groups or customers/{customer_id} for Google Groups.
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * The time when the Group was last updated.
     */
    readonly updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * An extended description to help users determine the purpose of a Group.
     * Must not be longer than 4,096 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the Group.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * EntityKey of the Group.
     * Structure is documented below.
     */
    readonly groupKey: pulumi.Input<inputs.cloudidentity.GroupGroupKey>;
    /**
     * The labels that apply to the Group.
     * Must not contain more than one entry. Must contain the entry
     * 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
     * 'system/groups/external': '' if the Group is an external-identity-mapped group.
     */
    readonly labels: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name of the entity under which this Group resides in the
     * Cloud Identity resource hierarchy.
     * Must be of the form identitysources/{identity_source_id} for external-identity-mapped
     * groups or customers/{customer_id} for Google Groups.
     */
    readonly parent: pulumi.Input<string>;
}
