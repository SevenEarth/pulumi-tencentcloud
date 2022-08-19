// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes
{
    public static class GetCharts
    {
        /// <summary>
        /// Use this data source to query detailed information of kubernetes cluster addons.
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
        ///         var name = Output.Create(Tencentcloud.Kubernetes.GetCharts.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetChartsResult> InvokeAsync(GetChartsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetChartsResult>("tencentcloud:Kubernetes/getCharts:getCharts", args ?? new GetChartsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of kubernetes cluster addons.
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
        ///         var name = Output.Create(Tencentcloud.Kubernetes.GetCharts.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetChartsResult> Invoke(GetChartsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetChartsResult>("tencentcloud:Kubernetes/getCharts:getCharts", args ?? new GetChartsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChartsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Operation system app supported. Available values: `arm32`, `arm64`, `amd64`.
        /// </summary>
        [Input("arch")]
        public string? Arch { get; set; }

        /// <summary>
        /// Cluster type. Available values: `tke`, `eks`.
        /// </summary>
        [Input("clusterType")]
        public string? ClusterType { get; set; }

        /// <summary>
        /// Kind of app chart. Available values: `log`, `scheduler`, `network`, `storage`, `monitor`, `dns`, `image`, `other`, `invisible`.
        /// </summary>
        [Input("kind")]
        public string? Kind { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetChartsArgs()
        {
        }
    }

    public sealed class GetChartsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Operation system app supported. Available values: `arm32`, `arm64`, `amd64`.
        /// </summary>
        [Input("arch")]
        public Input<string>? Arch { get; set; }

        /// <summary>
        /// Cluster type. Available values: `tke`, `eks`.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// Kind of app chart. Available values: `log`, `scheduler`, `network`, `storage`, `monitor`, `dns`, `image`, `other`, `invisible`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetChartsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetChartsResult
    {
        public readonly string? Arch;
        /// <summary>
        /// App chart list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetChartsChartListResult> ChartLists;
        public readonly string? ClusterType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Kind;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetChartsResult(
            string? arch,

            ImmutableArray<Outputs.GetChartsChartListResult> chartLists,

            string? clusterType,

            string id,

            string? kind,

            string? resultOutputFile)
        {
            Arch = arch;
            ChartLists = chartLists;
            ClusterType = clusterType;
            Id = id;
            Kind = kind;
            ResultOutputFile = resultOutputFile;
        }
    }
}
