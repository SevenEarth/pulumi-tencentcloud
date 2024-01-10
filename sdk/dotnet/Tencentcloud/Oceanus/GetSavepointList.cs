// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus
{
    public static class GetSavepointList
    {
        /// <summary>
        /// Use this data source to query detailed information of oceanus savepoint_list
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
        ///         var example = Output.Create(Tencentcloud.Oceanus.GetSavepointList.InvokeAsync(new Tencentcloud.Oceanus.GetSavepointListArgs
        ///         {
        ///             JobId = "cql-314rw6w0",
        ///             WorkSpaceId = "space-2idq8wbr",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSavepointListResult> InvokeAsync(GetSavepointListArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSavepointListResult>("tencentcloud:Oceanus/getSavepointList:getSavepointList", args ?? new GetSavepointListArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of oceanus savepoint_list
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
        ///         var example = Output.Create(Tencentcloud.Oceanus.GetSavepointList.InvokeAsync(new Tencentcloud.Oceanus.GetSavepointListArgs
        ///         {
        ///             JobId = "cql-314rw6w0",
        ///             WorkSpaceId = "space-2idq8wbr",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSavepointListResult> Invoke(GetSavepointListInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSavepointListResult>("tencentcloud:Oceanus/getSavepointList:getSavepointList", args ?? new GetSavepointListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSavepointListArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Job SerialId.
        /// </summary>
        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Workspace SerialId.
        /// </summary>
        [Input("workSpaceId")]
        public string? WorkSpaceId { get; set; }

        public GetSavepointListArgs()
        {
        }
    }

    public sealed class GetSavepointListInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Job SerialId.
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Workspace SerialId.
        /// </summary>
        [Input("workSpaceId")]
        public Input<string>? WorkSpaceId { get; set; }

        public GetSavepointListInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSavepointListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string JobId;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Snapshot listNote: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSavepointListSavepointResult> Savepoints;
        public readonly string? WorkSpaceId;

        [OutputConstructor]
        private GetSavepointListResult(
            string id,

            string jobId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSavepointListSavepointResult> savepoints,

            string? workSpaceId)
        {
            Id = id;
            JobId = jobId;
            ResultOutputFile = resultOutputFile;
            Savepoints = savepoints;
            WorkSpaceId = workSpaceId;
        }
    }
}
