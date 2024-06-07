# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceHostResourceArgs',
]

@pulumi.input_type
class InstanceHostResourceArgs:
    def __init__(__self__, *,
                 cpu_available_num: Optional[pulumi.Input[int]] = None,
                 cpu_total_num: Optional[pulumi.Input[int]] = None,
                 disk_available_size: Optional[pulumi.Input[int]] = None,
                 disk_total_size: Optional[pulumi.Input[int]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 memory_available_size: Optional[pulumi.Input[float]] = None,
                 memory_total_size: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[int] cpu_available_num: The number of available CPU cores of the instance.
        :param pulumi.Input[int] cpu_total_num: The number of total CPU cores of the instance.
        :param pulumi.Input[int] disk_available_size: Instance disk available capacity, unit in GB.
        :param pulumi.Input[int] disk_total_size: Instance disk total capacity, unit in GB.
        :param pulumi.Input[str] disk_type: Type of the disk.
        :param pulumi.Input[float] memory_available_size: Instance memory available capacity, unit in GB.
        :param pulumi.Input[float] memory_total_size: Instance memory total capacity, unit in GB.
        """
        if cpu_available_num is not None:
            pulumi.set(__self__, "cpu_available_num", cpu_available_num)
        if cpu_total_num is not None:
            pulumi.set(__self__, "cpu_total_num", cpu_total_num)
        if disk_available_size is not None:
            pulumi.set(__self__, "disk_available_size", disk_available_size)
        if disk_total_size is not None:
            pulumi.set(__self__, "disk_total_size", disk_total_size)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if memory_available_size is not None:
            pulumi.set(__self__, "memory_available_size", memory_available_size)
        if memory_total_size is not None:
            pulumi.set(__self__, "memory_total_size", memory_total_size)

    @property
    @pulumi.getter(name="cpuAvailableNum")
    def cpu_available_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of available CPU cores of the instance.
        """
        return pulumi.get(self, "cpu_available_num")

    @cpu_available_num.setter
    def cpu_available_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_available_num", value)

    @property
    @pulumi.getter(name="cpuTotalNum")
    def cpu_total_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of total CPU cores of the instance.
        """
        return pulumi.get(self, "cpu_total_num")

    @cpu_total_num.setter
    def cpu_total_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_total_num", value)

    @property
    @pulumi.getter(name="diskAvailableSize")
    def disk_available_size(self) -> Optional[pulumi.Input[int]]:
        """
        Instance disk available capacity, unit in GB.
        """
        return pulumi.get(self, "disk_available_size")

    @disk_available_size.setter
    def disk_available_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_available_size", value)

    @property
    @pulumi.getter(name="diskTotalSize")
    def disk_total_size(self) -> Optional[pulumi.Input[int]]:
        """
        Instance disk total capacity, unit in GB.
        """
        return pulumi.get(self, "disk_total_size")

    @disk_total_size.setter
    def disk_total_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_total_size", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the disk.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter(name="memoryAvailableSize")
    def memory_available_size(self) -> Optional[pulumi.Input[float]]:
        """
        Instance memory available capacity, unit in GB.
        """
        return pulumi.get(self, "memory_available_size")

    @memory_available_size.setter
    def memory_available_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "memory_available_size", value)

    @property
    @pulumi.getter(name="memoryTotalSize")
    def memory_total_size(self) -> Optional[pulumi.Input[float]]:
        """
        Instance memory total capacity, unit in GB.
        """
        return pulumi.get(self, "memory_total_size")

    @memory_total_size.setter
    def memory_total_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "memory_total_size", value)


