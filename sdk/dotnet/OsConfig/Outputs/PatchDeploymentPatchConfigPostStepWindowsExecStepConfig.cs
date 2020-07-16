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
    public sealed class PatchDeploymentPatchConfigPostStepWindowsExecStepConfig
    {
        /// <summary>
        /// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
        /// </summary>
        public readonly ImmutableArray<int> AllowedSuccessCodes;
        /// <summary>
        /// A Cloud Storage object containing the executable.  Structure is documented below.
        /// </summary>
        public readonly Outputs.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject? GcsObject;
        /// <summary>
        /// The script interpreter to use to run the script. If no interpreter is specified the script will
        /// be executed directly, which will likely only succeed for scripts with shebang lines.
        /// </summary>
        public readonly string? Interpreter;
        /// <summary>
        /// An absolute path to the executable on the VM.
        /// </summary>
        public readonly string? LocalPath;

        [OutputConstructor]
        private PatchDeploymentPatchConfigPostStepWindowsExecStepConfig(
            ImmutableArray<int> allowedSuccessCodes,

            Outputs.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject? gcsObject,

            string? interpreter,

            string? localPath)
        {
            AllowedSuccessCodes = allowedSuccessCodes;
            GcsObject = gcsObject;
            Interpreter = interpreter;
            LocalPath = localPath;
        }
    }
}
