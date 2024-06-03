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
    public static class GetFunctionAddress
    {
        /// <summary>
        /// Use this data source to query detailed information of scf function_address
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var functionAddress = Tencentcloud.Scf.GetFunctionAddress.Invoke(new()
        ///     {
        ///         FunctionName = "keep-1676351130",
        ///         Namespace = "default",
        ///         Qualifier = "$LATEST",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetFunctionAddressResult> InvokeAsync(GetFunctionAddressArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionAddressResult>("tencentcloud:Scf/getFunctionAddress:getFunctionAddress", args ?? new GetFunctionAddressArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of scf function_address
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var functionAddress = Tencentcloud.Scf.GetFunctionAddress.Invoke(new()
        ///     {
        ///         FunctionName = "keep-1676351130",
        ///         Namespace = "default",
        ///         Qualifier = "$LATEST",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetFunctionAddressResult> Invoke(GetFunctionAddressInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionAddressResult>("tencentcloud:Scf/getFunctionAddress:getFunctionAddress", args ?? new GetFunctionAddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionAddressArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// Function version.
        /// </summary>
        [Input("qualifier")]
        public string? Qualifier { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetFunctionAddressArgs()
        {
        }
        public static new GetFunctionAddressArgs Empty => new GetFunctionAddressArgs();
    }

    public sealed class GetFunctionAddressInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Function version.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetFunctionAddressInvokeArgs()
        {
        }
        public static new GetFunctionAddressInvokeArgs Empty => new GetFunctionAddressInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionAddressResult
    {
        /// <summary>
        /// SHA256 code of the function.
        /// </summary>
        public readonly string CodeSha256;
        public readonly string FunctionName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Namespace;
        public readonly string? Qualifier;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Cos address of the function.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetFunctionAddressResult(
            string codeSha256,

            string functionName,

            string id,

            string? @namespace,

            string? qualifier,

            string? resultOutputFile,

            string url)
        {
            CodeSha256 = codeSha256;
            FunctionName = functionName;
            Id = id;
            Namespace = @namespace;
            Qualifier = qualifier;
            ResultOutputFile = resultOutputFile;
            Url = url;
        }
    }
}
