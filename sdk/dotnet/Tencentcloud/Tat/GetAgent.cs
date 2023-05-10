// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tat
{
    public static class GetAgent
    {
        /// <summary>
        /// Use this data source to query detailed information of tat agent
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
        ///         var agent = Output.Create(Tencentcloud.Tat.GetAgent.InvokeAsync(new Tencentcloud.Tat.GetAgentArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Tat.Inputs.GetAgentFilterArgs
        ///                 {
        ///                     Name = "environment",
        ///                     Values = 
        ///                     {
        ///                         "Linux",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAgentResult> InvokeAsync(GetAgentArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAgentResult>("tencentcloud:Tat/getAgent:getAgent", args ?? new GetAgentArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tat agent
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
        ///         var agent = Output.Create(Tencentcloud.Tat.GetAgent.InvokeAsync(new Tencentcloud.Tat.GetAgentArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Tat.Inputs.GetAgentFilterArgs
        ///                 {
        ///                     Name = "environment",
        ///                     Values = 
        ///                     {
        ///                         "Linux",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAgentResult> Invoke(GetAgentInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAgentResult>("tencentcloud:Tat/getAgent:getAgent", args ?? new GetAgentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetAgentFilterArgs>? _filters;

        /// <summary>
        /// Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
        /// </summary>
        public List<Inputs.GetAgentFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetAgentFilterArgs>());
            set => _filters = value;
        }

        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// List of instance IDs for the query.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAgentArgs()
        {
        }
    }

    public sealed class GetAgentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetAgentFilterInputArgs>? _filters;

        /// <summary>
        /// Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
        /// </summary>
        public InputList<Inputs.GetAgentFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetAgentFilterInputArgs>());
            set => _filters = value;
        }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// List of instance IDs for the query.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAgentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAgentResult
    {
        /// <summary>
        /// List of agent message.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAgentAutomationAgentSetResult> AutomationAgentSets;
        public readonly ImmutableArray<Outputs.GetAgentFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetAgentResult(
            ImmutableArray<Outputs.GetAgentAutomationAgentSetResult> automationAgentSets,

            ImmutableArray<Outputs.GetAgentFilterResult> filters,

            string id,

            ImmutableArray<string> instanceIds,

            string? resultOutputFile)
        {
            AutomationAgentSets = automationAgentSets;
            Filters = filters;
            Id = id;
            InstanceIds = instanceIds;
            ResultOutputFile = resultOutputFile;
        }
    }
}
