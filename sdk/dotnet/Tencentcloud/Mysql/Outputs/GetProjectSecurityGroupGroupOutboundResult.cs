// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class GetProjectSecurityGroupGroupOutboundResult
    {
        /// <summary>
        /// Policy, ACCEPT or DROP.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Destination IP or IP segment, such as 172.16.0.0/12.
        /// </summary>
        public readonly string CidrIp;
        /// <summary>
        /// Rule description.
        /// </summary>
        public readonly string Desc;
        /// <summary>
        /// The direction defined by the rule, the inbound rule is OUTPUT.
        /// </summary>
        public readonly string Dir;
        /// <summary>
        /// Network protocol, support UDP, TCP, etc.
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// port or port range.
        /// </summary>
        public readonly string PortRange;

        [OutputConstructor]
        private GetProjectSecurityGroupGroupOutboundResult(
            string action,

            string cidrIp,

            string desc,

            string dir,

            string ipProtocol,

            string portRange)
        {
            Action = action;
            CidrIp = cidrIp;
            Desc = desc;
            Dir = dir;
            IpProtocol = ipProtocol;
            PortRange = portRange;
        }
    }
}
