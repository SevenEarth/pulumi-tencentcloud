// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod
{
    /// <summary>
    /// Provide a resource to create a VOD adaptive dynamic streaming template.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Vod.AdaptiveDynamicStreamingTemplate("foo", new Tencentcloud.Vod.AdaptiveDynamicStreamingTemplateArgs
    ///         {
    ///             Comment = "test",
    ///             DisableHigherVideoBitrate = false,
    ///             DisableHigherVideoResolution = false,
    ///             DrmType = "SimpleAES",
    ///             Format = "HLS",
    ///             StreamInfos = 
    ///             {
    ///                 new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs
    ///                 {
    ///                     Audio = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs
    ///                     {
    ///                         AudioChannel = "dual",
    ///                         Bitrate = 129,
    ///                         Codec = "libmp3lame",
    ///                         SampleRate = 44100,
    ///                     },
    ///                     RemoveAudio = false,
    ///                     Video = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs
    ///                     {
    ///                         Bitrate = 129,
    ///                         Codec = "libx265",
    ///                         FillType = "stretch",
    ///                         Fps = 4,
    ///                         Height = 128,
    ///                         ResolutionAdaptive = false,
    ///                         Width = 128,
    ///                     },
    ///                 },
    ///                 new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs
    ///                 {
    ///                     Audio = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs
    ///                     {
    ///                         Bitrate = 256,
    ///                         Codec = "libfdk_aac",
    ///                         SampleRate = 44100,
    ///                     },
    ///                     RemoveAudio = true,
    ///                     Video = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs
    ///                     {
    ///                         Bitrate = 256,
    ///                         Codec = "libx264",
    ///                         Fps = 4,
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VOD adaptive dynamic streaming template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate foo 169141
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate")]
    public partial class AdaptiveDynamicStreamingTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Template description. Length limit: 256 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Creation time of template in ISO date format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        /// </summary>
        [Output("disableHigherVideoBitrate")]
        public Output<bool?> DisableHigherVideoBitrate { get; private set; } = null!;

        /// <summary>
        /// Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        /// </summary>
        [Output("disableHigherVideoResolution")]
        public Output<bool?> DisableHigherVideoResolution { get; private set; } = null!;

        /// <summary>
        /// DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        /// </summary>
        [Output("drmType")]
        public Output<string?> DrmType { get; private set; } = null!;

        /// <summary>
        /// Adaptive bitstream format. Valid values: `HLS`.
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// Template name. Length limit: 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        /// </summary>
        [Output("streamInfos")]
        public Output<ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplateStreamInfo>> StreamInfos { get; private set; } = null!;

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Output("subAppId")]
        public Output<int?> SubAppId { get; private set; } = null!;

        /// <summary>
        /// Last modified time of template in ISO date format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AdaptiveDynamicStreamingTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdaptiveDynamicStreamingTemplate(string name, AdaptiveDynamicStreamingTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate", name, args ?? new AdaptiveDynamicStreamingTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdaptiveDynamicStreamingTemplate(string name, Input<string> id, AdaptiveDynamicStreamingTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AdaptiveDynamicStreamingTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdaptiveDynamicStreamingTemplate Get(string name, Input<string> id, AdaptiveDynamicStreamingTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new AdaptiveDynamicStreamingTemplate(name, id, state, options);
        }
    }

    public sealed class AdaptiveDynamicStreamingTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template description. Length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        /// </summary>
        [Input("disableHigherVideoBitrate")]
        public Input<bool>? DisableHigherVideoBitrate { get; set; }

        /// <summary>
        /// Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        /// </summary>
        [Input("disableHigherVideoResolution")]
        public Input<bool>? DisableHigherVideoResolution { get; set; }

        /// <summary>
        /// DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        /// </summary>
        [Input("drmType")]
        public Input<string>? DrmType { get; set; }

        /// <summary>
        /// Adaptive bitstream format. Valid values: `HLS`.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// Template name. Length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("streamInfos", required: true)]
        private InputList<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs>? _streamInfos;

        /// <summary>
        /// List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        /// </summary>
        public InputList<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs> StreamInfos
        {
            get => _streamInfos ?? (_streamInfos = new InputList<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs>());
            set => _streamInfos = value;
        }

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        public AdaptiveDynamicStreamingTemplateArgs()
        {
        }
    }

    public sealed class AdaptiveDynamicStreamingTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template description. Length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Creation time of template in ISO date format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        /// </summary>
        [Input("disableHigherVideoBitrate")]
        public Input<bool>? DisableHigherVideoBitrate { get; set; }

        /// <summary>
        /// Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        /// </summary>
        [Input("disableHigherVideoResolution")]
        public Input<bool>? DisableHigherVideoResolution { get; set; }

        /// <summary>
        /// DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        /// </summary>
        [Input("drmType")]
        public Input<string>? DrmType { get; set; }

        /// <summary>
        /// Adaptive bitstream format. Valid values: `HLS`.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Template name. Length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("streamInfos")]
        private InputList<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoGetArgs>? _streamInfos;

        /// <summary>
        /// List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        /// </summary>
        public InputList<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoGetArgs> StreamInfos
        {
            get => _streamInfos ?? (_streamInfos = new InputList<Inputs.AdaptiveDynamicStreamingTemplateStreamInfoGetArgs>());
            set => _streamInfos = value;
        }

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        /// <summary>
        /// Last modified time of template in ISO date format.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AdaptiveDynamicStreamingTemplateState()
        {
        }
    }
}
