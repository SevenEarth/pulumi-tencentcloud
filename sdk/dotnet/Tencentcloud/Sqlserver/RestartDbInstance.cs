// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    /// <summary>
    /// Provides a resource to create a sqlserver restart_db_instance
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
    ///             Product = "sqlserver",
    ///         }));
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             VpcId = vpc.Id,
    ///             CidrBlock = "10.0.0.0/16",
    ///             IsMulticast = false,
    ///         });
    ///         var securityGroup = new Tencentcloud.Security.Group("securityGroup", new Tencentcloud.Security.GroupArgs
    ///         {
    ///             Description = "desc.",
    ///         });
    ///         var exampleBasicInstance = new Tencentcloud.Sqlserver.BasicInstance("exampleBasicInstance", new Tencentcloud.Sqlserver.BasicInstanceArgs
    ///         {
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             ChargeType = "POSTPAID_BY_HOUR",
    ///             VpcId = vpc.Id,
    ///             SubnetId = subnet.Id,
    ///             ProjectId = 0,
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
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
    ///                 securityGroup.Id,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///         });
    ///         var exampleRestartDbInstance = new Tencentcloud.Sqlserver.RestartDbInstance("exampleRestartDbInstance", new Tencentcloud.Sqlserver.RestartDbInstanceArgs
    ///         {
    ///             InstanceId = exampleBasicInstance.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver restart_db_instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/restartDbInstance:RestartDbInstance example mssql-i9ma6oy7
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/restartDbInstance:RestartDbInstance")]
    public partial class RestartDbInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a RestartDbInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestartDbInstance(string name, RestartDbInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/restartDbInstance:RestartDbInstance", name, args ?? new RestartDbInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestartDbInstance(string name, Input<string> id, RestartDbInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/restartDbInstance:RestartDbInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RestartDbInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestartDbInstance Get(string name, Input<string> id, RestartDbInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new RestartDbInstance(name, id, state, options);
        }
    }

    public sealed class RestartDbInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public RestartDbInstanceArgs()
        {
        }
    }

    public sealed class RestartDbInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public RestartDbInstanceState()
        {
        }
    }
}
