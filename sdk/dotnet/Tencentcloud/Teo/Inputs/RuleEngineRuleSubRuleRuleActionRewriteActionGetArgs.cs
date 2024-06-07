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

    public sealed class RuleEngineRuleSubRuleRuleActionRewriteActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Feature name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&amp;!document=1) API to view the requirements for entering the feature name.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("parameters", required: true)]
        private InputList<Inputs.RuleEngineRuleSubRuleRuleActionRewriteActionParameterGetArgs>? _parameters;

        /// <summary>
        /// Parameter.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleSubRuleRuleActionRewriteActionParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RuleEngineRuleSubRuleRuleActionRewriteActionParameterGetArgs>());
            set => _parameters = value;
        }

        public RuleEngineRuleSubRuleRuleActionRewriteActionGetArgs()
        {
        }
        public static new RuleEngineRuleSubRuleRuleActionRewriteActionGetArgs Empty => new RuleEngineRuleSubRuleRuleActionRewriteActionGetArgs();
    }
}
