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
    /// Provides a SQL Server PublishSubscribe resource belongs to SQL Server instance.
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
    ///         var examplePubGeneralCloudInstance = new Tencentcloud.Sqlserver.GeneralCloudInstance("examplePubGeneralCloudInstance", new Tencentcloud.Sqlserver.GeneralCloudInstanceArgs
    ///         {
    ///             Zone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_HSSD",
    ///             InstanceChargeType = "POSTPAID",
    ///             ProjectId = 0,
    ///             SubnetId = subnet.Id,
    ///             VpcId = vpc.Id,
    ///             DbVersion = "2008R2",
    ///             SecurityGroupLists = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Weeklies = 
    ///             {
    ///                 1,
    ///                 2,
    ///                 3,
    ///                 5,
    ///                 6,
    ///                 7,
    ///             },
    ///             StartTime = "00:00",
    ///             Span = 6,
    ///             ResourceTags = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.GeneralCloudInstanceResourceTagArgs
    ///                 {
    ///                     TagKey = "test",
    ///                     TagValue = "test",
    ///                 },
    ///             },
    ///             Collation = "Chinese_PRC_CI_AS",
    ///             TimeZone = "China Standard Time",
    ///         });
    ///         var exampleSubGeneralCloudInstance = new Tencentcloud.Sqlserver.GeneralCloudInstance("exampleSubGeneralCloudInstance", new Tencentcloud.Sqlserver.GeneralCloudInstanceArgs
    ///         {
    ///             Zone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_HSSD",
    ///             InstanceChargeType = "POSTPAID",
    ///             ProjectId = 0,
    ///             SubnetId = subnet.Id,
    ///             VpcId = vpc.Id,
    ///             DbVersion = "2008R2",
    ///             SecurityGroupLists = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Weeklies = 
    ///             {
    ///                 1,
    ///                 2,
    ///                 3,
    ///                 5,
    ///                 6,
    ///                 7,
    ///             },
    ///             StartTime = "00:00",
    ///             Span = 6,
    ///             ResourceTags = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.GeneralCloudInstanceResourceTagArgs
    ///                 {
    ///                     TagKey = "test",
    ///                     TagValue = "test",
    ///                 },
    ///             },
    ///             Collation = "Chinese_PRC_CI_AS",
    ///             TimeZone = "China Standard Time",
    ///         });
    ///         var examplePubDb = new Tencentcloud.Sqlserver.Db("examplePubDb", new Tencentcloud.Sqlserver.DbArgs
    ///         {
    ///             InstanceId = examplePubGeneralCloudInstance.Id,
    ///             Charset = "Chinese_PRC_BIN",
    ///             Remark = "test-remark",
    ///         });
    ///         var exampleSubDb = new Tencentcloud.Sqlserver.Db("exampleSubDb", new Tencentcloud.Sqlserver.DbArgs
    ///         {
    ///             InstanceId = exampleSubGeneralCloudInstance.Id,
    ///             Charset = "Chinese_PRC_BIN",
    ///             Remark = "test-remark",
    ///         });
    ///         var example = new Tencentcloud.Sqlserver.PublishSubscribe("example", new Tencentcloud.Sqlserver.PublishSubscribeArgs
    ///         {
    ///             PublishInstanceId = examplePubGeneralCloudInstance.Id,
    ///             SubscribeInstanceId = exampleSubGeneralCloudInstance.Id,
    ///             PublishSubscribeName = "example",
    ///             DeleteSubscribeDb = false,
    ///             DatabaseTuples = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.PublishSubscribeDatabaseTupleArgs
    ///                 {
    ///                     PublishDatabase = examplePubDb.Name,
    ///                     SubscribeDatabase = exampleSubDb.Name,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SQL Server PublishSubscribe can be imported using the publish_sqlserver_id#subscribe_sqlserver_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/publishSubscribe:PublishSubscribe example publish_sqlserver_id#subscribe_sqlserver_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/publishSubscribe:PublishSubscribe")]
    public partial class PublishSubscribe : Pulumi.CustomResource
    {
        /// <summary>
        /// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
        /// </summary>
        [Output("databaseTuples")]
        public Output<ImmutableArray<Outputs.PublishSubscribeDatabaseTuple>> DatabaseTuples { get; private set; } = null!;

        /// <summary>
        /// Whether to delete the subscriber database when deleting the Publish and Subscribe. `true` for deletes the subscribe database, `false` for does not delete the subscribe database. default is `false`.
        /// </summary>
        [Output("deleteSubscribeDb")]
        public Output<bool?> DeleteSubscribeDb { get; private set; } = null!;

        /// <summary>
        /// ID of the SQL Server instance which publish.
        /// </summary>
        [Output("publishInstanceId")]
        public Output<string> PublishInstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the Publish and Subscribe. Default is `default_name`.
        /// </summary>
        [Output("publishSubscribeName")]
        public Output<string?> PublishSubscribeName { get; private set; } = null!;

        /// <summary>
        /// ID of the SQL Server instance which subscribe.
        /// </summary>
        [Output("subscribeInstanceId")]
        public Output<string> SubscribeInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a PublishSubscribe resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublishSubscribe(string name, PublishSubscribeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/publishSubscribe:PublishSubscribe", name, args ?? new PublishSubscribeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublishSubscribe(string name, Input<string> id, PublishSubscribeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/publishSubscribe:PublishSubscribe", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PublishSubscribe resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublishSubscribe Get(string name, Input<string> id, PublishSubscribeState? state = null, CustomResourceOptions? options = null)
        {
            return new PublishSubscribe(name, id, state, options);
        }
    }

    public sealed class PublishSubscribeArgs : Pulumi.ResourceArgs
    {
        [Input("databaseTuples", required: true)]
        private InputList<Inputs.PublishSubscribeDatabaseTupleArgs>? _databaseTuples;

        /// <summary>
        /// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
        /// </summary>
        public InputList<Inputs.PublishSubscribeDatabaseTupleArgs> DatabaseTuples
        {
            get => _databaseTuples ?? (_databaseTuples = new InputList<Inputs.PublishSubscribeDatabaseTupleArgs>());
            set => _databaseTuples = value;
        }

        /// <summary>
        /// Whether to delete the subscriber database when deleting the Publish and Subscribe. `true` for deletes the subscribe database, `false` for does not delete the subscribe database. default is `false`.
        /// </summary>
        [Input("deleteSubscribeDb")]
        public Input<bool>? DeleteSubscribeDb { get; set; }

        /// <summary>
        /// ID of the SQL Server instance which publish.
        /// </summary>
        [Input("publishInstanceId", required: true)]
        public Input<string> PublishInstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the Publish and Subscribe. Default is `default_name`.
        /// </summary>
        [Input("publishSubscribeName")]
        public Input<string>? PublishSubscribeName { get; set; }

        /// <summary>
        /// ID of the SQL Server instance which subscribe.
        /// </summary>
        [Input("subscribeInstanceId", required: true)]
        public Input<string> SubscribeInstanceId { get; set; } = null!;

        public PublishSubscribeArgs()
        {
        }
    }

    public sealed class PublishSubscribeState : Pulumi.ResourceArgs
    {
        [Input("databaseTuples")]
        private InputList<Inputs.PublishSubscribeDatabaseTupleGetArgs>? _databaseTuples;

        /// <summary>
        /// Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.
        /// </summary>
        public InputList<Inputs.PublishSubscribeDatabaseTupleGetArgs> DatabaseTuples
        {
            get => _databaseTuples ?? (_databaseTuples = new InputList<Inputs.PublishSubscribeDatabaseTupleGetArgs>());
            set => _databaseTuples = value;
        }

        /// <summary>
        /// Whether to delete the subscriber database when deleting the Publish and Subscribe. `true` for deletes the subscribe database, `false` for does not delete the subscribe database. default is `false`.
        /// </summary>
        [Input("deleteSubscribeDb")]
        public Input<bool>? DeleteSubscribeDb { get; set; }

        /// <summary>
        /// ID of the SQL Server instance which publish.
        /// </summary>
        [Input("publishInstanceId")]
        public Input<string>? PublishInstanceId { get; set; }

        /// <summary>
        /// The name of the Publish and Subscribe. Default is `default_name`.
        /// </summary>
        [Input("publishSubscribeName")]
        public Input<string>? PublishSubscribeName { get; set; }

        /// <summary>
        /// ID of the SQL Server instance which subscribe.
        /// </summary>
        [Input("subscribeInstanceId")]
        public Input<string>? SubscribeInstanceId { get; set; }

        public PublishSubscribeState()
        {
        }
    }
}
