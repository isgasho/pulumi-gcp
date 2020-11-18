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
    /// Represents a TargetTcpProxy resource, which is used by one or more
    /// global forwarding rule to route incoming TCP requests to a Backend
    /// service.
    /// 
    /// To get more information about TargetTcpProxy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetTcpProxies)
    /// * How-to Guides
    ///     * [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
    /// 
    /// ## Example Usage
    /// ### Target Tcp Proxy Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultHealthCheck = new Gcp.Compute.HealthCheck("defaultHealthCheck", new Gcp.Compute.HealthCheckArgs
    ///         {
    ///             TimeoutSec = 1,
    ///             CheckIntervalSec = 1,
    ///             TcpHealthCheck = new Gcp.Compute.Inputs.HealthCheckTcpHealthCheckArgs
    ///             {
    ///                 Port = 443,
    ///             },
    ///         });
    ///         var defaultBackendService = new Gcp.Compute.BackendService("defaultBackendService", new Gcp.Compute.BackendServiceArgs
    ///         {
    ///             Protocol = "TCP",
    ///             TimeoutSec = 10,
    ///             HealthChecks = 
    ///             {
    ///                 defaultHealthCheck.Id,
    ///             },
    ///         });
    ///         var defaultTargetTCPProxy = new Gcp.Compute.TargetTCPProxy("defaultTargetTCPProxy", new Gcp.Compute.TargetTCPProxyArgs
    ///         {
    ///             BackendService = defaultBackendService.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// TargetTcpProxy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default projects/{{project}}/global/targetTcpProxies/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{name}}
    /// ```
    /// </summary>
    public partial class TargetTCPProxy : Pulumi.CustomResource
    {
        /// <summary>
        /// A reference to the BackendService resource.
        /// </summary>
        [Output("backendService")]
        public Output<string> BackendService { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

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
        /// Specifies the type of proxy header to append before sending data to
        /// the backend.
        /// Default value is `NONE`.
        /// Possible values are `NONE` and `PROXY_V1`.
        /// </summary>
        [Output("proxyHeader")]
        public Output<string?> ProxyHeader { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the resource.
        /// </summary>
        [Output("proxyId")]
        public Output<int> ProxyId { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a TargetTCPProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetTCPProxy(string name, TargetTCPProxyArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/targetTCPProxy:TargetTCPProxy", name, args ?? new TargetTCPProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetTCPProxy(string name, Input<string> id, TargetTCPProxyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/targetTCPProxy:TargetTCPProxy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TargetTCPProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetTCPProxy Get(string name, Input<string> id, TargetTCPProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new TargetTCPProxy(name, id, state, options);
        }
    }

    public sealed class TargetTCPProxyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A reference to the BackendService resource.
        /// </summary>
        [Input("backendService", required: true)]
        public Input<string> BackendService { get; set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// Specifies the type of proxy header to append before sending data to
        /// the backend.
        /// Default value is `NONE`.
        /// Possible values are `NONE` and `PROXY_V1`.
        /// </summary>
        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        public TargetTCPProxyArgs()
        {
        }
    }

    public sealed class TargetTCPProxyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A reference to the BackendService resource.
        /// </summary>
        [Input("backendService")]
        public Input<string>? BackendService { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// Specifies the type of proxy header to append before sending data to
        /// the backend.
        /// Default value is `NONE`.
        /// Possible values are `NONE` and `PROXY_V1`.
        /// </summary>
        [Input("proxyHeader")]
        public Input<string>? ProxyHeader { get; set; }

        /// <summary>
        /// The unique identifier for the resource.
        /// </summary>
        [Input("proxyId")]
        public Input<int>? ProxyId { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public TargetTCPProxyState()
        {
        }
    }
}
