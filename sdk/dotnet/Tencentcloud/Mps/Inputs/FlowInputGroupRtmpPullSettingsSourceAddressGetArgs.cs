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

    public sealed class FlowInputGroupRtmpPullSettingsSourceAddressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// StreamKey information of the RTMP source site.
        /// </summary>
        [Input("streamKey", required: true)]
        public Input<string> StreamKey { get; set; } = null!;

        /// <summary>
        /// TcUrl address of the RTMP source server.
        /// </summary>
        [Input("tcUrl", required: true)]
        public Input<string> TcUrl { get; set; } = null!;

        public FlowInputGroupRtmpPullSettingsSourceAddressGetArgs()
        {
        }
    }
}
