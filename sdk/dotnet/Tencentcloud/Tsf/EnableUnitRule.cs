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
    /// <summary>
    /// Provides a resource to create a tsf enable_unit_rule
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var enableUnitRule = new Tencentcloud.Tsf.EnableUnitRule("enableUnitRule", new()
    ///     {
    ///         RuleId = "unit-rl-is9m4nxz",
    ///         Switch = "enabled",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// tsf enable_unit_rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Tsf/enableUnitRule:EnableUnitRule enable_unit_rule enable_unit_rule_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/enableUnitRule:EnableUnitRule")]
    public partial class EnableUnitRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// api ID.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// switch, on: `enabled`, off: `disabled`.
        /// </summary>
        [Output("switch")]
        public Output<string> Switch { get; private set; } = null!;


        /// <summary>
        /// Create a EnableUnitRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnableUnitRule(string name, EnableUnitRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/enableUnitRule:EnableUnitRule", name, args ?? new EnableUnitRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnableUnitRule(string name, Input<string> id, EnableUnitRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/enableUnitRule:EnableUnitRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnableUnitRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnableUnitRule Get(string name, Input<string> id, EnableUnitRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new EnableUnitRule(name, id, state, options);
        }
    }

    public sealed class EnableUnitRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// api ID.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        /// <summary>
        /// switch, on: `enabled`, off: `disabled`.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public EnableUnitRuleArgs()
        {
        }
        public static new EnableUnitRuleArgs Empty => new EnableUnitRuleArgs();
    }

    public sealed class EnableUnitRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// api ID.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// switch, on: `enabled`, off: `disabled`.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public EnableUnitRuleState()
        {
        }
        public static new EnableUnitRuleState Empty => new EnableUnitRuleState();
    }
}
