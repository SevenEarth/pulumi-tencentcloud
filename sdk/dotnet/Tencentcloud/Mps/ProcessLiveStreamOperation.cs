// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps
{
    /// <summary>
    /// Provides a resource to create a mps process_live_stream_operation
    /// 
    /// ## Example Usage
    /// ### Process mps live stream through CMQ
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var output = new Tencentcloud.Cos.Bucket("output", new Tencentcloud.Cos.BucketArgs
    ///         {
    ///             Bucket = $"tf-bucket-mps-process-live-stream-output-{local.App_id}",
    ///             ForceClean = true,
    ///             Acl = "public-read",
    ///         });
    ///         var operation = new Tencentcloud.Mps.ProcessLiveStreamOperation("operation", new Tencentcloud.Mps.ProcessLiveStreamOperationArgs
    ///         {
    ///             Url = "http://www.abc.com/abc.m3u8",
    ///             TaskNotifyConfig = new Tencentcloud.Mps.Inputs.ProcessLiveStreamOperationTaskNotifyConfigArgs
    ///             {
    ///                 CmqModel = "Queue",
    ///                 CmqRegion = "gz",
    ///                 QueueName = "test",
    ///                 TopicName = "test",
    ///                 NotifyType = "CMQ",
    ///             },
    ///             OutputStorage = new Tencentcloud.Mps.Inputs.ProcessLiveStreamOperationOutputStorageArgs
    ///             {
    ///                 Type = "COS",
    ///                 CosOutputStorage = new Tencentcloud.Mps.Inputs.ProcessLiveStreamOperationOutputStorageCosOutputStorageArgs
    ///                 {
    ///                     Bucket = output.CosBucket,
    ///                     Region = "%s",
    ///                 },
    ///             },
    ///             OutputDir = "/output/",
    ///             AiContentReviewTask = new Tencentcloud.Mps.Inputs.ProcessLiveStreamOperationAiContentReviewTaskArgs
    ///             {
    ///                 Definition = 10,
    ///             },
    ///             AiRecognitionTask = new Tencentcloud.Mps.Inputs.ProcessLiveStreamOperationAiRecognitionTaskArgs
    ///             {
    ///                 Definition = 10,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mps/processLiveStreamOperation:ProcessLiveStreamOperation")]
    public partial class ProcessLiveStreamOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// AI video intelligent analysis input parameter types.
        /// </summary>
        [Output("aiAnalysisTask")]
        public Output<Outputs.ProcessLiveStreamOperationAiAnalysisTask?> AiAnalysisTask { get; private set; } = null!;

        /// <summary>
        /// Type parameter of a video content audit task.
        /// </summary>
        [Output("aiContentReviewTask")]
        public Output<Outputs.ProcessLiveStreamOperationAiContentReviewTask?> AiContentReviewTask { get; private set; } = null!;

        /// <summary>
        /// The parameters for a video quality control task.
        /// </summary>
        [Output("aiQualityControlTask")]
        public Output<Outputs.ProcessLiveStreamOperationAiQualityControlTask?> AiQualityControlTask { get; private set; } = null!;

        /// <summary>
        /// Type parameter of video content recognition task.
        /// </summary>
        [Output("aiRecognitionTask")]
        public Output<Outputs.ProcessLiveStreamOperationAiRecognitionTask?> AiRecognitionTask { get; private set; } = null!;

        /// <summary>
        /// Target directory of a live stream processing output file, such as `/movie/201909/`. If this parameter is left empty, the `/` directory will be used.
        /// </summary>
        [Output("outputDir")]
        public Output<string?> OutputDir { get; private set; } = null!;

        /// <summary>
        /// Target bucket of a live stream processing output file. This parameter is required if a file will be output.
        /// </summary>
        [Output("outputStorage")]
        public Output<Outputs.ProcessLiveStreamOperationOutputStorage?> OutputStorage { get; private set; } = null!;

        /// <summary>
        /// The scheme ID.Note 1: About `OutputStorage` and `OutputDir`:If an output storage and directory are specified for a subtask of the scheme, those output settings will be applied.If an output storage and directory are not specified for the subtasks of a scheme, the output parameters passed in the `ProcessMedia` API will be applied.Note 2: If `TaskNotifyConfig` is specified, the specified settings will be used instead of the default callback settings of the scheme.
        /// </summary>
        [Output("scheduleId")]
        public Output<int?> ScheduleId { get; private set; } = null!;

        /// <summary>
        /// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
        /// </summary>
        [Output("sessionContext")]
        public Output<string?> SessionContext { get; private set; } = null!;

        /// <summary>
        /// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
        /// </summary>
        [Output("sessionId")]
        public Output<string?> SessionId { get; private set; } = null!;

        /// <summary>
        /// Event notification information of a task, which is used to specify the live stream processing result.
        /// </summary>
        [Output("taskNotifyConfig")]
        public Output<Outputs.ProcessLiveStreamOperationTaskNotifyConfig> TaskNotifyConfig { get; private set; } = null!;

        /// <summary>
        /// Live stream URL, which must be a live stream file address. RTMP, HLS, and FLV are supported.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ProcessLiveStreamOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProcessLiveStreamOperation(string name, ProcessLiveStreamOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/processLiveStreamOperation:ProcessLiveStreamOperation", name, args ?? new ProcessLiveStreamOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProcessLiveStreamOperation(string name, Input<string> id, ProcessLiveStreamOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mps/processLiveStreamOperation:ProcessLiveStreamOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProcessLiveStreamOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProcessLiveStreamOperation Get(string name, Input<string> id, ProcessLiveStreamOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ProcessLiveStreamOperation(name, id, state, options);
        }
    }

    public sealed class ProcessLiveStreamOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AI video intelligent analysis input parameter types.
        /// </summary>
        [Input("aiAnalysisTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiAnalysisTaskArgs>? AiAnalysisTask { get; set; }

        /// <summary>
        /// Type parameter of a video content audit task.
        /// </summary>
        [Input("aiContentReviewTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiContentReviewTaskArgs>? AiContentReviewTask { get; set; }

        /// <summary>
        /// The parameters for a video quality control task.
        /// </summary>
        [Input("aiQualityControlTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiQualityControlTaskArgs>? AiQualityControlTask { get; set; }

        /// <summary>
        /// Type parameter of video content recognition task.
        /// </summary>
        [Input("aiRecognitionTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiRecognitionTaskArgs>? AiRecognitionTask { get; set; }

        /// <summary>
        /// Target directory of a live stream processing output file, such as `/movie/201909/`. If this parameter is left empty, the `/` directory will be used.
        /// </summary>
        [Input("outputDir")]
        public Input<string>? OutputDir { get; set; }

        /// <summary>
        /// Target bucket of a live stream processing output file. This parameter is required if a file will be output.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.ProcessLiveStreamOperationOutputStorageArgs>? OutputStorage { get; set; }

        /// <summary>
        /// The scheme ID.Note 1: About `OutputStorage` and `OutputDir`:If an output storage and directory are specified for a subtask of the scheme, those output settings will be applied.If an output storage and directory are not specified for the subtasks of a scheme, the output parameters passed in the `ProcessMedia` API will be applied.Note 2: If `TaskNotifyConfig` is specified, the specified settings will be used instead of the default callback settings of the scheme.
        /// </summary>
        [Input("scheduleId")]
        public Input<int>? ScheduleId { get; set; }

        /// <summary>
        /// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
        /// </summary>
        [Input("sessionContext")]
        public Input<string>? SessionContext { get; set; }

        /// <summary>
        /// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
        /// </summary>
        [Input("sessionId")]
        public Input<string>? SessionId { get; set; }

        /// <summary>
        /// Event notification information of a task, which is used to specify the live stream processing result.
        /// </summary>
        [Input("taskNotifyConfig", required: true)]
        public Input<Inputs.ProcessLiveStreamOperationTaskNotifyConfigArgs> TaskNotifyConfig { get; set; } = null!;

        /// <summary>
        /// Live stream URL, which must be a live stream file address. RTMP, HLS, and FLV are supported.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ProcessLiveStreamOperationArgs()
        {
        }
    }

    public sealed class ProcessLiveStreamOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AI video intelligent analysis input parameter types.
        /// </summary>
        [Input("aiAnalysisTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiAnalysisTaskGetArgs>? AiAnalysisTask { get; set; }

        /// <summary>
        /// Type parameter of a video content audit task.
        /// </summary>
        [Input("aiContentReviewTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiContentReviewTaskGetArgs>? AiContentReviewTask { get; set; }

        /// <summary>
        /// The parameters for a video quality control task.
        /// </summary>
        [Input("aiQualityControlTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiQualityControlTaskGetArgs>? AiQualityControlTask { get; set; }

        /// <summary>
        /// Type parameter of video content recognition task.
        /// </summary>
        [Input("aiRecognitionTask")]
        public Input<Inputs.ProcessLiveStreamOperationAiRecognitionTaskGetArgs>? AiRecognitionTask { get; set; }

        /// <summary>
        /// Target directory of a live stream processing output file, such as `/movie/201909/`. If this parameter is left empty, the `/` directory will be used.
        /// </summary>
        [Input("outputDir")]
        public Input<string>? OutputDir { get; set; }

        /// <summary>
        /// Target bucket of a live stream processing output file. This parameter is required if a file will be output.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.ProcessLiveStreamOperationOutputStorageGetArgs>? OutputStorage { get; set; }

        /// <summary>
        /// The scheme ID.Note 1: About `OutputStorage` and `OutputDir`:If an output storage and directory are specified for a subtask of the scheme, those output settings will be applied.If an output storage and directory are not specified for the subtasks of a scheme, the output parameters passed in the `ProcessMedia` API will be applied.Note 2: If `TaskNotifyConfig` is specified, the specified settings will be used instead of the default callback settings of the scheme.
        /// </summary>
        [Input("scheduleId")]
        public Input<int>? ScheduleId { get; set; }

        /// <summary>
        /// The source context which is used to pass through the user request information. The task flow status change callback will return the value of this field. It can contain up to 1,000 characters.
        /// </summary>
        [Input("sessionContext")]
        public Input<string>? SessionContext { get; set; }

        /// <summary>
        /// The ID used for deduplication. If there was a request with the same ID in the last seven days, the current request will return an error. The ID can contain up to 50 characters. If this parameter is left empty or an empty string is entered, no deduplication will be performed.
        /// </summary>
        [Input("sessionId")]
        public Input<string>? SessionId { get; set; }

        /// <summary>
        /// Event notification information of a task, which is used to specify the live stream processing result.
        /// </summary>
        [Input("taskNotifyConfig")]
        public Input<Inputs.ProcessLiveStreamOperationTaskNotifyConfigGetArgs>? TaskNotifyConfig { get; set; }

        /// <summary>
        /// Live stream URL, which must be a live stream file address. RTMP, HLS, and FLV are supported.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProcessLiveStreamOperationState()
        {
        }
    }
}
