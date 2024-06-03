// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainIpFilterFilterRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Ip filter `blacklist`/`whitelist` type of filter rules.
        /// </summary>
        [Input("filterType", required: true)]
        public Input<string> FilterType { get; set; } = null!;

        [Input("filters", required: true)]
        private InputList<string>? _filters;

        /// <summary>
        /// Ip filter rule list, supports IPs in X.X.X.X format, or /8, /16, /24 format IP ranges. Up to 50 allowlists or blocklists can be entered.
        /// </summary>
        public InputList<string> Filters
        {
            get => _filters ?? (_filters = new InputList<string>());
            set => _filters = value;
        }

        [Input("rulePaths", required: true)]
        private InputList<string>? _rulePaths;

        /// <summary>
        /// Content list for each `rule_type`: `*` for `all`, file ext like `jpg` for `file`, `/dir/like/` for `directory` and `/path/index.html` for `path`.
        /// </summary>
        public InputList<string> RulePaths
        {
            get => _rulePaths ?? (_rulePaths = new InputList<string>());
            set => _rulePaths = value;
        }

        /// <summary>
        /// Ip filter rule type of filter rules, available: `all`, `file`, `directory`, `path`.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        public DomainIpFilterFilterRuleArgs()
        {
        }
        public static new DomainIpFilterFilterRuleArgs Empty => new DomainIpFilterFilterRuleArgs();
    }
}
