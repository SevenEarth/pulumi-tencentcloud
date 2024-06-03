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
    /// Provides a resource to create a mysql ro_stop_replication
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zones = Tencentcloud.Availability.GetZonesByProduct.Invoke(new()
    ///     {
    ///         Product = "cdb",
    ///     });
    /// 
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
    ///     {
    ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[0]?.Name),
    ///         VpcId = vpc.Id,
    ///         CidrBlock = "10.0.0.0/16",
    ///         IsMulticast = false,
    ///     });
    /// 
    ///     var securityGroup = new Tencentcloud.Security.Group("securityGroup", new()
    ///     {
    ///         Description = "mysql test",
    ///     });
    /// 
    ///     var exampleInstance = new Tencentcloud.Mysql.Instance("exampleInstance", new()
    ///     {
    ///         InternetService = 1,
    ///         EngineVersion = "5.7",
    ///         ChargeType = "POSTPAID",
    ///         RootPassword = "PassWord123",
    ///         SlaveDeployMode = 1,
    ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[0]?.Name),
    ///         FirstSlaveZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[1]?.Name),
    ///         SlaveSyncMode = 1,
    ///         InstanceName = "tf-example-mysql",
    ///         MemSize = 4000,
    ///         VolumeSize = 200,
    ///         VpcId = vpc.Id,
    ///         SubnetId = subnet.Id,
    ///         IntranetPort = 3306,
    ///         SecurityGroups = new[]
    ///         {
    ///             securityGroup.Id,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "name", "test" },
    ///         },
    ///         Parameters = 
    ///         {
    ///             { "character_set_server", "utf8" },
    ///             { "max_connections", "1000" },
    ///         },
    ///     });
    /// 
    ///     var exampleProxy = new Tencentcloud.Mysql.Proxy("exampleProxy", new()
    ///     {
    ///         InstanceId = exampleInstance.Id,
    ///         UniqVpcId = vpc.Id,
    ///         UniqSubnetId = subnet.Id,
    ///         ProxyNodeCustoms = new[]
    ///         {
    ///             new Tencentcloud.Mysql.Inputs.ProxyProxyNodeCustomArgs
    ///             {
    ///                 NodeCount = 1,
    ///                 Cpu = 2,
    ///                 Mem = 4000,
    ///                 Region = "ap-guangzhou",
    ///                 Zone = "ap-guangzhou-3",
    ///             },
    ///         },
    ///         SecurityGroups = new[]
    ///         {
    ///             securityGroup.Id,
    ///         },
    ///         Desc = "desc.",
    ///         ConnectionPoolLimit = 2,
    ///         Vip = "10.0.0.120",
    ///         Vport = 3306,
    ///     });
    /// 
    ///     var exampleRoStopReplication = new Tencentcloud.Mysql.RoStopReplication("exampleRoStopReplication", new()
    ///     {
    ///         InstanceId = exampleProxy.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/roStopReplication:RoStopReplication")]
    public partial class RoStopReplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Read-Only instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a RoStopReplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoStopReplication(string name, RoStopReplicationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/roStopReplication:RoStopReplication", name, args ?? new RoStopReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoStopReplication(string name, Input<string> id, RoStopReplicationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/roStopReplication:RoStopReplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoStopReplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoStopReplication Get(string name, Input<string> id, RoStopReplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new RoStopReplication(name, id, state, options);
        }
    }

    public sealed class RoStopReplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read-Only instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public RoStopReplicationArgs()
        {
        }
        public static new RoStopReplicationArgs Empty => new RoStopReplicationArgs();
    }

    public sealed class RoStopReplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read-Only instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public RoStopReplicationState()
        {
        }
        public static new RoStopReplicationState Empty => new RoStopReplicationState();
    }
}
