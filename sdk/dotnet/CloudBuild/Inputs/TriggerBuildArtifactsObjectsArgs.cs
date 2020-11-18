// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerBuildArtifactsObjectsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".
        /// Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
        /// this location as a prefix.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("paths")]
        private InputList<string>? _paths;

        /// <summary>
        /// Path globs used to match files in the build's workspace.
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        [Input("timings")]
        private InputList<Inputs.TriggerBuildArtifactsObjectsTimingArgs>? _timings;

        /// <summary>
        /// -
        /// Output only. Stores timing information for pushing all artifact objects.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.TriggerBuildArtifactsObjectsTimingArgs> Timings
        {
            get => _timings ?? (_timings = new InputList<Inputs.TriggerBuildArtifactsObjectsTimingArgs>());
            set => _timings = value;
        }

        public TriggerBuildArtifactsObjectsArgs()
        {
        }
    }
}