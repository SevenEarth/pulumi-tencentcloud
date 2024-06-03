// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos.Inputs
{

    public sealed class PortAclConfigAclConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action, can take values: drop, transmit, forward.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// end from port, with a range of 0~65535 values.
        /// </summary>
        [Input("dPortEnd", required: true)]
        public Input<int> DPortEnd { get; set; } = null!;

        /// <summary>
        /// Starting from port, with a range of 0~65535 values.
        /// </summary>
        [Input("dPortStart", required: true)]
        public Input<int> DPortStart { get; set; } = null!;

        /// <summary>
        /// Protocol type, can take TCP, udp, all values.
        /// </summary>
        [Input("forwardProtocol", required: true)]
        public Input<string> ForwardProtocol { get; set; } = null!;

        /// <summary>
        /// The policy priority, the smaller the number, the higher the level, and the higher the matching of the rule, with values ranging from 1 to 1000. Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// end from the source port, with a value range of 0~65535.
        /// </summary>
        [Input("sPortEnd", required: true)]
        public Input<int> SPortEnd { get; set; } = null!;

        /// <summary>
        /// Starting from the source port, with a value range of 0~65535.
        /// </summary>
        [Input("sPortStart", required: true)]
        public Input<int> SPortStart { get; set; } = null!;

        public PortAclConfigAclConfigArgs()
        {
        }
        public static new PortAclConfigAclConfigArgs Empty => new PortAclConfigAclConfigArgs();
    }
}
