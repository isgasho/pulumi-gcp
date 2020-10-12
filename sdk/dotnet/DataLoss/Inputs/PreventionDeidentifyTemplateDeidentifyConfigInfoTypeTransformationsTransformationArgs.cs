// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Inputs
{

    public sealed class PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs : Pulumi.ResourceArgs
    {
        [Input("infoTypes")]
        private InputList<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs>? _infoTypes;

        /// <summary>
        /// InfoTypes to apply the transformation to. Leaving this empty will apply the transformation to apply to
        /// all findings that correspond to infoTypes that were requested in InspectConfig.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs> InfoTypes
        {
            get => _infoTypes ?? (_infoTypes = new InputList<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoTypeArgs>());
            set => _infoTypes = value;
        }

        /// <summary>
        /// Primitive transformation to apply to the infoType.
        /// Structure is documented below.
        /// </summary>
        [Input("primitiveTransformation", required: true)]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationArgs> PrimitiveTransformation { get; set; } = null!;

        public PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationArgs()
        {
        }
    }
}
