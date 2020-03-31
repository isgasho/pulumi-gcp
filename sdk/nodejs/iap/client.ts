// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Contains the data that describes an Identity Aware Proxy owned client.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_client.html.markdown.
 */
export class Client extends pulumi.CustomResource {
    /**
     * Get an existing Client resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientState, opts?: pulumi.CustomResourceOptions): Client {
        return new Client(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/client:Client';

    /**
     * Returns true if the given object is an instance of Client.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Client {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Client.__pulumiType;
    }

    /**
     * Identifier of the brand to which this client is attached to. The format is
     * 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
     */
    public readonly brand!: pulumi.Output<string>;
    /**
     * Output only. Unique identifier of the OAuth client.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * Human-friendly name given to the OAuth client.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Output only. Client secret of the OAuth client.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;

    /**
     * Create a Client resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientArgs | ClientState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClientState | undefined;
            inputs["brand"] = state ? state.brand : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["secret"] = state ? state.secret : undefined;
        } else {
            const args = argsOrState as ClientArgs | undefined;
            if (!args || args.brand === undefined) {
                throw new Error("Missing required property 'brand'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["brand"] = args ? args.brand : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["clientId"] = undefined /*out*/;
            inputs["secret"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Client.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Client resources.
 */
export interface ClientState {
    /**
     * Identifier of the brand to which this client is attached to. The format is
     * 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
     */
    readonly brand?: pulumi.Input<string>;
    /**
     * Output only. Unique identifier of the OAuth client.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Human-friendly name given to the OAuth client.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Output only. Client secret of the OAuth client.
     */
    readonly secret?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    /**
     * Identifier of the brand to which this client is attached to. The format is
     * 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
     */
    readonly brand: pulumi.Input<string>;
    /**
     * Human-friendly name given to the OAuth client.
     */
    readonly displayName: pulumi.Input<string>;
}