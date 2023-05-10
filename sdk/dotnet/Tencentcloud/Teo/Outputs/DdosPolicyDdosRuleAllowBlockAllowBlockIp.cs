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
    public sealed class DdosPolicyDdosRuleAllowBlockAllowBlockIp
    {
        /// <summary>
        /// Valid value format:- ip, for example 1.1.1.1- ip range, for example 1.1.1.2-1.1.1.3- network segment, for example 1.2.1.0/24- network segment range, for example 1.2.1.0/24-1.2.2.0/24.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Valid values: `block`, `allow`.
        /// </summary>
        public readonly string Type;
        public readonly int? UpdateTime;

        [OutputConstructor]
        private DdosPolicyDdosRuleAllowBlockAllowBlockIp(
            string? ip,

            string type,

            int? updateTime)
        {
            Ip = ip;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}