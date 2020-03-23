// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_project.html.markdown.
    /// </summary>
    public partial class Project : Pulumi.CustomResource
    {
        /// <summary>
        /// Create the 'default' network automatically.  Default `true`.
        /// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
        /// will still need to have 1 network slot available to create the project successfully, even if
        /// you set `auto_create_network` to `false`, since the network will exist momentarily.
        /// </summary>
        [Output("autoCreateNetwork")]
        public Output<bool?> AutoCreateNetwork { get; private set; } = null!;

        /// <summary>
        /// The alphanumeric ID of the billing account this project
        /// belongs to. The user or service account performing this operation with the provider
        /// must have Billing Account Administrator privileges (`roles/billing.admin`) in
        /// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
        /// for more details.
        /// </summary>
        [Output("billingAccount")]
        public Output<string?> BillingAccount { get; private set; } = null!;

        /// <summary>
        /// The numeric ID of the folder this project should be
        /// created under. Only one of `org_id` or `folder_id` may be
        /// specified. If the `folder_id` is specified, then the project is
        /// created under the specified folder. Changing this forces the
        /// project to be migrated to the newly specified folder.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the project.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The display name of the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The numeric identifier of the project.
        /// </summary>
        [Output("number")]
        public Output<string> Number { get; private set; } = null!;

        /// <summary>
        /// The numeric ID of the organization this project belongs to.
        /// Changing this forces a new project to be created.  Only one of
        /// `org_id` or `folder_id` may be specified. If the `org_id` is
        /// specified then the project is created at the top level. Changing
        /// this forces the project to be migrated to the newly specified
        /// organization.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// The project ID. Changing this forces a new project to be created.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// If true, the resource can be deleted
        /// without deleting the Project via the Google API.
        /// </summary>
        [Output("skipDelete")]
        public Output<bool> SkipDelete { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("gcp:organizations/project:Project", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("gcp:organizations/project:Project", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create the 'default' network automatically.  Default `true`.
        /// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
        /// will still need to have 1 network slot available to create the project successfully, even if
        /// you set `auto_create_network` to `false`, since the network will exist momentarily.
        /// </summary>
        [Input("autoCreateNetwork")]
        public Input<bool>? AutoCreateNetwork { get; set; }

        /// <summary>
        /// The alphanumeric ID of the billing account this project
        /// belongs to. The user or service account performing this operation with the provider
        /// must have Billing Account Administrator privileges (`roles/billing.admin`) in
        /// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
        /// for more details.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// The numeric ID of the folder this project should be
        /// created under. Only one of `org_id` or `folder_id` may be
        /// specified. If the `folder_id` is specified, then the project is
        /// created under the specified folder. Changing this forces the
        /// project to be migrated to the newly specified folder.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the project.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The display name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The numeric ID of the organization this project belongs to.
        /// Changing this forces a new project to be created.  Only one of
        /// `org_id` or `folder_id` may be specified. If the `org_id` is
        /// specified then the project is created at the top level. Changing
        /// this forces the project to be migrated to the newly specified
        /// organization.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The project ID. Changing this forces a new project to be created.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// If true, the resource can be deleted
        /// without deleting the Project via the Google API.
        /// </summary>
        [Input("skipDelete")]
        public Input<bool>? SkipDelete { get; set; }

        public ProjectArgs()
        {
        }
    }

    public sealed class ProjectState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create the 'default' network automatically.  Default `true`.
        /// If set to `false`, the default network will be deleted.  Note that, for quota purposes, you
        /// will still need to have 1 network slot available to create the project successfully, even if
        /// you set `auto_create_network` to `false`, since the network will exist momentarily.
        /// </summary>
        [Input("autoCreateNetwork")]
        public Input<bool>? AutoCreateNetwork { get; set; }

        /// <summary>
        /// The alphanumeric ID of the billing account this project
        /// belongs to. The user or service account performing this operation with the provider
        /// must have Billing Account Administrator privileges (`roles/billing.admin`) in
        /// the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control)
        /// for more details.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// The numeric ID of the folder this project should be
        /// created under. Only one of `org_id` or `folder_id` may be
        /// specified. If the `folder_id` is specified, then the project is
        /// created under the specified folder. Changing this forces the
        /// project to be migrated to the newly specified folder.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the project.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The display name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The numeric identifier of the project.
        /// </summary>
        [Input("number")]
        public Input<string>? Number { get; set; }

        /// <summary>
        /// The numeric ID of the organization this project belongs to.
        /// Changing this forces a new project to be created.  Only one of
        /// `org_id` or `folder_id` may be specified. If the `org_id` is
        /// specified then the project is created at the top level. Changing
        /// this forces the project to be migrated to the newly specified
        /// organization.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The project ID. Changing this forces a new project to be created.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// If true, the resource can be deleted
        /// without deleting the Project via the Google API.
        /// </summary>
        [Input("skipDelete")]
        public Input<bool>? SkipDelete { get; set; }

        public ProjectState()
        {
        }
    }
}
