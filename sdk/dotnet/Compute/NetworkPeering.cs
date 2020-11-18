// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Manages a network peering within GCE. For more information see
    /// [the official documentation](https://cloud.google.com/compute/docs/vpc/vpc-peering)
    /// and
    /// [API](https://cloud.google.com/compute/docs/reference/latest/networks).
    /// 
    /// &gt; Both network must create a peering with each other for the peering
    /// to be functional.
    /// 
    /// &gt; Subnets IP ranges across peered VPC networks cannot overlap.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.Compute.Network("default", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var other = new Gcp.Compute.Network("other", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var peering1 = new Gcp.Compute.NetworkPeering("peering1", new Gcp.Compute.NetworkPeeringArgs
    ///         {
    ///             Network = @default.Id,
    ///             PeerNetwork = other.Id,
    ///         });
    ///         var peering2 = new Gcp.Compute.NetworkPeering("peering2", new Gcp.Compute.NetworkPeeringArgs
    ///         {
    ///             Network = other.Id,
    ///             PeerNetwork = @default.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC network peerings can be imported using the name and project of the primary network the peering exists in and the name of the network peering
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/networkPeering:NetworkPeering peering_network project-name/network-name/peering-name
    /// ```
    /// </summary>
    public partial class NetworkPeering : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to export the custom routes to the peer network. Defaults to `false`.
        /// </summary>
        [Output("exportCustomRoutes")]
        public Output<bool?> ExportCustomRoutes { get; private set; } = null!;

        /// <summary>
        /// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
        /// </summary>
        [Output("exportSubnetRoutesWithPublicIp")]
        public Output<bool?> ExportSubnetRoutesWithPublicIp { get; private set; } = null!;

        /// <summary>
        /// Whether to import the custom routes from the peer network. Defaults to `false`.
        /// </summary>
        [Output("importCustomRoutes")]
        public Output<bool?> ImportCustomRoutes { get; private set; } = null!;

        /// <summary>
        /// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
        /// </summary>
        [Output("importSubnetRoutesWithPublicIp")]
        public Output<bool?> ImportSubnetRoutesWithPublicIp { get; private set; } = null!;

        /// <summary>
        /// Name of the peering.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The primary network of the peering.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The peer network in the peering. The peer network
        /// may belong to a different project.
        /// </summary>
        [Output("peerNetwork")]
        public Output<string> PeerNetwork { get; private set; } = null!;

        /// <summary>
        /// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
        /// `ACTIVE` when there's a matching configuration in the peer network.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Details about the current state of the peering.
        /// </summary>
        [Output("stateDetails")]
        public Output<string> StateDetails { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkPeering(string name, NetworkPeeringArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/networkPeering:NetworkPeering", name, args ?? new NetworkPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkPeering(string name, Input<string> id, NetworkPeeringState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/networkPeering:NetworkPeering", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkPeering Get(string name, Input<string> id, NetworkPeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkPeering(name, id, state, options);
        }
    }

    public sealed class NetworkPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to export the custom routes to the peer network. Defaults to `false`.
        /// </summary>
        [Input("exportCustomRoutes")]
        public Input<bool>? ExportCustomRoutes { get; set; }

        /// <summary>
        /// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
        /// </summary>
        [Input("exportSubnetRoutesWithPublicIp")]
        public Input<bool>? ExportSubnetRoutesWithPublicIp { get; set; }

        /// <summary>
        /// Whether to import the custom routes from the peer network. Defaults to `false`.
        /// </summary>
        [Input("importCustomRoutes")]
        public Input<bool>? ImportCustomRoutes { get; set; }

        /// <summary>
        /// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
        /// </summary>
        [Input("importSubnetRoutesWithPublicIp")]
        public Input<bool>? ImportSubnetRoutesWithPublicIp { get; set; }

        /// <summary>
        /// Name of the peering.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The primary network of the peering.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The peer network in the peering. The peer network
        /// may belong to a different project.
        /// </summary>
        [Input("peerNetwork", required: true)]
        public Input<string> PeerNetwork { get; set; } = null!;

        public NetworkPeeringArgs()
        {
        }
    }

    public sealed class NetworkPeeringState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to export the custom routes to the peer network. Defaults to `false`.
        /// </summary>
        [Input("exportCustomRoutes")]
        public Input<bool>? ExportCustomRoutes { get; set; }

        /// <summary>
        /// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
        /// </summary>
        [Input("exportSubnetRoutesWithPublicIp")]
        public Input<bool>? ExportSubnetRoutesWithPublicIp { get; set; }

        /// <summary>
        /// Whether to import the custom routes from the peer network. Defaults to `false`.
        /// </summary>
        [Input("importCustomRoutes")]
        public Input<bool>? ImportCustomRoutes { get; set; }

        /// <summary>
        /// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
        /// </summary>
        [Input("importSubnetRoutesWithPublicIp")]
        public Input<bool>? ImportSubnetRoutesWithPublicIp { get; set; }

        /// <summary>
        /// Name of the peering.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The primary network of the peering.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The peer network in the peering. The peer network
        /// may belong to a different project.
        /// </summary>
        [Input("peerNetwork")]
        public Input<string>? PeerNetwork { get; set; }

        /// <summary>
        /// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
        /// `ACTIVE` when there's a matching configuration in the peer network.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Details about the current state of the peering.
        /// </summary>
        [Input("stateDetails")]
        public Input<string>? StateDetails { get; set; }

        public NetworkPeeringState()
        {
        }
    }
}
