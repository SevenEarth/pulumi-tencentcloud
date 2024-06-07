// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch
{
    public static class GetLogstashInstanceOperations
    {
        /// <summary>
        /// Use this data source to query detailed information of elasticsearch logstash_instance_operations
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
        ///     var logstashInstanceOperations = Tencentcloud.Elasticsearch.GetLogstashInstanceOperations.Invoke(new()
        ///     {
        ///         EndTime = "2023-10-31 10:12:45",
        ///         InstanceId = "ls-xxxxxx",
        ///         StartTime = "2018-01-01 00:00:00",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLogstashInstanceOperationsResult> InvokeAsync(GetLogstashInstanceOperationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogstashInstanceOperationsResult>("tencentcloud:Elasticsearch/getLogstashInstanceOperations:getLogstashInstanceOperations", args ?? new GetLogstashInstanceOperationsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of elasticsearch logstash_instance_operations
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
        ///     var logstashInstanceOperations = Tencentcloud.Elasticsearch.GetLogstashInstanceOperations.Invoke(new()
        ///     {
        ///         EndTime = "2023-10-31 10:12:45",
        ///         InstanceId = "ls-xxxxxx",
        ///         StartTime = "2018-01-01 00:00:00",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLogstashInstanceOperationsResult> Invoke(GetLogstashInstanceOperationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogstashInstanceOperationsResult>("tencentcloud:Elasticsearch/getLogstashInstanceOperations:getLogstashInstanceOperations", args ?? new GetLogstashInstanceOperationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogstashInstanceOperationsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// End time, e.g. 2019-03-30 20:18:03.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time, e.g. 2019-03-07 16:30:39.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetLogstashInstanceOperationsArgs()
        {
        }
        public static new GetLogstashInstanceOperationsArgs Empty => new GetLogstashInstanceOperationsArgs();
    }

    public sealed class GetLogstashInstanceOperationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// End time, e.g. 2019-03-30 20:18:03.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time, e.g. 2019-03-07 16:30:39.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetLogstashInstanceOperationsInvokeArgs()
        {
        }
        public static new GetLogstashInstanceOperationsInvokeArgs Empty => new GetLogstashInstanceOperationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogstashInstanceOperationsResult
    {
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Operation records.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLogstashInstanceOperationsOperationResult> Operations;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Start time.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GetLogstashInstanceOperationsResult(
            string endTime,

            string id,

            string instanceId,

            ImmutableArray<Outputs.GetLogstashInstanceOperationsOperationResult> operations,

            string? resultOutputFile,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Operations = operations;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}
