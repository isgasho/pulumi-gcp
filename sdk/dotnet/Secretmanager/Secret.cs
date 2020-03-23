// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecretManager
{
    /// <summary>
    /// A Secret is a logical secret whose value and versions can be accessed.
    /// 
    /// To get more information about Secret, see:
    /// 
    /// * [API documentation](https://cloud.google.com/secret-manager/docs/reference/rest/v1beta1/projects.secrets)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/secret_manager_secret.html.markdown.
    /// </summary>
    public partial class Secret : Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the Secret was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8
        /// encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
        /// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a
        /// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
        /// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be assigned to a given resource. An object containing
        /// a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Secret. Format: 'projects/{{project}}/secrets/{{secret_id}}'
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The replication policy of the secret data attached to the Secret. It cannot be changed after the Secret has
        /// been created.
        /// </summary>
        [Output("replication")]
        public Output<Outputs.SecretReplication> Replication { get; private set; } = null!;

        /// <summary>
        /// This must be unique within the project.
        /// </summary>
        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs args, CustomResourceOptions? options = null)
            : base("gcp:secretmanager/secret:Secret", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
            : base("gcp:secretmanager/secret:Secret", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, state, options);
        }
    }

    public sealed class SecretArgs : Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8
        /// encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
        /// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a
        /// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
        /// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be assigned to a given resource. An object containing
        /// a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The replication policy of the secret data attached to the Secret. It cannot be changed after the Secret has
        /// been created.
        /// </summary>
        [Input("replication", required: true)]
        public Input<Inputs.SecretReplicationArgs> Replication { get; set; } = null!;

        /// <summary>
        /// This must be unique within the project.
        /// </summary>
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        public SecretArgs()
        {
        }
    }

    public sealed class SecretState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time at which the Secret was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels assigned to this Secret. Label keys must be between 1 and 63 characters long, have a UTF-8
        /// encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
        /// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values must be between 0 and 63 characters long, have a
        /// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
        /// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be assigned to a given resource. An object containing
        /// a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name of the Secret. Format: 'projects/{{project}}/secrets/{{secret_id}}'
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The replication policy of the secret data attached to the Secret. It cannot be changed after the Secret has
        /// been created.
        /// </summary>
        [Input("replication")]
        public Input<Inputs.SecretReplicationGetArgs>? Replication { get; set; }

        /// <summary>
        /// This must be unique within the project.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        public SecretState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SecretReplicationArgs : Pulumi.ResourceArgs
    {
        [Input("automatic")]
        public Input<bool>? Automatic { get; set; }

        [Input("userManaged")]
        public Input<SecretReplicationUserManagedArgs>? UserManaged { get; set; }

        public SecretReplicationArgs()
        {
        }
    }

    public sealed class SecretReplicationGetArgs : Pulumi.ResourceArgs
    {
        [Input("automatic")]
        public Input<bool>? Automatic { get; set; }

        [Input("userManaged")]
        public Input<SecretReplicationUserManagedGetArgs>? UserManaged { get; set; }

        public SecretReplicationGetArgs()
        {
        }
    }

    public sealed class SecretReplicationUserManagedArgs : Pulumi.ResourceArgs
    {
        [Input("replicas", required: true)]
        private InputList<SecretReplicationUserManagedReplicasArgs>? _replicas;
        public InputList<SecretReplicationUserManagedReplicasArgs> Replicas
        {
            get => _replicas ?? (_replicas = new InputList<SecretReplicationUserManagedReplicasArgs>());
            set => _replicas = value;
        }

        public SecretReplicationUserManagedArgs()
        {
        }
    }

    public sealed class SecretReplicationUserManagedGetArgs : Pulumi.ResourceArgs
    {
        [Input("replicas", required: true)]
        private InputList<SecretReplicationUserManagedReplicasGetArgs>? _replicas;
        public InputList<SecretReplicationUserManagedReplicasGetArgs> Replicas
        {
            get => _replicas ?? (_replicas = new InputList<SecretReplicationUserManagedReplicasGetArgs>());
            set => _replicas = value;
        }

        public SecretReplicationUserManagedGetArgs()
        {
        }
    }

    public sealed class SecretReplicationUserManagedReplicasArgs : Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public SecretReplicationUserManagedReplicasArgs()
        {
        }
    }

    public sealed class SecretReplicationUserManagedReplicasGetArgs : Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public SecretReplicationUserManagedReplicasGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecretReplication
    {
        public readonly bool? Automatic;
        public readonly SecretReplicationUserManaged? UserManaged;

        [OutputConstructor]
        private SecretReplication(
            bool? automatic,
            SecretReplicationUserManaged? userManaged)
        {
            Automatic = automatic;
            UserManaged = userManaged;
        }
    }

    [OutputType]
    public sealed class SecretReplicationUserManaged
    {
        public readonly ImmutableArray<SecretReplicationUserManagedReplicas> Replicas;

        [OutputConstructor]
        private SecretReplicationUserManaged(ImmutableArray<SecretReplicationUserManagedReplicas> replicas)
        {
            Replicas = replicas;
        }
    }

    [OutputType]
    public sealed class SecretReplicationUserManagedReplicas
    {
        public readonly string Location;

        [OutputConstructor]
        private SecretReplicationUserManagedReplicas(string location)
        {
            Location = location;
        }
    }
    }
}
