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

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterAudioTemplateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audio channel mode, optional values:`1: single channel.2: Dual channel.6: Stereo.When the package format of the media is an audio format (flac, ogg, mp3, m4a), the number of channels is not allowed to be set to stereo.Default: 2.
        /// </summary>
        [Input("audioChannel")]
        public Input<int>? AudioChannel { get; set; }

        /// <summary>
        /// Bit rate of the audio stream, value range: 0 and [26, 256], unit: kbps.When the value is 0, it means that the audio bit rate is consistent with the original audio.
        /// </summary>
        [Input("bitrate")]
        public Input<int>? Bitrate { get; set; }

        /// <summary>
        /// Encoding format of frequency stream.When the outer parameter Container is mp3, the optional value is:libmp3lame.When the outer parameter Container is ogg or flac, the optional value is:flac.When the outer parameter Container is m4a, the optional value is:libfdk_aac.libmp3lame.ac3.When the outer parameter Container is mp4 or flv, the optional value is:libfdk_aac: more suitable for mp4.libmp3lame: more suitable for flv.When the outer parameter Container is hls, the optional value is:libfdk_aac.libmp3lame.
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        /// <summary>
        /// Sampling rate of audio stream, optional value.32000.44100.48000.Unit: Hz.
        /// </summary>
        [Input("sampleRate")]
        public Input<int>? SampleRate { get; set; }

        [Input("streamSelects")]
        private InputList<int>? _streamSelects;

        /// <summary>
        /// Specifies the audio track to preserve for the output. The default is to keep all sources.
        /// </summary>
        public InputList<int> StreamSelects
        {
            get => _streamSelects ?? (_streamSelects = new InputList<int>());
            set => _streamSelects = value;
        }

        public WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterAudioTemplateGetArgs()
        {
        }
    }
}
