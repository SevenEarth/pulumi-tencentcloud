// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kms
{
    /// <summary>
    /// Provides a resource to create a kms white_box_key
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
    ///     var example = new Tencentcloud.Kms.WhiteBoxKey("example", new()
    ///     {
    ///         Algorithm = "SM4",
    ///         Alias = "tf_example",
    ///         Description = "test desc.",
    ///         Status = "Enabled",
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// kms white_box_key can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Kms/whiteBoxKey:WhiteBoxKey example 244dab8c-6dad-11ea-80c6-5254006d0810
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kms/whiteBoxKey:WhiteBoxKey")]
    public partial class WhiteBoxKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// All algorithm types for creating keys, supported values: AES_256, SM4.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// As an alias for the key to be easier to identify and easier to understand, it cannot be empty and is a combination of 1-60 alphanumeric characters - _. The first character must be a letter or number. Alias are not repeatable.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Description of the key, up to 1024 bytes.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the key. Enabled or Disabled. Default is Enabled.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of Key.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a WhiteBoxKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WhiteBoxKey(string name, WhiteBoxKeyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kms/whiteBoxKey:WhiteBoxKey", name, args ?? new WhiteBoxKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WhiteBoxKey(string name, Input<string> id, WhiteBoxKeyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kms/whiteBoxKey:WhiteBoxKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WhiteBoxKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WhiteBoxKey Get(string name, Input<string> id, WhiteBoxKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new WhiteBoxKey(name, id, state, options);
        }
    }

    public sealed class WhiteBoxKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// All algorithm types for creating keys, supported values: AES_256, SM4.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// As an alias for the key to be easier to identify and easier to understand, it cannot be empty and is a combination of 1-60 alphanumeric characters - _. The first character must be a letter or number. Alias are not repeatable.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Description of the key, up to 1024 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable the key. Enabled or Disabled. Default is Enabled.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of Key.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public WhiteBoxKeyArgs()
        {
        }
        public static new WhiteBoxKeyArgs Empty => new WhiteBoxKeyArgs();
    }

    public sealed class WhiteBoxKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// All algorithm types for creating keys, supported values: AES_256, SM4.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// As an alias for the key to be easier to identify and easier to understand, it cannot be empty and is a combination of 1-60 alphanumeric characters - _. The first character must be a letter or number. Alias are not repeatable.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Description of the key, up to 1024 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable the key. Enabled or Disabled. Default is Enabled.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of Key.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public WhiteBoxKeyState()
        {
        }
        public static new WhiteBoxKeyState Empty => new WhiteBoxKeyState();
    }
}
