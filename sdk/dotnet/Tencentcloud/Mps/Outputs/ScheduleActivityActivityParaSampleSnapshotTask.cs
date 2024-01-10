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
    public sealed class ScheduleActivityActivityParaSampleSnapshotTask
    {
        /// <summary>
        /// Sampled screencapturing template ID.
        /// </summary>
        public readonly int Definition;
        /// <summary>
        /// Rule of the `{number}` variable in the sampled screenshot output path.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormat? ObjectNumberFormat;
        /// <summary>
        /// Output path to a generated sampled screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
        /// </summary>
        public readonly string? OutputObjectPath;
        /// <summary>
        /// Target bucket of a sampled screenshot. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaSampleSnapshotTaskOutputStorage? OutputStorage;
        /// <summary>
        /// List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private ScheduleActivityActivityParaSampleSnapshotTask(
            int definition,

            Outputs.ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormat? objectNumberFormat,

            string? outputObjectPath,

            Outputs.ScheduleActivityActivityParaSampleSnapshotTaskOutputStorage? outputStorage,

            ImmutableArray<Outputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSet> watermarkSets)
        {
            Definition = definition;
            ObjectNumberFormat = objectNumberFormat;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            WatermarkSets = watermarkSets;
        }
    }
}
