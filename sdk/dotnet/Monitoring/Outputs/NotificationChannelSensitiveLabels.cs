// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Outputs
{

    [OutputType]
    public sealed class NotificationChannelSensitiveLabels
    {
        public readonly string? AuthToken;
        public readonly string? Password;
        public readonly string? ServiceKey;

        [OutputConstructor]
        private NotificationChannelSensitiveLabels(
            string? authToken,

            string? password,

            string? serviceKey)
        {
            AuthToken = authToken;
            Password = password;
            ServiceKey = serviceKey;
        }
    }
}