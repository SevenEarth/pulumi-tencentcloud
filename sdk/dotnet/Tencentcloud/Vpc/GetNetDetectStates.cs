// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    public static class GetNetDetectStates
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc net_detect_states
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
        ///         var netDetectStates = Output.Create(Tencentcloud.Vpc.GetNetDetectStates.InvokeAsync(new Tencentcloud.Vpc.GetNetDetectStatesArgs
        ///         {
        ///             NetDetectIds = 
        ///             {
        ///                 "netd-12345678",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetDetectStatesResult> InvokeAsync(GetNetDetectStatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetDetectStatesResult>("tencentcloud:Vpc/getNetDetectStates:getNetDetectStates", args ?? new GetNetDetectStatesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc net_detect_states
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
        ///         var netDetectStates = Output.Create(Tencentcloud.Vpc.GetNetDetectStates.InvokeAsync(new Tencentcloud.Vpc.GetNetDetectStatesArgs
        ///         {
        ///             NetDetectIds = 
        ///             {
        ///                 "netd-12345678",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetDetectStatesResult> Invoke(GetNetDetectStatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetDetectStatesResult>("tencentcloud:Vpc/getNetDetectStates:getNetDetectStates", args ?? new GetNetDetectStatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetDetectStatesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetNetDetectStatesFilterArgs>? _filters;

        /// <summary>
        /// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.
        /// </summary>
        public List<Inputs.GetNetDetectStatesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetNetDetectStatesFilterArgs>());
            set => _filters = value;
        }

        [Input("netDetectIds")]
        private List<string>? _netDetectIds;

        /// <summary>
        /// The array of network detection instance `IDs`, such as [`netd-12345678`].
        /// </summary>
        public List<string> NetDetectIds
        {
            get => _netDetectIds ?? (_netDetectIds = new List<string>());
            set => _netDetectIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetNetDetectStatesArgs()
        {
        }
    }

    public sealed class GetNetDetectStatesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetNetDetectStatesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.
        /// </summary>
        public InputList<Inputs.GetNetDetectStatesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetNetDetectStatesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("netDetectIds")]
        private InputList<string>? _netDetectIds;

        /// <summary>
        /// The array of network detection instance `IDs`, such as [`netd-12345678`].
        /// </summary>
        public InputList<string> NetDetectIds
        {
            get => _netDetectIds ?? (_netDetectIds = new InputList<string>());
            set => _netDetectIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetNetDetectStatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetDetectStatesResult
    {
        public readonly ImmutableArray<Outputs.GetNetDetectStatesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> NetDetectIds;
        /// <summary>
        /// The array of network detection verification results that meet requirements.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetDetectStatesNetDetectStateSetResult> NetDetectStateSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetNetDetectStatesResult(
            ImmutableArray<Outputs.GetNetDetectStatesFilterResult> filters,

            string id,

            ImmutableArray<string> netDetectIds,

            ImmutableArray<Outputs.GetNetDetectStatesNetDetectStateSetResult> netDetectStateSets,

            string? resultOutputFile)
        {
            Filters = filters;
            Id = id;
            NetDetectIds = netDetectIds;
            NetDetectStateSets = netDetectStateSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
