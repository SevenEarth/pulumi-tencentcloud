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

    public sealed class SecurityPolicyConfigRateLimitConfigUserRuleConditionArgs : Pulumi.ResourceArgs
    {
        [Input("matchContent", required: true)]
        public Input<string> MatchContent { get; set; } = null!;

        [Input("matchFrom", required: true)]
        public Input<string> MatchFrom { get; set; } = null!;

        [Input("matchParam", required: true)]
        public Input<string> MatchParam { get; set; } = null!;

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public SecurityPolicyConfigRateLimitConfigUserRuleConditionArgs()
        {
        }
    }
}
