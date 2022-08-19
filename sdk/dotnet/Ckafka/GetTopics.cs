// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka
{
    public static class GetTopics
    {
        /// <summary>
        /// Use this data source to query detailed information of ckafka topic.
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
        ///         var foo = new Tencentcloud.Ckafka.Topic("foo", new Tencentcloud.Ckafka.TopicArgs
        ///         {
        ///             CleanUpPolicy = "delete",
        ///             EnableWhiteList = true,
        ///             InstanceId = "ckafka-f9ife4zz",
        ///             IpWhiteLists = 
        ///             {
        ///                 "ip1",
        ///                 "ip2",
        ///             },
        ///             MaxMessageBytes = 0,
        ///             Note = "topic note",
        ///             PartitionNum = 1,
        ///             ReplicaNum = 2,
        ///             Retention = 60000,
        ///             Segment = 3600000,
        ///             SyncReplicaMinNum = 1,
        ///             TopicName = "example",
        ///             UncleanLeaderElectionEnable = false,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTopicsResult> InvokeAsync(GetTopicsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicsResult>("tencentcloud:Ckafka/getTopics:getTopics", args ?? new GetTopicsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ckafka topic.
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
        ///         var foo = new Tencentcloud.Ckafka.Topic("foo", new Tencentcloud.Ckafka.TopicArgs
        ///         {
        ///             CleanUpPolicy = "delete",
        ///             EnableWhiteList = true,
        ///             InstanceId = "ckafka-f9ife4zz",
        ///             IpWhiteLists = 
        ///             {
        ///                 "ip1",
        ///                 "ip2",
        ///             },
        ///             MaxMessageBytes = 0,
        ///             Note = "topic note",
        ///             PartitionNum = 1,
        ///             ReplicaNum = 2,
        ///             Retention = 60000,
        ///             Segment = 3600000,
        ///             SyncReplicaMinNum = 1,
        ///             TopicName = "example",
        ///             UncleanLeaderElectionEnable = false,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTopicsResult> Invoke(GetTopicsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTopicsResult>("tencentcloud:Ckafka/getTopics:getTopics", args ?? new GetTopicsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Ckafka instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-). The length range is from 1 to 64.
        /// </summary>
        [Input("topicName")]
        public string? TopicName { get; set; }

        public GetTopicsArgs()
        {
        }
    }

    public sealed class GetTopicsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Ckafka instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-). The length range is from 1 to 64.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public GetTopicsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// A list of instances. Each element contains the following attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicsInstanceListResult> InstanceLists;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Name of the CKafka topic.
        /// </summary>
        public readonly string? TopicName;

        [OutputConstructor]
        private GetTopicsResult(
            string id,

            string instanceId,

            ImmutableArray<Outputs.GetTopicsInstanceListResult> instanceLists,

            string? resultOutputFile,

            string? topicName)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceLists = instanceLists;
            ResultOutputFile = resultOutputFile;
            TopicName = topicName;
        }
    }
}
