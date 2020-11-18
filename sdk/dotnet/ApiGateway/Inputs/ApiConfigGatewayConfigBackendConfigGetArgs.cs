// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ApiGateway.Inputs
{

    public sealed class ApiConfigGatewayConfigBackendConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google Cloud IAM service account used to sign OIDC tokens for backends that have authentication configured
        /// (https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backend).
        /// </summary>
        [Input("googleServiceAccount", required: true)]
        public Input<string> GoogleServiceAccount { get; set; } = null!;

        public ApiConfigGatewayConfigBackendConfigGetArgs()
        {
        }
    }
}