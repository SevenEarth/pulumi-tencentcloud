// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse.Outputs
{

    [OutputType]
    public sealed class InstanceContainerPublishPort
    {
        /// <summary>
        /// Container port.
        /// </summary>
        public readonly int ContainerPort;
        /// <summary>
        /// Host port.
        /// </summary>
        public readonly int HostPort;
        /// <summary>
        /// External IP. It defaults to 0.0.0.0.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// The protocol defaults to tcp. Valid values: tcp, udp and sctp.
        /// </summary>
        public readonly string? Protocol;

        [OutputConstructor]
        private InstanceContainerPublishPort(
            int containerPort,

            int hostPort,

            string? ip,

            string? protocol)
        {
            ContainerPort = containerPort;
            HostPort = hostPort;
            Ip = ip;
            Protocol = protocol;
        }
    }
}
