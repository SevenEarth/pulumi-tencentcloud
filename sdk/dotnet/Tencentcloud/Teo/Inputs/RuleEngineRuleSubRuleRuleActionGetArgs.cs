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

    public sealed class RuleEngineRuleSubRuleRuleActionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Define a code action.
        /// </summary>
        [Input("codeAction")]
        public Input<Inputs.RuleEngineRuleSubRuleRuleActionCodeActionGetArgs>? CodeAction { get; set; }

        /// <summary>
        /// Define a normal action.
        /// </summary>
        [Input("normalAction")]
        public Input<Inputs.RuleEngineRuleSubRuleRuleActionNormalActionGetArgs>? NormalAction { get; set; }

        /// <summary>
        /// Define a rewrite action.
        /// </summary>
        [Input("rewriteAction")]
        public Input<Inputs.RuleEngineRuleSubRuleRuleActionRewriteActionGetArgs>? RewriteAction { get; set; }

        public RuleEngineRuleSubRuleRuleActionGetArgs()
        {
        }
    }
}
