# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CustomHeaderHeaderArgs',
    'HttpRuleRealserverArgs',
    'Layer4ListenerRealserverBindSetArgs',
    'GetProxyGroupsFilterArgs',
    'GetProxyGroupsTagSetArgs',
]

@pulumi.input_type
class CustomHeaderHeaderArgs:
    def __init__(__self__, *,
                 header_name: pulumi.Input[str],
                 header_value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] header_name: Header name.
        :param pulumi.Input[str] header_value: Header value.
        """
        pulumi.set(__self__, "header_name", header_name)
        pulumi.set(__self__, "header_value", header_value)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> pulumi.Input[str]:
        """
        Header name.
        """
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter(name="headerValue")
    def header_value(self) -> pulumi.Input[str]:
        """
        Header value.
        """
        return pulumi.get(self, "header_value")

    @header_value.setter
    def header_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_value", value)


@pulumi.input_type
class HttpRuleRealserverArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int],
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] id: ID of the GAAP realserver.
        :param pulumi.Input[str] ip: IP of the GAAP realserver.
        :param pulumi.Input[int] port: Port of the GAAP realserver.
        :param pulumi.Input[int] weight: Scheduling weight, default value is `1`. Valid value ranges: (1~100).
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        ID of the GAAP realserver.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP of the GAAP realserver.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Port of the GAAP realserver.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        Scheduling weight, default value is `1`. Valid value ranges: (1~100).
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class Layer4ListenerRealserverBindSetArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int],
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] id: ID of the GAAP realserver.
        :param pulumi.Input[str] ip: IP of the GAAP realserver.
        :param pulumi.Input[int] port: Port of the GAAP realserver.
        :param pulumi.Input[int] weight: Scheduling weight, default value is `1`. The range of values is [1,100].
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        ID of the GAAP realserver.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP of the GAAP realserver.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Port of the GAAP realserver.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        Scheduling weight, default value is `1`. The range of values is [1,100].
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class GetProxyGroupsFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: Filter conditions.
        :param Sequence[str] values: filtering value.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Filter conditions.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        filtering value.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GetProxyGroupsTagSetArgs:
    def __init__(__self__, *,
                 tag_key: str,
                 tag_value: str):
        """
        :param str tag_key: Tag Key.
        :param str tag_value: Tag Value.
        """
        pulumi.set(__self__, "tag_key", tag_key)
        pulumi.set(__self__, "tag_value", tag_value)

    @property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> str:
        """
        Tag Key.
        """
        return pulumi.get(self, "tag_key")

    @tag_key.setter
    def tag_key(self, value: str):
        pulumi.set(self, "tag_key", value)

    @property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> str:
        """
        Tag Value.
        """
        return pulumi.get(self, "tag_value")

    @tag_value.setter
    def tag_value(self, value: str):
        pulumi.set(self, "tag_value", value)


