// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Wedata
{
    /// <summary>
    /// Provides a resource to create a wedata resource
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
    ///         var example = new Tencentcloud.Wedata.Resource("example", new Tencentcloud.Wedata.ResourceArgs
    ///         {
    ///             CosBucketName = "wedata-demo-1314991481",
    ///             CosRegion = "ap-guangzhou",
    ///             FileName = "tf_example",
    ///             FilePath = "/datastudio/resource/demo",
    ///             FilesSize = "8165",
    ///             ProjectId = "1612982498218618880",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// wedata resource can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Wedata/resource:Resource example 1612982498218618880#/datastudio/resource/demo#75431931-7d27-4034-b3de-3dc3348a220e
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Wedata/resource:Resource")]
    public partial class Resource : Pulumi.CustomResource
    {
        /// <summary>
        /// Cos bucket name.
        /// </summary>
        [Output("cosBucketName")]
        public Output<string> CosBucketName { get; private set; } = null!;

        /// <summary>
        /// Cos bucket region.
        /// </summary>
        [Output("cosRegion")]
        public Output<string> CosRegion { get; private set; } = null!;

        /// <summary>
        /// File name.
        /// </summary>
        [Output("fileName")]
        public Output<string> FileName { get; private set; } = null!;

        /// <summary>
        /// For file path:/datastudio/resource/projectId/folderName; for folder path:/datastudio/resource/folderName.
        /// </summary>
        [Output("filePath")]
        public Output<string> FilePath { get; private set; } = null!;

        /// <summary>
        /// File size.
        /// </summary>
        [Output("filesSize")]
        public Output<string> FilesSize { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Wedata/resource:Resource", name, args ?? new ResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Wedata/resource:Resource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, state, options);
        }
    }

    public sealed class ResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cos bucket name.
        /// </summary>
        [Input("cosBucketName", required: true)]
        public Input<string> CosBucketName { get; set; } = null!;

        /// <summary>
        /// Cos bucket region.
        /// </summary>
        [Input("cosRegion", required: true)]
        public Input<string> CosRegion { get; set; } = null!;

        /// <summary>
        /// File name.
        /// </summary>
        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        /// <summary>
        /// For file path:/datastudio/resource/projectId/folderName; for folder path:/datastudio/resource/folderName.
        /// </summary>
        [Input("filePath", required: true)]
        public Input<string> FilePath { get; set; } = null!;

        /// <summary>
        /// File size.
        /// </summary>
        [Input("filesSize", required: true)]
        public Input<string> FilesSize { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public ResourceArgs()
        {
        }
    }

    public sealed class ResourceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cos bucket name.
        /// </summary>
        [Input("cosBucketName")]
        public Input<string>? CosBucketName { get; set; }

        /// <summary>
        /// Cos bucket region.
        /// </summary>
        [Input("cosRegion")]
        public Input<string>? CosRegion { get; set; }

        /// <summary>
        /// File name.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// For file path:/datastudio/resource/projectId/folderName; for folder path:/datastudio/resource/folderName.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// File size.
        /// </summary>
        [Input("filesSize")]
        public Input<string>? FilesSize { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public ResourceState()
        {
        }
    }
}
