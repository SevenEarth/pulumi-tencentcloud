// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class TaskTaskRule
    {
        /// <summary>
        /// Cron type rule, cron expression.
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// time interval, in milliseconds.
        /// </summary>
        public readonly int? RepeatInterval;
        /// <summary>
        /// trigger rule type, Cron/Repeat.
        /// </summary>
        public readonly string RuleType;

        [OutputConstructor]
        private TaskTaskRule(
            string? expression,

            int? repeatInterval,

            string ruleType)
        {
            Expression = expression;
            RepeatInterval = repeatInterval;
            RuleType = ruleType;
        }
    }
}
