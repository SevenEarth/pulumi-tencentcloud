// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    /// <summary>
    /// Provides a resource to create a mariadb instance
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
    ///         var instance = new Tencentcloud.Mariadb.Instance("instance", new Tencentcloud.Mariadb.InstanceArgs
    ///         {
    ///             AutoRenewFlag = 1,
    ///             DbVersionId = "8.0",
    ///             DcnInstanceId = "",
    ///             DcnRegion = "",
    ///             InitParams = 
    ///             {
    ///                 new Tencentcloud.Mariadb.Inputs.InstanceInitParamArgs
    ///                 {
    ///                     Param = "character_set_server",
    ///                     Value = "utf8mb4",
    ///                 },
    ///                 new Tencentcloud.Mariadb.Inputs.InstanceInitParamArgs
    ///                 {
    ///                     Param = "lower_case_table_names",
    ///                     Value = "0",
    ///                 },
    ///                 new Tencentcloud.Mariadb.Inputs.InstanceInitParamArgs
    ///                 {
    ///                     Param = "innodb_page_size",
    ///                     Value = "16384",
    ///                 },
    ///                 new Tencentcloud.Mariadb.Inputs.InstanceInitParamArgs
    ///                 {
    ///                     Param = "sync_mode",
    ///                     Value = "1",
    ///                 },
    ///             },
    ///             InstanceName = "terraform-test",
    ///             Ipv6Flag = 0,
    ///             Memory = 8,
    ///             NodeCount = 2,
    ///             Period = 1,
    ///             Storage = 10,
    ///             SubnetId = "subnet-3ku415by",
    ///             Tags = 
    ///             {
    ///                 { "createby", "terrafrom-2" },
    ///             },
    ///             VpcId = "vpc-ii1jfbhl",
    ///             Zones = 
    ///             {
    ///                 "ap-guangzhou-3",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mariadb tencentcloud_mariadb_instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mariadb/instance:Instance instance tdsql-4pzs5b67
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mariadb/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the application to which the instance belongs.
        /// </summary>
        [Output("appId")]
        public Output<int> AppId { get; private set; } = null!;

        /// <summary>
        /// Automatic renewal flag, 1: automatic renewal, 2: no automatic renewal.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int> AutoRenewFlag { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically use the voucher for payment, the default is not used.
        /// </summary>
        [Output("autoVoucher")]
        public Output<bool?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Number of CPU cores of the instance.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// Instance creation time, the format is 2006-01-02 15:04:05.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Database Engine.
        /// </summary>
        [Output("dbEngine")]
        public Output<string> DbEngine { get; private set; } = null!;

        /// <summary>
        /// Database engine version, currently available: 8.0.18, 10.1.9, 5.7.17. If not passed, the default is Percona 5.7.17.
        /// </summary>
        [Output("dbVersionId")]
        public Output<string> DbVersionId { get; private set; } = null!;

        /// <summary>
        /// Number of DCN disaster recovery instances.
        /// </summary>
        [Output("dcnDstNum")]
        public Output<int> DcnDstNum { get; private set; } = null!;

        /// <summary>
        /// DCN flag, 0-none, 1-primary instance, 2-disaster backup instance.
        /// </summary>
        [Output("dcnFlag")]
        public Output<int> DcnFlag { get; private set; } = null!;

        /// <summary>
        /// DCN source instance ID.
        /// </summary>
        [Output("dcnInstanceId")]
        public Output<string?> DcnInstanceId { get; private set; } = null!;

        /// <summary>
        /// DCN source region.
        /// </summary>
        [Output("dcnRegion")]
        public Output<string?> DcnRegion { get; private set; } = null!;

        /// <summary>
        /// DCN status, 0-none, 1-creating, 2-synchronizing, 3-disconnected.
        /// </summary>
        [Output("dcnStatus")]
        public Output<int> DcnStatus { get; private set; } = null!;

        /// <summary>
        /// Exclusive cluster ID, if it is empty, it means a normal instance.
        /// </summary>
        [Output("exclusterId")]
        public Output<string> ExclusterId { get; private set; } = null!;

        /// <summary>
        /// Parameter list. The optional values of this interface are: character_set_server (character set, required) enum: utf8,latin1,gbk,utf8mb4,gb18030, lower_case_table_names (table name case sensitive, required, 0 - sensitive; 1 - insensitive), innodb_page_size (innodb data page, Default 16K), sync_mode (sync mode: 0 - asynchronous; 1 - strong synchronous; 2 - strong synchronous can degenerate. The default is strong synchronous can degenerate).
        /// </summary>
        [Output("initParams")]
        public Output<ImmutableArray<Outputs.InstanceInitParam>> InitParams { get; private set; } = null!;

        /// <summary>
        /// Instance ID, uniquely identifies a TDSQL instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance name, you can set the name of the instance independently through this field.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// 1: primary instance (exclusive), 2: primary instance, 3: disaster recovery instance, 4: disaster recovery instance (exclusive type).
        /// </summary>
        [Output("instanceType")]
        public Output<int> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Whether IPv6 is supported.
        /// </summary>
        [Output("ipv6Flag")]
        public Output<int> Ipv6Flag { get; private set; } = null!;

        /// <summary>
        /// Whether the instance supports auditing. 1-supported; 0-not supported.
        /// </summary>
        [Output("isAuditSupported")]
        public Output<int> IsAuditSupported { get; private set; } = null!;

        /// <summary>
        /// Whether data encryption is supported. 1-supported; 0-not supported.
        /// </summary>
        [Output("isEncryptSupported")]
        public Output<int> IsEncryptSupported { get; private set; } = null!;

        /// <summary>
        /// Whether it is a temporary instance, 0 means no, non-zero means yes.
        /// </summary>
        [Output("isTmp")]
        public Output<int> IsTmp { get; private set; } = null!;

        /// <summary>
        /// Asynchronous task process ID when the instance is in an asynchronous task.
        /// </summary>
        [Output("locker")]
        public Output<int> Locker { get; private set; } = null!;

        /// <summary>
        /// Machine Model.
        /// </summary>
        [Output("machine")]
        public Output<string> Machine { get; private set; } = null!;

        /// <summary>
        /// Memory size, unit: GB, can be obtained by querying instance specifications through DescribeDBInstanceSpecs.
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Number of nodes, 2 is one master and one slave, 3 is one master and two slaves.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        /// <summary>
        /// Payment Mode.
        /// </summary>
        [Output("paymode")]
        public Output<string> Paymode { get; private set; } = null!;

        /// <summary>
        /// The duration of the purchase, unit: month.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Instance expiration time, the format is 2006-01-02 15:04:05.
        /// </summary>
        [Output("periodEndTime")]
        public Output<string> PeriodEndTime { get; private set; } = null!;

        /// <summary>
        /// Product Type ID.
        /// </summary>
        [Output("pid")]
        public Output<int> Pid { get; private set; } = null!;

        /// <summary>
        /// Project ID, which can be obtained by viewing the project list, if not passed, it will be associated with the default project.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Maximum Qps value.
        /// </summary>
        [Output("qps")]
        public Output<int> Qps { get; private set; } = null!;

        /// <summary>
        /// The name of the region where the instance is located, such as ap-shanghai.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Security group ID list.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Instance status: 0 creating, 1 process processing, 2 running, 3 instance not initialized, -1 instance isolated, 4 instance initializing, 5 instance deleting, 6 instance restarting, 7 data migration.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Description of the current running state of the instance.
        /// </summary>
        [Output("statusDesc")]
        public Output<string> StatusDesc { get; private set; } = null!;

        /// <summary>
        /// Storage size, unit: GB. You can query instance specifications through DescribeDBInstanceSpecs to obtain the lower and upper limits of disk specifications corresponding to different memory sizes.
        /// </summary>
        [Output("storage")]
        public Output<int> Storage { get; private set; } = null!;

        /// <summary>
        /// Virtual private network subnet ID, required when VpcId is not empty.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// tag list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// TDSQL version information.
        /// </summary>
        [Output("tdsqlVersion")]
        public Output<string> TdsqlVersion { get; private set; } = null!;

        /// <summary>
        /// The account to which the instance belongs.
        /// </summary>
        [Output("uin")]
        public Output<string> Uin { get; private set; } = null!;

        /// <summary>
        /// The last update time of the instance in the format of 2006-01-02 15:04:05.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Intranet IP address.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// Intranet IPv6.
        /// </summary>
        [Output("vipv6")]
        public Output<string> Vipv6 { get; private set; } = null!;

        /// <summary>
        /// A list of voucher IDs. Currently, only one voucher can be specified.
        /// </summary>
        [Output("voucherIds")]
        public Output<ImmutableArray<string>> VoucherIds { get; private set; } = null!;

        /// <summary>
        /// Virtual private network ID, if not passed, it means that it is created as a basic network.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// Intranet port.
        /// </summary>
        [Output("vport")]
        public Output<int> Vport { get; private set; } = null!;

        /// <summary>
        /// The domain name accessed from the external network, which can be resolved by the public network.
        /// </summary>
        [Output("wanDomain")]
        public Output<string> WanDomain { get; private set; } = null!;

        /// <summary>
        /// Internet port.
        /// </summary>
        [Output("wanPort")]
        public Output<int> WanPort { get; private set; } = null!;

        /// <summary>
        /// Internet IPv6 port.
        /// </summary>
        [Output("wanPortIpv6")]
        public Output<int> WanPortIpv6 { get; private set; } = null!;

        /// <summary>
        /// External network status, 0-unopened; 1-opened; 2-closed; 3-opening.
        /// </summary>
        [Output("wanStatus")]
        public Output<int> WanStatus { get; private set; } = null!;

        /// <summary>
        /// Internet IPv6 status.
        /// </summary>
        [Output("wanStatusIpv6")]
        public Output<int> WanStatusIpv6 { get; private set; } = null!;

        /// <summary>
        /// Extranet IP address, accessible from the public network.
        /// </summary>
        [Output("wanVip")]
        public Output<string> WanVip { get; private set; } = null!;

        /// <summary>
        /// Internet IPv6.
        /// </summary>
        [Output("wanVipv6")]
        public Output<string> WanVipv6 { get; private set; } = null!;

        /// <summary>
        /// Instance node availability zone distribution, up to two availability zones can be filled. When the shard specification is one master and two slaves, two of the nodes are in the first availability zone.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatic renewal flag, 1: automatic renewal, 2: no automatic renewal.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Whether to automatically use the voucher for payment, the default is not used.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        /// <summary>
        /// Database engine version, currently available: 8.0.18, 10.1.9, 5.7.17. If not passed, the default is Percona 5.7.17.
        /// </summary>
        [Input("dbVersionId")]
        public Input<string>? DbVersionId { get; set; }

        /// <summary>
        /// DCN source instance ID.
        /// </summary>
        [Input("dcnInstanceId")]
        public Input<string>? DcnInstanceId { get; set; }

        /// <summary>
        /// DCN source region.
        /// </summary>
        [Input("dcnRegion")]
        public Input<string>? DcnRegion { get; set; }

        [Input("initParams")]
        private InputList<Inputs.InstanceInitParamArgs>? _initParams;

        /// <summary>
        /// Parameter list. The optional values of this interface are: character_set_server (character set, required) enum: utf8,latin1,gbk,utf8mb4,gb18030, lower_case_table_names (table name case sensitive, required, 0 - sensitive; 1 - insensitive), innodb_page_size (innodb data page, Default 16K), sync_mode (sync mode: 0 - asynchronous; 1 - strong synchronous; 2 - strong synchronous can degenerate. The default is strong synchronous can degenerate).
        /// </summary>
        public InputList<Inputs.InstanceInitParamArgs> InitParams
        {
            get => _initParams ?? (_initParams = new InputList<Inputs.InstanceInitParamArgs>());
            set => _initParams = value;
        }

        /// <summary>
        /// Instance name, you can set the name of the instance independently through this field.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Whether IPv6 is supported.
        /// </summary>
        [Input("ipv6Flag")]
        public Input<int>? Ipv6Flag { get; set; }

        /// <summary>
        /// Memory size, unit: GB, can be obtained by querying instance specifications through DescribeDBInstanceSpecs.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Number of nodes, 2 is one master and one slave, 3 is one master and two slaves.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        /// <summary>
        /// The duration of the purchase, unit: month.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Project ID, which can be obtained by viewing the project list, if not passed, it will be associated with the default project.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group ID list.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Storage size, unit: GB. You can query instance specifications through DescribeDBInstanceSpecs to obtain the lower and upper limits of disk specifications corresponding to different memory sizes.
        /// </summary>
        [Input("storage", required: true)]
        public Input<int> Storage { get; set; } = null!;

        /// <summary>
        /// Virtual private network subnet ID, required when VpcId is not empty.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// tag list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Intranet IP address.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// A list of voucher IDs. Currently, only one voucher can be specified.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// Virtual private network ID, if not passed, it means that it is created as a basic network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("zones", required: true)]
        private InputList<string>? _zones;

        /// <summary>
        /// Instance node availability zone distribution, up to two availability zones can be filled. When the shard specification is one master and two slaves, two of the nodes are in the first availability zone.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the application to which the instance belongs.
        /// </summary>
        [Input("appId")]
        public Input<int>? AppId { get; set; }

        /// <summary>
        /// Automatic renewal flag, 1: automatic renewal, 2: no automatic renewal.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Whether to automatically use the voucher for payment, the default is not used.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        /// <summary>
        /// Number of CPU cores of the instance.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// Instance creation time, the format is 2006-01-02 15:04:05.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Database Engine.
        /// </summary>
        [Input("dbEngine")]
        public Input<string>? DbEngine { get; set; }

        /// <summary>
        /// Database engine version, currently available: 8.0.18, 10.1.9, 5.7.17. If not passed, the default is Percona 5.7.17.
        /// </summary>
        [Input("dbVersionId")]
        public Input<string>? DbVersionId { get; set; }

        /// <summary>
        /// Number of DCN disaster recovery instances.
        /// </summary>
        [Input("dcnDstNum")]
        public Input<int>? DcnDstNum { get; set; }

        /// <summary>
        /// DCN flag, 0-none, 1-primary instance, 2-disaster backup instance.
        /// </summary>
        [Input("dcnFlag")]
        public Input<int>? DcnFlag { get; set; }

        /// <summary>
        /// DCN source instance ID.
        /// </summary>
        [Input("dcnInstanceId")]
        public Input<string>? DcnInstanceId { get; set; }

        /// <summary>
        /// DCN source region.
        /// </summary>
        [Input("dcnRegion")]
        public Input<string>? DcnRegion { get; set; }

        /// <summary>
        /// DCN status, 0-none, 1-creating, 2-synchronizing, 3-disconnected.
        /// </summary>
        [Input("dcnStatus")]
        public Input<int>? DcnStatus { get; set; }

        /// <summary>
        /// Exclusive cluster ID, if it is empty, it means a normal instance.
        /// </summary>
        [Input("exclusterId")]
        public Input<string>? ExclusterId { get; set; }

        [Input("initParams")]
        private InputList<Inputs.InstanceInitParamGetArgs>? _initParams;

        /// <summary>
        /// Parameter list. The optional values of this interface are: character_set_server (character set, required) enum: utf8,latin1,gbk,utf8mb4,gb18030, lower_case_table_names (table name case sensitive, required, 0 - sensitive; 1 - insensitive), innodb_page_size (innodb data page, Default 16K), sync_mode (sync mode: 0 - asynchronous; 1 - strong synchronous; 2 - strong synchronous can degenerate. The default is strong synchronous can degenerate).
        /// </summary>
        public InputList<Inputs.InstanceInitParamGetArgs> InitParams
        {
            get => _initParams ?? (_initParams = new InputList<Inputs.InstanceInitParamGetArgs>());
            set => _initParams = value;
        }

        /// <summary>
        /// Instance ID, uniquely identifies a TDSQL instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance name, you can set the name of the instance independently through this field.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// 1: primary instance (exclusive), 2: primary instance, 3: disaster recovery instance, 4: disaster recovery instance (exclusive type).
        /// </summary>
        [Input("instanceType")]
        public Input<int>? InstanceType { get; set; }

        /// <summary>
        /// Whether IPv6 is supported.
        /// </summary>
        [Input("ipv6Flag")]
        public Input<int>? Ipv6Flag { get; set; }

        /// <summary>
        /// Whether the instance supports auditing. 1-supported; 0-not supported.
        /// </summary>
        [Input("isAuditSupported")]
        public Input<int>? IsAuditSupported { get; set; }

        /// <summary>
        /// Whether data encryption is supported. 1-supported; 0-not supported.
        /// </summary>
        [Input("isEncryptSupported")]
        public Input<int>? IsEncryptSupported { get; set; }

        /// <summary>
        /// Whether it is a temporary instance, 0 means no, non-zero means yes.
        /// </summary>
        [Input("isTmp")]
        public Input<int>? IsTmp { get; set; }

        /// <summary>
        /// Asynchronous task process ID when the instance is in an asynchronous task.
        /// </summary>
        [Input("locker")]
        public Input<int>? Locker { get; set; }

        /// <summary>
        /// Machine Model.
        /// </summary>
        [Input("machine")]
        public Input<string>? Machine { get; set; }

        /// <summary>
        /// Memory size, unit: GB, can be obtained by querying instance specifications through DescribeDBInstanceSpecs.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Number of nodes, 2 is one master and one slave, 3 is one master and two slaves.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// Payment Mode.
        /// </summary>
        [Input("paymode")]
        public Input<string>? Paymode { get; set; }

        /// <summary>
        /// The duration of the purchase, unit: month.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Instance expiration time, the format is 2006-01-02 15:04:05.
        /// </summary>
        [Input("periodEndTime")]
        public Input<string>? PeriodEndTime { get; set; }

        /// <summary>
        /// Product Type ID.
        /// </summary>
        [Input("pid")]
        public Input<int>? Pid { get; set; }

        /// <summary>
        /// Project ID, which can be obtained by viewing the project list, if not passed, it will be associated with the default project.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Maximum Qps value.
        /// </summary>
        [Input("qps")]
        public Input<int>? Qps { get; set; }

        /// <summary>
        /// The name of the region where the instance is located, such as ap-shanghai.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security group ID list.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Instance status: 0 creating, 1 process processing, 2 running, 3 instance not initialized, -1 instance isolated, 4 instance initializing, 5 instance deleting, 6 instance restarting, 7 data migration.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Description of the current running state of the instance.
        /// </summary>
        [Input("statusDesc")]
        public Input<string>? StatusDesc { get; set; }

        /// <summary>
        /// Storage size, unit: GB. You can query instance specifications through DescribeDBInstanceSpecs to obtain the lower and upper limits of disk specifications corresponding to different memory sizes.
        /// </summary>
        [Input("storage")]
        public Input<int>? Storage { get; set; }

        /// <summary>
        /// Virtual private network subnet ID, required when VpcId is not empty.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// tag list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// TDSQL version information.
        /// </summary>
        [Input("tdsqlVersion")]
        public Input<string>? TdsqlVersion { get; set; }

        /// <summary>
        /// The account to which the instance belongs.
        /// </summary>
        [Input("uin")]
        public Input<string>? Uin { get; set; }

        /// <summary>
        /// The last update time of the instance in the format of 2006-01-02 15:04:05.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Intranet IP address.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Intranet IPv6.
        /// </summary>
        [Input("vipv6")]
        public Input<string>? Vipv6 { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// A list of voucher IDs. Currently, only one voucher can be specified.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// Virtual private network ID, if not passed, it means that it is created as a basic network.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Intranet port.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        /// <summary>
        /// The domain name accessed from the external network, which can be resolved by the public network.
        /// </summary>
        [Input("wanDomain")]
        public Input<string>? WanDomain { get; set; }

        /// <summary>
        /// Internet port.
        /// </summary>
        [Input("wanPort")]
        public Input<int>? WanPort { get; set; }

        /// <summary>
        /// Internet IPv6 port.
        /// </summary>
        [Input("wanPortIpv6")]
        public Input<int>? WanPortIpv6 { get; set; }

        /// <summary>
        /// External network status, 0-unopened; 1-opened; 2-closed; 3-opening.
        /// </summary>
        [Input("wanStatus")]
        public Input<int>? WanStatus { get; set; }

        /// <summary>
        /// Internet IPv6 status.
        /// </summary>
        [Input("wanStatusIpv6")]
        public Input<int>? WanStatusIpv6 { get; set; }

        /// <summary>
        /// Extranet IP address, accessible from the public network.
        /// </summary>
        [Input("wanVip")]
        public Input<string>? WanVip { get; set; }

        /// <summary>
        /// Internet IPv6.
        /// </summary>
        [Input("wanVipv6")]
        public Input<string>? WanVipv6 { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Instance node availability zone distribution, up to two availability zones can be filled. When the shard specification is one master and two slaves, two of the nodes are in the first availability zone.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public InstanceState()
        {
        }
    }
}
