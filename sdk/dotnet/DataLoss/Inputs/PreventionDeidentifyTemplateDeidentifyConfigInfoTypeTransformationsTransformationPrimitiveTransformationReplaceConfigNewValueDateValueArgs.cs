// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueDateValueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of month. Must be from 1 to 31 and valid for the year and month, or 0 if specifying a
        /// year by itself or a year and month where the day is not significant.
        /// </summary>
        [Input("day")]
        public Input<int>? Day { get; set; }

        /// <summary>
        /// Month of year. Must be from 1 to 12, or 0 if specifying a year without a month and day.
        /// </summary>
        [Input("month")]
        public Input<int>? Month { get; set; }

        /// <summary>
        /// Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year.
        /// </summary>
        [Input("year")]
        public Input<int>? Year { get; set; }

        public PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceConfigNewValueDateValueArgs()
        {
        }
    }
}
