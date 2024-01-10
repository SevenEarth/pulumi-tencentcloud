// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod.Outputs
{

    [OutputType]
    public sealed class GetDomainListDomainListResult
    {
        /// <summary>
        /// Whether to enable CNAME acceleration, enabled: ENABLE, disabled: DISABLE.
        /// </summary>
        public readonly string CnameSpeedup;
        /// <summary>
        /// Domain addition time.
        /// </summary>
        public readonly string CreatedOn;
        /// <summary>
        /// DNS settings status, error: DNSERROR, normal: empty string.
        /// </summary>
        public readonly string DnsStatus;
        /// <summary>
        /// Unique identifier assigned to the domain by the system.
        /// </summary>
        public readonly int DomainId;
        /// <summary>
        /// Valid DNS assigned to the domain by the system.
        /// </summary>
        public readonly ImmutableArray<string> EffectiveDns;
        /// <summary>
        /// Domain package level code.
        /// </summary>
        public readonly string Grade;
        /// <summary>
        /// Sequence number corresponding to the domain package level.
        /// </summary>
        public readonly int GradeLevel;
        /// <summary>
        /// Package name.
        /// </summary>
        public readonly string GradeTitle;
        /// <summary>
        /// Get domain names based on domain group id, which can be obtained through the GroupId field in DescribeDomain or DescribeDomainList interface.
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// Whether it is a paid package.
        /// </summary>
        public readonly string IsVip;
        /// <summary>
        /// Original format of the domain.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Domain owner account.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// Punycode encoded domain format.
        /// </summary>
        public readonly string Punycode;
        /// <summary>
        /// Number of records under the domain.
        /// </summary>
        public readonly int RecordCount;
        /// <summary>
        /// Get domain names based on remark information.
        /// </summary>
        public readonly string Remark;
        /// <summary>
        /// Whether to enable search engine push optimization, YES: YES, NO: NO.
        /// </summary>
        public readonly string SearchEnginePush;
        /// <summary>
        /// Get domain names based on domain status. Available values are ENABLE, LOCK, PAUSE, SPAM. ENABLE: Normal LOCK: Locked PAUSE: Paused SPAM: Banned.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Domain-related tag list Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainListDomainListTagListResult> TagLists;
        /// <summary>
        /// Default TTL value for domain resolution records.
        /// </summary>
        public readonly int Ttl;
        /// <summary>
        /// Domain update time.
        /// </summary>
        public readonly string UpdatedOn;
        /// <summary>
        /// Whether the domain has VIP auto-renewal enabled, YES: YES, NO: NO, DEFAULT: DEFAULT.
        /// </summary>
        public readonly string VipAutoRenew;
        /// <summary>
        /// Paid package expiration time.
        /// </summary>
        public readonly string VipEndAt;
        /// <summary>
        /// Paid package activation time.
        /// </summary>
        public readonly string VipStartAt;

        [OutputConstructor]
        private GetDomainListDomainListResult(
            string cnameSpeedup,

            string createdOn,

            string dnsStatus,

            int domainId,

            ImmutableArray<string> effectiveDns,

            string grade,

            int gradeLevel,

            string gradeTitle,

            int groupId,

            string isVip,

            string name,

            string owner,

            string punycode,

            int recordCount,

            string remark,

            string searchEnginePush,

            string status,

            ImmutableArray<Outputs.GetDomainListDomainListTagListResult> tagLists,

            int ttl,

            string updatedOn,

            string vipAutoRenew,

            string vipEndAt,

            string vipStartAt)
        {
            CnameSpeedup = cnameSpeedup;
            CreatedOn = createdOn;
            DnsStatus = dnsStatus;
            DomainId = domainId;
            EffectiveDns = effectiveDns;
            Grade = grade;
            GradeLevel = gradeLevel;
            GradeTitle = gradeTitle;
            GroupId = groupId;
            IsVip = isVip;
            Name = name;
            Owner = owner;
            Punycode = punycode;
            RecordCount = recordCount;
            Remark = remark;
            SearchEnginePush = searchEnginePush;
            Status = status;
            TagLists = tagLists;
            Ttl = ttl;
            UpdatedOn = updatedOn;
            VipAutoRenew = vipAutoRenew;
            VipEndAt = vipEndAt;
            VipStartAt = vipStartAt;
        }
    }
}
