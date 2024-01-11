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
    public sealed class ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitle
    {
        /// <summary>
        /// The subtitle file.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleSubtitle? Subtitle;
        /// <summary>
        /// The mode. Valid values:`subtitle-stream`: Add a subtitle track.`close-caption-708`: Embed CEA-708 subtitles in SEI frames.`close-caption-608`: Embed CEA-608 subtitles in SEI frames.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitle(
            Outputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleSubtitle? subtitle,

            string? type)
        {
            Subtitle = subtitle;
            Type = type;
        }
    }
}