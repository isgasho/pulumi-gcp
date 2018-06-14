// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * For more information, see [the official
 * documentation](https://cloud.google.com/source-repositories/) and
 * [API](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState): Repository {
        return new Repository(name, <any>state, { id });
    }

    /**
     * The name of the repository that will be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The size of the repository.
     */
    public /*out*/ readonly size: pulumi.Output<number>;
    /**
     * The url to clone the repository.
     */
    public /*out*/ readonly url: pulumi.Output<string>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RepositoryArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RepositoryState = argsOrState as RepositoryState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["size"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        super("gcp:sourcerepo/repository:Repository", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * The name of the repository that will be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The size of the repository.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * The url to clone the repository.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * The name of the repository that will be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}