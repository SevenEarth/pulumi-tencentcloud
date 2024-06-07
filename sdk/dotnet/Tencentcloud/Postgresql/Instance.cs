// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    /// <summary>
    /// Use this resource to create postgresql instance.
    /// 
    /// &gt; **Note:** To update the charge type, please update the `charge_type` and specify the `period` for the charging period. It only supports updating from `POSTPAID_BY_HOUR` to `PREPAID`, and the `period` field only valid in that upgrading case.
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
    ///     var config = new Config();
    ///     var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-1";
    ///     // create vpc
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     // create vpc subnet
    ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
    ///     {
    ///         AvailabilityZone = availabilityZone,
    ///         VpcId = vpc.Id,
    ///         CidrBlock = "10.0.20.0/28",
    ///         IsMulticast = false,
    ///     });
    /// 
    ///     // create postgresql
    ///     var foo = new Tencentcloud.Postgresql.Instance("foo", new()
    ///     {
    ///         AvailabilityZone = availabilityZone,
    ///         ChargeType = "POSTPAID_BY_HOUR",
    ///         VpcId = vpc.Id,
    ///         SubnetId = subnet.Id,
    ///         EngineVersion = "10.4",
    ///         RootUser = "root123",
    ///         RootPassword = "Root123$",
    ///         Charset = "UTF8",
    ///         ProjectId = 0,
    ///         Memory = 2,
    ///         Storage = 10,
    ///         Tags = 
    ///         {
    ///             { "test", "tf" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Create a multi available zone bucket
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
    ///     var config = new Config();
    ///     var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-6";
    ///     var standbyAvailabilityZone = config.Get("standbyAvailabilityZone") ?? "ap-guangzhou-7";
    ///     // create vpc
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     // create vpc subnet
    ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
    ///     {
    ///         AvailabilityZone = availabilityZone,
    ///         VpcId = vpc.Id,
    ///         CidrBlock = "10.0.20.0/28",
    ///         IsMulticast = false,
    ///     });
    /// 
    ///     // create postgresql
    ///     var foo = new Tencentcloud.Postgresql.Instance("foo", new()
    ///     {
    ///         AvailabilityZone = availabilityZone,
    ///         ChargeType = "POSTPAID_BY_HOUR",
    ///         VpcId = vpc.Id,
    ///         SubnetId = subnet.Id,
    ///         EngineVersion = "10.4",
    ///         RootUser = "root123",
    ///         RootPassword = "Root123$",
    ///         Charset = "UTF8",
    ///         ProjectId = 0,
    ///         Memory = 2,
    ///         Cpu = 1,
    ///         Storage = 10,
    ///         DbNodeSets = new[]
    ///         {
    ///             new Tencentcloud.Postgresql.Inputs.InstanceDbNodeSetArgs
    ///             {
    ///                 Role = "Primary",
    ///                 Zone = availabilityZone,
    ///             },
    ///             new Tencentcloud.Postgresql.Inputs.InstanceDbNodeSetArgs
    ///             {
    ///                 Zone = standbyAvailabilityZone,
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "test", "tf" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### create pgsql with kms key
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
    ///     var pg = new Tencentcloud.Postgresql.Instance("pg", new()
    ///     {
    ///         AvailabilityZone = "ap-guangzhou-6",
    ///         BackupPlan = new Tencentcloud.Postgresql.Inputs.InstanceBackupPlanArgs
    ///         {
    ///             BackupPeriods = new[]
    ///             {
    ///                 "tuesday",
    ///                 "wednesday",
    ///             },
    ///             BaseBackupRetentionPeriod = 7,
    ///             MaxBackupStartTime = "01:10:11",
    ///             MinBackupStartTime = "00:10:11",
    ///         },
    ///         ChargeType = "POSTPAID_BY_HOUR",
    ///         Charset = "LATIN1",
    ///         DbKernelVersion = "v11.12_r1.3",
    ///         EngineVersion = "11.12",
    ///         KmsKeyId = "788c606a-c7b7-11ec-82d1-5254001e5c4e",
    ///         KmsRegion = "ap-guangzhou",
    ///         Memory = 4,
    ///         NeedSupportTde = 1,
    ///         ProjectId = 0,
    ///         RootPassword = "xxxxxxxxxx",
    ///         Storage = 100,
    ///         SubnetId = "subnet-enm92y0m",
    ///         Tags = 
    ///         {
    ///             { "tf", "test" },
    ///         },
    ///         VpcId = "vpc-86v957zb",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### upgrade kernel version
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
    ///     var test = new Tencentcloud.Postgresql.Instance("test", new()
    ///     {
    ///         AvailabilityZone = data.Tencentcloud_availability_zones_by_product.Zone.Zones[5].Name,
    ///         ChargeType = "POSTPAID_BY_HOUR",
    ///         VpcId = local.Vpc_id,
    ///         SubnetId = local.Subnet_id,
    ///         EngineVersion = "13.3",
    ///         RootPassword = "*",
    ///         Charset = "LATIN1",
    ///         ProjectId = 0,
    ///         PublicAccessSwitch = false,
    ///         SecurityGroups = new[]
    ///         {
    ///             local.Sg_id,
    ///         },
    ///         Memory = 4,
    ///         Storage = 250,
    ///         BackupPlan = new Tencentcloud.Postgresql.Inputs.InstanceBackupPlanArgs
    ///         {
    ///             MinBackupStartTime = "01:10:11",
    ///             MaxBackupStartTime = "02:10:11",
    ///             BaseBackupRetentionPeriod = 5,
    ///             BackupPeriods = new[]
    ///             {
    ///                 "monday",
    ///                 "thursday",
    ///                 "sunday",
    ///             },
    ///         },
    ///         DbKernelVersion = "v13.3_r1.4",
    ///         Tags = 
    ///         {
    ///             { "tf", "teest" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// postgresql instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Postgresql/instance:Instance foo postgres-cda1iex1
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int?> AutoRenewFlag { get; private set; } = null!;

        /// <summary>
        /// Whether to use voucher, `1` for enabled.
        /// </summary>
        [Output("autoVoucher")]
        public Output<int?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Availability zone. NOTE: This field could not be modified, please use `db_node_set` instead of modification. The changes on this field will be suppressed when using the `db_node_set`.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Specify DB backup plan.
        /// </summary>
        [Output("backupPlan")]
        public Output<Outputs.InstanceBackupPlan?> BackupPlan { get; private set; } = null!;

        /// <summary>
        /// Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// Charset of the root account. Valid values are `UTF8`,`LATIN1`.
        /// </summary>
        [Output("charset")]
        public Output<string?> Charset { get; private set; } = null!;

        /// <summary>
        /// Number of CPU cores. Allowed value must be equal `cpu` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// Create time of the postgresql instance.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
        /// </summary>
        [Output("dbKernelVersion")]
        public Output<string> DbKernelVersion { get; private set; } = null!;

        /// <summary>
        /// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
        /// </summary>
        [Output("dbMajorVersion")]
        public Output<string> DbMajorVersion { get; private set; } = null!;

        /// <summary>
        /// `db_major_vesion` will be deprecated, use `db_major_version` instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
        /// </summary>
        [Output("dbMajorVesion")]
        public Output<string> DbMajorVesion { get; private set; } = null!;

        /// <summary>
        /// Specify instance node info for disaster migration.
        /// </summary>
        [Output("dbNodeSets")]
        public Output<ImmutableArray<Outputs.InstanceDbNodeSet>> DbNodeSets { get; private set; } = null!;

        /// <summary>
        /// Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string?> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// KeyId of the custom key.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Region of the custom key.
        /// </summary>
        [Output("kmsRegion")]
        public Output<string> KmsRegion { get; private set; } = null!;

        /// <summary>
        /// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
        /// </summary>
        [Output("maxStandbyArchiveDelay")]
        public Output<int> MaxStandbyArchiveDelay { get; private set; } = null!;

        /// <summary>
        /// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
        /// </summary>
        [Output("maxStandbyStreamingDelay")]
        public Output<int> MaxStandbyStreamingDelay { get; private set; } = null!;

        /// <summary>
        /// Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Name of the postgresql instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether to support data transparent encryption, 1: yes, 0: no (default).
        /// </summary>
        [Output("needSupportTde")]
        public Output<int> NeedSupportTde { get; private set; } = null!;

        /// <summary>
        /// Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// IP for private access.
        /// </summary>
        [Output("privateAccessIp")]
        public Output<string> PrivateAccessIp { get; private set; } = null!;

        /// <summary>
        /// Port for private access.
        /// </summary>
        [Output("privateAccessPort")]
        public Output<int> PrivateAccessPort { get; private set; } = null!;

        /// <summary>
        /// Project id, default value is `0`.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Host for public access.
        /// </summary>
        [Output("publicAccessHost")]
        public Output<string> PublicAccessHost { get; private set; } = null!;

        /// <summary>
        /// Port for public access.
        /// </summary>
        [Output("publicAccessPort")]
        public Output<int> PublicAccessPort { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to enable the access to an instance from public network or not.
        /// </summary>
        [Output("publicAccessSwitch")]
        public Output<bool?> PublicAccessSwitch { get; private set; } = null!;

        /// <summary>
        /// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
        /// </summary>
        [Output("rootPassword")]
        public Output<string> RootPassword { get; private set; } = null!;

        /// <summary>
        /// Instance root account name. This parameter is optional, Default value is `root`.
        /// </summary>
        [Output("rootUser")]
        public Output<string?> RootUser { get; private set; } = null!;

        /// <summary>
        /// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Output("storage")]
        public Output<int> Storage { get; private set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The available tags within this postgresql.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Uid of the postgresql instance.
        /// </summary>
        [Output("uid")]
        public Output<int> Uid { get; private set; } = null!;

        /// <summary>
        /// Specify Voucher Ids if `auto_voucher` was `1`, only support using 1 vouchers for now.
        /// </summary>
        [Output("voucherIds")]
        public Output<ImmutableArray<string>> VoucherIds { get; private set; } = null!;

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
                AdditionalSecretOutputs =
                {
                    "rootPassword",
                },
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

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Whether to use voucher, `1` for enabled.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Availability zone. NOTE: This field could not be modified, please use `db_node_set` instead of modification. The changes on this field will be suppressed when using the `db_node_set`.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// Specify DB backup plan.
        /// </summary>
        [Input("backupPlan")]
        public Input<Inputs.InstanceBackupPlanArgs>? BackupPlan { get; set; }

        /// <summary>
        /// Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Charset of the root account. Valid values are `UTF8`,`LATIN1`.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// Number of CPU cores. Allowed value must be equal `cpu` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
        /// </summary>
        [Input("dbKernelVersion")]
        public Input<string>? DbKernelVersion { get; set; }

        /// <summary>
        /// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
        /// </summary>
        [Input("dbMajorVersion")]
        public Input<string>? DbMajorVersion { get; set; }

        /// <summary>
        /// `db_major_vesion` will be deprecated, use `db_major_version` instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
        /// </summary>
        [Input("dbMajorVesion")]
        public Input<string>? DbMajorVesion { get; set; }

        [Input("dbNodeSets")]
        private InputList<Inputs.InstanceDbNodeSetArgs>? _dbNodeSets;

        /// <summary>
        /// Specify instance node info for disaster migration.
        /// </summary>
        public InputList<Inputs.InstanceDbNodeSetArgs> DbNodeSets
        {
            get => _dbNodeSets ?? (_dbNodeSets = new InputList<Inputs.InstanceDbNodeSetArgs>());
            set => _dbNodeSets = value;
        }

        /// <summary>
        /// Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// KeyId of the custom key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Region of the custom key.
        /// </summary>
        [Input("kmsRegion")]
        public Input<string>? KmsRegion { get; set; }

        /// <summary>
        /// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
        /// </summary>
        [Input("maxStandbyArchiveDelay")]
        public Input<int>? MaxStandbyArchiveDelay { get; set; }

        /// <summary>
        /// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
        /// </summary>
        [Input("maxStandbyStreamingDelay")]
        public Input<int>? MaxStandbyStreamingDelay { get; set; }

        /// <summary>
        /// Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Name of the postgresql instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to support data transparent encryption, 1: yes, 0: no (default).
        /// </summary>
        [Input("needSupportTde")]
        public Input<int>? NeedSupportTde { get; set; }

        /// <summary>
        /// Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Project id, default value is `0`.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Indicates whether to enable the access to an instance from public network or not.
        /// </summary>
        [Input("publicAccessSwitch")]
        public Input<bool>? PublicAccessSwitch { get; set; }

        [Input("rootPassword", required: true)]
        private Input<string>? _rootPassword;

        /// <summary>
        /// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
        /// </summary>
        public Input<string>? RootPassword
        {
            get => _rootPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _rootPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Instance root account name. This parameter is optional, Default value is `root`.
        /// </summary>
        [Input("rootUser")]
        public Input<string>? RootUser { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Input("storage", required: true)]
        public Input<int> Storage { get; set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The available tags within this postgresql.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// Specify Voucher Ids if `auto_voucher` was `1`, only support using 1 vouchers for now.
        /// </summary>
        public InputList<string> VoucherIds
        {
            get => _voucherIds ?? (_voucherIds = new InputList<string>());
            set => _voucherIds = value;
        }

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag, `1` for enabled. NOTES: Only support prepaid instance.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Whether to use voucher, `1` for enabled.
        /// </summary>
        [Input("autoVoucher")]
        public Input<int>? AutoVoucher { get; set; }

        /// <summary>
        /// Availability zone. NOTE: This field could not be modified, please use `db_node_set` instead of modification. The changes on this field will be suppressed when using the `db_node_set`.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Specify DB backup plan.
        /// </summary>
        [Input("backupPlan")]
        public Input<Inputs.InstanceBackupPlanGetArgs>? BackupPlan { get; set; }

        /// <summary>
        /// Pay type of the postgresql instance. Values `POSTPAID_BY_HOUR` (Default), `PREPAID`. It only support to update the type from `POSTPAID_BY_HOUR` to `PREPAID`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Charset of the root account. Valid values are `UTF8`,`LATIN1`.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// Number of CPU cores. Allowed value must be equal `cpu` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// Create time of the postgresql instance.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created. It supports updating the minor kernel version immediately.
        /// </summary>
        [Input("dbKernelVersion")]
        public Input<string>? DbKernelVersion { get; set; }

        /// <summary>
        /// PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
        /// </summary>
        [Input("dbMajorVersion")]
        public Input<string>? DbMajorVersion { get; set; }

        /// <summary>
        /// `db_major_vesion` will be deprecated, use `db_major_version` instead. PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.
        /// </summary>
        [Input("dbMajorVesion")]
        public Input<string>? DbMajorVesion { get; set; }

        [Input("dbNodeSets")]
        private InputList<Inputs.InstanceDbNodeSetGetArgs>? _dbNodeSets;

        /// <summary>
        /// Specify instance node info for disaster migration.
        /// </summary>
        public InputList<Inputs.InstanceDbNodeSetGetArgs> DbNodeSets
        {
            get => _dbNodeSets ?? (_dbNodeSets = new InputList<Inputs.InstanceDbNodeSetGetArgs>());
            set => _dbNodeSets = value;
        }

        /// <summary>
        /// Version of the postgresql database engine. Valid values: `10.4`, `11.8`, `12.4`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// KeyId of the custom key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Region of the custom key.
        /// </summary>
        [Input("kmsRegion")]
        public Input<string>? KmsRegion { get; set; }

        /// <summary>
        /// max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.
        /// </summary>
        [Input("maxStandbyArchiveDelay")]
        public Input<int>? MaxStandbyArchiveDelay { get; set; }

        /// <summary>
        /// max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.
        /// </summary>
        [Input("maxStandbyStreamingDelay")]
        public Input<int>? MaxStandbyStreamingDelay { get; set; }

        /// <summary>
        /// Memory size(in GB). Allowed value must be larger than `memory` that data source `tencentcloud.Postgresql.getSpecinfos` provides.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Name of the postgresql instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to support data transparent encryption, 1: yes, 0: no (default).
        /// </summary>
        [Input("needSupportTde")]
        public Input<int>? NeedSupportTde { get; set; }

        /// <summary>
        /// Specify Prepaid period in month. Default `1`. Values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`. This field is valid only when creating a `PREPAID` type instance, or updating the charge type from `POSTPAID_BY_HOUR` to `PREPAID`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// IP for private access.
        /// </summary>
        [Input("privateAccessIp")]
        public Input<string>? PrivateAccessIp { get; set; }

        /// <summary>
        /// Port for private access.
        /// </summary>
        [Input("privateAccessPort")]
        public Input<int>? PrivateAccessPort { get; set; }

        /// <summary>
        /// Project id, default value is `0`.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Host for public access.
        /// </summary>
        [Input("publicAccessHost")]
        public Input<string>? PublicAccessHost { get; set; }

        /// <summary>
        /// Port for public access.
        /// </summary>
        [Input("publicAccessPort")]
        public Input<int>? PublicAccessPort { get; set; }

        /// <summary>
        /// Indicates whether to enable the access to an instance from public network or not.
        /// </summary>
        [Input("publicAccessSwitch")]
        public Input<bool>? PublicAccessSwitch { get; set; }

        [Input("rootPassword")]
        private Input<string>? _rootPassword;

        /// <summary>
        /// Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.
        /// </summary>
        public Input<string>? RootPassword
        {
            get => _rootPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _rootPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Instance root account name. This parameter is optional, Default value is `root`.
        /// </summary>
        [Input("rootUser")]
        public Input<string>? RootUser { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of `storage_min` and `storage_max` which data source `tencentcloud.Postgresql.getSpecinfos` provides.
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
        /// The available tags within this postgresql.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Uid of the postgresql instance.
        /// </summary>
        [Input("uid")]
        public Input<int>? Uid { get; set; }

        [Input("voucherIds")]
        private InputList<string>? _voucherIds;

        /// <summary>
        /// Specify Voucher Ids if `auto_voucher` was `1`, only support using 1 vouchers for now.
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

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
