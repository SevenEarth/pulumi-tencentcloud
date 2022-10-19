// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class TmpTkeAlertPolicyAlertRuleRule
    {
        /// <summary>
        /// Refer to annotations in prometheus rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.TmpTkeAlertPolicyAlertRuleRuleAnnotation> Annotations;
        /// <summary>
        /// A description of the rule.
        /// </summary>
        public readonly string? Describe;
        /// <summary>
        /// Time of duration.
        /// </summary>
        public readonly string For;
        /// <summary>
        /// Extra labels.
        /// </summary>
        public readonly ImmutableArray<Outputs.TmpTkeAlertPolicyAlertRuleRuleLabel> Labels;
        /// <summary>
        /// Rule name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Prometheus statement.
        /// </summary>
        public readonly string Rule;
        /// <summary>
        /// Alarm rule status.
        /// </summary>
        public readonly int? RuleState;
        /// <summary>
        /// Alert sending template.
        /// </summary>
        public readonly string Template;

        [OutputConstructor]
        private TmpTkeAlertPolicyAlertRuleRule(
            ImmutableArray<Outputs.TmpTkeAlertPolicyAlertRuleRuleAnnotation> annotations,

            string? describe,

            string @for,

            ImmutableArray<Outputs.TmpTkeAlertPolicyAlertRuleRuleLabel> labels,

            string name,

            string rule,

            int? ruleState,

            string template)
        {
            Annotations = annotations;
            Describe = describe;
            For = @for;
            Labels = labels;
            Name = name;
            Rule = rule;
            RuleState = ruleState;
            Template = template;
        }
    }
}