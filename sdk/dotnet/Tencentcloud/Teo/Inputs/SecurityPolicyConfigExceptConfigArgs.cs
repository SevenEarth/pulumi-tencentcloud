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

    public sealed class SecurityPolicyConfigExceptConfigArgs : Pulumi.ResourceArgs
    {
        [Input("exceptUserRules")]
        private InputList<Inputs.SecurityPolicyConfigExceptConfigExceptUserRuleArgs>? _exceptUserRules;

        /// <summary>
        /// Exception rules.
        /// </summary>
        public InputList<Inputs.SecurityPolicyConfigExceptConfigExceptUserRuleArgs> ExceptUserRules
        {
            get => _exceptUserRules ?? (_exceptUserRules = new InputList<Inputs.SecurityPolicyConfigExceptConfigExceptUserRuleArgs>());
            set => _exceptUserRules = value;
        }

        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public SecurityPolicyConfigExceptConfigArgs()
        {
        }
    }
}
