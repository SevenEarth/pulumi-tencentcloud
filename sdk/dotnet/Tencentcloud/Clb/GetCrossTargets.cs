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
    public static class GetCrossTargets
    {
        /// <summary>
        /// Use this data source to query detailed information of clb cross_targets
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
        ///         var crossTargets = Output.Create(Tencentcloud.Clb.GetCrossTargets.InvokeAsync(new Tencentcloud.Clb.GetCrossTargetsArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Clb.Inputs.GetCrossTargetsFilterArgs
        ///                 {
        ///                     Name = "vpc-id",
        ///                     Values = 
        ///                     {
        ///                         "vpc-4owdpnwr",
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
        public static Task<GetCrossTargetsResult> InvokeAsync(GetCrossTargetsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCrossTargetsResult>("tencentcloud:Clb/getCrossTargets:getCrossTargets", args ?? new GetCrossTargetsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb cross_targets
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
        ///         var crossTargets = Output.Create(Tencentcloud.Clb.GetCrossTargets.InvokeAsync(new Tencentcloud.Clb.GetCrossTargetsArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Clb.Inputs.GetCrossTargetsFilterArgs
        ///                 {
        ///                     Name = "vpc-id",
        ///                     Values = 
        ///                     {
        ///                         "vpc-4owdpnwr",
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
        public static Output<GetCrossTargetsResult> Invoke(GetCrossTargetsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCrossTargetsResult>("tencentcloud:Clb/getCrossTargets:getCrossTargets", args ?? new GetCrossTargetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCrossTargetsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetCrossTargetsFilterArgs>? _filters;

        /// <summary>
        /// Filter conditions to query CVMs and ENIs: vpc-id - String - Required: No - (Filter condition) Filter by VPC ID, such as vpc-12345678. ip - String - Required: No - (Filter condition) Filter by real server IP, such as 192.168.0.1. listener-id - String - Required: No - (Filter condition) Filter by listener ID, such as lbl-12345678. location-id - String - Required: No - (Filter condition) Filter by forwarding rule ID of the layer-7 listener, such as loc-12345678.
        /// </summary>
        public List<Inputs.GetCrossTargetsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCrossTargetsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetCrossTargetsArgs()
        {
        }
    }

    public sealed class GetCrossTargetsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetCrossTargetsFilterInputArgs>? _filters;

        /// <summary>
        /// Filter conditions to query CVMs and ENIs: vpc-id - String - Required: No - (Filter condition) Filter by VPC ID, such as vpc-12345678. ip - String - Required: No - (Filter condition) Filter by real server IP, such as 192.168.0.1. listener-id - String - Required: No - (Filter condition) Filter by listener ID, such as lbl-12345678. location-id - String - Required: No - (Filter condition) Filter by forwarding rule ID of the layer-7 listener, such as loc-12345678.
        /// </summary>
        public InputList<Inputs.GetCrossTargetsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetCrossTargetsFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetCrossTargetsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCrossTargetsResult
    {
        /// <summary>
        /// Cross target set.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCrossTargetsCrossTargetSetResult> CrossTargetSets;
        public readonly ImmutableArray<Outputs.GetCrossTargetsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetCrossTargetsResult(
            ImmutableArray<Outputs.GetCrossTargetsCrossTargetSetResult> crossTargetSets,

            ImmutableArray<Outputs.GetCrossTargetsFilterResult> filters,

            string id,

            string? resultOutputFile)
        {
            CrossTargetSets = crossTargetSets;
            Filters = filters;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
