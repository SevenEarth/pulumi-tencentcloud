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

    public sealed class WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Transfer Adaptive Code Stream Template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// After converting to an adaptive stream, the output path of the manifest file can be a relative path or an absolute path. If not filled, the default is a relative path: `{inputName}_adaptiveDynamicStreaming_{definition}.{format}`.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// The target storage of the file after converting to the adaptive code stream, if not filled, it will inherit the OutputStorage value of the upper layer.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetOutputStorageArgs>? OutputStorage { get; set; }

        /// <summary>
        /// After converting to an adaptive stream (only HLS), the output path of the fragmented file can only be a relative path. If not filled, the default is a relative path: `{inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}_{segmentNumber}.{format}`.
        /// </summary>
        [Input("segmentObjectName")]
        public Input<string>? SegmentObjectName { get; set; }

        /// <summary>
        /// After converting to an adaptive stream, the output path of the sub-stream file can only be a relative path. If not filled, the default is a relative path: {inputName}_adaptiveDynamicStreaming_{definition}_{subStreamNumber}.{format}`.
        /// </summary>
        [Input("subStreamObjectName")]
        public Input<string>? SubStreamObjectName { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetArgs>? _watermarkSets;

        /// <summary>
        /// Watermark list, support multiple pictures or text watermarks, up to 10.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetArgs>());
            set => _watermarkSets = value;
        }

        public WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs()
        {
        }
        public static new WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs Empty => new WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs();
    }
}
