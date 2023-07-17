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
    public static class GetResources
    {
        /// <summary>
        /// Use this data source to query detailed information of clb resources
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
        ///         var resources = Output.Create(Tencentcloud.Clb.GetResources.InvokeAsync(new Tencentcloud.Clb.GetResourcesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Clb.Inputs.GetResourcesFilterArgs
        ///                 {
        ///                     Name = "isp",
        ///                     Values = 
        ///                     {
        ///                         "BGP",
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
        public static Task<GetResourcesResult> InvokeAsync(GetResourcesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcesResult>("tencentcloud:Clb/getResources:getResources", args ?? new GetResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb resources
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
        ///         var resources = Output.Create(Tencentcloud.Clb.GetResources.InvokeAsync(new Tencentcloud.Clb.GetResourcesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Clb.Inputs.GetResourcesFilterArgs
        ///                 {
        ///                     Name = "isp",
        ///                     Values = 
        ///                     {
        ///                         "BGP",
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
        public static Output<GetResourcesResult> Invoke(GetResourcesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourcesResult>("tencentcloud:Clb/getResources:getResources", args ?? new GetResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetResourcesFilterArgs>? _filters;

        /// <summary>
        /// Filter to query the list of AZ resources as detailed below: zone - String - Optional - Filter by AZ, such as ap-guangzhou-1. isp -- String - Optional - Filter by the ISP. Values: BGP, CMCC, CUCC and CTCC.
        /// </summary>
        public List<Inputs.GetResourcesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetResourcesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetResourcesArgs()
        {
        }
    }

    public sealed class GetResourcesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetResourcesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter to query the list of AZ resources as detailed below: zone - String - Optional - Filter by AZ, such as ap-guangzhou-1. isp -- String - Optional - Filter by the ISP. Values: BGP, CMCC, CUCC and CTCC.
        /// </summary>
        public InputList<Inputs.GetResourcesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetResourcesFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetResourcesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcesResult
    {
        public readonly ImmutableArray<Outputs.GetResourcesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// List of resources supported by the AZ.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourcesZoneResourceSetResult> ZoneResourceSets;

        [OutputConstructor]
        private GetResourcesResult(
            ImmutableArray<Outputs.GetResourcesFilterResult> filters,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetResourcesZoneResourceSetResult> zoneResourceSets)
        {
            Filters = filters;
            Id = id;
            ResultOutputFile = resultOutputFile;
            ZoneResourceSets = zoneResourceSets;
        }
    }
}
