// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain
{
    public static class GetSlowLogTopSqls
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_log_top_sqls
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
        ///         var test = Output.Create(Tencentcloud.Dbbrain.GetSlowLogTopSqls.InvokeAsync(new Tencentcloud.Dbbrain.GetSlowLogTopSqlsArgs
        ///         {
        ///             EndTime = "%s",
        ///             InstanceId = "%s",
        ///             OrderBy = "ASC",
        ///             Product = "mysql",
        ///             SortBy = "QueryTimeMax",
        ///             StartTime = "%s",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSlowLogTopSqlsResult> InvokeAsync(GetSlowLogTopSqlsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSlowLogTopSqlsResult>("tencentcloud:Dbbrain/getSlowLogTopSqls:getSlowLogTopSqls", args ?? new GetSlowLogTopSqlsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_log_top_sqls
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
        ///         var test = Output.Create(Tencentcloud.Dbbrain.GetSlowLogTopSqls.InvokeAsync(new Tencentcloud.Dbbrain.GetSlowLogTopSqlsArgs
        ///         {
        ///             EndTime = "%s",
        ///             InstanceId = "%s",
        ///             OrderBy = "ASC",
        ///             Product = "mysql",
        ///             SortBy = "QueryTimeMax",
        ///             StartTime = "%s",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSlowLogTopSqlsResult> Invoke(GetSlowLogTopSqlsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSlowLogTopSqlsResult>("tencentcloud:Dbbrain/getSlowLogTopSqls:getSlowLogTopSqls", args ?? new GetSlowLogTopSqlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowLogTopSqlsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The deadline, such as `2019-09-11 10:13:14`, the interval between the deadline and the start time is less than 7 days.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// The sorting method supports ASC (ascending) and DESC (descending). The default is DESC.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public string? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("schemaLists")]
        private List<Inputs.GetSlowLogTopSqlsSchemaListArgs>? _schemaLists;

        /// <summary>
        /// Array of database names.
        /// </summary>
        public List<Inputs.GetSlowLogTopSqlsSchemaListArgs> SchemaLists
        {
            get => _schemaLists ?? (_schemaLists = new List<Inputs.GetSlowLogTopSqlsSchemaListArgs>());
            set => _schemaLists = value;
        }

        /// <summary>
        /// Sort key, currently supports sort keys such as QueryTime, ExecTimes, RowsSent, LockTime and RowsExamined, the default is QueryTime.
        /// </summary>
        [Input("sortBy")]
        public string? SortBy { get; set; }

        /// <summary>
        /// Start time, such as `2019-09-10 12:13:14`.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetSlowLogTopSqlsArgs()
        {
        }
    }

    public sealed class GetSlowLogTopSqlsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The deadline, such as `2019-09-11 10:13:14`, the interval between the deadline and the start time is less than 7 days.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The sorting method supports ASC (ascending) and DESC (descending). The default is DESC.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("schemaLists")]
        private InputList<Inputs.GetSlowLogTopSqlsSchemaListInputArgs>? _schemaLists;

        /// <summary>
        /// Array of database names.
        /// </summary>
        public InputList<Inputs.GetSlowLogTopSqlsSchemaListInputArgs> SchemaLists
        {
            get => _schemaLists ?? (_schemaLists = new InputList<Inputs.GetSlowLogTopSqlsSchemaListInputArgs>());
            set => _schemaLists = value;
        }

        /// <summary>
        /// Sort key, currently supports sort keys such as QueryTime, ExecTimes, RowsSent, LockTime and RowsExamined, the default is QueryTime.
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// Start time, such as `2019-09-10 12:13:14`.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetSlowLogTopSqlsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSlowLogTopSqlsResult
    {
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? OrderBy;
        public readonly string? Product;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Slow log top sql list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlowLogTopSqlsRowResult> Rows;
        public readonly ImmutableArray<Outputs.GetSlowLogTopSqlsSchemaListResult> SchemaLists;
        public readonly string? SortBy;
        public readonly string StartTime;

        [OutputConstructor]
        private GetSlowLogTopSqlsResult(
            string endTime,

            string id,

            string instanceId,

            string? orderBy,

            string? product,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSlowLogTopSqlsRowResult> rows,

            ImmutableArray<Outputs.GetSlowLogTopSqlsSchemaListResult> schemaLists,

            string? sortBy,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            OrderBy = orderBy;
            Product = product;
            ResultOutputFile = resultOutputFile;
            Rows = rows;
            SchemaLists = schemaLists;
            SortBy = sortBy;
            StartTime = startTime;
        }
    }
}
