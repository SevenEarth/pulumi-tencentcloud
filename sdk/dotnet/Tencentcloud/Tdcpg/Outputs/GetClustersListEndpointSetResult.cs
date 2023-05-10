// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdcpg.Outputs
{

    [OutputType]
    public sealed class GetClustersListEndpointSetResult
    {
        /// <summary>
        /// cluster id.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// endpoint id.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// endpoint name.
        /// </summary>
        public readonly string EndpointName;
        /// <summary>
        /// endpoint type.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// private ip.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// private port.
        /// </summary>
        public readonly int PrivatePort;
        /// <summary>
        /// subnet id.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// vpc id.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// wan domain.
        /// </summary>
        public readonly string WanDomain;
        /// <summary>
        /// wan ip.
        /// </summary>
        public readonly string WanIp;
        /// <summary>
        /// wan port.
        /// </summary>
        public readonly int WanPort;

        [OutputConstructor]
        private GetClustersListEndpointSetResult(
            string clusterId,

            string endpointId,

            string endpointName,

            string endpointType,

            string privateIp,

            int privatePort,

            string subnetId,

            string vpcId,

            string wanDomain,

            string wanIp,

            int wanPort)
        {
            ClusterId = clusterId;
            EndpointId = endpointId;
            EndpointName = endpointName;
            EndpointType = endpointType;
            PrivateIp = privateIp;
            PrivatePort = privatePort;
            SubnetId = subnetId;
            VpcId = vpcId;
            WanDomain = wanDomain;
            WanIp = wanIp;
            WanPort = wanPort;
        }
    }
}
