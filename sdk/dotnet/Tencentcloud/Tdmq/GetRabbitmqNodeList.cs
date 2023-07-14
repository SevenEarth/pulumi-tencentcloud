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
    public static class GetRabbitmqNodeList
    {
        public static Task<GetRabbitmqNodeListResult> InvokeAsync(GetRabbitmqNodeListArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRabbitmqNodeListResult>("tencentcloud:Tdmq/getRabbitmqNodeList:getRabbitmqNodeList", args ?? new GetRabbitmqNodeListArgs(), options.WithDefaults());

        public static Output<GetRabbitmqNodeListResult> Invoke(GetRabbitmqNodeListInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRabbitmqNodeListResult>("tencentcloud:Tdmq/getRabbitmqNodeList:getRabbitmqNodeList", args ?? new GetRabbitmqNodeListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRabbitmqNodeListArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetRabbitmqNodeListFilterArgs>? _filters;
        public List<Inputs.GetRabbitmqNodeListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetRabbitmqNodeListFilterArgs>());
            set => _filters = value;
        }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("nodeName")]
        public string? NodeName { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("sortElement")]
        public string? SortElement { get; set; }

        [Input("sortOrder")]
        public string? SortOrder { get; set; }

        public GetRabbitmqNodeListArgs()
        {
        }
    }

    public sealed class GetRabbitmqNodeListInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetRabbitmqNodeListFilterInputArgs>? _filters;
        public InputList<Inputs.GetRabbitmqNodeListFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetRabbitmqNodeListFilterInputArgs>());
            set => _filters = value;
        }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("sortElement")]
        public Input<string>? SortElement { get; set; }

        [Input("sortOrder")]
        public Input<string>? SortOrder { get; set; }

        public GetRabbitmqNodeListInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRabbitmqNodeListResult
    {
        public readonly ImmutableArray<Outputs.GetRabbitmqNodeListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly ImmutableArray<Outputs.GetRabbitmqNodeListNodeListResult> NodeLists;
        public readonly string? NodeName;
        public readonly string? ResultOutputFile;
        public readonly string? SortElement;
        public readonly string? SortOrder;

        [OutputConstructor]
        private GetRabbitmqNodeListResult(
            ImmutableArray<Outputs.GetRabbitmqNodeListFilterResult> filters,

            string id,

            string instanceId,

            ImmutableArray<Outputs.GetRabbitmqNodeListNodeListResult> nodeLists,

            string? nodeName,

            string? resultOutputFile,

            string? sortElement,

            string? sortOrder)
        {
            Filters = filters;
            Id = id;
            InstanceId = instanceId;
            NodeLists = nodeLists;
            NodeName = nodeName;
            ResultOutputFile = resultOutputFile;
            SortElement = sortElement;
            SortOrder = sortOrder;
        }
    }
}
