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
    /// Provides a resource to create a sqlserver migration
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var srcAccount = new Tencentcloud.Sqlserver.Account("srcAccount", new Tencentcloud.Sqlserver.AccountArgs
    ///         {
    ///             InstanceId = local.Sqlserver_id,
    ///             Password = "password",
    ///             IsAdmin = true,
    ///         });
    ///         var srcAccountDbAttachment = new Tencentcloud.Sqlserver.AccountDbAttachment("srcAccountDbAttachment", new Tencentcloud.Sqlserver.AccountDbAttachmentArgs
    ///         {
    ///             InstanceId = local.Sqlserver_id,
    ///             AccountName = srcAccount.Name,
    ///             DbName = local.Sqlserver_db,
    ///             Privilege = "ReadWrite",
    ///         });
    ///         var dstInstance = new Tencentcloud.Sqlserver.Instance("dstInstance", new Tencentcloud.Sqlserver.InstanceArgs
    ///         {
    ///             AvailabilityZone = @var.Default_az,
    ///             ChargeType = "POSTPAID_BY_HOUR",
    ///             VpcId = local.Vpc_id,
    ///             SubnetId = local.Subnet_id,
    ///             SecurityGroups = 
    ///             {
    ///                 local.Sg_id,
    ///             },
    ///             ProjectId = 0,
    ///             Memory = 2,
    ///             Storage = 10,
    ///             MaintenanceWeekSets = 
    ///             {
    ///                 1,
    ///                 2,
    ///                 3,
    ///             },
    ///             MaintenanceStartTime = "09:00",
    ///             MaintenanceTimeSpan = 3,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///         });
    ///         var dstAccount = new Tencentcloud.Sqlserver.Account("dstAccount", new Tencentcloud.Sqlserver.AccountArgs
    ///         {
    ///             InstanceId = dstInstance.Id,
    ///             Password = "password",
    ///             IsAdmin = true,
    ///         });
    ///         var dstDb = new Tencentcloud.Sqlserver.Db("dstDb", new Tencentcloud.Sqlserver.DbArgs
    ///         {
    ///             InstanceId = dstInstance.Id,
    ///             Charset = "Chinese_PRC_BIN",
    ///             Remark = "testACC-remark",
    ///         });
    ///         var migration = new Tencentcloud.Sqlserver.Migration("migration", new Tencentcloud.Sqlserver.MigrationArgs
    ///         {
    ///             MigrateName = "tf_test_migration",
    ///             MigrateType = 1,
    ///             SourceType = 1,
    ///             Source = new Tencentcloud.Sqlserver.Inputs.MigrationSourceArgs
    ///             {
    ///                 InstanceId = local.Sqlserver_id,
    ///                 UserName = srcAccount.Name,
    ///                 Password = srcAccount.Password,
    ///             },
    ///             Target = new Tencentcloud.Sqlserver.Inputs.MigrationTargetArgs
    ///             {
    ///                 InstanceId = dstInstance.Id,
    ///                 UserName = dstAccount.Name,
    ///                 Password = dstAccount.Password,
    ///             },
    ///             MigrateDbSets = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.MigrationMigrateDbSetArgs
    ///                 {
    ///                     DbName = local.Sqlserver_db,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver migration can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/migration:Migration migration migration_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/migration:Migration")]
    public partial class Migration : Pulumi.CustomResource
    {
        /// <summary>
        /// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
        /// </summary>
        [Output("migrateDbSets")]
        public Output<ImmutableArray<Outputs.MigrationMigrateDbSet>> MigrateDbSets { get; private set; } = null!;

        /// <summary>
        /// Name of the migration task.
        /// </summary>
        [Output("migrateName")]
        public Output<string> MigrateName { get; private set; } = null!;

        /// <summary>
        /// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
        /// </summary>
        [Output("migrateType")]
        public Output<int> MigrateType { get; private set; } = null!;

        /// <summary>
        /// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
        /// </summary>
        [Output("renameRestores")]
        public Output<ImmutableArray<Outputs.MigrationRenameRestore>> RenameRestores { get; private set; } = null!;

        /// <summary>
        /// Migration source.
        /// </summary>
        [Output("source")]
        public Output<Outputs.MigrationSource> Source { get; private set; } = null!;

        /// <summary>
        /// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
        /// </summary>
        [Output("sourceType")]
        public Output<int> SourceType { get; private set; } = null!;

        /// <summary>
        /// Migration target.
        /// </summary>
        [Output("target")]
        public Output<Outputs.MigrationTarget> Target { get; private set; } = null!;


        /// <summary>
        /// Create a Migration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Migration(string name, MigrationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/migration:Migration", name, args ?? new MigrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Migration(string name, Input<string> id, MigrationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/migration:Migration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Migration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Migration Get(string name, Input<string> id, MigrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Migration(name, id, state, options);
        }
    }

    public sealed class MigrationArgs : Pulumi.ResourceArgs
    {
        [Input("migrateDbSets")]
        private InputList<Inputs.MigrationMigrateDbSetArgs>? _migrateDbSets;

        /// <summary>
        /// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
        /// </summary>
        public InputList<Inputs.MigrationMigrateDbSetArgs> MigrateDbSets
        {
            get => _migrateDbSets ?? (_migrateDbSets = new InputList<Inputs.MigrationMigrateDbSetArgs>());
            set => _migrateDbSets = value;
        }

        /// <summary>
        /// Name of the migration task.
        /// </summary>
        [Input("migrateName", required: true)]
        public Input<string> MigrateName { get; set; } = null!;

        /// <summary>
        /// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
        /// </summary>
        [Input("migrateType", required: true)]
        public Input<int> MigrateType { get; set; } = null!;

        [Input("renameRestores")]
        private InputList<Inputs.MigrationRenameRestoreArgs>? _renameRestores;

        /// <summary>
        /// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
        /// </summary>
        public InputList<Inputs.MigrationRenameRestoreArgs> RenameRestores
        {
            get => _renameRestores ?? (_renameRestores = new InputList<Inputs.MigrationRenameRestoreArgs>());
            set => _renameRestores = value;
        }

        /// <summary>
        /// Migration source.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.MigrationSourceArgs> Source { get; set; } = null!;

        /// <summary>
        /// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
        /// </summary>
        [Input("sourceType", required: true)]
        public Input<int> SourceType { get; set; } = null!;

        /// <summary>
        /// Migration target.
        /// </summary>
        [Input("target", required: true)]
        public Input<Inputs.MigrationTargetArgs> Target { get; set; } = null!;

        public MigrationArgs()
        {
        }
    }

    public sealed class MigrationState : Pulumi.ResourceArgs
    {
        [Input("migrateDbSets")]
        private InputList<Inputs.MigrationMigrateDbSetGetArgs>? _migrateDbSets;

        /// <summary>
        /// Migrate DB objects. Offline migration is not used (SourceType=4 or SourceType=5).
        /// </summary>
        public InputList<Inputs.MigrationMigrateDbSetGetArgs> MigrateDbSets
        {
            get => _migrateDbSets ?? (_migrateDbSets = new InputList<Inputs.MigrationMigrateDbSetGetArgs>());
            set => _migrateDbSets = value;
        }

        /// <summary>
        /// Name of the migration task.
        /// </summary>
        [Input("migrateName")]
        public Input<string>? MigrateName { get; set; }

        /// <summary>
        /// Migration type (1 structure migration 2 data migration 3 incremental synchronization).
        /// </summary>
        [Input("migrateType")]
        public Input<int>? MigrateType { get; set; }

        [Input("renameRestores")]
        private InputList<Inputs.MigrationRenameRestoreGetArgs>? _renameRestores;

        /// <summary>
        /// Restore and rename the database in ReNameRestoreDatabase. If it is not filled in, the restored database will be named by default and all databases will be restored. Valid if SourceType=5.
        /// </summary>
        public InputList<Inputs.MigrationRenameRestoreGetArgs> RenameRestores
        {
            get => _renameRestores ?? (_renameRestores = new InputList<Inputs.MigrationRenameRestoreGetArgs>());
            set => _renameRestores = value;
        }

        /// <summary>
        /// Migration source.
        /// </summary>
        [Input("source")]
        public Input<Inputs.MigrationSourceGetArgs>? Source { get; set; }

        /// <summary>
        /// Type of migration source 1 TencentDB for SQLServer 2 Cloud server self-built SQLServer database 4 SQLServer backup and restore 5 SQLServer backup and restore (COS mode).
        /// </summary>
        [Input("sourceType")]
        public Input<int>? SourceType { get; set; }

        /// <summary>
        /// Migration target.
        /// </summary>
        [Input("target")]
        public Input<Inputs.MigrationTargetGetArgs>? Target { get; set; }

        public MigrationState()
        {
        }
    }
}