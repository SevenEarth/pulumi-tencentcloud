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
    public sealed class ScheduleActivityActivityParaAdaptiveDynamicStreamingTask
    {
        /// <summary>
        /// Subtitle files to insert.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskAddOnSubtitle> AddOnSubtitles;
        /// <summary>
        /// Adaptive bitrate streaming template ID.
        /// </summary>
        public readonly int Definition;
        /// <summary>
        /// The relative or absolute output path of the manifest file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}.{format}`.
        /// </summary>
        public readonly string? OutputObjectPath;
        /// <summary>
        /// Target bucket of an output file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorage? OutputStorage;
        /// <summary>
        /// The relative output path of the segment file after being transcoded to adaptive bitrate streaming (in HLS format only). If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`.
        /// </summary>
        public readonly string? SegmentObjectName;
        /// <summary>
        /// The relative output path of the substream file after being transcoded to adaptive bitrate streaming. If this parameter is left empty, a relative path in the following format will be used by default: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`.
        /// </summary>
        public readonly string? SubStreamObjectName;
        /// <summary>
        /// List of up to 10 image or text watermarks.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private ScheduleActivityActivityParaAdaptiveDynamicStreamingTask(
            ImmutableArray<Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskAddOnSubtitle> addOnSubtitles,

            int definition,

            string? outputObjectPath,

            Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorage? outputStorage,

            string? segmentObjectName,

            string? subStreamObjectName,

            ImmutableArray<Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSet> watermarkSets)
        {
            AddOnSubtitles = addOnSubtitles;
            Definition = definition;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            SegmentObjectName = segmentObjectName;
            SubStreamObjectName = subStreamObjectName;
            WatermarkSets = watermarkSets;
        }
    }
}