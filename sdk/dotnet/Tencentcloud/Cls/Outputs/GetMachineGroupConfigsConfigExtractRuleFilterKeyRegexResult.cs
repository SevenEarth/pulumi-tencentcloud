// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class GetMachineGroupConfigsConfigExtractRuleFilterKeyRegexResult
    {
        /// <summary>
        /// tag key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Filter rule regex corresponding to key.
        /// </summary>
        public readonly string Regex;

        [OutputConstructor]
        private GetMachineGroupConfigsConfigExtractRuleFilterKeyRegexResult(
            string key,

            string regex)
        {
            Key = key;
            Regex = regex;
        }
    }
}
