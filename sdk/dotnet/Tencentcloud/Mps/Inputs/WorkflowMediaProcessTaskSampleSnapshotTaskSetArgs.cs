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

    public sealed class WorkflowMediaProcessTaskSampleSnapshotTaskSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sample screenshot template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// Rules for the `{number}` variable in the output path after sampling the screenshot.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("objectNumberFormat")]
        public Input<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetObjectNumberFormatArgs>? ObjectNumberFormat { get; set; }

        /// <summary>
        /// The output path of the image file after sampling the screenshot, which can be a relative path or an absolute path. If not filled, the default is a relative path: `{inputName}_sampleSnapshot_{definition}_{number}.{format}`.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// The target storage of the file after the screenshot at the time point, if not filled, it will inherit the OutputStorage value of the upper layer.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetOutputStorageArgs>? OutputStorage { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetWatermarkSetArgs>? _watermarkSets;

        /// <summary>
        /// Watermark list, support multiple pictures or text watermarks, up to 10.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetWatermarkSetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetWatermarkSetArgs>());
            set => _watermarkSets = value;
        }

        public WorkflowMediaProcessTaskSampleSnapshotTaskSetArgs()
        {
        }
    }
}
