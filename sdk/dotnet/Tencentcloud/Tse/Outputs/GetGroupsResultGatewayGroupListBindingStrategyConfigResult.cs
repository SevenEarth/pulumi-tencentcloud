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
    public sealed class GetGroupsResultGatewayGroupListBindingStrategyConfigResult
    {
        /// <summary>
        /// auto scaler IDNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string AutoScalerId;
        /// <summary>
        /// group create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// whether to enable timing auto scaling.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// maximum number of replicas.
        /// </summary>
        public readonly int MaxReplicas;
        /// <summary>
        /// metric listNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsResultGatewayGroupListBindingStrategyConfigMetricResult> Metrics;
        /// <summary>
        /// modify time.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// strategy ID.
        /// </summary>
        public readonly string StrategyId;

        [OutputConstructor]
        private GetGroupsResultGatewayGroupListBindingStrategyConfigResult(
            string autoScalerId,

            string createTime,

            bool enabled,

            int maxReplicas,

            ImmutableArray<Outputs.GetGroupsResultGatewayGroupListBindingStrategyConfigMetricResult> metrics,

            string modifyTime,

            string strategyId)
        {
            AutoScalerId = autoScalerId;
            CreateTime = createTime;
            Enabled = enabled;
            MaxReplicas = maxReplicas;
            Metrics = metrics;
            ModifyTime = modifyTime;
            StrategyId = strategyId;
        }
    }
}