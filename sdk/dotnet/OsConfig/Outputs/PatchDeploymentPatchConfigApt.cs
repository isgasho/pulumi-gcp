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
    public sealed class PatchDeploymentPatchConfigApt
    {
        /// <summary>
        /// List of KBs to exclude from update.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// An exclusive list of packages to be updated. These are the only packages that will be updated.
        /// If these packages are not installed, they will be ignored. This field cannot be specified with
        /// any other patch configuration fields.
        /// </summary>
        public readonly ImmutableArray<string> ExclusivePackages;
        /// <summary>
        /// By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private PatchDeploymentPatchConfigApt(
            ImmutableArray<string> excludes,

            ImmutableArray<string> exclusivePackages,

            string? type)
        {
            Excludes = excludes;
            ExclusivePackages = exclusivePackages;
            Type = type;
        }
    }
}
