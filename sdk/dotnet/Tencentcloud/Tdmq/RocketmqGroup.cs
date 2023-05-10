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
    /// <summary>
    /// Provides a resource to create a tdmqRocketmq group
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
    ///         var cluster = new Tencentcloud.Tdmq.RocketmqCluster("cluster", new Tencentcloud.Tdmq.RocketmqClusterArgs
    ///         {
    ///             ClusterName = "test_rocketmq",
    ///             Remark = "test recket mq",
    ///         });
    ///         var @namespace = new Tencentcloud.Tdmq.RocketmqNamespace("namespace", new Tencentcloud.Tdmq.RocketmqNamespaceArgs
    ///         {
    ///             ClusterId = cluster.ClusterId,
    ///             NamespaceName = "test_namespace",
    ///             Ttl = 65000,
    ///             RetentionTime = 65000,
    ///             Remark = "test namespace",
    ///         });
    ///         var @group = new Tencentcloud.Tdmq.RocketmqGroup("group", new Tencentcloud.Tdmq.RocketmqGroupArgs
    ///         {
    ///             GroupName = "test_rocketmq_group",
    ///             Namespace = @namespace.NamespaceName,
    ///             ReadEnable = true,
    ///             BroadcastEnable = true,
    ///             ClusterId = cluster.ClusterId,
    ///             Remark = "test rocketmq group",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tdmqRocketmq group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup group group_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup")]
    public partial class RocketmqGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable broadcast consumption.
        /// </summary>
        [Output("broadcastEnable")]
        public Output<bool> BroadcastEnable { get; private set; } = null!;

        /// <summary>
        /// Client protocol.
        /// </summary>
        [Output("clientProtocol")]
        public Output<string> ClientProtocol { get; private set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The number of online consumers.
        /// </summary>
        [Output("consumerNum")]
        public Output<int> ConsumerNum { get; private set; } = null!;

        /// <summary>
        /// Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
        /// </summary>
        [Output("consumerType")]
        public Output<string> ConsumerType { get; private set; } = null!;

        /// <summary>
        /// `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
        /// </summary>
        [Output("consumptionMode")]
        public Output<int> ConsumptionMode { get; private set; } = null!;

        /// <summary>
        /// Creation time in milliseconds.
        /// </summary>
        [Output("createTime")]
        public Output<int> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Group name (8-64 characters).
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Namespace. Currently, only one namespace is supported.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// Whether to enable consumption.
        /// </summary>
        [Output("readEnable")]
        public Output<bool> ReadEnable { get; private set; } = null!;

        /// <summary>
        /// Remarks (up to 128 characters).
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// The number of partitions in a retry topic.
        /// </summary>
        [Output("retryPartitionNum")]
        public Output<int> RetryPartitionNum { get; private set; } = null!;

        /// <summary>
        /// The total number of heaped messages.
        /// </summary>
        [Output("totalAccumulative")]
        public Output<int> TotalAccumulative { get; private set; } = null!;

        /// <summary>
        /// Consumption TPS.
        /// </summary>
        [Output("tps")]
        public Output<int> Tps { get; private set; } = null!;

        /// <summary>
        /// Modification time in milliseconds.
        /// </summary>
        [Output("updateTime")]
        public Output<int> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a RocketmqGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RocketmqGroup(string name, RocketmqGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup", name, args ?? new RocketmqGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RocketmqGroup(string name, Input<string> id, RocketmqGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RocketmqGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RocketmqGroup Get(string name, Input<string> id, RocketmqGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new RocketmqGroup(name, id, state, options);
        }
    }

    public sealed class RocketmqGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable broadcast consumption.
        /// </summary>
        [Input("broadcastEnable", required: true)]
        public Input<bool> BroadcastEnable { get; set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Group name (8-64 characters).
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Namespace. Currently, only one namespace is supported.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// Whether to enable consumption.
        /// </summary>
        [Input("readEnable", required: true)]
        public Input<bool> ReadEnable { get; set; } = null!;

        /// <summary>
        /// Remarks (up to 128 characters).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        public RocketmqGroupArgs()
        {
        }
    }

    public sealed class RocketmqGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable broadcast consumption.
        /// </summary>
        [Input("broadcastEnable")]
        public Input<bool>? BroadcastEnable { get; set; }

        /// <summary>
        /// Client protocol.
        /// </summary>
        [Input("clientProtocol")]
        public Input<string>? ClientProtocol { get; set; }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The number of online consumers.
        /// </summary>
        [Input("consumerNum")]
        public Input<int>? ConsumerNum { get; set; }

        /// <summary>
        /// Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
        /// </summary>
        [Input("consumerType")]
        public Input<string>? ConsumerType { get; set; }

        /// <summary>
        /// `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
        /// </summary>
        [Input("consumptionMode")]
        public Input<int>? ConsumptionMode { get; set; }

        /// <summary>
        /// Creation time in milliseconds.
        /// </summary>
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        /// <summary>
        /// Group name (8-64 characters).
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Namespace. Currently, only one namespace is supported.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Whether to enable consumption.
        /// </summary>
        [Input("readEnable")]
        public Input<bool>? ReadEnable { get; set; }

        /// <summary>
        /// Remarks (up to 128 characters).
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The number of partitions in a retry topic.
        /// </summary>
        [Input("retryPartitionNum")]
        public Input<int>? RetryPartitionNum { get; set; }

        /// <summary>
        /// The total number of heaped messages.
        /// </summary>
        [Input("totalAccumulative")]
        public Input<int>? TotalAccumulative { get; set; }

        /// <summary>
        /// Consumption TPS.
        /// </summary>
        [Input("tps")]
        public Input<int>? Tps { get; set; }

        /// <summary>
        /// Modification time in milliseconds.
        /// </summary>
        [Input("updateTime")]
        public Input<int>? UpdateTime { get; set; }

        public RocketmqGroupState()
        {
        }
    }
}
