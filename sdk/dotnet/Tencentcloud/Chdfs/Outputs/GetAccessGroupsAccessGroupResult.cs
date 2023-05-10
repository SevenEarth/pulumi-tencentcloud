// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Chdfs.Outputs
{

    [OutputType]
    public sealed class GetAccessGroupsAccessGroupResult
    {
        /// <summary>
        /// access group id.
        /// </summary>
        public readonly string AccessGroupId;
        /// <summary>
        /// access group name.
        /// </summary>
        public readonly string AccessGroupName;
        /// <summary>
        /// create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// access group description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// vpc network type(1:CVM, 2:BM 1.0).
        /// </summary>
        public readonly int VpcType;

        [OutputConstructor]
        private GetAccessGroupsAccessGroupResult(
            string accessGroupId,

            string accessGroupName,

            string createTime,

            string description,

            string vpcId,

            int vpcType)
        {
            AccessGroupId = accessGroupId;
            AccessGroupName = accessGroupName;
            CreateTime = createTime;
            Description = description;
            VpcId = vpcId;
            VpcType = vpcType;
        }
    }
}