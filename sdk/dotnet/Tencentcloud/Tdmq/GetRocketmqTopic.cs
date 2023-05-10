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
    public static class GetRocketmqTopic
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq topic
        /// </summary>
        public static Task<GetRocketmqTopicResult> InvokeAsync(GetRocketmqTopicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRocketmqTopicResult>("tencentcloud:Tdmq/getRocketmqTopic:getRocketmqTopic", args ?? new GetRocketmqTopicArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq topic
        /// </summary>
        public static Output<GetRocketmqTopicResult> Invoke(GetRocketmqTopicInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRocketmqTopicResult>("tencentcloud:Tdmq/getRocketmqTopic:getRocketmqTopic", args ?? new GetRocketmqTopicInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRocketmqTopicArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Search by topic name. Fuzzy query is supported.
        /// </summary>
        [Input("filterName")]
        public string? FilterName { get; set; }

        [Input("filterTypes")]
        private List<string>? _filterTypes;

        /// <summary>
        /// Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
        /// </summary>
        public List<string> FilterTypes
        {
            get => _filterTypes ?? (_filterTypes = new List<string>());
            set => _filterTypes = value;
        }

        /// <summary>
        /// Namespace.
        /// </summary>
        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRocketmqTopicArgs()
        {
        }
    }

    public sealed class GetRocketmqTopicInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Search by topic name. Fuzzy query is supported.
        /// </summary>
        [Input("filterName")]
        public Input<string>? FilterName { get; set; }

        [Input("filterTypes")]
        private InputList<string>? _filterTypes;

        /// <summary>
        /// Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
        /// </summary>
        public InputList<string> FilterTypes
        {
            get => _filterTypes ?? (_filterTypes = new InputList<string>());
            set => _filterTypes = value;
        }

        /// <summary>
        /// Namespace.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRocketmqTopicInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRocketmqTopicResult
    {
        public readonly string ClusterId;
        public readonly string? FilterName;
        public readonly ImmutableArray<string> FilterTypes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string NamespaceId;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// List of topic information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqTopicTopicResult> Topics;

        [OutputConstructor]
        private GetRocketmqTopicResult(
            string clusterId,

            string? filterName,

            ImmutableArray<string> filterTypes,

            string id,

            string namespaceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetRocketmqTopicTopicResult> topics)
        {
            ClusterId = clusterId;
            FilterName = filterName;
            FilterTypes = filterTypes;
            Id = id;
            NamespaceId = namespaceId;
            ResultOutputFile = resultOutputFile;
            Topics = topics;
        }
    }
}
