// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class ProcedureTemplateMediaProcessTaskTranscodeTaskListWatermarkList
    {
        /// <summary>
        /// Watermarking template ID.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// End time offset of a watermark in seconds. If this parameter is left blank or `0` is entered, the watermark will exist till the last video frame; If this value is greater than `0` (e.g., n), the watermark will exist till second n; If this value is smaller than `0` (e.g., -n), the watermark will exist till second n before the last video frame.
        /// </summary>
        public readonly double? EndTimeOffset;
        /// <summary>
        /// Start time offset of a watermark in seconds. If this parameter is left blank or `0` is entered, the watermark will appear upon the first video frame. If this parameter is left blank or `0` is entered, the watermark will appear upon the first video frame; If this value is greater than `0` (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than `0` (e.g., -n), the watermark will appear at second n before the last video frame.
        /// </summary>
        public readonly double? StartTimeOffset;
        /// <summary>
        /// SVG content of up to `2000000` characters. This needs to be entered only when the watermark type is `SVG`. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? SvgContent;
        /// <summary>
        /// Text content of up to `100` characters. This needs to be entered only when the watermark type is text. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? TextContent;

        [OutputConstructor]
        private ProcedureTemplateMediaProcessTaskTranscodeTaskListWatermarkList(
            string definition,

            double? endTimeOffset,

            double? startTimeOffset,

            string? svgContent,

            string? textContent)
        {
            Definition = definition;
            EndTimeOffset = endTimeOffset;
            StartTimeOffset = startTimeOffset;
            SvgContent = svgContent;
            TextContent = textContent;
        }
    }
}
