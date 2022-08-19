# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceDataDisk',
    'GetTypesFilterResult',
    'GetTypesInstanceTypeResult',
]

@pulumi.output_type
class InstanceDataDisk(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dataDiskSize":
            suggest = "data_disk_size"
        elif key == "dataDiskType":
            suggest = "data_disk_type"
        elif key == "dataDiskId":
            suggest = "data_disk_id"
        elif key == "dataDiskSnapshotId":
            suggest = "data_disk_snapshot_id"
        elif key == "deleteWithInstance":
            suggest = "delete_with_instance"
        elif key == "throughputPerformance":
            suggest = "throughput_performance"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceDataDisk. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceDataDisk.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceDataDisk.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 data_disk_size: int,
                 data_disk_type: str,
                 data_disk_id: Optional[str] = None,
                 data_disk_snapshot_id: Optional[str] = None,
                 delete_with_instance: Optional[bool] = None,
                 encrypt: Optional[bool] = None,
                 throughput_performance: Optional[int] = None):
        """
        :param int data_disk_size: Size of the data disk, and unit is GB. If disk type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
        :param str data_disk_type: Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_SSD`: SSD, `CLOUD_HSSD`: Enhanced SSD. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.
        :param str data_disk_id: Data disk ID used to initialize the data disk. When data disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
        :param str data_disk_snapshot_id: Snapshot ID of the data disk. The selected data disk snapshot size must be smaller than the data disk size.
        :param bool delete_with_instance: Decides whether the disk is deleted with instance(only applied to `CLOUD_BASIC`, `CLOUD_SSD` and `CLOUD_PREMIUM` disk with `POSTPAID_BY_HOUR` instance), default is true.
        :param bool encrypt: Decides whether the disk is encrypted. Default is `false`.
        :param int throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        pulumi.set(__self__, "data_disk_size", data_disk_size)
        pulumi.set(__self__, "data_disk_type", data_disk_type)
        if data_disk_id is not None:
            pulumi.set(__self__, "data_disk_id", data_disk_id)
        if data_disk_snapshot_id is not None:
            pulumi.set(__self__, "data_disk_snapshot_id", data_disk_snapshot_id)
        if delete_with_instance is not None:
            pulumi.set(__self__, "delete_with_instance", delete_with_instance)
        if encrypt is not None:
            pulumi.set(__self__, "encrypt", encrypt)
        if throughput_performance is not None:
            pulumi.set(__self__, "throughput_performance", throughput_performance)

    @property
    @pulumi.getter(name="dataDiskSize")
    def data_disk_size(self) -> int:
        """
        Size of the data disk, and unit is GB. If disk type is `CLOUD_SSD`, the size range is [100, 16000], and the others are [10-16000].
        """
        return pulumi.get(self, "data_disk_size")

    @property
    @pulumi.getter(name="dataDiskType")
    def data_disk_type(self) -> str:
        """
        Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_SSD`: SSD, `CLOUD_HSSD`: Enhanced SSD. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.
        """
        return pulumi.get(self, "data_disk_type")

    @property
    @pulumi.getter(name="dataDiskId")
    def data_disk_id(self) -> Optional[str]:
        """
        Data disk ID used to initialize the data disk. When data disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
        """
        return pulumi.get(self, "data_disk_id")

    @property
    @pulumi.getter(name="dataDiskSnapshotId")
    def data_disk_snapshot_id(self) -> Optional[str]:
        """
        Snapshot ID of the data disk. The selected data disk snapshot size must be smaller than the data disk size.
        """
        return pulumi.get(self, "data_disk_snapshot_id")

    @property
    @pulumi.getter(name="deleteWithInstance")
    def delete_with_instance(self) -> Optional[bool]:
        """
        Decides whether the disk is deleted with instance(only applied to `CLOUD_BASIC`, `CLOUD_SSD` and `CLOUD_PREMIUM` disk with `POSTPAID_BY_HOUR` instance), default is true.
        """
        return pulumi.get(self, "delete_with_instance")

    @property
    @pulumi.getter
    def encrypt(self) -> Optional[bool]:
        """
        Decides whether the disk is encrypted. Default is `false`.
        """
        return pulumi.get(self, "encrypt")

    @property
    @pulumi.getter(name="throughputPerformance")
    def throughput_performance(self) -> Optional[int]:
        """
        Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        return pulumi.get(self, "throughput_performance")


@pulumi.output_type
class GetTypesFilterResult(dict):
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: The filter name. Valid values: `zone`, `instance-family` and `instance-charge-type`.
        :param Sequence[str] values: The filter values.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The filter name. Valid values: `zone`, `instance-family` and `instance-charge-type`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        The filter values.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetTypesInstanceTypeResult(dict):
    def __init__(__self__, *,
                 availability_zone: str,
                 cpu_core_count: int,
                 family: str,
                 gpu_core_count: int,
                 instance_charge_type: str,
                 instance_type: str,
                 memory_size: int,
                 status: str):
        """
        :param str availability_zone: The available zone that the CVM instance locates at. This field is conflict with `filter`.
        :param int cpu_core_count: The number of CPU cores of the instance.
        :param str family: Type series of the instance.
        :param int gpu_core_count: The number of GPU cores of the instance.
        :param str instance_charge_type: Charge type of the instance.
        :param str instance_type: Type of the instance.
        :param int memory_size: Instance memory capacity, unit in GB.
        :param str status: Sell status of the instance.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "cpu_core_count", cpu_core_count)
        pulumi.set(__self__, "family", family)
        pulumi.set(__self__, "gpu_core_count", gpu_core_count)
        pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "memory_size", memory_size)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The available zone that the CVM instance locates at. This field is conflict with `filter`.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cpuCoreCount")
    def cpu_core_count(self) -> int:
        """
        The number of CPU cores of the instance.
        """
        return pulumi.get(self, "cpu_core_count")

    @property
    @pulumi.getter
    def family(self) -> str:
        """
        Type series of the instance.
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter(name="gpuCoreCount")
    def gpu_core_count(self) -> int:
        """
        The number of GPU cores of the instance.
        """
        return pulumi.get(self, "gpu_core_count")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> str:
        """
        Charge type of the instance.
        """
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        Type of the instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> int:
        """
        Instance memory capacity, unit in GB.
        """
        return pulumi.get(self, "memory_size")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Sell status of the instance.
        """
        return pulumi.get(self, "status")


