// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf
{
    public static class GetTlsVersions
    {
        /// <summary>
        /// Use this data source to query detailed information of waf tls_versions
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
        ///         var example = Output.Create(Tencentcloud.Waf.GetTlsVersions.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTlsVersionsResult> InvokeAsync(GetTlsVersionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTlsVersionsResult>("tencentcloud:Waf/getTlsVersions:getTlsVersions", args ?? new GetTlsVersionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of waf tls_versions
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
        ///         var example = Output.Create(Tencentcloud.Waf.GetTlsVersions.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTlsVersionsResult> Invoke(GetTlsVersionsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTlsVersionsResult>("tencentcloud:Waf/getTlsVersions:getTlsVersions", args ?? new GetTlsVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTlsVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetTlsVersionsArgs()
        {
        }
    }

    public sealed class GetTlsVersionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetTlsVersionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTlsVersionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// TLS key value.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTlsVersionsTlResult> Tls;

        [OutputConstructor]
        private GetTlsVersionsResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetTlsVersionsTlResult> tls)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            Tls = tls;
        }
    }
}