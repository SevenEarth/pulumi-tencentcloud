// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class DdosPolicyWatermarkFilter
    {
        /// <summary>
        /// Indicate whether to auto-remove the watermark or not.
        /// </summary>
        public readonly bool? AutoRemove;
        /// <summary>
        /// The offset of watermark. Valid value ranges: (0~1500).
        /// </summary>
        public readonly int? Offset;
        /// <summary>
        /// Indicate whether to open watermark or not. It muse be set `true` when any field of watermark was set.
        /// </summary>
        public readonly bool? OpenSwitch;
        /// <summary>
        /// Port range of TCP, the format is like `2000-3000`.
        /// </summary>
        public readonly ImmutableArray<string> TcpPortLists;
        /// <summary>
        /// Port range of TCP, the format is like `2000-3000`.
        /// </summary>
        public readonly ImmutableArray<string> UdpPortLists;

        [OutputConstructor]
        private DdosPolicyWatermarkFilter(
            bool? autoRemove,

            int? offset,

            bool? openSwitch,

            ImmutableArray<string> tcpPortLists,

            ImmutableArray<string> udpPortLists)
        {
            AutoRemove = autoRemove;
            Offset = offset;
            OpenSwitch = openSwitch;
            TcpPortLists = tcpPortLists;
            UdpPortLists = udpPortLists;
        }
    }
}
