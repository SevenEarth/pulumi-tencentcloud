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

    public sealed class ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of a watermarking template.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// End time offset of a watermark in seconds.If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame;If this value is greater than 0 (e.g., n), the watermark will exist till second n;If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.
        /// </summary>
        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        /// <summary>
        /// Custom watermark parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the watermark parameter preferably.Custom watermark parameter is not available for screenshot.
        /// </summary>
        [Input("rawParameter")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetRawParameterGetArgs>? RawParameter { get; set; }

        /// <summary>
        /// Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame.If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame;If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame;If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.
        /// </summary>
        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        /// <summary>
        /// SVG content of up to 2,000,000 characters. This field is required only when the watermark type is `SVG`.SVG watermark is not available for screenshot.
        /// </summary>
        [Input("svgContent")]
        public Input<string>? SvgContent { get; set; }

        /// <summary>
        /// Text content of up to 100 characters. This field is required only when the watermark type is text.Text watermark is not available for screenshot.
        /// </summary>
        [Input("textContent")]
        public Input<string>? TextContent { get; set; }

        public ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs()
        {
        }
    }
}
