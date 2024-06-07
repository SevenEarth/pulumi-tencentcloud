// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse.Inputs
{

    public sealed class InstanceContainerPublishPortGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container port.
        /// </summary>
        [Input("containerPort", required: true)]
        public Input<int> ContainerPort { get; set; } = null!;

        /// <summary>
        /// Host port.
        /// </summary>
        [Input("hostPort", required: true)]
        public Input<int> HostPort { get; set; } = null!;

        /// <summary>
        /// External IP. It defaults to 0.0.0.0.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The protocol defaults to tcp. Valid values: tcp, udp and sctp.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public InstanceContainerPublishPortGetArgs()
        {
        }
        public static new InstanceContainerPublishPortGetArgs Empty => new InstanceContainerPublishPortGetArgs();
    }
}
