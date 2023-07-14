// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    public static class GetBackupDownloadUrls
    {
        /// <summary>
        /// Use this data source to query detailed information of postgresql backup_download_urls
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
        ///         var logBackups = Output.Create(Tencentcloud.Postgresql.GetLogBackups.InvokeAsync(new Tencentcloud.Postgresql.GetLogBackupsArgs
        ///         {
        ///             MinFinishTime = "%s",
        ///             MaxFinishTime = "%s",
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Postgresql.Inputs.GetLogBackupsFilterArgs
        ///                 {
        ///                     Name = "db-instance-id",
        ///                     Values = 
        ///                     {
        ///                         local.Pgsql_id,
        ///                     },
        ///                 },
        ///             },
        ///             OrderBy = "StartTime",
        ///             OrderByType = "desc",
        ///         }));
        ///         var backupDownloadUrls = logBackups.Apply(logBackups =&gt; Output.Create(Tencentcloud.Postgresql.GetBackupDownloadUrls.InvokeAsync(new Tencentcloud.Postgresql.GetBackupDownloadUrlsArgs
        ///         {
        ///             DbInstanceId = local.Pgsql_id,
        ///             BackupType = "LogBackup",
        ///             BackupId = logBackups.LogBackupSets?[0]?.Id,
        ///             UrlExpireTime = 12,
        ///             BackupDownloadRestriction = new Tencentcloud.Postgresql.Inputs.GetBackupDownloadUrlsBackupDownloadRestrictionArgs
        ///             {
        ///                 RestrictionType = "NONE",
        ///                 VpcRestrictionEffect = "ALLOW",
        ///                 VpcIdSets = 
        ///                 {
        ///                     local.Vpc_id,
        ///                 },
        ///                 IpRestrictionEffect = "ALLOW",
        ///                 IpSets = 
        ///                 {
        ///                     "0.0.0.0",
        ///                 },
        ///             },
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupDownloadUrlsResult> InvokeAsync(GetBackupDownloadUrlsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupDownloadUrlsResult>("tencentcloud:Postgresql/getBackupDownloadUrls:getBackupDownloadUrls", args ?? new GetBackupDownloadUrlsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of postgresql backup_download_urls
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
        ///         var logBackups = Output.Create(Tencentcloud.Postgresql.GetLogBackups.InvokeAsync(new Tencentcloud.Postgresql.GetLogBackupsArgs
        ///         {
        ///             MinFinishTime = "%s",
        ///             MaxFinishTime = "%s",
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Postgresql.Inputs.GetLogBackupsFilterArgs
        ///                 {
        ///                     Name = "db-instance-id",
        ///                     Values = 
        ///                     {
        ///                         local.Pgsql_id,
        ///                     },
        ///                 },
        ///             },
        ///             OrderBy = "StartTime",
        ///             OrderByType = "desc",
        ///         }));
        ///         var backupDownloadUrls = logBackups.Apply(logBackups =&gt; Output.Create(Tencentcloud.Postgresql.GetBackupDownloadUrls.InvokeAsync(new Tencentcloud.Postgresql.GetBackupDownloadUrlsArgs
        ///         {
        ///             DbInstanceId = local.Pgsql_id,
        ///             BackupType = "LogBackup",
        ///             BackupId = logBackups.LogBackupSets?[0]?.Id,
        ///             UrlExpireTime = 12,
        ///             BackupDownloadRestriction = new Tencentcloud.Postgresql.Inputs.GetBackupDownloadUrlsBackupDownloadRestrictionArgs
        ///             {
        ///                 RestrictionType = "NONE",
        ///                 VpcRestrictionEffect = "ALLOW",
        ///                 VpcIdSets = 
        ///                 {
        ///                     local.Vpc_id,
        ///                 },
        ///                 IpRestrictionEffect = "ALLOW",
        ///                 IpSets = 
        ///                 {
        ///                     "0.0.0.0",
        ///                 },
        ///             },
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackupDownloadUrlsResult> Invoke(GetBackupDownloadUrlsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupDownloadUrlsResult>("tencentcloud:Postgresql/getBackupDownloadUrls:getBackupDownloadUrls", args ?? new GetBackupDownloadUrlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupDownloadUrlsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup download restriction.
        /// </summary>
        [Input("backupDownloadRestriction")]
        public Inputs.GetBackupDownloadUrlsBackupDownloadRestrictionArgs? BackupDownloadRestriction { get; set; }

        /// <summary>
        /// Unique backup ID.
        /// </summary>
        [Input("backupId", required: true)]
        public string BackupId { get; set; } = null!;

        /// <summary>
        /// Backup type. Valid values: `LogBackup`, `BaseBackup`.
        /// </summary>
        [Input("backupType", required: true)]
        public string BackupType { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public string DbInstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Validity period of a URL, which is 12 hours by default.
        /// </summary>
        [Input("urlExpireTime")]
        public int? UrlExpireTime { get; set; }

        public GetBackupDownloadUrlsArgs()
        {
        }
    }

    public sealed class GetBackupDownloadUrlsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup download restriction.
        /// </summary>
        [Input("backupDownloadRestriction")]
        public Input<Inputs.GetBackupDownloadUrlsBackupDownloadRestrictionInputArgs>? BackupDownloadRestriction { get; set; }

        /// <summary>
        /// Unique backup ID.
        /// </summary>
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        /// <summary>
        /// Backup type. Valid values: `LogBackup`, `BaseBackup`.
        /// </summary>
        [Input("backupType", required: true)]
        public Input<string> BackupType { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Validity period of a URL, which is 12 hours by default.
        /// </summary>
        [Input("urlExpireTime")]
        public Input<int>? UrlExpireTime { get; set; }

        public GetBackupDownloadUrlsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupDownloadUrlsResult
    {
        public readonly Outputs.GetBackupDownloadUrlsBackupDownloadRestrictionResult? BackupDownloadRestriction;
        /// <summary>
        /// Backup download URL.
        /// </summary>
        public readonly string BackupDownloadUrl;
        public readonly string BackupId;
        public readonly string BackupType;
        public readonly string DbInstanceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int? UrlExpireTime;

        [OutputConstructor]
        private GetBackupDownloadUrlsResult(
            Outputs.GetBackupDownloadUrlsBackupDownloadRestrictionResult? backupDownloadRestriction,

            string backupDownloadUrl,

            string backupId,

            string backupType,

            string dbInstanceId,

            string id,

            string? resultOutputFile,

            int? urlExpireTime)
        {
            BackupDownloadRestriction = backupDownloadRestriction;
            BackupDownloadUrl = backupDownloadUrl;
            BackupId = backupId;
            BackupType = backupType;
            DbInstanceId = dbInstanceId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            UrlExpireTime = urlExpireTime;
        }
    }
}
