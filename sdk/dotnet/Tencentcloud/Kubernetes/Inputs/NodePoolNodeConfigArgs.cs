// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Inputs
{

    public sealed class NodePoolNodeConfigArgs : Pulumi.ResourceArgs
    {
        [Input("dataDisks")]
        private InputList<Inputs.NodePoolNodeConfigDataDiskArgs>? _dataDisks;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public InputList<Inputs.NodePoolNodeConfigDataDiskArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.NodePoolNodeConfigDataDiskArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// Indicate to set desired pod number in node. valid when the cluster is podCIDR.
        /// </summary>
        [Input("desiredPodNum")]
        public Input<int>? DesiredPodNum { get; set; }

        /// <summary>
        /// Docker graph path. Default is `/var/lib/docker`.
        /// </summary>
        [Input("dockerGraphPath")]
        public Input<string>? DockerGraphPath { get; set; }

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;

        /// <summary>
        /// Custom parameter information related to the node. This is a white-list parameter.
        /// </summary>
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        /// <summary>
        /// GPU driver parameters.
        /// </summary>
        [Input("gpuArgs")]
        public Input<Inputs.NodePoolNodeConfigGpuArgsArgs>? GpuArgs { get; set; }

        /// <summary>
        /// Indicate to schedule the adding node or not. Default is true.
        /// </summary>
        [Input("isSchedule")]
        public Input<bool>? IsSchedule { get; set; }

        /// <summary>
        /// Mount target. Default is not mounting.
        /// </summary>
        [Input("mountTarget")]
        public Input<string>? MountTarget { get; set; }

        /// <summary>
        /// Base64-encoded User Data text, the length limit is 16KB.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public NodePoolNodeConfigArgs()
        {
        }
    }
}
