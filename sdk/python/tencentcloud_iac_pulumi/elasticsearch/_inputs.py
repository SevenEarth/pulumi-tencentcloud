# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceEsAclArgs',
    'InstanceMultiZoneInfoArgs',
    'InstanceNodeInfoListArgs',
    'InstanceWebNodeTypeInfoArgs',
]

@pulumi.input_type
class InstanceEsAclArgs:
    def __init__(__self__, *,
                 black_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 white_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] black_lists: Blacklist of kibana access.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] white_lists: Whitelist of kibana access.
        """
        if black_lists is not None:
            pulumi.set(__self__, "black_lists", black_lists)
        if white_lists is not None:
            pulumi.set(__self__, "white_lists", white_lists)

    @property
    @pulumi.getter(name="blackLists")
    def black_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Blacklist of kibana access.
        """
        return pulumi.get(self, "black_lists")

    @black_lists.setter
    def black_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "black_lists", value)

    @property
    @pulumi.getter(name="whiteLists")
    def white_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Whitelist of kibana access.
        """
        return pulumi.get(self, "white_lists")

    @white_lists.setter
    def white_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "white_lists", value)


@pulumi.input_type
class InstanceMultiZoneInfoArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] availability_zone: Availability zone.
        :param pulumi.Input[str] subnet_id: The ID of a VPC subnetwork.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        Availability zone.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of a VPC subnetwork.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class InstanceNodeInfoListArgs:
    def __init__(__self__, *,
                 node_num: pulumi.Input[int],
                 node_type: pulumi.Input[str],
                 disk_size: Optional[pulumi.Input[int]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 encrypt: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] node_num: Number of nodes.
        :param pulumi.Input[str] node_type: Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).
        :param pulumi.Input[int] disk_size: Node disk size. Unit is GB, and default value is `100`.
        :param pulumi.Input[str] disk_type: Node disk type. Valid values are `CLOUD_SSD` and `CLOUD_PREMIUM`. The default value is `CLOUD_SSD`.
        :param pulumi.Input[bool] encrypt: Decides to encrypt this disk or not.
        :param pulumi.Input[str] type: Node type. Valid values are `hotData`, `warmData` and `dedicatedMaster`. The default value is 'hotData`.
        """
        pulumi.set(__self__, "node_num", node_num)
        pulumi.set(__self__, "node_type", node_type)
        if disk_size is not None:
            pulumi.set(__self__, "disk_size", disk_size)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if encrypt is not None:
            pulumi.set(__self__, "encrypt", encrypt)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Input[int]:
        """
        Number of nodes.
        """
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        """
        Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="diskSize")
    def disk_size(self) -> Optional[pulumi.Input[int]]:
        """
        Node disk size. Unit is GB, and default value is `100`.
        """
        return pulumi.get(self, "disk_size")

    @disk_size.setter
    def disk_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        Node disk type. Valid values are `CLOUD_SSD` and `CLOUD_PREMIUM`. The default value is `CLOUD_SSD`.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter
    def encrypt(self) -> Optional[pulumi.Input[bool]]:
        """
        Decides to encrypt this disk or not.
        """
        return pulumi.get(self, "encrypt")

    @encrypt.setter
    def encrypt(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encrypt", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Node type. Valid values are `hotData`, `warmData` and `dedicatedMaster`. The default value is 'hotData`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class InstanceWebNodeTypeInfoArgs:
    def __init__(__self__, *,
                 node_num: pulumi.Input[int],
                 node_type: pulumi.Input[str]):
        """
        :param pulumi.Input[int] node_num: Visual node number.
        :param pulumi.Input[str] node_type: Visual node specifications.
        """
        pulumi.set(__self__, "node_num", node_num)
        pulumi.set(__self__, "node_type", node_type)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Input[int]:
        """
        Visual node number.
        """
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        """
        Visual node specifications.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)


