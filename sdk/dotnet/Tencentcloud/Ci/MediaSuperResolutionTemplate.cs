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
    /// Provides a resource to create a ci media_super_resolution_template
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
    ///     var mediaSuperResolutionTemplate = new Tencentcloud.Ci.MediaSuperResolutionTemplate("mediaSuperResolutionTemplate", new()
    ///     {
    ///         Bucket = "terraform-ci-1308919341",
    ///         EnableScaleUp = "true",
    ///         Resolution = "sdtohd",
    ///         Version = "Enhance",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ci media_super_resolution_template can be imported using the bucket#templateId, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate media_super_resolution_template terraform-ci-xxxxxx#t1d707eb2be3294e22b47123894f85cb8f
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate")]
    public partial class MediaSuperResolutionTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Auto scaling switch, off by default.
        /// </summary>
        [Output("enableScaleUp")]
        public Output<string?> EnableScaleUp { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resolution Options sdtohd: Standard Definition to Ultra Definition, hdto4k: HD to 4K.
        /// </summary>
        [Output("resolution")]
        public Output<string> Resolution { get; private set; } = null!;

        /// <summary>
        /// version, default value Base, Base: basic version, Enhance: enhanced version.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a MediaSuperResolutionTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaSuperResolutionTemplate(string name, MediaSuperResolutionTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate", name, args ?? new MediaSuperResolutionTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaSuperResolutionTemplate(string name, Input<string> id, MediaSuperResolutionTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaSuperResolutionTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaSuperResolutionTemplate Get(string name, Input<string> id, MediaSuperResolutionTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaSuperResolutionTemplate(name, id, state, options);
        }
    }

    public sealed class MediaSuperResolutionTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Auto scaling switch, off by default.
        /// </summary>
        [Input("enableScaleUp")]
        public Input<string>? EnableScaleUp { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resolution Options sdtohd: Standard Definition to Ultra Definition, hdto4k: HD to 4K.
        /// </summary>
        [Input("resolution", required: true)]
        public Input<string> Resolution { get; set; } = null!;

        /// <summary>
        /// version, default value Base, Base: basic version, Enhance: enhanced version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public MediaSuperResolutionTemplateArgs()
        {
        }
        public static new MediaSuperResolutionTemplateArgs Empty => new MediaSuperResolutionTemplateArgs();
    }

    public sealed class MediaSuperResolutionTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Auto scaling switch, off by default.
        /// </summary>
        [Input("enableScaleUp")]
        public Input<string>? EnableScaleUp { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resolution Options sdtohd: Standard Definition to Ultra Definition, hdto4k: HD to 4K.
        /// </summary>
        [Input("resolution")]
        public Input<string>? Resolution { get; set; }

        /// <summary>
        /// version, default value Base, Base: basic version, Enhance: enhanced version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public MediaSuperResolutionTemplateState()
        {
        }
        public static new MediaSuperResolutionTemplateState Empty => new MediaSuperResolutionTemplateState();
    }
}
