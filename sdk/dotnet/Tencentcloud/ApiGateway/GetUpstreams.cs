// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    public static class GetUpstreams
    {
        /// <summary>
        /// Use this data source to query detailed information of apigateway upstream
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
        ///         var example = Output.Create(Tencentcloud.ApiGateway.GetUpstreams.InvokeAsync(new Tencentcloud.ApiGateway.GetUpstreamsArgs
        ///         {
        ///             UpstreamId = "upstream-4n5bfklc",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUpstreamsResult> InvokeAsync(GetUpstreamsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUpstreamsResult>("tencentcloud:ApiGateway/getUpstreams:getUpstreams", args ?? new GetUpstreamsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of apigateway upstream
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
        ///         var example = Output.Create(Tencentcloud.ApiGateway.GetUpstreams.InvokeAsync(new Tencentcloud.ApiGateway.GetUpstreamsArgs
        ///         {
        ///             UpstreamId = "upstream-4n5bfklc",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUpstreamsResult> Invoke(GetUpstreamsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUpstreamsResult>("tencentcloud:ApiGateway/getUpstreams:getUpstreams", args ?? new GetUpstreamsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUpstreamsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetUpstreamsFilterArgs>? _filters;

        /// <summary>
        /// ServiceId and ApiId filtering queries.
        /// </summary>
        public List<Inputs.GetUpstreamsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetUpstreamsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Backend channel ID.
        /// </summary>
        [Input("upstreamId", required: true)]
        public string UpstreamId { get; set; } = null!;

        public GetUpstreamsArgs()
        {
        }
    }

    public sealed class GetUpstreamsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetUpstreamsFilterInputArgs>? _filters;

        /// <summary>
        /// ServiceId and ApiId filtering queries.
        /// </summary>
        public InputList<Inputs.GetUpstreamsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetUpstreamsFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Backend channel ID.
        /// </summary>
        [Input("upstreamId", required: true)]
        public Input<string> UpstreamId { get; set; } = null!;

        public GetUpstreamsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUpstreamsResult
    {
        public readonly ImmutableArray<Outputs.GetUpstreamsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Query Results.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsResultResult> Results;
        public readonly string UpstreamId;

        [OutputConstructor]
        private GetUpstreamsResult(
            ImmutableArray<Outputs.GetUpstreamsFilterResult> filters,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetUpstreamsResultResult> results,

            string upstreamId)
        {
            Filters = filters;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            UpstreamId = upstreamId;
        }
    }
}