// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Inputs
{

    public sealed class MetricMetricDescriptorGetArgs : Pulumi.ResourceArgs
    {
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputList<Inputs.MetricMetricDescriptorLabelGetArgs>? _labels;
        public InputList<Inputs.MetricMetricDescriptorLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.MetricMetricDescriptorLabelGetArgs>());
            set => _labels = value;
        }

        [Input("metricKind", required: true)]
        public Input<string> MetricKind { get; set; } = null!;

        [Input("unit")]
        public Input<string>? Unit { get; set; }

        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public MetricMetricDescriptorGetArgs()
        {
        }
    }
}