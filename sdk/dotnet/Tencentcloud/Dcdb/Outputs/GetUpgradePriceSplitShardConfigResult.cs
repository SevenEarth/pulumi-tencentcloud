// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb.Outputs
{

    [OutputType]
    public sealed class GetUpgradePriceSplitShardConfigResult
    {
        /// <summary>
        /// List of shard ID.
        /// </summary>
        public readonly ImmutableArray<string> ShardInstanceIds;
        /// <summary>
        /// Shard memory size in GB.
        /// </summary>
        public readonly int ShardMemory;
        /// <summary>
        /// Shard storage capacity in GB.
        /// </summary>
        public readonly int ShardStorage;
        /// <summary>
        /// Data split ratio, fixed at 50%.
        /// </summary>
        public readonly int SplitRate;

        [OutputConstructor]
        private GetUpgradePriceSplitShardConfigResult(
            ImmutableArray<string> shardInstanceIds,

            int shardMemory,

            int shardStorage,

            int splitRate)
        {
            ShardInstanceIds = shardInstanceIds;
            ShardMemory = shardMemory;
            ShardStorage = shardStorage;
            SplitRate = splitRate;
        }
    }
}
