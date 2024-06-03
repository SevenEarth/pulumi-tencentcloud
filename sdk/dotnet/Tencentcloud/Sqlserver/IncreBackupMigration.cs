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
    /// <summary>
    /// Provides a resource to create a sqlserver incre_backup_migration
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Tencentcloud.Sqlserver.IncreBackupMigration("example", new()
    ///     {
    ///         BackupFiles = new[] {},
    ///         BackupMigrationId = "mssql-backup-migration-9tj0sxnz",
    ///         InstanceId = "mssql-4gmc5805",
    ///         IsRecovery = "YES",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// sqlserver incre_backup_migration can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Sqlserver/increBackupMigration:IncreBackupMigration incre_backup_migration incre_backup_migration_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/increBackupMigration:IncreBackupMigration")]
    public partial class IncreBackupMigration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Incremental backup file. If the UploadType of a full backup file is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
        /// </summary>
        [Output("backupFiles")]
        public Output<ImmutableArray<string>> BackupFiles { get; private set; } = null!;

        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Output("backupMigrationId")]
        public Output<string> BackupMigrationId { get; private set; } = null!;

        /// <summary>
        /// Incremental import task ID.
        /// </summary>
        [Output("incrementalMigrationId")]
        public Output<string> IncrementalMigrationId { get; private set; } = null!;

        /// <summary>
        /// ID of imported target instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Whether restoration is required. No: not required. Yes: required. Not required by default.
        /// </summary>
        [Output("isRecovery")]
        public Output<string?> IsRecovery { get; private set; } = null!;


        /// <summary>
        /// Create a IncreBackupMigration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IncreBackupMigration(string name, IncreBackupMigrationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/increBackupMigration:IncreBackupMigration", name, args ?? new IncreBackupMigrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IncreBackupMigration(string name, Input<string> id, IncreBackupMigrationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/increBackupMigration:IncreBackupMigration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IncreBackupMigration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IncreBackupMigration Get(string name, Input<string> id, IncreBackupMigrationState? state = null, CustomResourceOptions? options = null)
        {
            return new IncreBackupMigration(name, id, state, options);
        }
    }

    public sealed class IncreBackupMigrationArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupFiles")]
        private InputList<string>? _backupFiles;

        /// <summary>
        /// Incremental backup file. If the UploadType of a full backup file is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
        /// </summary>
        public InputList<string> BackupFiles
        {
            get => _backupFiles ?? (_backupFiles = new InputList<string>());
            set => _backupFiles = value;
        }

        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Input("backupMigrationId", required: true)]
        public Input<string> BackupMigrationId { get; set; } = null!;

        /// <summary>
        /// ID of imported target instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Whether restoration is required. No: not required. Yes: required. Not required by default.
        /// </summary>
        [Input("isRecovery")]
        public Input<string>? IsRecovery { get; set; }

        public IncreBackupMigrationArgs()
        {
        }
        public static new IncreBackupMigrationArgs Empty => new IncreBackupMigrationArgs();
    }

    public sealed class IncreBackupMigrationState : global::Pulumi.ResourceArgs
    {
        [Input("backupFiles")]
        private InputList<string>? _backupFiles;

        /// <summary>
        /// Incremental backup file. If the UploadType of a full backup file is COS_URL, fill in URL here. If the UploadType is COS_UPLOAD, fill in the name of the backup file here. Only 1 backup file is supported, but a backup file can involve multiple databases.
        /// </summary>
        public InputList<string> BackupFiles
        {
            get => _backupFiles ?? (_backupFiles = new InputList<string>());
            set => _backupFiles = value;
        }

        /// <summary>
        /// Backup import task ID, which is returned through the API CreateBackupMigration.
        /// </summary>
        [Input("backupMigrationId")]
        public Input<string>? BackupMigrationId { get; set; }

        /// <summary>
        /// Incremental import task ID.
        /// </summary>
        [Input("incrementalMigrationId")]
        public Input<string>? IncrementalMigrationId { get; set; }

        /// <summary>
        /// ID of imported target instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Whether restoration is required. No: not required. Yes: required. Not required by default.
        /// </summary>
        [Input("isRecovery")]
        public Input<string>? IsRecovery { get; set; }

        public IncreBackupMigrationState()
        {
        }
        public static new IncreBackupMigrationState Empty => new IncreBackupMigrationState();
    }
}
