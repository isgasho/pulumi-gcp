// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a VPN gateway running in GCP. This virtual device is managed
 * by Google, but used only by you. This type of VPN Gateway allows for the creation
 * of VPN solutions with higher availability than classic Target VPN Gateways.
 * 
 * > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
 * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
 * 
 * To get more information about HaVpnGateway, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/vpnGateways)
 * * How-to Guides
 *     * [Choosing a VPN](https://cloud.google.com/vpn/docs/how-to/choosing-a-vpn)
 *     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
 */
export class HaVpnGateway extends pulumi.CustomResource {
    /**
     * Get an existing HaVpnGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HaVpnGatewayState, opts?: pulumi.CustomResourceOptions): HaVpnGateway {
        return new HaVpnGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/haVpnGateway:HaVpnGateway';

    /**
     * Returns true if the given object is an instance of HaVpnGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HaVpnGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HaVpnGateway.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly network!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public /*out*/ readonly vpnInterfaces!: pulumi.Output<{ id?: number, ipAddress?: string }[]>;

    /**
     * Create a HaVpnGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HaVpnGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HaVpnGatewayArgs | HaVpnGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HaVpnGatewayState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["vpnInterfaces"] = state ? state.vpnInterfaces : undefined;
        } else {
            const args = argsOrState as HaVpnGatewayArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["selfLink"] = undefined /*out*/;
            inputs["vpnInterfaces"] = undefined /*out*/;
        }
        super(HaVpnGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HaVpnGateway resources.
 */
export interface HaVpnGatewayState {
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly vpnInterfaces?: pulumi.Input<pulumi.Input<{ id?: pulumi.Input<number>, ipAddress?: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a HaVpnGateway resource.
 */
export interface HaVpnGatewayArgs {
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
}
