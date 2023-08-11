// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Outputs
{

    [OutputType]
    public sealed class GetResourcePackageListResourcePackageListResult
    {
        /// <summary>
        /// AppID note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int AppId;
        /// <summary>
        /// Note for binding instance information: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourcePackageListResourcePackageListBindInstanceInfoResult> BindInstanceInfos;
        /// <summary>
        /// Expiration time: August 1st, 2022 00:00:00 Attention: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// Resource package usage note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly bool HasQuota;
        /// <summary>
        /// Resource Package Unique ID.
        /// </summary>
        public readonly string PackageId;
        /// <summary>
        /// Resource Package Name.
        /// </summary>
        public readonly string PackageName;
        /// <summary>
        /// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
        /// </summary>
        public readonly string PackageRegion;
        /// <summary>
        /// Attention to the total amount of resource packages: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly double PackageTotalSpec;
        /// <summary>
        /// Resource package type CCU - Compute resource package, DISK - Storage resource package.
        /// </summary>
        public readonly string PackageType;
        /// <summary>
        /// Resource package usage note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly double PackageUsedSpec;
        /// <summary>
        /// Effective time: July 1st, 2022 00:00:00 Attention: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Resource package status creating - creating; Using - In use; Expired - has expired; Normal_ Finish - used up; Apply_ Refund - Applying for a refund; Refund - The fee has been refunded.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetResourcePackageListResourcePackageListResult(
            int appId,

            ImmutableArray<Outputs.GetResourcePackageListResourcePackageListBindInstanceInfoResult> bindInstanceInfos,

            string expireTime,

            bool hasQuota,

            string packageId,

            string packageName,

            string packageRegion,

            double packageTotalSpec,

            string packageType,

            double packageUsedSpec,

            string startTime,

            string status)
        {
            AppId = appId;
            BindInstanceInfos = bindInstanceInfos;
            ExpireTime = expireTime;
            HasQuota = hasQuota;
            PackageId = packageId;
            PackageName = packageName;
            PackageRegion = packageRegion;
            PackageTotalSpec = packageTotalSpec;
            PackageType = packageType;
            PackageUsedSpec = packageUsedSpec;
            StartTime = startTime;
            Status = status;
        }
    }
}