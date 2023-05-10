// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetRawParameterAudioTemplate
    {
        /// <summary>
        /// Audio channel mode, optional values:`1: single channel.2: Dual channel.6: Stereo.When the package format of the media is an audio format (flac, ogg, mp3, m4a), the number of channels is not allowed to be set to stereo.Default: 2.
        /// </summary>
        public readonly int? AudioChannel;
        /// <summary>
        /// Bit rate of the audio stream, value range: 0 and [26, 256], unit: kbps.When the value is 0, it means that the audio bit rate is consistent with the original audio.
        /// </summary>
        public readonly int Bitrate;
        /// <summary>
        /// Encoding format of frequency stream.When the outer parameter Container is mp3, the optional value is:libmp3lame.When the outer parameter Container is ogg or flac, the optional value is:flac.When the outer parameter Container is m4a, the optional value is:libfdk_aac.libmp3lame.ac3.When the outer parameter Container is mp4 or flv, the optional value is:libfdk_aac: more suitable for mp4.libmp3lame: more suitable for flv.When the outer parameter Container is hls, the optional value is:libfdk_aac.libmp3lame.
        /// </summary>
        public readonly string Codec;
        /// <summary>
        /// Sampling rate of audio stream, optional value.32000.44100.48000.Unit: Hz.
        /// </summary>
        public readonly int SampleRate;

        [OutputConstructor]
        private WorkflowMediaProcessTaskTranscodeTaskSetRawParameterAudioTemplate(
            int? audioChannel,

            int bitrate,

            string codec,

            int sampleRate)
        {
            AudioChannel = audioChannel;
            Bitrate = bitrate;
            Codec = codec;
            SampleRate = sampleRate;
        }
    }
}
