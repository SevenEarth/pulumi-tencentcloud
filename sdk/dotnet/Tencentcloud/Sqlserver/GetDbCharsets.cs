// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    public static class GetDbCharsets
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_d_b_charsets
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
        ///         var dbCharsets = Output.Create(Tencentcloud.Sqlserver.GetDbCharsets.InvokeAsync(new Tencentcloud.Sqlserver.GetDbCharsetsArgs
        ///         {
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDbCharsetsResult> InvokeAsync(GetDbCharsetsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbCharsetsResult>("tencentcloud:Sqlserver/getDbCharsets:getDbCharsets", args ?? new GetDbCharsetsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_d_b_charsets
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
        ///         var dbCharsets = Output.Create(Tencentcloud.Sqlserver.GetDbCharsets.InvokeAsync(new Tencentcloud.Sqlserver.GetDbCharsetsArgs
        ///         {
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDbCharsetsResult> Invoke(GetDbCharsetsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDbCharsetsResult>("tencentcloud:Sqlserver/getDbCharsets:getDbCharsets", args ?? new GetDbCharsetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDbCharsetsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance ID in the format of mssql-j8kv137v.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDbCharsetsArgs()
        {
        }
    }

    public sealed class GetDbCharsetsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance ID in the format of mssql-j8kv137v.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDbCharsetsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbCharsetsResult
    {
        /// <summary>
        /// Database character set list.
        /// </summary>
        public readonly ImmutableArray<string> DatabaseCharsets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDbCharsetsResult(
            ImmutableArray<string> databaseCharsets,

            string id,

            string instanceId,

            string? resultOutputFile)
        {
            DatabaseCharsets = databaseCharsets;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
