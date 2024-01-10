// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes
{
    public static class GetClusterNodePools
    {
        /// <summary>
        /// Use this data source to query detailed information of kubernetes cluster_node_pools
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var clusterNodePools = Output.Create(Tencentcloud.Kubernetes.GetClusterNodePools.InvokeAsync(new Tencentcloud.Kubernetes.GetClusterNodePoolsArgs
        ///         {
        ///             ClusterId = "cls-kzilgv5m",
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Kubernetes.Inputs.GetClusterNodePoolsFilterArgs
        ///                 {
        ///                     Name = "NodePoolsName",
        ///                     Values = 
        ///                     {
        ///                         "mynodepool_xxxx",
        ///                     },
        ///                 },
        ///                 new Tencentcloud.Kubernetes.Inputs.GetClusterNodePoolsFilterArgs
        ///                 {
        ///                     Name = "NodePoolsId",
        ///                     Values = 
        ///                     {
        ///                         "np-ngjwhdv4",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterNodePoolsResult> InvokeAsync(GetClusterNodePoolsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterNodePoolsResult>("tencentcloud:Kubernetes/getClusterNodePools:getClusterNodePools", args ?? new GetClusterNodePoolsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of kubernetes cluster_node_pools
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var clusterNodePools = Output.Create(Tencentcloud.Kubernetes.GetClusterNodePools.InvokeAsync(new Tencentcloud.Kubernetes.GetClusterNodePoolsArgs
        ///         {
        ///             ClusterId = "cls-kzilgv5m",
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Kubernetes.Inputs.GetClusterNodePoolsFilterArgs
        ///                 {
        ///                     Name = "NodePoolsName",
        ///                     Values = 
        ///                     {
        ///                         "mynodepool_xxxx",
        ///                     },
        ///                 },
        ///                 new Tencentcloud.Kubernetes.Inputs.GetClusterNodePoolsFilterArgs
        ///                 {
        ///                     Name = "NodePoolsId",
        ///                     Values = 
        ///                     {
        ///                         "np-ngjwhdv4",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterNodePoolsResult> Invoke(GetClusterNodePoolsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterNodePoolsResult>("tencentcloud:Kubernetes/getClusterNodePools:getClusterNodePools", args ?? new GetClusterNodePoolsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterNodePoolsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetClusterNodePoolsFilterArgs>? _filters;

        /// <summary>
        /// NodePoolsName, Filter according to the node pool name, type: String, required: no. NodePoolsId, Filter according to the node pool ID, type: String, required: no. tags, Filter according to the label key value pairs, type: String, required: no. tag:tag-key, Filter according to the label key value pairs, type: String, required: no.
        /// </summary>
        public List<Inputs.GetClusterNodePoolsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetClusterNodePoolsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetClusterNodePoolsArgs()
        {
        }
    }

    public sealed class GetClusterNodePoolsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetClusterNodePoolsFilterInputArgs>? _filters;

        /// <summary>
        /// NodePoolsName, Filter according to the node pool name, type: String, required: no. NodePoolsId, Filter according to the node pool ID, type: String, required: no. tags, Filter according to the label key value pairs, type: String, required: no. tag:tag-key, Filter according to the label key value pairs, type: String, required: no.
        /// </summary>
        public InputList<Inputs.GetClusterNodePoolsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetClusterNodePoolsFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetClusterNodePoolsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterNodePoolsResult
    {
        public readonly string ClusterId;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Node Pool List.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterNodePoolsNodePoolSetResult> NodePoolSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetClusterNodePoolsResult(
            string clusterId,

            ImmutableArray<Outputs.GetClusterNodePoolsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetClusterNodePoolsNodePoolSetResult> nodePoolSets,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            Filters = filters;
            Id = id;
            NodePoolSets = nodePoolSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
