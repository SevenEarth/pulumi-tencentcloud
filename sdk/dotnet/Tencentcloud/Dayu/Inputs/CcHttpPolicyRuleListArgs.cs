// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu.Inputs
{

    public sealed class CcHttpPolicyRuleListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Operator of the rule. Valid values: `include`, `not_include`, `equal`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// Key of the rule. Valid values: `host`, `cgi`, `ua`, `referer`.
        /// </summary>
        [Input("skey")]
        public Input<string>? Skey { get; set; }

        /// <summary>
        /// Rule value, then length should be less than 31 bytes.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public CcHttpPolicyRuleListArgs()
        {
        }
        public static new CcHttpPolicyRuleListArgs Empty => new CcHttpPolicyRuleListArgs();
    }
}
