// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb
{
    public static class GetClusterResources
    {
        /// <summary>
        /// Use this data source to query detailed information of clb cluster_resources
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
        ///         var clusterResources = Output.Create(Tencentcloud.Clb.GetClusterResources.InvokeAsync(new Tencentcloud.Clb.GetClusterResourcesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Clb.Inputs.GetClusterResourcesFilterArgs
        ///                 {
        ///                     Name = "idle",
        ///                     Values = 
        ///                     {
        ///                         "True",
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
        public static Task<GetClusterResourcesResult> InvokeAsync(GetClusterResourcesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResourcesResult>("tencentcloud:Clb/getClusterResources:getClusterResources", args ?? new GetClusterResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb cluster_resources
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
        ///         var clusterResources = Output.Create(Tencentcloud.Clb.GetClusterResources.InvokeAsync(new Tencentcloud.Clb.GetClusterResourcesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Clb.Inputs.GetClusterResourcesFilterArgs
        ///                 {
        ///                     Name = "idle",
        ///                     Values = 
        ///                     {
        ///                         "True",
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
        public static Output<GetClusterResourcesResult> Invoke(GetClusterResourcesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResourcesResult>("tencentcloud:Clb/getClusterResources:getClusterResources", args ?? new GetClusterResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterResourcesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetClusterResourcesFilterArgs>? _filters;

        /// <summary>
        /// Filter conditions to query cluster. cluster-id - String - Required: No - (Filter condition) Filter by cluster ID, such as tgw-12345678. vip - String - Required: No - (Filter condition) Filter by loadbalancer vip, such as 192.168.0.1. loadblancer-id - String - Required: No - (Filter condition) Filter by loadblancer ID, such as lbl-12345678. idle - String - Required: No - (Filter condition) Filter by Whether load balancing is idle, such as True, False.
        /// </summary>
        public List<Inputs.GetClusterResourcesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetClusterResourcesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetClusterResourcesArgs()
        {
        }
    }

    public sealed class GetClusterResourcesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetClusterResourcesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter conditions to query cluster. cluster-id - String - Required: No - (Filter condition) Filter by cluster ID, such as tgw-12345678. vip - String - Required: No - (Filter condition) Filter by loadbalancer vip, such as 192.168.0.1. loadblancer-id - String - Required: No - (Filter condition) Filter by loadblancer ID, such as lbl-12345678. idle - String - Required: No - (Filter condition) Filter by Whether load balancing is idle, such as True, False.
        /// </summary>
        public InputList<Inputs.GetClusterResourcesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetClusterResourcesFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetClusterResourcesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResourcesResult
    {
        /// <summary>
        /// Cluster resource set.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterResourcesClusterResourceSetResult> ClusterResourceSets;
        public readonly ImmutableArray<Outputs.GetClusterResourcesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetClusterResourcesResult(
            ImmutableArray<Outputs.GetClusterResourcesClusterResourceSetResult> clusterResourceSets,

            ImmutableArray<Outputs.GetClusterResourcesFilterResult> filters,

            string id,

            string? resultOutputFile)
        {
            ClusterResourceSets = clusterResourceSets;
            Filters = filters;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
