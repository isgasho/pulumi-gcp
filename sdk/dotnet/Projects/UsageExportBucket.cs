// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects
{
    /// <summary>
    /// Allows creation and management of a Google Cloud Platform project.
    /// 
    /// Projects created with this resource must be associated with an Organization.
    /// See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
    /// 
    /// The service account used to run this provider when creating a `gcp.organizations.Project`
    /// resource must have `roles/resourcemanager.projectCreator`. See the
    /// [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
    /// doc for more information.
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
    ///         var myProject = new Gcp.Organizations.Project("myProject", new Gcp.Organizations.ProjectArgs
    ///         {
    ///             OrgId = "1234567",
    ///             ProjectId = "your-project-id",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// To create a project under a specific folder
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var department1 = new Gcp.Organizations.Folder("department1", new Gcp.Organizations.FolderArgs
    ///         {
    ///             DisplayName = "Department 1",
    ///             Parent = "organizations/1234567",
    ///         });
    ///         var myProject_in_a_folder = new Gcp.Organizations.Project("myProject-in-a-folder", new Gcp.Organizations.ProjectArgs
    ///         {
    ///             ProjectId = "your-project-id",
    ///             FolderId = department1.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Projects can be imported using the `project_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
    /// ```
    /// </summary>
    public partial class UsageExportBucket : Pulumi.CustomResource
    {
        /// <summary>
        /// The bucket to store reports in.
        /// </summary>
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        /// <summary>
        /// A prefix for the reports, for instance, the project name.
        /// </summary>
        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;

        /// <summary>
        /// The project to set the export bucket on. If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a UsageExportBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsageExportBucket(string name, UsageExportBucketArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/usageExportBucket:UsageExportBucket", name, args ?? new UsageExportBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UsageExportBucket(string name, Input<string> id, UsageExportBucketState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/usageExportBucket:UsageExportBucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UsageExportBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsageExportBucket Get(string name, Input<string> id, UsageExportBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new UsageExportBucket(name, id, state, options);
        }
    }

    public sealed class UsageExportBucketArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket to store reports in.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// A prefix for the reports, for instance, the project name.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The project to set the export bucket on. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public UsageExportBucketArgs()
        {
        }
    }

    public sealed class UsageExportBucketState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket to store reports in.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// A prefix for the reports, for instance, the project name.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The project to set the export bucket on. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public UsageExportBucketState()
        {
        }
    }
}
