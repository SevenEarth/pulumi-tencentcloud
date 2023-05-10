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

    public sealed class SecurityPolicyConfigBotConfigIntelligenceRuleItemGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take. Valid values: `trans`, `monitor`, `alg`, `captcha`, `drop`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Bot label, valid values: `evil_bot`, `suspect_bot`, `good_bot`, `normal`.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        public SecurityPolicyConfigBotConfigIntelligenceRuleItemGetArgs()
        {
        }
    }
}
