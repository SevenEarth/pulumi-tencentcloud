// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Outputs
{

    [OutputType]
    public sealed class MediaTranscodeTemplateTransConfig
    {
        /// <summary>
        /// Resolution adjustment method, value scale, crop, pad, none, When the aspect ratio of the output video is different from the original video, adjust the resolution accordingly according to this parameter.
        /// </summary>
        public readonly string? AdjDarMethod;
        /// <summary>
        /// Audio bit rate adjustment mode, value 0, 1; when the output audio bit rate is greater than the original audio bit rate, 0 means use the original audio bit rate; 1 means return transcoding failed, Take effect when IsCheckAudioBitrate is true.
        /// </summary>
        public readonly string? AudioBitrateAdjMethod;
        /// <summary>
        /// Whether to delete the MetaData information in the file, true, false, When false, keep source file information.
        /// </summary>
        public readonly string? DeleteMetadata;
        /// <summary>
        /// hls encryption configuration.
        /// </summary>
        public readonly Outputs.MediaTranscodeTemplateTransConfigHlsEncrypt? HlsEncrypt;
        /// <summary>
        /// Whether to check the audio code rate, true, false, When false, transcode according to configuration parameters.
        /// </summary>
        public readonly string? IsCheckAudioBitrate;
        /// <summary>
        /// Whether to check the resolution, when it is false, transcode according to the configuration parameters.
        /// </summary>
        public readonly string? IsCheckReso;
        /// <summary>
        /// Whether to check the video code rate, when it is false, transcode according to the configuration parameters.
        /// </summary>
        public readonly string? IsCheckVideoBitrate;
        /// <summary>
        /// Whether to enable HDR to SDR true, false.
        /// </summary>
        public readonly string? IsHdr2Sdr;
        /// <summary>
        /// Resolution adjustment mode, value 0, 1; 0 means use the original video resolution; 1 means return transcoding failed, Take effect when IsCheckReso is true.
        /// </summary>
        public readonly string? ResoAdjMethod;
        /// <summary>
        /// Video bit rate adjustment method, value 0, 1; when the output video bit rate is greater than the original video bit rate, 0 means use the original video bit rate; 1 means return transcoding failed, Take effect when IsCheckVideoBitrate is true.
        /// </summary>
        public readonly string? VideoBitrateAdjMethod;

        [OutputConstructor]
        private MediaTranscodeTemplateTransConfig(
            string? adjDarMethod,

            string? audioBitrateAdjMethod,

            string? deleteMetadata,

            Outputs.MediaTranscodeTemplateTransConfigHlsEncrypt? hlsEncrypt,

            string? isCheckAudioBitrate,

            string? isCheckReso,

            string? isCheckVideoBitrate,

            string? isHdr2Sdr,

            string? resoAdjMethod,

            string? videoBitrateAdjMethod)
        {
            AdjDarMethod = adjDarMethod;
            AudioBitrateAdjMethod = audioBitrateAdjMethod;
            DeleteMetadata = deleteMetadata;
            HlsEncrypt = hlsEncrypt;
            IsCheckAudioBitrate = isCheckAudioBitrate;
            IsCheckReso = isCheckReso;
            IsCheckVideoBitrate = isCheckVideoBitrate;
            IsHdr2Sdr = isHdr2Sdr;
            ResoAdjMethod = resoAdjMethod;
            VideoBitrateAdjMethod = videoBitrateAdjMethod;
        }
    }
}
