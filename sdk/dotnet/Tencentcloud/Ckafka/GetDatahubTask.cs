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
    public static class GetDatahubTask
    {
        /// <summary>
        /// Use this data source to query detailed information of ckafka datahub_task
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
        ///         var datahubTask = Output.Create(Tencentcloud.Ckafka.GetDatahubTask.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatahubTaskResult> InvokeAsync(GetDatahubTaskArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatahubTaskResult>("tencentcloud:Ckafka/getDatahubTask:getDatahubTask", args ?? new GetDatahubTaskArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ckafka datahub_task
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
        ///         var datahubTask = Output.Create(Tencentcloud.Ckafka.GetDatahubTask.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatahubTaskResult> Invoke(GetDatahubTaskInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatahubTaskResult>("tencentcloud:Ckafka/getDatahubTask:getDatahubTask", args ?? new GetDatahubTaskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatahubTaskArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Resource.
        /// </summary>
        [Input("resource")]
        public string? Resource { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// search key.
        /// </summary>
        [Input("searchWord")]
        public string? SearchWord { get; set; }

        /// <summary>
        /// The source type.
        /// </summary>
        [Input("sourceType")]
        public string? SourceType { get; set; }

        /// <summary>
        /// Destination type of dump.
        /// </summary>
        [Input("targetType")]
        public string? TargetType { get; set; }

        /// <summary>
        /// Task type, SOURCE|SINK.
        /// </summary>
        [Input("taskType")]
        public string? TaskType { get; set; }

        public GetDatahubTaskArgs()
        {
        }
    }

    public sealed class GetDatahubTaskInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Resource.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// search key.
        /// </summary>
        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        /// <summary>
        /// The source type.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// Destination type of dump.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        /// <summary>
        /// Task type, SOURCE|SINK.
        /// </summary>
        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        public GetDatahubTaskInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatahubTaskResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The topic name of the topic sold separately.
        /// </summary>
        public readonly string? Resource;
        public readonly string? ResultOutputFile;
        public readonly string? SearchWord;
        public readonly string? SourceType;
        public readonly string? TargetType;
        /// <summary>
        /// Datahub task information list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatahubTaskTaskListResult> TaskLists;
        /// <summary>
        /// TaskType, SOURCE|SINK.
        /// </summary>
        public readonly string? TaskType;

        [OutputConstructor]
        private GetDatahubTaskResult(
            string id,

            string? resource,

            string? resultOutputFile,

            string? searchWord,

            string? sourceType,

            string? targetType,

            ImmutableArray<Outputs.GetDatahubTaskTaskListResult> taskLists,

            string? taskType)
        {
            Id = id;
            Resource = resource;
            ResultOutputFile = resultOutputFile;
            SearchWord = searchWord;
            SourceType = sourceType;
            TargetType = targetType;
            TaskLists = taskLists;
            TaskType = taskType;
        }
    }
}