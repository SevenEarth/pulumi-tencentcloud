// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm
{
    public static class GetChcDeniedActions
    {
        /// <summary>
        /// Use this data source to query detailed information of cvm chc_denied_actions
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
        ///         var chcDeniedActions = Output.Create(Tencentcloud.Cvm.GetChcDeniedActions.InvokeAsync(new Tencentcloud.Cvm.GetChcDeniedActionsArgs
        ///         {
        ///             ChcIds = 
        ///             {
        ///                 "chc-xxxxx",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetChcDeniedActionsResult> InvokeAsync(GetChcDeniedActionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetChcDeniedActionsResult>("tencentcloud:Cvm/getChcDeniedActions:getChcDeniedActions", args ?? new GetChcDeniedActionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cvm chc_denied_actions
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
        ///         var chcDeniedActions = Output.Create(Tencentcloud.Cvm.GetChcDeniedActions.InvokeAsync(new Tencentcloud.Cvm.GetChcDeniedActionsArgs
        ///         {
        ///             ChcIds = 
        ///             {
        ///                 "chc-xxxxx",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetChcDeniedActionsResult> Invoke(GetChcDeniedActionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetChcDeniedActionsResult>("tencentcloud:Cvm/getChcDeniedActions:getChcDeniedActions", args ?? new GetChcDeniedActionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChcDeniedActionsArgs : Pulumi.InvokeArgs
    {
        [Input("chcIds", required: true)]
        private List<string>? _chcIds;

        /// <summary>
        /// CHC host IDs.
        /// </summary>
        public List<string> ChcIds
        {
            get => _chcIds ?? (_chcIds = new List<string>());
            set => _chcIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetChcDeniedActionsArgs()
        {
        }
    }

    public sealed class GetChcDeniedActionsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("chcIds", required: true)]
        private InputList<string>? _chcIds;

        /// <summary>
        /// CHC host IDs.
        /// </summary>
        public InputList<string> ChcIds
        {
            get => _chcIds ?? (_chcIds = new InputList<string>());
            set => _chcIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetChcDeniedActionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetChcDeniedActionsResult
    {
        /// <summary>
        /// Actions not allowed for the CHC instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetChcDeniedActionsChcHostDeniedActionSetResult> ChcHostDeniedActionSets;
        public readonly ImmutableArray<string> ChcIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetChcDeniedActionsResult(
            ImmutableArray<Outputs.GetChcDeniedActionsChcHostDeniedActionSetResult> chcHostDeniedActionSets,

            ImmutableArray<string> chcIds,

            string id,

            string? resultOutputFile)
        {
            ChcHostDeniedActionSets = chcHostDeniedActionSets;
            ChcIds = chcIds;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
