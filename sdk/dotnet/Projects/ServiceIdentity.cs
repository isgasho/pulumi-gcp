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
    /// Generate service identity for a service.
    /// 
    /// &gt; **Note**: Once created, this resource cannot be updated or destroyed. These
    /// actions are a no-op.
    /// 
    /// To get more information about Service Identity, see:
    /// 
    /// * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1beta1/services/generateServiceIdentity)
    /// 
    /// ## Example Usage
    /// ### Service Identity Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = Output.Create(Gcp.Organizations.GetProject.InvokeAsync());
    ///         var hcSa = new Gcp.Projects.ServiceIdentity("hcSa", new Gcp.Projects.ServiceIdentityArgs
    ///         {
    ///             Project = project.Apply(project =&gt; project.ProjectId),
    ///             Service = "healthcare.googleapis.com",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var hcSaBqJobuser = new Gcp.Projects.IAMMember("hcSaBqJobuser", new Gcp.Projects.IAMMemberArgs
    ///         {
    ///             Project = project.Apply(project =&gt; project.ProjectId),
    ///             Role = "roles/bigquery.jobUser",
    ///             Member = hcSa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    public partial class ServiceIdentity : Pulumi.CustomResource
    {
        /// <summary>
        /// The email address of the Google managed service account.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The service to generate identity for.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceIdentity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceIdentity(string name, ServiceIdentityArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/serviceIdentity:ServiceIdentity", name, args ?? new ServiceIdentityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceIdentity(string name, Input<string> id, ServiceIdentityState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/serviceIdentity:ServiceIdentity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceIdentity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceIdentity Get(string name, Input<string> id, ServiceIdentityState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceIdentity(name, id, state, options);
        }
    }

    public sealed class ServiceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service to generate identity for.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ServiceIdentityArgs()
        {
        }
    }

    public sealed class ServiceIdentityState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email address of the Google managed service account.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service to generate identity for.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public ServiceIdentityState()
        {
        }
    }
}
