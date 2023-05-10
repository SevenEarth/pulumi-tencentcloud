// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class TranscodeTemplateEnhanceConfigVideoEnhanceFrameRate
    {
        /// <summary>
        /// Frame rate, value range: [0, 100], unit: Hz.Default value: 0.Note: For transcoding, this parameter will override the Fps inside the VideoTemplate.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? Fps;
        /// <summary>
        /// Capability configuration switch, optional value: ON/OFF.Default value: ON.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private TranscodeTemplateEnhanceConfigVideoEnhanceFrameRate(
            int? fps,

            string? @switch)
        {
            Fps = fps;
            Switch = @switch;
        }
    }
}
