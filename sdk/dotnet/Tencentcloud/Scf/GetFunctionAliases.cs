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
    public static class GetFunctionAliases
    {
        /// <summary>
        /// Use this data source to query detailed information of scf function_aliases
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
        ///         var functionAliases = Output.Create(Tencentcloud.Scf.GetFunctionAliases.InvokeAsync(new Tencentcloud.Scf.GetFunctionAliasesArgs
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
        public static Task<GetFunctionAliasesResult> InvokeAsync(GetFunctionAliasesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionAliasesResult>("tencentcloud:Scf/getFunctionAliases:getFunctionAliases", args ?? new GetFunctionAliasesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of scf function_aliases
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
        ///         var functionAliases = Output.Create(Tencentcloud.Scf.GetFunctionAliases.InvokeAsync(new Tencentcloud.Scf.GetFunctionAliasesArgs
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
        public static Output<GetFunctionAliasesResult> Invoke(GetFunctionAliasesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionAliasesResult>("tencentcloud:Scf/getFunctionAliases:getFunctionAliases", args ?? new GetFunctionAliasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionAliasesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        /// <summary>
        /// If this parameter is provided, only aliases associated with this function version will be returned.
        /// </summary>
        [Input("functionVersion")]
        public string? FunctionVersion { get; set; }

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetFunctionAliasesArgs()
        {
        }
    }

    public sealed class GetFunctionAliasesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// If this parameter is provided, only aliases associated with this function version will be returned.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? FunctionVersion { get; set; }

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetFunctionAliasesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionAliasesResult
    {
        /// <summary>
        /// Alias list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAliasesAliasResult> Aliases;
        public readonly string FunctionName;
        /// <summary>
        /// Master version pointed to by the alias.
        /// </summary>
        public readonly string? FunctionVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetFunctionAliasesResult(
            ImmutableArray<Outputs.GetFunctionAliasesAliasResult> aliases,

            string functionName,

            string? functionVersion,

            string id,

            string? @namespace,

            string? resultOutputFile)
        {
            Aliases = aliases;
            FunctionName = functionName;
            FunctionVersion = functionVersion;
            Id = id;
            Namespace = @namespace;
            ResultOutputFile = resultOutputFile;
        }
    }
}
