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
    public sealed class LaunchTemplateVersionVirtualPrivateCloud
    {
        /// <summary>
        /// Whether to use a CVM instance as a public gateway. The public gateway is only available when the instance has a public IP and resides in a VPC.
        /// </summary>
        public readonly bool? AsVpcGateway;
        /// <summary>
        /// Number of IPv6 addresses randomly generated for the ENI.
        /// </summary>
        public readonly int? Ipv6AddressCount;
        /// <summary>
        /// Array of VPC subnet IPs. You can use this parameter when creating instances or modifying VPC attributes of instances. Currently you can specify multiple IPs in one subnet only when creating multiple instances at the same time.
        /// </summary>
        public readonly ImmutableArray<string> PrivateIpAddresses;
        /// <summary>
        /// VPC subnet ID in the format subnet-xxx, if you specify DEFAULT for both VpcId and SubnetId when creating an instance, the default VPC will be used.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// VPC ID in the format of vpc-xxx, if you specify DEFAULT for both VpcId and SubnetId when creating an instance, the default VPC will be used.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private LaunchTemplateVersionVirtualPrivateCloud(
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
