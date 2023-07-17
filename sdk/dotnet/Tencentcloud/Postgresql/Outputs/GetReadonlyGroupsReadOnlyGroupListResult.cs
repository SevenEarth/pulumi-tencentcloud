// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql.Outputs
{

    [OutputType]
    public sealed class GetReadonlyGroupsReadOnlyGroupListResult
    {
        /// <summary>
        /// instance network connection information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReadonlyGroupsReadOnlyGroupListDbInstanceNetInfoResult> DbInstanceNetInfos;
        /// <summary>
        /// Master instance information, only returned when the instance is read-onlyNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string MasterDbInstanceId;
        /// <summary>
        /// delay time size threshold.
        /// </summary>
        public readonly double MaxReplayLag;
        /// <summary>
        /// delay space size threshold.
        /// </summary>
        public readonly int MaxReplayLatency;
        /// <summary>
        /// Minimum Number of Reserved InstancesNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int MinDelayEliminateReserve;
        /// <summary>
        /// Instance network information list (this field is obsolete)Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReadonlyGroupsReadOnlyGroupListNetworkAccessListResult> NetworkAccessLists;
        /// <summary>
        /// project ID.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// instance details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReadonlyGroupsReadOnlyGroupListReadOnlyDbInstanceListResult> ReadOnlyDbInstanceLists;
        /// <summary>
        /// read-only group idNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string ReadOnlyGroupId;
        /// <summary>
        /// read-only group nameNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string ReadOnlyGroupName;
        /// <summary>
        /// automatic load balancing switch.
        /// </summary>
        public readonly int Rebalance;
        /// <summary>
        /// region id.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// delay time switch.
        /// </summary>
        public readonly int ReplayLagEliminate;
        /// <summary>
        /// delay size switch.
        /// </summary>
        public readonly int ReplayLatencyEliminate;
        /// <summary>
        /// state.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// subnet-idNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// virtual network id.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// region id.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetReadonlyGroupsReadOnlyGroupListResult(
            ImmutableArray<Outputs.GetReadonlyGroupsReadOnlyGroupListDbInstanceNetInfoResult> dbInstanceNetInfos,

            string masterDbInstanceId,

            double maxReplayLag,

            int maxReplayLatency,

            int minDelayEliminateReserve,

            ImmutableArray<Outputs.GetReadonlyGroupsReadOnlyGroupListNetworkAccessListResult> networkAccessLists,

            int projectId,

            ImmutableArray<Outputs.GetReadonlyGroupsReadOnlyGroupListReadOnlyDbInstanceListResult> readOnlyDbInstanceLists,

            string readOnlyGroupId,

            string readOnlyGroupName,

            int rebalance,

            string region,

            int replayLagEliminate,

            int replayLatencyEliminate,

            string status,

            string subnetId,

            string vpcId,

            string zone)
        {
            DbInstanceNetInfos = dbInstanceNetInfos;
            MasterDbInstanceId = masterDbInstanceId;
            MaxReplayLag = maxReplayLag;
            MaxReplayLatency = maxReplayLatency;
            MinDelayEliminateReserve = minDelayEliminateReserve;
            NetworkAccessLists = networkAccessLists;
            ProjectId = projectId;
            ReadOnlyDbInstanceLists = readOnlyDbInstanceLists;
            ReadOnlyGroupId = readOnlyGroupId;
            ReadOnlyGroupName = readOnlyGroupName;
            Rebalance = rebalance;
            Region = region;
            ReplayLagEliminate = replayLagEliminate;
            ReplayLatencyEliminate = replayLatencyEliminate;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
            Zone = zone;
        }
    }
}
