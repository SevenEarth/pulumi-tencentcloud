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
    public sealed class GetClusterNodePoolsNodePoolSetGpuArgCudnnResult
    {
        /// <summary>
        /// Dev name of cuDNN.
        /// </summary>
        public readonly string DevName;
        /// <summary>
        /// Doc name of cuDNN.
        /// </summary>
        public readonly string DocName;
        /// <summary>
        /// The attribute name, if there are multiple filters, the relationship between the filters is a logical AND relationship.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// GPU driver or CUDA version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetClusterNodePoolsNodePoolSetGpuArgCudnnResult(
            string devName,

            string docName,

            string name,

            string version)
        {
            DevName = devName;
            DocName = docName;
            Name = name;
            Version = version;
        }
    }
}
