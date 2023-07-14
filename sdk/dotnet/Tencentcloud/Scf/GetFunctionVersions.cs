// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf
{
    public static class GetFunctionVersions
    {
        /// <summary>
        /// Use this data source to query detailed information of scf function_versions
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
        ///         var functionVersions = Output.Create(Tencentcloud.Scf.GetFunctionVersions.InvokeAsync(new Tencentcloud.Scf.GetFunctionVersionsArgs
        ///         {
        ///             FunctionName = "keep-1676351130",
        ///             Namespace = "default",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionVersionsResult> InvokeAsync(GetFunctionVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionVersionsResult>("tencentcloud:Scf/getFunctionVersions:getFunctionVersions", args ?? new GetFunctionVersionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of scf function_versions
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
        ///         var functionVersions = Output.Create(Tencentcloud.Scf.GetFunctionVersions.InvokeAsync(new Tencentcloud.Scf.GetFunctionVersionsArgs
        ///         {
        ///             FunctionName = "keep-1676351130",
        ///             Namespace = "default",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFunctionVersionsResult> Invoke(GetFunctionVersionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionVersionsResult>("tencentcloud:Scf/getFunctionVersions:getFunctionVersions", args ?? new GetFunctionVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function Name.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        /// <summary>
        /// The namespace where the function locates.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// It specifies whether to return the results in ascending or descending order. The value is `ASC` or `DESC`.
        /// </summary>
        [Input("order")]
        public string? Order { get; set; }

        /// <summary>
        /// It specifies the sorting order of the results according to a specified field, such as `AddTime`, `ModTime`.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetFunctionVersionsArgs()
        {
        }
    }

    public sealed class GetFunctionVersionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function Name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// The namespace where the function locates.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// It specifies whether to return the results in ascending or descending order. The value is `ASC` or `DESC`.
        /// </summary>
        [Input("order")]
        public Input<string>? Order { get; set; }

        /// <summary>
        /// It specifies the sorting order of the results according to a specified field, such as `AddTime`, `ModTime`.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetFunctionVersionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionVersionsResult
    {
        public readonly string FunctionName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        public readonly string? Order;
        public readonly string? OrderBy;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Function version listNote: This field may return null, indicating that no valid values is found.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionVersionsVersionResult> Versions;

        [OutputConstructor]
        private GetFunctionVersionsResult(
            string functionName,

            string id,

            string? @namespace,

            string? order,

            string? orderBy,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetFunctionVersionsVersionResult> versions)
        {
            FunctionName = functionName;
            Id = id;
            Namespace = @namespace;
            Order = order;
            OrderBy = orderBy;
            ResultOutputFile = resultOutputFile;
            Versions = versions;
        }
    }
}
