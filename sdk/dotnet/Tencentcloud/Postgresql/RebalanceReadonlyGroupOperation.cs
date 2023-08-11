// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    /// <summary>
    /// Provides a resource to create a postgresql rebalance_readonly_group_operation
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
    ///         var groupRebalance = new Tencentcloud.Postgresql.ReadonlyGroup("groupRebalance", new Tencentcloud.Postgresql.ReadonlyGroupArgs
    ///         {
    ///             MasterDbInstanceId = local.Pgsql_id,
    ///             ProjectId = 0,
    ///             VpcId = "vpc-86v957zb",
    ///             SubnetId = "subnet-enm92y0m",
    ///             ReplayLagEliminate = 1,
    ///             ReplayLatencyEliminate = 1,
    ///             MaxReplayLag = 100,
    ///             MaxReplayLatency = 512,
    ///             MinDelayEliminateReserve = 1,
    ///         });
    ///         var rebalanceReadonlyGroupOperation = new Tencentcloud.Postgresql.RebalanceReadonlyGroupOperation("rebalanceReadonlyGroupOperation", new Tencentcloud.Postgresql.RebalanceReadonlyGroupOperationArgs
    ///         {
    ///             ReadOnlyGroupId = groupRebalance.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/rebalanceReadonlyGroupOperation:RebalanceReadonlyGroupOperation")]
    public partial class RebalanceReadonlyGroupOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// readonly Group ID.
        /// </summary>
        [Output("readOnlyGroupId")]
        public Output<string> ReadOnlyGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a RebalanceReadonlyGroupOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RebalanceReadonlyGroupOperation(string name, RebalanceReadonlyGroupOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/rebalanceReadonlyGroupOperation:RebalanceReadonlyGroupOperation", name, args ?? new RebalanceReadonlyGroupOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RebalanceReadonlyGroupOperation(string name, Input<string> id, RebalanceReadonlyGroupOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/rebalanceReadonlyGroupOperation:RebalanceReadonlyGroupOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RebalanceReadonlyGroupOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RebalanceReadonlyGroupOperation Get(string name, Input<string> id, RebalanceReadonlyGroupOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RebalanceReadonlyGroupOperation(name, id, state, options);
        }
    }

    public sealed class RebalanceReadonlyGroupOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// readonly Group ID.
        /// </summary>
        [Input("readOnlyGroupId", required: true)]
        public Input<string> ReadOnlyGroupId { get; set; } = null!;

        public RebalanceReadonlyGroupOperationArgs()
        {
        }
    }

    public sealed class RebalanceReadonlyGroupOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// readonly Group ID.
        /// </summary>
        [Input("readOnlyGroupId")]
        public Input<string>? ReadOnlyGroupId { get; set; }

        public RebalanceReadonlyGroupOperationState()
        {
        }
    }
}