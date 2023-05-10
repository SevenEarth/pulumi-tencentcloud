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
    public sealed class ApiGroupBindedGatewayDeployGroup
    {
        /// <summary>
        /// application ID.
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// Application Name.
        /// </summary>
        public readonly string? ApplicationName;
        /// <summary>
        /// Application classification: V: virtual machine application, C: container application.
        /// </summary>
        public readonly string? ApplicationType;
        /// <summary>
        /// Cluster type, C: container, V: virtual machine.
        /// </summary>
        public readonly string? ClusterType;
        /// <summary>
        /// Gateway deployment group ID.
        /// </summary>
        public readonly string? DeployGroupId;
        /// <summary>
        /// Gateway deployment group name.
        /// </summary>
        public readonly string? DeployGroupName;
        /// <summary>
        /// Deployment group application status, values: Running, Waiting, Paused, Updating, RollingBack, Abnormal, Unknown.
        /// </summary>
        public readonly string? GroupStatus;

        [OutputConstructor]
        private ApiGroupBindedGatewayDeployGroup(
            string? applicationId,

            string? applicationName,

            string? applicationType,

            string? clusterType,

            string? deployGroupId,

            string? deployGroupName,

            string? groupStatus)
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
