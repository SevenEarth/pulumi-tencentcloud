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

    public sealed class DomainRefererGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("refererRules")]
        private InputList<Inputs.DomainRefererRefererRuleGetArgs>? _refererRules;

        /// <summary>
        /// List of referer rules.
        /// </summary>
        public InputList<Inputs.DomainRefererRefererRuleGetArgs> RefererRules
        {
            get => _refererRules ?? (_refererRules = new InputList<Inputs.DomainRefererRefererRuleGetArgs>());
            set => _refererRules = value;
        }

        /// <summary>
        /// Configuration switch, available values: `on`, `off` (default).
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public DomainRefererGetArgs()
        {
        }
        public static new DomainRefererGetArgs Empty => new DomainRefererGetArgs();
    }
}
