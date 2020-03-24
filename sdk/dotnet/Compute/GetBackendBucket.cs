// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        public static Task<GetBackendBucketResult> GetBackendBucket(GetBackendBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackendBucketResult>("gcp:compute/getBackendBucket:getBackendBucket", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetBackendBucketArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetBackendBucketArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetBackendBucketResult
    {
        public readonly string BucketName;
        public readonly ImmutableArray<Outputs.GetBackendBucketCdnPoliciesResult> CdnPolicies;
        public readonly string CreationTimestamp;
        public readonly string Description;
        public readonly bool EnableCdn;
        public readonly string Name;
        public readonly string? Project;
        public readonly string SelfLink;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetBackendBucketResult(
            string bucketName,
            ImmutableArray<Outputs.GetBackendBucketCdnPoliciesResult> cdnPolicies,
            string creationTimestamp,
            string description,
            bool enableCdn,
            string name,
            string? project,
            string selfLink,
            string id)
        {
            BucketName = bucketName;
            CdnPolicies = cdnPolicies;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EnableCdn = enableCdn;
            Name = name;
            Project = project;
            SelfLink = selfLink;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetBackendBucketCdnPoliciesResult
    {
        public readonly int SignedUrlCacheMaxAgeSec;

        [OutputConstructor]
        private GetBackendBucketCdnPoliciesResult(int signedUrlCacheMaxAgeSec)
        {
            SignedUrlCacheMaxAgeSec = signedUrlCacheMaxAgeSec;
        }
    }
    }
}