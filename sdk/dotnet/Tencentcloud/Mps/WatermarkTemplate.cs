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
    /// Provides a resource to create a mps watermark_template
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System;
    /// using System.IO;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    /// 	private static string ReadFileBase64(string path) {
    /// 		return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)))
    /// 	}
    /// 
    ///     public MyStack()
    ///     {
    ///         var watermarkTemplate = new Tencentcloud.Mps.WatermarkTemplate("watermarkTemplate", new Tencentcloud.Mps.WatermarkTemplateArgs
    ///         {
    ///             CoordinateOrigin = "TopLeft",
    ///             Type = "image",
    ///             XPos = "12%",
    ///             YPos = "21%",
    ///             ImageTemplate = new Tencentcloud.Mps.Inputs.WatermarkTemplateImageTemplateArgs
    ///             {
    ///                 Height = "17px",
    ///                 ImageContent = ReadFileBase64("./logo.png"),
    ///                 RepeatType = "repeat",
    ///                 Width = "12px",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mps watermark_template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mps/watermarkTemplate:WatermarkTemplate watermark_template watermark_template_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mps/watermarkTemplate:WatermarkTemplate")]
    public partial class WatermarkTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Template description information, length limit: 256 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
        /// </summary>
        [Output("coordinateOrigin")]
        public Output<string?> CoordinateOrigin { get; private set; } = null!;

        /// <summary>
        /// Image watermark template, only when Type is image, this field is required and valid.
        /// </summary>
        [Output("imageTemplate")]
        public Output<Outputs.WatermarkTemplateImageTemplate?> ImageTemplate { get; private set; } = null!;

        /// <summary>
        /// Watermark template name, length limit: 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SVG watermark template, only when Type is svg, this field is required and valid.
        /// </summary>
        [Output("svgTemplate")]
        public Output<Outputs.WatermarkTemplateSvgTemplate?> SvgTemplate { get; private set; } = null!;

        /// <summary>
        /// Text watermark template, only when Type is text, this field is required and valid.
        /// </summary>
        [Output("textTemplate")]
        public Output<Outputs.WatermarkTemplateTextTemplate?> TextTemplate { get; private set; } = null!;

        /// <summary>
        /// Watermark type, optional value:image, text, svg.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
        /// </summary>
        [Output("xPos")]
        public Output<string?> XPos { get; private set; } = null!;

        /// <summary>
        /// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
        /// </summary>
        [Output("yPos")]
        public Output<string?> YPos { get; private set; } = null!;


        /// <summary>
        /// Create a WatermarkTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WatermarkTemplate(string name, WatermarkTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/watermarkTemplate:WatermarkTemplate", name, args ?? new WatermarkTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WatermarkTemplate(string name, Input<string> id, WatermarkTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/watermarkTemplate:WatermarkTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WatermarkTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WatermarkTemplate Get(string name, Input<string> id, WatermarkTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new WatermarkTemplate(name, id, state, options);
        }
    }

    public sealed class WatermarkTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template description information, length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
        /// </summary>
        [Input("coordinateOrigin")]
        public Input<string>? CoordinateOrigin { get; set; }

        /// <summary>
        /// Image watermark template, only when Type is image, this field is required and valid.
        /// </summary>
        [Input("imageTemplate")]
        public Input<Inputs.WatermarkTemplateImageTemplateArgs>? ImageTemplate { get; set; }

        /// <summary>
        /// Watermark template name, length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SVG watermark template, only when Type is svg, this field is required and valid.
        /// </summary>
        [Input("svgTemplate")]
        public Input<Inputs.WatermarkTemplateSvgTemplateArgs>? SvgTemplate { get; set; }

        /// <summary>
        /// Text watermark template, only when Type is text, this field is required and valid.
        /// </summary>
        [Input("textTemplate")]
        public Input<Inputs.WatermarkTemplateTextTemplateArgs>? TextTemplate { get; set; }

        /// <summary>
        /// Watermark type, optional value:image, text, svg.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
        /// </summary>
        [Input("xPos")]
        public Input<string>? XPos { get; set; }

        /// <summary>
        /// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
        /// </summary>
        [Input("yPos")]
        public Input<string>? YPos { get; set; }

        public WatermarkTemplateArgs()
        {
        }
    }

    public sealed class WatermarkTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template description information, length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Origin position, optional value:TopLeft: Indicates that the origin of the coordinates is at the upper left corner of the video image, and the origin of the watermark is the upper left corner of the picture or text.TopRight: Indicates that the origin of the coordinates is at the upper right corner of the video image, and the origin of the watermark is at the upper right corner of the picture or text.BottomLeft: Indicates that the origin of the coordinates is at the lower left corner of the video image, and the origin of the watermark is the lower left corner of the picture or text.BottomRight: Indicates that the origin of the coordinates is at the lower right corner of the video image, and the origin of the watermark is at the lower right corner of the picture or text.Default value: TopLeft.
        /// </summary>
        [Input("coordinateOrigin")]
        public Input<string>? CoordinateOrigin { get; set; }

        /// <summary>
        /// Image watermark template, only when Type is image, this field is required and valid.
        /// </summary>
        [Input("imageTemplate")]
        public Input<Inputs.WatermarkTemplateImageTemplateGetArgs>? ImageTemplate { get; set; }

        /// <summary>
        /// Watermark template name, length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SVG watermark template, only when Type is svg, this field is required and valid.
        /// </summary>
        [Input("svgTemplate")]
        public Input<Inputs.WatermarkTemplateSvgTemplateGetArgs>? SvgTemplate { get; set; }

        /// <summary>
        /// Text watermark template, only when Type is text, this field is required and valid.
        /// </summary>
        [Input("textTemplate")]
        public Input<Inputs.WatermarkTemplateTextTemplateGetArgs>? TextTemplate { get; set; }

        /// <summary>
        /// Watermark type, optional value:image, text, svg.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The horizontal position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark XPos specifies a percentage for the video width, such as 10% means that XPos is 10% of the video width.When the string ends with px, it means that the watermark XPos is the specified pixel, such as 100px means that the XPos is 100 pixels.Default value: 0px.
        /// </summary>
        [Input("xPos")]
        public Input<string>? XPos { get; set; }

        /// <summary>
        /// The vertical position of the origin of the watermark from the origin of the coordinates of the video image. Support %, px two formats.When the string ends with %, it means that the watermark YPos specifies a percentage for the video height, such as 10% means that YPos is 10% of the video height.When the string ends with px, it means that the watermark YPos is the specified pixel, such as 100px means that the YPos is 100 pixels.Default value: 0px.
        /// </summary>
        [Input("yPos")]
        public Input<string>? YPos { get; set; }

        public WatermarkTemplateState()
        {
        }
    }
}
