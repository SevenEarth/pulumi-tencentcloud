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

    public sealed class SecurityPolicyConfigWafConfigWafRulesArgs : Pulumi.ResourceArgs
    {
        [Input("blockRuleIds", required: true)]
        private InputList<int>? _blockRuleIds;

        /// <summary>
        /// Block mode rules list. See details in data source `waf_managed_rules`.
        /// </summary>
        public InputList<int> BlockRuleIds
        {
            get => _blockRuleIds ?? (_blockRuleIds = new InputList<int>());
            set => _blockRuleIds = value;
        }

        [Input("observeRuleIds")]
        private InputList<int>? _observeRuleIds;

        /// <summary>
        /// Observe rules list. See details in data source `waf_managed_rules`.
        /// </summary>
        public InputList<int> ObserveRuleIds
        {
            get => _observeRuleIds ?? (_observeRuleIds = new InputList<int>());
            set => _observeRuleIds = value;
        }

        /// <summary>
        /// Whether to host the rules&amp;#39; configuration.- `on`: Enable.- `off`: Disable.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public SecurityPolicyConfigWafConfigWafRulesArgs()
        {
        }
    }
}
