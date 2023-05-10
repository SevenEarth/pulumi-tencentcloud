// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class DdosPolicyDdosRulePacketFilterArgs : Pulumi.ResourceArgs
    {
        [Input("packetFilters")]
        private InputList<Inputs.DdosPolicyDdosRulePacketFilterPacketFilterArgs>? _packetFilters;

        /// <summary>
        /// DDoS feature filtering configuration detail.
        /// </summary>
        public InputList<Inputs.DdosPolicyDdosRulePacketFilterPacketFilterArgs> PacketFilters
        {
            get => _packetFilters ?? (_packetFilters = new InputList<Inputs.DdosPolicyDdosRulePacketFilterPacketFilterArgs>());
            set => _packetFilters = value;
        }

        /// <summary>
        /// - `on`: Enable. `PacketFilters` parameter is required.- `off`: Disable.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public DdosPolicyDdosRulePacketFilterArgs()
        {
        }
    }
}