// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
    /// 
    /// * `gcp.healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
    /// * `gcp.healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
    /// * `gcp.healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
    /// 
    /// &gt; **Note:** `gcp.healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.Hl7StoreIamBinding` and `gcp.healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_healthcare\_hl7\_v2\_store\_iam\_policy
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
    ///             Binding = 
    ///             {
    ///                 
    ///                 {
    ///                     { "role", "roles/editor" },
    ///                     { "members", 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     } },
    ///                 },
    ///             },
    ///         }));
    ///         var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamPolicy("hl7V2Store", new Gcp.Healthcare.Hl7StoreIamPolicyArgs
    ///         {
    ///             Hl7V2StoreId = "your-hl7-v2-store-id",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_healthcare\_hl7\_v2\_store\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamBinding("hl7V2Store", new Gcp.Healthcare.Hl7StoreIamBindingArgs
    ///         {
    ///             Hl7V2StoreId = "your-hl7-v2-store-id",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_healthcare\_hl7\_v2\_store\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamMember("hl7V2Store", new Gcp.Healthcare.Hl7StoreIamMemberArgs
    ///         {
    ///             Hl7V2StoreId = "your-hl7-v2-store-id",
    ///             Member = "user:jane@example.com",
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Hl7StoreIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("hl7V2StoreId")]
        public Output<string> Hl7V2StoreId { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a Hl7StoreIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hl7StoreIamPolicy(string name, Hl7StoreIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, args ?? new Hl7StoreIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hl7StoreIamPolicy(string name, Input<string> id, Hl7StoreIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Hl7StoreIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hl7StoreIamPolicy Get(string name, Input<string> id, Hl7StoreIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Hl7StoreIamPolicy(name, id, state, options);
        }
    }

    public sealed class Hl7StoreIamPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId", required: true)]
        public Input<string> Hl7V2StoreId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public Hl7StoreIamPolicyArgs()
        {
        }
    }

    public sealed class Hl7StoreIamPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId")]
        public Input<string>? Hl7V2StoreId { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public Hl7StoreIamPolicyState()
        {
        }
    }
}
