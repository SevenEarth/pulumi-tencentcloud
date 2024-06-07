// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css
{
    public static class GetPullStreamTaskStatus
    {
        /// <summary>
        /// Use this data source to query detailed information of css pull_stream_task_status
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pullStreamTaskStatus = Tencentcloud.Css.GetPullStreamTaskStatus.Invoke(new()
        ///     {
        ///         TaskId = "63229997",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPullStreamTaskStatusResult> InvokeAsync(GetPullStreamTaskStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPullStreamTaskStatusResult>("tencentcloud:Css/getPullStreamTaskStatus:getPullStreamTaskStatus", args ?? new GetPullStreamTaskStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of css pull_stream_task_status
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pullStreamTaskStatus = Tencentcloud.Css.GetPullStreamTaskStatus.Invoke(new()
        ///     {
        ///         TaskId = "63229997",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPullStreamTaskStatusResult> Invoke(GetPullStreamTaskStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPullStreamTaskStatusResult>("tencentcloud:Css/getPullStreamTaskStatus:getPullStreamTaskStatus", args ?? new GetPullStreamTaskStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPullStreamTaskStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Task ID.
        /// </summary>
        [Input("taskId", required: true)]
        public string TaskId { get; set; } = null!;

        public GetPullStreamTaskStatusArgs()
        {
        }
        public static new GetPullStreamTaskStatusArgs Empty => new GetPullStreamTaskStatusArgs();
    }

    public sealed class GetPullStreamTaskStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Task ID.
        /// </summary>
        [Input("taskId", required: true)]
        public Input<string> TaskId { get; set; } = null!;

        public GetPullStreamTaskStatusInvokeArgs()
        {
        }
        public static new GetPullStreamTaskStatusInvokeArgs Empty => new GetPullStreamTaskStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetPullStreamTaskStatusResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string TaskId;
        /// <summary>
        /// Task status info.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPullStreamTaskStatusTaskStatusInfoResult> TaskStatusInfos;

        [OutputConstructor]
        private GetPullStreamTaskStatusResult(
            string id,

            string? resultOutputFile,

            string taskId,

            ImmutableArray<Outputs.GetPullStreamTaskStatusTaskStatusInfoResult> taskStatusInfos)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            TaskId = taskId;
            TaskStatusInfos = taskStatusInfos;
        }
    }
}
