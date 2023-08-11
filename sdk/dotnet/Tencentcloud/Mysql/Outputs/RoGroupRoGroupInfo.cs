// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class RoGroupRoGroupInfo
    {
        /// <summary>
        /// The minimum number of reserved instances. It can be set to any value less than or equal to the number of RO instances under this RO group. Note that if the setting value is greater than the number of RO instances, it will not be removed; if it is set to 0, all instances whose latency exceeds the limit will be removed.
        /// </summary>
        public readonly int? MinRoInGroup;
        /// <summary>
        /// Delayed replication time.
        /// </summary>
        public readonly int? ReplicationDelayTime;
        /// <summary>
        /// RO group name.
        /// </summary>
        public readonly string? RoGroupName;
        /// <summary>
        /// RO instance maximum latency threshold. The unit is seconds, the minimum value is 1. Note that the RO group must have enabled instance delay culling policy for this value to be valid.
        /// </summary>
        public readonly int? RoMaxDelayTime;
        /// <summary>
        /// Whether to enable delayed culling of instances. Supported values are: 1 - on; 0 - not on. Note that if you enable instance delay culling, you must set the delay threshold (RoMaxDelayTime) parameter.
        /// </summary>
        public readonly int? RoOfflineDelay;
        /// <summary>
        /// weight mode. Supported values include: `system` - automatically assigned by the system; `custom` - user-defined settings. Note that if the `custom` mode is set, the RO instance weight configuration (RoWeightValues) parameter must be set.
        /// </summary>
        public readonly string? WeightMode;

        [OutputConstructor]
        private RoGroupRoGroupInfo(
            int? minRoInGroup,

            int? replicationDelayTime,

            string? roGroupName,

            int? roMaxDelayTime,

            int? roOfflineDelay,

            string? weightMode)
        {
            MinRoInGroup = minRoInGroup;
            ReplicationDelayTime = replicationDelayTime;
            RoGroupName = roGroupName;
            RoMaxDelayTime = roMaxDelayTime;
            RoOfflineDelay = roOfflineDelay;
            WeightMode = weightMode;
        }
    }
}