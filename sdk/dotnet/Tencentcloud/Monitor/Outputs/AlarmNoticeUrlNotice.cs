// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class AlarmNoticeUrlNotice
    {
        /// <summary>
        /// Notification End Time Seconds at the start of a day.
        /// </summary>
        public readonly int? EndTime;
        /// <summary>
        /// Notification Start Time Number of seconds at the start of a day.
        /// </summary>
        public readonly int? StartTime;
        /// <summary>
        /// Callback URL (limited to 256 characters).
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Notification period 1-7 indicates Monday to Sunday.
        /// </summary>
        public readonly ImmutableArray<int> Weekdays;

        [OutputConstructor]
        private AlarmNoticeUrlNotice(
            int? endTime,

            int? startTime,

            string url,

            ImmutableArray<int> weekdays)
        {
            EndTime = endTime;
            StartTime = startTime;
            Url = url;
            Weekdays = weekdays;
        }
    }
}
