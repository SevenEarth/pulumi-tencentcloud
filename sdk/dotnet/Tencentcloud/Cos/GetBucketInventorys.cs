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
    public static class GetBucketInventorys
    {
        /// <summary>
        /// Use this data source to query the COS bucket inventorys.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var cosBucketInventorys = Tencentcloud.Cos.GetBucketInventorys.Invoke(new()
        ///     {
        ///         Bucket = "xxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBucketInventorysResult> InvokeAsync(GetBucketInventorysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBucketInventorysResult>("tencentcloud:Cos/getBucketInventorys:getBucketInventorys", args ?? new GetBucketInventorysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query the COS bucket inventorys.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var cosBucketInventorys = Tencentcloud.Cos.GetBucketInventorys.Invoke(new()
        ///     {
        ///         Bucket = "xxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBucketInventorysResult> Invoke(GetBucketInventorysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBucketInventorysResult>("tencentcloud:Cos/getBucketInventorys:getBucketInventorys", args ?? new GetBucketInventorysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketInventorysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBucketInventorysArgs()
        {
        }
        public static new GetBucketInventorysArgs Empty => new GetBucketInventorysArgs();
    }

    public sealed class GetBucketInventorysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBucketInventorysInvokeArgs()
        {
        }
        public static new GetBucketInventorysInvokeArgs Empty => new GetBucketInventorysInvokeArgs();
    }


    [OutputType]
    public sealed class GetBucketInventorysResult
    {
        /// <summary>
        /// Bucket name.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Multiple batch processing task information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketInventorysInventoryResult> Inventorys;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBucketInventorysResult(
            string bucket,

            string id,

            ImmutableArray<Outputs.GetBucketInventorysInventoryResult> inventorys,

            string? resultOutputFile)
        {
            Bucket = bucket;
            Id = id;
            Inventorys = inventorys;
            ResultOutputFile = resultOutputFile;
        }
    }
}
