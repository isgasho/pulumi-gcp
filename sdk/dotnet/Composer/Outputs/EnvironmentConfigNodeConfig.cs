// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class EnvironmentConfigNodeConfig
    {
        public readonly int? DiskSizeGb;
        public readonly Outputs.EnvironmentConfigNodeConfigIpAllocationPolicy? IpAllocationPolicy;
        public readonly string? MachineType;
        public readonly string? Network;
        public readonly ImmutableArray<string> OauthScopes;
        public readonly string? ServiceAccount;
        public readonly string? Subnetwork;
        public readonly ImmutableArray<string> Tags;
        public readonly string Zone;

        [OutputConstructor]
        private EnvironmentConfigNodeConfig(
            int? diskSizeGb,

            Outputs.EnvironmentConfigNodeConfigIpAllocationPolicy? ipAllocationPolicy,

            string? machineType,

            string? network,

            ImmutableArray<string> oauthScopes,

            string? serviceAccount,

            string? subnetwork,

            ImmutableArray<string> tags,

            string zone)
        {
            DiskSizeGb = diskSizeGb;
            IpAllocationPolicy = ipAllocationPolicy;
            MachineType = machineType;
            Network = network;
            OauthScopes = oauthScopes;
            ServiceAccount = serviceAccount;
            Subnetwork = subnetwork;
            Tags = tags;
            Zone = zone;
        }
    }
}