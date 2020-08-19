// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Network endpoint groups (NEGs) are zonal resources that represent
 * collections of IP address and port combinations for GCP resources within a
 * single subnet. Each IP address and port combination is called a network
 * endpoint.
 *
 * Network endpoint groups can be used as backends in backend services for
 * HTTP(S), TCP proxy, and SSL proxy load balancers. You cannot use NEGs as a
 * backend with internal load balancers. Because NEG backends allow you to
 * specify IP addresses and ports, you can distribute traffic in a granular
 * fashion among applications or containers running within VM instances.
 *
 *
 * To get more information about NetworkEndpointGroup, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
 *
 * ## Example Usage
 *
 * ### Network Endpoint Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false});
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const neg = new gcp.compute.NetworkEndpointGroup("neg", {
 *     network: defaultNetwork.id,
 *     subnetwork: defaultSubnetwork.id,
 *     defaultPort: "90",
 *     zone: "us-central1-a",
 * });
 * ```
 */
export class NetworkEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing NetworkEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkEndpointGroupState, opts?: pulumi.CustomResourceOptions): NetworkEndpointGroup {
        return new NetworkEndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/networkEndpointGroup:NetworkEndpointGroup';

    /**
     * Returns true if the given object is an instance of NetworkEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkEndpointGroup.__pulumiType;
    }

    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     */
    public readonly defaultPort!: pulumi.Output<number | undefined>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network to which all network endpoints in the NEG belong.
     * Uses "default" project network if unspecified.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Type of network endpoints in this network endpoint group.
     * Default value is `GCE_VM_IP_PORT`.
     * Possible values are `GCE_VM_IP_PORT`.
     */
    public readonly networkEndpointType!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Number of network endpoints in the network endpoint group.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Optional subnetwork to which all network endpoints in the NEG belong.
     */
    public readonly subnetwork!: pulumi.Output<string | undefined>;
    /**
     * Zone where the network endpoint group is located.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NetworkEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkEndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkEndpointGroupArgs | NetworkEndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkEndpointGroupState | undefined;
            inputs["defaultPort"] = state ? state.defaultPort : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkEndpointType"] = state ? state.networkEndpointType : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NetworkEndpointGroupArgs | undefined;
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["defaultPort"] = args ? args.defaultPort : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkEndpointType"] = args ? args.networkEndpointType : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["selfLink"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkEndpointGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkEndpointGroup resources.
 */
export interface NetworkEndpointGroupState {
    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     */
    readonly defaultPort?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network to which all network endpoints in the NEG belong.
     * Uses "default" project network if unspecified.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * Type of network endpoints in this network endpoint group.
     * Default value is `GCE_VM_IP_PORT`.
     * Possible values are `GCE_VM_IP_PORT`.
     */
    readonly networkEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Number of network endpoints in the network endpoint group.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * Optional subnetwork to which all network endpoints in the NEG belong.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * Zone where the network endpoint group is located.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkEndpointGroup resource.
 */
export interface NetworkEndpointGroupArgs {
    /**
     * The default port used if the port number is not specified in the
     * network endpoint.
     */
    readonly defaultPort?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network to which all network endpoints in the NEG belong.
     * Uses "default" project network if unspecified.
     */
    readonly network: pulumi.Input<string>;
    /**
     * Type of network endpoints in this network endpoint group.
     * Default value is `GCE_VM_IP_PORT`.
     * Possible values are `GCE_VM_IP_PORT`.
     */
    readonly networkEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Optional subnetwork to which all network endpoints in the NEG belong.
     */
    readonly subnetwork?: pulumi.Input<string>;
    /**
     * Zone where the network endpoint group is located.
     */
    readonly zone?: pulumi.Input<string>;
}
