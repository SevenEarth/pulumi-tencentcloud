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
    public sealed class GetClusterNativeNodePoolsNodePoolNativeManagementResult
    {
        /// <summary>
        /// Hosts configuration.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// Kernel parameter configuration.
        /// </summary>
        public readonly ImmutableArray<string> KernelArgs;
        /// <summary>
        /// Dns configuration.
        /// </summary>
        public readonly ImmutableArray<string> Nameservers;

        [OutputConstructor]
        private GetClusterNativeNodePoolsNodePoolNativeManagementResult(
            ImmutableArray<string> hosts,

            ImmutableArray<string> kernelArgs,

            ImmutableArray<string> nameservers)
        {
            Hosts = hosts;
            KernelArgs = kernelArgs;
            Nameservers = nameservers;
        }
    }
}
