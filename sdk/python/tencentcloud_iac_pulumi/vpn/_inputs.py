# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ConnectionSecurityGroupPolicyArgs',
    'CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs',
]

@pulumi.input_type
class ConnectionSecurityGroupPolicyArgs:
    def __init__(__self__, *,
                 local_cidr_block: pulumi.Input[str],
                 remote_cidr_blocks: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] local_cidr_block: Local cidr block.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] remote_cidr_blocks: Remote cidr block list.
        """
        pulumi.set(__self__, "local_cidr_block", local_cidr_block)
        pulumi.set(__self__, "remote_cidr_blocks", remote_cidr_blocks)

    @property
    @pulumi.getter(name="localCidrBlock")
    def local_cidr_block(self) -> pulumi.Input[str]:
        """
        Local cidr block.
        """
        return pulumi.get(self, "local_cidr_block")

    @local_cidr_block.setter
    def local_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_cidr_block", value)

    @property
    @pulumi.getter(name="remoteCidrBlocks")
    def remote_cidr_blocks(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Remote cidr block list.
        """
        return pulumi.get(self, "remote_cidr_blocks")

    @remote_cidr_blocks.setter
    def remote_cidr_blocks(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "remote_cidr_blocks", value)


@pulumi.input_type
class CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs:
    def __init__(__self__, *,
                 platform: pulumi.Input[str],
                 software_version: pulumi.Input[str],
                 vendor_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] platform: Platform.
        :param pulumi.Input[str] software_version: SoftwareVersion.
        :param pulumi.Input[str] vendor_name: VendorName.
        """
        pulumi.set(__self__, "platform", platform)
        pulumi.set(__self__, "software_version", software_version)
        pulumi.set(__self__, "vendor_name", vendor_name)

    @property
    @pulumi.getter
    def platform(self) -> pulumi.Input[str]:
        """
        Platform.
        """
        return pulumi.get(self, "platform")

    @platform.setter
    def platform(self, value: pulumi.Input[str]):
        pulumi.set(self, "platform", value)

    @property
    @pulumi.getter(name="softwareVersion")
    def software_version(self) -> pulumi.Input[str]:
        """
        SoftwareVersion.
        """
        return pulumi.get(self, "software_version")

    @software_version.setter
    def software_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "software_version", value)

    @property
    @pulumi.getter(name="vendorName")
    def vendor_name(self) -> pulumi.Input[str]:
        """
        VendorName.
        """
        return pulumi.get(self, "vendor_name")

    @vendor_name.setter
    def vendor_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vendor_name", value)


