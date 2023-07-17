// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    public static class GetTagRetentionExecutions
    {
        /// <summary>
        /// Use this data source to query detailed information of tcr tag_retention_executions
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
        ///         var tagRetentionExecutions = Output.Create(Tencentcloud.Tcr.GetTagRetentionExecutions.InvokeAsync(new Tencentcloud.Tcr.GetTagRetentionExecutionsArgs
        ///         {
        ///             RegistryId = "tcr_ins_id",
        ///             RetentionId = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTagRetentionExecutionsResult> InvokeAsync(GetTagRetentionExecutionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagRetentionExecutionsResult>("tencentcloud:Tcr/getTagRetentionExecutions:getTagRetentionExecutions", args ?? new GetTagRetentionExecutionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tcr tag_retention_executions
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
        ///         var tagRetentionExecutions = Output.Create(Tencentcloud.Tcr.GetTagRetentionExecutions.InvokeAsync(new Tencentcloud.Tcr.GetTagRetentionExecutionsArgs
        ///         {
        ///             RegistryId = "tcr_ins_id",
        ///             RetentionId = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTagRetentionExecutionsResult> Invoke(GetTagRetentionExecutionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTagRetentionExecutionsResult>("tencentcloud:Tcr/getTagRetentionExecutions:getTagRetentionExecutions", args ?? new GetTagRetentionExecutionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagRetentionExecutionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// retention id.
        /// </summary>
        [Input("retentionId", required: true)]
        public int RetentionId { get; set; }

        public GetTagRetentionExecutionsArgs()
        {
        }
    }

    public sealed class GetTagRetentionExecutionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// retention id.
        /// </summary>
        [Input("retentionId", required: true)]
        public Input<int> RetentionId { get; set; } = null!;

        public GetTagRetentionExecutionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagRetentionExecutionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RegistryId;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// list of version retention execution records.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTagRetentionExecutionsRetentionExecutionListResult> RetentionExecutionLists;
        /// <summary>
        /// retention id.
        /// </summary>
        public readonly int RetentionId;

        [OutputConstructor]
        private GetTagRetentionExecutionsResult(
            string id,

            string registryId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetTagRetentionExecutionsRetentionExecutionListResult> retentionExecutionLists,

            int retentionId)
        {
            Id = id;
            RegistryId = registryId;
            ResultOutputFile = resultOutputFile;
            RetentionExecutionLists = retentionExecutionLists;
            RetentionId = retentionId;
        }
    }
}
