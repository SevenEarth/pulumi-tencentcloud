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
    public sealed class LaunchTemplateInternetAccessible
    {
        /// <summary>
        /// The ID of bandwidth package.
        /// </summary>
        public readonly string? BandwidthPackageId;
        /// <summary>
        /// The type of internet charge.
        /// </summary>
        public readonly string? InternetChargeType;
        /// <summary>
        /// Internet outbound bandwidth upper limit, Mbps.
        /// </summary>
        public readonly int? InternetMaxBandwidthOut;
        /// <summary>
        /// Whether to allocate public network IP, TRUE or FALSE.
        /// </summary>
        public readonly bool? PublicIpAssigned;

        [OutputConstructor]
        private LaunchTemplateInternetAccessible(
            string? bandwidthPackageId,

            string? internetChargeType,

            int? internetMaxBandwidthOut,

            bool? publicIpAssigned)
        {
            BandwidthPackageId = bandwidthPackageId;
            InternetChargeType = internetChargeType;
            InternetMaxBandwidthOut = internetMaxBandwidthOut;
            PublicIpAssigned = publicIpAssigned;
        }
    }
}
