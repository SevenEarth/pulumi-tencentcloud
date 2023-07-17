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
    public static class GetPodInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf pod_instances
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
        ///         var podInstances = Output.Create(Tencentcloud.Tsf.GetPodInstances.InvokeAsync(new Tencentcloud.Tsf.GetPodInstancesArgs
        ///         {
        ///             GroupId = "group-ynd95rea",
        ///             PodNameLists = 
        ///             {
        ///                 "keep-terraform-6f8f977688-zvphm",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPodInstancesResult> InvokeAsync(GetPodInstancesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPodInstancesResult>("tencentcloud:Tsf/getPodInstances:getPodInstances", args ?? new GetPodInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf pod_instances
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
        ///         var podInstances = Output.Create(Tencentcloud.Tsf.GetPodInstances.InvokeAsync(new Tencentcloud.Tsf.GetPodInstancesArgs
        ///         {
        ///             GroupId = "group-ynd95rea",
        ///             PodNameLists = 
        ///             {
        ///                 "keep-terraform-6f8f977688-zvphm",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPodInstancesResult> Invoke(GetPodInstancesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPodInstancesResult>("tencentcloud:Tsf/getPodInstances:getPodInstances", args ?? new GetPodInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPodInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance&amp;amp;#39;s group ID.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        [Input("podNameLists")]
        private List<string>? _podNameLists;

        /// <summary>
        /// Filter, pod name list.
        /// </summary>
        public List<string> PodNameLists
        {
            get => _podNameLists ?? (_podNameLists = new List<string>());
            set => _podNameLists = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetPodInstancesArgs()
        {
        }
    }

    public sealed class GetPodInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance&amp;amp;#39;s group ID.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("podNameLists")]
        private InputList<string>? _podNameLists;

        /// <summary>
        /// Filter, pod name list.
        /// </summary>
        public InputList<string> PodNameLists
        {
            get => _podNameLists ?? (_podNameLists = new InputList<string>());
            set => _podNameLists = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetPodInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPodInstancesResult
    {
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> PodNameLists;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// pod instance list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPodInstancesResultResult> Results;

        [OutputConstructor]
        private GetPodInstancesResult(
            string groupId,

            string id,

            ImmutableArray<string> podNameLists,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetPodInstancesResultResult> results)
        {
            GroupId = groupId;
            Id = id;
            PodNameLists = podNameLists;
            ResultOutputFile = resultOutputFile;
            Results = results;
        }
    }
}
