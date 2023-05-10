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
    public sealed class GetInstanceZoneInfoReplicaGroupResult
    {
        /// <summary>
        /// Node group ID.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// Node group Name.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Node group node list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceZoneInfoReplicaGroupRedisNodeResult> RedisNodes;
        /// <summary>
        /// The node group type, master is the primary node, and replica is the replica node.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// he availability zone ID of the node, such as ap-guangzhou-1.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetInstanceZoneInfoReplicaGroupResult(
            int groupId,

            string groupName,

            ImmutableArray<Outputs.GetInstanceZoneInfoReplicaGroupRedisNodeResult> redisNodes,

            string role,

            string zoneId)
        {
            GroupId = groupId;
            GroupName = groupName;
            RedisNodes = redisNodes;
            Role = role;
            ZoneId = zoneId;
        }
    }
}