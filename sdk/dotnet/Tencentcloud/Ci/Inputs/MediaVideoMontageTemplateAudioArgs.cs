// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Inputs
{

    public sealed class MediaVideoMontageTemplateAudioArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Original audio bit rate, unit: Kbps, Value range: [8, 1000].
        /// </summary>
        [Input("bitrate")]
        public Input<string>? Bitrate { get; set; }

        /// <summary>
        /// number of channels- When Codec is set to aac, support 1, 2, 4, 5, 6, 8- When Codec is set to mp3, support 1, 2.
        /// </summary>
        [Input("channels")]
        public Input<string>? Channels { get; set; }

        /// <summary>
        /// Codec format, value aac, mp3.
        /// </summary>
        [Input("codec", required: true)]
        public Input<string> Codec { get; set; } = null!;

        /// <summary>
        /// Whether to delete the source audio stream, the value is true, false.
        /// </summary>
        [Input("remove")]
        public Input<string>? Remove { get; set; }

        /// <summary>
        /// Sampling Rate- Unit: Hz- Optional 11025, 22050, 32000, 44100, 48000, 96000- Different packages, mp3 supports different sampling rates, as shown in the table below.
        /// </summary>
        [Input("samplerate")]
        public Input<string>? Samplerate { get; set; }

        public MediaVideoMontageTemplateAudioArgs()
        {
        }
        public static new MediaVideoMontageTemplateAudioArgs Empty => new MediaVideoMontageTemplateAudioArgs();
    }
}
