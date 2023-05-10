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

    public sealed class WorkflowMediaProcessTaskGetArgs : Pulumi.ResourceArgs
    {
        [Input("adaptiveDynamicStreamingTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetGetArgs>? _adaptiveDynamicStreamingTaskSets;

        /// <summary>
        /// Transfer Adaptive Code Stream Task List.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetGetArgs> AdaptiveDynamicStreamingTaskSets
        {
            get => _adaptiveDynamicStreamingTaskSets ?? (_adaptiveDynamicStreamingTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetGetArgs>());
            set => _adaptiveDynamicStreamingTaskSets = value;
        }

        [Input("animatedGraphicTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetGetArgs>? _animatedGraphicTaskSets;

        /// <summary>
        /// Video Rotation Map Task List.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetGetArgs> AnimatedGraphicTaskSets
        {
            get => _animatedGraphicTaskSets ?? (_animatedGraphicTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetGetArgs>());
            set => _animatedGraphicTaskSets = value;
        }

        [Input("imageSpriteTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskImageSpriteTaskSetGetArgs>? _imageSpriteTaskSets;

        /// <summary>
        /// Sprite image capture task list for video.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskImageSpriteTaskSetGetArgs> ImageSpriteTaskSets
        {
            get => _imageSpriteTaskSets ?? (_imageSpriteTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskImageSpriteTaskSetGetArgs>());
            set => _imageSpriteTaskSets = value;
        }

        [Input("sampleSnapshotTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetGetArgs>? _sampleSnapshotTaskSets;

        /// <summary>
        /// Screenshot task list for video sampling.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetGetArgs> SampleSnapshotTaskSets
        {
            get => _sampleSnapshotTaskSets ?? (_sampleSnapshotTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskSampleSnapshotTaskSetGetArgs>());
            set => _sampleSnapshotTaskSets = value;
        }

        [Input("snapshotByTimeOffsetTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs>? _snapshotByTimeOffsetTaskSets;

        /// <summary>
        /// Screenshot the task list of the video according to the time point.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs> SnapshotByTimeOffsetTaskSets
        {
            get => _snapshotByTimeOffsetTaskSets ?? (_snapshotByTimeOffsetTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs>());
            set => _snapshotByTimeOffsetTaskSets = value;
        }

        [Input("transcodeTaskSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetGetArgs>? _transcodeTaskSets;

        /// <summary>
        /// Video Transcoding Task List.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetGetArgs> TranscodeTaskSets
        {
            get => _transcodeTaskSets ?? (_transcodeTaskSets = new InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetGetArgs>());
            set => _transcodeTaskSets = value;
        }

        public WorkflowMediaProcessTaskGetArgs()
        {
        }
    }
}
