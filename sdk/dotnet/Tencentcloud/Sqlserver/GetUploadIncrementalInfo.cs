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
    public static class GetUploadIncrementalInfo
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver upload_incremental_info
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
        ///     var example = Tencentcloud.Sqlserver.GetUploadIncrementalInfo.Invoke(new()
        ///     {
        ///         BackupMigrationId = "mssql-backup-migration-83t5u3tv",
        ///         IncrementalMigrationId = "mssql-incremental-migration-h36gkdxn",
        ///         InstanceId = "mssql-4tgeyeeh",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUploadIncrementalInfoResult> InvokeAsync(GetUploadIncrementalInfoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUploadIncrementalInfoResult>("tencentcloud:Sqlserver/getUploadIncrementalInfo:getUploadIncrementalInfo", args ?? new GetUploadIncrementalInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver upload_incremental_info
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
        ///     var example = Tencentcloud.Sqlserver.GetUploadIncrementalInfo.Invoke(new()
        ///     {
        ///         BackupMigrationId = "mssql-backup-migration-83t5u3tv",
        ///         IncrementalMigrationId = "mssql-incremental-migration-h36gkdxn",
        ///         InstanceId = "mssql-4tgeyeeh",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUploadIncrementalInfoResult> Invoke(GetUploadIncrementalInfoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUploadIncrementalInfoResult>("tencentcloud:Sqlserver/getUploadIncrementalInfo:getUploadIncrementalInfo", args ?? new GetUploadIncrementalInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUploadIncrementalInfoArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Input("backupMigrationId", required: true)]
        public string BackupMigrationId { get; set; } = null!;

        /// <summary>
        /// ID of the incremental import task.
        /// </summary>
        [Input("incrementalMigrationId", required: true)]
        public string IncrementalMigrationId { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetUploadIncrementalInfoArgs()
        {
        }
        public static new GetUploadIncrementalInfoArgs Empty => new GetUploadIncrementalInfoArgs();
    }

    public sealed class GetUploadIncrementalInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Input("backupMigrationId", required: true)]
        public Input<string> BackupMigrationId { get; set; } = null!;

        /// <summary>
        /// ID of the incremental import task.
        /// </summary>
        [Input("incrementalMigrationId", required: true)]
        public Input<string> IncrementalMigrationId { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetUploadIncrementalInfoInvokeArgs()
        {
        }
        public static new GetUploadIncrementalInfoInvokeArgs Empty => new GetUploadIncrementalInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetUploadIncrementalInfoResult
    {
        public readonly string BackupMigrationId;
        /// <summary>
        /// Bucket name.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// Temporary key expiration time.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IncrementalMigrationId;
        public readonly string InstanceId;
        /// <summary>
        /// Storage path.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Bucket location information.
        /// </summary>
        public readonly string Region;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Temporary key start time.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Temporary key ID.
        /// </summary>
        public readonly string TmpSecretId;
        /// <summary>
        /// Temporary key (Key).
        /// </summary>
        public readonly string TmpSecretKey;
        /// <summary>
        /// Temporary key (Token).
        /// </summary>
        public readonly string XCosSecurityToken;

        [OutputConstructor]
        private GetUploadIncrementalInfoResult(
            string backupMigrationId,

            string bucketName,

            string expiredTime,

            string id,

            string incrementalMigrationId,

            string instanceId,

            string path,

            string region,

            string? resultOutputFile,

            string startTime,

            string tmpSecretId,

            string tmpSecretKey,

            string xCosSecurityToken)
        {
            BackupMigrationId = backupMigrationId;
            BucketName = bucketName;
            ExpiredTime = expiredTime;
            Id = id;
            IncrementalMigrationId = incrementalMigrationId;
            InstanceId = instanceId;
            Path = path;
            Region = region;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            TmpSecretId = tmpSecretId;
            TmpSecretKey = tmpSecretKey;
            XCosSecurityToken = xCosSecurityToken;
        }
    }
}
