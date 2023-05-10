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

    public sealed class DdosPolicyDdosRuleAllowBlockAllowBlockIpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid value format:- ip, for example 1.1.1.1- ip range, for example 1.1.1.2-1.1.1.3- network segment, for example 1.2.1.0/24- network segment range, for example 1.2.1.0/24-1.2.2.0/24.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Valid values: `block`, `allow`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("updateTime")]
        public Input<int>? UpdateTime { get; set; }

        public DdosPolicyDdosRuleAllowBlockAllowBlockIpArgs()
        {
        }
    }
}
