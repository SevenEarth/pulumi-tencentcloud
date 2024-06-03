# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CustomerGatewayConfigurationDownloadArgs', 'CustomerGatewayConfigurationDownload']

@pulumi.input_type
class CustomerGatewayConfigurationDownloadArgs:
    def __init__(__self__, *,
                 customer_gateway_vendor: pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs'],
                 interface_name: pulumi.Input[str],
                 vpn_connection_id: pulumi.Input[str],
                 vpn_gateway_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a CustomerGatewayConfigurationDownload resource.
        :param pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs'] customer_gateway_vendor: Customer Gateway Vendor Info.
        :param pulumi.Input[str] interface_name: VPN connection access device physical interface name.
        :param pulumi.Input[str] vpn_connection_id: VPN Connection Instance id.
        :param pulumi.Input[str] vpn_gateway_id: VPN Gateway Instance ID.
        """
        pulumi.set(__self__, "customer_gateway_vendor", customer_gateway_vendor)
        pulumi.set(__self__, "interface_name", interface_name)
        pulumi.set(__self__, "vpn_connection_id", vpn_connection_id)
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="customerGatewayVendor")
    def customer_gateway_vendor(self) -> pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']:
        """
        Customer Gateway Vendor Info.
        """
        return pulumi.get(self, "customer_gateway_vendor")

    @customer_gateway_vendor.setter
    def customer_gateway_vendor(self, value: pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']):
        pulumi.set(self, "customer_gateway_vendor", value)

    @property
    @pulumi.getter(name="interfaceName")
    def interface_name(self) -> pulumi.Input[str]:
        """
        VPN connection access device physical interface name.
        """
        return pulumi.get(self, "interface_name")

    @interface_name.setter
    def interface_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface_name", value)

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> pulumi.Input[str]:
        """
        VPN Connection Instance id.
        """
        return pulumi.get(self, "vpn_connection_id")

    @vpn_connection_id.setter
    def vpn_connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpn_connection_id", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Input[str]:
        """
        VPN Gateway Instance ID.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpn_gateway_id", value)


@pulumi.input_type
class _CustomerGatewayConfigurationDownloadState:
    def __init__(__self__, *,
                 customer_gateway_configuration: Optional[pulumi.Input[str]] = None,
                 customer_gateway_vendor: Optional[pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']] = None,
                 interface_name: Optional[pulumi.Input[str]] = None,
                 vpn_connection_id: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomerGatewayConfigurationDownload resources.
        :param pulumi.Input[str] customer_gateway_configuration: xml configuration.
        :param pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs'] customer_gateway_vendor: Customer Gateway Vendor Info.
        :param pulumi.Input[str] interface_name: VPN connection access device physical interface name.
        :param pulumi.Input[str] vpn_connection_id: VPN Connection Instance id.
        :param pulumi.Input[str] vpn_gateway_id: VPN Gateway Instance ID.
        """
        if customer_gateway_configuration is not None:
            pulumi.set(__self__, "customer_gateway_configuration", customer_gateway_configuration)
        if customer_gateway_vendor is not None:
            pulumi.set(__self__, "customer_gateway_vendor", customer_gateway_vendor)
        if interface_name is not None:
            pulumi.set(__self__, "interface_name", interface_name)
        if vpn_connection_id is not None:
            pulumi.set(__self__, "vpn_connection_id", vpn_connection_id)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="customerGatewayConfiguration")
    def customer_gateway_configuration(self) -> Optional[pulumi.Input[str]]:
        """
        xml configuration.
        """
        return pulumi.get(self, "customer_gateway_configuration")

    @customer_gateway_configuration.setter
    def customer_gateway_configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_gateway_configuration", value)

    @property
    @pulumi.getter(name="customerGatewayVendor")
    def customer_gateway_vendor(self) -> Optional[pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']]:
        """
        Customer Gateway Vendor Info.
        """
        return pulumi.get(self, "customer_gateway_vendor")

    @customer_gateway_vendor.setter
    def customer_gateway_vendor(self, value: Optional[pulumi.Input['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']]):
        pulumi.set(self, "customer_gateway_vendor", value)

    @property
    @pulumi.getter(name="interfaceName")
    def interface_name(self) -> Optional[pulumi.Input[str]]:
        """
        VPN connection access device physical interface name.
        """
        return pulumi.get(self, "interface_name")

    @interface_name.setter
    def interface_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_name", value)

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPN Connection Instance id.
        """
        return pulumi.get(self, "vpn_connection_id")

    @vpn_connection_id.setter
    def vpn_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_connection_id", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPN Gateway Instance ID.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)


class CustomerGatewayConfigurationDownload(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 customer_gateway_vendor: Optional[pulumi.Input[pulumi.InputType['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']]] = None,
                 interface_name: Optional[pulumi.Input[str]] = None,
                 vpn_connection_id: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc vpn_customer_gateway_configuration_download

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        vpn_customer_gateway_configuration_download = tencentcloud.vpn.CustomerGatewayConfigurationDownload("vpnCustomerGatewayConfigurationDownload",
            customer_gateway_vendor=tencentcloud.vpn.CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs(
                platform="comware",
                software_version="V1.0",
                vendor_name="h3c",
            ),
            interface_name="test",
            vpn_connection_id="vpnx-kme2tx8m",
            vpn_gateway_id="vpngw-gt8bianl")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']] customer_gateway_vendor: Customer Gateway Vendor Info.
        :param pulumi.Input[str] interface_name: VPN connection access device physical interface name.
        :param pulumi.Input[str] vpn_connection_id: VPN Connection Instance id.
        :param pulumi.Input[str] vpn_gateway_id: VPN Gateway Instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomerGatewayConfigurationDownloadArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc vpn_customer_gateway_configuration_download

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        vpn_customer_gateway_configuration_download = tencentcloud.vpn.CustomerGatewayConfigurationDownload("vpnCustomerGatewayConfigurationDownload",
            customer_gateway_vendor=tencentcloud.vpn.CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs(
                platform="comware",
                software_version="V1.0",
                vendor_name="h3c",
            ),
            interface_name="test",
            vpn_connection_id="vpnx-kme2tx8m",
            vpn_gateway_id="vpngw-gt8bianl")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param CustomerGatewayConfigurationDownloadArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomerGatewayConfigurationDownloadArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 customer_gateway_vendor: Optional[pulumi.Input[pulumi.InputType['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']]] = None,
                 interface_name: Optional[pulumi.Input[str]] = None,
                 vpn_connection_id: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomerGatewayConfigurationDownloadArgs.__new__(CustomerGatewayConfigurationDownloadArgs)

            if customer_gateway_vendor is None and not opts.urn:
                raise TypeError("Missing required property 'customer_gateway_vendor'")
            __props__.__dict__["customer_gateway_vendor"] = customer_gateway_vendor
            if interface_name is None and not opts.urn:
                raise TypeError("Missing required property 'interface_name'")
            __props__.__dict__["interface_name"] = interface_name
            if vpn_connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_connection_id'")
            __props__.__dict__["vpn_connection_id"] = vpn_connection_id
            if vpn_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
            __props__.__dict__["customer_gateway_configuration"] = None
        super(CustomerGatewayConfigurationDownload, __self__).__init__(
            'tencentcloud:Vpn/customerGatewayConfigurationDownload:CustomerGatewayConfigurationDownload',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            customer_gateway_configuration: Optional[pulumi.Input[str]] = None,
            customer_gateway_vendor: Optional[pulumi.Input[pulumi.InputType['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']]] = None,
            interface_name: Optional[pulumi.Input[str]] = None,
            vpn_connection_id: Optional[pulumi.Input[str]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None) -> 'CustomerGatewayConfigurationDownload':
        """
        Get an existing CustomerGatewayConfigurationDownload resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] customer_gateway_configuration: xml configuration.
        :param pulumi.Input[pulumi.InputType['CustomerGatewayConfigurationDownloadCustomerGatewayVendorArgs']] customer_gateway_vendor: Customer Gateway Vendor Info.
        :param pulumi.Input[str] interface_name: VPN connection access device physical interface name.
        :param pulumi.Input[str] vpn_connection_id: VPN Connection Instance id.
        :param pulumi.Input[str] vpn_gateway_id: VPN Gateway Instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomerGatewayConfigurationDownloadState.__new__(_CustomerGatewayConfigurationDownloadState)

        __props__.__dict__["customer_gateway_configuration"] = customer_gateway_configuration
        __props__.__dict__["customer_gateway_vendor"] = customer_gateway_vendor
        __props__.__dict__["interface_name"] = interface_name
        __props__.__dict__["vpn_connection_id"] = vpn_connection_id
        __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        return CustomerGatewayConfigurationDownload(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customerGatewayConfiguration")
    def customer_gateway_configuration(self) -> pulumi.Output[str]:
        """
        xml configuration.
        """
        return pulumi.get(self, "customer_gateway_configuration")

    @property
    @pulumi.getter(name="customerGatewayVendor")
    def customer_gateway_vendor(self) -> pulumi.Output['outputs.CustomerGatewayConfigurationDownloadCustomerGatewayVendor']:
        """
        Customer Gateway Vendor Info.
        """
        return pulumi.get(self, "customer_gateway_vendor")

    @property
    @pulumi.getter(name="interfaceName")
    def interface_name(self) -> pulumi.Output[str]:
        """
        VPN connection access device physical interface name.
        """
        return pulumi.get(self, "interface_name")

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> pulumi.Output[str]:
        """
        VPN Connection Instance id.
        """
        return pulumi.get(self, "vpn_connection_id")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[str]:
        """
        VPN Gateway Instance ID.
        """
        return pulumi.get(self, "vpn_gateway_id")

