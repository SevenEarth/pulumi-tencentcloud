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
    /// Provides a resource to create a sqlserver general_backup
    /// 
    /// ## Example Usage
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
    ///         var zones = Output.Create(Tencentcloud.Availability.GetZonesByProduct.InvokeAsync(new Tencentcloud.Availability.GetZonesByProductArgs
    ///         {
    ///             Product = "sqlserver",
    ///         }));
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             VpcId = vpc.Id,
    ///             CidrBlock = "10.0.0.0/16",
    ///             IsMulticast = false,
    ///         });
    ///         var securityGroup = new Tencentcloud.Security.Group("securityGroup", new Tencentcloud.Security.GroupArgs
    ///         {
    ///             Description = "desc.",
    ///         });
    ///         var exampleBasicInstance = new Tencentcloud.Sqlserver.BasicInstance("exampleBasicInstance", new Tencentcloud.Sqlserver.BasicInstanceArgs
    ///         {
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             ChargeType = "POSTPAID_BY_HOUR",
    ///             VpcId = vpc.Id,
    ///             SubnetId = subnet.Id,
    ///             ProjectId = 0,
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_PREMIUM",
    ///             MaintenanceWeekSets = 
    ///             {
    ///                 1,
    ///                 2,
    ///                 3,
    ///             },
    ///             MaintenanceStartTime = "09:00",
    ///             MaintenanceTimeSpan = 3,
    ///             SecurityGroups = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///         });
    ///         var exampleDb = new Tencentcloud.Sqlserver.Db("exampleDb", new Tencentcloud.Sqlserver.DbArgs
    ///         {
    ///             InstanceId = exampleBasicInstance.Id,
    ///             Charset = "Chinese_PRC_BIN",
    ///             Remark = "test-remark",
    ///         });
    ///         var exampleGeneralBackup = new Tencentcloud.Sqlserver.GeneralBackup("exampleGeneralBackup", new Tencentcloud.Sqlserver.GeneralBackupArgs
    ///         {
    ///             InstanceId = exampleDb.InstanceId,
    ///             BackupName = "tf_example_backup",
    ///             Strategy = 0,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver general_backups can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/generalBackup:GeneralBackup example mssql-qelbzgwf#3512621#5293#2020-07-31 14:28:51#2020-07-31 15:10:27#autoed_instance_58037_20200728011545.bak.tar
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/generalBackup:GeneralBackup")]
    public partial class GeneralBackup : Pulumi.CustomResource
    {
        /// <summary>
        /// Backup name. If this parameter is left empty, a backup name in the format of [Instance ID]_[Backup start timestamp] will be automatically generated.
        /// </summary>
        [Output("backupName")]
        public Output<string> BackupName { get; private set; } = null!;

        /// <summary>
        /// List of names of databases to be backed up (required only for multi-database backup).
        /// </summary>
        [Output("dbNames")]
        public Output<ImmutableArray<string>> DbNames { get; private set; } = null!;

        /// <summary>
        /// flow id.
        /// </summary>
        [Output("flowId")]
        public Output<string> FlowId { get; private set; } = null!;

        /// <summary>
        /// Instance ID in the format of mssql-i1z41iwd.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Backup policy (0: instance backup, 1: multi-database backup).
        /// </summary>
        [Output("strategy")]
        public Output<int> Strategy { get; private set; } = null!;


        /// <summary>
        /// Create a GeneralBackup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeneralBackup(string name, GeneralBackupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/generalBackup:GeneralBackup", name, args ?? new GeneralBackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeneralBackup(string name, Input<string> id, GeneralBackupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/generalBackup:GeneralBackup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GeneralBackup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeneralBackup Get(string name, Input<string> id, GeneralBackupState? state = null, CustomResourceOptions? options = null)
        {
            return new GeneralBackup(name, id, state, options);
        }
    }

    public sealed class GeneralBackupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backup name. If this parameter is left empty, a backup name in the format of [Instance ID]_[Backup start timestamp] will be automatically generated.
        /// </summary>
        [Input("backupName")]
        public Input<string>? BackupName { get; set; }

        [Input("dbNames")]
        private InputList<string>? _dbNames;

        /// <summary>
        /// List of names of databases to be backed up (required only for multi-database backup).
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        /// <summary>
        /// Instance ID in the format of mssql-i1z41iwd.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Backup policy (0: instance backup, 1: multi-database backup).
        /// </summary>
        [Input("strategy")]
        public Input<int>? Strategy { get; set; }

        public GeneralBackupArgs()
        {
        }
    }

    public sealed class GeneralBackupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backup name. If this parameter is left empty, a backup name in the format of [Instance ID]_[Backup start timestamp] will be automatically generated.
        /// </summary>
        [Input("backupName")]
        public Input<string>? BackupName { get; set; }

        [Input("dbNames")]
        private InputList<string>? _dbNames;

        /// <summary>
        /// List of names of databases to be backed up (required only for multi-database backup).
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        /// <summary>
        /// flow id.
        /// </summary>
        [Input("flowId")]
        public Input<string>? FlowId { get; set; }

        /// <summary>
        /// Instance ID in the format of mssql-i1z41iwd.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Backup policy (0: instance backup, 1: multi-database backup).
        /// </summary>
        [Input("strategy")]
        public Input<int>? Strategy { get; set; }

        public GeneralBackupState()
        {
        }
    }
}