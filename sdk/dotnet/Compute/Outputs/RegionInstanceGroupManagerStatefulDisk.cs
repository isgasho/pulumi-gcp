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
    public sealed class RegionInstanceGroupManagerStatefulDisk
    {
        /// <summary>
        /// , A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are `NEVER` and `ON_PERMANENT_INSTANCE_DELETION`. `NEVER` detatch the disk when the VM is deleted, but not delete the disk. `ON_PERMANENT_INSTANCE_DELETION` will delete the stateful disk when the VM is permanently deleted from the instance group. The default is `NEVER`.
        /// </summary>
        public readonly string? DeleteRule;
        /// <summary>
        /// , The device name of the disk to be attached.
        /// </summary>
        public readonly string DeviceName;

        [OutputConstructor]
        private RegionInstanceGroupManagerStatefulDisk(
            string? deleteRule,

            string deviceName)
        {
            DeleteRule = deleteRule;
            DeviceName = deviceName;
        }
    }
}
