// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class RuleEngineRuleSubRuleRuleActionCodeActionParameter
    {
        /// <summary>
        /// Parameter name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status code.
        /// </summary>
        public readonly int StatusCode;
        /// <summary>
        /// Parameter value.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private RuleEngineRuleSubRuleRuleActionCodeActionParameter(
            string name,

            int statusCode,

            ImmutableArray<string> values)
        {
            Name = name;
            StatusCode = statusCode;
            Values = values;
        }
    }
}
