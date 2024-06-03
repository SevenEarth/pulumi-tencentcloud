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

    public sealed class DdosPolicyV2WaterPrintConfigListenerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Forwarding protocol, value [TCP, UDP].
        /// </summary>
        [Input("forwardProtocol", required: true)]
        public Input<string> ForwardProtocol { get; set; } = null!;

        /// <summary>
        /// Lower limit of forwarding listening port. Values: [1-65535].
        /// </summary>
        [Input("frontendPort", required: true)]
        public Input<int> FrontendPort { get; set; } = null!;

        /// <summary>
        /// Upper limit of forwarding listening port. Values: [1-65535].
        /// </summary>
        [Input("frontendPortEnd", required: true)]
        public Input<int> FrontendPortEnd { get; set; } = null!;

        public DdosPolicyV2WaterPrintConfigListenerGetArgs()
        {
        }
        public static new DdosPolicyV2WaterPrintConfigListenerGetArgs Empty => new DdosPolicyV2WaterPrintConfigListenerGetArgs();
    }
}
