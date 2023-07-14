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
    public sealed class GetDeliveryConfigsResultContentGroupResult
    {
        /// <summary>
        /// Associate Time. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string AssociateTime;
        /// <summary>
        /// Cluster ID. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Cluster Name. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// Cluster type.
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// Group Id.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// Group Name.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Namespace Name. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string NamespaceName;

        [OutputConstructor]
        private GetDeliveryConfigsResultContentGroupResult(
            string associateTime,

            string clusterId,

            string clusterName,

            string clusterType,

            string groupId,

            string groupName,

            string namespaceName)
        {
            AssociateTime = associateTime;
            ClusterId = clusterId;
            ClusterName = clusterName;
            ClusterType = clusterType;
            GroupId = groupId;
            GroupName = groupName;
            NamespaceName = namespaceName;
        }
    }
}
