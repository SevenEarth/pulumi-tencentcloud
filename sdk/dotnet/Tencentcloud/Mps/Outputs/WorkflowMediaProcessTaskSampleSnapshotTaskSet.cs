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
    public sealed class WorkflowMediaProcessTaskSampleSnapshotTaskSet
    {
        /// <summary>
        /// Sample screenshot template ID.
        /// </summary>
        public readonly int Definition;
        /// <summary>
        /// Rules for the `{number}` variable in the output path after sampling the screenshot.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetObjectNumberFormat? ObjectNumberFormat;
        /// <summary>
        /// The output path of the image file after sampling the screenshot, which can be a relative path or an absolute path. If not filled, the default is a relative path: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
        /// </summary>
        public readonly string? OutputObjectPath;
        /// <summary>
        /// The target storage of the file after the screenshot at the time point, if not filled, it will inherit the OutputStorage value of the upper layer.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetOutputStorage? OutputStorage;
        /// <summary>
        /// Watermark list, support multiple pictures or text watermarks, up to 10.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private WorkflowMediaProcessTaskSampleSnapshotTaskSet(
            int definition,

            Outputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetObjectNumberFormat? objectNumberFormat,

            string? outputObjectPath,

            Outputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetOutputStorage? outputStorage,

            ImmutableArray<Outputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetWatermarkSet> watermarkSets)
        {
            Definition = definition;
            ObjectNumberFormat = objectNumberFormat;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            WatermarkSets = watermarkSets;
        }
    }
}