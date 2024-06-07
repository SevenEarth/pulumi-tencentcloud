// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos
{
    /// <summary>
    /// Provides a resource to abort multipart upload
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
    ///     var abortMultipartUpload = new Tencentcloud.Cos.ObjectAbortMultipartUploadOperation("abortMultipartUpload", new()
    ///     {
    ///         Bucket = "keep-test-xxxxxx",
    ///         Key = "object",
    ///         UploadId = "xxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cos/objectAbortMultipartUploadOperation:ObjectAbortMultipartUploadOperation")]
    public partial class ObjectAbortMultipartUploadOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Object key.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Multipart uploaded id.
        /// </summary>
        [Output("uploadId")]
        public Output<string> UploadId { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectAbortMultipartUploadOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectAbortMultipartUploadOperation(string name, ObjectAbortMultipartUploadOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cos/objectAbortMultipartUploadOperation:ObjectAbortMultipartUploadOperation", name, args ?? new ObjectAbortMultipartUploadOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectAbortMultipartUploadOperation(string name, Input<string> id, ObjectAbortMultipartUploadOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cos/objectAbortMultipartUploadOperation:ObjectAbortMultipartUploadOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectAbortMultipartUploadOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectAbortMultipartUploadOperation Get(string name, Input<string> id, ObjectAbortMultipartUploadOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectAbortMultipartUploadOperation(name, id, state, options);
        }
    }

    public sealed class ObjectAbortMultipartUploadOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Object key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Multipart uploaded id.
        /// </summary>
        [Input("uploadId", required: true)]
        public Input<string> UploadId { get; set; } = null!;

        public ObjectAbortMultipartUploadOperationArgs()
        {
        }
        public static new ObjectAbortMultipartUploadOperationArgs Empty => new ObjectAbortMultipartUploadOperationArgs();
    }

    public sealed class ObjectAbortMultipartUploadOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Object key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Multipart uploaded id.
        /// </summary>
        [Input("uploadId")]
        public Input<string>? UploadId { get; set; }

        public ObjectAbortMultipartUploadOperationState()
        {
        }
        public static new ObjectAbortMultipartUploadOperationState Empty => new ObjectAbortMultipartUploadOperationState();
    }
}
