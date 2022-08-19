// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Inputs
{

    public sealed class DdosPolicyDropOptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of new connections based on destination IP that trigger suppression of connections. Valid value ranges: (0~4294967295).
        /// </summary>
        [Input("badConnThreshold", required: true)]
        public Input<int> BadConnThreshold { get; set; } = null!;

        /// <summary>
        /// Indicate whether to check null connection or not.
        /// </summary>
        [Input("checkSyncConn", required: true)]
        public Input<bool> CheckSyncConn { get; set; } = null!;

        /// <summary>
        /// Connection timeout of abnormal connection check. Valid value ranges: (0~65535).
        /// </summary>
        [Input("connTimeout", required: true)]
        public Input<int> ConnTimeout { get; set; } = null!;

        /// <summary>
        /// The limit of concurrent connections based on destination IP. Valid value ranges: (0~4294967295).
        /// </summary>
        [Input("dConnLimit", required: true)]
        public Input<int> DConnLimit { get; set; } = null!;

        /// <summary>
        /// The limit of new connections based on destination IP. Valid value ranges: (0~4294967295).
        /// </summary>
        [Input("dNewLimit", required: true)]
        public Input<int> DNewLimit { get; set; } = null!;

        /// <summary>
        /// Indicate whether to drop abroad traffic or not.
        /// </summary>
        [Input("dropAbroad", required: true)]
        public Input<bool> DropAbroad { get; set; } = null!;

        /// <summary>
        /// Indicate whether to drop ICMP protocol or not.
        /// </summary>
        [Input("dropIcmp", required: true)]
        public Input<bool> DropIcmp { get; set; } = null!;

        /// <summary>
        /// Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.
        /// </summary>
        [Input("dropOther", required: true)]
        public Input<bool> DropOther { get; set; } = null!;

        /// <summary>
        /// Indicate whether to drop TCP protocol or not.
        /// </summary>
        [Input("dropTcp", required: true)]
        public Input<bool> DropTcp { get; set; } = null!;

        /// <summary>
        /// Indicate to drop UDP protocol or not.
        /// </summary>
        [Input("dropUdp", required: true)]
        public Input<bool> DropUdp { get; set; } = null!;

        /// <summary>
        /// The limit of ICMP traffic rate. Valid value ranges: (0~4294967295)(Mbps).
        /// </summary>
        [Input("icmpMbpsLimit", required: true)]
        public Input<int> IcmpMbpsLimit { get; set; } = null!;

        /// <summary>
        /// Indicate to enable null connection or not.
        /// </summary>
        [Input("nullConnEnable", required: true)]
        public Input<bool> NullConnEnable { get; set; } = null!;

        /// <summary>
        /// The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate. Valid value ranges: (0~4294967295)(Mbps).
        /// </summary>
        [Input("otherMbpsLimit", required: true)]
        public Input<int> OtherMbpsLimit { get; set; } = null!;

        /// <summary>
        /// The limit of concurrent connections based on source IP. Valid value ranges: (0~4294967295).
        /// </summary>
        [Input("sConnLimit", required: true)]
        public Input<int> SConnLimit { get; set; } = null!;

        /// <summary>
        /// The limit of new connections based on source IP. Valid value ranges: (0~4294967295).
        /// </summary>
        [Input("sNewLimit", required: true)]
        public Input<int> SNewLimit { get; set; } = null!;

        /// <summary>
        /// The limit of syn of abnormal connection check. Valid value ranges: (0~100).
        /// </summary>
        [Input("synLimit", required: true)]
        public Input<int> SynLimit { get; set; } = null!;

        /// <summary>
        /// The percentage of syn in ack of abnormal connection check. Valid value ranges: (0~100).
        /// </summary>
        [Input("synRate")]
        public Input<int>? SynRate { get; set; }

        /// <summary>
        /// The limit of TCP traffic. Valid value ranges: (0~4294967295)(Mbps).
        /// </summary>
        [Input("tcpMbpsLimit", required: true)]
        public Input<int> TcpMbpsLimit { get; set; } = null!;

        /// <summary>
        /// The limit of UDP traffic rate. Valid value ranges: (0~4294967295)(Mbps).
        /// </summary>
        [Input("udpMbpsLimit", required: true)]
        public Input<int> UdpMbpsLimit { get; set; } = null!;

        public DdosPolicyDropOptionGetArgs()
        {
        }
    }
}
