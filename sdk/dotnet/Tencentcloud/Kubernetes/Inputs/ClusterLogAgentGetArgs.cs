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

    public sealed class ClusterLogAgentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the log agent enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Kubelet root directory as the literal.
        /// </summary>
        [Input("kubeletRootDir")]
        public Input<string>? KubeletRootDir { get; set; }

        public ClusterLogAgentGetArgs()
        {
        }
        public static new ClusterLogAgentGetArgs Empty => new ClusterLogAgentGetArgs();
    }
}
