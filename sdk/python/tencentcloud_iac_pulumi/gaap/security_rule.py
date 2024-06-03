# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecurityRuleArgs', 'SecurityRule']

@pulumi.input_type
class SecurityRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 cidr_ip: pulumi.Input[str],
                 policy_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityRule resource.
        :param pulumi.Input[str] action: Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] cidr_ip: A network address block of the request source.
        :param pulumi.Input[str] policy_id: ID of the security policy.
        :param pulumi.Input[str] name: Name of the security policy rule. Maximum length is 30.
        :param pulumi.Input[str] port: Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        :param pulumi.Input[str] protocol: Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "cidr_ip", cidr_ip)
        pulumi.set(__self__, "policy_id", policy_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> pulumi.Input[str]:
        """
        A network address block of the request source.
        """
        return pulumi.get(self, "cidr_ip")

    @cidr_ip.setter
    def cidr_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr_ip", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        ID of the security policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the security policy rule. Maximum length is 30.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


@pulumi.input_type
class _SecurityRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityRule resources.
        :param pulumi.Input[str] action: Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] cidr_ip: A network address block of the request source.
        :param pulumi.Input[str] name: Name of the security policy rule. Maximum length is 30.
        :param pulumi.Input[str] policy_id: ID of the security policy.
        :param pulumi.Input[str] port: Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        :param pulumi.Input[str] protocol: Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if cidr_ip is not None:
            pulumi.set(__self__, "cidr_ip", cidr_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> Optional[pulumi.Input[str]]:
        """
        A network address block of the request source.
        """
        return pulumi.get(self, "cidr_ip")

    @cidr_ip.setter
    def cidr_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the security policy rule. Maximum length is 30.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the security policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


class SecurityRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a security policy rule.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
            bandwidth=10,
            concurrent=2,
            access_region="SouthChina",
            realserver_region="NorthChina")
        foo_security_policy = tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy",
            proxy_id=foo_proxy.id,
            action="ACCEPT")
        foo_security_rule = tencentcloud.gaap.SecurityRule("fooSecurityRule",
            policy_id=foo_security_policy.id,
            cidr_ip="1.1.1.1",
            action="ACCEPT",
            protocol="TCP")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GAAP security rule can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Gaap/securityRule:SecurityRule tencentcloud_gaap_security_rule.foo sr-xxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] cidr_ip: A network address block of the request source.
        :param pulumi.Input[str] name: Name of the security policy rule. Maximum length is 30.
        :param pulumi.Input[str] policy_id: ID of the security policy.
        :param pulumi.Input[str] port: Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        :param pulumi.Input[str] protocol: Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a security policy rule.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
            bandwidth=10,
            concurrent=2,
            access_region="SouthChina",
            realserver_region="NorthChina")
        foo_security_policy = tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy",
            proxy_id=foo_proxy.id,
            action="ACCEPT")
        foo_security_rule = tencentcloud.gaap.SecurityRule("fooSecurityRule",
            policy_id=foo_security_policy.id,
            cidr_ip="1.1.1.1",
            action="ACCEPT",
            protocol="TCP")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GAAP security rule can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Gaap/securityRule:SecurityRule tencentcloud_gaap_security_rule.foo sr-xxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param SecurityRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 cidr_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityRuleArgs.__new__(SecurityRuleArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            if cidr_ip is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_ip'")
            __props__.__dict__["cidr_ip"] = cidr_ip
            __props__.__dict__["name"] = name
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["port"] = port
            __props__.__dict__["protocol"] = protocol
        super(SecurityRule, __self__).__init__(
            'tencentcloud:Gaap/securityRule:SecurityRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            cidr_ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None) -> 'SecurityRule':
        """
        Get an existing SecurityRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] cidr_ip: A network address block of the request source.
        :param pulumi.Input[str] name: Name of the security policy rule. Maximum length is 30.
        :param pulumi.Input[str] policy_id: ID of the security policy.
        :param pulumi.Input[str] port: Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        :param pulumi.Input[str] protocol: Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityRuleState.__new__(_SecurityRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["cidr_ip"] = cidr_ip
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["port"] = port
        __props__.__dict__["protocol"] = protocol
        return SecurityRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> pulumi.Output[str]:
        """
        A network address block of the request source.
        """
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the security policy rule. Maximum length is 30.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        ID of the security policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[str]]:
        """
        Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        """
        Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        """
        return pulumi.get(self, "protocol")

