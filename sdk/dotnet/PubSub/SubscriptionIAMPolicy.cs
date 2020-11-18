// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
    /// 
    /// * `gcp.pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
    /// * `gcp.pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
    /// * `gcp.pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
    /// 
    /// &gt; **Note:** `gcp.pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `gcp.pubsub.SubscriptionIAMBinding` and `gcp.pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `gcp.pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_pubsub\_subscription\_iam\_policy
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
    ///                     Role = "roles/editor",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var editor = new Gcp.PubSub.SubscriptionIAMPolicy("editor", new Gcp.PubSub.SubscriptionIAMPolicyArgs
    ///         {
    ///             Subscription = "your-subscription-name",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_pubsub\_subscription\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var editor = new Gcp.PubSub.SubscriptionIAMBinding("editor", new Gcp.PubSub.SubscriptionIAMBindingArgs
    ///         {
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Role = "roles/editor",
    ///             Subscription = "your-subscription-name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_pubsub\_subscription\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var editor = new Gcp.PubSub.SubscriptionIAMMember("editor", new Gcp.PubSub.SubscriptionIAMMemberArgs
    ///         {
    ///             Member = "user:jane@example.com",
    ///             Role = "roles/editor",
    ///             Subscription = "your-subscription-name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Pubsub subscription IAM resources can be imported using the project, subscription name, role and member.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy editor projects/{your-project-id}/subscriptions/{your-subscription-name}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    public partial class SubscriptionIAMPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the subscription's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The subscription name or id to bind to attach IAM policy to.
        /// </summary>
        [Output("subscription")]
        public Output<string> Subscription { get; private set; } = null!;


        /// <summary>
        /// Create a SubscriptionIAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubscriptionIAMPolicy(string name, SubscriptionIAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy", name, args ?? new SubscriptionIAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubscriptionIAMPolicy(string name, Input<string> id, SubscriptionIAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubscriptionIAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubscriptionIAMPolicy Get(string name, Input<string> id, SubscriptionIAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SubscriptionIAMPolicy(name, id, state, options);
        }
    }

    public sealed class SubscriptionIAMPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The subscription name or id to bind to attach IAM policy to.
        /// </summary>
        [Input("subscription", required: true)]
        public Input<string> Subscription { get; set; } = null!;

        public SubscriptionIAMPolicyArgs()
        {
        }
    }

    public sealed class SubscriptionIAMPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the subscription's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The subscription name or id to bind to attach IAM policy to.
        /// </summary>
        [Input("subscription")]
        public Input<string>? Subscription { get; set; }

        public SubscriptionIAMPolicyState()
        {
        }
    }
}
