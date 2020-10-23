// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class ClusterClusterConfigGceClusterConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, clusters are not restricted to internal IP addresses,
        /// and will have ephemeral external IP addresses assigned to each instance. If set to true, all
        /// instances in the cluster will only have internal IP addresses. Note: Private Google Access
        /// (also known as `privateIpGoogleAccess`) must be enabled on the subnetwork that the cluster
        /// will be launched in.
        /// </summary>
        [Input("internalIpOnly")]
        public Input<bool>? InternalIpOnly { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A map of the Compute Engine metadata entries to add to all instances
        /// (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name or self_link of the Google Compute Engine
        /// network to the cluster will be part of. Conflicts with `subnetwork`.
        /// If neither is specified, this defaults to the "default" network.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The service account to be used by the Node VMs.
        /// If not specified, the "default" service account is used.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("serviceAccountScopes")]
        private InputList<string>? _serviceAccountScopes;

        /// <summary>
        /// The set of Google API scopes
        /// to be made available on all of the node VMs under the `service_account`
        /// specified. These can be	either FQDNs, or scope aliases. The following scopes
        /// must be set if any other scopes are set. They're necessary to ensure the
        /// correct functioning ofthe cluster, and are set automatically by the API:
        /// </summary>
        public InputList<string> ServiceAccountScopes
        {
            get => _serviceAccountScopes ?? (_serviceAccountScopes = new InputList<string>());
            set => _serviceAccountScopes = value;
        }

        /// <summary>
        /// The name or self_link of the Google Compute Engine
        /// subnetwork the cluster will be part of. Conflicts with `network`.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of instance tags applied to instances in the cluster.
        /// Tags are used to identify valid sources or targets for network firewalls.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The GCP zone where your data is stored and used (i.e. where
        /// the master and the worker nodes will be created in). If `region` is set to 'global' (default)
        /// then `zone` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone)
        /// to determine this automatically for you.
        /// Note: This setting additionally determines and restricts
        /// which computing resources are available for use with other configs such as
        /// `cluster_config.master_config.machine_type` and `cluster_config.worker_config.machine_type`.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ClusterClusterConfigGceClusterConfigArgs()
        {
        }
    }
}
