// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Outputs
{

    [OutputType]
    public sealed class PatchDeploymentRecurringScheduleTimeZone
    {
        /// <summary>
        /// IANA Time Zone Database time zone, e.g. "America/New_York".
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IANA Time Zone Database version number, e.g. "2019a".
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private PatchDeploymentRecurringScheduleTimeZone(
            string id,

            string? version)
        {
            Id = id;
            Version = version;
        }
    }
}
