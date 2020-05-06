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
    public sealed class ReservationSpecificReservationInstancePropertiesLocalSsd
    {
        /// <summary>
        /// The size of the disk in base-2 GB.
        /// </summary>
        public readonly int DiskSizeGb;
        /// <summary>
        /// The disk interface to use for attaching this disk.
        /// </summary>
        public readonly string? Interface;

        [OutputConstructor]
        private ReservationSpecificReservationInstancePropertiesLocalSsd(
            int diskSizeGb,

            string? @interface)
        {
            DiskSizeGb = diskSizeGb;
            Interface = @interface;
        }
    }
}
