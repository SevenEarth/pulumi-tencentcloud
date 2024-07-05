// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Outputs
{

    [OutputType]
    public sealed class CngwStrategyConfig
    {
        /// <summary>
        /// behavior configuration of metric
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly Outputs.CngwStrategyConfigBehavior? Behavior;
        /// <summary>
        /// create time
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string? CreateTime;
        /// <summary>
        /// max number of replica for metric scaling.
        /// </summary>
        public readonly int? MaxReplicas;
        /// <summary>
        /// metric list.
        /// </summary>
        public readonly ImmutableArray<Outputs.CngwStrategyConfigMetric> Metrics;
        /// <summary>
        /// modify time
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string? ModifyTime;
        /// <summary>
        /// strategy ID
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string? StrategyId;

        [OutputConstructor]
        private CngwStrategyConfig(
            Outputs.CngwStrategyConfigBehavior? behavior,

            string? createTime,

            int? maxReplicas,

            ImmutableArray<Outputs.CngwStrategyConfigMetric> metrics,

            string? modifyTime,

            string? strategyId)
        {
            Behavior = behavior;
            CreateTime = createTime;
            MaxReplicas = maxReplicas;
            Metrics = metrics;
            ModifyTime = modifyTime;
            StrategyId = strategyId;
        }
    }
}