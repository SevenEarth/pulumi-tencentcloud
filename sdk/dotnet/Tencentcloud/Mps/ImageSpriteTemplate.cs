// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps
{
    /// <summary>
    /// Provides a resource to create a mps image_sprite_template
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
    ///     var imageSpriteTemplate = new Tencentcloud.Mps.ImageSpriteTemplate("imageSpriteTemplate", new()
    ///     {
    ///         ColumnCount = 10,
    ///         FillType = "stretch",
    ///         Format = "jpg",
    ///         Height = 143,
    ///         ResolutionAdaptive = "open",
    ///         RowCount = 10,
    ///         SampleInterval = 10,
    ///         SampleType = "Time",
    ///         Width = 182,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// mps image_sprite_template can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate image_sprite_template image_sprite_template_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate")]
    public partial class ImageSpriteTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of columns in the small image in the sprite.
        /// </summary>
        [Output("columnCount")]
        public Output<int> ColumnCount { get; private set; } = null!;

        /// <summary>
        /// Template description information, length limit: 256 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
        /// </summary>
        [Output("fillType")]
        public Output<string?> FillType { get; private set; } = null!;

        /// <summary>
        /// Image format, the value can be jpg, png, webp. Default is jpg.
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        /// </summary>
        [Output("height")]
        public Output<int?> Height { get; private set; } = null!;

        /// <summary>
        /// Image sprite template name, length limit: 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        /// </summary>
        [Output("resolutionAdaptive")]
        public Output<string?> ResolutionAdaptive { get; private set; } = null!;

        /// <summary>
        /// The number of rows in the small image in the sprite.
        /// </summary>
        [Output("rowCount")]
        public Output<int> RowCount { get; private set; } = null!;

        /// <summary>
        /// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
        /// </summary>
        [Output("sampleInterval")]
        public Output<int> SampleInterval { get; private set; } = null!;

        /// <summary>
        /// Sampling type, optional value:Percent/Time.
        /// </summary>
        [Output("sampleType")]
        public Output<string> SampleType { get; private set; } = null!;

        /// <summary>
        /// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        /// </summary>
        [Output("width")]
        public Output<int?> Width { get; private set; } = null!;


        /// <summary>
        /// Create a ImageSpriteTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageSpriteTemplate(string name, ImageSpriteTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate", name, args ?? new ImageSpriteTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageSpriteTemplate(string name, Input<string> id, ImageSpriteTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/imageSpriteTemplate:ImageSpriteTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImageSpriteTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageSpriteTemplate Get(string name, Input<string> id, ImageSpriteTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageSpriteTemplate(name, id, state, options);
        }
    }

    public sealed class ImageSpriteTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of columns in the small image in the sprite.
        /// </summary>
        [Input("columnCount", required: true)]
        public Input<int> ColumnCount { get; set; } = null!;

        /// <summary>
        /// Template description information, length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
        /// </summary>
        [Input("fillType")]
        public Input<string>? FillType { get; set; }

        /// <summary>
        /// Image format, the value can be jpg, png, webp. Default is jpg.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Image sprite template name, length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        /// </summary>
        [Input("resolutionAdaptive")]
        public Input<string>? ResolutionAdaptive { get; set; }

        /// <summary>
        /// The number of rows in the small image in the sprite.
        /// </summary>
        [Input("rowCount", required: true)]
        public Input<int> RowCount { get; set; } = null!;

        /// <summary>
        /// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
        /// </summary>
        [Input("sampleInterval", required: true)]
        public Input<int> SampleInterval { get; set; } = null!;

        /// <summary>
        /// Sampling type, optional value:Percent/Time.
        /// </summary>
        [Input("sampleType", required: true)]
        public Input<string> SampleType { get; set; } = null!;

        /// <summary>
        /// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public ImageSpriteTemplateArgs()
        {
        }
        public static new ImageSpriteTemplateArgs Empty => new ImageSpriteTemplateArgs();
    }

    public sealed class ImageSpriteTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of columns in the small image in the sprite.
        /// </summary>
        [Input("columnCount")]
        public Input<int>? ColumnCount { get; set; }

        /// <summary>
        /// Template description information, length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.Default value: black.
        /// </summary>
        [Input("fillType")]
        public Input<string>? FillType { get; set; }

        /// <summary>
        /// Image format, the value can be jpg, png, webp. Default is jpg.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// The maximum value of the height (or short side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Image sprite template name, length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        /// </summary>
        [Input("resolutionAdaptive")]
        public Input<string>? ResolutionAdaptive { get; set; }

        /// <summary>
        /// The number of rows in the small image in the sprite.
        /// </summary>
        [Input("rowCount")]
        public Input<int>? RowCount { get; set; }

        /// <summary>
        /// Sampling interval.When SampleType is Percent, specify the percentage of the sampling interval.When SampleType is Time, specify the sampling interval time in seconds.
        /// </summary>
        [Input("sampleInterval")]
        public Input<int>? SampleInterval { get; set; }

        /// <summary>
        /// Sampling type, optional value:Percent/Time.
        /// </summary>
        [Input("sampleType")]
        public Input<string>? SampleType { get; set; }

        /// <summary>
        /// The maximum value of the width (or long side) of the small image in the sprite image, value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public ImageSpriteTemplateState()
        {
        }
        public static new ImageSpriteTemplateState Empty => new ImageSpriteTemplateState();
    }
}
