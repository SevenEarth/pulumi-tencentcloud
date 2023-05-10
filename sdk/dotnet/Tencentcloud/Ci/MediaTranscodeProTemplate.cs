// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci
{
    /// <summary>
    /// Provides a resource to create a ci media_transcode_pro_template
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
    ///         var mediaTranscodeProTemplate = new Tencentcloud.Ci.MediaTranscodeProTemplate("mediaTranscodeProTemplate", new Tencentcloud.Ci.MediaTranscodeProTemplateArgs
    ///         {
    ///             Audio = new Tencentcloud.Ci.Inputs.MediaTranscodeProTemplateAudioArgs
    ///             {
    ///                 Codec = "pcm_s24le",
    ///                 Remove = "true",
    ///             },
    ///             Bucket = "terraform-ci-xxxxxx",
    ///             Container = new Tencentcloud.Ci.Inputs.MediaTranscodeProTemplateContainerArgs
    ///             {
    ///                 Format = "mxf",
    ///             },
    ///             TimeInterval = new Tencentcloud.Ci.Inputs.MediaTranscodeProTemplateTimeIntervalArgs
    ///             {
    ///                 Duration = "",
    ///                 Start = "",
    ///             },
    ///             TransConfig = new Tencentcloud.Ci.Inputs.MediaTranscodeProTemplateTransConfigArgs
    ///             {
    ///                 AdjDarMethod = "scale",
    ///                 AudioBitrateAdjMethod = "0",
    ///                 DeleteMetadata = "false",
    ///                 IsCheckAudioBitrate = "false",
    ///                 IsCheckReso = "false",
    ///                 IsCheckVideoBitrate = "false",
    ///                 IsHdr2Sdr = "false",
    ///                 ResoAdjMethod = "1",
    ///                 VideoBitrateAdjMethod = "0",
    ///             },
    ///             Video = new Tencentcloud.Ci.Inputs.MediaTranscodeProTemplateVideoArgs
    ///             {
    ///                 Bitrate = "50000",
    ///                 Codec = "xavc",
    ///                 Fps = "30000/1001",
    ///                 Height = "1080",
    ///                 Interlaced = "true",
    ///                 Profile = "XAVC-HD_422_10bit",
    ///                 Width = "1920",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ci media_transcode_pro_template can be imported using the bucket#templateId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ci/mediaTranscodeProTemplate:MediaTranscodeProTemplate media_transcode_pro_template terraform-ci-xxxxxx#t13ed9af009da0414e9c7c63456ec8f4d2
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ci/mediaTranscodeProTemplate:MediaTranscodeProTemplate")]
    public partial class MediaTranscodeProTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        /// </summary>
        [Output("audio")]
        public Output<Outputs.MediaTranscodeProTemplateAudio?> Audio { get; private set; } = null!;

        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// container format.
        /// </summary>
        [Output("container")]
        public Output<Outputs.MediaTranscodeProTemplateContainer> Container { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// time interval.
        /// </summary>
        [Output("timeInterval")]
        public Output<Outputs.MediaTranscodeProTemplateTimeInterval?> TimeInterval { get; private set; } = null!;

        /// <summary>
        /// transcoding configuration.
        /// </summary>
        [Output("transConfig")]
        public Output<Outputs.MediaTranscodeProTemplateTransConfig?> TransConfig { get; private set; } = null!;

        /// <summary>
        /// video information, do not upload Video, which is equivalent to deleting video information.
        /// </summary>
        [Output("video")]
        public Output<Outputs.MediaTranscodeProTemplateVideo?> Video { get; private set; } = null!;


        /// <summary>
        /// Create a MediaTranscodeProTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaTranscodeProTemplate(string name, MediaTranscodeProTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaTranscodeProTemplate:MediaTranscodeProTemplate", name, args ?? new MediaTranscodeProTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaTranscodeProTemplate(string name, Input<string> id, MediaTranscodeProTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaTranscodeProTemplate:MediaTranscodeProTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaTranscodeProTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaTranscodeProTemplate Get(string name, Input<string> id, MediaTranscodeProTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaTranscodeProTemplate(name, id, state, options);
        }
    }

    public sealed class MediaTranscodeProTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        /// </summary>
        [Input("audio")]
        public Input<Inputs.MediaTranscodeProTemplateAudioArgs>? Audio { get; set; }

        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// container format.
        /// </summary>
        [Input("container", required: true)]
        public Input<Inputs.MediaTranscodeProTemplateContainerArgs> Container { get; set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// time interval.
        /// </summary>
        [Input("timeInterval")]
        public Input<Inputs.MediaTranscodeProTemplateTimeIntervalArgs>? TimeInterval { get; set; }

        /// <summary>
        /// transcoding configuration.
        /// </summary>
        [Input("transConfig")]
        public Input<Inputs.MediaTranscodeProTemplateTransConfigArgs>? TransConfig { get; set; }

        /// <summary>
        /// video information, do not upload Video, which is equivalent to deleting video information.
        /// </summary>
        [Input("video")]
        public Input<Inputs.MediaTranscodeProTemplateVideoArgs>? Video { get; set; }

        public MediaTranscodeProTemplateArgs()
        {
        }
    }

    public sealed class MediaTranscodeProTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audio information, do not transmit Audio, which is equivalent to deleting audio information.
        /// </summary>
        [Input("audio")]
        public Input<Inputs.MediaTranscodeProTemplateAudioGetArgs>? Audio { get; set; }

        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// container format.
        /// </summary>
        [Input("container")]
        public Input<Inputs.MediaTranscodeProTemplateContainerGetArgs>? Container { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// time interval.
        /// </summary>
        [Input("timeInterval")]
        public Input<Inputs.MediaTranscodeProTemplateTimeIntervalGetArgs>? TimeInterval { get; set; }

        /// <summary>
        /// transcoding configuration.
        /// </summary>
        [Input("transConfig")]
        public Input<Inputs.MediaTranscodeProTemplateTransConfigGetArgs>? TransConfig { get; set; }

        /// <summary>
        /// video information, do not upload Video, which is equivalent to deleting video information.
        /// </summary>
        [Input("video")]
        public Input<Inputs.MediaTranscodeProTemplateVideoGetArgs>? Video { get; set; }

        public MediaTranscodeProTemplateState()
        {
        }
    }
}
