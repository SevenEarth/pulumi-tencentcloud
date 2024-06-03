// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class NodePoolNodeConfig
    {
        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodePoolNodeConfigDataDisk> DataDisks;
        /// <summary>
        /// Indicate to set desired pod number in node. valid when the cluster is podCIDR.
        /// </summary>
        public readonly int? DesiredPodNum;
        /// <summary>
        /// Docker graph path. Default is `/var/lib/docker`.
        /// </summary>
        public readonly string? DockerGraphPath;
        /// <summary>
        /// Custom parameter information related to the node. This is a white-list parameter.
        /// </summary>
        public readonly ImmutableArray<string> ExtraArgs;
        /// <summary>
        /// GPU driver parameters.
        /// </summary>
        public readonly Outputs.NodePoolNodeConfigGpuArgs? GpuArgs;
        /// <summary>
        /// Indicate to schedule the adding node or not. Default is true.
        /// </summary>
        public readonly bool? IsSchedule;
        /// <summary>
        /// Mount target. Default is not mounting.
        /// </summary>
        public readonly string? MountTarget;
        /// <summary>
        /// Base64-encoded user script, executed before initializing the node, currently only effective for adding existing nodes.
        /// </summary>
        public readonly string? PreStartUserScript;
        /// <summary>
        /// Base64-encoded User Data text, the length limit is 16KB.
        /// </summary>
        public readonly string? UserData;

        [OutputConstructor]
        private NodePoolNodeConfig(
            ImmutableArray<Outputs.NodePoolNodeConfigDataDisk> dataDisks,

            int? desiredPodNum,

            string? dockerGraphPath,

            ImmutableArray<string> extraArgs,

            Outputs.NodePoolNodeConfigGpuArgs? gpuArgs,

            bool? isSchedule,

            string? mountTarget,

            string? preStartUserScript,

            string? userData)
        {
            DataDisks = dataDisks;
            DesiredPodNum = desiredPodNum;
            DockerGraphPath = dockerGraphPath;
            ExtraArgs = extraArgs;
            GpuArgs = gpuArgs;
            IsSchedule = isSchedule;
            MountTarget = mountTarget;
            PreStartUserScript = preStartUserScript;
            UserData = userData;
        }
    }
}
