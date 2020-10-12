// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterNotificationConfigPubsubGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable the PodSecurityPolicy controller for this cluster.
        /// If enabled, pods must be valid under a PodSecurityPolicy to be created.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public ClusterNotificationConfigPubsubGetArgs()
        {
        }
    }
}
