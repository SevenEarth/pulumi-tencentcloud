# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSnapshotPoliciesSnapshotPolicyListResult',
    'GetSnapshotsSnapshotListResult',
    'GetStoragesSetStorageListResult',
    'GetStoragesStorageListResult',
]

@pulumi.output_type
class GetSnapshotPoliciesSnapshotPolicyListResult(dict):
    def __init__(__self__, *,
                 attached_storage_ids: Sequence[str],
                 create_time: str,
                 repeat_hours: Sequence[int],
                 repeat_weekdays: Sequence[int],
                 retention_days: int,
                 snapshot_policy_id: str,
                 snapshot_policy_name: str,
                 status: str):
        """
        :param Sequence[str] attached_storage_ids: Storage IDs that the snapshot policy attached.
        :param str create_time: Create time of the snapshot policy.
        :param Sequence[int] repeat_hours: Trigger hours of periodic snapshot.
        :param Sequence[int] repeat_weekdays: Trigger days of periodic snapshot.
        :param int retention_days: Retention days of the snapshot.
        :param str snapshot_policy_id: ID of the snapshot policy to be queried.
        :param str snapshot_policy_name: Name of the snapshot policy to be queried.
        :param str status: Status of the snapshot policy.
        """
        pulumi.set(__self__, "attached_storage_ids", attached_storage_ids)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "repeat_hours", repeat_hours)
        pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        pulumi.set(__self__, "retention_days", retention_days)
        pulumi.set(__self__, "snapshot_policy_id", snapshot_policy_id)
        pulumi.set(__self__, "snapshot_policy_name", snapshot_policy_name)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="attachedStorageIds")
    def attached_storage_ids(self) -> Sequence[str]:
        """
        Storage IDs that the snapshot policy attached.
        """
        return pulumi.get(self, "attached_storage_ids")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Create time of the snapshot policy.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="repeatHours")
    def repeat_hours(self) -> Sequence[int]:
        """
        Trigger hours of periodic snapshot.
        """
        return pulumi.get(self, "repeat_hours")

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> Sequence[int]:
        """
        Trigger days of periodic snapshot.
        """
        return pulumi.get(self, "repeat_weekdays")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> int:
        """
        Retention days of the snapshot.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> str:
        """
        ID of the snapshot policy to be queried.
        """
        return pulumi.get(self, "snapshot_policy_id")

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> str:
        """
        Name of the snapshot policy to be queried.
        """
        return pulumi.get(self, "snapshot_policy_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the snapshot policy.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetSnapshotsSnapshotListResult(dict):
    def __init__(__self__, *,
                 availability_zone: str,
                 create_time: str,
                 encrypt: bool,
                 percent: int,
                 project_id: int,
                 snapshot_id: str,
                 snapshot_name: str,
                 storage_id: str,
                 storage_size: int,
                 storage_usage: str):
        """
        :param str availability_zone: The available zone that the CBS instance locates at.
        :param str create_time: Creation time of snapshot.
        :param bool encrypt: Indicates whether the snapshot is encrypted.
        :param int percent: Snapshot creation progress percentage.
        :param int project_id: ID of the project within the snapshot.
        :param str snapshot_id: ID of the snapshot to be queried.
        :param str snapshot_name: Name of the snapshot to be queried.
        :param str storage_id: ID of the the CBS which this snapshot created from.
        :param int storage_size: Volume of storage which this snapshot created from.
        :param str storage_usage: Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "encrypt", encrypt)
        pulumi.set(__self__, "percent", percent)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        pulumi.set(__self__, "snapshot_name", snapshot_name)
        pulumi.set(__self__, "storage_id", storage_id)
        pulumi.set(__self__, "storage_size", storage_size)
        pulumi.set(__self__, "storage_usage", storage_usage)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The available zone that the CBS instance locates at.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of snapshot.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def encrypt(self) -> bool:
        """
        Indicates whether the snapshot is encrypted.
        """
        return pulumi.get(self, "encrypt")

    @property
    @pulumi.getter
    def percent(self) -> int:
        """
        Snapshot creation progress percentage.
        """
        return pulumi.get(self, "percent")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        """
        ID of the project within the snapshot.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> str:
        """
        ID of the snapshot to be queried.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> str:
        """
        Name of the snapshot to be queried.
        """
        return pulumi.get(self, "snapshot_name")

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> str:
        """
        ID of the the CBS which this snapshot created from.
        """
        return pulumi.get(self, "storage_id")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> int:
        """
        Volume of storage which this snapshot created from.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="storageUsage")
    def storage_usage(self) -> str:
        """
        Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.
        """
        return pulumi.get(self, "storage_usage")


@pulumi.output_type
class GetStoragesSetStorageListResult(dict):
    def __init__(__self__, *,
                 attached: bool,
                 availability_zone: str,
                 charge_type: str,
                 create_time: str,
                 encrypt: bool,
                 instance_id: str,
                 prepaid_renew_flag: str,
                 project_id: int,
                 status: str,
                 storage_id: str,
                 storage_name: str,
                 storage_size: int,
                 storage_type: str,
                 storage_usage: str,
                 tags: Mapping[str, Any],
                 throughput_performance: int):
        """
        :param bool attached: Indicates whether the CBS is mounted the CVM.
        :param str availability_zone: The available zone that the CBS instance locates at.
        :param str charge_type: List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
        :param str create_time: Creation time of CBS.
        :param bool encrypt: Indicates whether CBS is encrypted.
        :param str instance_id: ID of the CVM instance that be mounted by this CBS.
        :param str prepaid_renew_flag: The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
        :param int project_id: ID of the project with which the CBS is associated.
        :param str status: Status of CBS.
        :param str storage_id: ID of the CBS to be queried.
        :param str storage_name: Name of the CBS to be queried.
        :param int storage_size: Volume of CBS.
        :param str storage_type: Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
        :param str storage_usage: Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
        :param Mapping[str, Any] tags: The available tags within this CBS.
        :param int throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        pulumi.set(__self__, "attached", attached)
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "encrypt", encrypt)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "prepaid_renew_flag", prepaid_renew_flag)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage_id", storage_id)
        pulumi.set(__self__, "storage_name", storage_name)
        pulumi.set(__self__, "storage_size", storage_size)
        pulumi.set(__self__, "storage_type", storage_type)
        pulumi.set(__self__, "storage_usage", storage_usage)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "throughput_performance", throughput_performance)

    @property
    @pulumi.getter
    def attached(self) -> bool:
        """
        Indicates whether the CBS is mounted the CVM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The available zone that the CBS instance locates at.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of CBS.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def encrypt(self) -> bool:
        """
        Indicates whether CBS is encrypted.
        """
        return pulumi.get(self, "encrypt")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the CVM instance that be mounted by this CBS.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="prepaidRenewFlag")
    def prepaid_renew_flag(self) -> str:
        """
        The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
        """
        return pulumi.get(self, "prepaid_renew_flag")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        """
        ID of the project with which the CBS is associated.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of CBS.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> str:
        """
        ID of the CBS to be queried.
        """
        return pulumi.get(self, "storage_id")

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> str:
        """
        Name of the CBS to be queried.
        """
        return pulumi.get(self, "storage_name")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> int:
        """
        Volume of CBS.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> str:
        """
        Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="storageUsage")
    def storage_usage(self) -> str:
        """
        Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
        """
        return pulumi.get(self, "storage_usage")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        """
        The available tags within this CBS.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="throughputPerformance")
    def throughput_performance(self) -> int:
        """
        Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        return pulumi.get(self, "throughput_performance")


