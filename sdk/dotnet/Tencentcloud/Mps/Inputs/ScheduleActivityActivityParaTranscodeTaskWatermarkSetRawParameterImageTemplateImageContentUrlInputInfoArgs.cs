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

    public sealed class ScheduleActivityActivityParaTranscodeTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL of a video.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ScheduleActivityActivityParaTranscodeTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs()
        {
        }
        public static new ScheduleActivityActivityParaTranscodeTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs Empty => new ScheduleActivityActivityParaTranscodeTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs();
    }
}
