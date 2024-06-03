// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse
{
    public static class GetInstanceShards
    {
        /// <summary>
        /// Use this data source to query detailed information of clickhouse instance_shards
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
        ///     var instanceShards = Tencentcloud.Clickhouse.GetInstanceShards.Invoke(new()
        ///     {
        ///         InstanceId = "cdwch-datuhk3z",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInstanceShardsResult> InvokeAsync(GetInstanceShardsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceShardsResult>("tencentcloud:Clickhouse/getInstanceShards:getInstanceShards", args ?? new GetInstanceShardsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clickhouse instance_shards
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
        ///     var instanceShards = Tencentcloud.Clickhouse.GetInstanceShards.Invoke(new()
        ///     {
        ///         InstanceId = "cdwch-datuhk3z",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInstanceShardsResult> Invoke(GetInstanceShardsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceShardsResult>("tencentcloud:Clickhouse/getInstanceShards:getInstanceShards", args ?? new GetInstanceShardsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceShardsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceShardsArgs()
        {
        }
        public static new GetInstanceShardsArgs Empty => new GetInstanceShardsArgs();
    }

    public sealed class GetInstanceShardsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceShardsInvokeArgs()
        {
        }
        public static new GetInstanceShardsInvokeArgs Empty => new GetInstanceShardsInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceShardsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Instance shard information.
        /// </summary>
        public readonly string InstanceShardsList;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceShardsResult(
            string id,

            string instanceId,

            string instanceShardsList,

            string? resultOutputFile)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceShardsList = instanceShardsList;
            ResultOutputFile = resultOutputFile;
        }
    }
}
