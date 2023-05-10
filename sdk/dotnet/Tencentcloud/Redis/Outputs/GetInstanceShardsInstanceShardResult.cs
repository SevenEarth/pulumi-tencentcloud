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
    public sealed class GetInstanceShardsInstanceShardResult
    {
        /// <summary>
        /// Service status: 0-down;1-on.
        /// </summary>
        public readonly int Connected;
        /// <summary>
        /// Number of keys.
        /// </summary>
        public readonly int Keys;
        /// <summary>
        /// role.
        /// </summary>
        public readonly int Role;
        /// <summary>
        /// The node ID of the instance runtime.
        /// </summary>
        public readonly string Runid;
        /// <summary>
        /// Shard node ID.
        /// </summary>
        public readonly string ShardId;
        /// <summary>
        /// Shard node name.
        /// </summary>
        public readonly string ShardName;
        /// <summary>
        /// Slot information.
        /// </summary>
        public readonly string Slots;
        /// <summary>
        /// Used capacity.
        /// </summary>
        public readonly int Storage;
        /// <summary>
        /// Capacity tilt.
        /// </summary>
        public readonly double StorageSlope;

        [OutputConstructor]
        private GetInstanceShardsInstanceShardResult(
            int connected,

            int keys,

            int role,

            string runid,

            string shardId,

            string shardName,

            string slots,

            int storage,

            double storageSlope)
        {
            Connected = connected;
            Keys = keys;
            Role = role;
            Runid = runid;
            ShardId = shardId;
            ShardName = shardName;
            Slots = slots;
            Storage = storage;
            StorageSlope = storageSlope;
        }
    }
}