// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Emr
{
    public static class GetInstance
    {
        /// <summary>
        /// Provides an available EMR for the user.
        /// 
        /// The EMR data source fetch proper EMR from user's EMR pool.
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
        ///         var myEmr = Output.Create(Tencentcloud.Emr.GetInstance.InvokeAsync(new Tencentcloud.Emr.GetInstanceArgs
        ///         {
        ///             DisplayStrategy = "clusterList",
        ///             InstanceIds = 
        ///             {
        ///                 "emr-rnzqrleq",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("tencentcloud:Emr/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Provides an available EMR for the user.
        /// 
        /// The EMR data source fetch proper EMR from user's EMR pool.
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
        ///         var myEmr = Output.Create(Tencentcloud.Emr.GetInstance.InvokeAsync(new Tencentcloud.Emr.GetInstanceArgs
        ///         {
        ///             DisplayStrategy = "clusterList",
        ///             InstanceIds = 
        ///             {
        ///                 "emr-rnzqrleq",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("tencentcloud:Emr/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Display strategy(e.g.:clusterList, monitorManage).
        /// </summary>
        [Input("displayStrategy", required: true)]
        public string DisplayStrategy { get; set; } = null!;

        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// fetch all instances with same prefix(e.g.:emr-xxxxxx).
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Fetch all instances which owner same project. Default 0 meaning use default project id.
        /// </summary>
        [Input("projectId")]
        public int? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceArgs()
        {
        }
    }

    public sealed class GetInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Display strategy(e.g.:clusterList, monitorManage).
        /// </summary>
        [Input("displayStrategy", required: true)]
        public Input<string> DisplayStrategy { get; set; } = null!;

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// fetch all instances with same prefix(e.g.:emr-xxxxxx).
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Fetch all instances which owner same project. Default 0 meaning use default project id.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// A list of clusters will be exported and its every element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceClusterResult> Clusters;
        public readonly string DisplayStrategy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        /// <summary>
        /// Project id of instance.
        /// </summary>
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<Outputs.GetInstanceClusterResult> clusters,

            string displayStrategy,

            string id,

            ImmutableArray<string> instanceIds,

            int? projectId,

            string? resultOutputFile)
        {
            Clusters = clusters;
            DisplayStrategy = displayStrategy;
            Id = id;
            InstanceIds = instanceIds;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
