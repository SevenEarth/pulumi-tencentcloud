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

    public sealed class WatermarkTemplateTextTemplateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text transparency, value range: (0, 1].0: fully transparent.1: fully opaque.Default value: 1.
        /// </summary>
        [Input("fontAlpha", required: true)]
        public Input<double> FontAlpha { get; set; } = null!;

        /// <summary>
        /// Font color, format: 0xRRGGBB, default value: 0xFFFFFF (white).
        /// </summary>
        [Input("fontColor", required: true)]
        public Input<string> FontColor { get; set; } = null!;

        /// <summary>
        /// Font size, format: Npx, N is a number.
        /// </summary>
        [Input("fontSize", required: true)]
        public Input<string> FontSize { get; set; } = null!;

        /// <summary>
        /// Font type, currently supports two:simkai.ttf: can support Chinese and English.arial.ttf: English only.
        /// </summary>
        [Input("fontType", required: true)]
        public Input<string> FontType { get; set; } = null!;

        public WatermarkTemplateTextTemplateGetArgs()
        {
        }
    }
}
