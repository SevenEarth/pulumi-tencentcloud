// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Outputs
{

    [OutputType]
    public sealed class CngwGatewayInternetConfig
    {
        /// <summary>
        /// description of clb.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// internet type. Reference value: `IPV4`(default value), `IPV6`.
        /// </summary>
        public readonly string? InternetAddressVersion;
        /// <summary>
        /// public network bandwidth.
        /// </summary>
        public readonly int? InternetMaxBandwidthOut;
        /// <summary>
        /// trade type of internet. Reference value: `BANDWIDTH`, `TRAFFIC`(default value).
        /// </summary>
        public readonly string? InternetPayMode;
        /// <summary>
        /// primary availability zone.
        /// </summary>
        public readonly string? MasterZoneId;
        /// <summary>
        /// Whether load balancing has multiple availability zones.
        /// </summary>
        public readonly bool? MultiZoneFlag;
        /// <summary>
        /// specification type of clb. Default shared type when this parameter is empty. Reference value:- SLA LCU-supported.
        /// </summary>
        public readonly string? SlaType;
        /// <summary>
        /// alternate availability zone.
        /// </summary>
        public readonly string? SlaveZoneId;

        [OutputConstructor]
        private CngwGatewayInternetConfig(
            string? description,

            string? internetAddressVersion,

            int? internetMaxBandwidthOut,

            string? internetPayMode,

            string? masterZoneId,

            bool? multiZoneFlag,

            string? slaType,

            string? slaveZoneId)
        {
            Description = description;
            InternetAddressVersion = internetAddressVersion;
            InternetMaxBandwidthOut = internetMaxBandwidthOut;
            InternetPayMode = internetPayMode;
            MasterZoneId = masterZoneId;
            MultiZoneFlag = multiZoneFlag;
            SlaType = slaType;
            SlaveZoneId = slaveZoneId;
        }
    }
}
