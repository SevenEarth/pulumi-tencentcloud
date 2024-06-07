// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    public static class GetReplicationInstanceSyncStatus
    {
        /// <summary>
        /// Use this data source to query detailed information of tcr replication_instance_sync_status
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var syncStatus = Tencentcloud.Tcr.GetReplicationInstanceSyncStatus.Invoke(new()
        ///     {
        ///         RegistryId = local.Src_registry_id,
        ///         ReplicationRegistryId = local.Dst_registry_id,
        ///         ReplicationRegionId = local.Dst_region_id,
        ///         ShowReplicationLog = false,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetReplicationInstanceSyncStatusResult> InvokeAsync(GetReplicationInstanceSyncStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReplicationInstanceSyncStatusResult>("tencentcloud:Tcr/getReplicationInstanceSyncStatus:getReplicationInstanceSyncStatus", args ?? new GetReplicationInstanceSyncStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tcr replication_instance_sync_status
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var syncStatus = Tencentcloud.Tcr.GetReplicationInstanceSyncStatus.Invoke(new()
        ///     {
        ///         RegistryId = local.Src_registry_id,
        ///         ReplicationRegistryId = local.Dst_registry_id,
        ///         ReplicationRegionId = local.Dst_region_id,
        ///         ShowReplicationLog = false,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetReplicationInstanceSyncStatusResult> Invoke(GetReplicationInstanceSyncStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplicationInstanceSyncStatusResult>("tencentcloud:Tcr/getReplicationInstanceSyncStatus:getReplicationInstanceSyncStatus", args ?? new GetReplicationInstanceSyncStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplicationInstanceSyncStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// master registry id.
        /// </summary>
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        /// <summary>
        /// synchronization instance region id.
        /// </summary>
        [Input("replicationRegionId")]
        public int? ReplicationRegionId { get; set; }

        /// <summary>
        /// synchronization instance id.
        /// </summary>
        [Input("replicationRegistryId", required: true)]
        public string ReplicationRegistryId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// whether to display the synchronization log.
        /// </summary>
        [Input("showReplicationLog")]
        public bool? ShowReplicationLog { get; set; }

        public GetReplicationInstanceSyncStatusArgs()
        {
        }
        public static new GetReplicationInstanceSyncStatusArgs Empty => new GetReplicationInstanceSyncStatusArgs();
    }

    public sealed class GetReplicationInstanceSyncStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// master registry id.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// synchronization instance region id.
        /// </summary>
        [Input("replicationRegionId")]
        public Input<int>? ReplicationRegionId { get; set; }

        /// <summary>
        /// synchronization instance id.
        /// </summary>
        [Input("replicationRegistryId", required: true)]
        public Input<string> ReplicationRegistryId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// whether to display the synchronization log.
        /// </summary>
        [Input("showReplicationLog")]
        public Input<bool>? ShowReplicationLog { get; set; }

        public GetReplicationInstanceSyncStatusInvokeArgs()
        {
        }
        public static new GetReplicationInstanceSyncStatusInvokeArgs Empty => new GetReplicationInstanceSyncStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetReplicationInstanceSyncStatusResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RegistryId;
        /// <summary>
        /// sync log. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplicationInstanceSyncStatusReplicationLogResult> ReplicationLogs;
        public readonly int? ReplicationRegionId;
        public readonly string ReplicationRegistryId;
        /// <summary>
        /// sync status.
        /// </summary>
        public readonly string ReplicationStatus;
        /// <summary>
        /// sync complete time.
        /// </summary>
        public readonly string ReplicationTime;
        public readonly string? ResultOutputFile;
        public readonly bool? ShowReplicationLog;

        [OutputConstructor]
        private GetReplicationInstanceSyncStatusResult(
            string id,

            string registryId,

            ImmutableArray<Outputs.GetReplicationInstanceSyncStatusReplicationLogResult> replicationLogs,

            int? replicationRegionId,

            string replicationRegistryId,

            string replicationStatus,

            string replicationTime,

            string? resultOutputFile,

            bool? showReplicationLog)
        {
            Id = id;
            RegistryId = registryId;
            ReplicationLogs = replicationLogs;
            ReplicationRegionId = replicationRegionId;
            ReplicationRegistryId = replicationRegistryId;
            ReplicationStatus = replicationStatus;
            ReplicationTime = replicationTime;
            ResultOutputFile = resultOutputFile;
            ShowReplicationLog = showReplicationLog;
        }
    }
}
