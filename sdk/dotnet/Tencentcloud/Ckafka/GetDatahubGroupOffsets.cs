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
    public static class GetDatahubGroupOffsets
    {
        /// <summary>
        /// Use this data source to query detailed information of ckafka datahub_group_offsets
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
        ///         var datahubGroupOffsets = Output.Create(Tencentcloud.Ckafka.GetDatahubGroupOffsets.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatahubGroupOffsetsResult> InvokeAsync(GetDatahubGroupOffsetsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatahubGroupOffsetsResult>("tencentcloud:Ckafka/getDatahubGroupOffsets:getDatahubGroupOffsets", args ?? new GetDatahubGroupOffsetsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ckafka datahub_group_offsets
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
        ///         var datahubGroupOffsets = Output.Create(Tencentcloud.Ckafka.GetDatahubGroupOffsets.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatahubGroupOffsetsResult> Invoke(GetDatahubGroupOffsetsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatahubGroupOffsetsResult>("tencentcloud:Ckafka/getDatahubGroupOffsets:getDatahubGroupOffsets", args ?? new GetDatahubGroupOffsetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatahubGroupOffsetsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Kafka consumer group.
        /// </summary>
        [Input("group", required: true)]
        public string Group { get; set; } = null!;

        /// <summary>
        /// topic name that the task subscribe.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// fuzzy match topicName.
        /// </summary>
        [Input("searchWord")]
        public string? SearchWord { get; set; }

        public GetDatahubGroupOffsetsArgs()
        {
        }
    }

    public sealed class GetDatahubGroupOffsetsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Kafka consumer group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// topic name that the task subscribe.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// fuzzy match topicName.
        /// </summary>
        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        public GetDatahubGroupOffsetsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatahubGroupOffsetsResult
    {
        public readonly string Group;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? ResultOutputFile;
        public readonly string? SearchWord;
        /// <summary>
        /// The topic array, where each element is a json object.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatahubGroupOffsetsTopicListResult> TopicLists;

        [OutputConstructor]
        private GetDatahubGroupOffsetsResult(
            string group,

            string id,

            string name,

            string? resultOutputFile,

            string? searchWord,

            ImmutableArray<Outputs.GetDatahubGroupOffsetsTopicListResult> topicLists)
        {
            Group = group;
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
            SearchWord = searchWord;
            TopicLists = topicLists;
        }
    }
}
