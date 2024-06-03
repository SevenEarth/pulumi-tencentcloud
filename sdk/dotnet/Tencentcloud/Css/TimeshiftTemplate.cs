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
    /// Provides a resource to create a css timeshift_template
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
    ///     var timeshiftTemplate = new Tencentcloud.Css.TimeshiftTemplate("timeshiftTemplate", new()
    ///     {
    ///         Area = "Mainland",
    ///         Description = "timeshift template",
    ///         Duration = 604800,
    ///         ItemDuration = 5,
    ///         RemoveWatermark = true,
    ///         TemplateName = "tf-test",
    ///         TranscodeTemplateIds = new[] {},
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// css timeshift_template can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Css/timeshiftTemplate:TimeshiftTemplate timeshift_template templateId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Css/timeshiftTemplate:TimeshiftTemplate")]
    public partial class TimeshiftTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The region.`Mainland`: The Chinese mainland.`Overseas`: Outside the Chinese mainland.Default value: `Mainland`.
        /// </summary>
        [Output("area")]
        public Output<string?> Area { get; private set; } = null!;

        /// <summary>
        /// The template description.Only letters, numbers, underscores, and hyphens are supported.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The time shifting duration.Unit: Second.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// The segment size.Value range: 3-10.Unit: Second.Default value: 5.
        /// </summary>
        [Output("itemDuration")]
        public Output<int?> ItemDuration { get; private set; } = null!;

        /// <summary>
        /// Whether to remove watermarks.If you pass in `true`, the original stream will be recorded.Default value: `false`.
        /// </summary>
        [Output("removeWatermark")]
        public Output<bool?> RemoveWatermark { get; private set; } = null!;

        /// <summary>
        /// The template name.Maximum length: 255 bytes.Only letters, numbers, underscores, and hyphens are supported.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// The transcoding template IDs.This API works only if `RemoveWatermark` is `false`.
        /// </summary>
        [Output("transcodeTemplateIds")]
        public Output<ImmutableArray<int>> TranscodeTemplateIds { get; private set; } = null!;


        /// <summary>
        /// Create a TimeshiftTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TimeshiftTemplate(string name, TimeshiftTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/timeshiftTemplate:TimeshiftTemplate", name, args ?? new TimeshiftTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TimeshiftTemplate(string name, Input<string> id, TimeshiftTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/timeshiftTemplate:TimeshiftTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TimeshiftTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TimeshiftTemplate Get(string name, Input<string> id, TimeshiftTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new TimeshiftTemplate(name, id, state, options);
        }
    }

    public sealed class TimeshiftTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region.`Mainland`: The Chinese mainland.`Overseas`: Outside the Chinese mainland.Default value: `Mainland`.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// The template description.Only letters, numbers, underscores, and hyphens are supported.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The time shifting duration.Unit: Second.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// The segment size.Value range: 3-10.Unit: Second.Default value: 5.
        /// </summary>
        [Input("itemDuration")]
        public Input<int>? ItemDuration { get; set; }

        /// <summary>
        /// Whether to remove watermarks.If you pass in `true`, the original stream will be recorded.Default value: `false`.
        /// </summary>
        [Input("removeWatermark")]
        public Input<bool>? RemoveWatermark { get; set; }

        /// <summary>
        /// The template name.Maximum length: 255 bytes.Only letters, numbers, underscores, and hyphens are supported.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        [Input("transcodeTemplateIds")]
        private InputList<int>? _transcodeTemplateIds;

        /// <summary>
        /// The transcoding template IDs.This API works only if `RemoveWatermark` is `false`.
        /// </summary>
        public InputList<int> TranscodeTemplateIds
        {
            get => _transcodeTemplateIds ?? (_transcodeTemplateIds = new InputList<int>());
            set => _transcodeTemplateIds = value;
        }

        public TimeshiftTemplateArgs()
        {
        }
        public static new TimeshiftTemplateArgs Empty => new TimeshiftTemplateArgs();
    }

    public sealed class TimeshiftTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region.`Mainland`: The Chinese mainland.`Overseas`: Outside the Chinese mainland.Default value: `Mainland`.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// The template description.Only letters, numbers, underscores, and hyphens are supported.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The time shifting duration.Unit: Second.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// The segment size.Value range: 3-10.Unit: Second.Default value: 5.
        /// </summary>
        [Input("itemDuration")]
        public Input<int>? ItemDuration { get; set; }

        /// <summary>
        /// Whether to remove watermarks.If you pass in `true`, the original stream will be recorded.Default value: `false`.
        /// </summary>
        [Input("removeWatermark")]
        public Input<bool>? RemoveWatermark { get; set; }

        /// <summary>
        /// The template name.Maximum length: 255 bytes.Only letters, numbers, underscores, and hyphens are supported.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("transcodeTemplateIds")]
        private InputList<int>? _transcodeTemplateIds;

        /// <summary>
        /// The transcoding template IDs.This API works only if `RemoveWatermark` is `false`.
        /// </summary>
        public InputList<int> TranscodeTemplateIds
        {
            get => _transcodeTemplateIds ?? (_transcodeTemplateIds = new InputList<int>());
            set => _transcodeTemplateIds = value;
        }

        public TimeshiftTemplateState()
        {
        }
        public static new TimeshiftTemplateState Empty => new TimeshiftTemplateState();
    }
}
