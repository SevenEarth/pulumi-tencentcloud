// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka
{
    /// <summary>
    /// Provides a resource to create a ckafka consumer_group_modify_offset
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
    ///     var consumerGroupModifyOffset = new Tencentcloud.Ckafka.ConsumerGroupModifyOffset("consumerGroupModifyOffset", new()
    ///     {
    ///         Group = "xxxxxx",
    ///         InstanceId = "ckafka-xxxxxx",
    ///         Offset = 0,
    ///         Strategy = 2,
    ///         Topics = new[]
    ///         {
    ///             "xxxxxx",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ckafka/consumerGroupModifyOffset:ConsumerGroupModifyOffset")]
    public partial class ConsumerGroupModifyOffset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// kafka group.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// Kafka instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The offset location that needs to be reset. When strategy is 2, this field must be included.
        /// </summary>
        [Output("offset")]
        public Output<int?> Offset { get; private set; } = null!;

        /// <summary>
        /// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
        /// </summary>
        [Output("partitions")]
        public Output<ImmutableArray<int>> Partitions { get; private set; } = null!;

        /// <summary>
        /// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
        /// </summary>
        [Output("shift")]
        public Output<int?> Shift { get; private set; } = null!;

        /// <summary>
        /// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
        /// </summary>
        [Output("shiftTimestamp")]
        public Output<int?> ShiftTimestamp { get; private set; } = null!;

        /// <summary>
        /// Reset the policy of offset.
        /// `0`: Move the offset forward or backward shift bar;
        /// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
        /// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
        /// </summary>
        [Output("strategy")]
        public Output<int> Strategy { get; private set; } = null!;

        /// <summary>
        /// Indicates the topics that needs to be reset. Leave it empty means all.
        /// </summary>
        [Output("topics")]
        public Output<ImmutableArray<string>> Topics { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerGroupModifyOffset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerGroupModifyOffset(string name, ConsumerGroupModifyOffsetArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ckafka/consumerGroupModifyOffset:ConsumerGroupModifyOffset", name, args ?? new ConsumerGroupModifyOffsetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerGroupModifyOffset(string name, Input<string> id, ConsumerGroupModifyOffsetState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ckafka/consumerGroupModifyOffset:ConsumerGroupModifyOffset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerGroupModifyOffset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerGroupModifyOffset Get(string name, Input<string> id, ConsumerGroupModifyOffsetState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerGroupModifyOffset(name, id, state, options);
        }
    }

    public sealed class ConsumerGroupModifyOffsetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// kafka group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// Kafka instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The offset location that needs to be reset. When strategy is 2, this field must be included.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        [Input("partitions")]
        private InputList<int>? _partitions;

        /// <summary>
        /// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
        /// </summary>
        public InputList<int> Partitions
        {
            get => _partitions ?? (_partitions = new InputList<int>());
            set => _partitions = value;
        }

        /// <summary>
        /// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
        /// </summary>
        [Input("shift")]
        public Input<int>? Shift { get; set; }

        /// <summary>
        /// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
        /// </summary>
        [Input("shiftTimestamp")]
        public Input<int>? ShiftTimestamp { get; set; }

        /// <summary>
        /// Reset the policy of offset.
        /// `0`: Move the offset forward or backward shift bar;
        /// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
        /// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
        /// </summary>
        [Input("strategy", required: true)]
        public Input<int> Strategy { get; set; } = null!;

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// Indicates the topics that needs to be reset. Leave it empty means all.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public ConsumerGroupModifyOffsetArgs()
        {
        }
        public static new ConsumerGroupModifyOffsetArgs Empty => new ConsumerGroupModifyOffsetArgs();
    }

    public sealed class ConsumerGroupModifyOffsetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// kafka group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Kafka instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The offset location that needs to be reset. When strategy is 2, this field must be included.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        [Input("partitions")]
        private InputList<int>? _partitions;

        /// <summary>
        /// The list of partition that needs to be reset if no Topics parameter is specified. Resets the partition in the corresponding Partition list of all topics. When Topics is specified, the partition of the corresponding topic list of the specified Partitions list is reset.
        /// </summary>
        public InputList<int> Partitions
        {
            get => _partitions ?? (_partitions = new InputList<int>());
            set => _partitions = value;
        }

        /// <summary>
        /// This field must be included when strategy is 0. If it is greater than zero, the offset will be moved backward by shift bars, and if it is less than zero, the offset will be traced back to the number of shift entries. After the correct reset, the new offset should be (old_offset + shift). It should be noted that if the new offset is less than partition's earliest, it will be set to earliest, and if the latest greater than partition will be set to latest.
        /// </summary>
        [Input("shift")]
        public Input<int>? Shift { get; set; }

        /// <summary>
        /// Unit ms. When strategy is 1, you must include this field, where-2 means to reset the offset to the beginning,-1 means to reset to the latest position (equivalent to emptying), and other values represent the specified time. You will get the offset of the specified time in the topic and then reset it. If there is no message at the specified time, get the last offset.
        /// </summary>
        [Input("shiftTimestamp")]
        public Input<int>? ShiftTimestamp { get; set; }

        /// <summary>
        /// Reset the policy of offset.
        /// `0`: Move the offset forward or backward shift bar;
        /// `1`: Alignment reference (by-duration,to-datetime,to-earliest,to-latest), which means moving the offset to the location of the specified timestamp;
        /// `2`: Alignment reference (to-offset), which means to move the offset to the specified offset location.
        /// </summary>
        [Input("strategy")]
        public Input<int>? Strategy { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// Indicates the topics that needs to be reset. Leave it empty means all.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public ConsumerGroupModifyOffsetState()
        {
        }
        public static new ConsumerGroupModifyOffsetState Empty => new ConsumerGroupModifyOffsetState();
    }
}
