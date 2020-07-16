// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Outputs
{

    [OutputType]
    public sealed class GetGameServerDeploymentRolloutGameServerConfigOverrideResult
    {
        public readonly string ConfigVersion;
        public readonly ImmutableArray<Outputs.GetGameServerDeploymentRolloutGameServerConfigOverrideRealmsSelectorResult> RealmsSelectors;

        [OutputConstructor]
        private GetGameServerDeploymentRolloutGameServerConfigOverrideResult(
            string configVersion,

            ImmutableArray<Outputs.GetGameServerDeploymentRolloutGameServerConfigOverrideRealmsSelectorResult> realmsSelectors)
        {
            ConfigVersion = configVersion;
            RealmsSelectors = realmsSelectors;
        }
    }
}