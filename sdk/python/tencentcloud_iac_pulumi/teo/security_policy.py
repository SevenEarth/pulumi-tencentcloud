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

__all__ = ['SecurityPolicyArgs', 'SecurityPolicy']

@pulumi.input_type
class SecurityPolicyArgs:
    def __init__(__self__, *,
                 entity: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 config: Optional[pulumi.Input['SecurityPolicyConfigArgs']] = None):
        """
        The set of arguments for constructing a SecurityPolicy resource.
        :param pulumi.Input[str] entity: Subdomain.
        :param pulumi.Input[str] zone_id: Site ID.
        :param pulumi.Input['SecurityPolicyConfigArgs'] config: Security policy configuration.
        """
        pulumi.set(__self__, "entity", entity)
        pulumi.set(__self__, "zone_id", zone_id)
        if config is not None:
            pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter
    def entity(self) -> pulumi.Input[str]:
        """
        Subdomain.
        """
        return pulumi.get(self, "entity")

    @entity.setter
    def entity(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['SecurityPolicyConfigArgs']]:
        """
        Security policy configuration.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['SecurityPolicyConfigArgs']]):
        pulumi.set(self, "config", value)


@pulumi.input_type
class _SecurityPolicyState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input['SecurityPolicyConfigArgs']] = None,
                 entity: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityPolicy resources.
        :param pulumi.Input['SecurityPolicyConfigArgs'] config: Security policy configuration.
        :param pulumi.Input[str] entity: Subdomain.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if entity is not None:
            pulumi.set(__self__, "entity", entity)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['SecurityPolicyConfigArgs']]:
        """
        Security policy configuration.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['SecurityPolicyConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def entity(self) -> Optional[pulumi.Input[str]]:
        """
        Subdomain.
        """
        return pulumi.get(self, "entity")

    @entity.setter
    def entity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class SecurityPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyConfigArgs']]] = None,
                 entity: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SecurityPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SecurityPolicyConfigArgs']] config: Security policy configuration.
        :param pulumi.Input[str] entity: Subdomain.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecurityPolicy resource with the given unique name, props, and options.
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
                 config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyConfigArgs']]] = None,
                 entity: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SecurityPolicyArgs.__new__(SecurityPolicyArgs)

            __props__.__dict__["config"] = config
            if entity is None and not opts.urn:
                raise TypeError("Missing required property 'entity'")
            __props__.__dict__["entity"] = entity
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(SecurityPolicy, __self__).__init__(
            'tencentcloud:Teo/securityPolicy:SecurityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[pulumi.InputType['SecurityPolicyConfigArgs']]] = None,
            entity: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'SecurityPolicy':
        """
        Get an existing SecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SecurityPolicyConfigArgs']] config: Security policy configuration.
        :param pulumi.Input[str] entity: Subdomain.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityPolicyState.__new__(_SecurityPolicyState)

        __props__.__dict__["config"] = config
        __props__.__dict__["entity"] = entity
        __props__.__dict__["zone_id"] = zone_id
        return SecurityPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.SecurityPolicyConfig']:
        """
        Security policy configuration.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def entity(self) -> pulumi.Output[str]:
        """
        Subdomain.
        """
        return pulumi.get(self, "entity")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

