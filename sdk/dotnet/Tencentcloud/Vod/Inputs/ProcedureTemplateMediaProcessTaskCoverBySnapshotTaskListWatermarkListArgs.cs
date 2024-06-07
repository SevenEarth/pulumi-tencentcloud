// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Watermarking template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        /// <summary>
        /// End time offset of a watermark in seconds. If this parameter is left blank or `0` is entered, the watermark will exist till the last video frame; If this value is greater than `0` (e.g., n), the watermark will exist till second n; If this value is smaller than `0` (e.g., -n), the watermark will exist till second n before the last video frame.
        /// </summary>
        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        /// <summary>
        /// Start time offset of a watermark in seconds. If this parameter is left blank or `0` is entered, the watermark will appear upon the first video frame. If this parameter is left blank or `0` is entered, the watermark will appear upon the first video frame; If this value is greater than `0` (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than `0` (e.g., -n), the watermark will appear at second n before the last video frame.
        /// </summary>
        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        /// <summary>
        /// SVG content of up to `2000000` characters. This needs to be entered only when the watermark type is `SVG`. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("svgContent")]
        public Input<string>? SvgContent { get; set; }

        /// <summary>
        /// Text content of up to `100` characters. This needs to be entered only when the watermark type is text. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("textContent")]
        public Input<string>? TextContent { get; set; }

        public ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListArgs()
        {
        }
        public static new ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListArgs Empty => new ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListArgs();
    }
}
