// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cbs.Outputs
{

    [OutputType]
    public sealed class GetStoragesSetStorageListResult
    {
        /// <summary>
        /// Indicates whether the CBS is mounted the CVM.
        /// </summary>
        public readonly bool Attached;
        /// <summary>
        /// The available zone that the CBS instance locates at.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// Creation time of CBS.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Indicates whether CBS is encrypted.
        /// </summary>
        public readonly bool Encrypt;
        /// <summary>
        /// ID of the CVM instance that be mounted by this CBS.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
        /// </summary>
        public readonly string PrepaidRenewFlag;
        /// <summary>
        /// ID of the project with which the CBS is associated.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Status of CBS.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// ID of the CBS to be queried.
        /// </summary>
        public readonly string StorageId;
        /// <summary>
        /// Name of the CBS to be queried.
        /// </summary>
        public readonly string StorageName;
        /// <summary>
        /// Volume of CBS.
        /// </summary>
        public readonly int StorageSize;
        /// <summary>
        /// Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
        /// </summary>
        public readonly string StorageType;
        /// <summary>
        /// Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
        /// </summary>
        public readonly string StorageUsage;
        /// <summary>
        /// The available tags within this CBS.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        /// </summary>
        public readonly int ThroughputPerformance;

        [OutputConstructor]
        private GetStoragesSetStorageListResult(
            bool attached,

            string availabilityZone,

            string chargeType,

            string createTime,

            bool encrypt,

            string instanceId,

            string prepaidRenewFlag,

            int projectId,

            string status,

            string storageId,

            string storageName,

            int storageSize,

            string storageType,

            string storageUsage,

            ImmutableDictionary<string, object> tags,

            int throughputPerformance)
        {
            Attached = attached;
            AvailabilityZone = availabilityZone;
            ChargeType = chargeType;
            CreateTime = createTime;
            Encrypt = encrypt;
            InstanceId = instanceId;
            PrepaidRenewFlag = prepaidRenewFlag;
            ProjectId = projectId;
            Status = status;
            StorageId = storageId;
            StorageName = storageName;
            StorageSize = storageSize;
            StorageType = storageType;
            StorageUsage = storageUsage;
            Tags = tags;
            ThroughputPerformance = throughputPerformance;
        }
    }
}
