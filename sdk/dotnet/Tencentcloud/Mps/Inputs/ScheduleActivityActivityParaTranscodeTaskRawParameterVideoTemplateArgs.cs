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

    public sealed class ScheduleActivityActivityParaTranscodeTaskRawParameterVideoTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The video bitrate (Kbps). Value range: 0 and [128, 35000].If the value is 0, the bitrate of the video will be the same as that of the source video.
        /// </summary>
        [Input("bitrate", required: true)]
        public Input<int> Bitrate { get; set; } = null!;

        /// <summary>
        /// The video codec. Valid values: `libx264`: H.264 `libx265`: H.265 `av1`: AOMedia Video 1Note: You must specify a resolution (not higher than 640 x 480) if the H.265 codec is used.Note: You can only use the AOMedia Video 1 codec for MP4 files.
        /// </summary>
        [Input("codec", required: true)]
        public Input<string> Codec { get; set; } = null!;

        /// <summary>
        /// The fill mode, which indicates how a video is resized when the video's original aspect ratio is different from the target aspect ratio. Valid values: stretch: Stretch the image frame by frame to fill the entire screen. The video image may become squashed or stretched after transcoding. black: Keep the image&amp;#39;s original aspect ratio and fill the blank space with black bars. white: Keep the image's original aspect ratio and fill the blank space with white bars. gauss: Keep the image's original aspect ratio and apply Gaussian blur to the blank space.Default value: black.Note: Only `stretch` and `black` are supported for adaptive bitrate streaming.
        /// </summary>
        [Input("fillType")]
        public Input<string>? FillType { get; set; }

        /// <summary>
        /// The video frame rate (Hz). Value range: [0, 100].If the value is 0, the frame rate will be the same as that of the source video.Note: For adaptive bitrate streaming, the value range of this parameter is [0, 60].
        /// </summary>
        [Input("fps", required: true)]
        public Input<int> Fps { get; set; } = null!;

        /// <summary>
        /// Frame interval between I keyframes. Value range: 0 and [1,100000].If this parameter is 0 or left empty, the system will automatically set the GOP length.
        /// </summary>
        [Input("gop")]
        public Input<int>? Gop { get; set; }

        /// <summary>
        /// Maximum value of the height (or short side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Resolution adaption. Valid values: open: Enabled. When resolution adaption is enabled, `Width` indicates the long side of a video, while `Height` indicates the short side. close: Disabled. When resolution adaption is disabled, `Width` indicates the width of a video, while `Height` indicates the height.Default value: open.Note: When resolution adaption is enabled, `Width` cannot be smaller than `Height`.
        /// </summary>
        [Input("resolutionAdaptive")]
        public Input<string>? ResolutionAdaptive { get; set; }

        /// <summary>
        /// The control factor of video constant bitrate. Value range: [1, 51]If this parameter is specified, CRF (a bitrate control method) will be used for transcoding. (Video bitrate will no longer take effect.)It is not recommended to specify this parameter if there are no special requirements.
        /// </summary>
        [Input("vcrf")]
        public Input<int>? Vcrf { get; set; }

        /// <summary>
        /// Maximum value of the width (or long side) of a video stream in px. Value range: 0 and [128, 4,096]. If both `Width` and `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0, `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled; If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public ScheduleActivityActivityParaTranscodeTaskRawParameterVideoTemplateArgs()
        {
        }
        public static new ScheduleActivityActivityParaTranscodeTaskRawParameterVideoTemplateArgs Empty => new ScheduleActivityActivityParaTranscodeTaskRawParameterVideoTemplateArgs();
    }
}
