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
    /// Provides a resource to create a mps content_review_template
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
    ///         var template = new Tencentcloud.Mps.ContentReviewTemplate("template", new Tencentcloud.Mps.ContentReviewTemplateArgs
    ///         {
    ///             Comment = "tf test content review temp",
    ///             PoliticalConfigure = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePoliticalConfigureArgs
    ///             {
    ///                 AsrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePoliticalConfigureAsrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 ImgReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePoliticalConfigureImgReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     LabelSet = 
    ///                     {
    ///                         "violation_photo",
    ///                         "politician",
    ///                     },
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 OcrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePoliticalConfigureOcrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///             },
    ///             PornConfigure = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePornConfigureArgs
    ///             {
    ///                 AsrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePornConfigureAsrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 ImgReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePornConfigureImgReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     LabelSet = 
    ///                     {
    ///                         "porn",
    ///                         "vulgar",
    ///                     },
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 OcrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplatePornConfigureOcrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///             },
    ///             ProhibitedConfigure = new Tencentcloud.Mps.Inputs.ContentReviewTemplateProhibitedConfigureArgs
    ///             {
    ///                 AsrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateProhibitedConfigureAsrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 OcrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateProhibitedConfigureOcrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///             },
    ///             TerrorismConfigure = new Tencentcloud.Mps.Inputs.ContentReviewTemplateTerrorismConfigureArgs
    ///             {
    ///                 ImgReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateTerrorismConfigureImgReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     LabelSet = 
    ///                     {
    ///                         "guns",
    ///                         "crowd",
    ///                     },
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 OcrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateTerrorismConfigureOcrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///             },
    ///             UserDefineConfigure = new Tencentcloud.Mps.Inputs.ContentReviewTemplateUserDefineConfigureArgs
    ///             {
    ///                 AsrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateUserDefineConfigureAsrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     LabelSet = 
    ///                     {
    ///                         "VOICE_1",
    ///                         "VOICE_2",
    ///                     },
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 FaceReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateUserDefineConfigureFaceReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     LabelSet = 
    ///                     {
    ///                         "FACE_1",
    ///                         "FACE_2",
    ///                     },
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
    ///                 },
    ///                 OcrReviewInfo = new Tencentcloud.Mps.Inputs.ContentReviewTemplateUserDefineConfigureOcrReviewInfoArgs
    ///                 {
    ///                     BlockConfidence = 60,
    ///                     LabelSet = 
    ///                     {
    ///                         "VIDEO_1",
    ///                         "VIDEO_2",
    ///                     },
    ///                     ReviewConfidence = 100,
    ///                     Switch = "ON",
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
    /// mps content_review_template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate content_review_template definition
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate")]
    public partial class ContentReviewTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Content review template description information, length limit: 256 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Content review template name, length limit: 64 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Political control parameters.
        /// </summary>
        [Output("politicalConfigure")]
        public Output<Outputs.ContentReviewTemplatePoliticalConfigure?> PoliticalConfigure { get; private set; } = null!;

        /// <summary>
        /// Control parameters for porn image.
        /// </summary>
        [Output("pornConfigure")]
        public Output<Outputs.ContentReviewTemplatePornConfigure?> PornConfigure { get; private set; } = null!;

        /// <summary>
        /// Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        /// </summary>
        [Output("prohibitedConfigure")]
        public Output<Outputs.ContentReviewTemplateProhibitedConfigure?> ProhibitedConfigure { get; private set; } = null!;

        /// <summary>
        /// Control parameters for unsafe information.
        /// </summary>
        [Output("terrorismConfigure")]
        public Output<Outputs.ContentReviewTemplateTerrorismConfigure?> TerrorismConfigure { get; private set; } = null!;

        /// <summary>
        /// User-Defined Content Moderation Control Parameters.
        /// </summary>
        [Output("userDefineConfigure")]
        public Output<Outputs.ContentReviewTemplateUserDefineConfigure?> UserDefineConfigure { get; private set; } = null!;


        /// <summary>
        /// Create a ContentReviewTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContentReviewTemplate(string name, ContentReviewTemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate", name, args ?? new ContentReviewTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContentReviewTemplate(string name, Input<string> id, ContentReviewTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContentReviewTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContentReviewTemplate Get(string name, Input<string> id, ContentReviewTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ContentReviewTemplate(name, id, state, options);
        }
    }

    public sealed class ContentReviewTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content review template description information, length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Content review template name, length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Political control parameters.
        /// </summary>
        [Input("politicalConfigure")]
        public Input<Inputs.ContentReviewTemplatePoliticalConfigureArgs>? PoliticalConfigure { get; set; }

        /// <summary>
        /// Control parameters for porn image.
        /// </summary>
        [Input("pornConfigure")]
        public Input<Inputs.ContentReviewTemplatePornConfigureArgs>? PornConfigure { get; set; }

        /// <summary>
        /// Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        /// </summary>
        [Input("prohibitedConfigure")]
        public Input<Inputs.ContentReviewTemplateProhibitedConfigureArgs>? ProhibitedConfigure { get; set; }

        /// <summary>
        /// Control parameters for unsafe information.
        /// </summary>
        [Input("terrorismConfigure")]
        public Input<Inputs.ContentReviewTemplateTerrorismConfigureArgs>? TerrorismConfigure { get; set; }

        /// <summary>
        /// User-Defined Content Moderation Control Parameters.
        /// </summary>
        [Input("userDefineConfigure")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureArgs>? UserDefineConfigure { get; set; }

        public ContentReviewTemplateArgs()
        {
        }
    }

    public sealed class ContentReviewTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content review template description information, length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Content review template name, length limit: 64 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Political control parameters.
        /// </summary>
        [Input("politicalConfigure")]
        public Input<Inputs.ContentReviewTemplatePoliticalConfigureGetArgs>? PoliticalConfigure { get; set; }

        /// <summary>
        /// Control parameters for porn image.
        /// </summary>
        [Input("pornConfigure")]
        public Input<Inputs.ContentReviewTemplatePornConfigureGetArgs>? PornConfigure { get; set; }

        /// <summary>
        /// Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        /// </summary>
        [Input("prohibitedConfigure")]
        public Input<Inputs.ContentReviewTemplateProhibitedConfigureGetArgs>? ProhibitedConfigure { get; set; }

        /// <summary>
        /// Control parameters for unsafe information.
        /// </summary>
        [Input("terrorismConfigure")]
        public Input<Inputs.ContentReviewTemplateTerrorismConfigureGetArgs>? TerrorismConfigure { get; set; }

        /// <summary>
        /// User-Defined Content Moderation Control Parameters.
        /// </summary>
        [Input("userDefineConfigure")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureGetArgs>? UserDefineConfigure { get; set; }

        public ContentReviewTemplateState()
        {
        }
    }
}
