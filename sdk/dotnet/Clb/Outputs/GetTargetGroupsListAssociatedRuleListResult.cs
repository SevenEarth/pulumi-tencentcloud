// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb.Outputs
{

    [OutputType]
    public sealed class GetTargetGroupsListAssociatedRuleListResult
    {
        /// <summary>
        /// Forwarding rule domain.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// Listener ID.
        /// </summary>
        public readonly string ListenerId;
        /// <summary>
        /// Listener name.
        /// </summary>
        public readonly string ListenerName;
        /// <summary>
        /// Listener port.
        /// </summary>
        public readonly int ListenerPort;
        /// <summary>
        /// Load balance ID.
        /// </summary>
        public readonly string LoadBalancerId;
        /// <summary>
        /// Load balance name.
        /// </summary>
        public readonly string LoadBalancerName;
        /// <summary>
        /// Forwarding rule ID.
        /// </summary>
        public readonly string LocationId;
        /// <summary>
        /// Listener protocol type.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Forwarding rule URL.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetTargetGroupsListAssociatedRuleListResult(
            string domain,

            string listenerId,

            string listenerName,

            int listenerPort,

            string loadBalancerId,

            string loadBalancerName,

            string locationId,

            string protocol,

            string url)
        {
            Domain = domain;
            ListenerId = listenerId;
            ListenerName = listenerName;
            ListenerPort = listenerPort;
            LoadBalancerId = loadBalancerId;
            LoadBalancerName = loadBalancerName;
            LocationId = locationId;
            Protocol = protocol;
            Url = url;
        }
    }
}
