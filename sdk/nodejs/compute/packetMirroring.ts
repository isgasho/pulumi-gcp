// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Packet Mirroring mirrors traffic to and from particular VM instances.
 * You can use the collected traffic to help you detect security threats
 * and monitor application performance.
 * 
 * To get more information about PacketMirroring, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/packetMirroring)
 * * How-to Guides
 *     * [Using Packet Mirroring](https://cloud.google.com/vpc/docs/using-packet-mirroring#creating)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_packet_mirroring.html.markdown.
 */
export class PacketMirroring extends pulumi.CustomResource {
    /**
     * Get an existing PacketMirroring resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PacketMirroringState, opts?: pulumi.CustomResourceOptions): PacketMirroring {
        return new PacketMirroring(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/packetMirroring:PacketMirroring';

    /**
     * Returns true if the given object is an instance of PacketMirroring.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PacketMirroring {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PacketMirroring.__pulumiType;
    }

    /**
     * The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
     * traffic. The specified forwarding rule must have is_mirroring_collector set to true.
     */
    public readonly collectorIlb!: pulumi.Output<outputs.compute.PacketMirroringCollectorIlb>;
    /**
     * A human-readable description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A filter for mirrored traffic. If unset, all traffic is mirrored.
     */
    public readonly filter!: pulumi.Output<outputs.compute.PacketMirroringFilter | undefined>;
    /**
     * A means of specifying which resources to mirror.
     */
    public readonly mirroredResources!: pulumi.Output<outputs.compute.PacketMirroringMirroredResources>;
    /**
     * The name of the packet mirroring rule
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a
     * NIC in the given network. All mirrored subnetworks should belong to the given network.
     */
    public readonly network!: pulumi.Output<outputs.compute.PacketMirroringNetwork>;
    /**
     * Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to
     * the same instances.
     */
    public readonly priority!: pulumi.Output<number>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The Region in which the created address should reside. If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a PacketMirroring resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PacketMirroringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PacketMirroringArgs | PacketMirroringState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PacketMirroringState | undefined;
            inputs["collectorIlb"] = state ? state.collectorIlb : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["mirroredResources"] = state ? state.mirroredResources : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as PacketMirroringArgs | undefined;
            if (!args || args.collectorIlb === undefined) {
                throw new Error("Missing required property 'collectorIlb'");
            }
            if (!args || args.mirroredResources === undefined) {
                throw new Error("Missing required property 'mirroredResources'");
            }
            if (!args || args.network === undefined) {
                throw new Error("Missing required property 'network'");
            }
            inputs["collectorIlb"] = args ? args.collectorIlb : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["mirroredResources"] = args ? args.mirroredResources : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PacketMirroring.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PacketMirroring resources.
 */
export interface PacketMirroringState {
    /**
     * The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
     * traffic. The specified forwarding rule must have is_mirroring_collector set to true.
     */
    readonly collectorIlb?: pulumi.Input<inputs.compute.PacketMirroringCollectorIlb>;
    /**
     * A human-readable description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A filter for mirrored traffic. If unset, all traffic is mirrored.
     */
    readonly filter?: pulumi.Input<inputs.compute.PacketMirroringFilter>;
    /**
     * A means of specifying which resources to mirror.
     */
    readonly mirroredResources?: pulumi.Input<inputs.compute.PacketMirroringMirroredResources>;
    /**
     * The name of the packet mirroring rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a
     * NIC in the given network. All mirrored subnetworks should belong to the given network.
     */
    readonly network?: pulumi.Input<inputs.compute.PacketMirroringNetwork>;
    /**
     * Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to
     * the same instances.
     */
    readonly priority?: pulumi.Input<number>;
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside. If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PacketMirroring resource.
 */
export interface PacketMirroringArgs {
    /**
     * The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for mirrored
     * traffic. The specified forwarding rule must have is_mirroring_collector set to true.
     */
    readonly collectorIlb: pulumi.Input<inputs.compute.PacketMirroringCollectorIlb>;
    /**
     * A human-readable description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A filter for mirrored traffic. If unset, all traffic is mirrored.
     */
    readonly filter?: pulumi.Input<inputs.compute.PacketMirroringFilter>;
    /**
     * A means of specifying which resources to mirror.
     */
    readonly mirroredResources: pulumi.Input<inputs.compute.PacketMirroringMirroredResources>;
    /**
     * The name of the packet mirroring rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a
     * NIC in the given network. All mirrored subnetworks should belong to the given network.
     */
    readonly network: pulumi.Input<inputs.compute.PacketMirroringNetwork>;
    /**
     * Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to
     * the same instances.
     */
    readonly priority?: pulumi.Input<number>;
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside. If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
}
