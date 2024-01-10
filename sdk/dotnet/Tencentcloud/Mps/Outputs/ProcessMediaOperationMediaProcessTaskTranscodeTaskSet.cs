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
    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSet
    {
        /// <summary>
        /// ID of a video transcoding template.
        /// </summary>
        public readonly int Definition;
        /// <summary>
        /// End time offset of a transcoded video, in seconds.If this parameter is left empty or set to 0, the transcoded video will end at the same time as the original video.If this parameter is set to a positive number (n for example), the transcoded video will end at the nth second of the original video.If this parameter is set to a negative number (-n for example), the transcoded video will end at the nth second before the end of the original video.
        /// </summary>
        public readonly double? EndTimeOffset;
        /// <summary>
        /// Opening and closing credits parametersNote: this field may return `null`, indicating that no valid value was found.
        /// </summary>
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameter? HeadTailParameter;
        /// <summary>
        /// List of blurs. Up to 10 ones can be supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSet> MosaicSets;
        /// <summary>
        /// Rule of the `{number}` variable in the output path after transcoding.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetObjectNumberFormat? ObjectNumberFormat;
        /// <summary>
        /// Path to a primary output file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}.{format}`.
        /// </summary>
        public readonly string? OutputObjectPath;
        /// <summary>
        /// Target bucket of an output file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorage? OutputStorage;
        /// <summary>
        /// Video transcoding custom parameter, which is valid when `Definition` is not 0.When any parameters in this structure are entered, they will be used to override corresponding parameters in templates.This parameter is used in highly customized scenarios. We recommend you only use `Definition` to specify the transcoding parameter.Note: this field may return `null`, indicating that no valid value was found.
        /// </summary>
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOverrideParameter? OverrideParameter;
        /// <summary>
        /// Custom video transcoding parameter, which is valid if `Definition` is 0.This parameter is used in highly customized scenarios. We recommend you use `Definition` to specify the transcoding parameter preferably.
        /// </summary>
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetRawParameter? RawParameter;
        /// <summary>
        /// Path to an output file part (the path to ts during transcoding to HLS), which can only be a relative path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_transcode_{definition}_{number}.{format}`.
        /// </summary>
        public readonly string? SegmentObjectName;
        /// <summary>
        /// Start time offset of a transcoded video, in seconds.If this parameter is left empty or set to 0, the transcoded video will start at the same time as the original video.If this parameter is set to a positive number (n for example), the transcoded video will start at the nth second of the original video.If this parameter is set to a negative number (-n for example), the transcoded video will start at the nth second before the end of the original video.
        /// </summary>
        public readonly double? StartTimeOffset;
        /// <summary>
        /// List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskTranscodeTaskSet(
            int definition,

            double? endTimeOffset,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameter? headTailParameter,

            ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetMosaicSet> mosaicSets,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetObjectNumberFormat? objectNumberFormat,

            string? outputObjectPath,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOutputStorage? outputStorage,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetOverrideParameter? overrideParameter,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetRawParameter? rawParameter,

            string? segmentObjectName,

            double? startTimeOffset,

            ImmutableArray<Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSet> watermarkSets)
        {
            Definition = definition;
            EndTimeOffset = endTimeOffset;
            HeadTailParameter = headTailParameter;
            MosaicSets = mosaicSets;
            ObjectNumberFormat = objectNumberFormat;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            OverrideParameter = overrideParameter;
            RawParameter = rawParameter;
            SegmentObjectName = segmentObjectName;
            StartTimeOffset = startTimeOffset;
            WatermarkSets = watermarkSets;
        }
    }
}
