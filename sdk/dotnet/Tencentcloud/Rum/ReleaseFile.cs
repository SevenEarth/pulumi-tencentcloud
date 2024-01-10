// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Rum
{
    [TencentcloudResourceType("tencentcloud:Rum/releaseFile:ReleaseFile")]
    public partial class ReleaseFile : Pulumi.CustomResource
    {
        /// <summary>
        /// Release file hash.
        /// </summary>
        [Output("fileHash")]
        public Output<string> FileHash { get; private set; } = null!;

        /// <summary>
        /// Release file unique key.
        /// </summary>
        [Output("fileKey")]
        public Output<string> FileKey { get; private set; } = null!;

        /// <summary>
        /// Release file name.
        /// </summary>
        [Output("fileName")]
        public Output<string> FileName { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Release file id.
        /// </summary>
        [Output("releaseFileId")]
        public Output<int> ReleaseFileId { get; private set; } = null!;

        /// <summary>
        /// Release File version.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ReleaseFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReleaseFile(string name, ReleaseFileArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Rum/releaseFile:ReleaseFile", name, args ?? new ReleaseFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReleaseFile(string name, Input<string> id, ReleaseFileState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Rum/releaseFile:ReleaseFile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReleaseFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReleaseFile Get(string name, Input<string> id, ReleaseFileState? state = null, CustomResourceOptions? options = null)
        {
            return new ReleaseFile(name, id, state, options);
        }
    }

    public sealed class ReleaseFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Release file hash.
        /// </summary>
        [Input("fileHash", required: true)]
        public Input<string> FileHash { get; set; } = null!;

        /// <summary>
        /// Release file unique key.
        /// </summary>
        [Input("fileKey", required: true)]
        public Input<string> FileKey { get; set; } = null!;

        /// <summary>
        /// Release file name.
        /// </summary>
        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Release file id.
        /// </summary>
        [Input("releaseFileId", required: true)]
        public Input<int> ReleaseFileId { get; set; } = null!;

        /// <summary>
        /// Release File version.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ReleaseFileArgs()
        {
        }
    }

    public sealed class ReleaseFileState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Release file hash.
        /// </summary>
        [Input("fileHash")]
        public Input<string>? FileHash { get; set; }

        /// <summary>
        /// Release file unique key.
        /// </summary>
        [Input("fileKey")]
        public Input<string>? FileKey { get; set; }

        /// <summary>
        /// Release file name.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Release file id.
        /// </summary>
        [Input("releaseFileId")]
        public Input<int>? ReleaseFileId { get; set; }

        /// <summary>
        /// Release File version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ReleaseFileState()
        {
        }
    }
}