@pulumi.output_type
class GetStoragesStorageListResult(dict):
    def __init__(__self__, *,
                 attached: bool,
                 availability_zone: str,
                 charge_type: str,
                 create_time: str,
                 encrypt: bool,
                 instance_id: str,
                 prepaid_renew_flag: str,
                 project_id: int,
                 status: str,
                 storage_id: str,
                 storage_name: str,
                 storage_size: int,
                 storage_type: str,
                 storage_usage: str,
                 tags: Mapping[str, Any],
                 throughput_performance: int):
        """
        :param bool attached: Indicates whether the CBS is mounted the CVM.
        :param str availability_zone: The available zone that the CBS instance locates at.
        :param str charge_type: List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
        :param str create_time: Creation time of CBS.
        :param bool encrypt: Indicates whether CBS is encrypted.
        :param str instance_id: ID of the CVM instance that be mounted by this CBS.
        :param str prepaid_renew_flag: The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
        :param int project_id: ID of the project with which the CBS is associated.
        :param str status: Status of CBS.
        :param str storage_id: ID of the CBS to be queried.
        :param str storage_name: Name of the CBS to be queried.
        :param int storage_size: Volume of CBS.
        :param str storage_type: Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
        :param str storage_usage: Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
        :param Mapping[str, Any] tags: The available tags within this CBS.
        :param int throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        pulumi.set(__self__, "attached", attached)
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "encrypt", encrypt)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "prepaid_renew_flag", prepaid_renew_flag)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage_id", storage_id)
        pulumi.set(__self__, "storage_name", storage_name)
        pulumi.set(__self__, "storage_size", storage_size)
        pulumi.set(__self__, "storage_type", storage_type)
        pulumi.set(__self__, "storage_usage", storage_usage)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "throughput_performance", throughput_performance)

    @property
    @pulumi.getter
    def attached(self) -> bool:
        """
        Indicates whether the CBS is mounted the CVM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The available zone that the CBS instance locates at.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of CBS.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def encrypt(self) -> bool:
        """
        Indicates whether CBS is encrypted.
        """
        return pulumi.get(self, "encrypt")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of the CVM instance that be mounted by this CBS.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="prepaidRenewFlag")
    def prepaid_renew_flag(self) -> str:
        """
        The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.
        """
        return pulumi.get(self, "prepaid_renew_flag")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        """
        ID of the project with which the CBS is associated.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of CBS.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> str:
        """
        ID of the CBS to be queried.
        """
        return pulumi.get(self, "storage_id")

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> str:
        """
        Name of the CBS to be queried.
        """
        return pulumi.get(self, "storage_name")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> int:
        """
        Volume of CBS.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> str:
        """
        Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="storageUsage")
    def storage_usage(self) -> str:
        """
        Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).
        """
        return pulumi.get(self, "storage_usage")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        """
        The available tags within this CBS.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="throughputPerformance")
    def throughput_performance(self) -> int:
        """
        Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        return pulumi.get(self, "throughput_performance")


