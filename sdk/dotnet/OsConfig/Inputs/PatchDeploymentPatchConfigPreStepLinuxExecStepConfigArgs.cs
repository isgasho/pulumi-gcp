// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class PatchDeploymentPatchConfigPreStepLinuxExecStepConfigArgs : Pulumi.ResourceArgs
    {
        [Input("allowedSuccessCodes")]
        private InputList<int>? _allowedSuccessCodes;

        /// <summary>
        /// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
        /// </summary>
        public InputList<int> AllowedSuccessCodes
        {
            get => _allowedSuccessCodes ?? (_allowedSuccessCodes = new InputList<int>());
            set => _allowedSuccessCodes = value;
        }

        /// <summary>
        /// A Cloud Storage object containing the executable.  Structure is documented below.
        /// </summary>
        [Input("gcsObject")]
        public Input<Inputs.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObjectArgs>? GcsObject { get; set; }

        /// <summary>
        /// The script interpreter to use to run the script. If no interpreter is specified the script will
        /// be executed directly, which will likely only succeed for scripts with shebang lines.
        /// </summary>
        [Input("interpreter")]
        public Input<string>? Interpreter { get; set; }

        /// <summary>
        /// An absolute path to the executable on the VM.
        /// </summary>
        [Input("localPath")]
        public Input<string>? LocalPath { get; set; }

        public PatchDeploymentPatchConfigPreStepLinuxExecStepConfigArgs()
        {
        }
    }
}
