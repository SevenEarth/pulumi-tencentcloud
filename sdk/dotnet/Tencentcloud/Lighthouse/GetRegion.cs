// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse
{
    public static class GetRegion
    {
        /// <summary>
        /// Use this data source to query detailed information of lighthouse region
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
        ///         var region = Output.Create(Tencentcloud.Lighthouse.GetRegion.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegionResult> InvokeAsync(GetRegionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionResult>("tencentcloud:Lighthouse/getRegion:getRegion", args ?? new GetRegionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of lighthouse region
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
        ///         var region = Output.Create(Tencentcloud.Lighthouse.GetRegion.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegionResult> Invoke(GetRegionInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionResult>("tencentcloud:Lighthouse/getRegion:getRegion", args ?? new GetRegionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRegionArgs()
        {
        }
    }

    public sealed class GetRegionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRegionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Region information list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegionRegionSetResult> RegionSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRegionResult(
            string id,

            ImmutableArray<Outputs.GetRegionRegionSetResult> regionSets,

            string? resultOutputFile)
        {
            Id = id;
            RegionSets = regionSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
