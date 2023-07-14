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
    public sealed class GetPodInstancesResultContentResult
    {
        /// <summary>
        /// Instance start time.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Instance available status.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string InstanceAvailableStatus;
        /// <summary>
        /// Instance status.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string InstanceStatus;
        /// <summary>
        /// Instance ip.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Instance node id.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string NodeInstanceId;
        /// <summary>
        /// Instance node ip.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string NodeIp;
        /// <summary>
        /// Instance id (corresponding to the pod instance id in Kubernetes).Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string PodId;
        /// <summary>
        /// Instance name (corresponding to the pod name in Kubernetes).Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string PodName;
        /// <summary>
        /// Instance ready count.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly int ReadyCount;
        /// <summary>
        /// Instance reason for current status.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// Instance restart count.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly int RestartCount;
        /// <summary>
        /// Instance run time.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// Instance serve status.Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string ServiceInstanceStatus;
        /// <summary>
        /// Instance status. Please refer to the definition of instance and container status below. Starting (pod not ready): Starting; Running: Running; Abnormal: Abnormal; Stopped: Stopped;Note: This field may return null, which means no valid value was found.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetPodInstancesResultContentResult(
            string createdAt,

            string instanceAvailableStatus,

            string instanceStatus,

            string ip,

            string nodeInstanceId,

            string nodeIp,

            string podId,

            string podName,

            int readyCount,

            string reason,

            int restartCount,

            string runtime,

            string serviceInstanceStatus,

            string status)
        {
            CreatedAt = createdAt;
            InstanceAvailableStatus = instanceAvailableStatus;
            InstanceStatus = instanceStatus;
            Ip = ip;
            NodeInstanceId = nodeInstanceId;
            NodeIp = nodeIp;
            PodId = podId;
            PodName = podName;
            ReadyCount = readyCount;
            Reason = reason;
            RestartCount = restartCount;
            Runtime = runtime;
            ServiceInstanceStatus = serviceInstanceStatus;
            Status = status;
        }
    }
}
