# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetInstanceDnatListResult',
]

@pulumi.output_type
class GetInstanceDnatListResult(dict):
    def __init__(__self__, *,
                 elastic_ip: str,
                 elastic_port: str,
                 nat_id: str,
                 private_ip: str,
                 private_port: str,
                 protocol: str,
                 vpc_id: str,
                 description: Optional[str] = None):
        """
        :param str elastic_ip: Network address of the EIP.
        :param str elastic_port: Port of the EIP.
        :param str nat_id: ID of the NAT gateway.
        :param str private_ip: Network address of the backend service.
        :param str private_port: Port of intranet.
        :param str protocol: Type of the network protocol. Valid values: `TCP` and `UDP`.
        :param str vpc_id: ID of the VPC.
        :param str description: Description of the NAT forward.
        """
        pulumi.set(__self__, "elastic_ip", elastic_ip)
        pulumi.set(__self__, "elastic_port", elastic_port)
        pulumi.set(__self__, "nat_id", nat_id)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "private_port", private_port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="elasticIp")
    def elastic_ip(self) -> str:
        """
        Network address of the EIP.
        """
        return pulumi.get(self, "elastic_ip")

    @property
    @pulumi.getter(name="elasticPort")
    def elastic_port(self) -> str:
        """
        Port of the EIP.
        """
        return pulumi.get(self, "elastic_port")

    @property
    @pulumi.getter(name="natId")
    def nat_id(self) -> str:
        """
        ID of the NAT gateway.
        """
        return pulumi.get(self, "nat_id")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        Network address of the backend service.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="privatePort")
    def private_port(self) -> str:
        """
        Port of intranet.
        """
        return pulumi.get(self, "private_port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Type of the network protocol. Valid values: `TCP` and `UDP`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the NAT forward.
        """
        return pulumi.get(self, "description")

