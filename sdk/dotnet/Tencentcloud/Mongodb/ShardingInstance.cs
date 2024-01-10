// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mongodb
{
    /// <summary>
    /// Provide a resource to create a Mongodb sharding instance.
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
    ///         var mongodb = new Tencentcloud.Mongodb.ShardingInstance("mongodb", new Tencentcloud.Mongodb.ShardingInstanceArgs
    ///         {
    ///             AvailableZone = "ap-guangzhou-3",
    ///             EngineVersion = "MONGO_36_WT",
    ///             InstanceName = "mongodb",
    ///             MachineType = "HIO10G",
    ///             Memory = 4,
    ///             MongosCpu = 1,
    ///             MongosMemory = 2,
    ///             MongosNodeNum = 3,
    ///             NodesPerShard = 3,
    ///             Password = "password1234",
    ///             ProjectId = 0,
    ///             ShardQuantity = 2,
    ///             SubnetId = "subnet-lk0svi3p",
    ///             Volume = 100,
    ///             VpcId = "vpc-mz3efvbw",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Mongodb sharding instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mongodb/shardingInstance:ShardingInstance mongodb cmgo-41s6jwy4
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mongodb/shardingInstance:ShardingInstance")]
    public partial class ShardingInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int?> AutoRenewFlag { get; private set; } = null!;

        /// <summary>
        /// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
        /// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
        /// - Version 4.2 and above are not supported.
        /// - Read-only disaster recovery instances are not supported.
        /// - Basic network cannot be selected.
        /// </summary>
        [Output("availabilityZoneLists")]
        public Output<ImmutableArray<string>> AvailabilityZoneLists { get; private set; } = null!;

        /// <summary>
        /// The available zone of the Mongodb.
        /// </summary>
        [Output("availableZone")]
        public Output<string> AvailableZone { get; private set; } = null!;

        /// <summary>
        /// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// Creation time of the Mongodb instance.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
        /// </summary>
        [Output("hiddenZone")]
        public Output<string> HiddenZone { get; private set; } = null!;

        /// <summary>
        /// Name of the Mongodb instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Number of mongos cpu.
        /// </summary>
        [Output("mongosCpu")]
        public Output<int> MongosCpu { get; private set; } = null!;

        /// <summary>
        /// Mongos memory size in GB.
        /// </summary>
        [Output("mongosMemory")]
        public Output<int> MongosMemory { get; private set; } = null!;

        /// <summary>
        /// Number of mongos.
        /// </summary>
        [Output("mongosNodeNum")]
        public Output<int> MongosNodeNum { get; private set; } = null!;

        /// <summary>
        /// Number of nodes per shard, at least 3(one master and two slaves).
        /// </summary>
        [Output("nodesPerShard")]
        public Output<int> NodesPerShard { get; private set; } = null!;

        /// <summary>
        /// Password of this Mongodb account.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to `PREPAID`.
        /// </summary>
        [Output("prepaidPeriod")]
        public Output<int?> PrepaidPeriod { get; private set; } = null!;

        /// <summary>
        /// ID of the project which the instance belongs.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// ID of the security group.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Number of sharding.
        /// </summary>
        [Output("shardQuantity")]
        public Output<int> ShardQuantity { get; private set; } = null!;

        /// <summary>
        /// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// ID of the subnet within this VPC. The value is required if `vpc_id` is set.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// IP of the Mongodb instance.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
        /// </summary>
        [Output("volume")]
        public Output<int> Volume { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// IP port of the Mongodb instance.
        /// </summary>
        [Output("vport")]
        public Output<int> Vport { get; private set; } = null!;


        /// <summary>
        /// Create a ShardingInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ShardingInstance(string name, ShardingInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mongodb/shardingInstance:ShardingInstance", name, args ?? new ShardingInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ShardingInstance(string name, Input<string> id, ShardingInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mongodb/shardingInstance:ShardingInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ShardingInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ShardingInstance Get(string name, Input<string> id, ShardingInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ShardingInstance(name, id, state, options);
        }
    }

    public sealed class ShardingInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        [Input("availabilityZoneLists")]
        private InputList<string>? _availabilityZoneLists;

        /// <summary>
        /// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
        /// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
        /// - Version 4.2 and above are not supported.
        /// - Read-only disaster recovery instances are not supported.
        /// - Basic network cannot be selected.
        /// </summary>
        public InputList<string> AvailabilityZoneLists
        {
            get => _availabilityZoneLists ?? (_availabilityZoneLists = new InputList<string>());
            set => _availabilityZoneLists = value;
        }

        /// <summary>
        /// The available zone of the Mongodb.
        /// </summary>
        [Input("availableZone", required: true)]
        public Input<string> AvailableZone { get; set; } = null!;

        /// <summary>
        /// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
        /// </summary>
        [Input("hiddenZone")]
        public Input<string>? HiddenZone { get; set; }

        /// <summary>
        /// Name of the Mongodb instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
        /// </summary>
        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        /// <summary>
        /// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Number of mongos cpu.
        /// </summary>
        [Input("mongosCpu")]
        public Input<int>? MongosCpu { get; set; }

        /// <summary>
        /// Mongos memory size in GB.
        /// </summary>
        [Input("mongosMemory")]
        public Input<int>? MongosMemory { get; set; }

        /// <summary>
        /// Number of mongos.
        /// </summary>
        [Input("mongosNodeNum")]
        public Input<int>? MongosNodeNum { get; set; }

        /// <summary>
        /// Number of nodes per shard, at least 3(one master and two slaves).
        /// </summary>
        [Input("nodesPerShard", required: true)]
        public Input<int> NodesPerShard { get; set; } = null!;

        /// <summary>
        /// Password of this Mongodb account.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to `PREPAID`.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// ID of the project which the instance belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// ID of the security group.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Number of sharding.
        /// </summary>
        [Input("shardQuantity", required: true)]
        public Input<int> ShardQuantity { get; set; } = null!;

        /// <summary>
        /// ID of the subnet within this VPC. The value is required if `vpc_id` is set.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
        /// </summary>
        [Input("volume", required: true)]
        public Input<int> Volume { get; set; } = null!;

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ShardingInstanceArgs()
        {
        }
    }

    public sealed class ShardingInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. Valid values are `0`(NOTIFY_AND_MANUAL_RENEW), `1`(NOTIFY_AND_AUTO_RENEW) and `2`(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is `0`. Note: only works for PREPAID instance. Only supports`0` and `1` for creation.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        [Input("availabilityZoneLists")]
        private InputList<string>? _availabilityZoneLists;

        /// <summary>
        /// A list of nodes deployed in multiple availability zones. For more information, please use the API DescribeSpecInfo.
        /// - Multi-availability zone deployment nodes can only be deployed in 3 different availability zones. It is not supported to deploy most nodes of the cluster in the same availability zone. For example, a 3-node cluster does not support the deployment of 2 nodes in the same zone.
        /// - Version 4.2 and above are not supported.
        /// - Read-only disaster recovery instances are not supported.
        /// - Basic network cannot be selected.
        /// </summary>
        public InputList<string> AvailabilityZoneLists
        {
            get => _availabilityZoneLists ?? (_availabilityZoneLists = new InputList<string>());
            set => _availabilityZoneLists = value;
        }

        /// <summary>
        /// The available zone of the Mongodb.
        /// </summary>
        [Input("availableZone")]
        public Input<string>? AvailableZone { get; set; }

        /// <summary>
        /// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR`. Caution that update operation on this field will delete old instances and create new one with new charge type.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Creation time of the Mongodb instance.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Version of the Mongodb, and available values include `MONGO_36_WT` (MongoDB 3.6 WiredTiger Edition), `MONGO_40_WT` (MongoDB 4.0 WiredTiger Edition) and `MONGO_42_WT`  (MongoDB 4.2 WiredTiger Edition). NOTE: `MONGO_3_WT` (MongoDB 3.2 WiredTiger Edition) and `MONGO_3_ROCKS` (MongoDB 3.2 RocksDB Edition) will deprecated.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The availability zone to which the Hidden node belongs. This parameter must be configured to deploy instances across availability zones.
        /// </summary>
        [Input("hiddenZone")]
        public Input<string>? HiddenZone { get; set; }

        /// <summary>
        /// Name of the Mongodb instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Type of Mongodb instance, and available values include `HIO`(or `GIO` which will be deprecated, represents high IO) and `HIO10G`(or `TGIO` which will be deprecated, represents 10-gigabit high IO).
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Number of mongos cpu.
        /// </summary>
        [Input("mongosCpu")]
        public Input<int>? MongosCpu { get; set; }

        /// <summary>
        /// Mongos memory size in GB.
        /// </summary>
        [Input("mongosMemory")]
        public Input<int>? MongosMemory { get; set; }

        /// <summary>
        /// Number of mongos.
        /// </summary>
        [Input("mongosNodeNum")]
        public Input<int>? MongosNodeNum { get; set; }

        /// <summary>
        /// Number of nodes per shard, at least 3(one master and two slaves).
        /// </summary>
        [Input("nodesPerShard")]
        public Input<int>? NodesPerShard { get; set; }

        /// <summary>
        /// Password of this Mongodb account.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to `PREPAID`.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// ID of the project which the instance belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// ID of the security group.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Number of sharding.
        /// </summary>
        [Input("shardQuantity")]
        public Input<int>? ShardQuantity { get; set; }

        /// <summary>
        /// Status of the Mongodb instance, and available values include pending initialization(expressed with 0),  processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// ID of the subnet within this VPC. The value is required if `vpc_id` is set.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the Mongodb. Key name `project` is system reserved and can't be used.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// IP of the Mongodb instance.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.
        /// </summary>
        [Input("volume")]
        public Input<int>? Volume { get; set; }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// IP port of the Mongodb instance.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        public ShardingInstanceState()
        {
        }
    }
}
