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

__all__ = ['ProvisionedConcurrencyConfigArgs', 'ProvisionedConcurrencyConfig']

@pulumi.input_type
class ProvisionedConcurrencyConfigArgs:
    def __init__(__self__, *,
                 function_name: pulumi.Input[str],
                 qualifier: pulumi.Input[str],
                 version_provisioned_concurrency_num: pulumi.Input[int],
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 provisioned_type: Optional[pulumi.Input[str]] = None,
                 tracking_target: Optional[pulumi.Input[float]] = None,
                 trigger_actions: Optional[pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]]] = None):
        """
        The set of arguments for constructing a ProvisionedConcurrencyConfig resource.
        :param pulumi.Input[str] function_name: Name of the function for which to set the provisioned concurrency.
        :param pulumi.Input[str] qualifier: Function version number. Note: the $LATEST version does not support provisioned concurrency.
        :param pulumi.Input[int] version_provisioned_concurrency_num: Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        :param pulumi.Input[int] max_capacity: The maximum number of instances.
        :param pulumi.Input[int] min_capacity: The minimum number of instances. It can not be smaller than 1.
        :param pulumi.Input[str] namespace: Function namespace. Default value: default.
        :param pulumi.Input[str] provisioned_type: Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        :param pulumi.Input[float] tracking_target: The target concurrency utilization. Range: (0,1) (two decimal places).
        :param pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]] trigger_actions: Scheduled provisioned concurrency scaling action.
        """
        pulumi.set(__self__, "function_name", function_name)
        pulumi.set(__self__, "qualifier", qualifier)
        pulumi.set(__self__, "version_provisioned_concurrency_num", version_provisioned_concurrency_num)
        if max_capacity is not None:
            pulumi.set(__self__, "max_capacity", max_capacity)
        if min_capacity is not None:
            pulumi.set(__self__, "min_capacity", min_capacity)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if provisioned_type is not None:
            pulumi.set(__self__, "provisioned_type", provisioned_type)
        if tracking_target is not None:
            pulumi.set(__self__, "tracking_target", tracking_target)
        if trigger_actions is not None:
            pulumi.set(__self__, "trigger_actions", trigger_actions)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Input[str]:
        """
        Name of the function for which to set the provisioned concurrency.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter
    def qualifier(self) -> pulumi.Input[str]:
        """
        Function version number. Note: the $LATEST version does not support provisioned concurrency.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "qualifier", value)

    @property
    @pulumi.getter(name="versionProvisionedConcurrencyNum")
    def version_provisioned_concurrency_num(self) -> pulumi.Input[int]:
        """
        Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        """
        return pulumi.get(self, "version_provisioned_concurrency_num")

    @version_provisioned_concurrency_num.setter
    def version_provisioned_concurrency_num(self, value: pulumi.Input[int]):
        pulumi.set(self, "version_provisioned_concurrency_num", value)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of instances.
        """
        return pulumi.get(self, "max_capacity")

    @max_capacity.setter
    def max_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_capacity", value)

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of instances. It can not be smaller than 1.
        """
        return pulumi.get(self, "min_capacity")

    @min_capacity.setter
    def min_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_capacity", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Function namespace. Default value: default.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="provisionedType")
    def provisioned_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        """
        return pulumi.get(self, "provisioned_type")

    @provisioned_type.setter
    def provisioned_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioned_type", value)

    @property
    @pulumi.getter(name="trackingTarget")
    def tracking_target(self) -> Optional[pulumi.Input[float]]:
        """
        The target concurrency utilization. Range: (0,1) (two decimal places).
        """
        return pulumi.get(self, "tracking_target")

    @tracking_target.setter
    def tracking_target(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "tracking_target", value)

    @property
    @pulumi.getter(name="triggerActions")
    def trigger_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]]]:
        """
        Scheduled provisioned concurrency scaling action.
        """
        return pulumi.get(self, "trigger_actions")

    @trigger_actions.setter
    def trigger_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]]]):
        pulumi.set(self, "trigger_actions", value)


@pulumi.input_type
class _ProvisionedConcurrencyConfigState:
    def __init__(__self__, *,
                 function_name: Optional[pulumi.Input[str]] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 provisioned_type: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 tracking_target: Optional[pulumi.Input[float]] = None,
                 trigger_actions: Optional[pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]]] = None,
                 version_provisioned_concurrency_num: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProvisionedConcurrencyConfig resources.
        :param pulumi.Input[str] function_name: Name of the function for which to set the provisioned concurrency.
        :param pulumi.Input[int] max_capacity: The maximum number of instances.
        :param pulumi.Input[int] min_capacity: The minimum number of instances. It can not be smaller than 1.
        :param pulumi.Input[str] namespace: Function namespace. Default value: default.
        :param pulumi.Input[str] provisioned_type: Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        :param pulumi.Input[str] qualifier: Function version number. Note: the $LATEST version does not support provisioned concurrency.
        :param pulumi.Input[float] tracking_target: The target concurrency utilization. Range: (0,1) (two decimal places).
        :param pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]] trigger_actions: Scheduled provisioned concurrency scaling action.
        :param pulumi.Input[int] version_provisioned_concurrency_num: Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        """
        if function_name is not None:
            pulumi.set(__self__, "function_name", function_name)
        if max_capacity is not None:
            pulumi.set(__self__, "max_capacity", max_capacity)
        if min_capacity is not None:
            pulumi.set(__self__, "min_capacity", min_capacity)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if provisioned_type is not None:
            pulumi.set(__self__, "provisioned_type", provisioned_type)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)
        if tracking_target is not None:
            pulumi.set(__self__, "tracking_target", tracking_target)
        if trigger_actions is not None:
            pulumi.set(__self__, "trigger_actions", trigger_actions)
        if version_provisioned_concurrency_num is not None:
            pulumi.set(__self__, "version_provisioned_concurrency_num", version_provisioned_concurrency_num)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the function for which to set the provisioned concurrency.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of instances.
        """
        return pulumi.get(self, "max_capacity")

    @max_capacity.setter
    def max_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_capacity", value)

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of instances. It can not be smaller than 1.
        """
        return pulumi.get(self, "min_capacity")

    @min_capacity.setter
    def min_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_capacity", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Function namespace. Default value: default.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="provisionedType")
    def provisioned_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        """
        return pulumi.get(self, "provisioned_type")

    @provisioned_type.setter
    def provisioned_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provisioned_type", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        Function version number. Note: the $LATEST version does not support provisioned concurrency.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)

    @property
    @pulumi.getter(name="trackingTarget")
    def tracking_target(self) -> Optional[pulumi.Input[float]]:
        """
        The target concurrency utilization. Range: (0,1) (two decimal places).
        """
        return pulumi.get(self, "tracking_target")

    @tracking_target.setter
    def tracking_target(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "tracking_target", value)

    @property
    @pulumi.getter(name="triggerActions")
    def trigger_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]]]:
        """
        Scheduled provisioned concurrency scaling action.
        """
        return pulumi.get(self, "trigger_actions")

    @trigger_actions.setter
    def trigger_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProvisionedConcurrencyConfigTriggerActionArgs']]]]):
        pulumi.set(self, "trigger_actions", value)

    @property
    @pulumi.getter(name="versionProvisionedConcurrencyNum")
    def version_provisioned_concurrency_num(self) -> Optional[pulumi.Input[int]]:
        """
        Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        """
        return pulumi.get(self, "version_provisioned_concurrency_num")

    @version_provisioned_concurrency_num.setter
    def version_provisioned_concurrency_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version_provisioned_concurrency_num", value)


class ProvisionedConcurrencyConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 provisioned_type: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 tracking_target: Optional[pulumi.Input[float]] = None,
                 trigger_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProvisionedConcurrencyConfigTriggerActionArgs']]]]] = None,
                 version_provisioned_concurrency_num: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a scf provisioned_concurrency_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        provisioned_concurrency_config = tencentcloud.scf.ProvisionedConcurrencyConfig("provisionedConcurrencyConfig",
            function_name="keep-1676351130",
            max_capacity=2,
            min_capacity=1,
            namespace="default",
            provisioned_type="Default",
            qualifier="2",
            tracking_target=0.5,
            trigger_actions=[tencentcloud.scf.ProvisionedConcurrencyConfigTriggerActionArgs(
                provisioned_type="Default",
                trigger_cron_config="29 45 12 29 05 * 2023",
                trigger_name="test",
                trigger_provisioned_concurrency_num=2,
            )],
            version_provisioned_concurrency_num=2)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_name: Name of the function for which to set the provisioned concurrency.
        :param pulumi.Input[int] max_capacity: The maximum number of instances.
        :param pulumi.Input[int] min_capacity: The minimum number of instances. It can not be smaller than 1.
        :param pulumi.Input[str] namespace: Function namespace. Default value: default.
        :param pulumi.Input[str] provisioned_type: Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        :param pulumi.Input[str] qualifier: Function version number. Note: the $LATEST version does not support provisioned concurrency.
        :param pulumi.Input[float] tracking_target: The target concurrency utilization. Range: (0,1) (two decimal places).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProvisionedConcurrencyConfigTriggerActionArgs']]]] trigger_actions: Scheduled provisioned concurrency scaling action.
        :param pulumi.Input[int] version_provisioned_concurrency_num: Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProvisionedConcurrencyConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a scf provisioned_concurrency_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        provisioned_concurrency_config = tencentcloud.scf.ProvisionedConcurrencyConfig("provisionedConcurrencyConfig",
            function_name="keep-1676351130",
            max_capacity=2,
            min_capacity=1,
            namespace="default",
            provisioned_type="Default",
            qualifier="2",
            tracking_target=0.5,
            trigger_actions=[tencentcloud.scf.ProvisionedConcurrencyConfigTriggerActionArgs(
                provisioned_type="Default",
                trigger_cron_config="29 45 12 29 05 * 2023",
                trigger_name="test",
                trigger_provisioned_concurrency_num=2,
            )],
            version_provisioned_concurrency_num=2)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ProvisionedConcurrencyConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProvisionedConcurrencyConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 max_capacity: Optional[pulumi.Input[int]] = None,
                 min_capacity: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 provisioned_type: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 tracking_target: Optional[pulumi.Input[float]] = None,
                 trigger_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProvisionedConcurrencyConfigTriggerActionArgs']]]]] = None,
                 version_provisioned_concurrency_num: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProvisionedConcurrencyConfigArgs.__new__(ProvisionedConcurrencyConfigArgs)

            if function_name is None and not opts.urn:
                raise TypeError("Missing required property 'function_name'")
            __props__.__dict__["function_name"] = function_name
            __props__.__dict__["max_capacity"] = max_capacity
            __props__.__dict__["min_capacity"] = min_capacity
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["provisioned_type"] = provisioned_type
            if qualifier is None and not opts.urn:
                raise TypeError("Missing required property 'qualifier'")
            __props__.__dict__["qualifier"] = qualifier
            __props__.__dict__["tracking_target"] = tracking_target
            __props__.__dict__["trigger_actions"] = trigger_actions
            if version_provisioned_concurrency_num is None and not opts.urn:
                raise TypeError("Missing required property 'version_provisioned_concurrency_num'")
            __props__.__dict__["version_provisioned_concurrency_num"] = version_provisioned_concurrency_num
        super(ProvisionedConcurrencyConfig, __self__).__init__(
            'tencentcloud:Scf/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            function_name: Optional[pulumi.Input[str]] = None,
            max_capacity: Optional[pulumi.Input[int]] = None,
            min_capacity: Optional[pulumi.Input[int]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            provisioned_type: Optional[pulumi.Input[str]] = None,
            qualifier: Optional[pulumi.Input[str]] = None,
            tracking_target: Optional[pulumi.Input[float]] = None,
            trigger_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProvisionedConcurrencyConfigTriggerActionArgs']]]]] = None,
            version_provisioned_concurrency_num: Optional[pulumi.Input[int]] = None) -> 'ProvisionedConcurrencyConfig':
        """
        Get an existing ProvisionedConcurrencyConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_name: Name of the function for which to set the provisioned concurrency.
        :param pulumi.Input[int] max_capacity: The maximum number of instances.
        :param pulumi.Input[int] min_capacity: The minimum number of instances. It can not be smaller than 1.
        :param pulumi.Input[str] namespace: Function namespace. Default value: default.
        :param pulumi.Input[str] provisioned_type: Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        :param pulumi.Input[str] qualifier: Function version number. Note: the $LATEST version does not support provisioned concurrency.
        :param pulumi.Input[float] tracking_target: The target concurrency utilization. Range: (0,1) (two decimal places).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProvisionedConcurrencyConfigTriggerActionArgs']]]] trigger_actions: Scheduled provisioned concurrency scaling action.
        :param pulumi.Input[int] version_provisioned_concurrency_num: Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProvisionedConcurrencyConfigState.__new__(_ProvisionedConcurrencyConfigState)

        __props__.__dict__["function_name"] = function_name
        __props__.__dict__["max_capacity"] = max_capacity
        __props__.__dict__["min_capacity"] = min_capacity
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["provisioned_type"] = provisioned_type
        __props__.__dict__["qualifier"] = qualifier
        __props__.__dict__["tracking_target"] = tracking_target
        __props__.__dict__["trigger_actions"] = trigger_actions
        __props__.__dict__["version_provisioned_concurrency_num"] = version_provisioned_concurrency_num
        return ProvisionedConcurrencyConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        """
        Name of the function for which to set the provisioned concurrency.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum number of instances.
        """
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum number of instances. It can not be smaller than 1.
        """
        return pulumi.get(self, "min_capacity")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Function namespace. Default value: default.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="provisionedType")
    def provisioned_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
        """
        return pulumi.get(self, "provisioned_type")

    @property
    @pulumi.getter
    def qualifier(self) -> pulumi.Output[str]:
        """
        Function version number. Note: the $LATEST version does not support provisioned concurrency.
        """
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter(name="trackingTarget")
    def tracking_target(self) -> pulumi.Output[Optional[float]]:
        """
        The target concurrency utilization. Range: (0,1) (two decimal places).
        """
        return pulumi.get(self, "tracking_target")

    @property
    @pulumi.getter(name="triggerActions")
    def trigger_actions(self) -> pulumi.Output[Optional[Sequence['outputs.ProvisionedConcurrencyConfigTriggerAction']]]:
        """
        Scheduled provisioned concurrency scaling action.
        """
        return pulumi.get(self, "trigger_actions")

    @property
    @pulumi.getter(name="versionProvisionedConcurrencyNum")
    def version_provisioned_concurrency_num(self) -> pulumi.Output[int]:
        """
        Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
        """
        return pulumi.get(self, "version_provisioned_concurrency_num")

