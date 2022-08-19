// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainRequestHeaderGetArgs : Pulumi.ResourceArgs
    {
        [Input("headerRules")]
        private InputList<Inputs.DomainRequestHeaderHeaderRuleGetArgs>? _headerRules;

        /// <summary>
        /// Custom request header configuration rules.
        /// </summary>
        public InputList<Inputs.DomainRequestHeaderHeaderRuleGetArgs> HeaderRules
        {
            get => _headerRules ?? (_headerRules = new InputList<Inputs.DomainRequestHeaderHeaderRuleGetArgs>());
            set => _headerRules = value;
        }

        /// <summary>
        /// Custom request header configuration switch. Valid values are `on` and `off`. and default value is `off`.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public DomainRequestHeaderGetArgs()
        {
        }
    }
}
