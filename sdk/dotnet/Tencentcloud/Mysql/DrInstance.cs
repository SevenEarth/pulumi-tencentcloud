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
    /// Provides a mysql instance resource to create read-only database instances.
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
    ///     var mysqlDr = new Tencentcloud.Mysql.DrInstance("mysqlDr", new()
    ///     {
    ///         AutoRenewFlag = 0,
    ///         AvailabilityZone = "ap-shanghai-3",
    ///         ChargeType = "POSTPAID",
    ///         Cpu = 4,
    ///         DeviceType = "UNIVERSAL",
    ///         FirstSlaveZone = "ap-shanghai-4",
    ///         InstanceName = "mysql-dr-test-up",
    ///         IntranetPort = 3360,
    ///         MasterInstanceId = "cdb-adjdu3t5",
    ///         MasterRegion = "ap-guangzhou",
    ///         MemSize = 8000,
    ///         PrepaidPeriod = 1,
    ///         ProjectId = 0,
    ///         SecurityGroups = new[]
    ///         {
    ///             "sg-q4d821qk",
    ///         },
    ///         SlaveDeployMode = 1,
    ///         SlaveSyncMode = 0,
    ///         SubnetId = "subnet-5vfntba5",
    ///         Tags = 
    ///         {
    ///             { "test", "test-tf" },
    ///         },
    ///         VolumeSize = 100,
    ///         VpcId = "vpc-h6s1s3aa",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// mysql dr database instances can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Mysql/drInstance:DrInstance mysql_dr cdb-bcet7sdb
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/drInstance:DrInstance")]
    public partial class DrInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew flag. NOTES: Only supported prepaid instance.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int?> AutoRenewFlag { get; private set; } = null!;

        /// <summary>
        /// Indicates which availability zone will be used.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// CPU cores.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
        /// </summary>
        [Output("deviceType")]
        public Output<string> DeviceType { get; private set; } = null!;

        /// <summary>
        /// Zone information about first slave instance.
        /// </summary>
        [Output("firstSlaveZone")]
        public Output<string> FirstSlaveZone { get; private set; } = null!;

        /// <summary>
        /// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// The name of a mysql instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// instance intranet IP.
        /// </summary>
        [Output("intranetIp")]
        public Output<string> IntranetIp { get; private set; } = null!;

        /// <summary>
        /// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
        /// </summary>
        [Output("intranetPort")]
        public Output<int?> IntranetPort { get; private set; } = null!;

        /// <summary>
        /// Indicates the master instance ID of recovery instances.
        /// </summary>
        [Output("masterInstanceId")]
        public Output<string> MasterInstanceId { get; private set; } = null!;

        /// <summary>
        /// The zone information of the primary instance is required when you purchase a disaster recovery instance.
        /// </summary>
        [Output("masterRegion")]
        public Output<string> MasterRegion { get; private set; } = null!;

        /// <summary>
        /// Memory size (in MB).
        /// </summary>
        [Output("memSize")]
        public Output<int> MemSize { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `charge_type` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
        /// </summary>
        [Output("payType")]
        public Output<int?> PayType { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `prepaid_period` instead. Period of instance. NOTES: Only supported prepaid instance.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Period of instance. NOTES: Only supported prepaid instance.
        /// </summary>
        [Output("prepaidPeriod")]
        public Output<int?> PrepaidPeriod { get; private set; } = null!;

        /// <summary>
        /// Project ID, default value is 0.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Zone information about second slave instance.
        /// </summary>
        [Output("secondSlaveZone")]
        public Output<string?> SecondSlaveZone { get; private set; } = null!;

        /// <summary>
        /// Security groups to use.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
        /// </summary>
        [Output("slaveDeployMode")]
        public Output<int?> SlaveDeployMode { get; private set; } = null!;

        /// <summary>
        /// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
        /// </summary>
        [Output("slaveSyncMode")]
        public Output<int?> SlaveSyncMode { get; private set; } = null!;

        /// <summary>
        /// Private network ID. If `vpc_id` is set, this value is required.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Instance tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Disk size (in GB).
        /// </summary>
        [Output("volumeSize")]
        public Output<int> VolumeSize { get; private set; } = null!;

        /// <summary>
        /// ID of VPC, which can be modified once every 24 hours and can't be removed.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a DrInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DrInstance(string name, DrInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/drInstance:DrInstance", name, args ?? new DrInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DrInstance(string name, Input<string> id, DrInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/drInstance:DrInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DrInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DrInstance Get(string name, Input<string> id, DrInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new DrInstance(name, id, state, options);
        }
    }

    public sealed class DrInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Indicates which availability zone will be used.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// CPU cores.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// Zone information about first slave instance.
        /// </summary>
        [Input("firstSlaveZone")]
        public Input<string>? FirstSlaveZone { get; set; }

        /// <summary>
        /// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The name of a mysql instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
        /// </summary>
        [Input("intranetPort")]
        public Input<int>? IntranetPort { get; set; }

        /// <summary>
        /// Indicates the master instance ID of recovery instances.
        /// </summary>
        [Input("masterInstanceId", required: true)]
        public Input<string> MasterInstanceId { get; set; } = null!;

        /// <summary>
        /// The zone information of the primary instance is required when you purchase a disaster recovery instance.
        /// </summary>
        [Input("masterRegion", required: true)]
        public Input<string> MasterRegion { get; set; } = null!;

        /// <summary>
        /// Memory size (in MB).
        /// </summary>
        [Input("memSize", required: true)]
        public Input<int> MemSize { get; set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `charge_type` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
        /// </summary>
        [Input("payType")]
        public Input<int>? PayType { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `prepaid_period` instead. Period of instance. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Period of instance. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// Project ID, default value is 0.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Zone information about second slave instance.
        /// </summary>
        [Input("secondSlaveZone")]
        public Input<string>? SecondSlaveZone { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security groups to use.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
        /// </summary>
        [Input("slaveDeployMode")]
        public Input<int>? SlaveDeployMode { get; set; }

        /// <summary>
        /// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
        /// </summary>
        [Input("slaveSyncMode")]
        public Input<int>? SlaveSyncMode { get; set; }

        /// <summary>
        /// Private network ID. If `vpc_id` is set, this value is required.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Instance tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Disk size (in GB).
        /// </summary>
        [Input("volumeSize", required: true)]
        public Input<int> VolumeSize { get; set; } = null!;

        /// <summary>
        /// ID of VPC, which can be modified once every 24 hours and can't be removed.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DrInstanceArgs()
        {
        }
        public static new DrInstanceArgs Empty => new DrInstanceArgs();
    }

    public sealed class DrInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Indicates which availability zone will be used.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Pay type of instance. Valid values:`PREPAID`, `POSTPAID`. Default is `POSTPAID`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// CPU cores.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// Specify device type, available values: `UNIVERSAL` (default), `EXCLUSIVE`, `BASIC`.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// Zone information about first slave instance.
        /// </summary>
        [Input("firstSlaveZone")]
        public Input<string>? FirstSlaveZone { get; set; }

        /// <summary>
        /// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The name of a mysql instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// instance intranet IP.
        /// </summary>
        [Input("intranetIp")]
        public Input<string>? IntranetIp { get; set; }

        /// <summary>
        /// Public access port. Valid value ranges: [1024~65535]. The default value is `3306`.
        /// </summary>
        [Input("intranetPort")]
        public Input<int>? IntranetPort { get; set; }

        /// <summary>
        /// Indicates the master instance ID of recovery instances.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        /// <summary>
        /// The zone information of the primary instance is required when you purchase a disaster recovery instance.
        /// </summary>
        [Input("masterRegion")]
        public Input<string>? MasterRegion { get; set; }

        /// <summary>
        /// Memory size (in MB).
        /// </summary>
        [Input("memSize")]
        public Input<int>? MemSize { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `charge_type` instead. Pay type of instance. Valid values: `0`, `1`. `0`: prepaid, `1`: postpaid.
        /// </summary>
        [Input("payType")]
        public Input<int>? PayType { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `prepaid_period` instead. Period of instance. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Period of instance. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// Project ID, default value is 0.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Zone information about second slave instance.
        /// </summary>
        [Input("secondSlaveZone")]
        public Input<string>? SecondSlaveZone { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Security groups to use.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.
        /// </summary>
        [Input("slaveDeployMode")]
        public Input<int>? SlaveDeployMode { get; set; }

        /// <summary>
        /// Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.
        /// </summary>
        [Input("slaveSyncMode")]
        public Input<int>? SlaveSyncMode { get; set; }

        /// <summary>
        /// Private network ID. If `vpc_id` is set, this value is required.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Instance tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Disk size (in GB).
        /// </summary>
        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        /// <summary>
        /// ID of VPC, which can be modified once every 24 hours and can't be removed.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DrInstanceState()
        {
        }
        public static new DrInstanceState Empty => new DrInstanceState();
    }
}
