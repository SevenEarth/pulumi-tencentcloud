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

    public sealed class WorkflowTaskNotifyConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CMQ or TDMQ-CMQ model, there are two kinds of Queue and Topic.
        /// </summary>
        [Input("cmqModel")]
        public Input<string>? CmqModel { get; set; }

        /// <summary>
        /// Region of CMQ or TDMQ-CMQ, such as sh, bj, etc.
        /// </summary>
        [Input("cmqRegion")]
        public Input<string>? CmqRegion { get; set; }

        /// <summary>
        /// The mode of the workflow notification, the possible values are Finish and Change, leaving blank means Finish.
        /// </summary>
        [Input("notifyMode")]
        public Input<string>? NotifyMode { get; set; }

        /// <summary>
        /// Notification type, optional value:CMQ: offline, it is recommended to switch to TDMQ-CMQ.TDMQ-CMQ: message queue.URL: When the URL is specified, the HTTP callback is pushed to the address specified by NotifyUrl, the callback protocol is http+json, and the package body content is the same as the output parameters of the parsing event notification interface.SCF: not recommended, additional configuration of SCF in the console is required.Note: CMQ is the default when not filled or empty, if you need to use other types, you need to fill in the corresponding type value.
        /// </summary>
        [Input("notifyType")]
        public Input<string>? NotifyType { get; set; }

        /// <summary>
        /// HTTP callback address, required when NotifyType is URL.
        /// </summary>
        [Input("notifyUrl")]
        public Input<string>? NotifyUrl { get; set; }

        /// <summary>
        /// Valid when the model is Queue, indicating the queue name of the CMQ or TDMQ-CMQ that receives the event notification.
        /// </summary>
        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        /// <summary>
        /// Valid when the model is a Topic, indicating the topic name of the CMQ or TDMQ-CMQ that receives event notifications.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public WorkflowTaskNotifyConfigArgs()
        {
        }
    }
}
