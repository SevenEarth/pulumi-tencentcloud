// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class CngwGatewayInstancePortGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Http port range.
        /// </summary>
        [Input("httpPort")]
        public Input<string>? HttpPort { get; set; }

        /// <summary>
        /// Https port range.
        /// </summary>
        [Input("httpsPort")]
        public Input<string>? HttpsPort { get; set; }

        /// <summary>
        /// Tcp port range.
        /// </summary>
        [Input("tcpPort")]
        public Input<string>? TcpPort { get; set; }

        /// <summary>
        /// Udp port range.
        /// </summary>
        [Input("udpPort")]
        public Input<string>? UdpPort { get; set; }

        public CngwGatewayInstancePortGetArgs()
        {
        }
    }
}
