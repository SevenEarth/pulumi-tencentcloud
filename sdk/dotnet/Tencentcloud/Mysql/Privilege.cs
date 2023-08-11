// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a mysql account privilege resource to grant different access privilege to different database. A database can be granted by multiple account.
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
    ///             Product = "cdb",
    ///         }));
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[0]?.Name),
    ///             VpcId = vpc.Id,
    ///             CidrBlock = "10.0.0.0/16",
    ///             IsMulticast = false,
    ///         });
    ///         var securityGroup = new Tencentcloud.Security.Group("securityGroup", new Tencentcloud.Security.GroupArgs
    ///         {
    ///             Description = "mysql test",
    ///         });
    ///         var exampleInstance = new Tencentcloud.Mysql.Instance("exampleInstance", new Tencentcloud.Mysql.InstanceArgs
    ///         {
    ///             InternetService = 1,
    ///             EngineVersion = "5.7",
    ///             ChargeType = "POSTPAID",
    ///             RootPassword = "PassWord123",
    ///             SlaveDeployMode = 0,
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[0]?.Name),
    ///             SlaveSyncMode = 1,
    ///             InstanceName = "tf-example-mysql",
    ///             MemSize = 4000,
    ///             VolumeSize = 200,
    ///             VpcId = vpc.Id,
    ///             SubnetId = subnet.Id,
    ///             IntranetPort = 3306,
    ///             SecurityGroups = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "name", "test" },
    ///             },
    ///             Parameters = 
    ///             {
    ///                 { "character_set_server", "utf8" },
    ///                 { "max_connections", "1000" },
    ///             },
    ///         });
    ///         var exampleAccount = new Tencentcloud.Mysql.Account("exampleAccount", new Tencentcloud.Mysql.AccountArgs
    ///         {
    ///             MysqlId = exampleInstance.Id,
    ///             Password = "Qwer@234",
    ///             Description = "desc.",
    ///             MaxUserConnections = 10,
    ///         });
    ///         var examplePrivilege = new Tencentcloud.Mysql.Privilege("examplePrivilege", new Tencentcloud.Mysql.PrivilegeArgs
    ///         {
    ///             MysqlId = exampleInstance.Id,
    ///             AccountName = exampleAccount.Name,
    ///             Globals = 
    ///             {
    ///                 "TRIGGER",
    ///             },
    ///             Databases = 
    ///             {
    ///                 new Tencentcloud.Mysql.Inputs.PrivilegeDatabaseArgs
    ///                 {
    ///                     Privileges = 
    ///                     {
    ///                         "SELECT",
    ///                         "INSERT",
    ///                         "UPDATE",
    ///                         "DELETE",
    ///                         "CREATE",
    ///                     },
    ///                     DatabaseName = "sys",
    ///                 },
    ///                 new Tencentcloud.Mysql.Inputs.PrivilegeDatabaseArgs
    ///                 {
    ///                     Privileges = 
    ///                     {
    ///                         "SELECT",
    ///                     },
    ///                     DatabaseName = "performance_schema",
    ///                 },
    ///             },
    ///             Tables = 
    ///             {
    ///                 new Tencentcloud.Mysql.Inputs.PrivilegeTableArgs
    ///                 {
    ///                     Privileges = 
    ///                     {
    ///                         "SELECT",
    ///                         "INSERT",
    ///                         "UPDATE",
    ///                         "DELETE",
    ///                         "CREATE",
    ///                     },
    ///                     DatabaseName = "mysql",
    ///                     TableName = "slow_log",
    ///                 },
    ///                 new Tencentcloud.Mysql.Inputs.PrivilegeTableArgs
    ///                 {
    ///                     Privileges = 
    ///                     {
    ///                         "SELECT",
    ///                         "INSERT",
    ///                         "UPDATE",
    ///                     },
    ///                     DatabaseName = "mysql",
    ///                     TableName = "user",
    ///                 },
    ///             },
    ///             Columns = 
    ///             {
    ///                 new Tencentcloud.Mysql.Inputs.PrivilegeColumnArgs
    ///                 {
    ///                     Privileges = 
    ///                     {
    ///                         "SELECT",
    ///                         "INSERT",
    ///                         "UPDATE",
    ///                         "REFERENCES",
    ///                     },
    ///                     DatabaseName = "mysql",
    ///                     TableName = "user",
    ///                     ColumnName = "host",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/privilege:Privilege")]
    public partial class Privilege : Pulumi.CustomResource
    {
        /// <summary>
        /// Account host, default is `%`.
        /// </summary>
        [Output("accountHost")]
        public Output<string?> AccountHost { get; private set; } = null!;

        /// <summary>
        /// Account name.the forbidden value is:root,mysql.sys,tencentroot.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Column privileges list.
        /// </summary>
        [Output("columns")]
        public Output<ImmutableArray<Outputs.PrivilegeColumn>> Columns { get; private set; } = null!;

        /// <summary>
        /// Database privileges list.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<Outputs.PrivilegeDatabase>> Databases { get; private set; } = null!;

        /// <summary>
        /// Global privileges. available values for Privileges:ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES,CREATE USER,CREATE VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK TABLES,PROCESS,REFERENCES,RELOAD,REPLICATION CLIENT,REPLICATION SLAVE,SELECT,SHOW DATABASES,SHOW VIEW,TRIGGER,UPDATE.
        /// </summary>
        [Output("globals")]
        public Output<ImmutableArray<string>> Globals { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("mysqlId")]
        public Output<string> MysqlId { get; private set; } = null!;

        /// <summary>
        /// Table privileges list.
        /// </summary>
        [Output("tables")]
        public Output<ImmutableArray<Outputs.PrivilegeTable>> Tables { get; private set; } = null!;


        /// <summary>
        /// Create a Privilege resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Privilege(string name, PrivilegeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/privilege:Privilege", name, args ?? new PrivilegeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Privilege(string name, Input<string> id, PrivilegeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/privilege:Privilege", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Privilege resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Privilege Get(string name, Input<string> id, PrivilegeState? state = null, CustomResourceOptions? options = null)
        {
            return new Privilege(name, id, state, options);
        }
    }

    public sealed class PrivilegeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account host, default is `%`.
        /// </summary>
        [Input("accountHost")]
        public Input<string>? AccountHost { get; set; }

        /// <summary>
        /// Account name.the forbidden value is:root,mysql.sys,tencentroot.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("columns")]
        private InputList<Inputs.PrivilegeColumnArgs>? _columns;

        /// <summary>
        /// Column privileges list.
        /// </summary>
        public InputList<Inputs.PrivilegeColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.PrivilegeColumnArgs>());
            set => _columns = value;
        }

        [Input("databases")]
        private InputList<Inputs.PrivilegeDatabaseArgs>? _databases;

        /// <summary>
        /// Database privileges list.
        /// </summary>
        public InputList<Inputs.PrivilegeDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.PrivilegeDatabaseArgs>());
            set => _databases = value;
        }

        [Input("globals", required: true)]
        private InputList<string>? _globals;

        /// <summary>
        /// Global privileges. available values for Privileges:ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES,CREATE USER,CREATE VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK TABLES,PROCESS,REFERENCES,RELOAD,REPLICATION CLIENT,REPLICATION SLAVE,SELECT,SHOW DATABASES,SHOW VIEW,TRIGGER,UPDATE.
        /// </summary>
        public InputList<string> Globals
        {
            get => _globals ?? (_globals = new InputList<string>());
            set => _globals = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("mysqlId", required: true)]
        public Input<string> MysqlId { get; set; } = null!;

        [Input("tables")]
        private InputList<Inputs.PrivilegeTableArgs>? _tables;

        /// <summary>
        /// Table privileges list.
        /// </summary>
        public InputList<Inputs.PrivilegeTableArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.PrivilegeTableArgs>());
            set => _tables = value;
        }

        public PrivilegeArgs()
        {
        }
    }

    public sealed class PrivilegeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account host, default is `%`.
        /// </summary>
        [Input("accountHost")]
        public Input<string>? AccountHost { get; set; }

        /// <summary>
        /// Account name.the forbidden value is:root,mysql.sys,tencentroot.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("columns")]
        private InputList<Inputs.PrivilegeColumnGetArgs>? _columns;

        /// <summary>
        /// Column privileges list.
        /// </summary>
        public InputList<Inputs.PrivilegeColumnGetArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.PrivilegeColumnGetArgs>());
            set => _columns = value;
        }

        [Input("databases")]
        private InputList<Inputs.PrivilegeDatabaseGetArgs>? _databases;

        /// <summary>
        /// Database privileges list.
        /// </summary>
        public InputList<Inputs.PrivilegeDatabaseGetArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.PrivilegeDatabaseGetArgs>());
            set => _databases = value;
        }

        [Input("globals")]
        private InputList<string>? _globals;

        /// <summary>
        /// Global privileges. available values for Privileges:ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES,CREATE USER,CREATE VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK TABLES,PROCESS,REFERENCES,RELOAD,REPLICATION CLIENT,REPLICATION SLAVE,SELECT,SHOW DATABASES,SHOW VIEW,TRIGGER,UPDATE.
        /// </summary>
        public InputList<string> Globals
        {
            get => _globals ?? (_globals = new InputList<string>());
            set => _globals = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("mysqlId")]
        public Input<string>? MysqlId { get; set; }

        [Input("tables")]
        private InputList<Inputs.PrivilegeTableGetArgs>? _tables;

        /// <summary>
        /// Table privileges list.
        /// </summary>
        public InputList<Inputs.PrivilegeTableGetArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.PrivilegeTableGetArgs>());
            set => _tables = value;
        }

        public PrivilegeState()
        {
        }
    }
}
