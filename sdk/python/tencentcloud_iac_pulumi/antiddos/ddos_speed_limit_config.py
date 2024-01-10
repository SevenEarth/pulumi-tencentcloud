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

__all__ = ['DdosSpeedLimitConfigArgs', 'DdosSpeedLimitConfig']

@pulumi.input_type
class DdosSpeedLimitConfigArgs:
    def __init__(__self__, *,
                 ddos_speed_limit_config: pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs'],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DdosSpeedLimitConfig resource.
        :param pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs'] ddos_speed_limit_config: Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        :param pulumi.Input[str] instance_id: InstanceId.
        """
        pulumi.set(__self__, "ddos_speed_limit_config", ddos_speed_limit_config)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="ddosSpeedLimitConfig")
    def ddos_speed_limit_config(self) -> pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']:
        """
        Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        """
        return pulumi.get(self, "ddos_speed_limit_config")

    @ddos_speed_limit_config.setter
    def ddos_speed_limit_config(self, value: pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']):
        pulumi.set(self, "ddos_speed_limit_config", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        InstanceId.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _DdosSpeedLimitConfigState:
    def __init__(__self__, *,
                 ddos_speed_limit_config: Optional[pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DdosSpeedLimitConfig resources.
        :param pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs'] ddos_speed_limit_config: Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        :param pulumi.Input[str] instance_id: InstanceId.
        """
        if ddos_speed_limit_config is not None:
            pulumi.set(__self__, "ddos_speed_limit_config", ddos_speed_limit_config)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="ddosSpeedLimitConfig")
    def ddos_speed_limit_config(self) -> Optional[pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']]:
        """
        Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        """
        return pulumi.get(self, "ddos_speed_limit_config")

    @ddos_speed_limit_config.setter
    def ddos_speed_limit_config(self, value: Optional[pulumi.Input['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']]):
        pulumi.set(self, "ddos_speed_limit_config", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        InstanceId.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class DdosSpeedLimitConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ddos_speed_limit_config: Optional[pulumi.Input[pulumi.InputType['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a antiddos ddos speed limit config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        ddos_speed_limit_config = tencentcloud.antiddos.DdosSpeedLimitConfig("ddosSpeedLimitConfig",
            ddos_speed_limit_config=tencentcloud.antiddos.DdosSpeedLimitConfigDdosSpeedLimitConfigArgs(
                dst_port_list="8000",
                mode=1,
                protocol_list="ALL",
                speed_values=[
                    tencentcloud.antiddos.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueArgs(
                        type=1,
                        value=1,
                    ),
                    tencentcloud.antiddos.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueArgs(
                        type=2,
                        value=2,
                    ),
                ],
            ),
            instance_id="bgp-xxxxxx")
        ```

        ## Import

        antiddos ddos_speed_limit_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig ddos_speed_limit_config ${instanceId}#${configId}s
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']] ddos_speed_limit_config: Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        :param pulumi.Input[str] instance_id: InstanceId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DdosSpeedLimitConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a antiddos ddos speed limit config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        ddos_speed_limit_config = tencentcloud.antiddos.DdosSpeedLimitConfig("ddosSpeedLimitConfig",
            ddos_speed_limit_config=tencentcloud.antiddos.DdosSpeedLimitConfigDdosSpeedLimitConfigArgs(
                dst_port_list="8000",
                mode=1,
                protocol_list="ALL",
                speed_values=[
                    tencentcloud.antiddos.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueArgs(
                        type=1,
                        value=1,
                    ),
                    tencentcloud.antiddos.DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValueArgs(
                        type=2,
                        value=2,
                    ),
                ],
            ),
            instance_id="bgp-xxxxxx")
        ```

        ## Import

        antiddos ddos_speed_limit_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig ddos_speed_limit_config ${instanceId}#${configId}s
        ```

        :param str resource_name: The name of the resource.
        :param DdosSpeedLimitConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DdosSpeedLimitConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ddos_speed_limit_config: Optional[pulumi.Input[pulumi.InputType['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DdosSpeedLimitConfigArgs.__new__(DdosSpeedLimitConfigArgs)

            if ddos_speed_limit_config is None and not opts.urn:
                raise TypeError("Missing required property 'ddos_speed_limit_config'")
            __props__.__dict__["ddos_speed_limit_config"] = ddos_speed_limit_config
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(DdosSpeedLimitConfig, __self__).__init__(
            'tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ddos_speed_limit_config: Optional[pulumi.Input[pulumi.InputType['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'DdosSpeedLimitConfig':
        """
        Get an existing DdosSpeedLimitConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DdosSpeedLimitConfigDdosSpeedLimitConfigArgs']] ddos_speed_limit_config: Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        :param pulumi.Input[str] instance_id: InstanceId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DdosSpeedLimitConfigState.__new__(_DdosSpeedLimitConfigState)

        __props__.__dict__["ddos_speed_limit_config"] = ddos_speed_limit_config
        __props__.__dict__["instance_id"] = instance_id
        return DdosSpeedLimitConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ddosSpeedLimitConfig")
    def ddos_speed_limit_config(self) -> pulumi.Output['outputs.DdosSpeedLimitConfigDdosSpeedLimitConfig']:
        """
        Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        """
        return pulumi.get(self, "ddos_speed_limit_config")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        InstanceId.
        """
        return pulumi.get(self, "instance_id")

