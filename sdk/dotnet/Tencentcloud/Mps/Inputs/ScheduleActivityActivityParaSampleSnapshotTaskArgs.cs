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

    public sealed class ScheduleActivityActivityParaSampleSnapshotTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sampled screencapturing template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// Rule of the `{number}` variable in the sampled screenshot output path.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("objectNumberFormat")]
        public Input<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormatArgs>? ObjectNumberFormat { get; set; }

        /// <summary>
        /// Output path to a generated sampled screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// Target bucket of a sampled screenshot. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskOutputStorageArgs>? OutputStorage { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetArgs>? _watermarkSets;

        /// <summary>
        /// List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetArgs>());
            set => _watermarkSets = value;
        }

        public ScheduleActivityActivityParaSampleSnapshotTaskArgs()
        {
        }
        public static new ScheduleActivityActivityParaSampleSnapshotTaskArgs Empty => new ScheduleActivityActivityParaSampleSnapshotTaskArgs();
    }
}
