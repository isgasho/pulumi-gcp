// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigTable
{
    /// <summary>
    /// Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:
    /// 
    /// * `gcp.bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
    /// * `gcp.bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
    /// * `gcp.bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
    /// 
    /// &gt; **Note:** `gcp.bigtable.TableIamPolicy` **cannot** be used in conjunction with `gcp.bigtable.TableIamBinding` and `gcp.bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `gcp.bigtable.TableIamPolicy` replaces the entire policy.
    /// 
    /// &gt; **Note:** `gcp.bigtable.TableIamBinding` resources **can be** used in conjunction with `gcp.bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_bigtable\_instance\_iam\_policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var admin = Output.Create(Gcp.Organizations.GetIAMPolicy.InvokeAsync(new Gcp.Organizations.GetIAMPolicyArgs
    ///         {
    ///             Bindings = 
    ///             {
    ///                 new Gcp.Organizations.Inputs.GetIAMPolicyBindingArgs
    ///                 {
    ///                     Role = "roles/bigtable.user",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var editor = new Gcp.BigTable.TableIamPolicy("editor", new Gcp.BigTable.TableIamPolicyArgs
    ///         {
    ///             Project = "your-project",
    ///             Instance = "your-bigtable-instance",
    ///             Table = "your-bigtable-table",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_bigtable\_instance\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var editor = new Gcp.BigTable.TableIamBinding("editor", new Gcp.BigTable.TableIamBindingArgs
    ///         {
    ///             Instance = "your-bigtable-instance",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Role = "roles/bigtable.user",
    ///             Table = "your-bigtable-table",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_bigtable\_instance\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var editor = new Gcp.BigTable.TableIamMember("editor", new Gcp.BigTable.TableIamMemberArgs
    ///         {
    ///             Instance = "your-bigtable-instance",
    ///             Member = "user:jane@example.com",
    ///             Role = "roles/bigtable.user",
    ///             Table = "your-bigtable-table",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance IAM resources can be imported using the project, table name, role and/or member.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/tableIamPolicy:TableIamPolicy editor "projects/{project}/tables/{table}"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/tableIamPolicy:TableIamPolicy editor "projects/{project}/tables/{table} roles/editor"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/tableIamPolicy:TableIamPolicy editor "projects/{project}/tables/{table} roles/editor user:jane@example.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    public partial class TableIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the tables's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name or relative resource id of the instance that owns the table.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name or relative resource id of the table to manage IAM policies for.
        /// </summary>
        [Output("table")]
        public Output<string> Table { get; private set; } = null!;


        /// <summary>
        /// Create a TableIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TableIamPolicy(string name, TableIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigtable/tableIamPolicy:TableIamPolicy", name, args ?? new TableIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TableIamPolicy(string name, Input<string> id, TableIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/tableIamPolicy:TableIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TableIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TableIamPolicy Get(string name, Input<string> id, TableIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new TableIamPolicy(name, id, state, options);
        }
    }

    public sealed class TableIamPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or relative resource id of the instance that owns the table.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name or relative resource id of the table to manage IAM policies for.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public TableIamPolicyArgs()
        {
        }
    }

    public sealed class TableIamPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the tables's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name or relative resource id of the instance that owns the table.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name or relative resource id of the table to manage IAM policies for.
        /// </summary>
        [Input("table")]
        public Input<string>? Table { get; set; }

        public TableIamPolicyState()
        {
        }
    }
}
