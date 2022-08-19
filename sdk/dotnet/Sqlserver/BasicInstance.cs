// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver
{
    /// <summary>
    /// Provides a SQL Server instance resource to create basic database instances.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Sqlserver.BasicInstance("foo", new Tencentcloud.Sqlserver.BasicInstanceArgs
    ///         {
    ///             AvailabilityZone = @var.Availability_zone,
    ///             ChargeType = "POSTPAID_BY_HOUR",
    ///             VpcId = "vpc-26w7r56z",
    ///             SubnetId = "subnet-lvlr6eeu",
    ///             ProjectId = 0,
    ///             Memory = 2,
    ///             Storage = 20,
    ///             Cpu = 1,
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
    ///                 "sg-nltpbqg1",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQL Server basic instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/basicInstance:BasicInstance foo mssql-3cdq7kx5
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/basicInstance:BasicInstance")]
    public partial class BasicInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
        /// </summary>
        [Output("autoRenew")]
        public Output<int?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
        /// </summary>
        [Output("autoVoucher")]
        public Output<int?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The CPU number of the SQL Server basic instance.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// Create time of the SQL Server basic instance.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string?> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk.
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// Start time of the maintenance in one day, format like `HH:mm`.
        /// </summary>
        [Output("maintenanceStartTime")]
        public Output<string> MaintenanceStartTime { get; private set; } = null!;

        /// <summary>
        /// The timespan of maintenance in one day, unit is hour.
        /// </summary>
        [Output("maintenanceTimeSpan")]
        public Output<int> MaintenanceTimeSpan { get; private set; } = null!;

        /// <summary>
        /// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
        /// </summary>
        [Output("maintenanceWeekSets")]
        public Output<ImmutableArray<int>> MaintenanceWeekSets { get; private set; } = null!;

        /// <summary>
        /// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Name of the SQL Server basic instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Project ID, default value is 0.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Security group bound to the instance.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Output("storage")]
        public Output<int> Storage { get; private set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The tags of the SQL Server basic instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// IP for private access.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// An array of voucher IDs, currently only one can be used for a single order.
        /// </summary>
        [Output("voucherIds")]
        public Output<ImmutableArray<string>> VoucherIds { get; private set; } = null!;

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Port for private access.
        /// </summary>
        [Output("vport")]
        public Output<int> Vport { get; private set; } = null!;


        /// <summary>
        /// Create a BasicInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BasicInstance(string name, BasicInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/basicInstance:BasicInstance", name, args ?? new BasicInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BasicInstance(string name, Input<string> id, BasicInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/basicInstance:BasicInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BasicInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BasicInstance Get(string name, Input<string> id, BasicInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new BasicInstance(name, id, state, options);
        }
    }

    public sealed class BasicInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
        /// </summary>
        [Input("autoRenew")]
        public Input<int>? AutoRenew { get; set; }

        /// <summary>
        /// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The CPU number of the SQL Server basic instance.
        /// </summary>
        [Input("cpu", required: true)]
        public Input<int> Cpu { get; set; } = null!;

        /// <summary>
        /// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk.
        /// </summary>
        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        /// <summary>
        /// Start time of the maintenance in one day, format like `HH:mm`.
        /// </summary>
        [Input("maintenanceStartTime")]
        public Input<string>? MaintenanceStartTime { get; set; }

        /// <summary>
        /// The timespan of maintenance in one day, unit is hour.
        /// </summary>
        [Input("maintenanceTimeSpan")]
        public Input<int>? MaintenanceTimeSpan { get; set; }

        [Input("maintenanceWeekSets")]
        private InputList<int>? _maintenanceWeekSets;

        /// <summary>
        /// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
        /// </summary>
        public InputList<int> MaintenanceWeekSets
        {
            get => _maintenanceWeekSets ?? (_maintenanceWeekSets = new InputList<int>());
            set => _maintenanceWeekSets = value;
        }

        /// <summary>
        /// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Name of the SQL Server basic instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Project ID, default value is 0.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security group bound to the instance.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Input("storage", required: true)]
        public Input<int> Storage { get; set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the SQL Server basic instance.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// An array of voucher IDs, currently only one can be used for a single order.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public BasicInstanceArgs()
        {
        }
    }

    public sealed class BasicInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.
        /// </summary>
        [Input("autoRenew")]
        public Input<int>? AutoRenew { get; set; }

        /// <summary>
        /// Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Pay type of the SQL Server basic instance. For now, only `POSTPAID_BY_HOUR` is valid.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The CPU number of the SQL Server basic instance.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// Create time of the SQL Server basic instance.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Version of the SQL Server basic database engine. Allowed values are `2008R2`(SQL Server 2008 Enterprise), `2012SP3`(SQL Server 2012 Enterprise), `2016SP1` (SQL Server 2016 Enterprise), `201602`(SQL Server 2016 Standard) and `2017`(SQL Server 2017 Enterprise). Default is `2008R2`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The host type of the purchased instance, `CLOUD_PREMIUM` for virtual machine high-performance cloud disk, `CLOUD_SSD` for virtual machine SSD cloud disk.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// Start time of the maintenance in one day, format like `HH:mm`.
        /// </summary>
        [Input("maintenanceStartTime")]
        public Input<string>? MaintenanceStartTime { get; set; }

        /// <summary>
        /// The timespan of maintenance in one day, unit is hour.
        /// </summary>
        [Input("maintenanceTimeSpan")]
        public Input<int>? MaintenanceTimeSpan { get; set; }

        [Input("maintenanceWeekSets")]
        private InputList<int>? _maintenanceWeekSets;

        /// <summary>
        /// A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.
        /// </summary>
        public InputList<int> MaintenanceWeekSets
        {
            get => _maintenanceWeekSets ?? (_maintenanceWeekSets = new InputList<int>());
            set => _maintenanceWeekSets = value;
        }

        /// <summary>
        /// Memory size (in GB). Allowed value must be larger than `memory` that data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Name of the SQL Server basic instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Project ID, default value is 0.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security group bound to the instance.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud_sqlserver_specinfos` provides.
        /// </summary>
        [Input("storage")]
        public Input<int>? Storage { get; set; }

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the SQL Server basic instance.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// IP for private access.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// An array of voucher IDs, currently only one can be used for a single order.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Port for private access.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        public BasicInstanceState()
        {
        }
    }
}
