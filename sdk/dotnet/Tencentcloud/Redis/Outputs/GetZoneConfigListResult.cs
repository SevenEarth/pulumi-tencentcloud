// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis.Outputs
{

    [OutputType]
    public sealed class GetZoneConfigListResult
    {
        /// <summary>
        /// (**Deprecated**) It has been deprecated from version 1.26.0. Use `shard_memories` instead. The memory volume of an available instance(in MB).
        /// </summary>
        public readonly ImmutableArray<int> MemSizes;
        /// <summary>
        /// The support numbers of instance copies.
        /// </summary>
        public readonly ImmutableArray<int> RedisReplicasNums;
        /// <summary>
        /// The support numbers of instance shard.
        /// </summary>
        public readonly ImmutableArray<int> RedisShardNums;
        /// <summary>
        /// The memory volume list of an available instance shard(in MB).
        /// </summary>
        public readonly ImmutableArray<int> ShardMemories;
        /// <summary>
        /// (**Deprecated**) It has been deprecated from version 1.33.1. Please use 'type_id' instead. Instance type. Available values: `master_slave_redis`, `master_slave_ckv`, `cluster_ckv`, `cluster_redis` and `standalone_redis`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Instance type ID.
        /// </summary>
        public readonly int TypeId;
        /// <summary>
        /// Version description of an available instance. Possible values: `Redis 3.2`, `Redis 4.0`.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// ID of available zone.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetZoneConfigListResult(
            ImmutableArray<int> memSizes,

            ImmutableArray<int> redisReplicasNums,

            ImmutableArray<int> redisShardNums,

            ImmutableArray<int> shardMemories,

            string type,

            int typeId,

            string version,

            string zone)
        {
            MemSizes = memSizes;
            RedisReplicasNums = redisReplicasNums;
            RedisShardNums = redisShardNums;
            ShardMemories = shardMemories;
            Type = type;
            TypeId = typeId;
            Version = version;
            Zone = zone;
        }
    }
}