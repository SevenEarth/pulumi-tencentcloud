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

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterVideoTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bit rate of the video stream, value range: 0 and [128, 35000], unit: kbps.When the value is 0, it means that the video bit rate is consistent with the original video.
        /// </summary>
        [Input("bitrate")]
        public Input<int>? Bitrate { get; set; }

        /// <summary>
        /// Encoding format of the video stream, optional value:libx264: H.264 encoding.libx265: H.265 encoding.av1: AOMedia Video 1 encoding.Note: Currently H.265 encoding must specify a resolution, and it needs to be within 640*480.Note: av1 encoded containers currently only support mp4.
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        /// <summary>
        /// Content Adaptive Encoding. optional value:0: not open.1: open.Default: 0.When this parameter is turned on, multiple code streams with different resolutions and different bit rates will be adaptively generated. The width and height of the VideoTemplate are the maximum resolutions among the multiple code streams, and the bit rates in the VideoTemplate are multiple code rates. The highest bit rate in the stream, the vcrf in VideoTemplate is the highest quality among multiple bit streams. When the resolution, bit rate and vcrf are not set, the highest resolution generated by the ContentAdaptStream parameter is the resolution of the video source, and the video quality is close to vmaf95. To enable this parameter or learn about billing details, please contact your Tencent Cloud Business.
        /// </summary>
        [Input("contentAdaptStream")]
        public Input<int>? ContentAdaptStream { get; set; }

        /// <summary>
        /// Filling method, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling method:stretch: Stretch, stretch each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched;.black: Leave black, keep the aspect ratio of the video unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video unchanged, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and fill the rest of the edge with Gaussian blur.Default: black.Note: Adaptive stream only supports stretch, black.
        /// </summary>
        [Input("fillType")]
        public Input<string>? FillType { get; set; }

        /// <summary>
        /// Video frame rate, value range: [0, 100], unit: Hz.When the value is 0, it means that the frame rate is consistent with the original video.Note: The value range for adaptive code rate is [0, 60].
        /// </summary>
        [Input("fps")]
        public Input<int>? Fps { get; set; }

        /// <summary>
        /// The interval between keyframe I frames, value range: 0 and [1, 100000], unit: number of frames.When filling 0 or not filling, the system will automatically set the gop length.
        /// </summary>
        [Input("gop")]
        public Input<int>? Gop { get; set; }

        /// <summary>
        /// The maximum value of video stream height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default: 0.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Adaptive resolution, optional values:```open: open, at this time, Width represents the long side of the video, Height represents the short side of the video.close: close, at this time, Width represents the width of the video, and Height represents the height of the video.Default: open.Note: In adaptive mode, Width cannot be smaller than Height.
        /// </summary>
        [Input("resolutionAdaptive")]
        public Input<string>? ResolutionAdaptive { get; set; }

        /// <summary>
        /// Video constant bit rate control factor, the value range is [1, 51].If this parameter is specified, the code rate control method of CRF will be used for transcoding (the video code rate will no longer take effect).If there is no special requirement, it is not recommended to specify this parameter.
        /// </summary>
        [Input("vcrf")]
        public Input<int>? Vcrf { get; set; }

        /// <summary>
        /// The maximum value of video stream width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default: 0.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterVideoTemplateGetArgs()
        {
        }
        public static new WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterVideoTemplateGetArgs Empty => new WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterVideoTemplateGetArgs();
    }
}
