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
    public sealed class GetListListenerLayer4ListenerResult
    {
        /// <summary>
        /// Origin port, value 1~65535.
        /// </summary>
        public readonly int BackendPort;
        /// <summary>
        /// Forwarding port, value 1~65535.
        /// </summary>
        public readonly int FrontendPort;
        /// <summary>
        /// Resource instance to which the rule belongs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailRuleResult> InstanceDetailRules;
        /// <summary>
        /// InstanceDetails.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailResult> InstanceDetails;
        /// <summary>
        /// Protocol.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Source server list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListListenerLayer4ListenerRealServerResult> RealServers;

        [OutputConstructor]
        private GetListListenerLayer4ListenerResult(
            int backendPort,

            int frontendPort,

            ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailRuleResult> instanceDetailRules,

            ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailResult> instanceDetails,

            string protocol,

            ImmutableArray<Outputs.GetListListenerLayer4ListenerRealServerResult> realServers)
        {
            BackendPort = backendPort;
            FrontendPort = frontendPort;
            InstanceDetailRules = instanceDetailRules;
            InstanceDetails = instanceDetails;
            Protocol = protocol;
            RealServers = realServers;
        }
    }
}
