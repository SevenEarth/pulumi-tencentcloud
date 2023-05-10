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

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterSubtitleTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text transparency, value range: (0, 1].0: fully transparent.1: fully opaque.Default: 1.
        /// </summary>
        [Input("fontAlpha")]
        public Input<double>? FontAlpha { get; set; }

        /// <summary>
        /// Font color, format: 0xRRGGBB, default value: 0xFFFFFF (white).
        /// </summary>
        [Input("fontColor")]
        public Input<string>? FontColor { get; set; }

        /// <summary>
        /// Font size, format: Npx, N is a value, if not specified, the subtitle file shall prevail.
        /// </summary>
        [Input("fontSize")]
        public Input<string>? FontSize { get; set; }

        /// <summary>
        /// Font type.hei.ttf, song.ttf, simkai.ttf, arial.ttf.Default: hei.ttf.
        /// </summary>
        [Input("fontType")]
        public Input<string>? FontType { get; set; }

        /// <summary>
        /// The address of the subtitle file to be compressed into the video.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the subtitle track to be compressed into the video. If there is a specified Path, the Path has a higher priority. Path and StreamIndex specify at least one.
        /// </summary>
        [Input("streamIndex")]
        public Input<int>? StreamIndex { get; set; }

        public WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterSubtitleTemplateArgs()
        {
        }
    }
}
