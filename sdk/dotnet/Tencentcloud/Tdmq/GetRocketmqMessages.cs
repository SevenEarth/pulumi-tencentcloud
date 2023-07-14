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
    public static class GetRocketmqMessages
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmq message
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var message = Output.Create(Tencentcloud.Tdmq.GetRocketmqMessages.InvokeAsync(new Tencentcloud.Tdmq.GetRocketmqMessagesArgs
        ///         {
        ///             ClusterId = "rocketmq-rkrbm52djmro",
        ///             EnvironmentId = "keep_ns",
        ///             MsgId = "A9FE8D0567FE15DB97425FC08EEF0000",
        ///             QueryDlqMsg = false,
        ///             TopicName = "keep-topic",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRocketmqMessagesResult> InvokeAsync(GetRocketmqMessagesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRocketmqMessagesResult>("tencentcloud:Tdmq/getRocketmqMessages:getRocketmqMessages", args ?? new GetRocketmqMessagesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmq message
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var message = Output.Create(Tencentcloud.Tdmq.GetRocketmqMessages.InvokeAsync(new Tencentcloud.Tdmq.GetRocketmqMessagesArgs
        ///         {
        ///             ClusterId = "rocketmq-rkrbm52djmro",
        ///             EnvironmentId = "keep_ns",
        ///             MsgId = "A9FE8D0567FE15DB97425FC08EEF0000",
        ///             QueryDlqMsg = false,
        ///             TopicName = "keep-topic",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRocketmqMessagesResult> Invoke(GetRocketmqMessagesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRocketmqMessagesResult>("tencentcloud:Tdmq/getRocketmqMessages:getRocketmqMessages", args ?? new GetRocketmqMessagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRocketmqMessagesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster id.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Environment.
        /// </summary>
        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Message ID.
        /// </summary>
        [Input("msgId", required: true)]
        public string MsgId { get; set; } = null!;

        /// <summary>
        /// The value is true when querying dead letters, only valid for Rocketmq.
        /// </summary>
        [Input("queryDlqMsg")]
        public bool? QueryDlqMsg { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Topic, groupId is passed when querying dead letters.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetRocketmqMessagesArgs()
        {
        }
    }

    public sealed class GetRocketmqMessagesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster id.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Environment.
        /// </summary>
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Message ID.
        /// </summary>
        [Input("msgId", required: true)]
        public Input<string> MsgId { get; set; } = null!;

        /// <summary>
        /// The value is true when querying dead letters, only valid for Rocketmq.
        /// </summary>
        [Input("queryDlqMsg")]
        public Input<bool>? QueryDlqMsg { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Topic, groupId is passed when querying dead letters.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public GetRocketmqMessagesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRocketmqMessagesResult
    {
        /// <summary>
        /// Message body.
        /// </summary>
        public readonly string Body;
        public readonly string ClusterId;
        public readonly string EnvironmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Consumer Group ConsumptionNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqMessagesMessageTrackResult> MessageTracks;
        public readonly string MsgId;
        /// <summary>
        /// Production time.
        /// </summary>
        public readonly string ProduceTime;
        /// <summary>
        /// Producer address.
        /// </summary>
        public readonly string ProducerAddr;
        /// <summary>
        /// Detailed parameters.
        /// </summary>
        public readonly string Properties;
        public readonly bool? QueryDlqMsg;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// The topic name displayed on the details pageNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string ShowTopicName;
        public readonly string TopicName;

        [OutputConstructor]
        private GetRocketmqMessagesResult(
            string body,

            string clusterId,

            string environmentId,

            string id,

            ImmutableArray<Outputs.GetRocketmqMessagesMessageTrackResult> messageTracks,

            string msgId,

            string produceTime,

            string producerAddr,

            string properties,

            bool? queryDlqMsg,

            string? resultOutputFile,

            string showTopicName,

            string topicName)
        {
            Body = body;
            ClusterId = clusterId;
            EnvironmentId = environmentId;
            Id = id;
            MessageTracks = messageTracks;
            MsgId = msgId;
            ProduceTime = produceTime;
            ProducerAddr = producerAddr;
            Properties = properties;
            QueryDlqMsg = queryDlqMsg;
            ResultOutputFile = resultOutputFile;
            ShowTopicName = showTopicName;
            TopicName = topicName;
        }
    }
}
