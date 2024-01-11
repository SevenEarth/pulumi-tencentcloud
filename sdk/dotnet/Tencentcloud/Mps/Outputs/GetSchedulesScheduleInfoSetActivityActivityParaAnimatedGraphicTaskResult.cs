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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaAnimatedGraphicTaskResult
    {
        /// <summary>
        /// ID of a watermarking template.
        /// </summary>
        public readonly int Definition;
        /// <summary>
        /// End time offset of a watermark in seconds.If this parameter is left empty or 0 is entered, the watermark will exist till the last video frame;If this value is greater than 0 (e.g., n), the watermark will exist till second n;If this value is smaller than 0 (e.g., -n), the watermark will exist till second n before the last video frame.
        /// </summary>
        public readonly double EndTimeOffset;
        /// <summary>
        /// Path to a primary output file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}.{format}`.
        /// </summary>
        public readonly string OutputObjectPath;
        /// <summary>
        /// The bucket to save the output file.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAnimatedGraphicTaskOutputStorageResult> OutputStorages;
        /// <summary>
        /// Start time offset of a watermark in seconds. If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame.If this parameter is left empty or 0 is entered, the watermark will appear upon the first video frame;If this value is greater than 0 (e.g., n), the watermark will appear at second n after the first video frame;If this value is smaller than 0 (e.g., -n), the watermark will appear at second n before the last video frame.
        /// </summary>
        public readonly double StartTimeOffset;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaAnimatedGraphicTaskResult(
            int definition,

            double endTimeOffset,

            string outputObjectPath,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaAnimatedGraphicTaskOutputStorageResult> outputStorages,

            double startTimeOffset)
        {
            Definition = definition;
            EndTimeOffset = endTimeOffset;
            OutputObjectPath = outputObjectPath;
            OutputStorages = outputStorages;
            StartTimeOffset = startTimeOffset;
        }
    }
}