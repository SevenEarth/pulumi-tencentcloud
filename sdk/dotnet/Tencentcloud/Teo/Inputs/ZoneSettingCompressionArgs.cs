// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class ZoneSettingCompressionArgs : Pulumi.ResourceArgs
    {
        [Input("algorithms")]
        private InputList<string>? _algorithms;

        /// <summary>
        /// Compression algorithms to select. Valid values: `brotli`, `gzip`.
        /// </summary>
        public InputList<string> Algorithms
        {
            get => _algorithms ?? (_algorithms = new InputList<string>());
            set => _algorithms = value;
        }

        /// <summary>
        /// Whether to enable Smart compression.- `on`: Enable.- `off`: Disable.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public ZoneSettingCompressionArgs()
        {
        }
    }
}
