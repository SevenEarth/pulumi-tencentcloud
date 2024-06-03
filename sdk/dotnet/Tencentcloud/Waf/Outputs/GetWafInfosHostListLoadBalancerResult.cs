// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf.Outputs
{

    [OutputType]
    public sealed class GetWafInfosHostListLoadBalancerResult
    {
        /// <summary>
        /// Unique ID of listener in LB.
        /// </summary>
        public readonly string ListenerId;
        /// <summary>
        /// Listener name.
        /// </summary>
        public readonly string ListenerName;
        /// <summary>
        /// LoadBalancer ID.
        /// </summary>
        public readonly string LoadBalancerId;
        /// <summary>
        /// LoadBalancer name.
        /// </summary>
        public readonly string LoadBalancerName;
        /// <summary>
        /// Network type for load balancerNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string LoadBalancerType;
        /// <summary>
        /// VPCID for load balancer, public network is -1, and internal network is filled in according to actual conditionsNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int NumericalVpcId;
        /// <summary>
        /// Protocol of listener，http or https.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// LoadBalancer region.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// LoadBalancer IP.
        /// </summary>
        public readonly string Vip;
        /// <summary>
        /// LoadBalancer port.
        /// </summary>
        public readonly int Vport;
        /// <summary>
        /// LoadBalancer zone.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetWafInfosHostListLoadBalancerResult(
            string listenerId,

            string listenerName,

            string loadBalancerId,

            string loadBalancerName,

            string loadBalancerType,

            int numericalVpcId,

            string protocol,

            string region,

            string vip,

            int vport,

            string zone)
        {
            ListenerId = listenerId;
            ListenerName = listenerName;
            LoadBalancerId = loadBalancerId;
            LoadBalancerName = loadBalancerName;
            LoadBalancerType = loadBalancerType;
            NumericalVpcId = numericalVpcId;
            Protocol = protocol;
            Region = region;
            Vip = vip;
            Vport = vport;
            Zone = zone;
        }
    }
}
