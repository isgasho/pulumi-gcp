// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Cloud Firebase web application instance
 *
 * To get more information about WebApp, see:
 *
 * * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
 * * How-to Guides
 *     * [Official Documentation](https://firebase.google.com/)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * WebApp can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:firebase/webApp:WebApp default {{name}}
 * ```
 */
export class WebApp extends pulumi.CustomResource {
    /**
     * Get an existing WebApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAppState, opts?: pulumi.CustomResourceOptions): WebApp {
        return new WebApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:firebase/webApp:WebApp';

    /**
     * Returns true if the given object is an instance of WebApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebApp.__pulumiType;
    }

    /**
     * Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
     * token, as the data format is not specified.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * The user-assigned display name of the App.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a WebApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppArgs | WebAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebAppState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as WebAppArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["appId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WebApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebApp resources.
 */
export interface WebAppState {
    /**
     * Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
     * token, as the data format is not specified.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * The user-assigned display name of the App.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebApp resource.
 */
export interface WebAppArgs {
    /**
     * The user-assigned display name of the App.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
