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

    public sealed class WorkflowMediaProcessTaskArgs : global::Pulumi.ResourceArgs
    {
        [Input("adaptiveDynamicStreamingTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs>? _adaptiveDynamicStreamingTaskSets;

        /// <summary>
        /// Transfer Adaptive Code Stream Task List.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs> AdaptiveDynamicStreamingTaskSets
        {
            get => _adaptiveDynamicStreamingTaskSets ?? (_adaptiveDynamicStreamingTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetArgs>());
            set => _adaptiveDynamicStreamingTaskSets = value;
        }

        [Input("animatedGraphicTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetArgs>? _animatedGraphicTaskSets;

        /// <summary>
        /// Video Rotation Map Task List.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetArgs> AnimatedGraphicTaskSets
        {
            get => _animatedGraphicTaskSets ?? (_animatedGraphicTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetArgs>());
            set => _animatedGraphicTaskSets = value;
        }

        [Input("imageSpriteTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskImageSpriteTaskSetArgs>? _imageSpriteTaskSets;

        /// <summary>
        /// Sprite image capture task list for video.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskImageSpriteTaskSetArgs> ImageSpriteTaskSets
        {
            get => _imageSpriteTaskSets ?? (_imageSpriteTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskImageSpriteTaskSetArgs>());
            set => _imageSpriteTaskSets = value;
        }

        [Input("sampleSnapshotTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetArgs>? _sampleSnapshotTaskSets;

        /// <summary>
        /// Screenshot task list for video sampling.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetArgs> SampleSnapshotTaskSets
        {
            get => _sampleSnapshotTaskSets ?? (_sampleSnapshotTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetArgs>());
            set => _sampleSnapshotTaskSets = value;
        }

        [Input("snapshotByTimeOffsetTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetArgs>? _snapshotByTimeOffsetTaskSets;

        /// <summary>
        /// Screenshot the task list of the video according to the time point.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetArgs> SnapshotByTimeOffsetTaskSets
        {
            get => _snapshotByTimeOffsetTaskSets ?? (_snapshotByTimeOffsetTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetArgs>());
            set => _snapshotByTimeOffsetTaskSets = value;
        }

        [Input("transcodeTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetArgs>? _transcodeTaskSets;

        /// <summary>
        /// Video Transcoding Task List.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetArgs> TranscodeTaskSets
        {
            get => _transcodeTaskSets ?? (_transcodeTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetArgs>());
            set => _transcodeTaskSets = value;
        }

        public WorkflowMediaProcessTaskArgs()
        {
        }
        public static new WorkflowMediaProcessTaskArgs Empty => new WorkflowMediaProcessTaskArgs();
    }
}
