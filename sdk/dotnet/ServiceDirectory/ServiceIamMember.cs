// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceDirectory
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Service Directory Service. Each of these resources serves a different use case:
    /// 
    /// * `gcp.servicedirectory.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
    /// * `gcp.servicedirectory.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
    /// * `gcp.servicedirectory.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
    /// 
    /// &gt; **Note:** `gcp.servicedirectory.ServiceIamPolicy` **cannot** be used in conjunction with `gcp.servicedirectory.ServiceIamBinding` and `gcp.servicedirectory.ServiceIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.servicedirectory.ServiceIamBinding` resources **can be** used in conjunction with `gcp.servicedirectory.ServiceIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_service\_directory\_service\_iam\_policy
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
    ///                     Role = "roles/viewer",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.ServiceDirectory.ServiceIamPolicy("policy", new Gcp.ServiceDirectory.ServiceIamPolicyArgs
    ///         {
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_service\_directory\_service\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.ServiceDirectory.ServiceIamBinding("binding", new Gcp.ServiceDirectory.ServiceIamBindingArgs
    ///         {
    ///             Role = "roles/viewer",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_service\_directory\_service\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var member = new Gcp.ServiceDirectory.ServiceIamMember("member", new Gcp.ServiceDirectory.ServiceIamMemberArgs
    ///         {
    ///             Role = "roles/viewer",
    ///             Member = "user:jane@example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} * {{project}}/{{location}}/{{namespace_id}}/{{service_id}} * {{location}}/{{namespace_id}}/{{service_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:servicedirectory/serviceIamMember:ServiceIamMember editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    public partial class ServiceIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.ServiceIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceIamMember(string name, ServiceIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/serviceIamMember:ServiceIamMember", name, args ?? new ServiceIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceIamMember(string name, Input<string> id, ServiceIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/serviceIamMember:ServiceIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceIamMember Get(string name, Input<string> id, ServiceIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceIamMember(name, id, state, options);
        }
    }

    public sealed class ServiceIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.ServiceIamMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public ServiceIamMemberArgs()
        {
        }
    }

    public sealed class ServiceIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.ServiceIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.servicedirectory.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public ServiceIamMemberState()
        {
        }
    }
}
