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
    public static class GetUsableUnitNamespaces
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf usable_unit_namespaces
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
        ///         var usableUnitNamespaces = Output.Create(Tencentcloud.Tsf.GetUsableUnitNamespaces.InvokeAsync(new Tencentcloud.Tsf.GetUsableUnitNamespacesArgs
        ///         {
        ///             SearchWord = "",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUsableUnitNamespacesResult> InvokeAsync(GetUsableUnitNamespacesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUsableUnitNamespacesResult>("tencentcloud:Tsf/getUsableUnitNamespaces:getUsableUnitNamespaces", args ?? new GetUsableUnitNamespacesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf usable_unit_namespaces
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
        ///         var usableUnitNamespaces = Output.Create(Tencentcloud.Tsf.GetUsableUnitNamespaces.InvokeAsync(new Tencentcloud.Tsf.GetUsableUnitNamespacesArgs
        ///         {
        ///             SearchWord = "",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUsableUnitNamespacesResult> Invoke(GetUsableUnitNamespacesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUsableUnitNamespacesResult>("tencentcloud:Tsf/getUsableUnitNamespaces:getUsableUnitNamespaces", args ?? new GetUsableUnitNamespacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsableUnitNamespacesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// search by namespace id or namespace Name.
        /// </summary>
        [Input("searchWord")]
        public string? SearchWord { get; set; }

        public GetUsableUnitNamespacesArgs()
        {
        }
    }

    public sealed class GetUsableUnitNamespacesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// search by namespace id or namespace Name.
        /// </summary>
        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        public GetUsableUnitNamespacesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUsableUnitNamespacesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// namespace object list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsableUnitNamespacesResultResult> Results;
        public readonly string? SearchWord;

        [OutputConstructor]
        private GetUsableUnitNamespacesResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetUsableUnitNamespacesResultResult> results,

            string? searchWord)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            SearchWord = searchWord;
        }
    }
}
