// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    public static class GetCluster
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf cluster
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
        ///         var cluster = Output.Create(Tencentcloud.Tsf.GetCluster.InvokeAsync(new Tencentcloud.Tsf.GetClusterArgs
        ///         {
        ///             ClusterIdLists = 
        ///             {
        ///                 "cluster-vwgj5e6y",
        ///             },
        ///             ClusterType = "V",
        ///             DisableProgramAuthCheck = true,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("tencentcloud:Tsf/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf cluster
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
        ///         var cluster = Output.Create(Tencentcloud.Tsf.GetCluster.InvokeAsync(new Tencentcloud.Tsf.GetClusterArgs
        ///         {
        ///             ClusterIdLists = 
        ///             {
        ///                 "cluster-vwgj5e6y",
        ///             },
        ///             ClusterType = "V",
        ///             DisableProgramAuthCheck = true,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResult>("tencentcloud:Tsf/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        [Input("clusterIdLists")]
        private List<string>? _clusterIdLists;

        /// <summary>
        /// Cluster ID list to be queried, if not filled in or passed, all content will be queried.
        /// </summary>
        public List<string> ClusterIdLists
        {
            get => _clusterIdLists ?? (_clusterIdLists = new List<string>());
            set => _clusterIdLists = value;
        }

        /// <summary>
        /// The type of cluster to be queried, if left blank or not passed, all content will be queried. C: container, V: virtual machine.
        /// </summary>
        [Input("clusterType")]
        public string? ClusterType { get; set; }

        /// <summary>
        /// Whether to disable dataset authentication.
        /// </summary>
        [Input("disableProgramAuthCheck")]
        public bool? DisableProgramAuthCheck { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Filter by keywords for Cluster Id or name.
        /// </summary>
        [Input("searchWord")]
        public string? SearchWord { get; set; }

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clusterIdLists")]
        private InputList<string>? _clusterIdLists;

        /// <summary>
        /// Cluster ID list to be queried, if not filled in or passed, all content will be queried.
        /// </summary>
        public InputList<string> ClusterIdLists
        {
            get => _clusterIdLists ?? (_clusterIdLists = new InputList<string>());
            set => _clusterIdLists = value;
        }

        /// <summary>
        /// The type of cluster to be queried, if left blank or not passed, all content will be queried. C: container, V: virtual machine.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// Whether to disable dataset authentication.
        /// </summary>
        [Input("disableProgramAuthCheck")]
        public Input<bool>? DisableProgramAuthCheck { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Filter by keywords for Cluster Id or name.
        /// </summary>
        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        public GetClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly ImmutableArray<string> ClusterIdLists;
        /// <summary>
        /// Cluster type. Note: This field may return null, indicating no valid value.
        /// </summary>
        public readonly string? ClusterType;
        public readonly bool? DisableProgramAuthCheck;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// TSF cluster pagination object. Note: This field may return null, indicating no valid value.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterResultResult> Results;
        public readonly string? SearchWord;

        [OutputConstructor]
        private GetClusterResult(
            ImmutableArray<string> clusterIdLists,

            string? clusterType,

            bool? disableProgramAuthCheck,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetClusterResultResult> results,

            string? searchWord)
        {
            ClusterIdLists = clusterIdLists;
            ClusterType = clusterType;
            DisableProgramAuthCheck = disableProgramAuthCheck;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            SearchWord = searchWord;
        }
    }
}
