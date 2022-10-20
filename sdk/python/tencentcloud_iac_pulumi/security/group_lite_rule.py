# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupLiteRuleArgs', 'GroupLiteRule']

@pulumi.input_type
class GroupLiteRuleArgs:
    def __init__(__self__, *,
                 security_group_id: pulumi.Input[str],
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a GroupLiteRule resource.
        :param pulumi.Input[str] security_group_id: ID of the security group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] egresses: Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ingresses: Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        pulumi.set(__self__, "security_group_id", security_group_id)
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        ID of the security group.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ingresses", value)


@pulumi.input_type
class _GroupLiteRuleState:
    def __init__(__self__, *,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupLiteRule resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] egresses: Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ingresses: Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[str] security_group_id: ID of the security group.
        """
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the security group.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)


class GroupLiteRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide a resource to create security group some lite rules quickly.

        > **NOTE:** It can't be used with tencentcloud_security_group_rule, and don't create multiple Security.GroupRule resources, otherwise it may cause problems.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo_group = tencentcloud.security.Group("fooGroup")
        foo_group_lite_rule = tencentcloud.security.GroupLiteRule("fooGroupLiteRule",
            security_group_id=foo_group.id,
            ingresses=[
                "ACCEPT#192.168.1.0/24#80#TCP",
                "DROP#8.8.8.8#80,90#UDP",
                "ACCEPT#0.0.0.0/0#80-90#TCP",
                "ACCEPT#sg-7ixn3foj#80-90#TCP",
                "ACCEPT#ipm-epjq5kn0#80-90#TCP",
                "ACCEPT#ipmg-3loavam6#80-90#TCP",
            ],
            egresses=[
                "ACCEPT#192.168.0.0/16#ALL#TCP",
                "ACCEPT#10.0.0.0/8#ALL#ICMP",
                "DROP#0.0.0.0/0#ALL#ALL",
            ])
        ```

        ## Import

        Security group lite rule can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Security/groupLiteRule:GroupLiteRule tencentcloud_security_group_lite_rule.foo sg-ey3wmiz1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] egresses: Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ingresses: Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[str] security_group_id: ID of the security group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupLiteRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create security group some lite rules quickly.

        > **NOTE:** It can't be used with tencentcloud_security_group_rule, and don't create multiple Security.GroupRule resources, otherwise it may cause problems.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo_group = tencentcloud.security.Group("fooGroup")
        foo_group_lite_rule = tencentcloud.security.GroupLiteRule("fooGroupLiteRule",
            security_group_id=foo_group.id,
            ingresses=[
                "ACCEPT#192.168.1.0/24#80#TCP",
                "DROP#8.8.8.8#80,90#UDP",
                "ACCEPT#0.0.0.0/0#80-90#TCP",
                "ACCEPT#sg-7ixn3foj#80-90#TCP",
                "ACCEPT#ipm-epjq5kn0#80-90#TCP",
                "ACCEPT#ipmg-3loavam6#80-90#TCP",
            ],
            egresses=[
                "ACCEPT#192.168.0.0/16#ALL#TCP",
                "ACCEPT#10.0.0.0/8#ALL#ICMP",
                "DROP#0.0.0.0/0#ALL#ALL",
            ])
        ```

        ## Import

        Security group lite rule can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Security/groupLiteRule:GroupLiteRule tencentcloud_security_group_lite_rule.foo sg-ey3wmiz1
        ```

        :param str resource_name: The name of the resource.
        :param GroupLiteRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupLiteRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = GroupLiteRuleArgs.__new__(GroupLiteRuleArgs)

            __props__.__dict__["egresses"] = egresses
            __props__.__dict__["ingresses"] = ingresses
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
        super(GroupLiteRule, __self__).__init__(
            'tencentcloud:Security/groupLiteRule:GroupLiteRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            egresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None) -> 'GroupLiteRule':
        """
        Get an existing GroupLiteRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] egresses: Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ingresses: Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        :param pulumi.Input[str] security_group_id: ID of the security group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupLiteRuleState.__new__(_GroupLiteRuleState)

        __props__.__dict__["egresses"] = egresses
        __props__.__dict__["ingresses"] = ingresses
        __props__.__dict__["security_group_id"] = security_group_id
        return GroupLiteRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def egresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        return pulumi.get(self, "egresses")

    @property
    @pulumi.getter
    def ingresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is `ACCEPT` and `DROP`. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is `80`, `80,443`, `80-90` or `ALL`. The available value of 'protocol' is `TCP`, `UDP`, `ICMP` and `ALL`. When 'protocol' is `ICMP` or `ALL`, the 'port' must be `ALL`.
        """
        return pulumi.get(self, "ingresses")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        ID of the security group.
        """
        return pulumi.get(self, "security_group_id")

