// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a mysql instance resource to create read-only database instances.
    /// 
    /// &gt; **NOTE:** Read-only instances can be purchased only for two-node or three-node source instances on MySQL 5.6 or above with the InnoDB engine at a specification of 1 GB memory and 50 GB disk capacity or above.
    /// **NOTE:** The terminate operation of read only mysql does NOT take effect immediately, maybe takes for several hours. so during that time, VPCs associated with that mysql instance can't be terminated also.
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
    ///         var @default = new Tencentcloud.Mysql.ReadonlyInstance("default", new Tencentcloud.Mysql.ReadonlyInstanceArgs
    ///         {
    ///             InstanceName = "myTestMysql",
    ///             IntranetPort = 3306,
    ///             MasterInstanceId = "cdb-dnqksd9f",
    ///             MemSize = 128000,
    ///             SecurityGroups = 
    ///             {
    ///                 "sg-ot8eclwz",
    ///             },
    ///             SubnetId = "subnet-9uivyb1g",
    ///             Tags = 
    ///             {
    ///                 { "name", "test" },
    ///             },
    ///             VolumeSize = 255,
    ///             VpcId = "vpc-12mt3l31",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/readonlyInstance:ReadonlyInstance")]
    public partial class ReadonlyInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew flag. NOTES: Only supported prepaid instance.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int?> AutoRenewFlag { get; private set; } = null!;

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
        public Output<string?> DeviceType { get; private set; } = null!;

        /// <summary>
        /// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
        /// </summary>
        [Output("fastUpgrade")]
        public Output<int?> FastUpgrade { get; private set; } = null!;

        /// <summary>
        /// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
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
        /// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
        /// </summary>
        [Output("locked")]
        public Output<int> Locked { get; private set; } = null!;

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
        /// Specify parameter template id.
        /// </summary>
        [Output("paramTemplateId")]
        public Output<int?> ParamTemplateId { get; private set; } = null!;

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
        /// Security groups to use.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Private network ID. If `vpc_id` is set, this value is required.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Instance tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Indicates which kind of operations is being executed.
        /// </summary>
        [Output("taskStatus")]
        public Output<int> TaskStatus { get; private set; } = null!;

        /// <summary>
        /// Disk size (in GB).
        /// </summary>
        [Output("volumeSize")]
        public Output<int> VolumeSize { get; private set; } = null!;

        /// <summary>
        /// ID of VPC, which can be modified once every 24 hours and can't be removed.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a ReadonlyInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadonlyInstance(string name, ReadonlyInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/readonlyInstance:ReadonlyInstance", name, args ?? new ReadonlyInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadonlyInstance(string name, Input<string> id, ReadonlyInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/readonlyInstance:ReadonlyInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReadonlyInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReadonlyInstance Get(string name, Input<string> id, ReadonlyInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ReadonlyInstance(name, id, state, options);
        }
    }

    public sealed class ReadonlyInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

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
        /// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
        /// </summary>
        [Input("fastUpgrade")]
        public Input<int>? FastUpgrade { get; set; }

        /// <summary>
        /// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
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
        [Input("masterRegion")]
        public Input<string>? MasterRegion { get; set; }

        /// <summary>
        /// Memory size (in MB).
        /// </summary>
        [Input("memSize", required: true)]
        public Input<int> MemSize { get; set; } = null!;

        /// <summary>
        /// Specify parameter template id.
        /// </summary>
        [Input("paramTemplateId")]
        public Input<int>? ParamTemplateId { get; set; }

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

        /// <summary>
        /// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ReadonlyInstanceArgs()
        {
        }
    }

    public sealed class ReadonlyInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew flag. NOTES: Only supported prepaid instance.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

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
        /// Specify whether to enable fast upgrade when upgrade instance spec, available value: `1` - enabled, `0` - disabled.
        /// </summary>
        [Input("fastUpgrade")]
        public Input<int>? FastUpgrade { get; set; }

        /// <summary>
        /// Indicate whether to delete instance directly or not. Default is `false`. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.
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
        /// Indicates whether the instance is locked. Valid values: `0`, `1`. `0` - No; `1` - Yes.
        /// </summary>
        [Input("locked")]
        public Input<int>? Locked { get; set; }

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
        /// Specify parameter template id.
        /// </summary>
        [Input("paramTemplateId")]
        public Input<int>? ParamTemplateId { get; set; }

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
        /// Instance status. Valid values: `0`, `1`, `4`, `5`. `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

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
        /// Indicates which kind of operations is being executed.
        /// </summary>
        [Input("taskStatus")]
        public Input<int>? TaskStatus { get; set; }

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

        /// <summary>
        /// Zone information, this parameter defaults to, the system automatically selects an Availability Zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ReadonlyInstanceState()
        {
        }
    }
}
