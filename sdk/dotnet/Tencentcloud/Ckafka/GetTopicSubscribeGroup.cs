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
    public static class GetTopicSubscribeGroup
    {
        /// <summary>
        /// Use this data source to query detailed information of ckafka topic_subscribe_group
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
        ///         var topicSubscribeGroup = Output.Create(Tencentcloud.Ckafka.GetTopicSubscribeGroup.InvokeAsync(new Tencentcloud.Ckafka.GetTopicSubscribeGroupArgs
        ///         {
        ///             InstanceId = "ckafka-xxxxxx",
        ///             TopicName = "xxxxxx",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTopicSubscribeGroupResult> InvokeAsync(GetTopicSubscribeGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicSubscribeGroupResult>("tencentcloud:Ckafka/getTopicSubscribeGroup:getTopicSubscribeGroup", args ?? new GetTopicSubscribeGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ckafka topic_subscribe_group
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
        ///         var topicSubscribeGroup = Output.Create(Tencentcloud.Ckafka.GetTopicSubscribeGroup.InvokeAsync(new Tencentcloud.Ckafka.GetTopicSubscribeGroupArgs
        ///         {
        ///             InstanceId = "ckafka-xxxxxx",
        ///             TopicName = "xxxxxx",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTopicSubscribeGroupResult> Invoke(GetTopicSubscribeGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTopicSubscribeGroupResult>("tencentcloud:Ckafka/getTopicSubscribeGroup:getTopicSubscribeGroup", args ?? new GetTopicSubscribeGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicSubscribeGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// TopicName.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetTopicSubscribeGroupArgs()
        {
        }
    }

    public sealed class GetTopicSubscribeGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// TopicName.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public GetTopicSubscribeGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicSubscribeGroupResult
    {
        /// <summary>
        /// Consumer group information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicSubscribeGroupGroupsInfoResult> GroupsInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly string TopicName;

        [OutputConstructor]
        private GetTopicSubscribeGroupResult(
            ImmutableArray<Outputs.GetTopicSubscribeGroupGroupsInfoResult> groupsInfos,

            string id,

            string instanceId,

            string? resultOutputFile,

            string topicName)
        {
            GroupsInfos = groupsInfos;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            TopicName = topicName;
        }
    }
}
