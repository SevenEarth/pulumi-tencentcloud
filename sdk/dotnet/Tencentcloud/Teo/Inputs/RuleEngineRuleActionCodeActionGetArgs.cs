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

    public sealed class RuleEngineRuleActionCodeActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Feature name. You can call the [DescribeRulesSetting](https://tcloud4api.woa.com/document/product/1657/79433?!preview&amp;!document=1) API to view the requirements for entering the feature name.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("parameters", required: true)]
        private InputList<Inputs.RuleEngineRuleActionCodeActionParameterGetArgs>? _parameters;

        /// <summary>
        /// Operation parameter.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleActionCodeActionParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RuleEngineRuleActionCodeActionParameterGetArgs>());
            set => _parameters = value;
        }

        public RuleEngineRuleActionCodeActionGetArgs()
        {
        }
        public static new RuleEngineRuleActionCodeActionGetArgs Empty => new RuleEngineRuleActionCodeActionGetArgs();
    }
}
