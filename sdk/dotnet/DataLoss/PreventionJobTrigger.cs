// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss
{
    /// <summary>
    /// A job trigger configuration.
    /// 
    /// To get more information about JobTrigger, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.jobTriggers)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-job-triggers)
    /// 
    /// ## Example Usage
    /// ### Dlp Job Trigger Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var basic = new Gcp.DataLoss.PreventionJobTrigger("basic", new Gcp.DataLoss.PreventionJobTriggerArgs
    ///         {
    ///             Description = "Description",
    ///             DisplayName = "Displayname",
    ///             InspectJob = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobArgs
    ///             {
    ///                 Actions = 
    ///                 {
    ///                     new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobActionArgs
    ///                     {
    ///                         SaveFindings = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobActionSaveFindingsArgs
    ///                         {
    ///                             OutputConfig = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobActionSaveFindingsOutputConfigArgs
    ///                             {
    ///                                 Table = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobActionSaveFindingsOutputConfigTableArgs
    ///                                 {
    ///                                     DatasetId = "asdf",
    ///                                     ProjectId = "asdf",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 InspectTemplateName = "fake",
    ///                 StorageConfig = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobStorageConfigArgs
    ///                 {
    ///                     CloudStorageOptions = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsArgs
    ///                     {
    ///                         FileSet = new Gcp.DataLoss.Inputs.PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetArgs
    ///                         {
    ///                             Url = "gs://mybucket/directory/",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Parent = "projects/my-project-name",
    ///             Triggers = 
    ///             {
    ///                 new Gcp.DataLoss.Inputs.PreventionJobTriggerTriggerArgs
    ///                 {
    ///                     Schedule = new Gcp.DataLoss.Inputs.PreventionJobTriggerTriggerScheduleArgs
    ///                     {
    ///                         RecurrencePeriodDuration = "86400s",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// JobTrigger can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataloss/preventionJobTrigger:PreventionJobTrigger default {{parent}}/jobTriggers/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataloss/preventionJobTrigger:PreventionJobTrigger default {{parent}}/{{name}}
    /// ```
    /// </summary>
    public partial class PreventionJobTrigger : Pulumi.CustomResource
    {
        /// <summary>
        /// A description of the job trigger.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User set display name of the job trigger.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Controls what and how to inspect for findings.
        /// Structure is documented below.
        /// </summary>
        [Output("inspectJob")]
        public Output<Outputs.PreventionJobTriggerInspectJob?> InspectJob { get; private set; } = null!;

        /// <summary>
        /// The timestamp of the last time this trigger executed.
        /// </summary>
        [Output("lastRunTime")]
        public Output<string> LastRunTime { get; private set; } = null!;

        /// <summary>
        /// The name of the Datastore kind.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent of the trigger, either in the format `projects/{{project}}`
        /// or `projects/{{project}}/locations/{{location}}`
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Whether the trigger is currently active.
        /// Default value is `HEALTHY`.
        /// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// What event needs to occur for a new job to be started.
        /// Structure is documented below.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<Outputs.PreventionJobTriggerTrigger>> Triggers { get; private set; } = null!;


        /// <summary>
        /// Create a PreventionJobTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PreventionJobTrigger(string name, PreventionJobTriggerArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataloss/preventionJobTrigger:PreventionJobTrigger", name, args ?? new PreventionJobTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PreventionJobTrigger(string name, Input<string> id, PreventionJobTriggerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataloss/preventionJobTrigger:PreventionJobTrigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PreventionJobTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PreventionJobTrigger Get(string name, Input<string> id, PreventionJobTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new PreventionJobTrigger(name, id, state, options);
        }
    }

    public sealed class PreventionJobTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the job trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User set display name of the job trigger.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Controls what and how to inspect for findings.
        /// Structure is documented below.
        /// </summary>
        [Input("inspectJob")]
        public Input<Inputs.PreventionJobTriggerInspectJobArgs>? InspectJob { get; set; }

        /// <summary>
        /// The parent of the trigger, either in the format `projects/{{project}}`
        /// or `projects/{{project}}/locations/{{location}}`
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Whether the trigger is currently active.
        /// Default value is `HEALTHY`.
        /// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("triggers", required: true)]
        private InputList<Inputs.PreventionJobTriggerTriggerArgs>? _triggers;

        /// <summary>
        /// What event needs to occur for a new job to be started.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PreventionJobTriggerTriggerArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.PreventionJobTriggerTriggerArgs>());
            set => _triggers = value;
        }

        public PreventionJobTriggerArgs()
        {
        }
    }

    public sealed class PreventionJobTriggerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the job trigger.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User set display name of the job trigger.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Controls what and how to inspect for findings.
        /// Structure is documented below.
        /// </summary>
        [Input("inspectJob")]
        public Input<Inputs.PreventionJobTriggerInspectJobGetArgs>? InspectJob { get; set; }

        /// <summary>
        /// The timestamp of the last time this trigger executed.
        /// </summary>
        [Input("lastRunTime")]
        public Input<string>? LastRunTime { get; set; }

        /// <summary>
        /// The name of the Datastore kind.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of the trigger, either in the format `projects/{{project}}`
        /// or `projects/{{project}}/locations/{{location}}`
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Whether the trigger is currently active.
        /// Default value is `HEALTHY`.
        /// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("triggers")]
        private InputList<Inputs.PreventionJobTriggerTriggerGetArgs>? _triggers;

        /// <summary>
        /// What event needs to occur for a new job to be started.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PreventionJobTriggerTriggerGetArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.PreventionJobTriggerTriggerGetArgs>());
            set => _triggers = value;
        }

        public PreventionJobTriggerState()
        {
        }
    }
}
