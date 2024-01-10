// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse
{
    public static class GetGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of tse groups
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
        ///         var groups = Output.Create(Tencentcloud.Tse.GetGroups.InvokeAsync(new Tencentcloud.Tse.GetGroupsArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Tse.Inputs.GetGroupsFilterArgs
        ///                 {
        ///                     Name = "GroupId",
        ///                     Values = 
        ///                     {
        ///                         "group-013c0d8e",
        ///                     },
        ///                 },
        ///             },
        ///             GatewayId = "gateway-ddbb709b",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("tencentcloud:Tse/getGroups:getGroups", args ?? new GetGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tse groups
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
        ///         var groups = Output.Create(Tencentcloud.Tse.GetGroups.InvokeAsync(new Tencentcloud.Tse.GetGroupsArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Tse.Inputs.GetGroupsFilterArgs
        ///                 {
        ///                     Name = "GroupId",
        ///                     Values = 
        ///                     {
        ///                         "group-013c0d8e",
        ///                     },
        ///                 },
        ///             },
        ///             GatewayId = "gateway-ddbb709b",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupsResult> Invoke(GetGroupsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupsResult>("tencentcloud:Tse/getGroups:getGroups", args ?? new GetGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetGroupsFilterArgs>? _filters;

        /// <summary>
        /// filter conditions, valid value:Name,GroupId.
        /// </summary>
        public List<Inputs.GetGroupsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetGroupsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public string GatewayId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetGroupsArgs()
        {
        }
    }

    public sealed class GetGroupsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetGroupsFilterInputArgs>? _filters;

        /// <summary>
        /// filter conditions, valid value:Name,GroupId.
        /// </summary>
        public InputList<Inputs.GetGroupsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetGroupsFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetGroupsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        public readonly ImmutableArray<Outputs.GetGroupsFilterResult> Filters;
        /// <summary>
        /// gateway ID.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// groups information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsResultResult> Results;

        [OutputConstructor]
        private GetGroupsResult(
            ImmutableArray<Outputs.GetGroupsFilterResult> filters,

            string gatewayId,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetGroupsResultResult> results)
        {
            Filters = filters;
            GatewayId = gatewayId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
        }
    }
}
