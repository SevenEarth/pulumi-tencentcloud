// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver
{
    public static class GetBackups
    {
        /// <summary>
        /// Use this data source to query the list of SQL Server backups.
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
        ///         var foo = Output.Create(Tencentcloud.Sqlserver.GetBackups.InvokeAsync(new Tencentcloud.Sqlserver.GetBackupsArgs
        ///         {
        ///             EndTime = "2020-06-22 00:00:00",
        ///             InstanceId = "mssql-3cdq7kx5",
        ///             StartTime = "2020-06-17 00:00:00",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupsResult> InvokeAsync(GetBackupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupsResult>("tencentcloud:Sqlserver/getBackups:getBackups", args ?? new GetBackupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query the list of SQL Server backups.
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
        ///         var foo = Output.Create(Tencentcloud.Sqlserver.GetBackups.InvokeAsync(new Tencentcloud.Sqlserver.GetBackupsArgs
        ///         {
        ///             EndTime = "2020-06-22 00:00:00",
        ///             InstanceId = "mssql-3cdq7kx5",
        ///             StartTime = "2020-06-17 00:00:00",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackupsResult> Invoke(GetBackupsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupsResult>("tencentcloud:Sqlserver/getBackups:getBackups", args ?? new GetBackupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// End time of the instance list, like yyyy-MM-dd HH:mm:ss.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time of the instance list, like yyyy-MM-dd HH:mm:ss.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetBackupsArgs()
        {
        }
    }

    public sealed class GetBackupsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// End time of the instance list, like yyyy-MM-dd HH:mm:ss.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time of the instance list, like yyyy-MM-dd HH:mm:ss.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetBackupsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupsResult
    {
        /// <summary>
        /// End time of the backup.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// A list of SQL Server backup. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackupsListResult> Lists;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Start time of the backup.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GetBackupsResult(
            string endTime,

            string id,

            string instanceId,

            ImmutableArray<Outputs.GetBackupsListResult> lists,

            string? resultOutputFile,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}
