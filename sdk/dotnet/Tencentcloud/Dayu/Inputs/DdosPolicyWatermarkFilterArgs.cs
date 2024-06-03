// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu.Inputs
{

    public sealed class DdosPolicyWatermarkFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicate whether to auto-remove the watermark or not.
        /// </summary>
        [Input("autoRemove")]
        public Input<bool>? AutoRemove { get; set; }

        /// <summary>
        /// The offset of watermark. Valid value ranges: (0~1500).
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        /// <summary>
        /// Indicate whether to open watermark or not. It muse be set `true` when any field of watermark was set.
        /// </summary>
        [Input("openSwitch")]
        public Input<bool>? OpenSwitch { get; set; }

        [Input("tcpPortLists")]
        private InputList<string>? _tcpPortLists;

        /// <summary>
        /// Port range of TCP, the format is like `2000-3000`.
        /// </summary>
        public InputList<string> TcpPortLists
        {
            get => _tcpPortLists ?? (_tcpPortLists = new InputList<string>());
            set => _tcpPortLists = value;
        }

        [Input("udpPortLists")]
        private InputList<string>? _udpPortLists;

        /// <summary>
        /// Port range of TCP, the format is like `2000-3000`.
        /// </summary>
        public InputList<string> UdpPortLists
        {
            get => _udpPortLists ?? (_udpPortLists = new InputList<string>());
            set => _udpPortLists = value;
        }

        public DdosPolicyWatermarkFilterArgs()
        {
        }
        public static new DdosPolicyWatermarkFilterArgs Empty => new DdosPolicyWatermarkFilterArgs();
    }
}
