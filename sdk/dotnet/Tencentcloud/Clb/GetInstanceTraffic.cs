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
    public static class GetInstanceTraffic
    {
        /// <summary>
        /// Use this data source to query detailed information of clb instance_traffic
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
        ///         var instanceTraffic = Output.Create(Tencentcloud.Clb.GetInstanceTraffic.InvokeAsync(new Tencentcloud.Clb.GetInstanceTrafficArgs
        ///         {
        ///             LoadBalancerRegion = "ap-guangzhou",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceTrafficResult> InvokeAsync(GetInstanceTrafficArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTrafficResult>("tencentcloud:Clb/getInstanceTraffic:getInstanceTraffic", args ?? new GetInstanceTrafficArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb instance_traffic
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
        ///         var instanceTraffic = Output.Create(Tencentcloud.Clb.GetInstanceTraffic.InvokeAsync(new Tencentcloud.Clb.GetInstanceTrafficArgs
        ///         {
        ///             LoadBalancerRegion = "ap-guangzhou",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceTrafficResult> Invoke(GetInstanceTrafficInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceTrafficResult>("tencentcloud:Clb/getInstanceTraffic:getInstanceTraffic", args ?? new GetInstanceTrafficInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceTrafficArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
        /// </summary>
        [Input("loadBalancerRegion")]
        public string? LoadBalancerRegion { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceTrafficArgs()
        {
        }
    }

    public sealed class GetInstanceTrafficInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
        /// </summary>
        [Input("loadBalancerRegion")]
        public Input<string>? LoadBalancerRegion { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceTrafficInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceTrafficResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? LoadBalancerRegion;
        /// <summary>
        /// Information of CLB instances sorted by outbound bandwidth from highest to lowest. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTrafficLoadBalancerTrafficResult> LoadBalancerTraffics;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceTrafficResult(
            string id,

            string? loadBalancerRegion,

            ImmutableArray<Outputs.GetInstanceTrafficLoadBalancerTrafficResult> loadBalancerTraffics,

            string? resultOutputFile)
        {
            Id = id;
            LoadBalancerRegion = loadBalancerRegion;
            LoadBalancerTraffics = loadBalancerTraffics;
            ResultOutputFile = resultOutputFile;
        }
    }
}