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

    public sealed class SecurityPolicyConfigBotConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("intelligenceRule")]
        public Input<Inputs.SecurityPolicyConfigBotConfigIntelligenceRuleGetArgs>? IntelligenceRule { get; set; }

        [Input("managedRule")]
        public Input<Inputs.SecurityPolicyConfigBotConfigManagedRuleGetArgs>? ManagedRule { get; set; }

        [Input("portraitRule")]
        public Input<Inputs.SecurityPolicyConfigBotConfigPortraitRuleGetArgs>? PortraitRule { get; set; }

        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public SecurityPolicyConfigBotConfigGetArgs()
        {
        }
    }
}
