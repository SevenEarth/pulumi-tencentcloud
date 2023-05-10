// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class WorkflowTaskNotifyConfig
    {
        /// <summary>
        /// CMQ or TDMQ-CMQ model, there are two kinds of Queue and Topic.
        /// </summary>
        public readonly string? CmqModel;
        /// <summary>
        /// Region of CMQ or TDMQ-CMQ, such as sh, bj, etc.
        /// </summary>
        public readonly string? CmqRegion;
        /// <summary>
        /// The mode of the workflow notification, the possible values are Finish and Change, leaving blank means Finish.
        /// </summary>
        public readonly string? NotifyMode;
        /// <summary>
        /// Notification type, optional value:CMQ: offline, it is recommended to switch to TDMQ-CMQ.TDMQ-CMQ: message queue.URL: When the URL is specified, the HTTP callback is pushed to the address specified by NotifyUrl, the callback protocol is http+json, and the package body content is the same as the output parameters of the parsing event notification interface.SCF: not recommended, additional configuration of SCF in the console is required.Note: CMQ is the default when not filled or empty, if you need to use other types, you need to fill in the corresponding type value.
        /// </summary>
        public readonly string? NotifyType;
        /// <summary>
        /// HTTP callback address, required when NotifyType is URL.
        /// </summary>
        public readonly string? NotifyUrl;
        /// <summary>
        /// Valid when the model is Queue, indicating the queue name of the CMQ or TDMQ-CMQ that receives the event notification.
        /// </summary>
        public readonly string? QueueName;
        /// <summary>
        /// Valid when the model is a Topic, indicating the topic name of the CMQ or TDMQ-CMQ that receives event notifications.
        /// </summary>
        public readonly string? TopicName;

        [OutputConstructor]
        private WorkflowTaskNotifyConfig(
            string? cmqModel,

            string? cmqRegion,

            string? notifyMode,

            string? notifyType,

            string? notifyUrl,

            string? queueName,

            string? topicName)
        {
            CmqModel = cmqModel;
            CmqRegion = cmqRegion;
            NotifyMode = notifyMode;
            NotifyType = notifyType;
            NotifyUrl = notifyUrl;
            QueueName = queueName;
            TopicName = topicName;
        }
    }
}