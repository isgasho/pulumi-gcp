// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class PatchDeploymentRecurringScheduleMonthlyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.
        /// Months without the target day will be skipped. For example, a schedule to run "every month on the 31st"
        /// will not run in February, April, June, etc.
        /// </summary>
        [Input("monthDay")]
        public Input<int>? MonthDay { get; set; }

        /// <summary>
        /// Week day in a month.  Structure is documented below.
        /// </summary>
        [Input("weekDayOfMonth")]
        public Input<Inputs.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthGetArgs>? WeekDayOfMonth { get; set; }

        public PatchDeploymentRecurringScheduleMonthlyGetArgs()
        {
        }
    }
}
