// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a monitor tmpAlertRule
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tmpAlertRule = new Tencentcloud.Monitor.TmpAlertRule("tmpAlertRule", new Tencentcloud.Monitor.TmpAlertRuleArgs
    ///         {
    ///             Annotations = 
    ///             {
    ///                 new Tencentcloud.Monitor.Inputs.TmpAlertRuleAnnotationArgs
    ///                 {
    ///                     Key = "hello2",
    ///                     Value = "world2",
    ///                 },
    ///             },
    ///             Duration = "4m",
    ///             Expr = "up{service=\"rig-prometheus-agent\"}&gt;0",
    ///             InstanceId = "prom-c89b3b3u",
    ///             Labels = 
    ///             {
    ///                 new Tencentcloud.Monitor.Inputs.TmpAlertRuleLabelArgs
    ///                 {
    ///                     Key = "hello1",
    ///                     Value = "world1",
    ///                 },
    ///             },
    ///             Receivers = 
    ///             {
    ///                 "notice-l9ziyxw6",
    ///             },
    ///             RuleName = "test123",
    ///             RuleState = 2,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// monitor tmpAlertRule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Monitor/tmpAlertRule:TmpAlertRule tmpAlertRule instanceId#Rule_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpAlertRule:TmpAlertRule")]
    public partial class TmpAlertRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableArray<Outputs.TmpAlertRuleAnnotation>> Annotations { get; private set; } = null!;

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        [Output("duration")]
        public Output<string?> Duration { get; private set; } = null!;

        /// <summary>
        /// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
        /// </summary>
        [Output("expr")]
        public Output<string> Expr { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.TmpAlertRuleLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// Alarm notification template id list.
        /// </summary>
        [Output("receivers")]
        public Output<ImmutableArray<string>> Receivers { get; private set; } = null!;

        /// <summary>
        /// Rule name.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;

        /// <summary>
        /// Rule state code.
        /// </summary>
        [Output("ruleState")]
        public Output<int?> RuleState { get; private set; } = null!;

        /// <summary>
        /// Alarm Policy Template Classification.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TmpAlertRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpAlertRule(string name, TmpAlertRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpAlertRule:TmpAlertRule", name, args ?? new TmpAlertRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpAlertRule(string name, Input<string> id, TmpAlertRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpAlertRule:TmpAlertRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpAlertRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpAlertRule Get(string name, Input<string> id, TmpAlertRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpAlertRule(name, id, state, options);
        }
    }

    public sealed class TmpAlertRuleArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<Inputs.TmpAlertRuleAnnotationArgs>? _annotations;

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        public InputList<Inputs.TmpAlertRuleAnnotationArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.TmpAlertRuleAnnotationArgs>());
            set => _annotations = value;
        }

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
        /// </summary>
        [Input("expr", required: true)]
        public Input<string> Expr { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("labels")]
        private InputList<Inputs.TmpAlertRuleLabelArgs>? _labels;

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        public InputList<Inputs.TmpAlertRuleLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.TmpAlertRuleLabelArgs>());
            set => _labels = value;
        }

        [Input("receivers", required: true)]
        private InputList<string>? _receivers;

        /// <summary>
        /// Alarm notification template id list.
        /// </summary>
        public InputList<string> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<string>());
            set => _receivers = value;
        }

        /// <summary>
        /// Rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        /// <summary>
        /// Rule state code.
        /// </summary>
        [Input("ruleState")]
        public Input<int>? RuleState { get; set; }

        /// <summary>
        /// Alarm Policy Template Classification.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TmpAlertRuleArgs()
        {
        }
    }

    public sealed class TmpAlertRuleState : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<Inputs.TmpAlertRuleAnnotationGetArgs>? _annotations;

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        public InputList<Inputs.TmpAlertRuleAnnotationGetArgs> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<Inputs.TmpAlertRuleAnnotationGetArgs>());
            set => _annotations = value;
        }

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
        /// </summary>
        [Input("expr")]
        public Input<string>? Expr { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("labels")]
        private InputList<Inputs.TmpAlertRuleLabelGetArgs>? _labels;

        /// <summary>
        /// Rule alarm duration.
        /// </summary>
        public InputList<Inputs.TmpAlertRuleLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.TmpAlertRuleLabelGetArgs>());
            set => _labels = value;
        }

        [Input("receivers")]
        private InputList<string>? _receivers;

        /// <summary>
        /// Alarm notification template id list.
        /// </summary>
        public InputList<string> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<string>());
            set => _receivers = value;
        }

        /// <summary>
        /// Rule name.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        /// <summary>
        /// Rule state code.
        /// </summary>
        [Input("ruleState")]
        public Input<int>? RuleState { get; set; }

        /// <summary>
        /// Alarm Policy Template Classification.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TmpAlertRuleState()
        {
        }
    }
}
