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

    public sealed class MediaTranscodeTemplateTransConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resolution adjustment method, value scale, crop, pad, none, When the aspect ratio of the output video is different from the original video, adjust the resolution accordingly according to this parameter.
        /// </summary>
        [Input("adjDarMethod")]
        public Input<string>? AdjDarMethod { get; set; }

        /// <summary>
        /// Audio bit rate adjustment mode, value 0, 1; when the output audio bit rate is greater than the original audio bit rate, 0 means use the original audio bit rate; 1 means return transcoding failed, Take effect when IsCheckAudioBitrate is true.
        /// </summary>
        [Input("audioBitrateAdjMethod")]
        public Input<string>? AudioBitrateAdjMethod { get; set; }

        /// <summary>
        /// Whether to delete the MetaData information in the file, true, false, When false, keep source file information.
        /// </summary>
        [Input("deleteMetadata")]
        public Input<string>? DeleteMetadata { get; set; }

        /// <summary>
        /// hls encryption configuration.
        /// </summary>
        [Input("hlsEncrypt")]
        public Input<Inputs.MediaTranscodeTemplateTransConfigHlsEncryptGetArgs>? HlsEncrypt { get; set; }

        /// <summary>
        /// Whether to check the audio code rate, true, false, When false, transcode according to configuration parameters.
        /// </summary>
        [Input("isCheckAudioBitrate")]
        public Input<string>? IsCheckAudioBitrate { get; set; }

        /// <summary>
        /// Whether to check the resolution, when it is false, transcode according to the configuration parameters.
        /// </summary>
        [Input("isCheckReso")]
        public Input<string>? IsCheckReso { get; set; }

        /// <summary>
        /// Whether to check the video code rate, when it is false, transcode according to the configuration parameters.
        /// </summary>
        [Input("isCheckVideoBitrate")]
        public Input<string>? IsCheckVideoBitrate { get; set; }

        /// <summary>
        /// Whether to enable HDR to SDR true, false.
        /// </summary>
        [Input("isHdr2Sdr")]
        public Input<string>? IsHdr2Sdr { get; set; }

        /// <summary>
        /// Resolution adjustment mode, value 0, 1; 0 means use the original video resolution; 1 means return transcoding failed, Take effect when IsCheckReso is true.
        /// </summary>
        [Input("resoAdjMethod")]
        public Input<string>? ResoAdjMethod { get; set; }

        /// <summary>
        /// Video bit rate adjustment method, value 0, 1; when the output video bit rate is greater than the original video bit rate, 0 means use the original video bit rate; 1 means return transcoding failed, Take effect when IsCheckVideoBitrate is true.
        /// </summary>
        [Input("videoBitrateAdjMethod")]
        public Input<string>? VideoBitrateAdjMethod { get; set; }

        public MediaTranscodeTemplateTransConfigGetArgs()
        {
        }
    }
}