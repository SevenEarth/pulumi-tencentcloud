// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Outputs
{

    [OutputType]
    public sealed class GetEventRulesRuleDeadLetterConfigResult
    {
        /// <summary>
        /// After setting the DLQ mode, this option is required. The error message will be delivered to the corresponding kafka topic Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEventRulesRuleDeadLetterConfigCkafkaDeliveryParamResult> CkafkaDeliveryParams;
        /// <summary>
        /// Support three modes of dlq, discarding, ignoring errors and continuing to pass, corresponding to: DLQ, DROP, IGNORE_ERROR.
        /// </summary>
        public readonly string DisposeMethod;

        [OutputConstructor]
        private GetEventRulesRuleDeadLetterConfigResult(
            ImmutableArray<Outputs.GetEventRulesRuleDeadLetterConfigCkafkaDeliveryParamResult> ckafkaDeliveryParams,

            string disposeMethod)
        {
            CkafkaDeliveryParams = ckafkaDeliveryParams;
            DisposeMethod = disposeMethod;
        }
    }
}