// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class PatchDeploymentRecurringScheduleWeeklyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A day of the week.
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<string> DayOfWeek { get; set; } = null!;

        public PatchDeploymentRecurringScheduleWeeklyArgs()
        {
        }
    }
}
