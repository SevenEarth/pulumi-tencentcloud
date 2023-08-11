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
    /// Provides a resource to create a mysql time_window
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
    ///         var exampleTimeWindow = new Tencentcloud.Mysql.TimeWindow("exampleTimeWindow", new Tencentcloud.Mysql.TimeWindowArgs
    ///         {
    ///             InstanceId = exampleInstance.Id,
    ///             MaxDelayTime = 10,
    ///             TimeRanges = 
    ///             {
    ///                 "01:00-02:01",
    ///             },
    ///             Weekdays = 
    ///             {
    ///                 "friday",
    ///                 "monday",
    ///                 "saturday",
    ///                 "thursday",
    ///                 "tuesday",
    ///                 "wednesday",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mysql time_window can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mysql/timeWindow:TimeWindow time_window instanceId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/timeWindow:TimeWindow")]
    public partial class TimeWindow : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        /// </summary>
        [Output("maxDelayTime")]
        public Output<int?> MaxDelayTime { get; private set; } = null!;

        /// <summary>
        /// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        /// </summary>
        [Output("timeRanges")]
        public Output<ImmutableArray<string>> TimeRanges { get; private set; } = null!;

        /// <summary>
        /// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        /// </summary>
        [Output("weekdays")]
        public Output<ImmutableArray<string>> Weekdays { get; private set; } = null!;


        /// <summary>
        /// Create a TimeWindow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TimeWindow(string name, TimeWindowArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/timeWindow:TimeWindow", name, args ?? new TimeWindowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TimeWindow(string name, Input<string> id, TimeWindowState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/timeWindow:TimeWindow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TimeWindow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TimeWindow Get(string name, Input<string> id, TimeWindowState? state = null, CustomResourceOptions? options = null)
        {
            return new TimeWindow(name, id, state, options);
        }
    }

    public sealed class TimeWindowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        /// </summary>
        [Input("maxDelayTime")]
        public Input<int>? MaxDelayTime { get; set; }

        [Input("timeRanges", required: true)]
        private InputList<string>? _timeRanges;

        /// <summary>
        /// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        /// </summary>
        public InputList<string> TimeRanges
        {
            get => _timeRanges ?? (_timeRanges = new InputList<string>());
            set => _timeRanges = value;
        }

        [Input("weekdays")]
        private InputList<string>? _weekdays;

        /// <summary>
        /// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        /// </summary>
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public TimeWindowArgs()
        {
        }
    }

    public sealed class TimeWindowState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Data delay threshold. It takes effect only for source instance and disaster recovery instance. Default value: 10.
        /// </summary>
        [Input("maxDelayTime")]
        public Input<int>? MaxDelayTime { get; set; }

        [Input("timeRanges")]
        private InputList<string>? _timeRanges;

        /// <summary>
        /// Time period available for maintenance after modification in the format of 10:00-12:00. Each period lasts from half an hour to three hours, with the start time and end time aligned by half-hour. Up to two time periods can be set. Start and end time range: [00:00, 24:00].
        /// </summary>
        public InputList<string> TimeRanges
        {
            get => _timeRanges ?? (_timeRanges = new InputList<string>());
            set => _timeRanges = value;
        }

        [Input("weekdays")]
        private InputList<string>? _weekdays;

        /// <summary>
        /// Specifies for which day to modify the time period. Value range: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. If it is not specified or is left blank, the time period will be modified for every day by default.
        /// </summary>
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public TimeWindowState()
        {
        }
    }
}
