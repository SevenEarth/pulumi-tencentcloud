// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse
{
    /// <summary>
    /// Provides a resource to create a lighthouse key_pair_attachment
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
    ///     var keyPairAttachment = new Tencentcloud.Lighthouse.KeyPairAttachment("keyPairAttachment", new()
    ///     {
    ///         InstanceId = "lhins-xxxxxx",
    ///         KeyId = "lhkp-xxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// lighthouse key_pair_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Lighthouse/keyPairAttachment:KeyPairAttachment key_pair_attachment key_pair_attachment_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Lighthouse/keyPairAttachment:KeyPairAttachment")]
    public partial class KeyPairAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Key pair ID.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;


        /// <summary>
        /// Create a KeyPairAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyPairAttachment(string name, KeyPairAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Lighthouse/keyPairAttachment:KeyPairAttachment", name, args ?? new KeyPairAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyPairAttachment(string name, Input<string> id, KeyPairAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Lighthouse/keyPairAttachment:KeyPairAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyPairAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyPairAttachment Get(string name, Input<string> id, KeyPairAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyPairAttachment(name, id, state, options);
        }
    }

    public sealed class KeyPairAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Key pair ID.
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public KeyPairAttachmentArgs()
        {
        }
        public static new KeyPairAttachmentArgs Empty => new KeyPairAttachmentArgs();
    }

    public sealed class KeyPairAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Key pair ID.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        public KeyPairAttachmentState()
        {
        }
        public static new KeyPairAttachmentState Empty => new KeyPairAttachmentState();
    }
}
