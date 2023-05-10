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
    public sealed class RuleEngineRuleSubRuleRuleOrAnd
    {
        /// <summary>
        /// Whether to ignore the case of the parameter value, the default value is false.
        /// </summary>
        public readonly bool? IgnoreCase;
        /// <summary>
        /// The parameter name corresponding to the matching type is valid when the Target value is the following, and the valid value cannot be empty: `query_string` (query string): The parameter name of the query string in the URL request under the current site, such as lang and version in lang=cn&amp;version=1; `request_header` (HTTP request header): HTTP request header field name, such as Accept-Language in Accept-Language:zh-CN,zh;q=0.9.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Condition operator. Valid values are `equal`, `notequal`.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Condition target. Valid values:- `host`: Host of the URL.- `filename`: filename of the URL.- `extension`: file extension of the URL.- `full_url`: full url.- `url`: path of the URL.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// Condition Value.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private RuleEngineRuleSubRuleRuleOrAnd(
            bool? ignoreCase,

            string? name,

            string @operator,

            string target,

            ImmutableArray<string> values)
        {
            IgnoreCase = ignoreCase;
            Name = name;
            Operator = @operator;
            Target = target;
            Values = values;
        }
    }
}
