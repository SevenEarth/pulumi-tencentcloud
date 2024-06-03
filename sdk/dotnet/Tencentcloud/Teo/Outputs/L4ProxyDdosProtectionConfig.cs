// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class L4ProxyDdosProtectionConfig
    {
        /// <summary>
        /// Exclusive DDoS protection specifications in the Chinese mainland. For details, see [Dedicated DDoS Mitigation Fee (Pay-as-You-Go)] (https://intl.cloud.tencent.com/document/product/1552/94162?from_cn_redirect=1). `PLATFORM`: Default protection of the platform, i.e., Exclusive DDoS protection is not enabled; `BASE30_MAX300`: Exclusive DDoS protection enabled, providing a baseline protection bandwidth of 30 Gbps and an elastic protection bandwidth of up to 300 Gbps; `BASE60_MAX600`: Exclusive DDoS protection enabled, providing a baseline protection bandwidth of 60 Gbps and an elastic protection bandwidth of up to 600 Gbps. If no parameters are filled, the default value PLATFORM is used.
        /// </summary>
        public readonly string? LevelMainland;
        /// <summary>
        /// Exclusive DDoS protection specifications in the worldwide region (excluding the Chinese mainland). `PLATFORM`: Default protection of the platform, i.e., Exclusive DDoS protection is not enabled; `ANYCAST300`: Exclusive DDoS protection enabled, offering a total maximum protection bandwidth of 300 Gbps; `ANYCAST_ALLIN`: Exclusive DDoS protection enabled, utilizing all available protection resources for protection. When no parameters are filled, the default value PLATFORM is used.
        /// </summary>
        public readonly string? LevelOverseas;
        /// <summary>
        /// Configuration of elastic protection bandwidth for exclusive DDoS protection in the Chinese mainland.Valid only when exclusive DDoS protection in the Chinese mainland is enabled (refer to the LevelMainland parameter configuration), and the value has the following limitations: When exclusive DDoS protection is enabled in the Chinese mainland and the 30 Gbps baseline protection bandwidth is used (the LevelMainland parameter value is BASE30_MAX300): the value range is 30 to 300 in Gbps; When exclusive DDoS protection is enabled in the Chinese mainland and the 60 Gbps baseline protection bandwidth is used (the LevelMainland parameter value is BASE60_MAX600): the value range is 60 to 600 in Gbps; When the default protection of the platform is used (the LevelMainland parameter value is PLATFORM): configuration is not supported, and the value of this parameter is invalid.
        /// </summary>
        public readonly int? MaxBandwidthMainland;

        [OutputConstructor]
        private L4ProxyDdosProtectionConfig(
            string? levelMainland,

            string? levelOverseas,

            int? maxBandwidthMainland)
        {
            LevelMainland = levelMainland;
            LevelOverseas = levelOverseas;
            MaxBandwidthMainland = maxBandwidthMainland;
        }
    }
}
