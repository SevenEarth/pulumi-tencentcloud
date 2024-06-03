// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Inputs
{

    public sealed class Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `IPv6` subnet segment. Such as: `3402:4e00:20:1001::/64`.
        /// </summary>
        [Input("ipv6CidrBlock", required: true)]
        public Input<string> Ipv6CidrBlock { get; set; } = null!;

        /// <summary>
        /// Subnet instance `ID`. Such as:`subnet-pxir56ns`.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksGetArgs()
        {
        }
        public static new Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksGetArgs Empty => new Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksGetArgs();
    }
}
