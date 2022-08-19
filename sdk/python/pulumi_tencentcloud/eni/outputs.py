# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceIpv4',
    'InstanceIpv4Info',
]

@pulumi.output_type
class InstanceIpv4(dict):
    def __init__(__self__, *,
                 ip: str,
                 primary: bool,
                 description: Optional[str] = None):
        """
        :param str ip: Intranet IP.
        :param bool primary: Indicates whether the IP is primary.
        :param str description: Description of the IP, maximum length 25.
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "primary", primary)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        Intranet IP.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def primary(self) -> bool:
        """
        Indicates whether the IP is primary.
        """
        return pulumi.get(self, "primary")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the IP, maximum length 25.
        """
        return pulumi.get(self, "description")


@pulumi.output_type
class InstanceIpv4Info(dict):
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 ip: Optional[str] = None,
                 primary: Optional[bool] = None):
        """
        :param str description: Description of the IP, maximum length 25.
        :param str ip: Intranet IP.
        :param bool primary: Indicates whether the IP is primary.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if primary is not None:
            pulumi.set(__self__, "primary", primary)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the IP, maximum length 25.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        Intranet IP.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def primary(self) -> Optional[bool]:
        """
        Indicates whether the IP is primary.
        """
        return pulumi.get(self, "primary")


