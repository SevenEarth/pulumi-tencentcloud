// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb.Outputs
{

    [OutputType]
    public sealed class GetSecurityGroupsListOutboundResult
    {
        /// <summary>
        /// policy action.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// cidr ip.
        /// </summary>
        public readonly string CidrIp;
        /// <summary>
        /// internet protocol.
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// port range.
        /// </summary>
        public readonly string PortRange;

        [OutputConstructor]
        private GetSecurityGroupsListOutboundResult(
            string action,

            string cidrIp,

            string ipProtocol,

            string portRange)
        {
            Action = action;
            CidrIp = cidrIp;
            IpProtocol = ipProtocol;
            PortRange = portRange;
        }
    }
}
