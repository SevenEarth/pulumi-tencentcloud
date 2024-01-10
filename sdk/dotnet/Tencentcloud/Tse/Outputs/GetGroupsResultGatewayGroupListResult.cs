// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Outputs
{

    [OutputType]
    public sealed class GetGroupsResultGatewayGroupListResult
    {
        /// <summary>
        /// associated strategy informationNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsResultGatewayGroupListBindingStrategyResult> BindingStrategies;
        /// <summary>
        /// group create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// group description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// gateway ID.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// group Id.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// public network outbound traffic bandwidth.
        /// </summary>
        public readonly int InternetMaxBandwidthOut;
        /// <summary>
        /// whether it is the default group- 0: false.- 1: yes.
        /// </summary>
        public readonly int IsFirstGroup;
        /// <summary>
        /// modify time.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// filter name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// group node configration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsResultGatewayGroupListNodeConfigResult> NodeConfigs;
        /// <summary>
        /// group status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// subnet IDs.
        /// </summary>
        public readonly string SubnetIds;

        [OutputConstructor]
        private GetGroupsResultGatewayGroupListResult(
            ImmutableArray<Outputs.GetGroupsResultGatewayGroupListBindingStrategyResult> bindingStrategies,

            string createTime,

            string description,

            string gatewayId,

            string groupId,

            int internetMaxBandwidthOut,

            int isFirstGroup,

            string modifyTime,

            string name,

            ImmutableArray<Outputs.GetGroupsResultGatewayGroupListNodeConfigResult> nodeConfigs,

            string status,

            string subnetIds)
        {
            BindingStrategies = bindingStrategies;
            CreateTime = createTime;
            Description = description;
            GatewayId = gatewayId;
            GroupId = groupId;
            InternetMaxBandwidthOut = internetMaxBandwidthOut;
            IsFirstGroup = isFirstGroup;
            ModifyTime = modifyTime;
            Name = name;
            NodeConfigs = nodeConfigs;
            Status = status;
            SubnetIds = subnetIds;
        }
    }
}
