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
    public static class GetIdleInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of clb idle_loadbalancers
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
        ///         var idleInstance = Output.Create(Tencentcloud.Clb.GetIdleInstances.InvokeAsync(new Tencentcloud.Clb.GetIdleInstancesArgs
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
        public static Task<GetIdleInstancesResult> InvokeAsync(GetIdleInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdleInstancesResult>("tencentcloud:Clb/getIdleInstances:getIdleInstances", args ?? new GetIdleInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb idle_loadbalancers
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
        ///         var idleInstance = Output.Create(Tencentcloud.Clb.GetIdleInstances.InvokeAsync(new Tencentcloud.Clb.GetIdleInstancesArgs
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
        public static Output<GetIdleInstancesResult> Invoke(GetIdleInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIdleInstancesResult>("tencentcloud:Clb/getIdleInstances:getIdleInstances", args ?? new GetIdleInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdleInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// CLB instance region.
        /// </summary>
        [Input("loadBalancerRegion")]
        public string? LoadBalancerRegion { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetIdleInstancesArgs()
        {
        }
    }

    public sealed class GetIdleInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// CLB instance region.
        /// </summary>
        [Input("loadBalancerRegion")]
        public Input<string>? LoadBalancerRegion { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetIdleInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIdleInstancesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of idle CLBs. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIdleInstancesIdleLoadBalancerResult> IdleLoadBalancers;
        public readonly string? LoadBalancerRegion;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetIdleInstancesResult(
            string id,

            ImmutableArray<Outputs.GetIdleInstancesIdleLoadBalancerResult> idleLoadBalancers,

            string? loadBalancerRegion,

            string? resultOutputFile)
        {
            Id = id;
            IdleLoadBalancers = idleLoadBalancers;
            LoadBalancerRegion = loadBalancerRegion;
            ResultOutputFile = resultOutputFile;
        }
    }
}
