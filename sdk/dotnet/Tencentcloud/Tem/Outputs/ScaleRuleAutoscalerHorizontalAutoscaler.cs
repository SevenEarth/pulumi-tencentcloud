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
    public sealed class ScaleRuleAutoscalerHorizontalAutoscaler
    {
        /// <summary>
        /// enable scaler.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// maximal replica number.
        /// </summary>
        public readonly int MaxReplicas;
        /// <summary>
        /// metric name.
        /// </summary>
        public readonly string Metrics;
        /// <summary>
        /// minimal replica number.
        /// </summary>
        public readonly int MinReplicas;
        /// <summary>
        /// metric threshold.
        /// </summary>
        public readonly int Threshold;

        [OutputConstructor]
        private ScaleRuleAutoscalerHorizontalAutoscaler(
            bool enabled,

            int maxReplicas,

            string metrics,

            int minReplicas,

            int threshold)
        {
            Enabled = enabled;
            MaxReplicas = maxReplicas;
            Metrics = metrics;
            MinReplicas = minReplicas;
            Threshold = threshold;
        }
    }
}
