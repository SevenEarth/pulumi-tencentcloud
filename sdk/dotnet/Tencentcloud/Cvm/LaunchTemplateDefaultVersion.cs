// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm
{
    /// <summary>
    /// Provides a resource to create a cvm launch_template_default_version
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
    ///     var launchTemplateDefaultVersion = new Tencentcloud.Cvm.LaunchTemplateDefaultVersion("launchTemplateDefaultVersion", new()
    ///     {
    ///         DefaultVersion = 2,
    ///         LaunchTemplateId = "lt-34vaef8fe",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// cvm launch_template_default_version can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion launch_template_default_version launch_template_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion")]
    public partial class LaunchTemplateDefaultVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of the version that you want to set as the default version.
        /// </summary>
        [Output("defaultVersion")]
        public Output<int> DefaultVersion { get; private set; } = null!;

        /// <summary>
        /// Instance launch template ID.
        /// </summary>
        [Output("launchTemplateId")]
        public Output<string> LaunchTemplateId { get; private set; } = null!;


        /// <summary>
        /// Create a LaunchTemplateDefaultVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LaunchTemplateDefaultVersion(string name, LaunchTemplateDefaultVersionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion", name, args ?? new LaunchTemplateDefaultVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LaunchTemplateDefaultVersion(string name, Input<string> id, LaunchTemplateDefaultVersionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LaunchTemplateDefaultVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LaunchTemplateDefaultVersion Get(string name, Input<string> id, LaunchTemplateDefaultVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new LaunchTemplateDefaultVersion(name, id, state, options);
        }
    }

    public sealed class LaunchTemplateDefaultVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of the version that you want to set as the default version.
        /// </summary>
        [Input("defaultVersion", required: true)]
        public Input<int> DefaultVersion { get; set; } = null!;

        /// <summary>
        /// Instance launch template ID.
        /// </summary>
        [Input("launchTemplateId", required: true)]
        public Input<string> LaunchTemplateId { get; set; } = null!;

        public LaunchTemplateDefaultVersionArgs()
        {
        }
        public static new LaunchTemplateDefaultVersionArgs Empty => new LaunchTemplateDefaultVersionArgs();
    }

    public sealed class LaunchTemplateDefaultVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of the version that you want to set as the default version.
        /// </summary>
        [Input("defaultVersion")]
        public Input<int>? DefaultVersion { get; set; }

        /// <summary>
        /// Instance launch template ID.
        /// </summary>
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        public LaunchTemplateDefaultVersionState()
        {
        }
        public static new LaunchTemplateDefaultVersionState Empty => new LaunchTemplateDefaultVersionState();
    }
}
