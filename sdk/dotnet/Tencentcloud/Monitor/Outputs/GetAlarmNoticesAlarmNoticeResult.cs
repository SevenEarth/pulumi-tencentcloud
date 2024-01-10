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
    public sealed class GetAlarmNoticesAlarmNoticeResult
    {
        /// <summary>
        /// AMP consumer ID.
        /// </summary>
        public readonly string AmpConsumerId;
        /// <summary>
        /// A maximum of one alarm notification can be pushed to the CLS service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmNoticesAlarmNoticeClsNoticeResult> ClsNotices;
        /// <summary>
        /// Alarm notification template ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether it is the system default notification template 0=No 1=Yes.
        /// </summary>
        public readonly int IsPreset;
        /// <summary>
        /// Alarm notification template name Used for fuzzy search.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notification language zh-CN=Chinese en-US=English.
        /// </summary>
        public readonly string NoticeLanguage;
        /// <summary>
        /// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        /// </summary>
        public readonly string NoticeType;
        /// <summary>
        /// List of alarm policy IDs bound to the alarm notification template.
        /// </summary>
        public readonly ImmutableArray<string> PolicyIds;
        /// <summary>
        /// Last modified time.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// Last Modified By.
        /// </summary>
        public readonly string UpdatedBy;
        /// <summary>
        /// The maximum number of callback notifications is 3.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmNoticesAlarmNoticeUrlNoticeResult> UrlNotices;
        /// <summary>
        /// Alarm notification template list.(At most five).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmNoticesAlarmNoticeUserNoticeResult> UserNotices;

        [OutputConstructor]
        private GetAlarmNoticesAlarmNoticeResult(
            string ampConsumerId,

            ImmutableArray<Outputs.GetAlarmNoticesAlarmNoticeClsNoticeResult> clsNotices,

            string id,

            int isPreset,

            string name,

            string noticeLanguage,

            string noticeType,

            ImmutableArray<string> policyIds,

            string updatedAt,

            string updatedBy,

            ImmutableArray<Outputs.GetAlarmNoticesAlarmNoticeUrlNoticeResult> urlNotices,

            ImmutableArray<Outputs.GetAlarmNoticesAlarmNoticeUserNoticeResult> userNotices)
        {
            AmpConsumerId = ampConsumerId;
            ClsNotices = clsNotices;
            Id = id;
            IsPreset = isPreset;
            Name = name;
            NoticeLanguage = noticeLanguage;
            NoticeType = noticeType;
            PolicyIds = policyIds;
            UpdatedAt = updatedAt;
            UpdatedBy = updatedBy;
            UrlNotices = urlNotices;
            UserNotices = userNotices;
        }
    }
}
