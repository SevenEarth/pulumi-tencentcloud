// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq.Outputs
{

    [OutputType]
    public sealed class GetRocketmqClusterClusterListResult
    {
        /// <summary>
        /// Cluster configuration information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqClusterClusterListConfigResult> Configs;
        /// <summary>
        /// Basic cluster information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqClusterClusterListInfoResult> Infos;
        /// <summary>
        /// Cluster status. `0`: Creating; `1`: Normal; `2`: Terminating; `3`: Deleted; `4`: Isolated; `5`: Creation failed; `6`: Deletion failed.
        /// </summary>
        public readonly int Status;

        [OutputConstructor]
        private GetRocketmqClusterClusterListResult(
            ImmutableArray<Outputs.GetRocketmqClusterClusterListConfigResult> configs,

            ImmutableArray<Outputs.GetRocketmqClusterClusterListInfoResult> infos,

            int status)
        {
            Configs = configs;
            Infos = infos;
            Status = status;
        }
    }
}