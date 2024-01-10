// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class GetListListenerLayer7ListenerInstanceDetailRuleResult
    {
        /// <summary>
        /// Cname.
        /// </summary>
        public readonly string Cname;
        /// <summary>
        /// Instance ip list.
        /// </summary>
        public readonly ImmutableArray<string> EipLists;
        /// <summary>
        /// Instance id.
        /// </summary>
        public readonly string InstanceId;

        [OutputConstructor]
        private GetListListenerLayer7ListenerInstanceDetailRuleResult(
            string cname,

            ImmutableArray<string> eipLists,

            string instanceId)
        {
            Cname = cname;
            EipLists = eipLists;
            InstanceId = instanceId;
        }
    }
}
