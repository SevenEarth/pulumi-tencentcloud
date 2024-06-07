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

    public sealed class RuleEngineRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.RuleEngineRuleActionGetArgs>? _actions;

        /// <summary>
        /// Feature to be executed.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.RuleEngineRuleActionGetArgs>());
            set => _actions = value;
        }

        [Input("ors", required: true)]
        private InputList<Inputs.RuleEngineRuleOrGetArgs>? _ors;

        /// <summary>
        /// OR Conditions list of the rule. Rule would be triggered if any of the condition is true.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleOrGetArgs> Ors
        {
            get => _ors ?? (_ors = new InputList<Inputs.RuleEngineRuleOrGetArgs>());
            set => _ors = value;
        }

        [Input("subRules")]
        private InputList<Inputs.RuleEngineRuleSubRuleGetArgs>? _subRules;

        /// <summary>
        /// The nested rule.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleSubRuleGetArgs> SubRules
        {
            get => _subRules ?? (_subRules = new InputList<Inputs.RuleEngineRuleSubRuleGetArgs>());
            set => _subRules = value;
        }

        public RuleEngineRuleGetArgs()
        {
        }
        public static new RuleEngineRuleGetArgs Empty => new RuleEngineRuleGetArgs();
    }
}
