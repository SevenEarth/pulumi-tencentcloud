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
    public static class GetApiDetail
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf api_detail
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
        ///         var apiDetail = Output.Create(Tencentcloud.Tsf.GetApiDetail.InvokeAsync(new Tencentcloud.Tsf.GetApiDetailArgs
        ///         {
        ///             ApplicationId = "application-a24x29xv",
        ///             Method = "GET",
        ///             MicroserviceId = "ms-yq3jo6jd",
        ///             Path = "/printRequest",
        ///             PkgVersion = "20210625192923",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApiDetailResult> InvokeAsync(GetApiDetailArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiDetailResult>("tencentcloud:Tsf/getApiDetail:getApiDetail", args ?? new GetApiDetailArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf api_detail
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
        ///         var apiDetail = Output.Create(Tencentcloud.Tsf.GetApiDetail.InvokeAsync(new Tencentcloud.Tsf.GetApiDetailArgs
        ///         {
        ///             ApplicationId = "application-a24x29xv",
        ///             Method = "GET",
        ///             MicroserviceId = "ms-yq3jo6jd",
        ///             Path = "/printRequest",
        ///             PkgVersion = "20210625192923",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApiDetailResult> Invoke(GetApiDetailInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApiDetailResult>("tencentcloud:Tsf/getApiDetail:getApiDetail", args ?? new GetApiDetailInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiDetailArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// application id.
        /// </summary>
        [Input("applicationId", required: true)]
        public string ApplicationId { get; set; } = null!;

        /// <summary>
        /// request method.
        /// </summary>
        [Input("method", required: true)]
        public string Method { get; set; } = null!;

        /// <summary>
        /// microservice id.
        /// </summary>
        [Input("microserviceId", required: true)]
        public string MicroserviceId { get; set; } = null!;

        /// <summary>
        /// api path.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// pkg version.
        /// </summary>
        [Input("pkgVersion", required: true)]
        public string PkgVersion { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetApiDetailArgs()
        {
        }
    }

    public sealed class GetApiDetailInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// application id.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// request method.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// microservice id.
        /// </summary>
        [Input("microserviceId", required: true)]
        public Input<string> MicroserviceId { get; set; } = null!;

        /// <summary>
        /// api path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// pkg version.
        /// </summary>
        [Input("pkgVersion", required: true)]
        public Input<string> PkgVersion { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetApiDetailInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiDetailResult
    {
        public readonly string ApplicationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Method;
        public readonly string MicroserviceId;
        public readonly string Path;
        public readonly string PkgVersion;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// api detail.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApiDetailResultResult> Results;

        [OutputConstructor]
        private GetApiDetailResult(
            string applicationId,

            string id,

            string method,

            string microserviceId,

            string path,

            string pkgVersion,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetApiDetailResultResult> results)
        {
            ApplicationId = applicationId;
            Id = id;
            Method = method;
            MicroserviceId = microserviceId;
            Path = path;
            PkgVersion = pkgVersion;
            ResultOutputFile = resultOutputFile;
            Results = results;
        }
    }
}
