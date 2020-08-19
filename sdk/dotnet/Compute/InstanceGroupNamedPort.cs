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
    /// Mange the named ports setting for a managed instance group without
    /// managing the group as whole. This resource is primarily intended for use
    /// with GKE-generated groups that shouldn't otherwise be managed by other
    /// tools.
    /// 
    /// 
    /// To get more information about InstanceGroupNamedPort, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroup)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/)
    /// </summary>
    public partial class InstanceGroupNamedPort : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the instance group.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// The name for this named port. The name must be 1-63 characters
        /// long, and comply with RFC1035.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The port number, which can be a value between 1 and 65535.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The zone of the instance group.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceGroupNamedPort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceGroupNamedPort(string name, InstanceGroupNamedPortArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort", name, args ?? new InstanceGroupNamedPortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceGroupNamedPort(string name, Input<string> id, InstanceGroupNamedPortState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceGroupNamedPort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceGroupNamedPort Get(string name, Input<string> id, InstanceGroupNamedPortState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceGroupNamedPort(name, id, state, options);
        }
    }

    public sealed class InstanceGroupNamedPortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the instance group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// The name for this named port. The name must be 1-63 characters
        /// long, and comply with RFC1035.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port number, which can be a value between 1 and 65535.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The zone of the instance group.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceGroupNamedPortArgs()
        {
        }
    }

    public sealed class InstanceGroupNamedPortState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the instance group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// The name for this named port. The name must be 1-63 characters
        /// long, and comply with RFC1035.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port number, which can be a value between 1 and 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The zone of the instance group.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceGroupNamedPortState()
        {
        }
    }
}
