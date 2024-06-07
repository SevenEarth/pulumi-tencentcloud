// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    /// <summary>
    /// Use this resource to create tcr long term token.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a token for tcr instance
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
    ///     var exampleInstance = new Tencentcloud.Tcr.Instance("exampleInstance", new()
    ///     {
    ///         InstanceType = "basic",
    ///         DeleteBucket = true,
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///     });
    /// 
    ///     var exampleToken = new Tencentcloud.Tcr.Token("exampleToken", new()
    ///     {
    ///         InstanceId = exampleInstance.Id,
    ///         Description = "example for the tcr token",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// tcr token can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Tcr/token:Token example instance_id#token_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/token:Token")]
    public partial class Token : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the token. Valid length is [0~255].
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicate to enable this token or not.
        /// </summary>
        [Output("enable")]
        public Output<bool?> Enable { get; private set; } = null!;

        /// <summary>
        /// ID of the TCR instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The content of the token.
        /// </summary>
        [Output("token")]
        public Output<string> TcrToken { get; private set; } = null!;

        /// <summary>
        /// Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
        /// </summary>
        [Output("tokenId")]
        public Output<string> TokenId { get; private set; } = null!;

        /// <summary>
        /// User name of the token.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a Token resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Token(string name, TokenArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/token:Token", name, args ?? new TokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Token(string name, Input<string> id, TokenState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/token:Token", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Token resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Token Get(string name, Input<string> id, TokenState? state = null, CustomResourceOptions? options = null)
        {
            return new Token(name, id, state, options);
        }
    }

    public sealed class TokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the token. Valid length is [0~255].
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicate to enable this token or not.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// ID of the TCR instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public TokenArgs()
        {
        }
        public static new TokenArgs Empty => new TokenArgs();
    }

    public sealed class TokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the token. Valid length is [0~255].
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicate to enable this token or not.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// ID of the TCR instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The content of the token.
        /// </summary>
        [Input("token")]
        public Input<string>? TcrToken { get; set; }

        /// <summary>
        /// Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
        /// </summary>
        [Input("tokenId")]
        public Input<string>? TokenId { get; set; }

        /// <summary>
        /// User name of the token.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public TokenState()
        {
        }
        public static new TokenState Empty => new TokenState();
    }
}
