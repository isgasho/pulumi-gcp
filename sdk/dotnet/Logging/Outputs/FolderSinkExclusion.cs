// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Outputs
{

    [OutputType]
    public sealed class FolderSinkExclusion
    {
        public readonly string? Description;
        public readonly bool? Disabled;
        /// <summary>
        /// The filter to apply when exporting logs. Only log entries that match the filter are exported.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// The name of the logging sink.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private FolderSinkExclusion(
            string? description,

            bool? disabled,

            string filter,

            string name)
        {
            Description = description;
            Disabled = disabled;
            Filter = filter;
            Name = name;
        }
    }
}
