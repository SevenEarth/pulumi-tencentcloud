// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    /// <summary>
    /// Provides a resource to create a dcdb cancel_dcn_job_operation
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
    ///         var @internal = Output.Create(Tencentcloud.Security.GetGroups.InvokeAsync(new Tencentcloud.Security.GetGroupsArgs
    ///         {
    ///             Name = "default",
    ///         }));
    ///         var vpc = Output.Create(Tencentcloud.Vpc.GetInstances.InvokeAsync(new Tencentcloud.Vpc.GetInstancesArgs
    ///         {
    ///             Name = "Default-VPC",
    ///         }));
    ///         var subnet = vpc.Apply(vpc =&gt; Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             VpcId = vpc.InstanceLists?[0]?.VpcId,
    ///         })));
    ///         var vpcId = subnet.Apply(subnet =&gt; subnet.InstanceLists?[0]?.VpcId);
    ///         var subnetId = subnet.Apply(subnet =&gt; subnet.InstanceLists?[0]?.SubnetId);
    ///         var sgId = @internal.Apply(@internal =&gt; @internal.SecurityGroups?[0]?.SecurityGroupId);
    ///         var hourdbInstanceDcn = new Tencentcloud.Dcdb.HourdbInstance("hourdbInstanceDcn", new Tencentcloud.Dcdb.HourdbInstanceArgs
    ///         {
    ///             InstanceName = "test_dcdb_db_hourdb_instance_dcn",
    ///             Zones = 
    ///             {
    ///                 @var.Default_az,
    ///             },
    ///             ShardMemory = 2,
    ///             ShardStorage = 10,
    ///             ShardNodeCount = 2,
    ///             ShardCount = 2,
    ///             VpcId = vpcId,
    ///             SubnetId = subnetId,
    ///             SecurityGroupId = sgId,
    ///             DbVersionId = "8.0",
    ///             DcnRegion = "ap-guangzhou",
    ///             DcnInstanceId = local.Dcdb_id,
    ///             ResourceTags = 
    ///             {
    ///                 new Tencentcloud.Dcdb.Inputs.HourdbInstanceResourceTagArgs
    ///                 {
    ///                     TagKey = "aaa",
    ///                     TagValue = "bbb",
    ///                 },
    ///             },
    ///         });
    ///         var dcnDcdbId = hourdbInstanceDcn.Id;
    ///         var cancelOperation = new Tencentcloud.Dcdb.CancelDcnJobOperation("cancelOperation", new Tencentcloud.Dcdb.CancelDcnJobOperationArgs
    ///         {
    ///             InstanceId = dcnDcdbId,
    ///         });
    ///         // cancel the dcn/readonly instance
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dcdb/cancelDcnJobOperation:CancelDcnJobOperation")]
    public partial class CancelDcnJobOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a CancelDcnJobOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CancelDcnJobOperation(string name, CancelDcnJobOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/cancelDcnJobOperation:CancelDcnJobOperation", name, args ?? new CancelDcnJobOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CancelDcnJobOperation(string name, Input<string> id, CancelDcnJobOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/cancelDcnJobOperation:CancelDcnJobOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CancelDcnJobOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CancelDcnJobOperation Get(string name, Input<string> id, CancelDcnJobOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new CancelDcnJobOperation(name, id, state, options);
        }
    }

    public sealed class CancelDcnJobOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public CancelDcnJobOperationArgs()
        {
        }
    }

    public sealed class CancelDcnJobOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public CancelDcnJobOperationState()
        {
        }
    }
}