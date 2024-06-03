// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf
{
    /// <summary>
    /// Provides a resource to create a waf custom_white_rule
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
    ///     var example = new Tencentcloud.Waf.CustomWhiteRule("example", new()
    ///     {
    ///         Bypass = "geoip,cc,owasp",
    ///         Domain = "test.com",
    ///         ExpireTime = "0",
    ///         SortId = "30",
    ///         Status = "1",
    ///         Strategies = new[]
    ///         {
    ///             new Tencentcloud.Waf.Inputs.CustomWhiteRuleStrategyArgs
    ///             {
    ///                 Arg = "",
    ///                 CompareFunc = "ipmatch",
    ///                 Content = "1.1.1.1",
    ///                 Field = "IP",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// waf custom_white_rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Waf/customWhiteRule:CustomWhiteRule example test.com#1100310837
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Waf/customWhiteRule:CustomWhiteRule")]
    public partial class CustomWhiteRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Details of bypass.
        /// </summary>
        [Output("bypass")]
        public Output<string> Bypass { get; private set; } = null!;

        /// <summary>
        /// Domain name that needs to add policy.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Expiration time, measured in seconds, such as 1677254399, which means the expiration time is 2023-02-24 23:59:59 0 means never expires.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// rule ID.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// Priority, value range 1-100, The smaller the number, the higher the execution priority of this rule.
        /// </summary>
        [Output("sortId")]
        public Output<string> SortId { get; private set; } = null!;

        /// <summary>
        /// The status of the switch, 1 is on, 0 is off, default 1.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Strategies detail.
        /// </summary>
        [Output("strategies")]
        public Output<ImmutableArray<Outputs.CustomWhiteRuleStrategy>> Strategies { get; private set; } = null!;


        /// <summary>
        /// Create a CustomWhiteRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomWhiteRule(string name, CustomWhiteRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/customWhiteRule:CustomWhiteRule", name, args ?? new CustomWhiteRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomWhiteRule(string name, Input<string> id, CustomWhiteRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/customWhiteRule:CustomWhiteRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomWhiteRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomWhiteRule Get(string name, Input<string> id, CustomWhiteRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomWhiteRule(name, id, state, options);
        }
    }

    public sealed class CustomWhiteRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of bypass.
        /// </summary>
        [Input("bypass", required: true)]
        public Input<string> Bypass { get; set; } = null!;

        /// <summary>
        /// Domain name that needs to add policy.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Expiration time, measured in seconds, such as 1677254399, which means the expiration time is 2023-02-24 23:59:59 0 means never expires.
        /// </summary>
        [Input("expireTime", required: true)]
        public Input<string> ExpireTime { get; set; } = null!;

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority, value range 1-100, The smaller the number, the higher the execution priority of this rule.
        /// </summary>
        [Input("sortId", required: true)]
        public Input<string> SortId { get; set; } = null!;

        /// <summary>
        /// The status of the switch, 1 is on, 0 is off, default 1.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("strategies", required: true)]
        private InputList<Inputs.CustomWhiteRuleStrategyArgs>? _strategies;

        /// <summary>
        /// Strategies detail.
        /// </summary>
        public InputList<Inputs.CustomWhiteRuleStrategyArgs> Strategies
        {
            get => _strategies ?? (_strategies = new InputList<Inputs.CustomWhiteRuleStrategyArgs>());
            set => _strategies = value;
        }

        public CustomWhiteRuleArgs()
        {
        }
        public static new CustomWhiteRuleArgs Empty => new CustomWhiteRuleArgs();
    }

    public sealed class CustomWhiteRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of bypass.
        /// </summary>
        [Input("bypass")]
        public Input<string>? Bypass { get; set; }

        /// <summary>
        /// Domain name that needs to add policy.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Expiration time, measured in seconds, such as 1677254399, which means the expiration time is 2023-02-24 23:59:59 0 means never expires.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// Priority, value range 1-100, The smaller the number, the higher the execution priority of this rule.
        /// </summary>
        [Input("sortId")]
        public Input<string>? SortId { get; set; }

        /// <summary>
        /// The status of the switch, 1 is on, 0 is off, default 1.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("strategies")]
        private InputList<Inputs.CustomWhiteRuleStrategyGetArgs>? _strategies;

        /// <summary>
        /// Strategies detail.
        /// </summary>
        public InputList<Inputs.CustomWhiteRuleStrategyGetArgs> Strategies
        {
            get => _strategies ?? (_strategies = new InputList<Inputs.CustomWhiteRuleStrategyGetArgs>());
            set => _strategies = value;
        }

        public CustomWhiteRuleState()
        {
        }
        public static new CustomWhiteRuleState Empty => new CustomWhiteRuleState();
    }
}
