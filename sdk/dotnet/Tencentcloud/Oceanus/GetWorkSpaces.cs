// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus
{
    public static class GetWorkSpaces
    {
        /// <summary>
        /// Use this data source to query detailed information of oceanus work_spaces
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
        ///         var example = Output.Create(Tencentcloud.Oceanus.GetWorkSpaces.InvokeAsync(new Tencentcloud.Oceanus.GetWorkSpacesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Oceanus.Inputs.GetWorkSpacesFilterArgs
        ///                 {
        ///                     Name = "WorkSpaceName",
        ///                     Values = 
        ///                     {
        ///                         "tf_example",
        ///                     },
        ///                 },
        ///             },
        ///             OrderType = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWorkSpacesResult> InvokeAsync(GetWorkSpacesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkSpacesResult>("tencentcloud:Oceanus/getWorkSpaces:getWorkSpaces", args ?? new GetWorkSpacesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of oceanus work_spaces
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
        ///         var example = Output.Create(Tencentcloud.Oceanus.GetWorkSpaces.InvokeAsync(new Tencentcloud.Oceanus.GetWorkSpacesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Oceanus.Inputs.GetWorkSpacesFilterArgs
        ///                 {
        ///                     Name = "WorkSpaceName",
        ///                     Values = 
        ///                     {
        ///                         "tf_example",
        ///                     },
        ///                 },
        ///             },
        ///             OrderType = 1,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWorkSpacesResult> Invoke(GetWorkSpacesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWorkSpacesResult>("tencentcloud:Oceanus/getWorkSpaces:getWorkSpaces", args ?? new GetWorkSpacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkSpacesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetWorkSpacesFilterArgs>? _filters;

        /// <summary>
        /// Filter rules.
        /// </summary>
        public List<Inputs.GetWorkSpacesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetWorkSpacesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// 1:sort by creation time in descending order (default); 2:sort by creation time in ascending order; 3:sort by status in descending order; 4:sort by status in ascending order; default is 0.
        /// </summary>
        [Input("orderType")]
        public int? OrderType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetWorkSpacesArgs()
        {
        }
    }

    public sealed class GetWorkSpacesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetWorkSpacesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter rules.
        /// </summary>
        public InputList<Inputs.GetWorkSpacesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetWorkSpacesFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// 1:sort by creation time in descending order (default); 2:sort by creation time in ascending order; 3:sort by status in descending order; 4:sort by status in ascending order; default is 0.
        /// </summary>
        [Input("orderType")]
        public Input<int>? OrderType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetWorkSpacesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkSpacesResult
    {
        public readonly ImmutableArray<Outputs.GetWorkSpacesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? OrderType;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// List of workspace detailsNote: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWorkSpacesWorkSpaceSetItemResult> WorkSpaceSetItems;

        [OutputConstructor]
        private GetWorkSpacesResult(
            ImmutableArray<Outputs.GetWorkSpacesFilterResult> filters,

            string id,

            int? orderType,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetWorkSpacesWorkSpaceSetItemResult> workSpaceSetItems)
        {
            Filters = filters;
            Id = id;
            OrderType = orderType;
            ResultOutputFile = resultOutputFile;
            WorkSpaceSetItems = workSpaceSetItems;
        }
    }
}
