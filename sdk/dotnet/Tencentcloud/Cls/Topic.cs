// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls
{
    /// <summary>
    /// Provides a resource to create a cls topic.
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
    ///     var exampleLogset = new Tencentcloud.Cls.Logset("exampleLogset", new()
    ///     {
    ///         LogsetName = "tf_example",
    ///         Tags = 
    ///         {
    ///             { "demo", "test" },
    ///         },
    ///     });
    /// 
    ///     var exampleTopic = new Tencentcloud.Cls.Topic("exampleTopic", new()
    ///     {
    ///         TopicName = "tf_example",
    ///         LogsetId = exampleLogset.Id,
    ///         AutoSplit = false,
    ///         MaxSplitPartitions = 20,
    ///         PartitionCount = 1,
    ///         Period = 30,
    ///         StorageType = "hot",
    ///         Describes = "Test Demo.",
    ///         HotPeriod = 10,
    ///         Tags = 
    ///         {
    ///             { "test", "test" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// cls topic can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cls/topic:Topic example 2f5764c1-c833-44c5-84c7-950979b2a278
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cls/topic:Topic")]
    public partial class Topic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable automatic split. Default value: true.
        /// </summary>
        [Output("autoSplit")]
        public Output<bool> AutoSplit { get; private set; } = null!;

        /// <summary>
        /// Log Topic Description.
        /// </summary>
        [Output("describes")]
        public Output<string?> Describes { get; private set; } = null!;

        /// <summary>
        /// 0: Turn off log sinking. Non 0: The number of days of standard storage after enabling log settling. HotPeriod needs to be greater than or equal to 7 and less than Period. Only effective when StorageType is hot.
        /// </summary>
        [Output("hotPeriod")]
        public Output<int> HotPeriod { get; private set; } = null!;

        /// <summary>
        /// Logset ID.
        /// </summary>
        [Output("logsetId")]
        public Output<string> LogsetId { get; private set; } = null!;

        /// <summary>
        /// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
        /// </summary>
        [Output("maxSplitPartitions")]
        public Output<int> MaxSplitPartitions { get; private set; } = null!;

        /// <summary>
        /// Number of log topic partitions. Default value: 1. Maximum value: 10.
        /// </summary>
        [Output("partitionCount")]
        public Output<int> PartitionCount { get; private set; } = null!;

        /// <summary>
        /// Lifecycle in days. Value range: 1~366. Default value: 30.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;

        /// <summary>
        /// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is passed in, please contact the customer service to add the log topic to the allowlist first.
        /// </summary>
        [Output("storageType")]
        public Output<string> StorageType { get; private set; } = null!;

        /// <summary>
        /// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Log topic name.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic split. Default value: true.
        /// </summary>
        [Input("autoSplit")]
        public Input<bool>? AutoSplit { get; set; }

        /// <summary>
        /// Log Topic Description.
        /// </summary>
        [Input("describes")]
        public Input<string>? Describes { get; set; }

        /// <summary>
        /// 0: Turn off log sinking. Non 0: The number of days of standard storage after enabling log settling. HotPeriod needs to be greater than or equal to 7 and less than Period. Only effective when StorageType is hot.
        /// </summary>
        [Input("hotPeriod")]
        public Input<int>? HotPeriod { get; set; }

        /// <summary>
        /// Logset ID.
        /// </summary>
        [Input("logsetId", required: true)]
        public Input<string> LogsetId { get; set; } = null!;

        /// <summary>
        /// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
        /// </summary>
        [Input("maxSplitPartitions")]
        public Input<int>? MaxSplitPartitions { get; set; }

        /// <summary>
        /// Number of log topic partitions. Default value: 1. Maximum value: 10.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// Lifecycle in days. Value range: 1~366. Default value: 30.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is passed in, please contact the customer service to add the log topic to the allowlist first.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Log topic name.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public TopicArgs()
        {
        }
        public static new TopicArgs Empty => new TopicArgs();
    }

    public sealed class TopicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic split. Default value: true.
        /// </summary>
        [Input("autoSplit")]
        public Input<bool>? AutoSplit { get; set; }

        /// <summary>
        /// Log Topic Description.
        /// </summary>
        [Input("describes")]
        public Input<string>? Describes { get; set; }

        /// <summary>
        /// 0: Turn off log sinking. Non 0: The number of days of standard storage after enabling log settling. HotPeriod needs to be greater than or equal to 7 and less than Period. Only effective when StorageType is hot.
        /// </summary>
        [Input("hotPeriod")]
        public Input<int>? HotPeriod { get; set; }

        /// <summary>
        /// Logset ID.
        /// </summary>
        [Input("logsetId")]
        public Input<string>? LogsetId { get; set; }

        /// <summary>
        /// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
        /// </summary>
        [Input("maxSplitPartitions")]
        public Input<int>? MaxSplitPartitions { get; set; }

        /// <summary>
        /// Number of log topic partitions. Default value: 1. Maximum value: 10.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// Lifecycle in days. Value range: 1~366. Default value: 30.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is passed in, please contact the customer service to add the log topic to the allowlist first.
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Log topic name.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public TopicState()
        {
        }
        public static new TopicState Empty => new TopicState();
    }
}
