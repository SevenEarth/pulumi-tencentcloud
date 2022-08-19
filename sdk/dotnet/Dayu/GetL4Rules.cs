// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu
{
    public static class GetL4Rules
    {
        /// <summary>
        /// Use this data source to query dayu layer 4 rules
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
        ///         var nameTest = Output.Create(Tencentcloud.Dayu.GetL4Rules.InvokeAsync(new Tencentcloud.Dayu.GetL4RulesArgs
        ///         {
        ///             ResourceType = tencentcloud_dayu_l4_rule.Test_rule.Resource_type,
        ///             ResourceId = tencentcloud_dayu_l4_rule.Test_rule.Resource_id,
        ///             Name = tencentcloud_dayu_l4_rule.Test_rule.Name,
        ///         }));
        ///         var idTest = Output.Create(Tencentcloud.Dayu.GetL4Rules.InvokeAsync(new Tencentcloud.Dayu.GetL4RulesArgs
        ///         {
        ///             ResourceType = tencentcloud_dayu_l4_rule.Test_rule.Resource_type,
        ///             ResourceId = tencentcloud_dayu_l4_rule.Test_rule.Resource_id,
        ///             RuleId = tencentcloud_dayu_l4_rule.Test_rule.Rule_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetL4RulesResult> InvokeAsync(GetL4RulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetL4RulesResult>("tencentcloud:Dayu/getL4Rules:getL4Rules", args ?? new GetL4RulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query dayu layer 4 rules
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
        ///         var nameTest = Output.Create(Tencentcloud.Dayu.GetL4Rules.InvokeAsync(new Tencentcloud.Dayu.GetL4RulesArgs
        ///         {
        ///             ResourceType = tencentcloud_dayu_l4_rule.Test_rule.Resource_type,
        ///             ResourceId = tencentcloud_dayu_l4_rule.Test_rule.Resource_id,
        ///             Name = tencentcloud_dayu_l4_rule.Test_rule.Name,
        ///         }));
        ///         var idTest = Output.Create(Tencentcloud.Dayu.GetL4Rules.InvokeAsync(new Tencentcloud.Dayu.GetL4RulesArgs
        ///         {
        ///             ResourceType = tencentcloud_dayu_l4_rule.Test_rule.Resource_type,
        ///             ResourceId = tencentcloud_dayu_l4_rule.Test_rule.Resource_id,
        ///             RuleId = tencentcloud_dayu_l4_rule.Test_rule.Rule_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetL4RulesResult> Invoke(GetL4RulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetL4RulesResult>("tencentcloud:Dayu/getL4Rules:getL4Rules", args ?? new GetL4RulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetL4RulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the layer 4 rule to be queried.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Id of the resource that the layer 4 rule works for.
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        /// <summary>
        /// Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
        /// </summary>
        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Id of the layer 4 rule to be queried.
        /// </summary>
        [Input("ruleId")]
        public string? RuleId { get; set; }

        public GetL4RulesArgs()
        {
        }
    }

    public sealed class GetL4RulesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the layer 4 rule to be queried.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Id of the resource that the layer 4 rule works for.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Id of the layer 4 rule to be queried.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        public GetL4RulesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetL4RulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of layer 4 rules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetL4RulesListResult> Lists;
        /// <summary>
        /// Name of the rule.
        /// </summary>
        public readonly string? Name;
        public readonly string ResourceId;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// ID of the 4 layer rule.
        /// </summary>
        public readonly string? RuleId;

        [OutputConstructor]
        private GetL4RulesResult(
            string id,

            ImmutableArray<Outputs.GetL4RulesListResult> lists,

            string? name,

            string resourceId,

            string resourceType,

            string? resultOutputFile,

            string? ruleId)
        {
            Id = id;
            Lists = lists;
            Name = name;
            ResourceId = resourceId;
            ResourceType = resourceType;
            ResultOutputFile = resultOutputFile;
            RuleId = ruleId;
        }
    }
}
