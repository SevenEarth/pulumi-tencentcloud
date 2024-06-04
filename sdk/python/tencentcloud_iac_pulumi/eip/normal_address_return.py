# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NormalAddressReturnArgs', 'NormalAddressReturn']

@pulumi.input_type
class NormalAddressReturnArgs:
    def __init__(__self__, *,
                 address_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a NormalAddressReturn resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_ips: The IP address of the EIP, example: 101.35.139.183.
        """
        if address_ips is not None:
            pulumi.set(__self__, "address_ips", address_ips)

    @property
    @pulumi.getter(name="addressIps")
    def address_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IP address of the EIP, example: 101.35.139.183.
        """
        return pulumi.get(self, "address_ips")

    @address_ips.setter
    def address_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_ips", value)


@pulumi.input_type
class _NormalAddressReturnState:
    def __init__(__self__, *,
                 address_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering NormalAddressReturn resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_ips: The IP address of the EIP, example: 101.35.139.183.
        """
        if address_ips is not None:
            pulumi.set(__self__, "address_ips", address_ips)

    @property
    @pulumi.getter(name="addressIps")
    def address_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IP address of the EIP, example: 101.35.139.183.
        """
        return pulumi.get(self, "address_ips")

    @address_ips.setter
    def address_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_ips", value)


class NormalAddressReturn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc normal_address_return

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.eip.NormalAddressReturn("example", address_ips=["172.16.17.32"])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_ips: The IP address of the EIP, example: 101.35.139.183.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NormalAddressReturnArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc normal_address_return

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.eip.NormalAddressReturn("example", address_ips=["172.16.17.32"])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param NormalAddressReturnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NormalAddressReturnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NormalAddressReturnArgs.__new__(NormalAddressReturnArgs)

            __props__.__dict__["address_ips"] = address_ips
        super(NormalAddressReturn, __self__).__init__(
            'tencentcloud:Eip/normalAddressReturn:NormalAddressReturn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'NormalAddressReturn':
        """
        Get an existing NormalAddressReturn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_ips: The IP address of the EIP, example: 101.35.139.183.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NormalAddressReturnState.__new__(_NormalAddressReturnState)

        __props__.__dict__["address_ips"] = address_ips
        return NormalAddressReturn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressIps")
    def address_ips(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IP address of the EIP, example: 101.35.139.183.
        """
        return pulumi.get(self, "address_ips")

