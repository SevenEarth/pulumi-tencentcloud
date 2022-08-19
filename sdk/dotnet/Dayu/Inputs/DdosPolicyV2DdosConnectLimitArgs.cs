// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Inputs
{

    public sealed class DdosPolicyV2DdosConnectLimitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Based on connection suppression trigger threshold, value range [0,4294967295].
        /// </summary>
        [Input("badConnThreshold", required: true)]
        public Input<int> BadConnThreshold { get; set; } = null!;

        /// <summary>
        /// Abnormal connection detection condition, connection timeout, value range [0,65535].
        /// </summary>
        [Input("connTimeout", required: true)]
        public Input<int> ConnTimeout { get; set; } = null!;

        /// <summary>
        /// Concurrent connection control based on destination IP+ destination port.
        /// </summary>
        [Input("dstConnLimit", required: true)]
        public Input<int> DstConnLimit { get; set; } = null!;

        /// <summary>
        /// Limit on the number of news per second based on the destination IP.
        /// </summary>
        [Input("dstNewLimit", required: true)]
        public Input<int> DstNewLimit { get; set; } = null!;

        /// <summary>
        /// Abnormal connection detection conditions, empty connection guard switch, value range[0,1].
        /// </summary>
        [Input("nullConnEnable", required: true)]
        public Input<int> NullConnEnable { get; set; } = null!;

        /// <summary>
        /// Concurrent connection control based on source IP + destination IP.
        /// </summary>
        [Input("sdConnLimit", required: true)]
        public Input<int> SdConnLimit { get; set; } = null!;

        /// <summary>
        /// The limit on the number of news per second based on source IP + destination IP.
        /// </summary>
        [Input("sdNewLimit", required: true)]
        public Input<int> SdNewLimit { get; set; } = null!;

        /// <summary>
        /// Anomaly connection detection condition, syn threshold, value range [0,100].
        /// </summary>
        [Input("synLimit", required: true)]
        public Input<int> SynLimit { get; set; } = null!;

        /// <summary>
        /// Anomalous connection detection condition, percentage of syn ack, value range [0,100].
        /// </summary>
        [Input("synRate", required: true)]
        public Input<int> SynRate { get; set; } = null!;

        public DdosPolicyV2DdosConnectLimitArgs()
        {
        }
    }
}
