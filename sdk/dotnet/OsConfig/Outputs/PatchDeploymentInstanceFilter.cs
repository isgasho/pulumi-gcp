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
    public sealed class PatchDeploymentInstanceFilter
    {
        /// <summary>
        /// Target all VM instances in the project. If true, no other criteria is permitted.
        /// </summary>
        public readonly bool? All;
        /// <summary>
        /// Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.PatchDeploymentInstanceFilterGroupLabel> GroupLabels;
        /// <summary>
        /// Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group
        /// VMs when targeting configs, for example prefix="prod-".
        /// </summary>
        public readonly ImmutableArray<string> InstanceNamePrefixes;
        /// <summary>
        /// Targets any of the VM instances specified. Instances are specified by their URI in the `form zones/{{zone}}/instances/{{instance_name}}`,
        /// `projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}`, or
        /// `https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}`
        /// </summary>
        public readonly ImmutableArray<string> Instances;
        /// <summary>
        /// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private PatchDeploymentInstanceFilter(
            bool? all,

            ImmutableArray<Outputs.PatchDeploymentInstanceFilterGroupLabel> groupLabels,

            ImmutableArray<string> instanceNamePrefixes,

            ImmutableArray<string> instances,

            ImmutableArray<string> zones)
        {
            All = all;
            GroupLabels = groupLabels;
            InstanceNamePrefixes = instanceNamePrefixes;
            Instances = instances;
            Zones = zones;
        }
    }
}
