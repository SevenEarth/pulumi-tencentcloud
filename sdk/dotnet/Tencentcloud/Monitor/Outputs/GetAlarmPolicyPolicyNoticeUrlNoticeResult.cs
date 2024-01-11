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
    public sealed class GetAlarmPolicyPolicyNoticeUrlNoticeResult
    {
        /// <summary>
        /// Notification end time, which is expressed by the number of seconds since 00:00:00. Value range: 0-86399Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int EndTime;
        /// <summary>
        /// Whether verification is passed. Valid values: 0 (no), 1 (yes)Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int IsValid;
        /// <summary>
        /// Notification start time, which is expressed by the number of seconds since 00:00:00. Value range: 0-86399Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int StartTime;
        /// <summary>
        /// Callback URL, which can contain up to 256 charactersNote: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Verification codeNote: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ValidationCode;
        /// <summary>
        /// Notification cycle. The values 1-7 indicate Monday to Sunday.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<int> Weekdays;

        [OutputConstructor]
        private GetAlarmPolicyPolicyNoticeUrlNoticeResult(
            int endTime,

            int isValid,

            int startTime,

            string url,

            string validationCode,

            ImmutableArray<int> weekdays)
        {
            EndTime = endTime;
            IsValid = isValid;
            StartTime = startTime;
            Url = url;
            ValidationCode = validationCode;
            Weekdays = weekdays;
        }
    }
}