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
    public sealed class ScheduleActivityActivityParaTranscodeTaskOverrideParameterVideoTemplate
    {
        /// <summary>
        /// The video bitrate (Kbps). Value range: 0 and [128, 35000].If the value is 0, the bitrate of the video will be the same as that of the source video.
        /// </summary>
        public readonly int? Bitrate;
        /// <summary>
        /// The video codec. Valid values: `libx264`: H.264 `libx265`: H.265 `av1`: AOMedia Video 1Note: You must specify a resolution (not higher than 640 x 480) if the H.265 codec is used.Note: You can only use the AOMedia Video 1 codec for MP4 files.
        /// </summary>
        public readonly string? Codec;
        /// <summary>
        /// Whether to enable adaptive encoding. Valid values: 0: Disable 1: EnableDefault value: 0. If this parameter is set to `1`, multiple streams with different resolutions and bitrates will be generated automatically. The highest resolution, bitrate, and quality of the streams are determined by the values of `width` and `height`, `Bitrate`, and `Vcrf` in `VideoTemplate` respectively. If these parameters are not set in `VideoTemplate`, the highest resolution generated will be the same as that of the source video, and the highest video quality will be close to VMAF 95. To use this parameter or learn about the billing details of adaptive encoding, please contact your sales rep.
        /// </summary>
        public readonly int? ContentAdaptStream;
        /// <summary>
        /// The fill mode, which indicates how a video is resized when the video's original aspect ratio is different from the target aspect ratio. Valid values: stretch: Stretch the image frame by frame to fill the entire screen. The video image may become squashed or stretched after transcoding. black: Keep the image&amp;#39;s original aspect ratio and fill the blank space with black bars. white: Keep the image's original aspect ratio and fill the blank space with white bars. gauss: Keep the image's original aspect ratio and apply Gaussian blur to the blank space.Default value: black.Note: Only `stretch` and `black` are supported for adaptive bitrate streaming.
        /// </summary>
        public readonly string? FillType;
        /// <summary>
        /// The video frame rate (Hz). Value range: [0, 100].If the value is 0, the frame rate will be the same as that of the source video.Note: For adaptive bitrate streaming, the value range of this parameter is [0, 60].
        /// </summary>
        public readonly int? Fps;
        /// <summary>
        /// Frame interval between I keyframes. Value range: 0 and [1,100000].If this parameter is 0 or left empty, the system will automatically set the GOP length.
        /// </summary>
        public readonly int? Gop;
        /// <summary>
        /// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
        /// </summary>
        public readonly int? Height;
        /// <summary>
        /// Resolution adaption. Valid values: open: Enabled. When resolution adaption is enabled, `Width` indicates the long side of a video, while `Height` indicates the short side. close: Disabled. When resolution adaption is disabled, `Width` indicates the width of a video, while `Height` indicates the height.Default value: open.Note: When resolution adaption is enabled, `Width` cannot be smaller than `Height`.
        /// </summary>
        public readonly string? ResolutionAdaptive;
        /// <summary>
        /// The control factor of video constant bitrate. Value range: [1, 51]If this parameter is specified, CRF (a bitrate control method) will be used for transcoding. (Video bitrate will no longer take effect.)It is not recommended to specify this parameter if there are no special requirements.
        /// </summary>
        public readonly int? Vcrf;
        /// <summary>
        /// Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
        /// </summary>
        public readonly int? Width;

        [OutputConstructor]
        private ScheduleActivityActivityParaTranscodeTaskOverrideParameterVideoTemplate(
            int? bitrate,

            string? codec,

            int? contentAdaptStream,

            string? fillType,

            int? fps,

            int? gop,

            int? height,

            string? resolutionAdaptive,

            int? vcrf,

            int? width)
        {
            Bitrate = bitrate;
            Codec = codec;
            ContentAdaptStream = contentAdaptStream;
            FillType = fillType;
            Fps = fps;
            Gop = gop;
            Height = height;
            ResolutionAdaptive = resolutionAdaptive;
            Vcrf = vcrf;
            Width = width;
        }
    }
}
