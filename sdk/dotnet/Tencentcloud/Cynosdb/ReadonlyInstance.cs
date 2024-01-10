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
    /// Provide a resource to create a CynosDB readonly instance.
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
    ///         var foo = new Tencentcloud.Cynosdb.ReadonlyInstance("foo", new Tencentcloud.Cynosdb.ReadonlyInstanceArgs
    ///         {
    ///             ClusterId = cynosdbmysql_dzj5l8gz,
    ///             InstanceName = "tf-cynosdb-readonly-instance",
    ///             ForceDelete = true,
    ///             InstanceCpuCore = 2,
    ///             InstanceMemorySize = 4,
    ///             InstanceMaintainDuration = 7200,
    ///             InstanceMaintainStartTime = 21600,
    ///             InstanceMaintainWeekdays = 
    ///             {
    ///                 "Fri",
    ///                 "Mon",
    ///                 "Sat",
    ///                 "Sun",
    ///                 "Thu",
    ///                 "Wed",
    ///                 "Tue",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CynosDB readonly instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance foo cynosdbmysql-ins-dhwynib6
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance")]
    public partial class ReadonlyInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID which the readonly instance belongs to.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Output("instanceCpuCore")]
        public Output<int?> InstanceCpuCore { get; private set; } = null!;

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
        /// Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Output("instanceMemorySize")]
        public Output<int?> InstanceMemorySize { get; private set; } = null!;

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
        /// ID of the subnet within this VPC.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a ReadonlyInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadonlyInstance(string name, ReadonlyInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance", name, args ?? new ReadonlyInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadonlyInstance(string name, Input<string> id, ReadonlyInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance", name, state, MakeResourceOptions(options, id))
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
        /// Cluster ID which the readonly instance belongs to.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceCpuCore")]
        public Input<int>? InstanceCpuCore { get; set; }

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
        /// Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceMemorySize")]
        public Input<int>? InstanceMemorySize { get; set; }

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// ID of the subnet within this VPC.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ReadonlyInstanceArgs()
        {
        }
    }

    public sealed class ReadonlyInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID which the readonly instance belongs to.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        /// </summary>
        [Input("instanceCpuCore")]
        public Input<int>? InstanceCpuCore { get; set; }

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
        /// Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
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

        /// <summary>
        /// ID of the subnet within this VPC.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ReadonlyInstanceState()
        {
        }
    }
}
