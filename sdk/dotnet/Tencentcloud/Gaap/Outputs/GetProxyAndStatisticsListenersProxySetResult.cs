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
    public sealed class GetProxyAndStatisticsListenersProxySetResult
    {
        /// <summary>
        /// Listener List.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyAndStatisticsListenersProxySetListenerListResult> ListenerLists;
        /// <summary>
        /// Proxy Id.
        /// </summary>
        public readonly string ProxyId;
        /// <summary>
        /// Proxy Name.
        /// </summary>
        public readonly string ProxyName;

        [OutputConstructor]
        private GetProxyAndStatisticsListenersProxySetResult(
            ImmutableArray<Outputs.GetProxyAndStatisticsListenersProxySetListenerListResult> listenerLists,

            string proxyId,

            string proxyName)
        {
            ListenerLists = listenerLists;
            ProxyId = proxyId;
            ProxyName = proxyName;
        }
    }
}
