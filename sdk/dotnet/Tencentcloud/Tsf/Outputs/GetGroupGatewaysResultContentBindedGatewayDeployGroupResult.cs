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
    public sealed class GetGroupGatewaysResultContentBindedGatewayDeployGroupResult
    {
        /// <summary>
        /// application ID.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// application name.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ApplicationName;
        /// <summary>
        /// Application category: V: virtual machine application, C: container application.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ApplicationType;
        /// <summary>
        /// Cluster type, with possible values: C: container, V: virtual machine.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// Gateway deployment group ID.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string DeployGroupId;
        /// <summary>
        /// Gateway deployment group name.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string DeployGroupName;
        /// <summary>
        /// Application status of the deployment group, with possible values: Running, Waiting, Paused, Updating, RollingBack, Abnormal, Unknown.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string GroupStatus;

        [OutputConstructor]
        private GetGroupGatewaysResultContentBindedGatewayDeployGroupResult(
            string applicationId,

            string applicationName,

            string applicationType,

            string clusterType,

            string deployGroupId,

            string deployGroupName,

            string groupStatus)
        {
            ApplicationId = applicationId;
            ApplicationName = applicationName;
            ApplicationType = applicationType;
            ClusterType = clusterType;
            DeployGroupId = deployGroupId;
            DeployGroupName = deployGroupName;
            GroupStatus = groupStatus;
        }
    }
}