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
    public sealed class ScheduleActivityActivityParaTranscodeTaskOverrideParameterTehdConfig
    {
        /// <summary>
        /// Maximum bitrate, which is valid when `Type` is `TESHD`.If this parameter is left empty or 0 is entered, there will be no upper limit for bitrate.
        /// </summary>
        public readonly int? MaxVideoBitrate;
        /// <summary>
        /// TESHD type. Valid values: TEHD-100: TESHD-100.If this parameter is left empty, TESHD will not be enabled.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ScheduleActivityActivityParaTranscodeTaskOverrideParameterTehdConfig(
            int? maxVideoBitrate,

            string? type)
        {
            MaxVideoBitrate = maxVideoBitrate;
            Type = type;
        }
    }
}
