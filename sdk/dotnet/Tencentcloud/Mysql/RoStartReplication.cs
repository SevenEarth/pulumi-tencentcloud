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
    /// Provides a resource to create a mysql ro_start_replication
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
    ///             SlaveDeployMode = 1,
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[0]?.Name),
    ///             FirstSlaveZone = zones.Apply(zones =&gt; zones.Zones?[1]?.Name),
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
    ///         var exampleProxy = new Tencentcloud.Mysql.Proxy("exampleProxy", new Tencentcloud.Mysql.ProxyArgs
    ///         {
    ///             InstanceId = exampleInstance.Id,
    ///             UniqVpcId = vpc.Id,
    ///             UniqSubnetId = subnet.Id,
    ///             ProxyNodeCustoms = 
    ///             {
    ///                 new Tencentcloud.Mysql.Inputs.ProxyProxyNodeCustomArgs
    ///                 {
    ///                     NodeCount = 1,
    ///                     Cpu = 2,
    ///                     Mem = 4000,
    ///                     Region = "ap-guangzhou",
    ///                     Zone = "ap-guangzhou-3",
    ///                 },
    ///             },
    ///             SecurityGroups = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Desc = "desc.",
    ///             ConnectionPoolLimit = 2,
    ///             Vip = "10.0.0.120",
    ///             Vport = 3306,
    ///         });
    ///         var exampleRoStartReplication = new Tencentcloud.Mysql.RoStartReplication("exampleRoStartReplication", new Tencentcloud.Mysql.RoStartReplicationArgs
    ///         {
    ///             InstanceId = exampleProxy.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/roStartReplication:RoStartReplication")]
    public partial class RoStartReplication : Pulumi.CustomResource
    {
        /// <summary>
        /// Read-Only instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a RoStartReplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoStartReplication(string name, RoStartReplicationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/roStartReplication:RoStartReplication", name, args ?? new RoStartReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoStartReplication(string name, Input<string> id, RoStartReplicationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/roStartReplication:RoStartReplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoStartReplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoStartReplication Get(string name, Input<string> id, RoStartReplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new RoStartReplication(name, id, state, options);
        }
    }

    public sealed class RoStartReplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read-Only instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public RoStartReplicationArgs()
        {
        }
    }

    public sealed class RoStartReplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read-Only instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public RoStartReplicationState()
        {
        }
    }
}