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
    /// Provides a resource to create a tsf application_file_config_release
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var applicationFileConfigRelease = new Tencentcloud.Tsf.ApplicationFileConfigRelease("applicationFileConfigRelease", new Tencentcloud.Tsf.ApplicationFileConfigReleaseArgs
    ///         {
    ///             ConfigId = "dcfg-f-123456",
    ///             GroupId = "group-123456",
    ///             ReleaseDesc = "product release",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tsf applicationfile_config_release can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease application_file_config_release application_file_config_release_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease")]
    public partial class ApplicationFileConfigRelease : Pulumi.CustomResource
    {
        /// <summary>
        /// File config id.
        /// </summary>
        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Group Id.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// release Description.
        /// </summary>
        [Output("releaseDesc")]
        public Output<string?> ReleaseDesc { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationFileConfigRelease resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationFileConfigRelease(string name, ApplicationFileConfigReleaseArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease", name, args ?? new ApplicationFileConfigReleaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationFileConfigRelease(string name, Input<string> id, ApplicationFileConfigReleaseState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationFileConfigRelease resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationFileConfigRelease Get(string name, Input<string> id, ApplicationFileConfigReleaseState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationFileConfigRelease(name, id, state, options);
        }
    }

    public sealed class ApplicationFileConfigReleaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File config id.
        /// </summary>
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// Group Id.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// release Description.
        /// </summary>
        [Input("releaseDesc")]
        public Input<string>? ReleaseDesc { get; set; }

        public ApplicationFileConfigReleaseArgs()
        {
        }
    }

    public sealed class ApplicationFileConfigReleaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File config id.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// Group Id.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// release Description.
        /// </summary>
        [Input("releaseDesc")]
        public Input<string>? ReleaseDesc { get; set; }

        public ApplicationFileConfigReleaseState()
        {
        }
    }
}
