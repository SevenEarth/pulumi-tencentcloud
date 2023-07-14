# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DhcpIpArgs', 'DhcpIp']

@pulumi.input_type
class DhcpIpArgs:
    def __init__(__self__, *,
                 dhcp_ip_name: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DhcpIp resource.
        :param pulumi.Input[str] dhcp_ip_name: `DhcpIp` name.
        :param pulumi.Input[str] subnet_id: Subnet `ID`.
        :param pulumi.Input[str] vpc_id: The private network `ID`.
        """
        pulumi.set(__self__, "dhcp_ip_name", dhcp_ip_name)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="dhcpIpName")
    def dhcp_ip_name(self) -> pulumi.Input[str]:
        """
        `DhcpIp` name.
        """
        return pulumi.get(self, "dhcp_ip_name")

    @dhcp_ip_name.setter
    def dhcp_ip_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "dhcp_ip_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Subnet `ID`.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The private network `ID`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _DhcpIpState:
    def __init__(__self__, *,
                 dhcp_ip_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DhcpIp resources.
        :param pulumi.Input[str] dhcp_ip_name: `DhcpIp` name.
        :param pulumi.Input[str] subnet_id: Subnet `ID`.
        :param pulumi.Input[str] vpc_id: The private network `ID`.
        """
        if dhcp_ip_name is not None:
            pulumi.set(__self__, "dhcp_ip_name", dhcp_ip_name)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="dhcpIpName")
    def dhcp_ip_name(self) -> Optional[pulumi.Input[str]]:
        """
        `DhcpIp` name.
        """
        return pulumi.get(self, "dhcp_ip_name")

    @dhcp_ip_name.setter
    def dhcp_ip_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dhcp_ip_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet `ID`.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The private network `ID`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class DhcpIp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp_ip_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc dhcp_ip

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        dhcp_ip = tencentcloud.vpc.DhcpIp("dhcpIp",
            dhcp_ip_name="terraform-test",
            subnet_id="subnet-h7av55g8",
            vpc_id="vpc-1yg5ua6l")
        ```

        ## Import

        vpc dhcp_ip can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/dhcpIp:DhcpIp dhcp_ip dhcp_ip_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dhcp_ip_name: `DhcpIp` name.
        :param pulumi.Input[str] subnet_id: Subnet `ID`.
        :param pulumi.Input[str] vpc_id: The private network `ID`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DhcpIpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc dhcp_ip

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        dhcp_ip = tencentcloud.vpc.DhcpIp("dhcpIp",
            dhcp_ip_name="terraform-test",
            subnet_id="subnet-h7av55g8",
            vpc_id="vpc-1yg5ua6l")
        ```

        ## Import

        vpc dhcp_ip can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/dhcpIp:DhcpIp dhcp_ip dhcp_ip_id
        ```

        :param str resource_name: The name of the resource.
        :param DhcpIpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DhcpIpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp_ip_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DhcpIpArgs.__new__(DhcpIpArgs)

            if dhcp_ip_name is None and not opts.urn:
                raise TypeError("Missing required property 'dhcp_ip_name'")
            __props__.__dict__["dhcp_ip_name"] = dhcp_ip_name
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(DhcpIp, __self__).__init__(
            'tencentcloud:Vpc/dhcpIp:DhcpIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dhcp_ip_name: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'DhcpIp':
        """
        Get an existing DhcpIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dhcp_ip_name: `DhcpIp` name.
        :param pulumi.Input[str] subnet_id: Subnet `ID`.
        :param pulumi.Input[str] vpc_id: The private network `ID`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DhcpIpState.__new__(_DhcpIpState)

        __props__.__dict__["dhcp_ip_name"] = dhcp_ip_name
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["vpc_id"] = vpc_id
        return DhcpIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dhcpIpName")
    def dhcp_ip_name(self) -> pulumi.Output[str]:
        """
        `DhcpIp` name.
        """
        return pulumi.get(self, "dhcp_ip_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Subnet `ID`.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The private network `ID`.
        """
        return pulumi.get(self, "vpc_id")

