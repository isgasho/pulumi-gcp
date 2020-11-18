// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class AutoscalarAutoscalingPolicyScaleInControl
    {
        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.AutoscalarAutoscalingPolicyScaleInControlMaxScaledInReplicas? MaxScaledInReplicas;
        /// <summary>
        /// How long back autoscaling should look when computing recommendations
        /// to include directives regarding slower scale down, as described above.
        /// </summary>
        public readonly int? TimeWindowSec;

        [OutputConstructor]
        private AutoscalarAutoscalingPolicyScaleInControl(
            Outputs.AutoscalarAutoscalingPolicyScaleInControlMaxScaledInReplicas? maxScaledInReplicas,

            int? timeWindowSec)
        {
            MaxScaledInReplicas = maxScaledInReplicas;
            TimeWindowSec = timeWindowSec;
        }
    }
}
