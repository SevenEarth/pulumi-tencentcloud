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

    public sealed class SecurityPolicyConfigWafConfigWafRulesGetArgs : Pulumi.ResourceArgs
    {
        [Input("blockRuleIds", required: true)]
        private InputList<int>? _blockRuleIds;
        public InputList<int> BlockRuleIds
        {
            get => _blockRuleIds ?? (_blockRuleIds = new InputList<int>());
            set => _blockRuleIds = value;
        }

        [Input("observeRuleIds")]
        private InputList<int>? _observeRuleIds;
        public InputList<int> ObserveRuleIds
        {
            get => _observeRuleIds ?? (_observeRuleIds = new InputList<int>());
            set => _observeRuleIds = value;
        }

        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public SecurityPolicyConfigWafConfigWafRulesGetArgs()
        {
        }
    }
}
