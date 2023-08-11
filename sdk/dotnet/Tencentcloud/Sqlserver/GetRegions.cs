// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    public static class GetRegions
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_regions
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
        ///         var example = Output.Create(Tencentcloud.Sqlserver.GetRegions.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegionsResult> InvokeAsync(GetRegionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionsResult>("tencentcloud:Sqlserver/getRegions:getRegions", args ?? new GetRegionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_regions
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
        ///         var example = Output.Create(Tencentcloud.Sqlserver.GetRegions.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegionsResult> Invoke(GetRegionsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionsResult>("tencentcloud:Sqlserver/getRegions:getRegions", args ?? new GetRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRegionsArgs()
        {
        }
    }

    public sealed class GetRegionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRegionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Region information array.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegionsRegionSetResult> RegionSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRegionsResult(
            string id,

            ImmutableArray<Outputs.GetRegionsRegionSetResult> regionSets,

            string? resultOutputFile)
        {
            Id = id;
            RegionSets = regionSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
