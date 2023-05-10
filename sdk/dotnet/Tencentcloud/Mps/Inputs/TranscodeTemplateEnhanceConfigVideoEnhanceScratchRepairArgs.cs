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

    public sealed class TranscodeTemplateEnhanceConfigVideoEnhanceScratchRepairArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Intensity, value range: 0.0~1.0.Default value: 0.0.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("intensity")]
        public Input<double>? Intensity { get; set; }

        /// <summary>
        /// Capability configuration switch, optional value: ON/OFF.Default value: ON.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public TranscodeTemplateEnhanceConfigVideoEnhanceScratchRepairArgs()
        {
        }
    }
}
