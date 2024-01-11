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
    public sealed class GetEventRulesRuleDeadLetterConfigCkafkaDeliveryParamResult
    {
        /// <summary>
        /// ckafka resource qcs six-segment.
        /// </summary>
        public readonly string ResourceDescription;
        /// <summary>
        /// ckafka topic name.
        /// </summary>
        public readonly string TopicName;

        [OutputConstructor]
        private GetEventRulesRuleDeadLetterConfigCkafkaDeliveryParamResult(
            string resourceDescription,

            string topicName)
        {
            ResourceDescription = resourceDescription;
            TopicName = topicName;
        }
    }
}