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
    public static class GetDescribeHostLiveInstanceList
    {
        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_live_instance_list
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
        ///         var describeHostLiveInstanceList = Output.Create(Tencentcloud.Ssl.GetDescribeHostLiveInstanceList.InvokeAsync(new Tencentcloud.Ssl.GetDescribeHostLiveInstanceListArgs
        ///         {
        ///             CertificateId = "8u8DII0l",
        ///             ResourceType = "live",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDescribeHostLiveInstanceListResult> InvokeAsync(GetDescribeHostLiveInstanceListArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostLiveInstanceListResult>("tencentcloud:Ssl/getDescribeHostLiveInstanceList:getDescribeHostLiveInstanceList", args ?? new GetDescribeHostLiveInstanceListArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_live_instance_list
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
        ///         var describeHostLiveInstanceList = Output.Create(Tencentcloud.Ssl.GetDescribeHostLiveInstanceList.InvokeAsync(new Tencentcloud.Ssl.GetDescribeHostLiveInstanceListArgs
        ///         {
        ///             CertificateId = "8u8DII0l",
        ///             ResourceType = "live",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDescribeHostLiveInstanceListResult> Invoke(GetDescribeHostLiveInstanceListInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDescribeHostLiveInstanceListResult>("tencentcloud:Ssl/getDescribeHostLiveInstanceList:getDescribeHostLiveInstanceList", args ?? new GetDescribeHostLiveInstanceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostLiveInstanceListArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate ID to be deployed.
        /// </summary>
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetDescribeHostLiveInstanceListFilterArgs>? _filters;

        /// <summary>
        /// List of filtering parameters; Filterkey: domainmatch.
        /// </summary>
        public List<Inputs.GetDescribeHostLiveInstanceListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDescribeHostLiveInstanceListFilterArgs>());
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

        public GetDescribeHostLiveInstanceListArgs()
        {
        }
    }

    public sealed class GetDescribeHostLiveInstanceListInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Certificate ID to be deployed.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetDescribeHostLiveInstanceListFilterInputArgs>? _filters;

        /// <summary>
        /// List of filtering parameters; Filterkey: domainmatch.
        /// </summary>
        public InputList<Inputs.GetDescribeHostLiveInstanceListFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDescribeHostLiveInstanceListFilterInputArgs>());
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

        public GetDescribeHostLiveInstanceListInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDescribeHostLiveInstanceListResult
    {
        public readonly string CertificateId;
        public readonly ImmutableArray<Outputs.GetDescribeHostLiveInstanceListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Live instance listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeHostLiveInstanceListInstanceListResult> InstanceLists;
        public readonly int? IsCache;
        public readonly string? OldCertificateId;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeHostLiveInstanceListResult(
            string certificateId,

            ImmutableArray<Outputs.GetDescribeHostLiveInstanceListFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetDescribeHostLiveInstanceListInstanceListResult> instanceLists,

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
