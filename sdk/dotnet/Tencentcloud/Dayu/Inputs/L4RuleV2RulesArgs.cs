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

    public sealed class L4RuleV2RulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// session hold switch.
        /// </summary>
        [Input("keepEnable", required: true)]
        public Input<bool> KeepEnable { get; set; } = null!;

        /// <summary>
        /// The keeptime of the layer 4 rule.
        /// </summary>
        [Input("keeptime", required: true)]
        public Input<int> Keeptime { get; set; } = null!;

        /// <summary>
        /// LB type of the rule, `1` for weight cycling and `2` for IP hash.
        /// </summary>
        [Input("lbType", required: true)]
        public Input<int> LbType { get; set; } = null!;

        /// <summary>
        /// Protocol of the rule.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Corresponding regional information.
        /// </summary>
        [Input("region", required: true)]
        public Input<int> Region { get; set; } = null!;

        /// <summary>
        /// Remove the watermark state.
        /// </summary>
        [Input("removeSwitch", required: true)]
        public Input<bool> RemoveSwitch { get; set; } = null!;

        /// <summary>
        /// Name of the rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        [Input("sourceLists", required: true)]
        private InputList<Inputs.L4RuleV2RulesSourceListArgs>? _sourceLists;

        /// <summary>
        /// Source list of the rule.
        /// </summary>
        public InputList<Inputs.L4RuleV2RulesSourceListArgs> SourceLists
        {
            get => _sourceLists ?? (_sourceLists = new InputList<Inputs.L4RuleV2RulesSourceListArgs>());
            set => _sourceLists = value;
        }

        /// <summary>
        /// The source port of the layer 4 rule.
        /// </summary>
        [Input("sourcePort", required: true)]
        public Input<int> SourcePort { get; set; } = null!;

        /// <summary>
        /// Source type, `1` for source of host, `2` for source of IP.
        /// </summary>
        [Input("sourceType", required: true)]
        public Input<int> SourceType { get; set; } = null!;

        /// <summary>
        /// The virtual port of the layer 4 rule.
        /// </summary>
        [Input("virtualPort", required: true)]
        public Input<int> VirtualPort { get; set; } = null!;

        public L4RuleV2RulesArgs()
        {
        }
        public static new L4RuleV2RulesArgs Empty => new L4RuleV2RulesArgs();
    }
}
