// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class CosRechargeExtractRuleInfoFilterKeyRegexGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// need filter log key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// need filter log regex.
        /// </summary>
        [Input("regex", required: true)]
        public Input<string> Regex { get; set; } = null!;

        public CosRechargeExtractRuleInfoFilterKeyRegexGetArgs()
        {
        }
        public static new CosRechargeExtractRuleInfoFilterKeyRegexGetArgs Empty => new CosRechargeExtractRuleInfoFilterKeyRegexGetArgs();
    }
}
