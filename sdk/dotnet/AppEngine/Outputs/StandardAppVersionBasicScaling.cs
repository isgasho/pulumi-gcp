// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Outputs
{

    [OutputType]
    public sealed class StandardAppVersionBasicScaling
    {
        /// <summary>
        /// Duration of time after the last request that an instance must wait before the instance is shut down.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
        /// </summary>
        public readonly string? IdleTimeout;
        /// <summary>
        /// Maximum number of instances to create for this version. Must be in the range [1.0, 200.0].
        /// </summary>
        public readonly int MaxInstances;

        [OutputConstructor]
        private StandardAppVersionBasicScaling(
            string? idleTimeout,

            int maxInstances)
        {
            IdleTimeout = idleTimeout;
            MaxInstances = maxInstances;
        }
    }
}
