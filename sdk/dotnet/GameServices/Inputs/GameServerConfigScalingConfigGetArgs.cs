// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Inputs
{

    public sealed class GameServerConfigScalingConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("fleetAutoscalerSpec", required: true)]
        public Input<string> FleetAutoscalerSpec { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("schedules")]
        private InputList<Inputs.GameServerConfigScalingConfigScheduleGetArgs>? _schedules;
        public InputList<Inputs.GameServerConfigScalingConfigScheduleGetArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.GameServerConfigScalingConfigScheduleGetArgs>());
            set => _schedules = value;
        }

        [Input("selectors")]
        private InputList<Inputs.GameServerConfigScalingConfigSelectorGetArgs>? _selectors;
        public InputList<Inputs.GameServerConfigScalingConfigSelectorGetArgs> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Inputs.GameServerConfigScalingConfigSelectorGetArgs>());
            set => _selectors = value;
        }

        public GameServerConfigScalingConfigGetArgs()
        {
        }
    }
}