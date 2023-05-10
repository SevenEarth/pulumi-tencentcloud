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
    /// Provides a resource to create a ci media_watermark_template
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
    ///         var mediaWatermarkTemplate = new Tencentcloud.Ci.MediaWatermarkTemplate("mediaWatermarkTemplate", new Tencentcloud.Ci.MediaWatermarkTemplateArgs
    ///         {
    ///             Bucket = "terraform-ci-1308919341",
    ///             Watermark = new Tencentcloud.Ci.Inputs.MediaWatermarkTemplateWatermarkArgs
    ///             {
    ///                 Dx = "128",
    ///                 Dy = "128",
    ///                 EndTime = "100.5",
    ///                 LocMode = "Absolute",
    ///                 Pos = "TopRight",
    ///                 StartTime = "0",
    ///                 Text = new Tencentcloud.Ci.Inputs.MediaWatermarkTemplateWatermarkTextArgs
    ///                 {
    ///                     FontColor = "0xF0F8F0",
    ///                     FontSize = "30",
    ///                     FontType = "simfang.ttf",
    ///                     Text = "watermark-content",
    ///                     Transparency = "30",
    ///                 },
    ///                 Type = "Text",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ci media_watermark_template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate media_watermark_template media_watermark_template_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate")]
    public partial class MediaWatermarkTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// container format.
        /// </summary>
        [Output("watermark")]
        public Output<Outputs.MediaWatermarkTemplateWatermark> Watermark { get; private set; } = null!;


        /// <summary>
        /// Create a MediaWatermarkTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaWatermarkTemplate(string name, MediaWatermarkTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate", name, args ?? new MediaWatermarkTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaWatermarkTemplate(string name, Input<string> id, MediaWatermarkTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaWatermarkTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaWatermarkTemplate Get(string name, Input<string> id, MediaWatermarkTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaWatermarkTemplate(name, id, state, options);
        }
    }

    public sealed class MediaWatermarkTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// container format.
        /// </summary>
        [Input("watermark", required: true)]
        public Input<Inputs.MediaWatermarkTemplateWatermarkArgs> Watermark { get; set; } = null!;

        public MediaWatermarkTemplateArgs()
        {
        }
    }

    public sealed class MediaWatermarkTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// container format.
        /// </summary>
        [Input("watermark")]
        public Input<Inputs.MediaWatermarkTemplateWatermarkGetArgs>? Watermark { get; set; }

        public MediaWatermarkTemplateState()
        {
        }
    }
}
