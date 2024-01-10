// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Organization.Outputs
{

    [OutputType]
    public sealed class GetMembersItemResult
    {
        /// <summary>
        /// Security information binding status. Valid values: `Unbound`, `Valid`, `Success`, `Failed`.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string BindStatus;
        /// <summary>
        /// Creation timeNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string IsAllowQuit;
        /// <summary>
        /// Member type. Valid values: `Invite` (invited); `Create` (created).Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string MemberType;
        /// <summary>
        /// Member UINNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int MemberUin;
        /// <summary>
        /// Permission name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Node IDNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int NodeId;
        /// <summary>
        /// Node nameNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string NodeName;
        /// <summary>
        /// Management identityNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMembersItemOrgIdentityResult> OrgIdentities;
        /// <summary>
        /// Relationship policy permissionNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMembersItemOrgPermissionResult> OrgPermissions;
        /// <summary>
        /// Relationship policy nameNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string OrgPolicyName;
        /// <summary>
        /// Relationship policy typeNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string OrgPolicyType;
        /// <summary>
        /// Payer nameNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string PayName;
        /// <summary>
        /// Payer UINNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string PayUin;
        /// <summary>
        /// Member permission status. Valid values: `Confirmed`, `UnConfirmed`.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string PermissionStatus;
        /// <summary>
        /// RemarksNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Remark;
        /// <summary>
        /// Update timeNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetMembersItemResult(
            string bindStatus,

            string createTime,

            string isAllowQuit,

            string memberType,

            int memberUin,

            string name,

            int nodeId,

            string nodeName,

            ImmutableArray<Outputs.GetMembersItemOrgIdentityResult> orgIdentities,

            ImmutableArray<Outputs.GetMembersItemOrgPermissionResult> orgPermissions,

            string orgPolicyName,

            string orgPolicyType,

            string payName,

            string payUin,

            string permissionStatus,

            string remark,

            string updateTime)
        {
            BindStatus = bindStatus;
            CreateTime = createTime;
            IsAllowQuit = isAllowQuit;
            MemberType = memberType;
            MemberUin = memberUin;
            Name = name;
            NodeId = nodeId;
            NodeName = nodeName;
            OrgIdentities = orgIdentities;
            OrgPermissions = orgPermissions;
            OrgPolicyName = orgPolicyName;
            OrgPolicyType = orgPolicyType;
            PayName = payName;
            PayUin = payUin;
            PermissionStatus = permissionStatus;
            Remark = remark;
            UpdateTime = updateTime;
        }
    }
}
