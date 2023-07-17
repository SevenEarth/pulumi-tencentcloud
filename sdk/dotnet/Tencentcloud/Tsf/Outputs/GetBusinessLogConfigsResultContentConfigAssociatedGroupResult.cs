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
    public sealed class GetBusinessLogConfigsResultContentConfigAssociatedGroupResult
    {
        /// <summary>
        /// Application Id of Group. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// Application Name. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ApplicationName;
        /// <summary>
        /// Application Type. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ApplicationType;
        /// <summary>
        /// Time when the deployment group is associated with the log configuration.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string AssociatedTime;
        /// <summary>
        /// Cluster ID to which the deployment group belongs.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Cluster Name to which the deployment group belongs.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// Cluster type to which the deployment group belongs.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// Group Id. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// Group Name. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Namespace ID to which the deployment group belongs.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// Namespace Name to which the deployment group belongs.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string NamespaceName;

        [OutputConstructor]
        private GetBusinessLogConfigsResultContentConfigAssociatedGroupResult(
            string applicationId,

            string applicationName,

            string applicationType,

            string associatedTime,

            string clusterId,

            string clusterName,

            string clusterType,

            string groupId,

            string groupName,

            string namespaceId,

            string namespaceName)
        {
            ApplicationId = applicationId;
            ApplicationName = applicationName;
            ApplicationType = applicationType;
            AssociatedTime = associatedTime;
            ClusterId = clusterId;
            ClusterName = clusterName;
            ClusterType = clusterType;
            GroupId = groupId;
            GroupName = groupName;
            NamespaceId = namespaceId;
            NamespaceName = namespaceName;
        }
    }
}
