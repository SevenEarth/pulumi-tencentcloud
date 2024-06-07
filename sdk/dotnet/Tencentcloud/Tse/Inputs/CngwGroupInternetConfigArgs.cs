// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class CngwGroupInternetConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description of clb.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// internet type. Reference value:- IPV4 (default value)- IPV6.
        /// </summary>
        [Input("internetAddressVersion")]
        public Input<string>? InternetAddressVersion { get; set; }

        /// <summary>
        /// public network bandwidth.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// trade type of internet. Reference value:- BANDWIDTH- TRAFFIC (default value).
        /// </summary>
        [Input("internetPayMode")]
        public Input<string>? InternetPayMode { get; set; }

        /// <summary>
        /// primary availability zone.
        /// </summary>
        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        /// <summary>
        /// Whether load balancing has multiple availability zones.
        /// </summary>
        [Input("multiZoneFlag")]
        public Input<bool>? MultiZoneFlag { get; set; }

        /// <summary>
        /// specification type of clb. Default shared type when this parameter is empty. Reference value:- SLA LCU-supported.
        /// </summary>
        [Input("slaType")]
        public Input<string>? SlaType { get; set; }

        /// <summary>
        /// alternate availability zone.
        /// </summary>
        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        public CngwGroupInternetConfigArgs()
        {
        }
        public static new CngwGroupInternetConfigArgs Empty => new CngwGroupInternetConfigArgs();
    }
}
