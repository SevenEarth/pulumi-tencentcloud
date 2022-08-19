// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainIpFilterArgs : Pulumi.ResourceArgs
    {
        [Input("filterRules")]
        private InputList<Inputs.DomainIpFilterFilterRuleArgs>? _filterRules;

        /// <summary>
        /// Ip filter rules, This feature is only available to selected beta customers.
        /// </summary>
        public InputList<Inputs.DomainIpFilterFilterRuleArgs> FilterRules
        {
            get => _filterRules ?? (_filterRules = new InputList<Inputs.DomainIpFilterFilterRuleArgs>());
            set => _filterRules = value;
        }

        /// <summary>
        /// IP `blacklist`/`whitelist` type.
        /// </summary>
        [Input("filterType")]
        public Input<string>? FilterType { get; set; }

        [Input("filters")]
        private InputList<string>? _filters;

        /// <summary>
        /// Ip filter list, Supports IPs in X.X.X.X format, or /8, /16, /24 format IP ranges. Up to 50 allowlists or blocklists can be entered.
        /// </summary>
        public InputList<string> Filters
        {
            get => _filters ?? (_filters = new InputList<string>());
            set => _filters = value;
        }

        /// <summary>
        /// Return code, available values: 400-499.
        /// </summary>
        [Input("returnCode")]
        public Input<int>? ReturnCode { get; set; }

        /// <summary>
        /// Configuration switch, available values: `on`, `off` (default).
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public DomainIpFilterArgs()
        {
        }
    }
}
