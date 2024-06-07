// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class FlowInputGroupSrtSettingsSourceAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Peer IP.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// Peer port.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public FlowInputGroupSrtSettingsSourceAddressArgs()
        {
        }
        public static new FlowInputGroupSrtSettingsSourceAddressArgs Empty => new FlowInputGroupSrtSettingsSourceAddressArgs();
    }
}
