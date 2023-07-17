// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTargetResourceEsParamDropDlqGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// dlq type, CKAFKA|TOPIC.
        /// </summary>
        [Input("dlqType")]
        public Input<string>? DlqType { get; set; }

        /// <summary>
        /// Ckafka type dlq.
        /// </summary>
        [Input("kafkaParam")]
        public Input<Inputs.DatahubTaskTargetResourceEsParamDropDlqKafkaParamGetArgs>? KafkaParam { get; set; }

        /// <summary>
        /// retry times.
        /// </summary>
        [Input("maxRetryAttempts")]
        public Input<int>? MaxRetryAttempts { get; set; }

        /// <summary>
        /// retry interval.
        /// </summary>
        [Input("retryInterval")]
        public Input<int>? RetryInterval { get; set; }

        /// <summary>
        /// DIP Topic type dead letter queue.
        /// </summary>
        [Input("topicParam")]
        public Input<Inputs.DatahubTaskTargetResourceEsParamDropDlqTopicParamGetArgs>? TopicParam { get; set; }

        /// <summary>
        /// type, DLQ dead letter queue, IGNORE_ERROR|DROP.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatahubTaskTargetResourceEsParamDropDlqGetArgs()
        {
        }
    }
}
