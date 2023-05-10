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

    public sealed class DdosPolicyDdosRuleAclArgs : Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.DdosPolicyDdosRuleAclAclArgs>? _acls;

        /// <summary>
        /// DDoS ACL rule configuration detail.
        /// </summary>
        public InputList<Inputs.DdosPolicyDdosRuleAclAclArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.DdosPolicyDdosRuleAclAclArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// - `on`: Enable. `Acl` parameter is require.- `off`: Disable.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public DdosPolicyDdosRuleAclArgs()
        {
        }
    }
}
