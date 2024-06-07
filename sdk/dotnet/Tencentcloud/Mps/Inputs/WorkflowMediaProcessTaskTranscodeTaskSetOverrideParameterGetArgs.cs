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

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audio stream configuration parameters.
        /// </summary>
        [Input("audioTemplate")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterAudioTemplateGetArgs>? AudioTemplate { get; set; }

        /// <summary>
        /// Encapsulation format, optional values: mp4, flv, hls, mp3, flac, ogg, m4a. Among them, mp3, flac, ogg, m4a are pure audio files.
        /// </summary>
        [Input("container")]
        public Input<string>? Container { get; set; }

        /// <summary>
        /// Whether to remove audio data, value:0: reserved.1: remove.
        /// </summary>
        [Input("removeAudio")]
        public Input<int>? RemoveAudio { get; set; }

        /// <summary>
        /// Whether to remove video data, value:0: reserved.1: remove.
        /// </summary>
        [Input("removeVideo")]
        public Input<int>? RemoveVideo { get; set; }

        /// <summary>
        /// Subtitle Stream Configuration Parameters.
        /// </summary>
        [Input("subtitleTemplate")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterSubtitleTemplateGetArgs>? SubtitleTemplate { get; set; }

        /// <summary>
        /// Ultra-fast HD transcoding parameters.
        /// </summary>
        [Input("tehdConfig")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterTehdConfigGetArgs>? TehdConfig { get; set; }

        /// <summary>
        /// Video streaming configuration parameters.
        /// </summary>
        [Input("videoTemplate")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterVideoTemplateGetArgs>? VideoTemplate { get; set; }

        public WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterGetArgs()
        {
        }
        public static new WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterGetArgs Empty => new WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterGetArgs();
    }
}
