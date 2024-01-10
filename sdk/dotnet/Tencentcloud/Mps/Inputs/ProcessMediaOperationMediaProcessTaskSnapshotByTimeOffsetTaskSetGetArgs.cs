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

    public sealed class ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of a time point screencapturing template.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        [Input("extTimeOffsetSets")]
        private InputList<string>? _extTimeOffsetSets;

        /// <summary>
        /// List of screenshot time points in the format of `s` or `%`:If the string ends in `s`, it means that the time point is in seconds; for example, `3.5s` means that the time point is the 3.5th second;If the string ends in `%`, it means that the time point is the specified percentage of the video duration; for example, `10%` means that the time point is 10% of the video duration.
        /// </summary>
        public InputList<string> ExtTimeOffsetSets
        {
            get => _extTimeOffsetSets ?? (_extTimeOffsetSets = new InputList<string>());
            set => _extTimeOffsetSets = value;
        }

        /// <summary>
        /// Rule of the `{number}` variable in the time point screenshot output path.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("objectNumberFormat")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetObjectNumberFormatGetArgs>? ObjectNumberFormat { get; set; }

        /// <summary>
        /// Output path to a generated time point screenshot, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_snapshotByTimeOffset_{definition}_{number}.{format}`.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// Target bucket of a generated time point screenshot file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetOutputStorageGetArgs>? OutputStorage { get; set; }

        [Input("timeOffsetSets")]
        private InputList<double>? _timeOffsetSets;

        /// <summary>
        /// List of time points of screenshots in &amp;lt;font color=red&amp;gt;seconds&amp;lt;/font&amp;gt;.
        /// </summary>
        public InputList<double> TimeOffsetSets
        {
            get => _timeOffsetSets ?? (_timeOffsetSets = new InputList<double>());
            set => _timeOffsetSets = value;
        }

        [Input("watermarkSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs>? _watermarkSets;

        /// <summary>
        /// List of up to 10 image or text watermarks.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetWatermarkSetGetArgs>());
            set => _watermarkSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs()
        {
        }
    }
}
