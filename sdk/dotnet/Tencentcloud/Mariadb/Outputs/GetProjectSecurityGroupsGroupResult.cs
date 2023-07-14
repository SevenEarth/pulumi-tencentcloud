// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Outputs
{

    [OutputType]
    public sealed class GetProjectSecurityGroupsGroupResult
    {
        /// <summary>
        /// Creation time in the format of yyyy-mm-dd hh:mm:ss.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Inbound rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectSecurityGroupsGroupInboundResult> Inbounds;
        /// <summary>
        /// Outbound rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectSecurityGroupsGroupOutboundResult> Outbounds;
        /// <summary>
        /// Project ID.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Security group ID.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// Security group name.
        /// </summary>
        public readonly string SecurityGroupName;
        /// <summary>
        /// Security group remarks.
        /// </summary>
        public readonly string SecurityGroupRemark;

        [OutputConstructor]
        private GetProjectSecurityGroupsGroupResult(
            string createTime,

            ImmutableArray<Outputs.GetProjectSecurityGroupsGroupInboundResult> inbounds,

            ImmutableArray<Outputs.GetProjectSecurityGroupsGroupOutboundResult> outbounds,

            int projectId,

            string securityGroupId,

            string securityGroupName,

            string securityGroupRemark)
        {
            CreateTime = createTime;
            Inbounds = inbounds;
            Outbounds = outbounds;
            ProjectId = projectId;
            SecurityGroupId = securityGroupId;
            SecurityGroupName = securityGroupName;
            SecurityGroupRemark = securityGroupRemark;
        }
    }
}
