# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DhcpAssociateAddressArgs', 'DhcpAssociateAddress']

@pulumi.input_type
class DhcpAssociateAddressArgs:
    def __init__(__self__, *,
                 address_ip: pulumi.Input[str],
                 dhcp_ip_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DhcpAssociateAddress resource.
        :param pulumi.Input[str] address_ip: Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        :param pulumi.Input[str] dhcp_ip_id: `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        pulumi.set(__self__, "address_ip", address_ip)
        pulumi.set(__self__, "dhcp_ip_id", dhcp_ip_id)

    @property
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> pulumi.Input[str]:
        """
        Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        """
        return pulumi.get(self, "address_ip")

    @address_ip.setter
    def address_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_ip", value)

    @property
    @pulumi.getter(name="dhcpIpId")
    def dhcp_ip_id(self) -> pulumi.Input[str]:
        """
        `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        return pulumi.get(self, "dhcp_ip_id")

    @dhcp_ip_id.setter
    def dhcp_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dhcp_ip_id", value)


@pulumi.input_type
class _DhcpAssociateAddressState:
    def __init__(__self__, *,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_ip_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DhcpAssociateAddress resources.
        :param pulumi.Input[str] address_ip: Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        :param pulumi.Input[str] dhcp_ip_id: `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        if address_ip is not None:
            pulumi.set(__self__, "address_ip", address_ip)
        if dhcp_ip_id is not None:
            pulumi.set(__self__, "dhcp_ip_id", dhcp_ip_id)

    @property
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        """
        return pulumi.get(self, "address_ip")

    @address_ip.setter
    def address_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_ip", value)

    @property
    @pulumi.getter(name="dhcpIpId")
    def dhcp_ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        return pulumi.get(self, "dhcp_ip_id")

    @dhcp_ip_id.setter
    def dhcp_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_ip_id", value)


class DhcpAssociateAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_ip_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DhcpAssociateAddress resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_ip: Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        :param pulumi.Input[str] dhcp_ip_id: `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DhcpAssociateAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DhcpAssociateAddress resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DhcpAssociateAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DhcpAssociateAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 dhcp_ip_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DhcpAssociateAddressArgs.__new__(DhcpAssociateAddressArgs)

            if address_ip is None and not opts.urn:
                raise TypeError("Missing required property 'address_ip'")
            __props__.__dict__["address_ip"] = address_ip
            if dhcp_ip_id is None and not opts.urn:
                raise TypeError("Missing required property 'dhcp_ip_id'")
            __props__.__dict__["dhcp_ip_id"] = dhcp_ip_id
        super(DhcpAssociateAddress, __self__).__init__(
            'tencentcloud:Vpc/dhcpAssociateAddress:DhcpAssociateAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_ip: Optional[pulumi.Input[str]] = None,
            dhcp_ip_id: Optional[pulumi.Input[str]] = None) -> 'DhcpAssociateAddress':
        """
        Get an existing DhcpAssociateAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_ip: Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        :param pulumi.Input[str] dhcp_ip_id: `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DhcpAssociateAddressState.__new__(_DhcpAssociateAddressState)

        __props__.__dict__["address_ip"] = address_ip
        __props__.__dict__["dhcp_ip_id"] = dhcp_ip_id
        return DhcpAssociateAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> pulumi.Output[str]:
        """
        Elastic public network `IP`. Must be `EIP` not bound to `DhcpIp`.
        """
        return pulumi.get(self, "address_ip")

    @property
    @pulumi.getter(name="dhcpIpId")
    def dhcp_ip_id(self) -> pulumi.Output[str]:
        """
        `DhcpIp` unique `ID`, like: `dhcpip-9o233uri`. Must be a `DhcpIp` that is not bound to `EIP`.
        """
        return pulumi.get(self, "dhcp_ip_id")

