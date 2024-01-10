// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Gaap.Outputs
{

    [OutputType]
    public sealed class GetRuleRealServersRealServerSetResult
    {
        /// <summary>
        /// Is it on the banned blacklist? 0 indicates not on the blacklist, and 1 indicates on the blacklist.
        /// </summary>
        public readonly int InBanBlacklist;
        /// <summary>
        /// Project Id.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Real Server Id.
        /// </summary>
        public readonly string RealServerId;
        /// <summary>
        /// Real Server IP or domain.
        /// </summary>
        public readonly string RealServerIp;
        /// <summary>
        /// Real Server Name.
        /// </summary>
        public readonly string RealServerName;

        [OutputConstructor]
        private GetRuleRealServersRealServerSetResult(
            int inBanBlacklist,

            int projectId,

            string realServerId,

            string realServerIp,

            string realServerName)
        {
            InBanBlacklist = inBanBlacklist;
            ProjectId = projectId;
            RealServerId = realServerId;
            RealServerIp = realServerIp;
            RealServerName = realServerName;
        }
    }
}
