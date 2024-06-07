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
    public static class GetDescribeHostApiGatewayInstanceList
    {
        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_api_gateway_instance_list
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
        ///     var describeHostApiGatewayInstanceList = Tencentcloud.Ssl.GetDescribeHostApiGatewayInstanceList.Invoke(new()
        ///     {
        ///         CertificateId = "9Bpk7XOu",
        ///         ResourceType = "apiGateway",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDescribeHostApiGatewayInstanceListResult> InvokeAsync(GetDescribeHostApiGatewayInstanceListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostApiGatewayInstanceListResult>("tencentcloud:Ssl/getDescribeHostApiGatewayInstanceList:getDescribeHostApiGatewayInstanceList", args ?? new GetDescribeHostApiGatewayInstanceListArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_api_gateway_instance_list
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
        ///     var describeHostApiGatewayInstanceList = Tencentcloud.Ssl.GetDescribeHostApiGatewayInstanceList.Invoke(new()
        ///     {
        ///         CertificateId = "9Bpk7XOu",
        ///         ResourceType = "apiGateway",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDescribeHostApiGatewayInstanceListResult> Invoke(GetDescribeHostApiGatewayInstanceListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeHostApiGatewayInstanceListResult>("tencentcloud:Ssl/getDescribeHostApiGatewayInstanceList:getDescribeHostApiGatewayInstanceList", args ?? new GetDescribeHostApiGatewayInstanceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostApiGatewayInstanceListArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate ID to be deployed.
        /// </summary>
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetDescribeHostApiGatewayInstanceListFilterArgs>? _filters;

        /// <summary>
        /// List of filtering parameters; Filterkey: domainmatch.
        /// </summary>
        public List<Inputs.GetDescribeHostApiGatewayInstanceListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDescribeHostApiGatewayInstanceListFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
        /// </summary>
        [Input("isCache")]
        public int? IsCache { get; set; }

        /// <summary>
        /// Deployed certificate ID.
        /// </summary>
        [Input("oldCertificateId")]
        public string? OldCertificateId { get; set; }

        /// <summary>
        /// Deploy resource type.
        /// </summary>
        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeHostApiGatewayInstanceListArgs()
        {
        }
        public static new GetDescribeHostApiGatewayInstanceListArgs Empty => new GetDescribeHostApiGatewayInstanceListArgs();
    }

    public sealed class GetDescribeHostApiGatewayInstanceListInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate ID to be deployed.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetDescribeHostApiGatewayInstanceListFilterInputArgs>? _filters;

        /// <summary>
        /// List of filtering parameters; Filterkey: domainmatch.
        /// </summary>
        public InputList<Inputs.GetDescribeHostApiGatewayInstanceListFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDescribeHostApiGatewayInstanceListFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
        /// </summary>
        [Input("isCache")]
        public Input<int>? IsCache { get; set; }

        /// <summary>
        /// Deployed certificate ID.
        /// </summary>
        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        /// <summary>
        /// Deploy resource type.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeHostApiGatewayInstanceListInvokeArgs()
        {
        }
        public static new GetDescribeHostApiGatewayInstanceListInvokeArgs Empty => new GetDescribeHostApiGatewayInstanceListInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeHostApiGatewayInstanceListResult
    {
        public readonly string CertificateId;
        public readonly ImmutableArray<Outputs.GetDescribeHostApiGatewayInstanceListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Apigateway instance listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeHostApiGatewayInstanceListInstanceListResult> InstanceLists;
        public readonly int? IsCache;
        public readonly string? OldCertificateId;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeHostApiGatewayInstanceListResult(
            string certificateId,

            ImmutableArray<Outputs.GetDescribeHostApiGatewayInstanceListFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetDescribeHostApiGatewayInstanceListInstanceListResult> instanceLists,

            int? isCache,

            string? oldCertificateId,

            string resourceType,

            string? resultOutputFile)
        {
            CertificateId = certificateId;
            Filters = filters;
            Id = id;
            InstanceLists = instanceLists;
            IsCache = isCache;
            OldCertificateId = oldCertificateId;
            ResourceType = resourceType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
