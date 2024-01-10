// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl
{
    public static class GetDescribeCertificate
    {
        /// <summary>
        /// Use this data source to query detailed information of ssl describe_certificate
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
        ///         var describeCertificate = Output.Create(Tencentcloud.Ssl.GetDescribeCertificate.InvokeAsync(new Tencentcloud.Ssl.GetDescribeCertificateArgs
        ///         {
        ///             CertificateId = "8cj4g8h8",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDescribeCertificateResult> InvokeAsync(GetDescribeCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDescribeCertificateResult>("tencentcloud:Ssl/getDescribeCertificate:getDescribeCertificate", args ?? new GetDescribeCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssl describe_certificate
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
        ///         var describeCertificate = Output.Create(Tencentcloud.Ssl.GetDescribeCertificate.InvokeAsync(new Tencentcloud.Ssl.GetDescribeCertificateArgs
        ///         {
        ///             CertificateId = "8cj4g8h8",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDescribeCertificateResult> Invoke(GetDescribeCertificateInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDescribeCertificateResult>("tencentcloud:Ssl/getDescribeCertificate:getDescribeCertificate", args ?? new GetDescribeCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeCertificateArgs()
        {
        }
    }

    public sealed class GetDescribeCertificateInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeCertificateInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDescribeCertificateResult
    {
        public readonly string CertificateId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// result list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeCertificateResultResult> Results;

        [OutputConstructor]
        private GetDescribeCertificateResult(
            string certificateId,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetDescribeCertificateResultResult> results)
        {
            CertificateId = certificateId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
        }
    }
}
