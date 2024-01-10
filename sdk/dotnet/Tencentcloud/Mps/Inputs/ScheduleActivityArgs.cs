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

    public sealed class ScheduleActivityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameters of a subtask.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("activityPara")]
        public Input<Inputs.ScheduleActivityActivityParaArgs>? ActivityPara { get; set; }

        /// <summary>
        /// The subtask type. `input`: The start. `output`: The end. `action-trans`: Transcoding. `action-samplesnapshot`: Sampled screencapturing. `action-AIAnalysis`: Content analysis. `action-AIRecognition`: Content recognition. `action-aiReview`: Content moderation. `action-animated-graphics`: Animated screenshot generation. `action-image-sprite`: Image sprite generation. `action-snapshotByTimeOffset`: Time point screencapturing. `action-adaptive-substream`: Adaptive bitrate streaming.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("activityType", required: true)]
        public Input<string> ActivityType { get; set; } = null!;

        [Input("reardriveIndices")]
        private InputList<int>? _reardriveIndices;

        /// <summary>
        /// The indexes of the subsequent actions. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<int> ReardriveIndices
        {
            get => _reardriveIndices ?? (_reardriveIndices = new InputList<int>());
            set => _reardriveIndices = value;
        }

        public ScheduleActivityArgs()
        {
        }
    }
}
