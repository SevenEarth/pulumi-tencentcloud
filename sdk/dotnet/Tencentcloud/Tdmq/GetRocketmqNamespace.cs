// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq
{
    public static class GetRocketmqNamespace
    {
        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq namespace
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var cluster = new Tencentcloud.Tdmq.RocketmqCluster("cluster", new Tencentcloud.Tdmq.RocketmqClusterArgs
        ///         {
        ///             ClusterName = "test_rocketmq_namespace_sdatasource",
        ///             Remark = "test recket mq",
        ///         });
        ///         var namespacedata = new Tencentcloud.Tdmq.RocketmqNamespace("namespacedata", new Tencentcloud.Tdmq.RocketmqNamespaceArgs
        ///         {
        ///             ClusterId = cluster.ClusterId,
        ///             NamespaceName = "test_namespace_datasource",
        ///             Ttl = 65000,
        ///             RetentionTime = 65000,
        ///             Remark = "test namespace",
        ///         });
        ///         var @namespace = Tencentcloud.Tdmq.GetRocketmqNamespace.Invoke(new Tencentcloud.Tdmq.GetRocketmqNamespaceInvokeArgs
        ///         {
        ///             ClusterId = cluster.ClusterId,
        ///             NameKeyword = namespacedata.NamespaceName,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRocketmqNamespaceResult> InvokeAsync(GetRocketmqNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRocketmqNamespaceResult>("tencentcloud:Tdmq/getRocketmqNamespace:getRocketmqNamespace", args ?? new GetRocketmqNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdmqRocketmq namespace
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var cluster = new Tencentcloud.Tdmq.RocketmqCluster("cluster", new Tencentcloud.Tdmq.RocketmqClusterArgs
        ///         {
        ///             ClusterName = "test_rocketmq_namespace_sdatasource",
        ///             Remark = "test recket mq",
        ///         });
        ///         var namespacedata = new Tencentcloud.Tdmq.RocketmqNamespace("namespacedata", new Tencentcloud.Tdmq.RocketmqNamespaceArgs
        ///         {
        ///             ClusterId = cluster.ClusterId,
        ///             NamespaceName = "test_namespace_datasource",
        ///             Ttl = 65000,
        ///             RetentionTime = 65000,
        ///             Remark = "test namespace",
        ///         });
        ///         var @namespace = Tencentcloud.Tdmq.GetRocketmqNamespace.Invoke(new Tencentcloud.Tdmq.GetRocketmqNamespaceInvokeArgs
        ///         {
        ///             ClusterId = cluster.ClusterId,
        ///             NameKeyword = namespacedata.NamespaceName,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRocketmqNamespaceResult> Invoke(GetRocketmqNamespaceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRocketmqNamespaceResult>("tencentcloud:Tdmq/getRocketmqNamespace:getRocketmqNamespace", args ?? new GetRocketmqNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRocketmqNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Search by name.
        /// </summary>
        [Input("nameKeyword")]
        public string? NameKeyword { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRocketmqNamespaceArgs()
        {
        }
    }

    public sealed class GetRocketmqNamespaceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Search by name.
        /// </summary>
        [Input("nameKeyword")]
        public Input<string>? NameKeyword { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRocketmqNamespaceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRocketmqNamespaceResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameKeyword;
        /// <summary>
        /// List of namespaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRocketmqNamespaceNamespaceResult> Namespaces;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRocketmqNamespaceResult(
            string clusterId,

            string id,

            string? nameKeyword,

            ImmutableArray<Outputs.GetRocketmqNamespaceNamespaceResult> namespaces,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            Id = id;
            NameKeyword = nameKeyword;
            Namespaces = namespaces;
            ResultOutputFile = resultOutputFile;
        }
    }
}
