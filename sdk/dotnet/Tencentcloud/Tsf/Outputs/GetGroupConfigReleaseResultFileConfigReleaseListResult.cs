// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class GetGroupConfigReleaseResultFileConfigReleaseListResult
    {
        /// <summary>
        /// Configuration item release cluster ID.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Configuration item release cluster name.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// Configuration item  ID.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ConfigId;
        /// <summary>
        /// Configuration item name.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ConfigName;
        /// <summary>
        /// Configuration item release ID.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ConfigReleaseId;
        /// <summary>
        /// Configuration version.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ConfigVersion;
        /// <summary>
        /// groupId.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// Configuration item release group name.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Configuration item release namespace ID.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// Configuration item release namespace name.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string NamespaceName;
        /// <summary>
        /// Configuration item release description.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ReleaseDesc;
        /// <summary>
        /// Configuration item release time.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ReleaseTime;

        [OutputConstructor]
        private GetGroupConfigReleaseResultFileConfigReleaseListResult(
            string clusterId,

            string clusterName,

            string configId,

            string configName,

            string configReleaseId,

            string configVersion,

            string groupId,

            string groupName,

            string namespaceId,

            string namespaceName,

            string releaseDesc,

            string releaseTime)
        {
            ClusterId = clusterId;
            ClusterName = clusterName;
            ConfigId = configId;
            ConfigName = configName;
            ConfigReleaseId = configReleaseId;
            ConfigVersion = configVersion;
            GroupId = groupId;
            GroupName = groupName;
            NamespaceId = namespaceId;
            NamespaceName = namespaceName;
            ReleaseDesc = releaseDesc;
            ReleaseTime = releaseTime;
        }
    }
}
