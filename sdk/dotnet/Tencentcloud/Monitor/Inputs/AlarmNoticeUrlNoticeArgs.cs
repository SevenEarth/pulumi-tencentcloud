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

    public sealed class AlarmNoticeUrlNoticeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Notification End Time Seconds at the start of a day.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        /// <summary>
        /// If passed verification `0` is no, `1` is yes. Default `0`.
        /// </summary>
        [Input("isValid")]
        public Input<int>? IsValid { get; set; }

        /// <summary>
        /// Notification Start Time Number of seconds at the start of a day.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// Callback URL (limited to 256 characters).
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Verification code.
        /// </summary>
        [Input("validationCode")]
        public Input<string>? ValidationCode { get; set; }

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

        public AlarmNoticeUrlNoticeArgs()
        {
        }
        public static new AlarmNoticeUrlNoticeArgs Empty => new AlarmNoticeUrlNoticeArgs();
    }
}
