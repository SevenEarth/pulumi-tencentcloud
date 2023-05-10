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

    public sealed class TranscodeTemplateEnhanceConfigVideoEnhanceHdrArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Capability configuration switch, optional value: ON/OFF.Default value: ON.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        /// <summary>
        /// Type, optional value: HDR10/HLG.Default value: HDR10.Note: The encoding method of video needs to be libx265.Note: Video encoding bit depth is 10.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TranscodeTemplateEnhanceConfigVideoEnhanceHdrArgs()
        {
        }
    }
}
