// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class InstanceNetworkInterfaceGetArgs : Pulumi.ResourceArgs
    {
        [Input("accessConfigs")]
        private InputList<Inputs.InstanceNetworkInterfaceAccessConfigGetArgs>? _accessConfigs;

        /// <summary>
        /// Access configurations, i.e. IPs via which this
        /// instance can be accessed via the Internet. Omit to ensure that the instance
        /// is not accessible from the Internet. If omitted, ssh will not
        /// work unless this provider can send traffic to the instance's network (e.g. via
        /// tunnel or because it is running on another cloud instance on that network).
        /// This block can be repeated multiple times. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceNetworkInterfaceAccessConfigGetArgs> AccessConfigs
        {
            get => _accessConfigs ?? (_accessConfigs = new InputList<Inputs.InstanceNetworkInterfaceAccessConfigGetArgs>());
            set => _accessConfigs = value;
        }

        [Input("aliasIpRanges")]
        private InputList<Inputs.InstanceNetworkInterfaceAliasIpRangeGetArgs>? _aliasIpRanges;

        /// <summary>
        /// An
        /// array of alias IP ranges for this network interface. Can only be specified for network
        /// interfaces on subnet-mode networks. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceNetworkInterfaceAliasIpRangeGetArgs> AliasIpRanges
        {
            get => _aliasIpRanges ?? (_aliasIpRanges = new InputList<Inputs.InstanceNetworkInterfaceAliasIpRangeGetArgs>());
            set => _aliasIpRanges = value;
        }

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or self_link of the network to attach this interface to.
        /// Either `network` or `subnetwork` must be provided.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The private IP address to assign to the instance. If
        /// empty, the address will be automatically assigned.
        /// </summary>
        [Input("networkIp")]
        public Input<string>? NetworkIp { get; set; }

        /// <summary>
        /// The name or self_link of the subnetwork to attach this
        /// interface to. The subnetwork must exist in the same region this instance will be
        /// created in. If network isn't provided it will be inferred from the subnetwork.
        /// Either `network` or `subnetwork` must be provided.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// The project in which the subnetwork belongs.
        /// If the `subnetwork` is a self_link, this field is ignored in favor of the project
        /// defined in the subnetwork self_link. If the `subnetwork` is a name and this
        /// field is not provided, the provider project is used.
        /// </summary>
        [Input("subnetworkProject")]
        public Input<string>? SubnetworkProject { get; set; }

        public InstanceNetworkInterfaceGetArgs()
        {
        }
    }
}
