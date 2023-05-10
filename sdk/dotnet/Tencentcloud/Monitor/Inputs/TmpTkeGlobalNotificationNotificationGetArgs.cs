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

    public sealed class TmpTkeGlobalNotificationNotificationGetArgs : Pulumi.ResourceArgs
    {
        [Input("alertManagers")]
        private InputList<Inputs.TmpTkeGlobalNotificationNotificationAlertManagerGetArgs>? _alertManagers;

        /// <summary>
        /// Alert manager, if Type is `alertmanager`, this field is required.
        /// </summary>
        public InputList<Inputs.TmpTkeGlobalNotificationNotificationAlertManagerGetArgs> AlertManagers
        {
            get => _alertManagers ?? (_alertManagers = new InputList<Inputs.TmpTkeGlobalNotificationNotificationAlertManagerGetArgs>());
            set => _alertManagers = value;
        }

        /// <summary>
        /// Alarm notification switch.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("notifyWays")]
        private InputList<string>? _notifyWays;

        /// <summary>
        /// Alarm notification method, Valid values: `SMS`, `EMAIL`, `CALL`, `WECHAT`.
        /// </summary>
        public InputList<string> NotifyWays
        {
            get => _notifyWays ?? (_notifyWays = new InputList<string>());
            set => _notifyWays = value;
        }

        /// <summary>
        /// Phone Alarm Reach Notification, NotifyWay is `CALL`, and this parameter is used.
        /// </summary>
        [Input("phoneArriveNotice")]
        public Input<bool>? PhoneArriveNotice { get; set; }

        /// <summary>
        /// Telephone alarm off-wheel interval, NotifyWay is `CALL`, and this parameter is used.
        /// </summary>
        [Input("phoneCircleInterval")]
        public Input<int>? PhoneCircleInterval { get; set; }

        /// <summary>
        /// Number of phone alerts (user group), NotifyWay is `CALL`, and this parameter is used.
        /// </summary>
        [Input("phoneCircleTimes")]
        public Input<int>? PhoneCircleTimes { get; set; }

        /// <summary>
        /// Interval between telephone alarm rounds, NotifyWay is `CALL`, and this parameter is used.
        /// </summary>
        [Input("phoneInnerInterval")]
        public Input<int>? PhoneInnerInterval { get; set; }

        [Input("phoneNotifyOrders")]
        private InputList<int>? _phoneNotifyOrders;

        /// <summary>
        /// Phone alert sequence, NotifyWay is `CALL`, and this parameter is used.
        /// </summary>
        public InputList<int> PhoneNotifyOrders
        {
            get => _phoneNotifyOrders ?? (_phoneNotifyOrders = new InputList<int>());
            set => _phoneNotifyOrders = value;
        }

        [Input("receiverGroups")]
        private InputList<string>? _receiverGroups;

        /// <summary>
        /// Alarm receiving group(user group).
        /// </summary>
        public InputList<string> ReceiverGroups
        {
            get => _receiverGroups ?? (_receiverGroups = new InputList<string>());
            set => _receiverGroups = value;
        }

        /// <summary>
        /// Convergence time.
        /// </summary>
        [Input("repeatInterval")]
        public Input<string>? RepeatInterval { get; set; }

        /// <summary>
        /// Effective end time.
        /// </summary>
        [Input("timeRangeEnd")]
        public Input<string>? TimeRangeEnd { get; set; }

        /// <summary>
        /// Effective start time.
        /// </summary>
        [Input("timeRangeStart")]
        public Input<string>? TimeRangeStart { get; set; }

        /// <summary>
        /// Alarm notification type, Valid values: `amp`, `webhook`, `alertmanager`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Web hook, if Type is `webhook`, this field is required.
        /// </summary>
        [Input("webHook")]
        public Input<string>? WebHook { get; set; }

        public TmpTkeGlobalNotificationNotificationGetArgs()
        {
        }
    }
}
