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
    /// Represents a SSL policy. SSL policies give you the ability to control the
    /// features of SSL that your SSL proxy or HTTPS load balancer negotiates.
    /// 
    /// To get more information about SslPolicy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslPolicies)
    /// * How-to Guides
    ///     * [Using SSL Policies](https://cloud.google.com/compute/docs/load-balancing/ssl-policies)
    /// 
    /// ## Example Usage
    /// ### Ssl Policy Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var prod_ssl_policy = new Gcp.Compute.SSLPolicy("prod-ssl-policy", new Gcp.Compute.SSLPolicyArgs
    ///         {
    ///             Profile = "MODERN",
    ///         });
    ///         var nonprod_ssl_policy = new Gcp.Compute.SSLPolicy("nonprod-ssl-policy", new Gcp.Compute.SSLPolicyArgs
    ///         {
    ///             MinTlsVersion = "TLS_1_2",
    ///             Profile = "MODERN",
    ///         });
    ///         var custom_ssl_policy = new Gcp.Compute.SSLPolicy("custom-ssl-policy", new Gcp.Compute.SSLPolicyArgs
    ///         {
    ///             CustomFeatures = 
    ///             {
    ///                 "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
    ///                 "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
    ///             },
    ///             MinTlsVersion = "TLS_1_2",
    ///             Profile = "CUSTOM",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SslPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default projects/{{project}}/global/sslPolicies/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/sSLPolicy:SSLPolicy default {{name}}
    /// ```
    /// </summary>
    public partial class SSLPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Profile specifies the set of SSL features that can be used by the
        /// load balancer when negotiating SSL with clients. This can be one of
        /// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
        /// the set of SSL features to enable must be specified in the
        /// `customFeatures` field.
        /// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        /// for which ciphers are available to use. **Note**: this argument
        /// *must* be present when using the `CUSTOM` profile. This argument
        /// *must not* be present when using any other profile.
        /// </summary>
        [Output("customFeatures")]
        public Output<ImmutableArray<string>> CustomFeatures { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The list of features enabled in the SSL policy.
        /// </summary>
        [Output("enabledFeatures")]
        public Output<ImmutableArray<string>> EnabledFeatures { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The minimum version of SSL protocol that can be used by the clients
        /// to establish a connection with the load balancer.
        /// Default value is `TLS_1_0`.
        /// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        /// </summary>
        [Output("minTlsVersion")]
        public Output<string?> MinTlsVersion { get; private set; } = null!;

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
        /// Profile specifies the set of SSL features that can be used by the
        /// load balancer when negotiating SSL with clients. If using `CUSTOM`,
        /// the set of SSL features to enable must be specified in the
        /// `customFeatures` field.
        /// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        /// for information on what cipher suites each profile provides. If
        /// `CUSTOM` is used, the `custom_features` attribute **must be set**.
        /// Default value is `COMPATIBLE`.
        /// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a SSLPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SSLPolicy(string name, SSLPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/sSLPolicy:SSLPolicy", name, args ?? new SSLPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SSLPolicy(string name, Input<string> id, SSLPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/sSLPolicy:SSLPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SSLPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SSLPolicy Get(string name, Input<string> id, SSLPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SSLPolicy(name, id, state, options);
        }
    }

    public sealed class SSLPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("customFeatures")]
        private InputList<string>? _customFeatures;

        /// <summary>
        /// Profile specifies the set of SSL features that can be used by the
        /// load balancer when negotiating SSL with clients. This can be one of
        /// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
        /// the set of SSL features to enable must be specified in the
        /// `customFeatures` field.
        /// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        /// for which ciphers are available to use. **Note**: this argument
        /// *must* be present when using the `CUSTOM` profile. This argument
        /// *must not* be present when using any other profile.
        /// </summary>
        public InputList<string> CustomFeatures
        {
            get => _customFeatures ?? (_customFeatures = new InputList<string>());
            set => _customFeatures = value;
        }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The minimum version of SSL protocol that can be used by the clients
        /// to establish a connection with the load balancer.
        /// Default value is `TLS_1_0`.
        /// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        /// </summary>
        [Input("minTlsVersion")]
        public Input<string>? MinTlsVersion { get; set; }

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
        /// Profile specifies the set of SSL features that can be used by the
        /// load balancer when negotiating SSL with clients. If using `CUSTOM`,
        /// the set of SSL features to enable must be specified in the
        /// `customFeatures` field.
        /// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        /// for information on what cipher suites each profile provides. If
        /// `CUSTOM` is used, the `custom_features` attribute **must be set**.
        /// Default value is `COMPATIBLE`.
        /// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public SSLPolicyArgs()
        {
        }
    }

    public sealed class SSLPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("customFeatures")]
        private InputList<string>? _customFeatures;

        /// <summary>
        /// Profile specifies the set of SSL features that can be used by the
        /// load balancer when negotiating SSL with clients. This can be one of
        /// `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
        /// the set of SSL features to enable must be specified in the
        /// `customFeatures` field.
        /// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        /// for which ciphers are available to use. **Note**: this argument
        /// *must* be present when using the `CUSTOM` profile. This argument
        /// *must not* be present when using any other profile.
        /// </summary>
        public InputList<string> CustomFeatures
        {
            get => _customFeatures ?? (_customFeatures = new InputList<string>());
            set => _customFeatures = value;
        }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabledFeatures")]
        private InputList<string>? _enabledFeatures;

        /// <summary>
        /// The list of features enabled in the SSL policy.
        /// </summary>
        public InputList<string> EnabledFeatures
        {
            get => _enabledFeatures ?? (_enabledFeatures = new InputList<string>());
            set => _enabledFeatures = value;
        }

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The minimum version of SSL protocol that can be used by the clients
        /// to establish a connection with the load balancer.
        /// Default value is `TLS_1_0`.
        /// Possible values are `TLS_1_0`, `TLS_1_1`, and `TLS_1_2`.
        /// </summary>
        [Input("minTlsVersion")]
        public Input<string>? MinTlsVersion { get; set; }

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
        /// Profile specifies the set of SSL features that can be used by the
        /// load balancer when negotiating SSL with clients. If using `CUSTOM`,
        /// the set of SSL features to enable must be specified in the
        /// `customFeatures` field.
        /// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
        /// for information on what cipher suites each profile provides. If
        /// `CUSTOM` is used, the `custom_features` attribute **must be set**.
        /// Default value is `COMPATIBLE`.
        /// Possible values are `COMPATIBLE`, `MODERN`, `RESTRICTED`, and `CUSTOM`.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public SSLPolicyState()
        {
        }
    }
}
