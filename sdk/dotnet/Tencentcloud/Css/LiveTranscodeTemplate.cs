// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css
{
    /// <summary>
    /// Provides a resource to create a css live_transcode_template
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var liveTranscodeTemplate = new Tencentcloud.Css.LiveTranscodeTemplate("liveTranscodeTemplate", new()
    ///     {
    ///         Acodec = "aac",
    ///         AdaptBitratePercent = 0,
    ///         AiTransCode = 0,
    ///         AudioBitrate = 128,
    ///         BitrateToOrig = 0,
    ///         Description = "This_is_a_tf_test_temp.",
    ///         DrmTracks = "SD",
    ///         DrmType = "fairplay",
    ///         Fps = 0,
    ///         FpsToOrig = 0,
    ///         Gop = 2,
    ///         Height = 0,
    ///         HeightToOrig = 0,
    ///         NeedAudio = 1,
    ///         NeedVideo = 1,
    ///         Profile = "baseline",
    ///         Rotate = 0,
    ///         ShortEdgeAsHeight = 0,
    ///         TemplateName = "template_name",
    ///         Vcodec = "origin",
    ///         VideoBitrate = 100,
    ///         Width = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// css live_transcode_template can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Css/liveTranscodeTemplate:LiveTranscodeTemplate live_transcode_template liveTranscodeTemplate_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Css/liveTranscodeTemplate:LiveTranscodeTemplate")]
    public partial class LiveTranscodeTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// default aac, not support now.
        /// </summary>
        [Output("acodec")]
        public Output<string?> Acodec { get; private set; } = null!;

        /// <summary>
        /// high speed mode adapt bitrate, support 0 - 0.5.
        /// </summary>
        [Output("adaptBitratePercent")]
        public Output<double?> AdaptBitratePercent { get; private set; } = null!;

        /// <summary>
        /// enable high speed mode, default 0, 1 for enable, 0 for no.
        /// </summary>
        [Output("aiTransCode")]
        public Output<int?> AiTransCode { get; private set; } = null!;

        /// <summary>
        /// default 0, range 0 - 500.
        /// </summary>
        [Output("audioBitrate")]
        public Output<int?> AudioBitrate { get; private set; } = null!;

        /// <summary>
        /// base on origin bitrate if origin bitrate is lower than the setting bitrate. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Output("bitrateToOrig")]
        public Output<int?> BitrateToOrig { get; private set; } = null!;

        /// <summary>
        /// template desc.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// DRM tracks, support AUDIO/SD/HD/UHD1/UHD2.
        /// </summary>
        [Output("drmTracks")]
        public Output<string?> DrmTracks { get; private set; } = null!;

        /// <summary>
        /// DRM type, support fairplay/normalaes/widevine.
        /// </summary>
        [Output("drmType")]
        public Output<string?> DrmType { get; private set; } = null!;

        /// <summary>
        /// video fps, default 0, range 0 - 60.
        /// </summary>
        [Output("fps")]
        public Output<int?> Fps { get; private set; } = null!;

        /// <summary>
        /// base on origin fps if origin fps is lower than the setting fps. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Output("fpsToOrig")]
        public Output<int?> FpsToOrig { get; private set; } = null!;

        /// <summary>
        /// gop of the video, second, default origin of the video, range 2 - 6.
        /// </summary>
        [Output("gop")]
        public Output<int?> Gop { get; private set; } = null!;

        /// <summary>
        /// template height, default 0, range 0 - 3000, must be pow of 2, needed while AiTransCode = 1.
        /// </summary>
        [Output("height")]
        public Output<int?> Height { get; private set; } = null!;

        /// <summary>
        /// base on origin height if origin height is lower than the setting height. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Output("heightToOrig")]
        public Output<int?> HeightToOrig { get; private set; } = null!;

        /// <summary>
        /// keep audio or not, default 1 for yes, 0 for no.
        /// </summary>
        [Output("needAudio")]
        public Output<int?> NeedAudio { get; private set; } = null!;

        /// <summary>
        /// keep video or not, default 1 for yes, 0 for no.
        /// </summary>
        [Output("needVideo")]
        public Output<int?> NeedVideo { get; private set; } = null!;

        /// <summary>
        /// quality of the video, default baseline, support baseline/main/high.
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// roate degree, default 0, support 0/90/180/270.
        /// </summary>
        [Output("rotate")]
        public Output<int?> Rotate { get; private set; } = null!;

        /// <summary>
        /// let the short edge as the height.
        /// </summary>
        [Output("shortEdgeAsHeight")]
        public Output<int?> ShortEdgeAsHeight { get; private set; } = null!;

        /// <summary>
        /// template name, only support 0-9 and a-z.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// video codec, default origin, support h264/h265/origin.
        /// </summary>
        [Output("vcodec")]
        public Output<string?> Vcodec { get; private set; } = null!;

        /// <summary>
        /// video bitrate, 0 for origin, range 0kbps - 8000kbps.
        /// </summary>
        [Output("videoBitrate")]
        public Output<int> VideoBitrate { get; private set; } = null!;

        /// <summary>
        /// template width, default 0, range 0 - 3000, must be pow of 2.
        /// </summary>
        [Output("width")]
        public Output<int?> Width { get; private set; } = null!;


        /// <summary>
        /// Create a LiveTranscodeTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LiveTranscodeTemplate(string name, LiveTranscodeTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/liveTranscodeTemplate:LiveTranscodeTemplate", name, args ?? new LiveTranscodeTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LiveTranscodeTemplate(string name, Input<string> id, LiveTranscodeTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/liveTranscodeTemplate:LiveTranscodeTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LiveTranscodeTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LiveTranscodeTemplate Get(string name, Input<string> id, LiveTranscodeTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new LiveTranscodeTemplate(name, id, state, options);
        }
    }

    public sealed class LiveTranscodeTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// default aac, not support now.
        /// </summary>
        [Input("acodec")]
        public Input<string>? Acodec { get; set; }

        /// <summary>
        /// high speed mode adapt bitrate, support 0 - 0.5.
        /// </summary>
        [Input("adaptBitratePercent")]
        public Input<double>? AdaptBitratePercent { get; set; }

        /// <summary>
        /// enable high speed mode, default 0, 1 for enable, 0 for no.
        /// </summary>
        [Input("aiTransCode")]
        public Input<int>? AiTransCode { get; set; }

        /// <summary>
        /// default 0, range 0 - 500.
        /// </summary>
        [Input("audioBitrate")]
        public Input<int>? AudioBitrate { get; set; }

        /// <summary>
        /// base on origin bitrate if origin bitrate is lower than the setting bitrate. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Input("bitrateToOrig")]
        public Input<int>? BitrateToOrig { get; set; }

        /// <summary>
        /// template desc.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// DRM tracks, support AUDIO/SD/HD/UHD1/UHD2.
        /// </summary>
        [Input("drmTracks")]
        public Input<string>? DrmTracks { get; set; }

        /// <summary>
        /// DRM type, support fairplay/normalaes/widevine.
        /// </summary>
        [Input("drmType")]
        public Input<string>? DrmType { get; set; }

        /// <summary>
        /// video fps, default 0, range 0 - 60.
        /// </summary>
        [Input("fps")]
        public Input<int>? Fps { get; set; }

        /// <summary>
        /// base on origin fps if origin fps is lower than the setting fps. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Input("fpsToOrig")]
        public Input<int>? FpsToOrig { get; set; }

        /// <summary>
        /// gop of the video, second, default origin of the video, range 2 - 6.
        /// </summary>
        [Input("gop")]
        public Input<int>? Gop { get; set; }

        /// <summary>
        /// template height, default 0, range 0 - 3000, must be pow of 2, needed while AiTransCode = 1.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// base on origin height if origin height is lower than the setting height. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Input("heightToOrig")]
        public Input<int>? HeightToOrig { get; set; }

        /// <summary>
        /// keep audio or not, default 1 for yes, 0 for no.
        /// </summary>
        [Input("needAudio")]
        public Input<int>? NeedAudio { get; set; }

        /// <summary>
        /// keep video or not, default 1 for yes, 0 for no.
        /// </summary>
        [Input("needVideo")]
        public Input<int>? NeedVideo { get; set; }

        /// <summary>
        /// quality of the video, default baseline, support baseline/main/high.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// roate degree, default 0, support 0/90/180/270.
        /// </summary>
        [Input("rotate")]
        public Input<int>? Rotate { get; set; }

        /// <summary>
        /// let the short edge as the height.
        /// </summary>
        [Input("shortEdgeAsHeight")]
        public Input<int>? ShortEdgeAsHeight { get; set; }

        /// <summary>
        /// template name, only support 0-9 and a-z.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// video codec, default origin, support h264/h265/origin.
        /// </summary>
        [Input("vcodec")]
        public Input<string>? Vcodec { get; set; }

        /// <summary>
        /// video bitrate, 0 for origin, range 0kbps - 8000kbps.
        /// </summary>
        [Input("videoBitrate", required: true)]
        public Input<int> VideoBitrate { get; set; } = null!;

        /// <summary>
        /// template width, default 0, range 0 - 3000, must be pow of 2.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public LiveTranscodeTemplateArgs()
        {
        }
        public static new LiveTranscodeTemplateArgs Empty => new LiveTranscodeTemplateArgs();
    }

    public sealed class LiveTranscodeTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// default aac, not support now.
        /// </summary>
        [Input("acodec")]
        public Input<string>? Acodec { get; set; }

        /// <summary>
        /// high speed mode adapt bitrate, support 0 - 0.5.
        /// </summary>
        [Input("adaptBitratePercent")]
        public Input<double>? AdaptBitratePercent { get; set; }

        /// <summary>
        /// enable high speed mode, default 0, 1 for enable, 0 for no.
        /// </summary>
        [Input("aiTransCode")]
        public Input<int>? AiTransCode { get; set; }

        /// <summary>
        /// default 0, range 0 - 500.
        /// </summary>
        [Input("audioBitrate")]
        public Input<int>? AudioBitrate { get; set; }

        /// <summary>
        /// base on origin bitrate if origin bitrate is lower than the setting bitrate. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Input("bitrateToOrig")]
        public Input<int>? BitrateToOrig { get; set; }

        /// <summary>
        /// template desc.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// DRM tracks, support AUDIO/SD/HD/UHD1/UHD2.
        /// </summary>
        [Input("drmTracks")]
        public Input<string>? DrmTracks { get; set; }

        /// <summary>
        /// DRM type, support fairplay/normalaes/widevine.
        /// </summary>
        [Input("drmType")]
        public Input<string>? DrmType { get; set; }

        /// <summary>
        /// video fps, default 0, range 0 - 60.
        /// </summary>
        [Input("fps")]
        public Input<int>? Fps { get; set; }

        /// <summary>
        /// base on origin fps if origin fps is lower than the setting fps. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Input("fpsToOrig")]
        public Input<int>? FpsToOrig { get; set; }

        /// <summary>
        /// gop of the video, second, default origin of the video, range 2 - 6.
        /// </summary>
        [Input("gop")]
        public Input<int>? Gop { get; set; }

        /// <summary>
        /// template height, default 0, range 0 - 3000, must be pow of 2, needed while AiTransCode = 1.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// base on origin height if origin height is lower than the setting height. default 0, 1 for yes, 0 for no.
        /// </summary>
        [Input("heightToOrig")]
        public Input<int>? HeightToOrig { get; set; }

        /// <summary>
        /// keep audio or not, default 1 for yes, 0 for no.
        /// </summary>
        [Input("needAudio")]
        public Input<int>? NeedAudio { get; set; }

        /// <summary>
        /// keep video or not, default 1 for yes, 0 for no.
        /// </summary>
        [Input("needVideo")]
        public Input<int>? NeedVideo { get; set; }

        /// <summary>
        /// quality of the video, default baseline, support baseline/main/high.
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// roate degree, default 0, support 0/90/180/270.
        /// </summary>
        [Input("rotate")]
        public Input<int>? Rotate { get; set; }

        /// <summary>
        /// let the short edge as the height.
        /// </summary>
        [Input("shortEdgeAsHeight")]
        public Input<int>? ShortEdgeAsHeight { get; set; }

        /// <summary>
        /// template name, only support 0-9 and a-z.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// video codec, default origin, support h264/h265/origin.
        /// </summary>
        [Input("vcodec")]
        public Input<string>? Vcodec { get; set; }

        /// <summary>
        /// video bitrate, 0 for origin, range 0kbps - 8000kbps.
        /// </summary>
        [Input("videoBitrate")]
        public Input<int>? VideoBitrate { get; set; }

        /// <summary>
        /// template width, default 0, range 0 - 3000, must be pow of 2.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public LiveTranscodeTemplateState()
        {
        }
        public static new LiveTranscodeTemplateState Empty => new LiveTranscodeTemplateState();
    }
}
