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
    public static class GetDatabaseTable
    {
        /// <summary>
        /// Use this data source to query detailed information of mariadb database_table
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
        ///     var databaseTable = Tencentcloud.Mariadb.GetDatabaseTable.Invoke(new()
        ///     {
        ///         DbName = "mysql",
        ///         InstanceId = "tdsql-e9tklsgz",
        ///         Table = "server_cost",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDatabaseTableResult> InvokeAsync(GetDatabaseTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseTableResult>("tencentcloud:Mariadb/getDatabaseTable:getDatabaseTable", args ?? new GetDatabaseTableArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mariadb database_table
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
        ///     var databaseTable = Tencentcloud.Mariadb.GetDatabaseTable.Invoke(new()
        ///     {
        ///         DbName = "mysql",
        ///         InstanceId = "tdsql-e9tklsgz",
        ///         Table = "server_cost",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDatabaseTableResult> Invoke(GetDatabaseTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseTableResult>("tencentcloud:Mariadb/getDatabaseTable:getDatabaseTable", args ?? new GetDatabaseTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseTableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// database name.
        /// </summary>
        [Input("dbName", required: true)]
        public string DbName { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// table name.
        /// </summary>
        [Input("table", required: true)]
        public string Table { get; set; } = null!;

        public GetDatabaseTableArgs()
        {
        }
        public static new GetDatabaseTableArgs Empty => new GetDatabaseTableArgs();
    }

    public sealed class GetDatabaseTableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// database name.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// table name.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public GetDatabaseTableInvokeArgs()
        {
        }
        public static new GetDatabaseTableInvokeArgs Empty => new GetDatabaseTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseTableResult
    {
        /// <summary>
        /// column list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseTableColResult> Cols;
        public readonly string DbName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly string Table;

        [OutputConstructor]
        private GetDatabaseTableResult(
            ImmutableArray<Outputs.GetDatabaseTableColResult> cols,

            string dbName,

            string id,

            string instanceId,

            string? resultOutputFile,

            string table)
        {
            Cols = cols;
            DbName = dbName;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            Table = table;
        }
    }
}
