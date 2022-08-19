// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dnats.Outputs
{

    [OutputType]
    public sealed class GetInstanceDnatListResult
    {
        /// <summary>
        /// Description of the NAT forward.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Network address of the EIP.
        /// </summary>
        public readonly string ElasticIp;
        /// <summary>
        /// Port of the EIP.
        /// </summary>
        public readonly string ElasticPort;
        /// <summary>
        /// ID of the NAT gateway.
        /// </summary>
        public readonly string NatId;
        /// <summary>
        /// Network address of the backend service.
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// Port of intranet.
        /// </summary>
        public readonly string PrivatePort;
        /// <summary>
        /// Type of the network protocol. Valid values: `TCP` and `UDP`.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// ID of the VPC.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetInstanceDnatListResult(
            string? description,

            string elasticIp,

            string elasticPort,

            string natId,

            string privateIp,

            string privatePort,

            string protocol,

            string vpcId)
        {
            Description = description;
            ElasticIp = elasticIp;
            ElasticPort = elasticPort;
            NatId = natId;
            PrivateIp = privateIp;
            PrivatePort = privatePort;
            Protocol = protocol;
            VpcId = vpcId;
        }
    }
}
