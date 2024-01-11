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
    public sealed class ScheduleActivity
    {
        /// <summary>
        /// The parameters of a subtask.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityPara? ActivityPara;
        /// <summary>
        /// The subtask type. `input`: The start. `output`: The end. `action-trans`: Transcoding. `action-samplesnapshot`: Sampled screencapturing. `action-AIAnalysis`: Content analysis. `action-AIRecognition`: Content recognition. `action-aiReview`: Content moderation. `action-animated-graphics`: Animated screenshot generation. `action-image-sprite`: Image sprite generation. `action-snapshotByTimeOffset`: Time point screencapturing. `action-adaptive-substream`: Adaptive bitrate streaming.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ActivityType;
        /// <summary>
        /// The indexes of the subsequent actions. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<int> ReardriveIndices;

        [OutputConstructor]
        private ScheduleActivity(
            Outputs.ScheduleActivityActivityPara? activityPara,

            string activityType,

            ImmutableArray<int> reardriveIndices)
        {
            ActivityPara = activityPara;
            ActivityType = activityType;
            ReardriveIndices = reardriveIndices;
        }
    }
}