// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetGatewayFlowQosGatewayQosSetResult
    {
        /// <summary>
        /// bandwidth value.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// cvm ip address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// vpc id.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetGatewayFlowQosGatewayQosSetResult(
            int bandwidth,

            string createTime,

            string ipAddress,

            string vpcId)
        {
            Bandwidth = bandwidth;
            CreateTime = createTime;
            IpAddress = ipAddress;
            VpcId = vpcId;
        }
    }
}