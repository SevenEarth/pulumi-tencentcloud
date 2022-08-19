# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecurityPolicyArgs', 'SecurityPolicy']

@pulumi.input_type
class SecurityPolicyArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 proxy_id: pulumi.Input[str],
                 enable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SecurityPolicy resource.
        :param pulumi.Input[str] action: Default policy. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[str] proxy_id: ID of the GAAP proxy.
        :param pulumi.Input[bool] enable: Indicates whether policy is enable, default value is `true`.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "proxy_id", proxy_id)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Default policy. Valid value: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> pulumi.Input[str]:
        """
        ID of the GAAP proxy.
        """
        return pulumi.get(self, "proxy_id")

    @proxy_id.setter
    def proxy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "proxy_id", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether policy is enable, default value is `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)


@pulumi.input_type
class _SecurityPolicyState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 proxy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityPolicy resources.
        :param pulumi.Input[str] action: Default policy. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[bool] enable: Indicates whether policy is enable, default value is `true`.
        :param pulumi.Input[str] proxy_id: ID of the GAAP proxy.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if proxy_id is not None:
            pulumi.set(__self__, "proxy_id", proxy_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Default policy. Valid value: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether policy is enable, default value is `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the GAAP proxy.
        """
        return pulumi.get(self, "proxy_id")

    @proxy_id.setter
    def proxy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_id", value)


class SecurityPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 proxy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a security policy of GAAP proxy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
            bandwidth=10,
            concurrent=2,
            access_region="SouthChina",
            realserver_region="NorthChina")
        foo_security_policy = tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy",
            proxy_id=foo_proxy.id,
            action="DROP")
        ```

        ## Import

        GAAP security policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Gaap/securityPolicy:SecurityPolicy tencentcloud_gaap_security_policy.foo pl-xxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Default policy. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[bool] enable: Indicates whether policy is enable, default value is `true`.
        :param pulumi.Input[str] proxy_id: ID of the GAAP proxy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a security policy of GAAP proxy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
            bandwidth=10,
            concurrent=2,
            access_region="SouthChina",
            realserver_region="NorthChina")
        foo_security_policy = tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy",
            proxy_id=foo_proxy.id,
            action="DROP")
        ```

        ## Import

        GAAP security policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Gaap/securityPolicy:SecurityPolicy tencentcloud_gaap_security_policy.foo pl-xxxx
        ```

        :param str resource_name: The name of the resource.
        :param SecurityPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 proxy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityPolicyArgs.__new__(SecurityPolicyArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["enable"] = enable
            if proxy_id is None and not opts.urn:
                raise TypeError("Missing required property 'proxy_id'")
            __props__.__dict__["proxy_id"] = proxy_id
        super(SecurityPolicy, __self__).__init__(
            'tencentcloud:Gaap/securityPolicy:SecurityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            proxy_id: Optional[pulumi.Input[str]] = None) -> 'SecurityPolicy':
        """
        Get an existing SecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Default policy. Valid value: `ACCEPT` and `DROP`.
        :param pulumi.Input[bool] enable: Indicates whether policy is enable, default value is `true`.
        :param pulumi.Input[str] proxy_id: ID of the GAAP proxy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityPolicyState.__new__(_SecurityPolicyState)

        __props__.__dict__["action"] = action
        __props__.__dict__["enable"] = enable
        __props__.__dict__["proxy_id"] = proxy_id
        return SecurityPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Default policy. Valid value: `ACCEPT` and `DROP`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether policy is enable, default value is `true`.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> pulumi.Output[str]:
        """
        ID of the GAAP proxy.
        """
        return pulumi.get(self, "proxy_id")

