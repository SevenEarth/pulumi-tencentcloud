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

__all__ = ['DefaultAlarmThresholdArgs', 'DefaultAlarmThreshold']

@pulumi.input_type
class DefaultAlarmThresholdArgs:
    def __init__(__self__, *,
                 default_alarm_config: pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs'],
                 instance_type: pulumi.Input[str]):
        """
        The set of arguments for constructing a DefaultAlarmThreshold resource.
        :param pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs'] default_alarm_config: Alarm threshold configuration.
        :param pulumi.Input[str] instance_type: Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        pulumi.set(__self__, "default_alarm_config", default_alarm_config)
        pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="defaultAlarmConfig")
    def default_alarm_config(self) -> pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs']:
        """
        Alarm threshold configuration.
        """
        return pulumi.get(self, "default_alarm_config")

    @default_alarm_config.setter
    def default_alarm_config(self, value: pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs']):
        pulumi.set(self, "default_alarm_config", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class _DefaultAlarmThresholdState:
    def __init__(__self__, *,
                 default_alarm_config: Optional[pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs']] = None,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DefaultAlarmThreshold resources.
        :param pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs'] default_alarm_config: Alarm threshold configuration.
        :param pulumi.Input[str] instance_type: Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        if default_alarm_config is not None:
            pulumi.set(__self__, "default_alarm_config", default_alarm_config)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="defaultAlarmConfig")
    def default_alarm_config(self) -> Optional[pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs']]:
        """
        Alarm threshold configuration.
        """
        return pulumi.get(self, "default_alarm_config")

    @default_alarm_config.setter
    def default_alarm_config(self, value: Optional[pulumi.Input['DefaultAlarmThresholdDefaultAlarmConfigArgs']]):
        pulumi.set(self, "default_alarm_config", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


class DefaultAlarmThreshold(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_alarm_config: Optional[pulumi.Input[pulumi.InputType['DefaultAlarmThresholdDefaultAlarmConfigArgs']]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a antiddos default alarm threshold

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        default_alarm_threshold = tencentcloud.antiddos.DefaultAlarmThreshold("defaultAlarmThreshold",
            default_alarm_config=tencentcloud.antiddos.DefaultAlarmThresholdDefaultAlarmConfigArgs(
                alarm_threshold=2000,
                alarm_type=1,
            ),
            instance_type="bgp")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        antiddos default_alarm_threshold can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold default_alarm_threshold ${instanceType}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DefaultAlarmThresholdDefaultAlarmConfigArgs']] default_alarm_config: Alarm threshold configuration.
        :param pulumi.Input[str] instance_type: Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefaultAlarmThresholdArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a antiddos default alarm threshold

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        default_alarm_threshold = tencentcloud.antiddos.DefaultAlarmThreshold("defaultAlarmThreshold",
            default_alarm_config=tencentcloud.antiddos.DefaultAlarmThresholdDefaultAlarmConfigArgs(
                alarm_threshold=2000,
                alarm_type=1,
            ),
            instance_type="bgp")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        antiddos default_alarm_threshold can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold default_alarm_threshold ${instanceType}
        ```

        :param str resource_name: The name of the resource.
        :param DefaultAlarmThresholdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultAlarmThresholdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_alarm_config: Optional[pulumi.Input[pulumi.InputType['DefaultAlarmThresholdDefaultAlarmConfigArgs']]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultAlarmThresholdArgs.__new__(DefaultAlarmThresholdArgs)

            if default_alarm_config is None and not opts.urn:
                raise TypeError("Missing required property 'default_alarm_config'")
            __props__.__dict__["default_alarm_config"] = default_alarm_config
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
        super(DefaultAlarmThreshold, __self__).__init__(
            'tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_alarm_config: Optional[pulumi.Input[pulumi.InputType['DefaultAlarmThresholdDefaultAlarmConfigArgs']]] = None,
            instance_type: Optional[pulumi.Input[str]] = None) -> 'DefaultAlarmThreshold':
        """
        Get an existing DefaultAlarmThreshold resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DefaultAlarmThresholdDefaultAlarmConfigArgs']] default_alarm_config: Alarm threshold configuration.
        :param pulumi.Input[str] instance_type: Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultAlarmThresholdState.__new__(_DefaultAlarmThresholdState)

        __props__.__dict__["default_alarm_config"] = default_alarm_config
        __props__.__dict__["instance_type"] = instance_type
        return DefaultAlarmThreshold(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultAlarmConfig")
    def default_alarm_config(self) -> pulumi.Output['outputs.DefaultAlarmThresholdDefaultAlarmConfig']:
        """
        Alarm threshold configuration.
        """
        return pulumi.get(self, "default_alarm_config")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        """
        return pulumi.get(self, "instance_type")

