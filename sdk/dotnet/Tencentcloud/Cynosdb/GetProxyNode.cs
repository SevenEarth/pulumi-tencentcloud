// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    public static class GetProxyNode
    {
        /// <summary>
        /// Use this data source to query detailed information of cynosdb proxy_node
        /// </summary>
        public static Task<GetProxyNodeResult> InvokeAsync(GetProxyNodeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProxyNodeResult>("tencentcloud:Cynosdb/getProxyNode:getProxyNode", args ?? new GetProxyNodeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cynosdb proxy_node
        /// </summary>
        public static Output<GetProxyNodeResult> Invoke(GetProxyNodeInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProxyNodeResult>("tencentcloud:Cynosdb/getProxyNode:getProxyNode", args ?? new GetProxyNodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProxyNodeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetProxyNodeFilterArgs>? _filters;

        /// <summary>
        /// Search criteria, if there are multiple filters, the relationship between the filters is a logical AND relationship.
        /// </summary>
        public List<Inputs.GetProxyNodeFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetProxyNodeFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Sort field, value range:CREATETIME: creation time; PRIODENDTIME: expiration time.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// Sort type, value range:ASC: ascending sort; DESC: descending sort.
        /// </summary>
        [Input("orderByType")]
        public string? OrderByType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetProxyNodeArgs()
        {
        }
    }

    public sealed class GetProxyNodeInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetProxyNodeFilterInputArgs>? _filters;

        /// <summary>
        /// Search criteria, if there are multiple filters, the relationship between the filters is a logical AND relationship.
        /// </summary>
        public InputList<Inputs.GetProxyNodeFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetProxyNodeFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Sort field, value range:CREATETIME: creation time; PRIODENDTIME: expiration time.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// Sort type, value range:ASC: ascending sort; DESC: descending sort.
        /// </summary>
        [Input("orderByType")]
        public Input<string>? OrderByType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetProxyNodeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProxyNodeResult
    {
        public readonly ImmutableArray<Outputs.GetProxyNodeFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OrderBy;
        public readonly string? OrderByType;
        /// <summary>
        /// Database Agent Node List.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyNodeProxyNodeInfoResult> ProxyNodeInfos;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetProxyNodeResult(
            ImmutableArray<Outputs.GetProxyNodeFilterResult> filters,

            string id,

            string? orderBy,

            string? orderByType,

            ImmutableArray<Outputs.GetProxyNodeProxyNodeInfoResult> proxyNodeInfos,

            string? resultOutputFile)
        {
            Filters = filters;
            Id = id;
            OrderBy = orderBy;
            OrderByType = orderByType;
            ProxyNodeInfos = proxyNodeInfos;
            ResultOutputFile = resultOutputFile;
        }
    }
}
