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
    /// Manages a Cloud Router interface. For more information see
    /// [the official documentation](https://cloud.google.com/compute/docs/cloudrouter)
    /// and
    /// [API](https://cloud.google.com/compute/docs/reference/latest/routers).
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
    ///         var foobar = new Gcp.Compute.RouterInterface("foobar", new Gcp.Compute.RouterInterfaceArgs
    ///         {
    ///             IpRange = "169.254.1.1/30",
    ///             Region = "us-central1",
    ///             Router = "router-1",
    ///             VpnTunnel = "tunnel-1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router interfaces can be imported using the `region`, `router`, and `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/routerInterface:RouterInterface foobar us-central1/router-1/interface-1
    /// ```
    /// </summary>
    public partial class RouterInterface : Pulumi.CustomResource
    {
        /// <summary>
        /// The name or resource link to the
        /// VLAN interconnect for this interface. Changing this forces a new interface to
        /// be created. Only one of `vpn_tunnel` and `interconnect_attachment` can be
        /// specified.
        /// </summary>
        [Output("interconnectAttachment")]
        public Output<string?> InterconnectAttachment { get; private set; } = null!;

        /// <summary>
        /// IP address and range of the interface. The IP range must be
        /// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
        /// </summary>
        [Output("ipRange")]
        public Output<string?> IpRange { get; private set; } = null!;

        /// <summary>
        /// A unique name for the interface, required by GCE. Changing
        /// this forces a new interface to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which this interface's router belongs. If it
        /// is not provided, the provider project is used. Changing this forces a new interface to be created.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region this interface's router sits in. If not specified,
        /// the project region will be used. Changing this forces a new interface to be
        /// created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The name of the router this interface will be attached to.
        /// Changing this forces a new interface to be created.
        /// </summary>
        [Output("router")]
        public Output<string> Router { get; private set; } = null!;

        /// <summary>
        /// The name or resource link to the VPN tunnel this
        /// interface will be linked to. Changing this forces a new interface to be created. Only
        /// one of `vpn_tunnel` and `interconnect_attachment` can be specified.
        /// </summary>
        [Output("vpnTunnel")]
        public Output<string?> VpnTunnel { get; private set; } = null!;


        /// <summary>
        /// Create a RouterInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterInterface(string name, RouterInterfaceArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/routerInterface:RouterInterface", name, args ?? new RouterInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouterInterface(string name, Input<string> id, RouterInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/routerInterface:RouterInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterInterface Get(string name, Input<string> id, RouterInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new RouterInterface(name, id, state, options);
        }
    }

    public sealed class RouterInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or resource link to the
        /// VLAN interconnect for this interface. Changing this forces a new interface to
        /// be created. Only one of `vpn_tunnel` and `interconnect_attachment` can be
        /// specified.
        /// </summary>
        [Input("interconnectAttachment")]
        public Input<string>? InterconnectAttachment { get; set; }

        /// <summary>
        /// IP address and range of the interface. The IP range must be
        /// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// A unique name for the interface, required by GCE. Changing
        /// this forces a new interface to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which this interface's router belongs. If it
        /// is not provided, the provider project is used. Changing this forces a new interface to be created.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this interface's router sits in. If not specified,
        /// the project region will be used. Changing this forces a new interface to be
        /// created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the router this interface will be attached to.
        /// Changing this forces a new interface to be created.
        /// </summary>
        [Input("router", required: true)]
        public Input<string> Router { get; set; } = null!;

        /// <summary>
        /// The name or resource link to the VPN tunnel this
        /// interface will be linked to. Changing this forces a new interface to be created. Only
        /// one of `vpn_tunnel` and `interconnect_attachment` can be specified.
        /// </summary>
        [Input("vpnTunnel")]
        public Input<string>? VpnTunnel { get; set; }

        public RouterInterfaceArgs()
        {
        }
    }

    public sealed class RouterInterfaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or resource link to the
        /// VLAN interconnect for this interface. Changing this forces a new interface to
        /// be created. Only one of `vpn_tunnel` and `interconnect_attachment` can be
        /// specified.
        /// </summary>
        [Input("interconnectAttachment")]
        public Input<string>? InterconnectAttachment { get; set; }

        /// <summary>
        /// IP address and range of the interface. The IP range must be
        /// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// A unique name for the interface, required by GCE. Changing
        /// this forces a new interface to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which this interface's router belongs. If it
        /// is not provided, the provider project is used. Changing this forces a new interface to be created.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this interface's router sits in. If not specified,
        /// the project region will be used. Changing this forces a new interface to be
        /// created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the router this interface will be attached to.
        /// Changing this forces a new interface to be created.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// The name or resource link to the VPN tunnel this
        /// interface will be linked to. Changing this forces a new interface to be created. Only
        /// one of `vpn_tunnel` and `interconnect_attachment` can be specified.
        /// </summary>
        [Input("vpnTunnel")]
        public Input<string>? VpnTunnel { get; set; }

        public RouterInterfaceState()
        {
        }
    }
}
