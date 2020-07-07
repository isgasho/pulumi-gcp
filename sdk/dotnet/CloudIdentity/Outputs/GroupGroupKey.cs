// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudIdentity.Outputs
{

    [OutputType]
    public sealed class GroupGroupKey
    {
        /// <summary>
        /// The ID of the entity.
        /// For Google-managed entities, the id must be the email address of an existing
        /// group or user.
        /// For external-identity-mapped entities, the id must be a string conforming
        /// to the Identity Source's requirements.
        /// Must be unique within a namespace.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The namespace in which the entity exists.
        /// If not specified, the EntityKey represents a Google-managed entity
        /// such as a Google user or a Google Group.
        /// If specified, the EntityKey represents an external-identity-mapped group.
        /// The namespace must correspond to an identity source created in Admin Console
        /// and must be in the form of `identitysources/{identity_source_id}`.
        /// </summary>
        public readonly string? Namespace;

        [OutputConstructor]
        private GroupGroupKey(
            string id,

            string? @namespace)
        {
            Id = id;
            Namespace = @namespace;
        }
    }
}
