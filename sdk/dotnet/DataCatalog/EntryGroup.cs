// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataCatalog
{
    /// <summary>
    /// An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.
    /// 
    /// To get more information about EntryGroup, see:
    /// 
    /// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
    /// 
    /// ## Example Usage
    /// ### Data Catalog Entry Group Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var basicEntryGroup = new Gcp.DataCatalog.EntryGroup("basicEntryGroup", new Gcp.DataCatalog.EntryGroupArgs
    ///         {
    ///             EntryGroupId = "my_group",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Data Catalog Entry Group Full
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var basicEntryGroup = new Gcp.DataCatalog.EntryGroup("basicEntryGroup", new Gcp.DataCatalog.EntryGroupArgs
    ///         {
    ///             Description = "example entry group",
    ///             DisplayName = "entry group",
    ///             EntryGroupId = "my_group",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EntryGroup can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datacatalog/entryGroup:EntryGroup default {{name}}
    /// ```
    /// </summary>
    public partial class EntryGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A short name to identify the entry group, for example, "analytics data - jan 2011".
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The id of the entry group to create. The id must begin with a letter or underscore,
        /// contain only English letters, numbers and underscores, and be at most 64 characters.
        /// </summary>
        [Output("entryGroupId")]
        public Output<string> EntryGroupId { get; private set; } = null!;

        /// <summary>
        /// The resource name of the entry group in URL format. Example:
        /// projects/{project}/locations/{location}/entryGroups/{entryGroupId}
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
        /// EntryGroup location region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a EntryGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EntryGroup(string name, EntryGroupArgs args, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/entryGroup:EntryGroup", name, args ?? new EntryGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EntryGroup(string name, Input<string> id, EntryGroupState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/entryGroup:EntryGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EntryGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EntryGroup Get(string name, Input<string> id, EntryGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new EntryGroup(name, id, state, options);
        }
    }

    public sealed class EntryGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A short name to identify the entry group, for example, "analytics data - jan 2011".
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The id of the entry group to create. The id must begin with a letter or underscore,
        /// contain only English letters, numbers and underscores, and be at most 64 characters.
        /// </summary>
        [Input("entryGroupId", required: true)]
        public Input<string> EntryGroupId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// EntryGroup location region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public EntryGroupArgs()
        {
        }
    }

    public sealed class EntryGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A short name to identify the entry group, for example, "analytics data - jan 2011".
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The id of the entry group to create. The id must begin with a letter or underscore,
        /// contain only English letters, numbers and underscores, and be at most 64 characters.
        /// </summary>
        [Input("entryGroupId")]
        public Input<string>? EntryGroupId { get; set; }

        /// <summary>
        /// The resource name of the entry group in URL format. Example:
        /// projects/{project}/locations/{location}/entryGroups/{entryGroupId}
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
        /// EntryGroup location region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public EntryGroupState()
        {
        }
    }
}
