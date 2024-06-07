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

    public sealed class RuleEngineRuleSubRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.RuleEngineRuleSubRuleRuleGetArgs>? _rules;

        /// <summary>
        /// Nested rule settings.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleSubRuleRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RuleEngineRuleSubRuleRuleGetArgs>());
            set => _rules = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tag of the rule.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public RuleEngineRuleSubRuleGetArgs()
        {
        }
        public static new RuleEngineRuleSubRuleGetArgs Empty => new RuleEngineRuleSubRuleGetArgs();
    }
}
