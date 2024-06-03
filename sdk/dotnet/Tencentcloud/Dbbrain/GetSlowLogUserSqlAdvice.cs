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
    public static class GetSlowLogUserSqlAdvice
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_log_user_sql_advice
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
        ///     var test = Tencentcloud.Dbbrain.GetSlowLogUserSqlAdvice.Invoke(new()
        ///     {
        ///         InstanceId = "%s",
        ///         Product = "mysql",
        ///         SqlText = "%s",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSlowLogUserSqlAdviceResult> InvokeAsync(GetSlowLogUserSqlAdviceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSlowLogUserSqlAdviceResult>("tencentcloud:Dbbrain/getSlowLogUserSqlAdvice:getSlowLogUserSqlAdvice", args ?? new GetSlowLogUserSqlAdviceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_log_user_sql_advice
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
        ///     var test = Tencentcloud.Dbbrain.GetSlowLogUserSqlAdvice.Invoke(new()
        ///     {
        ///         InstanceId = "%s",
        ///         Product = "mysql",
        ///         SqlText = "%s",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSlowLogUserSqlAdviceResult> Invoke(GetSlowLogUserSqlAdviceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSlowLogUserSqlAdviceResult>("tencentcloud:Dbbrain/getSlowLogUserSqlAdvice:getSlowLogUserSqlAdvice", args ?? new GetSlowLogUserSqlAdviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowLogUserSqlAdviceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values: `mysql` - cloud database MySQL; `cynosdb` - cloud database TDSQL-C for MySQL; `dbbrain-mysql` - self-built MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public string? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// library name.
        /// </summary>
        [Input("schema")]
        public string? Schema { get; set; }

        /// <summary>
        /// SQL statements.
        /// </summary>
        [Input("sqlText", required: true)]
        public string SqlText { get; set; } = null!;

        public GetSlowLogUserSqlAdviceArgs()
        {
        }
        public static new GetSlowLogUserSqlAdviceArgs Empty => new GetSlowLogUserSqlAdviceArgs();
    }

    public sealed class GetSlowLogUserSqlAdviceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values: `mysql` - cloud database MySQL; `cynosdb` - cloud database TDSQL-C for MySQL; `dbbrain-mysql` - self-built MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// library name.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// SQL statements.
        /// </summary>
        [Input("sqlText", required: true)]
        public Input<string> SqlText { get; set; } = null!;

        public GetSlowLogUserSqlAdviceInvokeArgs()
        {
        }
        public static new GetSlowLogUserSqlAdviceInvokeArgs Empty => new GetSlowLogUserSqlAdviceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSlowLogUserSqlAdviceResult
    {
        /// <summary>
        /// SQL optimization suggestion, which can be parsed into a JSON array, and the output is empty when no optimization is required.
        /// </summary>
        public readonly string Advices;
        /// <summary>
        /// SQL optimization suggestion remarks, which can be parsed into a String array, and the output is empty when optimization is not required.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// The cost saving details after SQL optimization can be parsed as JSON, and the output is empty when no optimization is required.
        /// </summary>
        public readonly string Cost;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? Product;
        /// <summary>
        /// Unique request ID, returned for every request. The RequestId of the request needs to be provided when locating the problem.
        /// </summary>
        public readonly string RequestId;
        public readonly string? ResultOutputFile;
        public readonly string Schema;
        /// <summary>
        /// The SQL execution plan can be parsed into JSON, and the output is empty when no optimization is required.
        /// </summary>
        public readonly string SqlPlan;
        public readonly string SqlText;
        /// <summary>
        /// The DDL information of related tables can be parsed into a JSON array.
        /// </summary>
        public readonly string Tables;

        [OutputConstructor]
        private GetSlowLogUserSqlAdviceResult(
            string advices,

            string comments,

            string cost,

            string id,

            string instanceId,

            string? product,

            string requestId,

            string? resultOutputFile,

            string schema,

            string sqlPlan,

            string sqlText,

            string tables)
        {
            Advices = advices;
            Comments = comments;
            Cost = cost;
            Id = id;
            InstanceId = instanceId;
            Product = product;
            RequestId = requestId;
            ResultOutputFile = resultOutputFile;
            Schema = schema;
            SqlPlan = sqlPlan;
            SqlText = sqlText;
            Tables = tables;
        }
    }
}
