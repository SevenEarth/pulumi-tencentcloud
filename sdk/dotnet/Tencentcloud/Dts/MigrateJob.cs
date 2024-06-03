// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts
{
    /// <summary>
    /// Provides a resource to create a dts migrate_job
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
    ///     var foo = new Tencentcloud.Cynosdb.Cluster("foo", new()
    ///     {
    ///         AvailableZone = @var.Availability_zone,
    ///         VpcId = local.Vpc_id,
    ///         SubnetId = local.Subnet_id,
    ///         DbType = "MYSQL",
    ///         DbVersion = "5.7",
    ///         StorageLimit = 1000,
    ///         ClusterName = "tf-cynosdb-mysql",
    ///         Password = "cynos@123",
    ///         InstanceMaintainDuration = 3600,
    ///         InstanceMaintainStartTime = 10800,
    ///         InstanceMaintainWeekdays = new[]
    ///         {
    ///             "Fri",
    ///             "Mon",
    ///             "Sat",
    ///             "Sun",
    ///             "Thu",
    ///             "Wed",
    ///             "Tue",
    ///         },
    ///         InstanceCpuCore = 1,
    ///         InstanceMemorySize = 2,
    ///         ParamItems = new[]
    ///         {
    ///             new Tencentcloud.Cynosdb.Inputs.ClusterParamItemArgs
    ///             {
    ///                 Name = "character_set_server",
    ///                 CurrentValue = "utf8",
    ///             },
    ///             new Tencentcloud.Cynosdb.Inputs.ClusterParamItemArgs
    ///             {
    ///                 Name = "time_zone",
    ///                 CurrentValue = "+09:00",
    ///             },
    ///             new Tencentcloud.Cynosdb.Inputs.ClusterParamItemArgs
    ///             {
    ///                 Name = "lower_case_table_names",
    ///                 CurrentValue = "1",
    ///             },
    ///         },
    ///         ForceDelete = true,
    ///         RwGroupSgs = new[]
    ///         {
    ///             local.Sg_id,
    ///         },
    ///         RoGroupSgs = new[]
    ///         {
    ///             local.Sg_id,
    ///         },
    ///         PrarmTemplateId = @var.My_param_template,
    ///     });
    /// 
    ///     var service = new Tencentcloud.Dts.MigrateService("service", new()
    ///     {
    ///         SrcDatabaseType = "mysql",
    ///         DstDatabaseType = "cynosdbmysql",
    ///         SrcRegion = "ap-guangzhou",
    ///         DstRegion = "ap-guangzhou",
    ///         InstanceClass = "small",
    ///         JobName = "tf_test_migration_service_1",
    ///         Tags = new[]
    ///         {
    ///             new Tencentcloud.Dts.Inputs.MigrateServiceTagArgs
    ///             {
    ///                 TagKey = "aaa",
    ///                 TagValue = "bbb",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var job = new Tencentcloud.Dts.MigrateJob("job", new()
    ///     {
    ///         ServiceId = service.Id,
    ///         RunMode = "immediate",
    ///         MigrateOption = new Tencentcloud.Dts.Inputs.MigrateJobMigrateOptionArgs
    ///         {
    ///             DatabaseTable = new Tencentcloud.Dts.Inputs.MigrateJobMigrateOptionDatabaseTableArgs
    ///             {
    ///                 ObjectMode = "partial",
    ///                 Databases = new[]
    ///                 {
    ///                     new Tencentcloud.Dts.Inputs.MigrateJobMigrateOptionDatabaseTableDatabaseArgs
    ///                     {
    ///                         DbName = "tf_ci_test",
    ///                         DbMode = "partial",
    ///                         TableMode = "partial",
    ///                         Tables = new[]
    ///                         {
    ///                             new Tencentcloud.Dts.Inputs.MigrateJobMigrateOptionDatabaseTableDatabaseTableArgs
    ///                             {
    ///                                 TableName = "test",
    ///                                 NewTableName = "test_%s",
    ///                                 TableEditMode = "rename",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         SrcInfo = new Tencentcloud.Dts.Inputs.MigrateJobSrcInfoArgs
    ///         {
    ///             Region = "ap-guangzhou",
    ///             AccessType = "cdb",
    ///             DatabaseType = "mysql",
    ///             NodeType = "simple",
    ///             Infos = new[]
    ///             {
    ///                 new Tencentcloud.Dts.Inputs.MigrateJobSrcInfoInfoArgs
    ///                 {
    ///                     User = "user_name",
    ///                     Password = "your_pw",
    ///                     InstanceId = "cdb-fitq5t9h",
    ///                 },
    ///             },
    ///         },
    ///         DstInfo = new Tencentcloud.Dts.Inputs.MigrateJobDstInfoArgs
    ///         {
    ///             Region = "ap-guangzhou",
    ///             AccessType = "cdb",
    ///             DatabaseType = "cynosdbmysql",
    ///             NodeType = "simple",
    ///             Infos = new[]
    ///             {
    ///                 new Tencentcloud.Dts.Inputs.MigrateJobDstInfoInfoArgs
    ///                 {
    ///                     User = "user_name",
    ///                     Password = "your_pw",
    ///                     InstanceId = foo.Id,
    ///                 },
    ///             },
    ///         },
    ///         AutoRetryTimeRangeMinutes = 0,
    ///     });
    /// 
    ///     var start = new Tencentcloud.Dts.MigrateJobStartOperation("start", new()
    ///     {
    ///         JobId = job.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// dts migrate_job can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Dts/migrateJob:MigrateJob migrate_job migrate_config_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dts/migrateJob:MigrateJob")]
    public partial class MigrateJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AutoRetryTimeRangeMinutes.
        /// </summary>
        [Output("autoRetryTimeRangeMinutes")]
        public Output<int?> AutoRetryTimeRangeMinutes { get; private set; } = null!;

        /// <summary>
        /// DstInfo.
        /// </summary>
        [Output("dstInfo")]
        public Output<Outputs.MigrateJobDstInfo> DstInfo { get; private set; } = null!;

        /// <summary>
        /// ExpectRunTime.
        /// </summary>
        [Output("expectRunTime")]
        public Output<string> ExpectRunTime { get; private set; } = null!;

        /// <summary>
        /// Migration job configuration options, used to describe how the task performs migration.
        /// </summary>
        [Output("migrateOption")]
        public Output<Outputs.MigrateJobMigrateOption> MigrateOption { get; private set; } = null!;

        /// <summary>
        /// Run Mode. eg:immediate,timed.
        /// </summary>
        [Output("runMode")]
        public Output<string> RunMode { get; private set; } = null!;

        /// <summary>
        /// Migrate service Id from `tencentcloud.Dts.MigrateService`.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// SrcInfo.
        /// </summary>
        [Output("srcInfo")]
        public Output<Outputs.MigrateJobSrcInfo> SrcInfo { get; private set; } = null!;

        /// <summary>
        /// Migrate job status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a MigrateJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrateJob(string name, MigrateJobArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/migrateJob:MigrateJob", name, args ?? new MigrateJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrateJob(string name, Input<string> id, MigrateJobState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/migrateJob:MigrateJob", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MigrateJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrateJob Get(string name, Input<string> id, MigrateJobState? state = null, CustomResourceOptions? options = null)
        {
            return new MigrateJob(name, id, state, options);
        }
    }

    public sealed class MigrateJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AutoRetryTimeRangeMinutes.
        /// </summary>
        [Input("autoRetryTimeRangeMinutes")]
        public Input<int>? AutoRetryTimeRangeMinutes { get; set; }

        /// <summary>
        /// DstInfo.
        /// </summary>
        [Input("dstInfo", required: true)]
        public Input<Inputs.MigrateJobDstInfoArgs> DstInfo { get; set; } = null!;

        /// <summary>
        /// ExpectRunTime.
        /// </summary>
        [Input("expectRunTime")]
        public Input<string>? ExpectRunTime { get; set; }

        /// <summary>
        /// Migration job configuration options, used to describe how the task performs migration.
        /// </summary>
        [Input("migrateOption", required: true)]
        public Input<Inputs.MigrateJobMigrateOptionArgs> MigrateOption { get; set; } = null!;

        /// <summary>
        /// Run Mode. eg:immediate,timed.
        /// </summary>
        [Input("runMode", required: true)]
        public Input<string> RunMode { get; set; } = null!;

        /// <summary>
        /// Migrate service Id from `tencentcloud.Dts.MigrateService`.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// SrcInfo.
        /// </summary>
        [Input("srcInfo", required: true)]
        public Input<Inputs.MigrateJobSrcInfoArgs> SrcInfo { get; set; } = null!;

        public MigrateJobArgs()
        {
        }
        public static new MigrateJobArgs Empty => new MigrateJobArgs();
    }

    public sealed class MigrateJobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AutoRetryTimeRangeMinutes.
        /// </summary>
        [Input("autoRetryTimeRangeMinutes")]
        public Input<int>? AutoRetryTimeRangeMinutes { get; set; }

        /// <summary>
        /// DstInfo.
        /// </summary>
        [Input("dstInfo")]
        public Input<Inputs.MigrateJobDstInfoGetArgs>? DstInfo { get; set; }

        /// <summary>
        /// ExpectRunTime.
        /// </summary>
        [Input("expectRunTime")]
        public Input<string>? ExpectRunTime { get; set; }

        /// <summary>
        /// Migration job configuration options, used to describe how the task performs migration.
        /// </summary>
        [Input("migrateOption")]
        public Input<Inputs.MigrateJobMigrateOptionGetArgs>? MigrateOption { get; set; }

        /// <summary>
        /// Run Mode. eg:immediate,timed.
        /// </summary>
        [Input("runMode")]
        public Input<string>? RunMode { get; set; }

        /// <summary>
        /// Migrate service Id from `tencentcloud.Dts.MigrateService`.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// SrcInfo.
        /// </summary>
        [Input("srcInfo")]
        public Input<Inputs.MigrateJobSrcInfoGetArgs>? SrcInfo { get; set; }

        /// <summary>
        /// Migrate job status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public MigrateJobState()
        {
        }
        public static new MigrateJobState Empty => new MigrateJobState();
    }
}
