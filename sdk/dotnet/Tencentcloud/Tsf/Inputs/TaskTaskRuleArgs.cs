// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class TaskTaskRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron type rule, cron expression.
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        /// <summary>
        /// time interval, in milliseconds.
        /// </summary>
        [Input("repeatInterval")]
        public Input<int>? RepeatInterval { get; set; }

        /// <summary>
        /// trigger rule type, Cron/Repeat.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        public TaskTaskRuleArgs()
        {
        }
        public static new TaskTaskRuleArgs Empty => new TaskTaskRuleArgs();
    }
}
