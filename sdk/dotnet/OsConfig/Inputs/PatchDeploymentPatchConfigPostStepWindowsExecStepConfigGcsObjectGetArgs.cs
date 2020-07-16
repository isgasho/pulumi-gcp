// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket of the Cloud Storage object.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
        /// </summary>
        [Input("generationNumber", required: true)]
        public Input<string> GenerationNumber { get; set; } = null!;

        /// <summary>
        /// Name of the Cloud Storage object.
        /// </summary>
        [Input("object", required: true)]
        public Input<string> Object { get; set; } = null!;

        public PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectGetArgs()
        {
        }
    }
}
