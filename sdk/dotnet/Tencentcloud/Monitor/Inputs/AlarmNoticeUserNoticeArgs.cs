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

    public sealed class AlarmNoticeUserNoticeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds since the notification end time 00:00:00 (value range 0-86399).
        /// </summary>
        [Input("endTime", required: true)]
        public Input<int> EndTime { get; set; } = null!;

        [Input("groupIds")]
        private InputList<int>? _groupIds;

        /// <summary>
        /// User group ID list.
        /// </summary>
        public InputList<int> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<int>());
            set => _groupIds = value;
        }

        /// <summary>
        /// Contact notification required 0= No 1= Yes.
        /// </summary>
        [Input("needPhoneArriveNotice")]
        public Input<int>? NeedPhoneArriveNotice { get; set; }

        [Input("noticeWays", required: true)]
        private InputList<string>? _noticeWays;

        /// <summary>
        /// Notification Channel List EMAIL=Mail SMS=SMS CALL=Telephone WECHAT=WeChat RTX=Enterprise WeChat.
        /// </summary>
        public InputList<string> NoticeWays
        {
            get => _noticeWays ?? (_noticeWays = new InputList<string>());
            set => _noticeWays = value;
        }

        /// <summary>
        /// Call type SYNC= Simultaneous call CIRCLE= Round call If this parameter is not specified, the default value is round call.
        /// </summary>
        [Input("phoneCallType")]
        public Input<string>? PhoneCallType { get; set; }

        /// <summary>
        /// Number of seconds between polls (value range: 60-900).
        /// </summary>
        [Input("phoneCircleInterval")]
        public Input<int>? PhoneCircleInterval { get; set; }

        /// <summary>
        /// Number of telephone polls (value range: 1-5).
        /// </summary>
        [Input("phoneCircleTimes")]
        public Input<int>? PhoneCircleTimes { get; set; }

        /// <summary>
        /// Number of seconds between calls in a polling session (value range: 60-900).
        /// </summary>
        [Input("phoneInnerInterval")]
        public Input<int>? PhoneInnerInterval { get; set; }

        [Input("phoneOrders")]
        private InputList<int>? _phoneOrders;

        /// <summary>
        /// Telephone polling list.
        /// </summary>
        public InputList<int> PhoneOrders
        {
            get => _phoneOrders ?? (_phoneOrders = new InputList<int>());
            set => _phoneOrders = value;
        }

        /// <summary>
        /// Recipient Type USER=User GROUP=User Group.
        /// </summary>
        [Input("receiverType", required: true)]
        public Input<string> ReceiverType { get; set; } = null!;

        /// <summary>
        /// The number of seconds since the notification start time 00:00:00 (value range 0-86399).
        /// </summary>
        [Input("startTime", required: true)]
        public Input<int> StartTime { get; set; } = null!;

        [Input("userIds")]
        private InputList<int>? _userIds;

        /// <summary>
        /// User UID List.
        /// </summary>
        public InputList<int> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<int>());
            set => _userIds = value;
        }

        [Input("weekdays")]
        private InputList<int>? _weekdays;

        /// <summary>
        /// Notification period 1-7 indicates Monday to Sunday.
        /// </summary>
        public InputList<int> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<int>());
            set => _weekdays = value;
        }

        public AlarmNoticeUserNoticeArgs()
        {
        }
        public static new AlarmNoticeUserNoticeArgs Empty => new AlarmNoticeUserNoticeArgs();
    }
}
