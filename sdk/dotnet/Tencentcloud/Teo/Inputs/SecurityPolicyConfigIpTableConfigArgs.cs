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

    public sealed class SecurityPolicyConfigIpTableConfigArgs : Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.SecurityPolicyConfigIpTableConfigRuleArgs>? _rules;
        public InputList<Inputs.SecurityPolicyConfigIpTableConfigRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecurityPolicyConfigIpTableConfigRuleArgs>());
            set => _rules = value;
        }

        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public SecurityPolicyConfigIpTableConfigArgs()
        {
        }
    }
}
