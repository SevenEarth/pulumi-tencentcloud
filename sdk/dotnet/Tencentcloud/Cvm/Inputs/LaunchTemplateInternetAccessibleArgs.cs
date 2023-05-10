// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateInternetAccessibleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of bandwidth package.
        /// </summary>
        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        /// <summary>
        /// The type of internet charge.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// Internet outbound bandwidth upper limit, Mbps.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// Whether to allocate public network IP, TRUE or FALSE.
        /// </summary>
        [Input("publicIpAssigned")]
        public Input<bool>? PublicIpAssigned { get; set; }

        public LaunchTemplateInternetAccessibleArgs()
        {
        }
    }
}
