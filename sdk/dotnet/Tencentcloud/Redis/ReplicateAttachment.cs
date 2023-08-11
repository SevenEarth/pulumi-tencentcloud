// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis
{
    /// <summary>
    /// Provides a resource to create a redis replicate_attachment
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
    ///         var zone = Output.Create(Tencentcloud.Redis.GetZoneConfig.InvokeAsync(new Tencentcloud.Redis.GetZoneConfigArgs
    ///         {
    ///             TypeId = 7,
    ///             Region = "ap-guangzhou",
    ///         }));
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             VpcId = vpc.Id,
    ///             AvailabilityZone = zone.Apply(zone =&gt; zone.Lists?[2]?.Zone),
    ///             CidrBlock = "10.0.1.0/24",
    ///         });
    ///         var fooGroup = new Tencentcloud.Security.Group("fooGroup", new Tencentcloud.Security.GroupArgs
    ///         {
    ///         });
    ///         var fooGroupLiteRule = new Tencentcloud.Security.GroupLiteRule("fooGroupLiteRule", new Tencentcloud.Security.GroupLiteRuleArgs
    ///         {
    ///             SecurityGroupId = fooGroup.Id,
    ///             Ingresses = 
    ///             {
    ///                 "ACCEPT#192.168.1.0/24#80#TCP",
    ///                 "DROP#8.8.8.8#80,90#UDP",
    ///                 "DROP#0.0.0.0/0#80-90#TCP",
    ///             },
    ///             Egresses = 
    ///             {
    ///                 "ACCEPT#192.168.0.0/16#ALL#TCP",
    ///                 "ACCEPT#10.0.0.0/8#ALL#ICMP",
    ///                 "DROP#0.0.0.0/0#ALL#ALL",
    ///             },
    ///         });
    ///         var fooInstance = new Tencentcloud.Redis.Instance("fooInstance", new Tencentcloud.Redis.InstanceArgs
    ///         {
    ///             AvailabilityZone = zone.Apply(zone =&gt; zone.Lists?[2]?.Zone),
    ///             TypeId = zone.Apply(zone =&gt; zone.Lists?[2]?.TypeId),
    ///             Password = "test12345789",
    ///             MemSize = 8192,
    ///             RedisShardNum = zone.Apply(zone =&gt; zone.Lists?[2]?.RedisShardNums?[0]),
    ///             RedisReplicasNum = zone.Apply(zone =&gt; zone.Lists?[2]?.RedisReplicasNums?[0]),
    ///             Port = 6379,
    ///             VpcId = vpc.Id,
    ///             SubnetId = subnet.Id,
    ///             SecurityGroups = 
    ///             {
    ///                 fooGroup.Id,
    ///             },
    ///         });
    ///         var instance = new Tencentcloud.Redis.Instance("instance", new Tencentcloud.Redis.InstanceArgs
    ///         {
    ///             AvailabilityZone = zone.Apply(zone =&gt; zone.Lists?[2]?.Zone),
    ///             TypeId = zone.Apply(zone =&gt; zone.Lists?[2]?.TypeId),
    ///             Password = "test12345789",
    ///             MemSize = 8192,
    ///             RedisShardNum = zone.Apply(zone =&gt; zone.Lists?[2]?.RedisShardNums?[0]),
    ///             RedisReplicasNum = zone.Apply(zone =&gt; zone.Lists?[2]?.RedisReplicasNums?[0]),
    ///             Port = 6379,
    ///             VpcId = vpc.Id,
    ///             SubnetId = subnet.Id,
    ///             SecurityGroups = 
    ///             {
    ///                 fooGroup.Id,
    ///             },
    ///         });
    ///         var replicateAttachment = new Tencentcloud.Redis.ReplicateAttachment("replicateAttachment", new Tencentcloud.Redis.ReplicateAttachmentArgs
    ///         {
    ///             GroupId = "crs-rpl-orfiwmn5",
    ///             MasterInstanceId = fooInstance.Id,
    ///             InstanceIds = 
    ///             {
    ///                 instance.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// redis replicate_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Redis/replicateAttachment:ReplicateAttachment replicate_attachment replicate_attachment_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Redis/replicateAttachment:ReplicateAttachment")]
    public partial class ReplicateAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// All instance ids of the replication group.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// The ID of master instance.
        /// </summary>
        [Output("masterInstanceId")]
        public Output<string> MasterInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicateAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicateAttachment(string name, ReplicateAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/replicateAttachment:ReplicateAttachment", name, args ?? new ReplicateAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicateAttachment(string name, Input<string> id, ReplicateAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/replicateAttachment:ReplicateAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicateAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicateAttachment Get(string name, Input<string> id, ReplicateAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicateAttachment(name, id, state, options);
        }
    }

    public sealed class ReplicateAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// All instance ids of the replication group.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// The ID of master instance.
        /// </summary>
        [Input("masterInstanceId", required: true)]
        public Input<string> MasterInstanceId { get; set; } = null!;

        public ReplicateAttachmentArgs()
        {
        }
    }

    public sealed class ReplicateAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// All instance ids of the replication group.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// The ID of master instance.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        public ReplicateAttachmentState()
        {
        }
    }
}