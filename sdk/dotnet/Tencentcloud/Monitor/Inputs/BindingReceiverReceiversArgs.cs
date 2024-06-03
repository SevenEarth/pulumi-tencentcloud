// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class BindingReceiverReceiversArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of alarm period. Meaning with `start_time`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("notifyWays", required: true)]
        private InputList<string>? _notifyWays;

        /// <summary>
        /// Method of warning notification.Optional `CALL`,`EMAIL`,`SITE`,`SMS`,`WECHAT`.
        /// </summary>
        public InputList<string> NotifyWays
        {
            get => _notifyWays ?? (_notifyWays = new InputList<string>());
            set => _notifyWays = value;
        }

        /// <summary>
        /// Alert sending language. Optional `en-US`,`zh-CN`.
        /// </summary>
        [Input("receiveLanguage")]
        public Input<string>? ReceiveLanguage { get; set; }

        [Input("receiverGroupLists")]
        private InputList<int>? _receiverGroupLists;

        /// <summary>
        /// Alarm receive group ID list.
        /// </summary>
        public InputList<int> ReceiverGroupLists
        {
            get => _receiverGroupLists ?? (_receiverGroupLists = new InputList<int>());
            set => _receiverGroupLists = value;
        }

        /// <summary>
        /// Receive type. Optional `group`,`user`.
        /// </summary>
        [Input("receiverType", required: true)]
        public Input<string> ReceiverType { get; set; } = null!;

        [Input("receiverUserLists")]
        private InputList<int>? _receiverUserLists;

        /// <summary>
        /// Alarm receiver ID list.
        /// </summary>
        public InputList<int> ReceiverUserLists
        {
            get => _receiverUserLists ?? (_receiverUserLists = new InputList<int>());
            set => _receiverUserLists = value;
        }

        /// <summary>
        /// Alarm period start time. Valid value ranges: (0~86399). which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        public BindingReceiverReceiversArgs()
        {
        }
        public static new BindingReceiverReceiversArgs Empty => new BindingReceiverReceiversArgs();
    }
}
