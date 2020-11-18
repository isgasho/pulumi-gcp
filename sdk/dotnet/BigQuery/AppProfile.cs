// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// App profile is a configuration object describing how Cloud Bigtable should treat traffic from a particular end user application.
    /// 
    /// ## Example Usage
    /// ### Bigtable App Profile Multicluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var instance = new Gcp.BigTable.Instance("instance", new Gcp.BigTable.InstanceArgs
    ///         {
    ///             Clusters = 
    ///             {
    ///                 new Gcp.BigTable.Inputs.InstanceClusterArgs
    ///                 {
    ///                     ClusterId = "bt-instance",
    ///                     Zone = "us-central1-b",
    ///                     NumNodes = 3,
    ///                     StorageType = "HDD",
    ///                 },
    ///             },
    ///             DeletionProtection = true,
    ///         });
    ///         var ap = new Gcp.BigQuery.AppProfile("ap", new Gcp.BigQuery.AppProfileArgs
    ///         {
    ///             Instance = instance.Name,
    ///             AppProfileId = "bt-profile",
    ///             MultiClusterRoutingUseAny = true,
    ///             IgnoreWarnings = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Bigtable App Profile Singlecluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var instance = new Gcp.BigTable.Instance("instance", new Gcp.BigTable.InstanceArgs
    ///         {
    ///             Clusters = 
    ///             {
    ///                 new Gcp.BigTable.Inputs.InstanceClusterArgs
    ///                 {
    ///                     ClusterId = "bt-instance",
    ///                     Zone = "us-central1-b",
    ///                     NumNodes = 3,
    ///                     StorageType = "HDD",
    ///                 },
    ///             },
    ///             DeletionProtection = true,
    ///         });
    ///         var ap = new Gcp.BigQuery.AppProfile("ap", new Gcp.BigQuery.AppProfileArgs
    ///         {
    ///             Instance = instance.Name,
    ///             AppProfileId = "bt-profile",
    ///             SingleClusterRouting = new Gcp.BigQuery.Inputs.AppProfileSingleClusterRoutingArgs
    ///             {
    ///                 ClusterId = "bt-instance",
    ///                 AllowTransactionalWrites = true,
    ///             },
    ///             IgnoreWarnings = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AppProfile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/appProfile:AppProfile default projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/appProfile:AppProfile default {{project}}/{{instance}}/{{app_profile_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/appProfile:AppProfile default {{instance}}/{{app_profile_id}}
    /// ```
    /// </summary>
    public partial class AppProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        /// </summary>
        [Output("appProfileId")]
        public Output<string> AppProfileId { get; private set; } = null!;

        /// <summary>
        /// Long form description of the use case for this app profile.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If true, ignore safety checks when deleting/updating the app profile.
        /// </summary>
        [Output("ignoreWarnings")]
        public Output<bool?> IgnoreWarnings { get; private set; } = null!;

        /// <summary>
        /// The name of the instance to create the app profile within.
        /// </summary>
        [Output("instance")]
        public Output<string?> Instance { get; private set; } = null!;

        /// <summary>
        /// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
        /// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
        /// consistency to improve availability.
        /// </summary>
        [Output("multiClusterRoutingUseAny")]
        public Output<bool?> MultiClusterRoutingUseAny { get; private set; } = null!;

        /// <summary>
        /// The unique name of the requested app profile. Values are of the form
        /// 'projects/&lt;project&gt;/instances/&lt;instance&gt;/appProfiles/&lt;appProfileId&gt;'.
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
        /// Use a single-cluster routing policy.
        /// Structure is documented below.
        /// </summary>
        [Output("singleClusterRouting")]
        public Output<Outputs.AppProfileSingleClusterRouting?> SingleClusterRouting { get; private set; } = null!;


        /// <summary>
        /// Create a AppProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppProfile(string name, AppProfileArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/appProfile:AppProfile", name, args ?? new AppProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppProfile(string name, Input<string> id, AppProfileState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/appProfile:AppProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppProfile Get(string name, Input<string> id, AppProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new AppProfile(name, id, state, options);
        }
    }

    public sealed class AppProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        /// </summary>
        [Input("appProfileId", required: true)]
        public Input<string> AppProfileId { get; set; } = null!;

        /// <summary>
        /// Long form description of the use case for this app profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If true, ignore safety checks when deleting/updating the app profile.
        /// </summary>
        [Input("ignoreWarnings")]
        public Input<bool>? IgnoreWarnings { get; set; }

        /// <summary>
        /// The name of the instance to create the app profile within.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
        /// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
        /// consistency to improve availability.
        /// </summary>
        [Input("multiClusterRoutingUseAny")]
        public Input<bool>? MultiClusterRoutingUseAny { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Use a single-cluster routing policy.
        /// Structure is documented below.
        /// </summary>
        [Input("singleClusterRouting")]
        public Input<Inputs.AppProfileSingleClusterRoutingArgs>? SingleClusterRouting { get; set; }

        public AppProfileArgs()
        {
        }
    }

    public sealed class AppProfileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        /// </summary>
        [Input("appProfileId")]
        public Input<string>? AppProfileId { get; set; }

        /// <summary>
        /// Long form description of the use case for this app profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If true, ignore safety checks when deleting/updating the app profile.
        /// </summary>
        [Input("ignoreWarnings")]
        public Input<bool>? IgnoreWarnings { get; set; }

        /// <summary>
        /// The name of the instance to create the app profile within.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
        /// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
        /// consistency to improve availability.
        /// </summary>
        [Input("multiClusterRoutingUseAny")]
        public Input<bool>? MultiClusterRoutingUseAny { get; set; }

        /// <summary>
        /// The unique name of the requested app profile. Values are of the form
        /// 'projects/&lt;project&gt;/instances/&lt;instance&gt;/appProfiles/&lt;appProfileId&gt;'.
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
        /// Use a single-cluster routing policy.
        /// Structure is documented below.
        /// </summary>
        [Input("singleClusterRouting")]
        public Input<Inputs.AppProfileSingleClusterRoutingGetArgs>? SingleClusterRouting { get; set; }

        public AppProfileState()
        {
        }
    }
}
