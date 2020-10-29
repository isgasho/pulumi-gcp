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
    /// Manages a VPC network or legacy network resource on GCP.
    /// 
    /// To get more information about Network, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/vpc/docs/vpc)
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class Network : Pulumi.CustomResource
    {
        /// <summary>
        /// When set to `true`, the network is created in "auto subnet mode" and
        /// it will create a subnet for each region automatically across the
        /// `10.128.0.0/9` address range.
        /// When set to `false`, the network is created in "custom subnet mode" so
        /// the user can explicitly connect subnetwork resources.
        /// </summary>
        [Output("autoCreateSubnetworks")]
        public Output<bool?> AutoCreateSubnetworks { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, default routes (`0.0.0.0/0`) will be deleted
        /// immediately after network creation. Defaults to `false`.
        /// </summary>
        [Output("deleteDefaultRoutesOnCreate")]
        public Output<bool?> DeleteDefaultRoutesOnCreate { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. The resource must be
        /// recreated to modify this field.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The gateway address for default routing out of the network. This value is selected by GCP.
        /// </summary>
        [Output("gatewayIpv4")]
        public Output<string> GatewayIpv4 { get; private set; } = null!;

        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1460
        /// and the maximum value is 1500 bytes.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The network-wide routing mode to use. If set to `REGIONAL`, this
        /// network's cloud routers will only advertise routes with subnetworks
        /// of this network in the same region as the router. If set to `GLOBAL`,
        /// this network's cloud routers will advertise routes with all
        /// subnetworks of this network, across regions.
        /// Possible values are `REGIONAL` and `GLOBAL`.
        /// </summary>
        [Output("routingMode")]
        public Output<string> RoutingMode { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/network:Network", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to `true`, the network is created in "auto subnet mode" and
        /// it will create a subnet for each region automatically across the
        /// `10.128.0.0/9` address range.
        /// When set to `false`, the network is created in "custom subnet mode" so
        /// the user can explicitly connect subnetwork resources.
        /// </summary>
        [Input("autoCreateSubnetworks")]
        public Input<bool>? AutoCreateSubnetworks { get; set; }

        /// <summary>
        /// If set to `true`, default routes (`0.0.0.0/0`) will be deleted
        /// immediately after network creation. Defaults to `false`.
        /// </summary>
        [Input("deleteDefaultRoutesOnCreate")]
        public Input<bool>? DeleteDefaultRoutesOnCreate { get; set; }

        /// <summary>
        /// An optional description of this resource. The resource must be
        /// recreated to modify this field.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1460
        /// and the maximum value is 1500 bytes.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The network-wide routing mode to use. If set to `REGIONAL`, this
        /// network's cloud routers will only advertise routes with subnetworks
        /// of this network in the same region as the router. If set to `GLOBAL`,
        /// this network's cloud routers will advertise routes with all
        /// subnetworks of this network, across regions.
        /// Possible values are `REGIONAL` and `GLOBAL`.
        /// </summary>
        [Input("routingMode")]
        public Input<string>? RoutingMode { get; set; }

        public NetworkArgs()
        {
        }
    }

    public sealed class NetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to `true`, the network is created in "auto subnet mode" and
        /// it will create a subnet for each region automatically across the
        /// `10.128.0.0/9` address range.
        /// When set to `false`, the network is created in "custom subnet mode" so
        /// the user can explicitly connect subnetwork resources.
        /// </summary>
        [Input("autoCreateSubnetworks")]
        public Input<bool>? AutoCreateSubnetworks { get; set; }

        /// <summary>
        /// If set to `true`, default routes (`0.0.0.0/0`) will be deleted
        /// immediately after network creation. Defaults to `false`.
        /// </summary>
        [Input("deleteDefaultRoutesOnCreate")]
        public Input<bool>? DeleteDefaultRoutesOnCreate { get; set; }

        /// <summary>
        /// An optional description of this resource. The resource must be
        /// recreated to modify this field.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The gateway address for default routing out of the network. This value is selected by GCP.
        /// </summary>
        [Input("gatewayIpv4")]
        public Input<string>? GatewayIpv4 { get; set; }

        /// <summary>
        /// Maximum Transmission Unit in bytes. The minimum value for this field is 1460
        /// and the maximum value is 1500 bytes.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The network-wide routing mode to use. If set to `REGIONAL`, this
        /// network's cloud routers will only advertise routes with subnetworks
        /// of this network in the same region as the router. If set to `GLOBAL`,
        /// this network's cloud routers will advertise routes with all
        /// subnetworks of this network, across regions.
        /// Possible values are `REGIONAL` and `GLOBAL`.
        /// </summary>
        [Input("routingMode")]
        public Input<string>? RoutingMode { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public NetworkState()
        {
        }
    }
}
