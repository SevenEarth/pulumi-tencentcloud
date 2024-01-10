// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class GetDescribeHostClbInstanceListInstanceListResult
    {
        /// <summary>
        /// CLB listener listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeHostClbInstanceListInstanceListListenerResult> Listeners;
        /// <summary>
        /// CLB instance ID.
        /// </summary>
        public readonly string LoadBalancerId;
        /// <summary>
        /// CLB instance name name.
        /// </summary>
        public readonly string LoadBalancerName;

        [OutputConstructor]
        private GetDescribeHostClbInstanceListInstanceListResult(
            ImmutableArray<Outputs.GetDescribeHostClbInstanceListInstanceListListenerResult> listeners,

            string loadBalancerId,

            string loadBalancerName)
        {
            Listeners = listeners;
            LoadBalancerId = loadBalancerId;
            LoadBalancerName = loadBalancerName;
        }
    }
}
