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

    public sealed class ScheduleActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The text transparency. Value range: 0-1. 0: Completely transparent 1: Completely opaqueDefault value: 1.
        /// </summary>
        [Input("fontAlpha")]
        public Input<double>? FontAlpha { get; set; }

        /// <summary>
        /// The font color in 0xRRGGBB format. Default value: 0xFFFFFF (white).
        /// </summary>
        [Input("fontColor")]
        public Input<string>? FontColor { get; set; }

        /// <summary>
        /// The font size (pixels). If this is not specified, the font size in the subtitle file will be used.
        /// </summary>
        [Input("fontSize")]
        public Input<string>? FontSize { get; set; }

        /// <summary>
        /// The font type. Valid values: `hei.ttf` `song.ttf` `simkai.ttf` `arial.ttf` (for English only). The default is `hei.ttf`.
        /// </summary>
        [Input("fontType")]
        public Input<string>? FontType { get; set; }

        /// <summary>
        /// The URL of the subtitles to add to the video.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The subtitle track to add to the video. If both `Path` and `StreamIndex` are specified, `Path` will be used. You need to specify at least one of the two parameters.
        /// </summary>
        [Input("streamIndex")]
        public Input<int>? StreamIndex { get; set; }

        public ScheduleActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplateGetArgs()
        {
        }
    }
}
