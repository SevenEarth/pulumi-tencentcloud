// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskTargetResourceTopicParam
    {
        /// <summary>
        /// Whether to perform compression when writing a topic, if it is not enabled, fill in none, if it is enabled, you can choose one of gzip, snappy, lz4 to fill in.
        /// </summary>
        public readonly string? CompressionType;
        /// <summary>
        /// 1 source topic message is amplified into msg Multiple and written to the target topic (this parameter is currently only applicable to ckafka flowing into ckafka).
        /// </summary>
        public readonly int? MsgMultiple;
        /// <summary>
        /// Offset type, initial position earliest, latest position latest, time point position timestamp.
        /// </summary>
        public readonly string? OffsetType;
        /// <summary>
        /// The topic name of the topic sold separately.
        /// </summary>
        public readonly string Resource;
        /// <summary>
        /// It must be passed when the Offset type is timestamp, and the time stamp is passed, accurate to the second.
        /// </summary>
        public readonly int? StartTime;
        /// <summary>
        /// TopicId.
        /// </summary>
        public readonly string? TopicId;
        /// <summary>
        /// whether the used topic need to be automatically created (currently only supports SOURCE inflow tasks).
        /// </summary>
        public readonly bool? UseAutoCreateTopic;

        [OutputConstructor]
        private DatahubTaskTargetResourceTopicParam(
            string? compressionType,

            int? msgMultiple,

            string? offsetType,

            string resource,

            int? startTime,

            string? topicId,

            bool? useAutoCreateTopic)
        {
            CompressionType = compressionType;
            MsgMultiple = msgMultiple;
            OffsetType = offsetType;
            Resource = resource;
            StartTime = startTime;
            TopicId = topicId;
            UseAutoCreateTopic = useAutoCreateTopic;
        }
    }
}
