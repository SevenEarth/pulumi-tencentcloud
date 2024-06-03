// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis
{
    public static class GetParamRecords
    {
        /// <summary>
        /// Use this data source to query detailed information of redis param records
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
        ///     var paramRecords = Tencentcloud.Redis.GetParamRecords.Invoke(new()
        ///     {
        ///         InstanceId = "crs-c1nl9rpv",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetParamRecordsResult> InvokeAsync(GetParamRecordsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetParamRecordsResult>("tencentcloud:Redis/getParamRecords:getParamRecords", args ?? new GetParamRecordsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of redis param records
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
        ///     var paramRecords = Tencentcloud.Redis.GetParamRecords.Invoke(new()
        ///     {
        ///         InstanceId = "crs-c1nl9rpv",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetParamRecordsResult> Invoke(GetParamRecordsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetParamRecordsResult>("tencentcloud:Redis/getParamRecords:getParamRecords", args ?? new GetParamRecordsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetParamRecordsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetParamRecordsArgs()
        {
        }
        public static new GetParamRecordsArgs Empty => new GetParamRecordsArgs();
    }

    public sealed class GetParamRecordsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetParamRecordsInvokeArgs()
        {
        }
        public static new GetParamRecordsInvokeArgs Empty => new GetParamRecordsInvokeArgs();
    }


    [OutputType]
    public sealed class GetParamRecordsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// The parameter name.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetParamRecordsInstanceParamHistoryResult> InstanceParamHistories;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetParamRecordsResult(
            string id,

            string instanceId,

            ImmutableArray<Outputs.GetParamRecordsInstanceParamHistoryResult> instanceParamHistories,

            string? resultOutputFile)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceParamHistories = instanceParamHistories;
            ResultOutputFile = resultOutputFile;
        }
    }
}
