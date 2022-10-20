// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    /// <summary>
    /// Provide a resource to create a CynosDB cluster.
    /// 
    /// ## Import
    /// 
    /// CynosDB cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cynosdb/cluster:Cluster foo cynosdbmysql-dzj5l8gz
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/cluster:Cluster")]
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int?> AutoRenewFlag { get; private set; } = null!;

        /// <summary>
        /// The available zone of the CynosDB Cluster.
        /// </summary>
        [Output("availableZone")]
        public Output<string> AvailableZone { get; private set; } = null!;

        /// <summary>
        /// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// Charset used by CynosDB cluster.
        /// </summary>
        [Output("charset")]
        public Output<string> Charset { get; private set; } = null!;

        /// <summary>
        /// Name of CynosDB cluster.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Status of the Cynosdb cluster.
        /// </summary>
        [Output("clusterStatus")]
        public Output<string> ClusterStatus { get; private set; } = null!;

        /// <summary>
        /// Creation time of the CynosDB cluster.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Type of CynosDB, and available values include `MYSQL`.
        /// </summary>
        [Output("dbType")]
        public Output<string> DbType { get; private set; } = null!;

        /// <summary>
        /// Version of CynosDB, which is related to `db_type`. For `MYSQL`, available value is `5.7`.
        /// </summary>
        [Output("dbVersion")]
        public Output<string> DbVersion { get; private set; } = null!;

        /// <summary>
        /// Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Output("instanceCpuCore")]
        public Output<int> InstanceCpuCore { get; private set; } = null!;

        /// <summary>
        /// ID of instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Duration time for maintenance, unit in second. `3600` by default.
        /// </summary>
        [Output("instanceMaintainDuration")]
        public Output<int?> InstanceMaintainDuration { get; private set; } = null!;

        /// <summary>
        /// Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        /// </summary>
        [Output("instanceMaintainStartTime")]
        public Output<int?> InstanceMaintainStartTime { get; private set; } = null!;

        /// <summary>
        /// Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        /// </summary>
        [Output("instanceMaintainWeekdays")]
        public Output<ImmutableArray<string>> InstanceMaintainWeekdays { get; private set; } = null!;

        /// <summary>
        /// Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Output("instanceMemorySize")]
        public Output<int> InstanceMemorySize { get; private set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// Status of the instance.
        /// </summary>
        [Output("instanceStatus")]
        public Output<string> InstanceStatus { get; private set; } = null!;

        /// <summary>
        /// Storage size of the instance, unit in GB.
        /// </summary>
        [Output("instanceStorageSize")]
        public Output<int> InstanceStorageSize { get; private set; } = null!;

        /// <summary>
        /// Specify parameter list of database. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
        /// </summary>
        [Output("paramItems")]
        public Output<ImmutableArray<Outputs.ClusterParamItem>> ParamItems { get; private set; } = null!;

        /// <summary>
        /// Password of `root` account.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Port of CynosDB cluster.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when charge_type is set to `PREPAID`.
        /// </summary>
        [Output("prepaidPeriod")]
        public Output<int?> PrepaidPeriod { get; private set; } = null!;

        /// <summary>
        /// ID of the project. `0` by default.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Readonly addresses. Each element contains the following attributes:
        /// </summary>
        [Output("roGroupAddrs")]
        public Output<ImmutableArray<Outputs.ClusterRoGroupAddr>> RoGroupAddrs { get; private set; } = null!;

        /// <summary>
        /// ID of read-only instance group.
        /// </summary>
        [Output("roGroupId")]
        public Output<string> RoGroupId { get; private set; } = null!;

        /// <summary>
        /// List of instances in the read-only instance group.
        /// </summary>
        [Output("roGroupInstances")]
        public Output<ImmutableArray<Outputs.ClusterRoGroupInstance>> RoGroupInstances { get; private set; } = null!;

        /// <summary>
        /// IDs of security group for `ro_group`.
        /// </summary>
        [Output("roGroupSgs")]
        public Output<ImmutableArray<string>> RoGroupSgs { get; private set; } = null!;

        /// <summary>
        /// Read-write addresses. Each element contains the following attributes:
        /// </summary>
        [Output("rwGroupAddrs")]
        public Output<ImmutableArray<Outputs.ClusterRwGroupAddr>> RwGroupAddrs { get; private set; } = null!;

        /// <summary>
        /// ID of read-write instance group.
        /// </summary>
        [Output("rwGroupId")]
        public Output<string> RwGroupId { get; private set; } = null!;

        /// <summary>
        /// List of instances in the read-write instance group.
        /// </summary>
        [Output("rwGroupInstances")]
        public Output<ImmutableArray<Outputs.ClusterRwGroupInstance>> RwGroupInstances { get; private set; } = null!;

        /// <summary>
        /// IDs of security group for `rw_group`.
        /// </summary>
        [Output("rwGroupSgs")]
        public Output<ImmutableArray<string>> RwGroupSgs { get; private set; } = null!;

        /// <summary>
        /// Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If db_type is `MYSQL` and charge_type is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when charge_type is `POSTPAID_BY_HOUR`, this argument is unnecessary.
        /// </summary>
        [Output("storageLimit")]
        public Output<int?> StorageLimit { get; private set; } = null!;

        /// <summary>
        /// Used storage of CynosDB cluster, unit in MB.
        /// </summary>
        [Output("storageUsed")]
        public Output<int> StorageUsed { get; private set; } = null!;

        /// <summary>
        /// ID of the subnet within this VPC.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The tags of the CynosDB cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// The available zone of the CynosDB Cluster.
        /// </summary>
        [Input("availableZone", required: true)]
        public Input<string> AvailableZone { get; set; } = null!;

        /// <summary>
        /// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Name of CynosDB cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Type of CynosDB, and available values include `MYSQL`.
        /// </summary>
        [Input("dbType", required: true)]
        public Input<string> DbType { get; set; } = null!;

        /// <summary>
        /// Version of CynosDB, which is related to `db_type`. For `MYSQL`, available value is `5.7`.
        /// </summary>
        [Input("dbVersion", required: true)]
        public Input<string> DbVersion { get; set; } = null!;

        /// <summary>
        /// Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceCpuCore", required: true)]
        public Input<int> InstanceCpuCore { get; set; } = null!;

        /// <summary>
        /// Duration time for maintenance, unit in second. `3600` by default.
        /// </summary>
        [Input("instanceMaintainDuration")]
        public Input<int>? InstanceMaintainDuration { get; set; }

        /// <summary>
        /// Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        /// </summary>
        [Input("instanceMaintainStartTime")]
        public Input<int>? InstanceMaintainStartTime { get; set; }

        [Input("instanceMaintainWeekdays")]
        private InputList<string>? _instanceMaintainWeekdays;

        /// <summary>
        /// Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        /// </summary>
        public InputList<string> InstanceMaintainWeekdays
        {
            get => _instanceMaintainWeekdays ?? (_instanceMaintainWeekdays = new InputList<string>());
            set => _instanceMaintainWeekdays = value;
        }

        /// <summary>
        /// Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceMemorySize", required: true)]
        public Input<int> InstanceMemorySize { get; set; } = null!;

        [Input("paramItems")]
        private InputList<Inputs.ClusterParamItemArgs>? _paramItems;

        /// <summary>
        /// Specify parameter list of database. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
        /// </summary>
        public InputList<Inputs.ClusterParamItemArgs> ParamItems
        {
            get => _paramItems ?? (_paramItems = new InputList<Inputs.ClusterParamItemArgs>());
            set => _paramItems = value;
        }

        /// <summary>
        /// Password of `root` account.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Port of CynosDB cluster.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when charge_type is set to `PREPAID`.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// ID of the project. `0` by default.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("roGroupSgs")]
        private InputList<string>? _roGroupSgs;

        /// <summary>
        /// IDs of security group for `ro_group`.
        /// </summary>
        public InputList<string> RoGroupSgs
        {
            get => _roGroupSgs ?? (_roGroupSgs = new InputList<string>());
            set => _roGroupSgs = value;
        }

        [Input("rwGroupSgs")]
        private InputList<string>? _rwGroupSgs;

        /// <summary>
        /// IDs of security group for `rw_group`.
        /// </summary>
        public InputList<string> RwGroupSgs
        {
            get => _rwGroupSgs ?? (_rwGroupSgs = new InputList<string>());
            set => _rwGroupSgs = value;
        }

        /// <summary>
        /// Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If db_type is `MYSQL` and charge_type is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when charge_type is `POSTPAID_BY_HOUR`, this argument is unnecessary.
        /// </summary>
        [Input("storageLimit")]
        public Input<int>? StorageLimit { get; set; }

        /// <summary>
        /// ID of the subnet within this VPC.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the CynosDB cluster.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. Valid values are `0`(MANUAL_RENEW), `1`(AUTO_RENEW). Default value is `0`. Only works for PREPAID cluster.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// The available zone of the CynosDB Cluster.
        /// </summary>
        [Input("availableZone")]
        public Input<string>? AvailableZone { get; set; }

        /// <summary>
        /// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `POSTPAID_BY_HOUR`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Charset used by CynosDB cluster.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// Name of CynosDB cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Status of the Cynosdb cluster.
        /// </summary>
        [Input("clusterStatus")]
        public Input<string>? ClusterStatus { get; set; }

        /// <summary>
        /// Creation time of the CynosDB cluster.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Type of CynosDB, and available values include `MYSQL`.
        /// </summary>
        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        /// <summary>
        /// Version of CynosDB, which is related to `db_type`. For `MYSQL`, available value is `5.7`.
        /// </summary>
        [Input("dbVersion")]
        public Input<string>? DbVersion { get; set; }

        /// <summary>
        /// Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its `All RELATED INSTANCES` will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceCpuCore")]
        public Input<int>? InstanceCpuCore { get; set; }

        /// <summary>
        /// ID of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Duration time for maintenance, unit in second. `3600` by default.
        /// </summary>
        [Input("instanceMaintainDuration")]
        public Input<int>? InstanceMaintainDuration { get; set; }

        /// <summary>
        /// Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        /// </summary>
        [Input("instanceMaintainStartTime")]
        public Input<int>? InstanceMaintainStartTime { get; set; }

        [Input("instanceMaintainWeekdays")]
        private InputList<string>? _instanceMaintainWeekdays;

        /// <summary>
        /// Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        /// </summary>
        public InputList<string> InstanceMaintainWeekdays
        {
            get => _instanceMaintainWeekdays ?? (_instanceMaintainWeekdays = new InputList<string>());
            set => _instanceMaintainWeekdays = value;
        }

        /// <summary>
        /// Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceMemorySize")]
        public Input<int>? InstanceMemorySize { get; set; }

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Status of the instance.
        /// </summary>
        [Input("instanceStatus")]
        public Input<string>? InstanceStatus { get; set; }

        /// <summary>
        /// Storage size of the instance, unit in GB.
        /// </summary>
        [Input("instanceStorageSize")]
        public Input<int>? InstanceStorageSize { get; set; }

        [Input("paramItems")]
        private InputList<Inputs.ClusterParamItemGetArgs>? _paramItems;

        /// <summary>
        /// Specify parameter list of database. Use `data.tencentcloud_mysql_default_params` to query available parameter details.
        /// </summary>
        public InputList<Inputs.ClusterParamItemGetArgs> ParamItems
        {
            get => _paramItems ?? (_paramItems = new InputList<Inputs.ClusterParamItemGetArgs>());
            set => _paramItems = value;
        }

        /// <summary>
        /// Password of `root` account.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Port of CynosDB cluster.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The tenancy (time unit is month) of the prepaid instance. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. NOTE: it only works when charge_type is set to `PREPAID`.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// ID of the project. `0` by default.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("roGroupAddrs")]
        private InputList<Inputs.ClusterRoGroupAddrGetArgs>? _roGroupAddrs;

        /// <summary>
        /// Readonly addresses. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.ClusterRoGroupAddrGetArgs> RoGroupAddrs
        {
            get => _roGroupAddrs ?? (_roGroupAddrs = new InputList<Inputs.ClusterRoGroupAddrGetArgs>());
            set => _roGroupAddrs = value;
        }

        /// <summary>
        /// ID of read-only instance group.
        /// </summary>
        [Input("roGroupId")]
        public Input<string>? RoGroupId { get; set; }

        [Input("roGroupInstances")]
        private InputList<Inputs.ClusterRoGroupInstanceGetArgs>? _roGroupInstances;

        /// <summary>
        /// List of instances in the read-only instance group.
        /// </summary>
        public InputList<Inputs.ClusterRoGroupInstanceGetArgs> RoGroupInstances
        {
            get => _roGroupInstances ?? (_roGroupInstances = new InputList<Inputs.ClusterRoGroupInstanceGetArgs>());
            set => _roGroupInstances = value;
        }

        [Input("roGroupSgs")]
        private InputList<string>? _roGroupSgs;

        /// <summary>
        /// IDs of security group for `ro_group`.
        /// </summary>
        public InputList<string> RoGroupSgs
        {
            get => _roGroupSgs ?? (_roGroupSgs = new InputList<string>());
            set => _roGroupSgs = value;
        }

        [Input("rwGroupAddrs")]
        private InputList<Inputs.ClusterRwGroupAddrGetArgs>? _rwGroupAddrs;

        /// <summary>
        /// Read-write addresses. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.ClusterRwGroupAddrGetArgs> RwGroupAddrs
        {
            get => _rwGroupAddrs ?? (_rwGroupAddrs = new InputList<Inputs.ClusterRwGroupAddrGetArgs>());
            set => _rwGroupAddrs = value;
        }

        /// <summary>
        /// ID of read-write instance group.
        /// </summary>
        [Input("rwGroupId")]
        public Input<string>? RwGroupId { get; set; }

        [Input("rwGroupInstances")]
        private InputList<Inputs.ClusterRwGroupInstanceGetArgs>? _rwGroupInstances;

        /// <summary>
        /// List of instances in the read-write instance group.
        /// </summary>
        public InputList<Inputs.ClusterRwGroupInstanceGetArgs> RwGroupInstances
        {
            get => _rwGroupInstances ?? (_rwGroupInstances = new InputList<Inputs.ClusterRwGroupInstanceGetArgs>());
            set => _rwGroupInstances = value;
        }

        [Input("rwGroupSgs")]
        private InputList<string>? _rwGroupSgs;

        /// <summary>
        /// IDs of security group for `rw_group`.
        /// </summary>
        public InputList<string> RwGroupSgs
        {
            get => _rwGroupSgs ?? (_rwGroupSgs = new InputList<string>());
            set => _rwGroupSgs = value;
        }

        /// <summary>
        /// Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If db_type is `MYSQL` and charge_type is `PREPAID`, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when charge_type is `POSTPAID_BY_HOUR`, this argument is unnecessary.
        /// </summary>
        [Input("storageLimit")]
        public Input<int>? StorageLimit { get; set; }

        /// <summary>
        /// Used storage of CynosDB cluster, unit in MB.
        /// </summary>
        [Input("storageUsed")]
        public Input<int>? StorageUsed { get; set; }

        /// <summary>
        /// ID of the subnet within this VPC.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the CynosDB cluster.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClusterState()
        {
        }
    }
}
