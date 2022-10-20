# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GroupRuleArgs', 'GroupRule']

@pulumi.input_type
class GroupRuleArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input[str],
                 security_group_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 address_template: Optional[pulumi.Input['GroupRuleAddressTemplateArgs']] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 port_range: Optional[pulumi.Input[str]] = None,
                 protocol_template: Optional[pulumi.Input['GroupRuleProtocolTemplateArgs']] = None,
                 source_sgid: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupRule resource.
        :param pulumi.Input[str] policy: Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[str] type: Type of the security group rule. Valid values: `ingress` and `egress`.
        :param pulumi.Input['GroupRuleAddressTemplateArgs'] address_template: ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        :param pulumi.Input[str] cidr_ip: An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        :param pulumi.Input[str] description: Description of the security group rule.
        :param pulumi.Input[str] ip_protocol: Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        :param pulumi.Input[str] port_range: Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        :param pulumi.Input['GroupRuleProtocolTemplateArgs'] protocol_template: ID of the address template, and conflict with `ip_protocol`, `port_range`.
        :param pulumi.Input[str] source_sgid: ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        """
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "type", type)
        if address_template is not None:
            pulumi.set(__self__, "address_template", address_template)
        if cidr_ip is not None:
            pulumi.set(__self__, "cidr_ip", cidr_ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if port_range is not None:
            pulumi.set(__self__, "port_range", port_range)
        if protocol_template is not None:
            pulumi.set(__self__, "protocol_template", protocol_template)
        if source_sgid is not None:
            pulumi.set(__self__, "source_sgid", source_sgid)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        ID of the security group to be queried.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of the security group rule. Valid values: `ingress` and `egress`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="addressTemplate")
    def address_template(self) -> Optional[pulumi.Input['GroupRuleAddressTemplateArgs']]:
        """
        ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        """
        return pulumi.get(self, "address_template")

    @address_template.setter
    def address_template(self, value: Optional[pulumi.Input['GroupRuleAddressTemplateArgs']]):
        pulumi.set(self, "address_template", value)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        """
        return pulumi.get(self, "cidr_ip")

    @cidr_ip.setter
    def cidr_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_ip", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the security group rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[pulumi.Input[str]]:
        """
        Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        """
        return pulumi.get(self, "port_range")

    @port_range.setter
    def port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_range", value)

    @property
    @pulumi.getter(name="protocolTemplate")
    def protocol_template(self) -> Optional[pulumi.Input['GroupRuleProtocolTemplateArgs']]:
        """
        ID of the address template, and conflict with `ip_protocol`, `port_range`.
        """
        return pulumi.get(self, "protocol_template")

    @protocol_template.setter
    def protocol_template(self, value: Optional[pulumi.Input['GroupRuleProtocolTemplateArgs']]):
        pulumi.set(self, "protocol_template", value)

    @property
    @pulumi.getter(name="sourceSgid")
    def source_sgid(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        """
        return pulumi.get(self, "source_sgid")

    @source_sgid.setter
    def source_sgid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_sgid", value)


@pulumi.input_type
class _GroupRuleState:
    def __init__(__self__, *,
                 address_template: Optional[pulumi.Input['GroupRuleAddressTemplateArgs']] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 port_range: Optional[pulumi.Input[str]] = None,
                 protocol_template: Optional[pulumi.Input['GroupRuleProtocolTemplateArgs']] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 source_sgid: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupRule resources.
        :param pulumi.Input['GroupRuleAddressTemplateArgs'] address_template: ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        :param pulumi.Input[str] cidr_ip: An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        :param pulumi.Input[str] description: Description of the security group rule.
        :param pulumi.Input[str] ip_protocol: Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        :param pulumi.Input[str] policy: Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] port_range: Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        :param pulumi.Input['GroupRuleProtocolTemplateArgs'] protocol_template: ID of the address template, and conflict with `ip_protocol`, `port_range`.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[str] source_sgid: ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        :param pulumi.Input[str] type: Type of the security group rule. Valid values: `ingress` and `egress`.
        """
        if address_template is not None:
            pulumi.set(__self__, "address_template", address_template)
        if cidr_ip is not None:
            pulumi.set(__self__, "cidr_ip", cidr_ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if port_range is not None:
            pulumi.set(__self__, "port_range", port_range)
        if protocol_template is not None:
            pulumi.set(__self__, "protocol_template", protocol_template)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if source_sgid is not None:
            pulumi.set(__self__, "source_sgid", source_sgid)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="addressTemplate")
    def address_template(self) -> Optional[pulumi.Input['GroupRuleAddressTemplateArgs']]:
        """
        ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        """
        return pulumi.get(self, "address_template")

    @address_template.setter
    def address_template(self, value: Optional[pulumi.Input['GroupRuleAddressTemplateArgs']]):
        pulumi.set(self, "address_template", value)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        """
        return pulumi.get(self, "cidr_ip")

    @cidr_ip.setter
    def cidr_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_ip", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the security group rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> Optional[pulumi.Input[str]]:
        """
        Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        """
        return pulumi.get(self, "port_range")

    @port_range.setter
    def port_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_range", value)

    @property
    @pulumi.getter(name="protocolTemplate")
    def protocol_template(self) -> Optional[pulumi.Input['GroupRuleProtocolTemplateArgs']]:
        """
        ID of the address template, and conflict with `ip_protocol`, `port_range`.
        """
        return pulumi.get(self, "protocol_template")

    @protocol_template.setter
    def protocol_template(self, value: Optional[pulumi.Input['GroupRuleProtocolTemplateArgs']]):
        pulumi.set(self, "protocol_template", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the security group to be queried.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="sourceSgid")
    def source_sgid(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        """
        return pulumi.get(self, "source_sgid")

    @source_sgid.setter
    def source_sgid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_sgid", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the security group rule. Valid values: `ingress` and `egress`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class GroupRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_template: Optional[pulumi.Input[pulumi.InputType['GroupRuleAddressTemplateArgs']]] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 port_range: Optional[pulumi.Input[str]] = None,
                 protocol_template: Optional[pulumi.Input[pulumi.InputType['GroupRuleProtocolTemplateArgs']]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 source_sgid: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create security group rule.

        ## Example Usage

        Source is CIDR ip

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sglab1_group = tencentcloud.security.Group("sglab1Group",
            description="favourite sg_1",
            project_id=0)
        sglab1_group_rule = tencentcloud.security.GroupRule("sglab1GroupRule",
            security_group_id=sglab1_group.id,
            type="ingress",
            cidr_ip="10.0.0.0/16",
            ip_protocol="TCP",
            port_range="80",
            policy="ACCEPT",
            description="favourite sg rule_1")
        ```

        Source is a security group id

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sglab2_group = tencentcloud.security.Group("sglab2Group",
            description="favourite sg_2",
            project_id=0)
        sglab3 = tencentcloud.security.Group("sglab3",
            description="favourite sg_3",
            project_id=0)
        sglab2_group_rule = tencentcloud.security.GroupRule("sglab2GroupRule",
            security_group_id=sglab2_group.id,
            type="ingress",
            ip_protocol="TCP",
            port_range="80",
            policy="ACCEPT",
            source_sgid=sglab3.id,
            description="favourite sg rule_2")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GroupRuleAddressTemplateArgs']] address_template: ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        :param pulumi.Input[str] cidr_ip: An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        :param pulumi.Input[str] description: Description of the security group rule.
        :param pulumi.Input[str] ip_protocol: Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        :param pulumi.Input[str] policy: Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] port_range: Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        :param pulumi.Input[pulumi.InputType['GroupRuleProtocolTemplateArgs']] protocol_template: ID of the address template, and conflict with `ip_protocol`, `port_range`.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[str] source_sgid: ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        :param pulumi.Input[str] type: Type of the security group rule. Valid values: `ingress` and `egress`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create security group rule.

        ## Example Usage

        Source is CIDR ip

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sglab1_group = tencentcloud.security.Group("sglab1Group",
            description="favourite sg_1",
            project_id=0)
        sglab1_group_rule = tencentcloud.security.GroupRule("sglab1GroupRule",
            security_group_id=sglab1_group.id,
            type="ingress",
            cidr_ip="10.0.0.0/16",
            ip_protocol="TCP",
            port_range="80",
            policy="ACCEPT",
            description="favourite sg rule_1")
        ```

        Source is a security group id

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sglab2_group = tencentcloud.security.Group("sglab2Group",
            description="favourite sg_2",
            project_id=0)
        sglab3 = tencentcloud.security.Group("sglab3",
            description="favourite sg_3",
            project_id=0)
        sglab2_group_rule = tencentcloud.security.GroupRule("sglab2GroupRule",
            security_group_id=sglab2_group.id,
            type="ingress",
            ip_protocol="TCP",
            port_range="80",
            policy="ACCEPT",
            source_sgid=sglab3.id,
            description="favourite sg rule_2")
        ```

        :param str resource_name: The name of the resource.
        :param GroupRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_template: Optional[pulumi.Input[pulumi.InputType['GroupRuleAddressTemplateArgs']]] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 port_range: Optional[pulumi.Input[str]] = None,
                 protocol_template: Optional[pulumi.Input[pulumi.InputType['GroupRuleProtocolTemplateArgs']]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 source_sgid: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = GroupRuleArgs.__new__(GroupRuleArgs)

            __props__.__dict__["address_template"] = address_template
            __props__.__dict__["cidr_ip"] = cidr_ip
            __props__.__dict__["description"] = description
            __props__.__dict__["ip_protocol"] = ip_protocol
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["port_range"] = port_range
            __props__.__dict__["protocol_template"] = protocol_template
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
            __props__.__dict__["source_sgid"] = source_sgid
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(GroupRule, __self__).__init__(
            'tencentcloud:Security/groupRule:GroupRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_template: Optional[pulumi.Input[pulumi.InputType['GroupRuleAddressTemplateArgs']]] = None,
            cidr_ip: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_protocol: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            port_range: Optional[pulumi.Input[str]] = None,
            protocol_template: Optional[pulumi.Input[pulumi.InputType['GroupRuleProtocolTemplateArgs']]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            source_sgid: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'GroupRule':
        """
        Get an existing GroupRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GroupRuleAddressTemplateArgs']] address_template: ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        :param pulumi.Input[str] cidr_ip: An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        :param pulumi.Input[str] description: Description of the security group rule.
        :param pulumi.Input[str] ip_protocol: Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        :param pulumi.Input[str] policy: Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] port_range: Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        :param pulumi.Input[pulumi.InputType['GroupRuleProtocolTemplateArgs']] protocol_template: ID of the address template, and conflict with `ip_protocol`, `port_range`.
        :param pulumi.Input[str] security_group_id: ID of the security group to be queried.
        :param pulumi.Input[str] source_sgid: ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        :param pulumi.Input[str] type: Type of the security group rule. Valid values: `ingress` and `egress`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupRuleState.__new__(_GroupRuleState)

        __props__.__dict__["address_template"] = address_template
        __props__.__dict__["cidr_ip"] = cidr_ip
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_protocol"] = ip_protocol
        __props__.__dict__["policy"] = policy
        __props__.__dict__["port_range"] = port_range
        __props__.__dict__["protocol_template"] = protocol_template
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["source_sgid"] = source_sgid
        __props__.__dict__["type"] = type
        return GroupRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressTemplate")
    def address_template(self) -> pulumi.Output['outputs.GroupRuleAddressTemplate']:
        """
        ID of the address template, and confilicts with `source_sgid` and `cidr_ip`.
        """
        return pulumi.get(self, "address_template")

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> pulumi.Output[Optional[str]]:
        """
        An IP address network or segment, and conflict with `source_sgid` and `address_template`.
        """
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the security group rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> pulumi.Output[str]:
        """
        Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocol_template`.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> pulumi.Output[str]:
        """
        Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocol_template`.
        """
        return pulumi.get(self, "port_range")

    @property
    @pulumi.getter(name="protocolTemplate")
    def protocol_template(self) -> pulumi.Output['outputs.GroupRuleProtocolTemplate']:
        """
        ID of the address template, and conflict with `ip_protocol`, `port_range`.
        """
        return pulumi.get(self, "protocol_template")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        ID of the security group to be queried.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="sourceSgid")
    def source_sgid(self) -> pulumi.Output[str]:
        """
        ID of the nested security group, and conflicts with `cidr_ip` and `address_template`.
        """
        return pulumi.get(self, "source_sgid")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the security group rule. Valid values: `ingress` and `egress`.
        """
        return pulumi.get(self, "type")

