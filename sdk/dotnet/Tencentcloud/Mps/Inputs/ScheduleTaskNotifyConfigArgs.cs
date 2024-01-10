// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleTaskNotifyConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS SQS queue. This parameter is required if `NotifyType` is `AWS-SQS`.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("awsSqs")]
        public Input<Inputs.ScheduleTaskNotifyConfigAwsSqsArgs>? AwsSqs { get; set; }

        /// <summary>
        /// The CMQ or TDMQ-CMQ model. Valid values: Queue, Topic.
        /// </summary>
        [Input("cmqModel")]
        public Input<string>? CmqModel { get; set; }

        /// <summary>
        /// The CMQ or TDMQ-CMQ region, such as `sh` (Shanghai) or `bj` (Beijing).
        /// </summary>
        [Input("cmqRegion")]
        public Input<string>? CmqRegion { get; set; }

        /// <summary>
        /// Workflow notification method. Valid values: Finish, Change. If this parameter is left empty, `Finish` will be used.
        /// </summary>
        [Input("notifyMode")]
        public Input<string>? NotifyMode { get; set; }

        /// <summary>
        /// The notification type. Valid values: `CMQ`: This value is no longer used. Please use `TDMQ-CMQ` instead. `TDMQ-CMQ`: Message queue `URL`: If `NotifyType` is set to `URL`, HTTP callbacks are sent to the URL specified by `NotifyUrl`. HTTP and JSON are used for the callbacks. The packet contains the response parameters of the `ParseNotification` API. `SCF`: This notification type is not recommended. You need to configure it in the SCF console. `AWS-SQS`: AWS queue. This type is only supported for AWS tasks, and the queue must be in the same region as the AWS bucket.Note: If you do not pass this parameter or pass in an empty string, `CMQ` will be used. To use a different notification type, specify this parameter accordingly.
        /// </summary>
        [Input("notifyType")]
        public Input<string>? NotifyType { get; set; }

        /// <summary>
        /// HTTP callback URL, required if `NotifyType` is set to `URL`.
        /// </summary>
        [Input("notifyUrl")]
        public Input<string>? NotifyUrl { get; set; }

        /// <summary>
        /// The CMQ or TDMQ-CMQ queue to receive notifications. This parameter is valid when `CmqModel` is `Queue`.
        /// </summary>
        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        /// <summary>
        /// The CMQ or TDMQ-CMQ topic to receive notifications. This parameter is valid when `CmqModel` is `Topic`.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public ScheduleTaskNotifyConfigArgs()
        {
        }
    }
}
