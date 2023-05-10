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
    public static class GetSecurityAuditLogDownloadUrls
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain security_audit_log_download_urls
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var task = new Tencentcloud.Dbbrain.SecurityAuditLogExportTask("task", new Tencentcloud.Dbbrain.SecurityAuditLogExportTaskArgs
        ///         {
        ///             SecAuditGroupId = "%s",
        ///             StartTime = "%s",
        ///             EndTime = "%s",
        ///             Product = "mysql",
        ///             DangerLevels = 
        ///             {
        ///                 0,
        ///                 1,
        ///                 2,
        ///             },
        ///         });
        ///         var test = Tencentcloud.Dbbrain.GetSecurityAuditLogDownloadUrls.Invoke(new Tencentcloud.Dbbrain.GetSecurityAuditLogDownloadUrlsInvokeArgs
        ///         {
        ///             SecAuditGroupId = "%s",
        ///             AsyncRequestId = task.AsyncRequestId,
        ///             Product = "mysql",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecurityAuditLogDownloadUrlsResult> InvokeAsync(GetSecurityAuditLogDownloadUrlsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecurityAuditLogDownloadUrlsResult>("tencentcloud:Dbbrain/getSecurityAuditLogDownloadUrls:getSecurityAuditLogDownloadUrls", args ?? new GetSecurityAuditLogDownloadUrlsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain security_audit_log_download_urls
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var task = new Tencentcloud.Dbbrain.SecurityAuditLogExportTask("task", new Tencentcloud.Dbbrain.SecurityAuditLogExportTaskArgs
        ///         {
        ///             SecAuditGroupId = "%s",
        ///             StartTime = "%s",
        ///             EndTime = "%s",
        ///             Product = "mysql",
        ///             DangerLevels = 
        ///             {
        ///                 0,
        ///                 1,
        ///                 2,
        ///             },
        ///         });
        ///         var test = Tencentcloud.Dbbrain.GetSecurityAuditLogDownloadUrls.Invoke(new Tencentcloud.Dbbrain.GetSecurityAuditLogDownloadUrlsInvokeArgs
        ///         {
        ///             SecAuditGroupId = "%s",
        ///             AsyncRequestId = task.AsyncRequestId,
        ///             Product = "mysql",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecurityAuditLogDownloadUrlsResult> Invoke(GetSecurityAuditLogDownloadUrlsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecurityAuditLogDownloadUrlsResult>("tencentcloud:Dbbrain/getSecurityAuditLogDownloadUrls:getSecurityAuditLogDownloadUrls", args ?? new GetSecurityAuditLogDownloadUrlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityAuditLogDownloadUrlsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Asynchronous task ID.
        /// </summary>
        [Input("asyncRequestId", required: true)]
        public int AsyncRequestId { get; set; }

        /// <summary>
        /// Service product type, supported values: `mysql` - ApsaraDB for MySQL.
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Security audit group Id.
        /// </summary>
        [Input("secAuditGroupId", required: true)]
        public string SecAuditGroupId { get; set; } = null!;

        public GetSecurityAuditLogDownloadUrlsArgs()
        {
        }
    }

    public sealed class GetSecurityAuditLogDownloadUrlsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Asynchronous task ID.
        /// </summary>
        [Input("asyncRequestId", required: true)]
        public Input<int> AsyncRequestId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values: `mysql` - ApsaraDB for MySQL.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Security audit group Id.
        /// </summary>
        [Input("secAuditGroupId", required: true)]
        public Input<string> SecAuditGroupId { get; set; } = null!;

        public GetSecurityAuditLogDownloadUrlsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecurityAuditLogDownloadUrlsResult
    {
        public readonly int AsyncRequestId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Product;
        public readonly string? ResultOutputFile;
        public readonly string SecAuditGroupId;
        /// <summary>
        /// List of COS links to export results. When the result set is large, it may be divided into multiple urls for download.
        /// </summary>
        public readonly ImmutableArray<string> Urls;

        [OutputConstructor]
        private GetSecurityAuditLogDownloadUrlsResult(
            int asyncRequestId,

            string id,

            string product,

            string? resultOutputFile,

            string secAuditGroupId,

            ImmutableArray<string> urls)
        {
            AsyncRequestId = asyncRequestId;
            Id = id;
            Product = product;
            ResultOutputFile = resultOutputFile;
            SecAuditGroupId = secAuditGroupId;
            Urls = urls;
        }
    }
}
