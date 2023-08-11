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
    public static class GetContainerGroup
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf container_group
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
        ///         var containerGroup = Output.Create(Tencentcloud.Tsf.GetContainerGroup.InvokeAsync(new Tencentcloud.Tsf.GetContainerGroupArgs
        ///         {
        ///             ApplicationId = "application-a24x29xv",
        ///             ClusterId = "cluster-vwgj5e6y",
        ///             NamespaceId = "namespace-aemrg36v",
        ///             OrderBy = "createTime",
        ///             OrderType = 0,
        ///             SearchWord = "keep",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetContainerGroupResult> InvokeAsync(GetContainerGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContainerGroupResult>("tencentcloud:Tsf/getContainerGroup:getContainerGroup", args ?? new GetContainerGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf container_group
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
        ///         var containerGroup = Output.Create(Tencentcloud.Tsf.GetContainerGroup.InvokeAsync(new Tencentcloud.Tsf.GetContainerGroupArgs
        ///         {
        ///             ApplicationId = "application-a24x29xv",
        ///             ClusterId = "cluster-vwgj5e6y",
        ///             NamespaceId = "namespace-aemrg36v",
        ///             OrderBy = "createTime",
        ///             OrderType = 0,
        ///             SearchWord = "keep",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetContainerGroupResult> Invoke(GetContainerGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetContainerGroupResult>("tencentcloud:Tsf/getContainerGroup:getContainerGroup", args ?? new GetContainerGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ApplicationId, required.
        /// </summary>
        [Input("applicationId")]
        public string? ApplicationId { get; set; }

        /// <summary>
        /// Cluster Id.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// Namespace Id.
        /// </summary>
        [Input("namespaceId")]
        public string? NamespaceId { get; set; }

        /// <summary>
        /// The sorting field. By default, it is the createTime field. Supports id, name, createTime.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// The sorting order. By default, it is 1, indicating descending order. 0 indicates ascending order, and 1 indicates descending order.
        /// </summary>
        [Input("orderType")]
        public int? OrderType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// search word, support group name.
        /// </summary>
        [Input("searchWord")]
        public string? SearchWord { get; set; }

        public GetContainerGroupArgs()
        {
        }
    }

    public sealed class GetContainerGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ApplicationId, required.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Cluster Id.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Namespace Id.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// The sorting field. By default, it is the createTime field. Supports id, name, createTime.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// The sorting order. By default, it is 1, indicating descending order. 0 indicates ascending order, and 1 indicates descending order.
        /// </summary>
        [Input("orderType")]
        public Input<int>? OrderType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// search word, support group name.
        /// </summary>
        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        public GetContainerGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContainerGroupResult
    {
        public readonly string? ApplicationId;
        /// <summary>
        /// Cluster Id.Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly string? ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Namespace Id.Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly string? NamespaceId;
        public readonly string? OrderBy;
        public readonly int? OrderType;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// result list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContainerGroupResultResult> Results;
        public readonly string? SearchWord;

        [OutputConstructor]
        private GetContainerGroupResult(
            string? applicationId,

            string? clusterId,

            string id,

            string? namespaceId,

            string? orderBy,

            int? orderType,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetContainerGroupResultResult> results,

            string? searchWord)
        {
            ApplicationId = applicationId;
            ClusterId = clusterId;
            Id = id;
            NamespaceId = namespaceId;
            OrderBy = orderBy;
            OrderType = orderType;
            ResultOutputFile = resultOutputFile;
            Results = results;
            SearchWord = searchWord;
        }
    }
}