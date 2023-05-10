// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As.Outputs
{

    [OutputType]
    public sealed class GetInstancesInstanceListResult
    {
        /// <summary>
        /// The time when the instance joined the group.
        /// </summary>
        public readonly string AddTime;
        /// <summary>
        /// Auto scaling group ID.
        /// </summary>
        public readonly string AutoScalingGroupId;
        /// <summary>
        /// Auto scaling group name.
        /// </summary>
        public readonly string AutoScalingGroupName;
        /// <summary>
        /// Valid values: `AUTO_CREATION`, `MANUAL_ATTACHING`.
        /// </summary>
        public readonly string CreationType;
        /// <summary>
        /// Health status, the valid values are HEALTHY and UNHEALTHY.
        /// </summary>
        public readonly string HealthStatus;
        /// <summary>
        /// Instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        public readonly string LaunchConfigurationId;
        /// <summary>
        /// Launch configuration name.
        /// </summary>
        public readonly string LaunchConfigurationName;
        /// <summary>
        /// Life cycle state. Please refer to the link for field value details: https://cloud.tencent.com/document/api/377/20453#Instance.
        /// </summary>
        public readonly string LifeCycleState;
        /// <summary>
        /// Enable scale in protection.
        /// </summary>
        public readonly bool ProtectedFromScaleIn;
        /// <summary>
        /// Version ID.
        /// </summary>
        public readonly int VersionNumber;
        /// <summary>
        /// Available zone.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstancesInstanceListResult(
            string addTime,

            string autoScalingGroupId,

            string autoScalingGroupName,

            string creationType,

            string healthStatus,

            string instanceId,

            string instanceType,

            string launchConfigurationId,

            string launchConfigurationName,

            string lifeCycleState,

            bool protectedFromScaleIn,

            int versionNumber,

            string zone)
        {
            AddTime = addTime;
            AutoScalingGroupId = autoScalingGroupId;
            AutoScalingGroupName = autoScalingGroupName;
            CreationType = creationType;
            HealthStatus = healthStatus;
            InstanceId = instanceId;
            InstanceType = instanceType;
            LaunchConfigurationId = launchConfigurationId;
            LaunchConfigurationName = launchConfigurationName;
            LifeCycleState = lifeCycleState;
            ProtectedFromScaleIn = protectedFromScaleIn;
            VersionNumber = versionNumber;
            Zone = zone;
        }
    }
}
