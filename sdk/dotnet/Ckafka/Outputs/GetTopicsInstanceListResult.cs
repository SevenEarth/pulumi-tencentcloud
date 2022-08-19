// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetTopicsInstanceListResult
    {
        /// <summary>
        /// Clear log policy, log clear mode. `delete`: logs are deleted according to the storage time, `compact`: logs are compressed according to the key, `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
        /// </summary>
        public readonly string CleanUpPolicy;
        /// <summary>
        /// Create time of the CKafka topic.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Whether to open the IP Whitelist. `true`: open, `false`: close.
        /// </summary>
        public readonly bool EnableWhiteList;
        /// <summary>
        /// Data backup cos bucket: the bucket address that is dumped to cos.
        /// </summary>
        public readonly string ForwardCosBucket;
        /// <summary>
        /// Periodic frequency of data backup to cos.
        /// </summary>
        public readonly int ForwardInterval;
        /// <summary>
        /// Data backup cos status. `1`: do not open data backup, `0`: open data backup.
        /// </summary>
        public readonly int ForwardStatus;
        /// <summary>
        /// IP Whitelist count.
        /// </summary>
        public readonly int IpWhiteListCount;
        /// <summary>
        /// Max message bytes.
        /// </summary>
        public readonly int MaxMessageBytes;
        /// <summary>
        /// CKafka topic note description.
        /// </summary>
        public readonly string Note;
        /// <summary>
        /// The number of partition.
        /// </summary>
        public readonly int PartitionNum;
        /// <summary>
        /// The number of replica.
        /// </summary>
        public readonly int ReplicaNum;
        /// <summary>
        /// Message can be selected. Retention time(unit ms).
        /// </summary>
        public readonly int Retention;
        /// <summary>
        /// Segment scrolling time, in ms.
        /// </summary>
        public readonly int Segment;
        /// <summary>
        /// Number of bytes rolled by shard.
        /// </summary>
        public readonly int SegmentBytes;
        /// <summary>
        /// Min number of sync replicas.
        /// </summary>
        public readonly int SyncReplicaMinNum;
        /// <summary>
        /// ID of the CKafka topic.
        /// </summary>
        public readonly string TopicId;
        /// <summary>
        /// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-). The length range is from 1 to 64.
        /// </summary>
        public readonly string TopicName;
        /// <summary>
        /// Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
        /// </summary>
        public readonly bool UncleanLeaderElectionEnable;

        [OutputConstructor]
        private GetTopicsInstanceListResult(
            string cleanUpPolicy,

            string createTime,

            bool enableWhiteList,

            string forwardCosBucket,

            int forwardInterval,

            int forwardStatus,

            int ipWhiteListCount,

            int maxMessageBytes,

            string note,

            int partitionNum,

            int replicaNum,

            int retention,

            int segment,

            int segmentBytes,

            int syncReplicaMinNum,

            string topicId,

            string topicName,

            bool uncleanLeaderElectionEnable)
        {
            CleanUpPolicy = cleanUpPolicy;
            CreateTime = createTime;
            EnableWhiteList = enableWhiteList;
            ForwardCosBucket = forwardCosBucket;
            ForwardInterval = forwardInterval;
            ForwardStatus = forwardStatus;
            IpWhiteListCount = ipWhiteListCount;
            MaxMessageBytes = maxMessageBytes;
            Note = note;
            PartitionNum = partitionNum;
            ReplicaNum = replicaNum;
            Retention = retention;
            Segment = segment;
            SegmentBytes = segmentBytes;
            SyncReplicaMinNum = syncReplicaMinNum;
            TopicId = topicId;
            TopicName = topicName;
            UncleanLeaderElectionEnable = uncleanLeaderElectionEnable;
        }
    }
}
