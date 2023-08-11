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
    public static class GetBackupCommands
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_backup_command
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
        ///         var example = Output.Create(Tencentcloud.Sqlserver.GetBackupCommands.InvokeAsync(new Tencentcloud.Sqlserver.GetBackupCommandsArgs
        ///         {
        ///             BackupFileType = "FULL",
        ///             DataBaseName = "keep-publish-instance",
        ///             IsRecovery = "NO",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupCommandsResult> InvokeAsync(GetBackupCommandsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupCommandsResult>("tencentcloud:Sqlserver/getBackupCommands:getBackupCommands", args ?? new GetBackupCommandsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_backup_command
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
        ///         var example = Output.Create(Tencentcloud.Sqlserver.GetBackupCommands.InvokeAsync(new Tencentcloud.Sqlserver.GetBackupCommandsArgs
        ///         {
        ///             BackupFileType = "FULL",
        ///             DataBaseName = "keep-publish-instance",
        ///             IsRecovery = "NO",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackupCommandsResult> Invoke(GetBackupCommandsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupCommandsResult>("tencentcloud:Sqlserver/getBackupCommands:getBackupCommands", args ?? new GetBackupCommandsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupCommandsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
        /// </summary>
        [Input("backupFileType", required: true)]
        public string BackupFileType { get; set; } = null!;

        /// <summary>
        /// Database name.
        /// </summary>
        [Input("dataBaseName", required: true)]
        public string DataBaseName { get; set; } = null!;

        /// <summary>
        /// Whether restoration is required. No: not required. Yes: required.
        /// </summary>
        [Input("isRecovery", required: true)]
        public string IsRecovery { get; set; } = null!;

        /// <summary>
        /// Storage path of backup files. If this parameter is left empty, the default storage path will be D:.
        /// </summary>
        [Input("localPath")]
        public string? LocalPath { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBackupCommandsArgs()
        {
        }
    }

    public sealed class GetBackupCommandsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
        /// </summary>
        [Input("backupFileType", required: true)]
        public Input<string> BackupFileType { get; set; } = null!;

        /// <summary>
        /// Database name.
        /// </summary>
        [Input("dataBaseName", required: true)]
        public Input<string> DataBaseName { get; set; } = null!;

        /// <summary>
        /// Whether restoration is required. No: not required. Yes: required.
        /// </summary>
        [Input("isRecovery", required: true)]
        public Input<string> IsRecovery { get; set; } = null!;

        /// <summary>
        /// Storage path of backup files. If this parameter is left empty, the default storage path will be D:.
        /// </summary>
        [Input("localPath")]
        public Input<string>? LocalPath { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBackupCommandsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupCommandsResult
    {
        public readonly string BackupFileType;
        public readonly string DataBaseName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IsRecovery;
        /// <summary>
        /// Command list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackupCommandsListResult> Lists;
        public readonly string? LocalPath;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBackupCommandsResult(
            string backupFileType,

            string dataBaseName,

            string id,

            string isRecovery,

            ImmutableArray<Outputs.GetBackupCommandsListResult> lists,

            string? localPath,

            string? resultOutputFile)
        {
            BackupFileType = backupFileType;
            DataBaseName = dataBaseName;
            Id = id;
            IsRecovery = isRecovery;
            Lists = lists;
            LocalPath = localPath;
            ResultOutputFile = resultOutputFile;
        }
    }
}
