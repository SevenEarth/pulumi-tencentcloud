// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq
{
    public static class GetRocketmqCluster
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq cluster
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var exampleRocketmqCluster = Output.Create(Tencentcloud.Tdmq.GetRocketmqCluster.InvokeAsync(new Tencentcloud.Tdmq.GetRocketmqClusterArgs
        ///         {
        ///             NameKeyword = exampleTdmq / rocketmqClusterRocketmqCluster.ClusterName,
        ///         }));
        ///         var exampleTdmq_rocketmqClusterRocketmqCluster = new Tencentcloud.Tdmq.RocketmqCluster("exampleTdmq/rocketmqClusterRocketmqCluster", new Tencentcloud.Tdmq.RocketmqClusterArgs
        ///         {
        ///             ClusterName = "tf_example",
        ///             Remark = "remark.",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRocketmqClusterResult> InvokeAsync(GetRocketmqClusterArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRocketmqClusterResult>("tencentcloud:Tdmq/getRocketmqCluster:getRocketmqCluster", args ?? new GetRocketmqClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq cluster
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var exampleRocketmqCluster = Output.Create(Tencentcloud.Tdmq.GetRocketmqCluster.InvokeAsync(new Tencentcloud.Tdmq.GetRocketmqClusterArgs
        ///         {
        ///             NameKeyword = exampleTdmq / rocketmqClusterRocketmqCluster.ClusterName,
        ///         }));
        ///         var exampleTdmq_rocketmqClusterRocketmqCluster = new Tencentcloud.Tdmq.RocketmqCluster("exampleTdmq/rocketmqClusterRocketmqCluster", new Tencentcloud.Tdmq.RocketmqClusterArgs
        ///         {
        ///             ClusterName = "tf_example",
        ///             Remark = "remark.",
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRocketmqClusterResult> Invoke(GetRocketmqClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRocketmqClusterResult>("tencentcloud:Tdmq/getRocketmqCluster:getRocketmqCluster", args ?? new GetRocketmqClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRocketmqClusterArgs : Pulumi.InvokeArgs
    {
        [Input("clusterIdLists")]
        private List<string>? _clusterIdLists;

        /// <summary>
        /// Filter by cluster ID.
        /// </summary>
        public List<string> ClusterIdLists
        {
            get => _clusterIdLists ?? (_clusterIdLists = new List<string>());
            set => _clusterIdLists = value;
        }

        /// <summary>
        /// Search by cluster ID.
        /// </summary>
        [Input("idKeyword")]
        public string? IdKeyword { get; set; }

        /// <summary>
        /// Search by cluster name.
        /// </summary>
        [Input("nameKeyword")]
        public string? NameKeyword { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRocketmqClusterArgs()
        {
        }
    }

    public sealed class GetRocketmqClusterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clusterIdLists")]
        private InputList<string>? _clusterIdLists;

        /// <summary>
        /// Filter by cluster ID.
        /// </summary>
        public InputList<string> ClusterIdLists
        {
            get => _clusterIdLists ?? (_clusterIdLists = new InputList<string>());
            set => _clusterIdLists = value;
        }

        /// <summary>
        /// Search by cluster ID.
        /// </summary>
        [Input("idKeyword")]
        public Input<string>? IdKeyword { get; set; }

        /// <summary>
        /// Search by cluster name.
        /// </summary>
        [Input("nameKeyword")]
        public Input<string>? NameKeyword { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRocketmqClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRocketmqClusterResult
    {
        public readonly ImmutableArray<string> ClusterIdLists;
        /// <summary>
        /// Cluster information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqClusterClusterListResult> ClusterLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IdKeyword;
        public readonly string? NameKeyword;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRocketmqClusterResult(
            ImmutableArray<string> clusterIdLists,

            ImmutableArray<Outputs.GetRocketmqClusterClusterListResult> clusterLists,

            string id,

            string? idKeyword,

            string? nameKeyword,

            string? resultOutputFile)
        {
            ClusterIdLists = clusterIdLists;
            ClusterLists = clusterLists;
            Id = id;
            IdKeyword = idKeyword;
            NameKeyword = nameKeyword;
            ResultOutputFile = resultOutputFile;
        }
    }
}
