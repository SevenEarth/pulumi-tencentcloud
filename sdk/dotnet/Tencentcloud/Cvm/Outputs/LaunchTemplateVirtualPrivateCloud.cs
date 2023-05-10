// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateVirtualPrivateCloud
    {
        /// <summary>
        /// Is it used as a Public network gateway, TRUE or FALSE.
        /// </summary>
        public readonly bool? AsVpcGateway;
        /// <summary>
        /// The number of ipv6 addresses for Elastic Network Interface.
        /// </summary>
        public readonly int? Ipv6AddressCount;
        /// <summary>
        /// The address of private ip.
        /// </summary>
        public readonly ImmutableArray<string> PrivateIpAddresses;
        /// <summary>
        /// The id of subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The id of VPC.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private LaunchTemplateVirtualPrivateCloud(
            bool? asVpcGateway,

            int? ipv6AddressCount,

            ImmutableArray<string> privateIpAddresses,

            string subnetId,

            string vpcId)
        {
            AsVpcGateway = asVpcGateway;
            Ipv6AddressCount = ipv6AddressCount;
            PrivateIpAddresses = privateIpAddresses;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}