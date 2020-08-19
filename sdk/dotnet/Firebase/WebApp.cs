// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firebase
{
    /// <summary>
    /// A Google Cloud Firebase web application instance
    /// 
    /// To get more information about WebApp, see:
    /// 
    /// * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
    /// * How-to Guides
    ///     * [Official Documentation](https://firebase.google.com/)
    /// </summary>
    public partial class WebApp : Pulumi.CustomResource
    {
        /// <summary>
        /// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
        /// token, as the data format is not specified.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// The user-assigned display name of the App.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
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
        /// Create a WebApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebApp(string name, WebAppArgs args, CustomResourceOptions? options = null)
            : base("gcp:firebase/webApp:WebApp", name, args ?? new WebAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebApp(string name, Input<string> id, WebAppState? state = null, CustomResourceOptions? options = null)
            : base("gcp:firebase/webApp:WebApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebApp Get(string name, Input<string> id, WebAppState? state = null, CustomResourceOptions? options = null)
        {
            return new WebApp(name, id, state, options);
        }
    }

    public sealed class WebAppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user-assigned display name of the App.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public WebAppArgs()
        {
        }
    }

    public sealed class WebAppState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. The globally unique, Firebase-assigned identifier of the App. This identifier should be treated as an opaque
        /// token, as the data format is not specified.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// The user-assigned display name of the App.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The fully qualified resource name of the App, for example: projects/projectId/webApps/appId
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public WebAppState()
        {
        }
    }
}
