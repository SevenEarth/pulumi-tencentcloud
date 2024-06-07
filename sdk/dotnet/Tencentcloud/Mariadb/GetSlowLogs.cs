// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    public static class GetSlowLogs
    {
        /// <summary>
        /// Use this data source to query detailed information of mariadb slow_logs
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
        ///     var slowLogs = Tencentcloud.Mariadb.GetSlowLogs.Invoke(new()
        ///     {
        ///         InstanceId = "tdsql-9vqvls95",
        ///         OrderBy = "query_time_sum",
        ///         OrderByType = "desc",
        ///         Slave = 0,
        ///         StartTime = "2023-06-01 14:55:20",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSlowLogsResult> InvokeAsync(GetSlowLogsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSlowLogsResult>("tencentcloud:Mariadb/getSlowLogs:getSlowLogs", args ?? new GetSlowLogsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mariadb slow_logs
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
        ///     var slowLogs = Tencentcloud.Mariadb.GetSlowLogs.Invoke(new()
        ///     {
        ///         InstanceId = "tdsql-9vqvls95",
        ///         OrderBy = "query_time_sum",
        ///         OrderByType = "desc",
        ///         Slave = 0,
        ///         StartTime = "2023-06-01 14:55:20",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSlowLogsResult> Invoke(GetSlowLogsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSlowLogsResult>("tencentcloud:Mariadb/getSlowLogs:getSlowLogs", args ?? new GetSlowLogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowLogsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specific name of the database to be queried.
        /// </summary>
        [Input("db")]
        public string? Db { get; set; }

        /// <summary>
        /// Query end time in the format of 2016-08-22 14:55:20.
        /// </summary>
        [Input("endTime")]
        public string? EndTime { get; set; }

        /// <summary>
        /// Instance ID in the format of `tdsql-ow728lmc`.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Sorting metric. Valid values: query_time_sum, query_count.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Sorting order. Valid values: desc, asc.
        /// </summary>
        [Input("orderByType")]
        public string? OrderByType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica).
        /// </summary>
        [Input("slave")]
        public int? Slave { get; set; }

        /// <summary>
        /// Query start time in the format of 2016-07-23 14:55:20.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetSlowLogsArgs()
        {
        }
        public static new GetSlowLogsArgs Empty => new GetSlowLogsArgs();
    }

    public sealed class GetSlowLogsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specific name of the database to be queried.
        /// </summary>
        [Input("db")]
        public Input<string>? Db { get; set; }

        /// <summary>
        /// Query end time in the format of 2016-08-22 14:55:20.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Instance ID in the format of `tdsql-ow728lmc`.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Sorting metric. Valid values: query_time_sum, query_count.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Sorting order. Valid values: desc, asc.
        /// </summary>
        [Input("orderByType")]
        public Input<string>? OrderByType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica).
        /// </summary>
        [Input("slave")]
        public Input<int>? Slave { get; set; }

        /// <summary>
        /// Query start time in the format of 2016-07-23 14:55:20.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetSlowLogsInvokeArgs()
        {
        }
        public static new GetSlowLogsInvokeArgs Empty => new GetSlowLogsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSlowLogsResult
    {
        /// <summary>
        /// Slow query log data.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlowLogsDataResult> Datas;
        /// <summary>
        /// Database name.
        /// </summary>
        public readonly string? Db;
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Total statement lock time.
        /// </summary>
        public readonly double LockTimeSum;
        public readonly string? OrderBy;
        public readonly string? OrderByType;
        /// <summary>
        /// Total number of statement queries.
        /// </summary>
        public readonly int QueryCount;
        /// <summary>
        /// Total statement query time.
        /// </summary>
        public readonly double QueryTimeSum;
        public readonly string? ResultOutputFile;
        public readonly int? Slave;
        public readonly string StartTime;

        [OutputConstructor]
        private GetSlowLogsResult(
            ImmutableArray<Outputs.GetSlowLogsDataResult> datas,

            string? db,

            string? endTime,

            string id,

            string instanceId,

            double lockTimeSum,

            string? orderBy,

            string? orderByType,

            int queryCount,

            double queryTimeSum,

            string? resultOutputFile,

            int? slave,

            string startTime)
        {
            Datas = datas;
            Db = db;
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            LockTimeSum = lockTimeSum;
            OrderBy = orderBy;
            OrderByType = orderByType;
            QueryCount = queryCount;
            QueryTimeSum = queryTimeSum;
            ResultOutputFile = resultOutputFile;
            Slave = slave;
            StartTime = startTime;
        }
    }
}
