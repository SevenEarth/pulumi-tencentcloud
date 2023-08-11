// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class AlarmAlarmTarget
    {
        /// <summary>
        /// search end time of offset.
        /// </summary>
        public readonly int EndTimeOffset;
        /// <summary>
        /// logset id.
        /// </summary>
        public readonly string LogsetId;
        /// <summary>
        /// the number of alarm object.
        /// </summary>
        public readonly int Number;
        /// <summary>
        /// query rules.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// search start time of offset.
        /// </summary>
        public readonly int StartTimeOffset;
        /// <summary>
        /// topic id.
        /// </summary>
        public readonly string TopicId;

        [OutputConstructor]
        private AlarmAlarmTarget(
            int endTimeOffset,

            string logsetId,

            int number,

            string query,

            int startTimeOffset,

            string topicId)
        {
            EndTimeOffset = endTimeOffset;
            LogsetId = logsetId;
            Number = number;
            Query = query;
            StartTimeOffset = startTimeOffset;
            TopicId = topicId;
        }
    }
}