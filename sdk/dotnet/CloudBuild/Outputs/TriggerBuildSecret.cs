// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Outputs
{

    [OutputType]
    public sealed class TriggerBuildSecret
    {
        /// <summary>
        /// Cloud KMS key name to use to decrypt these envs.
        /// </summary>
        public readonly string KmsKeyName;
        /// <summary>
        /// A list of global environment variables, which are encrypted using a Cloud Key Management
        /// Service crypto key. These values must be specified in the build's Secret. These variables
        /// will be available to all build steps in this build.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? SecretEnv;

        [OutputConstructor]
        private TriggerBuildSecret(
            string kmsKeyName,

            ImmutableDictionary<string, string>? secretEnv)
        {
            KmsKeyName = kmsKeyName;
            SecretEnv = secretEnv;
        }
    }
}
