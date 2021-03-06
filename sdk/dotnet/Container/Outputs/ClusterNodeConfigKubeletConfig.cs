// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterNodeConfigKubeletConfig
    {
        /// <summary>
        /// If true, enables CPU CFS quota enforcement for
        /// containers that specify CPU limits.
        /// </summary>
        public readonly bool? CpuCfsQuota;
        /// <summary>
        /// The CPU CFS quota period value. Specified
        /// as a sequence of decimal numbers, each with optional fraction and a unit suffix,
        /// such as `"300ms"`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m",
        /// "h". The value must be a positive duration.
        /// </summary>
        public readonly string? CpuCfsQuotaPeriod;
        /// <summary>
        /// The CPU management policy on the node. See
        /// [K8S CPU Management Policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/).
        /// One of `"none"` or `"static"`. Defaults to `none` when `kubelet_config` is unset.
        /// </summary>
        public readonly string CpuManagerPolicy;

        [OutputConstructor]
        private ClusterNodeConfigKubeletConfig(
            bool? cpuCfsQuota,

            string? cpuCfsQuotaPeriod,

            string cpuManagerPolicy)
        {
            CpuCfsQuota = cpuCfsQuota;
            CpuCfsQuotaPeriod = cpuCfsQuotaPeriod;
            CpuManagerPolicy = cpuManagerPolicy;
        }
    }
}
