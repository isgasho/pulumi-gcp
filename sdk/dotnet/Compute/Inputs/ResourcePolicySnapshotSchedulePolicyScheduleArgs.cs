// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class ResourcePolicySnapshotSchedulePolicyScheduleArgs : Pulumi.ResourceArgs
    {
        [Input("dailySchedule")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs>? DailySchedule { get; set; }

        [Input("hourlySchedule")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleArgs>? HourlySchedule { get; set; }

        [Input("weeklySchedule")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleArgs>? WeeklySchedule { get; set; }

        public ResourcePolicySnapshotSchedulePolicyScheduleArgs()
        {
        }
    }
}