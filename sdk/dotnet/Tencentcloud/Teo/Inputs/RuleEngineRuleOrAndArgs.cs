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

    public sealed class RuleEngineRuleOrAndArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to ignore the case of the parameter value, the default value is false.
        /// </summary>
        [Input("ignoreCase")]
        public Input<bool>? IgnoreCase { get; set; }

        /// <summary>
        /// The parameter name corresponding to the matching type is valid when the Target value is the following, and the valid value cannot be empty: `query_string` (query string): The parameter name of the query string in the URL request under the current site, such as lang and version in lang=cn&amp;version=1; `request_header` (HTTP request header): HTTP request header field name, such as Accept-Language in Accept-Language:zh-CN,zh;q=0.9.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Condition operator. Valid values are `equal`, `notequal`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// Condition target. Valid values:- `host`: Host of the URL.- `filename`: filename of the URL.- `extension`: file extension of the URL.- `full_url`: full url.- `url`: path of the URL.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Condition Value.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public RuleEngineRuleOrAndArgs()
        {
        }
    }
}
