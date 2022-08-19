# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterParamItemArgs',
    'ClusterRoGroupAddrArgs',
    'ClusterRoGroupInstanceArgs',
    'ClusterRwGroupAddrArgs',
    'ClusterRwGroupInstanceArgs',
]

@pulumi.input_type
class ClusterParamItemArgs:
    def __init__(__self__, *,
                 current_value: pulumi.Input[str],
                 name: pulumi.Input[str],
                 old_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] current_value: Param expected value to set.
        :param pulumi.Input[str] name: Name of param, e.g. `character_set_server`.
        :param pulumi.Input[str] old_value: Param old value, indicates the value which already set, this value is required when modifying current_value.
        """
        pulumi.set(__self__, "current_value", current_value)
        pulumi.set(__self__, "name", name)
        if old_value is not None:
            pulumi.set(__self__, "old_value", old_value)

    @property
    @pulumi.getter(name="currentValue")
    def current_value(self) -> pulumi.Input[str]:
        """
        Param expected value to set.
        """
        return pulumi.get(self, "current_value")

    @current_value.setter
    def current_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "current_value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of param, e.g. `character_set_server`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="oldValue")
    def old_value(self) -> Optional[pulumi.Input[str]]:
        """
        Param old value, indicates the value which already set, this value is required when modifying current_value.
        """
        return pulumi.get(self, "old_value")

    @old_value.setter
    def old_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "old_value", value)


@pulumi.input_type
class ClusterRoGroupAddrArgs:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] ip: IP address for read-write connection.
        :param pulumi.Input[int] port: Port of CynosDB cluster.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address for read-write connection.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of CynosDB cluster.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class ClusterRoGroupInstanceArgs:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance_id: ID of instance.
        :param pulumi.Input[str] instance_name: Name of instance.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)


@pulumi.input_type
class ClusterRwGroupAddrArgs:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] ip: IP address for read-write connection.
        :param pulumi.Input[int] port: Port of CynosDB cluster.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address for read-write connection.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of CynosDB cluster.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class ClusterRwGroupInstanceArgs:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance_id: ID of instance.
        :param pulumi.Input[str] instance_name: Name of instance.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)


