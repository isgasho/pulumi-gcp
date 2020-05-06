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
    public sealed class SnapshotSourceDiskEncryptionKey
    {
        /// <summary>
        /// Specifies a 256-bit customer-supplied encryption key, encoded in
        /// RFC 4648 base64 to either encrypt or decrypt this resource.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public readonly string? RawKey;

        [OutputConstructor]
        private SnapshotSourceDiskEncryptionKey(string? rawKey)
        {
            RawKey = rawKey;
        }
    }
}
