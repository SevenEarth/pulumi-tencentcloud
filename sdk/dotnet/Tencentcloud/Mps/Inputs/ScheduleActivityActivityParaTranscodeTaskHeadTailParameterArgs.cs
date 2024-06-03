// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleActivityActivityParaTranscodeTaskHeadTailParameterArgs : global::Pulumi.ResourceArgs
    {
        [Input("headSets")]
        private InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetArgs>? _headSets;

        /// <summary>
        /// Opening credits list.
        /// </summary>
        public InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetArgs> HeadSets
        {
            get => _headSets ?? (_headSets = new InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetArgs>());
            set => _headSets = value;
        }

        [Input("tailSets")]
        private InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterTailSetArgs>? _tailSets;

        /// <summary>
        /// Closing credits list.
        /// </summary>
        public InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterTailSetArgs> TailSets
        {
            get => _tailSets ?? (_tailSets = new InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterTailSetArgs>());
            set => _tailSets = value;
        }

        public ScheduleActivityActivityParaTranscodeTaskHeadTailParameterArgs()
        {
        }
        public static new ScheduleActivityActivityParaTranscodeTaskHeadTailParameterArgs Empty => new ScheduleActivityActivityParaTranscodeTaskHeadTailParameterArgs();
    }
}
