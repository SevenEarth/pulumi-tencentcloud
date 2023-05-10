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

    public sealed class RuleEngineRuleActionCodeActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action name.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("parameters", required: true)]
        private InputList<Inputs.RuleEngineRuleActionCodeActionParameterArgs>? _parameters;

        /// <summary>
        /// Action parameters.
        /// </summary>
        public InputList<Inputs.RuleEngineRuleActionCodeActionParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.RuleEngineRuleActionCodeActionParameterArgs>());
            set => _parameters = value;
        }

        public RuleEngineRuleActionCodeActionArgs()
        {
        }
    }
}
