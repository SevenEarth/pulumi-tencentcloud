// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class AdaptiveDynamicStreamingTemplateStreamInfoAudioGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audio channel system. Valid values: mono, dual, stereo. Default value: dual.
        /// </summary>
        [Input("audioChannel")]
        public Input<string>? AudioChannel { get; set; }

        /// <summary>
        /// Audio stream bitrate in Kbps. Value range: `0` and `[26, 256]`. If the value is `0`, the bitrate of the audio stream will be the same as that of the original audio.
        /// </summary>
        [Input("bitrate", required: true)]
        public Input<int> Bitrate { get; set; } = null!;

        /// <summary>
        /// Audio stream encoder. Valid value are: `libfdk_aac` and `libmp3lame`. while `libfdk_aac` is recommended.
        /// </summary>
        [Input("codec", required: true)]
        public Input<string> Codec { get; set; } = null!;

        /// <summary>
        /// Audio stream sample rate. Valid values: `32000`, `44100`, `48000`Hz.
        /// </summary>
        [Input("sampleRate", required: true)]
        public Input<int> SampleRate { get; set; } = null!;

        public AdaptiveDynamicStreamingTemplateStreamInfoAudioGetArgs()
        {
        }
    }
}
