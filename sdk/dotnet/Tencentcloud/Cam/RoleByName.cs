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
    /// Provides a resource to create a CAM role.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create normally
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
    ///     var foo = new Tencentcloud.Cam.RoleByName("foo", new()
    ///     {
    ///         ConsoleLogin = true,
    ///         Description = "test",
    ///         Document = @"{
    ///   ""version"": ""2.0"",
    ///   ""statement"": [
    ///     {
    ///       ""action"": [""name/sts:AssumeRole""],
    ///       ""effect"": ""allow"",
    ///       ""principal"": {
    ///         ""qcs"": [""qcs::cam::uin/&lt;your-account-id&gt;:uin/&lt;your-account-id&gt;""]
    ///       }
    ///     }
    ///   ]
    /// }
    /// 
    /// ",
    ///         Tags = 
    ///         {
    ///             { "test", "tf-cam-role" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Create with SAML provider
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
    ///     var boo = new Tencentcloud.Cam.RoleByName("boo", new()
    ///     {
    ///         ConsoleLogin = true,
    ///         Description = "test",
    ///         Document = @"{
    ///   ""version"": ""2.0"",
    ///   ""statement"": [
    ///     {
    ///       ""action"": [""name/sts:AssumeRole"", ""name/sts:AssumeRoleWithWebIdentity""],
    ///       ""effect"": ""allow"",
    ///       ""principal"": {
    ///         ""federated"": [""qcs::cam::uin/&lt;your-account-id&gt;:saml-provider/&lt;your-name&gt;""]
    ///       }
    ///     }
    ///   ]
    /// }
    /// 
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// CAM role can be imported using the name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cam/roleByName:RoleByName foo cam-role-test
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cam/roleByName:RoleByName")]
    public partial class RoleByName : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the CAM role can login or not.
        /// </summary>
        [Output("consoleLogin")]
        public Output<bool?> ConsoleLogin { get; private set; } = null!;

        /// <summary>
        /// Create time of the CAM role.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the CAM role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Document of the CAM role. The syntax refers to CAM POLICY Name of CAM role.
        /// </summary>
        [Output("document")]
        public Output<string> Document { get; private set; } = null!;

        /// <summary>
        /// Name of CAM role.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The last update time of the CAM role.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a RoleByName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleByName(string name, RoleByNameArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/roleByName:RoleByName", name, args ?? new RoleByNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoleByName(string name, Input<string> id, RoleByNameState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/roleByName:RoleByName", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoleByName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleByName Get(string name, Input<string> id, RoleByNameState? state = null, CustomResourceOptions? options = null)
        {
            return new RoleByName(name, id, state, options);
        }
    }

    public sealed class RoleByNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the CAM role can login or not.
        /// </summary>
        [Input("consoleLogin")]
        public Input<bool>? ConsoleLogin { get; set; }

        /// <summary>
        /// Description of the CAM role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Document of the CAM role. The syntax refers to CAM POLICY Name of CAM role.
        /// </summary>
        [Input("document", required: true)]
        public Input<string> Document { get; set; } = null!;

        /// <summary>
        /// Name of CAM role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public RoleByNameArgs()
        {
        }
        public static new RoleByNameArgs Empty => new RoleByNameArgs();
    }

    public sealed class RoleByNameState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the CAM role can login or not.
        /// </summary>
        [Input("consoleLogin")]
        public Input<bool>? ConsoleLogin { get; set; }

        /// <summary>
        /// Create time of the CAM role.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the CAM role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Document of the CAM role. The syntax refers to CAM POLICY Name of CAM role.
        /// </summary>
        [Input("document")]
        public Input<string>? Document { get; set; }

        /// <summary>
        /// Name of CAM role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The last update time of the CAM role.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public RoleByNameState()
        {
        }
        public static new RoleByNameState Empty => new RoleByNameState();
    }
}
