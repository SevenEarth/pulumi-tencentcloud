// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cam
{
    /// <summary>
    /// Provides a resource to create a CAM-ROLE-SSO (Only support OIDC).
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
    ///     var foo = new Tencentcloud.Cam.RoleSso("foo", new()
    ///     {
    ///         ClientIds = new[]
    ///         {
    ///             "...",
    ///         },
    ///         Description = "this is a description",
    ///         IdentityKey = "...",
    ///         IdentityUrl = "https://login.microsoftonline.com/.../v2.0",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// CAM-ROLE-SSO can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cam/roleSso:RoleSso foo "test"
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cam/roleSso:RoleSso")]
    public partial class RoleSso : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Client ids.
        /// </summary>
        [Output("clientIds")]
        public Output<ImmutableArray<string>> ClientIds { get; private set; } = null!;

        /// <summary>
        /// The description of resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Sign the public key.
        /// </summary>
        [Output("identityKey")]
        public Output<string> IdentityKey { get; private set; } = null!;

        /// <summary>
        /// Identity provider URL.
        /// </summary>
        [Output("identityUrl")]
        public Output<string> IdentityUrl { get; private set; } = null!;

        /// <summary>
        /// The name of resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a RoleSso resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleSso(string name, RoleSsoArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/roleSso:RoleSso", name, args ?? new RoleSsoArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoleSso(string name, Input<string> id, RoleSsoState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/roleSso:RoleSso", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoleSso resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleSso Get(string name, Input<string> id, RoleSsoState? state = null, CustomResourceOptions? options = null)
        {
            return new RoleSso(name, id, state, options);
        }
    }

    public sealed class RoleSsoArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientIds", required: true)]
        private InputList<string>? _clientIds;

        /// <summary>
        /// Client ids.
        /// </summary>
        public InputList<string> ClientIds
        {
            get => _clientIds ?? (_clientIds = new InputList<string>());
            set => _clientIds = value;
        }

        /// <summary>
        /// The description of resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sign the public key.
        /// </summary>
        [Input("identityKey", required: true)]
        public Input<string> IdentityKey { get; set; } = null!;

        /// <summary>
        /// Identity provider URL.
        /// </summary>
        [Input("identityUrl", required: true)]
        public Input<string> IdentityUrl { get; set; } = null!;

        /// <summary>
        /// The name of resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RoleSsoArgs()
        {
        }
        public static new RoleSsoArgs Empty => new RoleSsoArgs();
    }

    public sealed class RoleSsoState : global::Pulumi.ResourceArgs
    {
        [Input("clientIds")]
        private InputList<string>? _clientIds;

        /// <summary>
        /// Client ids.
        /// </summary>
        public InputList<string> ClientIds
        {
            get => _clientIds ?? (_clientIds = new InputList<string>());
            set => _clientIds = value;
        }

        /// <summary>
        /// The description of resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sign the public key.
        /// </summary>
        [Input("identityKey")]
        public Input<string>? IdentityKey { get; set; }

        /// <summary>
        /// Identity provider URL.
        /// </summary>
        [Input("identityUrl")]
        public Input<string>? IdentityUrl { get; set; }

        /// <summary>
        /// The name of resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RoleSsoState()
        {
        }
        public static new RoleSsoState Empty => new RoleSsoState();
    }
}
