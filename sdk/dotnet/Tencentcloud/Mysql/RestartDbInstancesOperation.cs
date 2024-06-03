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
    /// Provides a resource to create a mysql restart_db_instances_operation
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
    ///         SlaveDeployMode = 0,
    ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[0]?.Name),
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
    ///     var exampleRestartDbInstancesOperation = new Tencentcloud.Mysql.RestartDbInstancesOperation("exampleRestartDbInstancesOperation", new()
    ///     {
    ///         InstanceId = exampleInstance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// mysql restart_db_instances_operation can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation restart_db_instances_operation restart_db_instances_operation_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation")]
    public partial class RestartDbInstancesOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance status.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RestartDbInstancesOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestartDbInstancesOperation(string name, RestartDbInstancesOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation", name, args ?? new RestartDbInstancesOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestartDbInstancesOperation(string name, Input<string> id, RestartDbInstancesOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RestartDbInstancesOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestartDbInstancesOperation Get(string name, Input<string> id, RestartDbInstancesOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RestartDbInstancesOperation(name, id, state, options);
        }
    }

    public sealed class RestartDbInstancesOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public RestartDbInstancesOperationArgs()
        {
        }
        public static new RestartDbInstancesOperationArgs Empty => new RestartDbInstancesOperationArgs();
    }

    public sealed class RestartDbInstancesOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public RestartDbInstancesOperationState()
        {
        }
        public static new RestartDbInstancesOperationState Empty => new RestartDbInstancesOperationState();
    }
}
