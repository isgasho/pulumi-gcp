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
    /// A VPC network is a virtual version of the traditional physical networks
    /// that exist within and between physical data centers. A VPC network
    /// provides connectivity for your Compute Engine virtual machine (VM)
    /// instances, Container Engine containers, App Engine Flex services, and
    /// other network-related resources.
    /// 
    /// Each GCP project contains one or more VPC networks. Each VPC network is a
    /// global entity spanning all GCP regions. This global VPC network allows VM
    /// instances and other resources to communicate with each other via internal,
    /// private IP addresses.
    /// 
    /// Each VPC network is subdivided into subnets, and each subnet is contained
    /// within a single region. You can have more than one subnet in a region for
    /// a given VPC network. Each subnet has a contiguous private RFC1918 IP
    /// space. You create instances, containers, and the like in these subnets.
    /// When you create an instance, you must create it in a subnet, and the
    /// instance draws its internal IP address from that subnet.
    /// 
    /// Virtual machine (VM) instances in a VPC network can communicate with
    /// instances in all other subnets of the same VPC network, regardless of
    /// region, using their RFC1918 private IP addresses. You can isolate portions
    /// of the network, even entire subnets, using firewall rules.
    /// 
    /// 
    /// To get more information about Subnetwork, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/subnetworks)
    /// * How-to Guides
    ///     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
    ///     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
    /// 
    /// ## Example Usage
    /// 
    /// ### Subnetwork Internal L7lb
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var custom_test = new Gcp.Compute.Network("custom-test", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var network_for_l7lb = new Gcp.Compute.Subnetwork("network-for-l7lb", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             IpCidrRange = "10.0.0.0/22",
    ///             Region = "us-central1",
    ///             Purpose = "INTERNAL_HTTPS_LOAD_BALANCER",
    ///             Role = "ACTIVE",
    ///             Network = custom_test.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Subnetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource. This field can be set only at resource
        /// creation time.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. This field is used internally during updates of this resource.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        [Output("gatewayAddress")]
        public Output<string> GatewayAddress { get; private set; } = null!;

        /// <summary>
        /// The range of IP addresses belonging to this subnetwork secondary
        /// range. Provide this property when you create the subnetwork.
        /// Ranges must be unique and non-overlapping with all primary and
        /// secondary IP ranges within a network. Only IPv4 is supported.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// Denotes the logging options for the subnetwork flow logs. If logging is enabled
        /// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
        /// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
        /// Structure is documented below.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.SubnetworkLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially
        /// creating the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters
        /// long and match the regular expression `a-z?` which
        /// means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network this subnet belongs to.
        /// Only networks that are in the distributed mode can have subnetworks.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// When enabled, VMs in this subnetwork without external IP addresses can
        /// access Google APIs and services by using Private Google Access.
        /// </summary>
        [Output("privateIpGoogleAccess")]
        public Output<bool?> PrivateIpGoogleAccess { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE
        /// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
        /// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
        /// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
        /// purpose defaults to PRIVATE.
        /// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        /// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

        /// <summary>
        /// The GCP region for this subnetwork.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when
        /// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
        /// or BACKUP. An ACTIVE subnetwork is one that is currently being used
        /// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
        /// is ready to be promoted to ACTIVE or is currently draining.
        /// Possible values are `ACTIVE` and `BACKUP`.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances
        /// contained in this subnetwork. The primary IP of such VM must belong
        /// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
        /// to either primary or secondary ranges. Structure is documented below.
        /// </summary>
        [Output("secondaryIpRanges")]
        public Output<ImmutableArray<Outputs.SubnetworkSecondaryIpRange>> SecondaryIpRanges { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Subnetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnetwork(string name, SubnetworkArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/subnetwork:Subnetwork", name, args ?? new SubnetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnetwork(string name, Input<string> id, SubnetworkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/subnetwork:Subnetwork", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subnetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnetwork Get(string name, Input<string> id, SubnetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Subnetwork(name, id, state, options);
        }
    }

    public sealed class SubnetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource. This field can be set only at resource
        /// creation time.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The range of IP addresses belonging to this subnetwork secondary
        /// range. Provide this property when you create the subnetwork.
        /// Ranges must be unique and non-overlapping with all primary and
        /// secondary IP ranges within a network. Only IPv4 is supported.
        /// </summary>
        [Input("ipCidrRange", required: true)]
        public Input<string> IpCidrRange { get; set; } = null!;

        /// <summary>
        /// Denotes the logging options for the subnetwork flow logs. If logging is enabled
        /// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
        /// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
        /// Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.SubnetworkLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially
        /// creating the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters
        /// long and match the regular expression `a-z?` which
        /// means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this subnet belongs to.
        /// Only networks that are in the distributed mode can have subnetworks.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// When enabled, VMs in this subnetwork without external IP addresses can
        /// access Google APIs and services by using Private Google Access.
        /// </summary>
        [Input("privateIpGoogleAccess")]
        public Input<bool>? PrivateIpGoogleAccess { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE
        /// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
        /// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
        /// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
        /// purpose defaults to PRIVATE.
        /// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        /// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// The GCP region for this subnetwork.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when
        /// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
        /// or BACKUP. An ACTIVE subnetwork is one that is currently being used
        /// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
        /// is ready to be promoted to ACTIVE or is currently draining.
        /// Possible values are `ACTIVE` and `BACKUP`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("secondaryIpRanges")]
        private InputList<Inputs.SubnetworkSecondaryIpRangeArgs>? _secondaryIpRanges;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances
        /// contained in this subnetwork. The primary IP of such VM must belong
        /// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
        /// to either primary or secondary ranges. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SubnetworkSecondaryIpRangeArgs> SecondaryIpRanges
        {
            get => _secondaryIpRanges ?? (_secondaryIpRanges = new InputList<Inputs.SubnetworkSecondaryIpRangeArgs>());
            set => _secondaryIpRanges = value;
        }

        public SubnetworkArgs()
        {
        }
    }

    public sealed class SubnetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource. This field can be set only at resource
        /// creation time.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Fingerprint of this resource. This field is used internally during updates of this resource.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        [Input("gatewayAddress")]
        public Input<string>? GatewayAddress { get; set; }

        /// <summary>
        /// The range of IP addresses belonging to this subnetwork secondary
        /// range. Provide this property when you create the subnetwork.
        /// Ranges must be unique and non-overlapping with all primary and
        /// secondary IP ranges within a network. Only IPv4 is supported.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// Denotes the logging options for the subnetwork flow logs. If logging is enabled
        /// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
        /// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
        /// Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.SubnetworkLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially
        /// creating the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters
        /// long and match the regular expression `a-z?` which
        /// means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this subnet belongs to.
        /// Only networks that are in the distributed mode can have subnetworks.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// When enabled, VMs in this subnetwork without external IP addresses can
        /// access Google APIs and services by using Private Google Access.
        /// </summary>
        [Input("privateIpGoogleAccess")]
        public Input<bool>? PrivateIpGoogleAccess { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE
        /// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
        /// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
        /// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
        /// purpose defaults to PRIVATE.
        /// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        /// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// The GCP region for this subnetwork.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when
        /// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
        /// or BACKUP. An ACTIVE subnetwork is one that is currently being used
        /// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
        /// is ready to be promoted to ACTIVE or is currently draining.
        /// Possible values are `ACTIVE` and `BACKUP`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("secondaryIpRanges")]
        private InputList<Inputs.SubnetworkSecondaryIpRangeGetArgs>? _secondaryIpRanges;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances
        /// contained in this subnetwork. The primary IP of such VM must belong
        /// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
        /// to either primary or secondary ranges. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SubnetworkSecondaryIpRangeGetArgs> SecondaryIpRanges
        {
            get => _secondaryIpRanges ?? (_secondaryIpRanges = new InputList<Inputs.SubnetworkSecondaryIpRangeGetArgs>());
            set => _secondaryIpRanges = value;
        }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public SubnetworkState()
        {
        }
    }
}
