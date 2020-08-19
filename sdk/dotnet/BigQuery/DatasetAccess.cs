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
    /// Gives dataset access for a single entity. This resource is intended to be used in cases where
    /// it is not possible to compile a full list of access blocks to include in a
    /// `gcp.bigquery.Dataset` resource, to enable them to be added separately.
    /// 
    /// &gt; **Note:** If this resource is used alongside a `gcp.bigquery.Dataset` resource, the
    /// dataset resource must either have no defined `access` blocks or a `lifecycle` block with
    /// `ignore_changes = [access]` so they don't fight over which accesses should be on the dataset.
    /// 
    /// 
    /// To get more information about DatasetAccess, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets)
    /// * How-to Guides
    ///     * [Controlling access to datasets](https://cloud.google.com/bigquery/docs/dataset-access-controls)
    /// 
    /// ## Example Usage
    /// 
    /// ### Bigquery Dataset Access Basic User
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dataset = new Gcp.BigQuery.Dataset("dataset", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "example_dataset",
    ///         });
    ///         var bqowner = new Gcp.ServiceAccount.Account("bqowner", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "bqowner",
    ///         });
    ///         var access = new Gcp.BigQuery.DatasetAccess("access", new Gcp.BigQuery.DatasetAccessArgs
    ///         {
    ///             DatasetId = dataset.DatasetId,
    ///             Role = "OWNER",
    ///             UserByEmail = bqowner.Email,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Bigquery Dataset Access View
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @private = new Gcp.BigQuery.Dataset("private", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "example_dataset",
    ///         });
    ///         var publicDataset = new Gcp.BigQuery.Dataset("publicDataset", new Gcp.BigQuery.DatasetArgs
    ///         {
    ///             DatasetId = "example_dataset2",
    ///         });
    ///         var publicTable = new Gcp.BigQuery.Table("publicTable", new Gcp.BigQuery.TableArgs
    ///         {
    ///             DatasetId = publicDataset.DatasetId,
    ///             TableId = "example_table",
    ///             View = new Gcp.BigQuery.Inputs.TableViewArgs
    ///             {
    ///                 Query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
    ///                 UseLegacySql = false,
    ///             },
    ///         });
    ///         var access = new Gcp.BigQuery.DatasetAccess("access", new Gcp.BigQuery.DatasetAccessArgs
    ///         {
    ///             DatasetId = @private.DatasetId,
    ///             View = new Gcp.BigQuery.Inputs.DatasetAccessViewArgs
    ///             {
    ///                 ProjectId = publicTable.Project,
    ///                 DatasetId = publicDataset.DatasetId,
    ///                 TableId = publicTable.TableId,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DatasetAccess : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the dataset containing this table.
        /// </summary>
        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// A domain to grant access to. Any users signed in with the
        /// domain specified will be granted the specified access
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// An email address of a Google Group to grant access to.
        /// </summary>
        [Output("groupByEmail")]
        public Output<string?> GroupByEmail { get; private set; } = null!;

        /// <summary>
        /// Some other type of member that appears in the IAM Policy but isn't a user,
        /// group, domain, or special group. For example: `allUsers`
        /// </summary>
        [Output("iamMember")]
        public Output<string?> IamMember { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Describes the rights granted to the user specified by the other
        /// member of the access object. Primitive, Predefined and custom
        /// roles are supported. Predefined roles that have equivalent
        /// primitive roles are swapped by the API to their Primitive
        /// counterparts, and will show a diff post-create. See
        /// [official docs](https://cloud.google.com/bigquery/docs/access-control).
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// A special group to grant access to. Possible values include:
        /// </summary>
        [Output("specialGroup")]
        public Output<string?> SpecialGroup { get; private set; } = null!;

        /// <summary>
        /// An email address of a user to grant access to. For example:
        /// fred@example.com
        /// </summary>
        [Output("userByEmail")]
        public Output<string?> UserByEmail { get; private set; } = null!;

        /// <summary>
        /// A view from a different dataset to grant access to. Queries
        /// executed against that view will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that view is updated by any user, access to the view
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Output("view")]
        public Output<Outputs.DatasetAccessView?> View { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetAccess(string name, DatasetAccessArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/datasetAccess:DatasetAccess", name, args ?? new DatasetAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetAccess(string name, Input<string> id, DatasetAccessState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/datasetAccess:DatasetAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetAccess Get(string name, Input<string> id, DatasetAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetAccess(name, id, state, options);
        }
    }

    public sealed class DatasetAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dataset containing this table.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// A domain to grant access to. Any users signed in with the
        /// domain specified will be granted the specified access
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// An email address of a Google Group to grant access to.
        /// </summary>
        [Input("groupByEmail")]
        public Input<string>? GroupByEmail { get; set; }

        /// <summary>
        /// Some other type of member that appears in the IAM Policy but isn't a user,
        /// group, domain, or special group. For example: `allUsers`
        /// </summary>
        [Input("iamMember")]
        public Input<string>? IamMember { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes the rights granted to the user specified by the other
        /// member of the access object. Primitive, Predefined and custom
        /// roles are supported. Predefined roles that have equivalent
        /// primitive roles are swapped by the API to their Primitive
        /// counterparts, and will show a diff post-create. See
        /// [official docs](https://cloud.google.com/bigquery/docs/access-control).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// A special group to grant access to. Possible values include:
        /// </summary>
        [Input("specialGroup")]
        public Input<string>? SpecialGroup { get; set; }

        /// <summary>
        /// An email address of a user to grant access to. For example:
        /// fred@example.com
        /// </summary>
        [Input("userByEmail")]
        public Input<string>? UserByEmail { get; set; }

        /// <summary>
        /// A view from a different dataset to grant access to. Queries
        /// executed against that view will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that view is updated by any user, access to the view
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Input("view")]
        public Input<Inputs.DatasetAccessViewArgs>? View { get; set; }

        public DatasetAccessArgs()
        {
        }
    }

    public sealed class DatasetAccessState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the dataset containing this table.
        /// </summary>
        [Input("datasetId")]
        public Input<string>? DatasetId { get; set; }

        /// <summary>
        /// A domain to grant access to. Any users signed in with the
        /// domain specified will be granted the specified access
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// An email address of a Google Group to grant access to.
        /// </summary>
        [Input("groupByEmail")]
        public Input<string>? GroupByEmail { get; set; }

        /// <summary>
        /// Some other type of member that appears in the IAM Policy but isn't a user,
        /// group, domain, or special group. For example: `allUsers`
        /// </summary>
        [Input("iamMember")]
        public Input<string>? IamMember { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes the rights granted to the user specified by the other
        /// member of the access object. Primitive, Predefined and custom
        /// roles are supported. Predefined roles that have equivalent
        /// primitive roles are swapped by the API to their Primitive
        /// counterparts, and will show a diff post-create. See
        /// [official docs](https://cloud.google.com/bigquery/docs/access-control).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// A special group to grant access to. Possible values include:
        /// </summary>
        [Input("specialGroup")]
        public Input<string>? SpecialGroup { get; set; }

        /// <summary>
        /// An email address of a user to grant access to. For example:
        /// fred@example.com
        /// </summary>
        [Input("userByEmail")]
        public Input<string>? UserByEmail { get; set; }

        /// <summary>
        /// A view from a different dataset to grant access to. Queries
        /// executed against that view will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that view is updated by any user, access to the view
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Input("view")]
        public Input<Inputs.DatasetAccessViewGetArgs>? View { get; set; }

        public DatasetAccessState()
        {
        }
    }
}
