// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetViewsClusterViewResult
    {
        /// <summary>
        /// Average cpu utilization.
        /// </summary>
        public readonly double AvgCpuUsage;
        /// <summary>
        /// Average disk utilization.
        /// </summary>
        public readonly double AvgDiskUsage;
        /// <summary>
        /// Average memory utilization.
        /// </summary>
        public readonly double AvgMemUsage;
        /// <summary>
        /// Whether or not to break.
        /// </summary>
        public readonly double Break;
        /// <summary>
        /// Number of data nodes.
        /// </summary>
        public readonly int DataNodeNum;
        /// <summary>
        /// Bytes used on disk.
        /// </summary>
        public readonly int DiskUsedInBytes;
        /// <summary>
        /// Number of documents.
        /// </summary>
        public readonly int DocNum;
        /// <summary>
        /// Cluster health status.
        /// </summary>
        public readonly double Health;
        /// <summary>
        /// Index number.
        /// </summary>
        public readonly int IndexNum;
        /// <summary>
        /// Initializing shard number.
        /// </summary>
        public readonly int InitializingShardNum;
        /// <summary>
        /// Number of online nodes.
        /// </summary>
        public readonly int NodeNum;
        /// <summary>
        /// Primary shard number.
        /// </summary>
        public readonly int PrimaryShardNum;
        /// <summary>
        /// Relocating shard number.
        /// </summary>
        public readonly int RelocatingShardNum;
        /// <summary>
        /// Enterprise cluster can search the appid to which snapshot cos belongs.
        /// </summary>
        public readonly string SearchableSnapshotCosAppId;
        /// <summary>
        /// Enterprise cluster searchable bucket name stored in snapshot cos.
        /// </summary>
        public readonly string SearchableSnapshotCosBucket;
        /// <summary>
        /// Number of node fragments.
        /// </summary>
        public readonly int ShardNum;
        /// <summary>
        /// Client request node.
        /// </summary>
        public readonly ImmutableArray<string> TargetNodeTypes;
        /// <summary>
        /// Storage capacity of COS Enterprise Edition (in GB).
        /// </summary>
        public readonly int TotalCosStorage;
        /// <summary>
        /// Total storage size of cluster.
        /// </summary>
        public readonly int TotalDiskSize;
        /// <summary>
        /// Total number of nodes.
        /// </summary>
        public readonly int TotalNodeNum;
        /// <summary>
        /// Unassigned shard number.
        /// </summary>
        public readonly int UnassignedShardNum;
        /// <summary>
        /// Whether the node is visible.
        /// </summary>
        public readonly double Visible;

        [OutputConstructor]
        private GetViewsClusterViewResult(
            double avgCpuUsage,

            double avgDiskUsage,

            double avgMemUsage,

            double @break,

            int dataNodeNum,

            int diskUsedInBytes,

            int docNum,

            double health,

            int indexNum,

            int initializingShardNum,

            int nodeNum,

            int primaryShardNum,

            int relocatingShardNum,

            string searchableSnapshotCosAppId,

            string searchableSnapshotCosBucket,

            int shardNum,

            ImmutableArray<string> targetNodeTypes,

            int totalCosStorage,

            int totalDiskSize,

            int totalNodeNum,

            int unassignedShardNum,

            double visible)
        {
            AvgCpuUsage = avgCpuUsage;
            AvgDiskUsage = avgDiskUsage;
            AvgMemUsage = avgMemUsage;
            Break = @break;
            DataNodeNum = dataNodeNum;
            DiskUsedInBytes = diskUsedInBytes;
            DocNum = docNum;
            Health = health;
            IndexNum = indexNum;
            InitializingShardNum = initializingShardNum;
            NodeNum = nodeNum;
            PrimaryShardNum = primaryShardNum;
            RelocatingShardNum = relocatingShardNum;
            SearchableSnapshotCosAppId = searchableSnapshotCosAppId;
            SearchableSnapshotCosBucket = searchableSnapshotCosBucket;
            ShardNum = shardNum;
            TargetNodeTypes = targetNodeTypes;
            TotalCosStorage = totalCosStorage;
            TotalDiskSize = totalDiskSize;
            TotalNodeNum = totalNodeNum;
            UnassignedShardNum = unassignedShardNum;
            Visible = visible;
        }
    }
}
