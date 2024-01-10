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

    public sealed class ScheduleActivityActivityParaTranscodeTaskOverrideParameterGetArgs : Pulumi.ResourceArgs
    {
        [Input("addOnSubtitles")]
        private InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleGetArgs>? _addOnSubtitles;

        /// <summary>
        /// Subtitle files to insert.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleGetArgs> AddOnSubtitles
        {
            get => _addOnSubtitles ?? (_addOnSubtitles = new InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitleGetArgs>());
            set => _addOnSubtitles = value;
        }

        [Input("addonAudioStreams")]
        private InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStreamGetArgs>? _addonAudioStreams;

        /// <summary>
        /// The information of the external audio track to add.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStreamGetArgs> AddonAudioStreams
        {
            get => _addonAudioStreams ?? (_addonAudioStreams = new InputList<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStreamGetArgs>());
            set => _addonAudioStreams = value;
        }

        /// <summary>
        /// Audio stream configuration parameter.
        /// </summary>
        [Input("audioTemplate")]
        public Input<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAudioTemplateGetArgs>? AudioTemplate { get; set; }

        /// <summary>
        /// Container format. Valid values: mp4, flv, hls, mp3, flac, ogg, and m4a; mp3, flac, ogg, and m4a are formats of audio files.
        /// </summary>
        [Input("container")]
        public Input<string>? Container { get; set; }

        /// <summary>
        /// Whether to remove audio data. Valid values: 0: retain 1: remove.
        /// </summary>
        [Input("removeAudio")]
        public Input<int>? RemoveAudio { get; set; }

        /// <summary>
        /// Whether to remove video data. Valid values: 0: retain 1: remove.
        /// </summary>
        [Input("removeVideo")]
        public Input<int>? RemoveVideo { get; set; }

        /// <summary>
        /// Transcoding extension field.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("stdExtInfo")]
        public Input<string>? StdExtInfo { get; set; }

        /// <summary>
        /// The subtitle settings.
        /// </summary>
        [Input("subtitleTemplate")]
        public Input<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplateGetArgs>? SubtitleTemplate { get; set; }

        /// <summary>
        /// TESHD transcoding parameter.
        /// </summary>
        [Input("tehdConfig")]
        public Input<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterTehdConfigGetArgs>? TehdConfig { get; set; }

        /// <summary>
        /// Video stream configuration parameter.
        /// </summary>
        [Input("videoTemplate")]
        public Input<Inputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterVideoTemplateGetArgs>? VideoTemplate { get; set; }

        public ScheduleActivityActivityParaTranscodeTaskOverrideParameterGetArgs()
        {
        }
    }
}
