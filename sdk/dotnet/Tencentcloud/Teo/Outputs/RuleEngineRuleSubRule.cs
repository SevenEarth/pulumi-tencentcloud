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
    public sealed class RuleEngineRuleSubRule
    {
        /// <summary>
        /// Rule items list.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleEngineRuleSubRuleRule> Rules;
        /// <summary>
        /// rule tag list.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private RuleEngineRuleSubRule(
            ImmutableArray<Outputs.RuleEngineRuleSubRuleRule> rules,

            ImmutableArray<string> tags)
        {
            Rules = rules;
            Tags = tags;
        }
    }
}
