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
    /// Provides a resource to create a css callback_template
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
    ///     var callbackTemplate = new Tencentcloud.Css.CallbackTemplate("callbackTemplate", new()
    ///     {
    ///         CallbackKey = "adasda131312",
    ///         Description = "this is demo",
    ///         PornCensorshipNotifyUrl = "http://www.yourdomain.com/api/notify?action=porn",
    ///         PushExceptionNotifyUrl = "http://www.yourdomain.com/api/notify?action=pushException",
    ///         RecordNotifyUrl = "http://www.yourdomain.com/api/notify?action=record",
    ///         SnapshotNotifyUrl = "http://www.yourdomain.com/api/notify?action=snapshot",
    ///         StreamBeginNotifyUrl = "http://www.yourdomain.com/api/notify?action=streamBegin",
    ///         StreamEndNotifyUrl = "http://www.yourdomain.com/api/notify?action=streamEnd",
    ///         TemplateName = "tf-test",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// css callback_template can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Css/callbackTemplate:CallbackTemplate callback_template templateId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Css/callbackTemplate:CallbackTemplate")]
    public partial class CallbackTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Callback Key, public callback URL.
        /// </summary>
        [Output("callbackKey")]
        public Output<string?> CallbackKey { get; private set; } = null!;

        /// <summary>
        /// Description information.Maximum length: 1024 bytes.Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// PornCensorship callback URL.
        /// </summary>
        [Output("pornCensorshipNotifyUrl")]
        public Output<string?> PornCensorshipNotifyUrl { get; private set; } = null!;

        /// <summary>
        /// Streaming Exception Callback URL.
        /// </summary>
        [Output("pushExceptionNotifyUrl")]
        public Output<string?> PushExceptionNotifyUrl { get; private set; } = null!;

        /// <summary>
        /// Recording callback URL.
        /// </summary>
        [Output("recordNotifyUrl")]
        public Output<string?> RecordNotifyUrl { get; private set; } = null!;

        /// <summary>
        /// Snapshot callback URL.
        /// </summary>
        [Output("snapshotNotifyUrl")]
        public Output<string?> SnapshotNotifyUrl { get; private set; } = null!;

        /// <summary>
        /// Launch callback URL.
        /// </summary>
        [Output("streamBeginNotifyUrl")]
        public Output<string?> StreamBeginNotifyUrl { get; private set; } = null!;

        /// <summary>
        /// Cutoff callback URL.
        /// </summary>
        [Output("streamEndNotifyUrl")]
        public Output<string?> StreamEndNotifyUrl { get; private set; } = null!;

        /// <summary>
        /// Template name.Maximum length: 255 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;


        /// <summary>
        /// Create a CallbackTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CallbackTemplate(string name, CallbackTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/callbackTemplate:CallbackTemplate", name, args ?? new CallbackTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CallbackTemplate(string name, Input<string> id, CallbackTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/callbackTemplate:CallbackTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CallbackTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CallbackTemplate Get(string name, Input<string> id, CallbackTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new CallbackTemplate(name, id, state, options);
        }
    }

    public sealed class CallbackTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Callback Key, public callback URL.
        /// </summary>
        [Input("callbackKey")]
        public Input<string>? CallbackKey { get; set; }

        /// <summary>
        /// Description information.Maximum length: 1024 bytes.Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// PornCensorship callback URL.
        /// </summary>
        [Input("pornCensorshipNotifyUrl")]
        public Input<string>? PornCensorshipNotifyUrl { get; set; }

        /// <summary>
        /// Streaming Exception Callback URL.
        /// </summary>
        [Input("pushExceptionNotifyUrl")]
        public Input<string>? PushExceptionNotifyUrl { get; set; }

        /// <summary>
        /// Recording callback URL.
        /// </summary>
        [Input("recordNotifyUrl")]
        public Input<string>? RecordNotifyUrl { get; set; }

        /// <summary>
        /// Snapshot callback URL.
        /// </summary>
        [Input("snapshotNotifyUrl")]
        public Input<string>? SnapshotNotifyUrl { get; set; }

        /// <summary>
        /// Launch callback URL.
        /// </summary>
        [Input("streamBeginNotifyUrl")]
        public Input<string>? StreamBeginNotifyUrl { get; set; }

        /// <summary>
        /// Cutoff callback URL.
        /// </summary>
        [Input("streamEndNotifyUrl")]
        public Input<string>? StreamEndNotifyUrl { get; set; }

        /// <summary>
        /// Template name.Maximum length: 255 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public CallbackTemplateArgs()
        {
        }
        public static new CallbackTemplateArgs Empty => new CallbackTemplateArgs();
    }

    public sealed class CallbackTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Callback Key, public callback URL.
        /// </summary>
        [Input("callbackKey")]
        public Input<string>? CallbackKey { get; set; }

        /// <summary>
        /// Description information.Maximum length: 1024 bytes.Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// PornCensorship callback URL.
        /// </summary>
        [Input("pornCensorshipNotifyUrl")]
        public Input<string>? PornCensorshipNotifyUrl { get; set; }

        /// <summary>
        /// Streaming Exception Callback URL.
        /// </summary>
        [Input("pushExceptionNotifyUrl")]
        public Input<string>? PushExceptionNotifyUrl { get; set; }

        /// <summary>
        /// Recording callback URL.
        /// </summary>
        [Input("recordNotifyUrl")]
        public Input<string>? RecordNotifyUrl { get; set; }

        /// <summary>
        /// Snapshot callback URL.
        /// </summary>
        [Input("snapshotNotifyUrl")]
        public Input<string>? SnapshotNotifyUrl { get; set; }

        /// <summary>
        /// Launch callback URL.
        /// </summary>
        [Input("streamBeginNotifyUrl")]
        public Input<string>? StreamBeginNotifyUrl { get; set; }

        /// <summary>
        /// Cutoff callback URL.
        /// </summary>
        [Input("streamEndNotifyUrl")]
        public Input<string>? StreamEndNotifyUrl { get; set; }

        /// <summary>
        /// Template name.Maximum length: 255 bytes. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        public CallbackTemplateState()
        {
        }
        public static new CallbackTemplateState Empty => new CallbackTemplateState();
    }
}
