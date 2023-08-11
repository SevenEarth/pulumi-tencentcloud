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
    /// Provides a resource to create a sqlserver database_tde
    /// 
    /// ## Example Usage
    /// ### Open database tde encryption
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
    ///         var exampleDatabaseTde = new Tencentcloud.Sqlserver.DatabaseTde("exampleDatabaseTde", new Tencentcloud.Sqlserver.DatabaseTdeArgs
    ///         {
    ///             InstanceId = exampleBasicInstance.Id,
    ///             DbNames = 
    ///             {
    ///                 exampleDb.Name,
    ///             },
    ///             Encryption = "enable",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Close database tde encryption
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Tencentcloud.Sqlserver.DatabaseTde("example", new Tencentcloud.Sqlserver.DatabaseTdeArgs
    ///         {
    ///             InstanceId = tencentcloud_sqlserver_instance.Example.Id,
    ///             DbNames = 
    ///             {
    ///                 tencentcloud_sqlserver_db.Example.Name,
    ///             },
    ///             Encryption = "disable",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver database_tde can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/databaseTde:DatabaseTde example mssql-farjz9tz#tf_example_db
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/databaseTde:DatabaseTde")]
    public partial class DatabaseTde : Pulumi.CustomResource
    {
        /// <summary>
        /// Database name list.
        /// </summary>
        [Output("dbNames")]
        public Output<ImmutableArray<string>> DbNames { get; private set; } = null!;

        /// <summary>
        /// `enable` - enable encryption, `disable` - disable encryption.
        /// </summary>
        [Output("encryption")]
        public Output<string> Encryption { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseTde resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseTde(string name, DatabaseTdeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/databaseTde:DatabaseTde", name, args ?? new DatabaseTdeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseTde(string name, Input<string> id, DatabaseTdeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/databaseTde:DatabaseTde", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseTde resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseTde Get(string name, Input<string> id, DatabaseTdeState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseTde(name, id, state, options);
        }
    }

    public sealed class DatabaseTdeArgs : Pulumi.ResourceArgs
    {
        [Input("dbNames", required: true)]
        private InputList<string>? _dbNames;

        /// <summary>
        /// Database name list.
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        /// <summary>
        /// `enable` - enable encryption, `disable` - disable encryption.
        /// </summary>
        [Input("encryption", required: true)]
        public Input<string> Encryption { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public DatabaseTdeArgs()
        {
        }
    }

    public sealed class DatabaseTdeState : Pulumi.ResourceArgs
    {
        [Input("dbNames")]
        private InputList<string>? _dbNames;

        /// <summary>
        /// Database name list.
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        /// <summary>
        /// `enable` - enable encryption, `disable` - disable encryption.
        /// </summary>
        [Input("encryption")]
        public Input<string>? Encryption { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public DatabaseTdeState()
        {
        }
    }
}
