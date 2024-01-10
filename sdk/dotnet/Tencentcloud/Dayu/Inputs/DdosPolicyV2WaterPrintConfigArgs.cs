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

    public sealed class DdosPolicyV2WaterPrintConfigArgs : Pulumi.ResourceArgs
    {
        [Input("listeners", required: true)]
        private InputList<Inputs.DdosPolicyV2WaterPrintConfigListenerArgs>? _listeners;

        /// <summary>
        /// List of forwarding listeners to which the watermark belongs.
        /// </summary>
        public InputList<Inputs.DdosPolicyV2WaterPrintConfigListenerArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.DdosPolicyV2WaterPrintConfigListenerArgs>());
            set => _listeners = value;
        }

        /// <summary>
        /// Watermark offset, value range: [0-100].
        /// </summary>
        [Input("offset", required: true)]
        public Input<int> Offset { get; set; } = null!;

        /// <summary>
        /// Whether it is enabled, value [0 (manual open), 1 (immediate operation)].
        /// </summary>
        [Input("openStatus", required: true)]
        public Input<int> OpenStatus { get; set; } = null!;

        /// <summary>
        /// Watermark check mode, value [`checkall`(normal mode), `shortfpcheckall`(simplified mode)].
        /// </summary>
        [Input("verify", required: true)]
        public Input<string> Verify { get; set; } = null!;

        public DdosPolicyV2WaterPrintConfigArgs()
        {
        }
    }
}
