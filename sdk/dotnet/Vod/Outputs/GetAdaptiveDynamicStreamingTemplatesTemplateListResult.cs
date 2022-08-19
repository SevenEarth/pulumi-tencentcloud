// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class GetAdaptiveDynamicStreamingTemplatesTemplateListResult
    {
        /// <summary>
        /// Template description.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Creation time of template in ISO date format.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Unique ID filter of adaptive dynamic streaming template.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// Whether to prohibit transcoding video from low bitrate to high bitrate. `false`: no, `true`: yes.
        /// </summary>
        public readonly bool DisableHigherVideoBitrate;
        /// <summary>
        /// Whether to prohibit transcoding from low resolution to high resolution. `false`: no, `true`: yes.
        /// </summary>
        public readonly bool DisableHigherVideoResolution;
        /// <summary>
        /// DRM scheme type.
        /// </summary>
        public readonly string DrmType;
        /// <summary>
        /// Adaptive bitstream format.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// Template name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdaptiveDynamicStreamingTemplatesTemplateListStreamInfoResult> StreamInfos;
        /// <summary>
        /// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Last modified time of template in ISO date format.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAdaptiveDynamicStreamingTemplatesTemplateListResult(
            string comment,

            string createTime,

            string definition,

            bool disableHigherVideoBitrate,

            bool disableHigherVideoResolution,

            string drmType,

            string format,

            string name,

            ImmutableArray<Outputs.GetAdaptiveDynamicStreamingTemplatesTemplateListStreamInfoResult> streamInfos,

            string type,

            string updateTime)
        {
            Comment = comment;
            CreateTime = createTime;
            Definition = definition;
            DisableHigherVideoBitrate = disableHigherVideoBitrate;
            DisableHigherVideoResolution = disableHigherVideoResolution;
            DrmType = drmType;
            Format = format;
            Name = name;
            StreamInfos = streamInfos;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
