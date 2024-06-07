// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    /// <summary>
    /// Provides a resource to create a tsf application_public_config_release
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
    ///     var applicationPublicConfigRelease = new Tencentcloud.Tsf.ApplicationPublicConfigRelease("applicationPublicConfigRelease", new()
    ///     {
    ///         ConfigId = "dcfg-p-123456",
    ///         NamespaceId = "namespace-123456",
    ///         ReleaseDesc = "product version",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// tsf application_public_config_release can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease application_public_config_release application_public_config_attachment_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease")]
    public partial class ApplicationPublicConfigRelease : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ConfigId.
        /// </summary>
        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        /// <summary>
        /// namespace-id.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Release description.
        /// </summary>
        [Output("releaseDesc")]
        public Output<string?> ReleaseDesc { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationPublicConfigRelease resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationPublicConfigRelease(string name, ApplicationPublicConfigReleaseArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease", name, args ?? new ApplicationPublicConfigReleaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationPublicConfigRelease(string name, Input<string> id, ApplicationPublicConfigReleaseState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationPublicConfigRelease resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationPublicConfigRelease Get(string name, Input<string> id, ApplicationPublicConfigReleaseState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationPublicConfigRelease(name, id, state, options);
        }
    }

    public sealed class ApplicationPublicConfigReleaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ConfigId.
        /// </summary>
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// namespace-id.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// Release description.
        /// </summary>
        [Input("releaseDesc")]
        public Input<string>? ReleaseDesc { get; set; }

        public ApplicationPublicConfigReleaseArgs()
        {
        }
        public static new ApplicationPublicConfigReleaseArgs Empty => new ApplicationPublicConfigReleaseArgs();
    }

    public sealed class ApplicationPublicConfigReleaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ConfigId.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// namespace-id.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// Release description.
        /// </summary>
        [Input("releaseDesc")]
        public Input<string>? ReleaseDesc { get; set; }

        public ApplicationPublicConfigReleaseState()
        {
        }
        public static new ApplicationPublicConfigReleaseState Empty => new ApplicationPublicConfigReleaseState();
    }
}
