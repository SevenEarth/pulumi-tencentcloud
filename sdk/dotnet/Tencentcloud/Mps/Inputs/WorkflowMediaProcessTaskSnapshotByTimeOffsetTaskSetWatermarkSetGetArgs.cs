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

    public sealed class WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Watermark Template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// End time offset of watermark, unit: second.Do not fill in or fill in 0, indicating that the watermark lasts until the end of the screen.When the value is greater than 0 (assumed to be n), it means that the watermark lasts until the nth second and disappears.When the value is less than 0 (assumed to be -n), it means that the watermark lasts until it disappears n seconds before the end of the screen.
        /// </summary>
        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        /// <summary>
        /// Watermark custom parameters, valid when Definition is filled with 0.This parameter is used in highly customized scenarios, it is recommended that you use Definition to specify watermark parameters first.Watermark custom parameters do not support screenshot watermarking.
        /// </summary>
        [Input("rawParameter")]
        public Input<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetRawParameterGetArgs>? RawParameter { get; set; }

        /// <summary>
        /// The start time offset of the watermark, unit: second. Do not fill in or fill in 0, which means that the watermark will start to appear when the screen appears.Do not fill in or fill in 0, which means the watermark will appear from the beginning of the screen.When the value is greater than 0 (assumed to be n), it means that the watermark appears from the nth second of the screen.When the value is less than 0 (assumed to be -n), it means that the watermark starts to appear n seconds before the end of the screen.
        /// </summary>
        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        /// <summary>
        /// SVG content. The length cannot exceed 2000000 characters. Fill in only if the watermark type is SVG watermark.SVG watermark does not support screenshot watermarking.
        /// </summary>
        [Input("svgContent")]
        public Input<string>? SvgContent { get; set; }

        /// <summary>
        /// Text content, the length does not exceed 100 characters. Fill in only when the watermark type is text watermark.Text watermark does not support screenshot watermarking.
        /// </summary>
        [Input("textContent")]
        public Input<string>? TextContent { get; set; }

        public WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs()
        {
        }
    }
}
