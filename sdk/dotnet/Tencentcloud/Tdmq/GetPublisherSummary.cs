// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq
{
    public static class GetPublisherSummary
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmq publisher_summary
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
        ///     var publisherSummary = Tencentcloud.Tdmq.GetPublisherSummary.Invoke(new()
        ///     {
        ///         ClusterId = "pulsar-9n95ax58b9vn",
        ///         Namespace = "keep-ns",
        ///         Topic = "keep-topic",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPublisherSummaryResult> InvokeAsync(GetPublisherSummaryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublisherSummaryResult>("tencentcloud:Tdmq/getPublisherSummary:getPublisherSummary", args ?? new GetPublisherSummaryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmq publisher_summary
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
        ///     var publisherSummary = Tencentcloud.Tdmq.GetPublisherSummary.Invoke(new()
        ///     {
        ///         ClusterId = "pulsar-9n95ax58b9vn",
        ///         Namespace = "keep-ns",
        ///         Topic = "keep-topic",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPublisherSummaryResult> Invoke(GetPublisherSummaryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublisherSummaryResult>("tencentcloud:Tdmq/getPublisherSummary:getPublisherSummary", args ?? new GetPublisherSummaryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublisherSummaryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespace", required: true)]
        public string Namespace { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// subject name.
        /// </summary>
        [Input("topic", required: true)]
        public string Topic { get; set; } = null!;

        public GetPublisherSummaryArgs()
        {
        }
        public static new GetPublisherSummaryArgs Empty => new GetPublisherSummaryArgs();
    }

    public sealed class GetPublisherSummaryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// subject name.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public GetPublisherSummaryInvokeArgs()
        {
        }
        public static new GetPublisherSummaryInvokeArgs Empty => new GetPublisherSummaryInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublisherSummaryResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Production rate (units per second)Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly double MsgRateIn;
        /// <summary>
        /// Production rate (bytes per second)Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly double MsgThroughputIn;
        public readonly string Namespace;
        /// <summary>
        /// number of producersNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int PublisherCount;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Message store size in bytesNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int StorageSize;
        public readonly string Topic;

        [OutputConstructor]
        private GetPublisherSummaryResult(
            string clusterId,

            string id,

            double msgRateIn,

            double msgThroughputIn,

            string @namespace,

            int publisherCount,

            string? resultOutputFile,

            int storageSize,

            string topic)
        {
            ClusterId = clusterId;
            Id = id;
            MsgRateIn = msgRateIn;
            MsgThroughputIn = msgThroughputIn;
            Namespace = @namespace;
            PublisherCount = publisherCount;
            ResultOutputFile = resultOutputFile;
            StorageSize = storageSize;
            Topic = topic;
        }
    }
}
