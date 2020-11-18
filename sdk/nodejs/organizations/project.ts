// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows creation and management of a Google Cloud Platform project.
 *
 * Projects created with this resource must be associated with an Organization.
 * See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
 *
 * The service account used to run this provider when creating a `gcp.organizations.Project`
 * resource must have `roles/resourcemanager.projectCreator`. See the
 * [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
 * doc for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myProject = new gcp.organizations.Project("my_project", {
 *     orgId: "1234567",
 *     projectId: "your-project-id",
 * });
 * ```
 *
 * To create a project under a specific folder
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const department1 = new gcp.organizations.Folder("department1", {
 *     displayName: "Department 1",
 *     parent: "organizations/1234567",
 * });
 * const myProject_in_a_folder = new gcp.organizations.Project("myProject-in-a-folder", {
 *     projectId: "your-project-id",
 *     folderId: department1.name,
 * });
 * ```
 *
 * ## Import
 *
 * Projects can be imported using the `project_id`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:organizations/project:Project my_project your-project-id
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:organizations/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Create the 'default' network automatically.  Default `true`.
     * If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
     * will still need to have 1 network slot available to create the project successfully, even if
     * you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
     */
    public readonly autoCreateNetwork!: pulumi.Output<boolean | undefined>;
    /**
     * The alphanumeric ID of the billing account this project
     * belongs to. The user or service account performing this operation with the provider
     * must have Billing Account Administrator privileges (`roles/billing.admin`) in
     * the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
     * for more details.
     */
    public readonly billingAccount!: pulumi.Output<string | undefined>;
    /**
     * The numeric ID of the folder this project should be
     * created under. Only one of `orgId` or `folderId` may be
     * specified. If the `folderId` is specified, then the project is
     * created under the specified folder. Changing this forces the
     * project to be migrated to the newly specified folder.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the project.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The display name of the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The numeric identifier of the project.
     */
    public /*out*/ readonly number!: pulumi.Output<string>;
    /**
     * The numeric ID of the organization this project belongs to.
     * Changing this forces a new project to be created.  Only one of
     * `orgId` or `folderId` may be specified. If the `orgId` is
     * specified then the project is created at the top level. Changing
     * this forces the project to be migrated to the newly specified
     * organization.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The project ID. Changing this forces a new project to be created.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * If true, the resource can be deleted
     * without deleting the Project via the Google API.
     */
    public readonly skipDelete!: pulumi.Output<boolean>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["autoCreateNetwork"] = state ? state.autoCreateNetwork : undefined;
            inputs["billingAccount"] = state ? state.billingAccount : undefined;
            inputs["folderId"] = state ? state.folderId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["number"] = state ? state.number : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["skipDelete"] = state ? state.skipDelete : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["autoCreateNetwork"] = args ? args.autoCreateNetwork : undefined;
            inputs["billingAccount"] = args ? args.billingAccount : undefined;
            inputs["folderId"] = args ? args.folderId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["skipDelete"] = args ? args.skipDelete : undefined;
            inputs["number"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Create the 'default' network automatically.  Default `true`.
     * If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
     * will still need to have 1 network slot available to create the project successfully, even if
     * you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
     */
    readonly autoCreateNetwork?: pulumi.Input<boolean>;
    /**
     * The alphanumeric ID of the billing account this project
     * belongs to. The user or service account performing this operation with the provider
     * must have Billing Account Administrator privileges (`roles/billing.admin`) in
     * the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
     * for more details.
     */
    readonly billingAccount?: pulumi.Input<string>;
    /**
     * The numeric ID of the folder this project should be
     * created under. Only one of `orgId` or `folderId` may be
     * specified. If the `folderId` is specified, then the project is
     * created under the specified folder. Changing this forces the
     * project to be migrated to the newly specified folder.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the project.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The display name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The numeric identifier of the project.
     */
    readonly number?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization this project belongs to.
     * Changing this forces a new project to be created.  Only one of
     * `orgId` or `folderId` may be specified. If the `orgId` is
     * specified then the project is created at the top level. Changing
     * this forces the project to be migrated to the newly specified
     * organization.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * The project ID. Changing this forces a new project to be created.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * If true, the resource can be deleted
     * without deleting the Project via the Google API.
     */
    readonly skipDelete?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Create the 'default' network automatically.  Default `true`.
     * If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
     * will still need to have 1 network slot available to create the project successfully, even if
     * you set `autoCreateNetwork` to `false`, since the network will exist momentarily.
     */
    readonly autoCreateNetwork?: pulumi.Input<boolean>;
    /**
     * The alphanumeric ID of the billing account this project
     * belongs to. The user or service account performing this operation with the provider
     * must have Billing Account Administrator privileges (`roles/billing.admin`) in
     * the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
     * for more details.
     */
    readonly billingAccount?: pulumi.Input<string>;
    /**
     * The numeric ID of the folder this project should be
     * created under. Only one of `orgId` or `folderId` may be
     * specified. If the `folderId` is specified, then the project is
     * created under the specified folder. Changing this forces the
     * project to be migrated to the newly specified folder.
     */
    readonly folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the project.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The display name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization this project belongs to.
     * Changing this forces a new project to be created.  Only one of
     * `orgId` or `folderId` may be specified. If the `orgId` is
     * specified then the project is created at the top level. Changing
     * this forces the project to be migrated to the newly specified
     * organization.
     */
    readonly orgId?: pulumi.Input<string>;
    /**
     * The project ID. Changing this forces a new project to be created.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * If true, the resource can be deleted
     * without deleting the Project via the Google API.
     */
    readonly skipDelete?: pulumi.Input<boolean>;
}
