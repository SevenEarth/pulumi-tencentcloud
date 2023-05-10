// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Outputs
{

    [OutputType]
    public sealed class ScaleRuleAutoscaler
    {
        /// <summary>
        /// name.
        /// </summary>
        public readonly string AutoscalerName;
        /// <summary>
        /// scaler based on cron configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScaleRuleAutoscalerCronHorizontalAutoscaler> CronHorizontalAutoscalers;
        /// <summary>
        /// description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// enable AutoScaler.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// scaler based on metrics.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScaleRuleAutoscalerHorizontalAutoscaler> HorizontalAutoscalers;
        /// <summary>
        /// maximal replica number.
        /// </summary>
        public readonly int MaxReplicas;
        /// <summary>
        /// minimal replica number.
        /// </summary>
        public readonly int MinReplicas;

        [OutputConstructor]
        private ScaleRuleAutoscaler(
            string autoscalerName,

            ImmutableArray<Outputs.ScaleRuleAutoscalerCronHorizontalAutoscaler> cronHorizontalAutoscalers,

            string? description,

            bool enabled,

            ImmutableArray<Outputs.ScaleRuleAutoscalerHorizontalAutoscaler> horizontalAutoscalers,

            int maxReplicas,

            int minReplicas)
        {
            AutoscalerName = autoscalerName;
            CronHorizontalAutoscalers = cronHorizontalAutoscalers;
            Description = description;
            Enabled = enabled;
            HorizontalAutoscalers = horizontalAutoscalers;
            MaxReplicas = maxReplicas;
            MinReplicas = minReplicas;
        }
    }
}
