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

    public sealed class ProcessMediaOperationMediaProcessTaskGetArgs : Pulumi.ResourceArgs
    {
        [Input("adaptiveDynamicStreamingTaskSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetGetArgs>? _adaptiveDynamicStreamingTaskSets;

        /// <summary>
        /// List of adaptive bitrate streaming tasks.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetGetArgs> AdaptiveDynamicStreamingTaskSets
        {
            get => _adaptiveDynamicStreamingTaskSets ?? (_adaptiveDynamicStreamingTaskSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetGetArgs>());
            set => _adaptiveDynamicStreamingTaskSets = value;
        }

        [Input("animatedGraphicTaskSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskAnimatedGraphicTaskSetGetArgs>? _animatedGraphicTaskSets;

        /// <summary>
        /// List of animated image generating tasks.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskAnimatedGraphicTaskSetGetArgs> AnimatedGraphicTaskSets
        {
            get => _animatedGraphicTaskSets ?? (_animatedGraphicTaskSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskAnimatedGraphicTaskSetGetArgs>());
            set => _animatedGraphicTaskSets = value;
        }

        [Input("imageSpriteTaskSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskImageSpriteTaskSetGetArgs>? _imageSpriteTaskSets;

        /// <summary>
        /// List of image sprite generating tasks.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskImageSpriteTaskSetGetArgs> ImageSpriteTaskSets
        {
            get => _imageSpriteTaskSets ?? (_imageSpriteTaskSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskImageSpriteTaskSetGetArgs>());
            set => _imageSpriteTaskSets = value;
        }

        [Input("sampleSnapshotTaskSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetGetArgs>? _sampleSnapshotTaskSets;

        /// <summary>
        /// List of sampled screencapturing tasks.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetGetArgs> SampleSnapshotTaskSets
        {
            get => _sampleSnapshotTaskSets ?? (_sampleSnapshotTaskSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetGetArgs>());
            set => _sampleSnapshotTaskSets = value;
        }

        [Input("snapshotByTimeOffsetTaskSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs>? _snapshotByTimeOffsetTaskSets;

        /// <summary>
        /// List of time point screencapturing tasks.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs> SnapshotByTimeOffsetTaskSets
        {
            get => _snapshotByTimeOffsetTaskSets ?? (_snapshotByTimeOffsetTaskSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskSnapshotByTimeOffsetTaskSetGetArgs>());
            set => _snapshotByTimeOffsetTaskSets = value;
        }

        [Input("transcodeTaskSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs>? _transcodeTaskSets;

        /// <summary>
        /// List of transcoding tasks.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs> TranscodeTaskSets
        {
            get => _transcodeTaskSets ?? (_transcodeTaskSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetGetArgs>());
            set => _transcodeTaskSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskGetArgs()
        {
        }
    }
}
