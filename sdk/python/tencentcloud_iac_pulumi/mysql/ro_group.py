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

__all__ = ['RoGroupArgs', 'RoGroup']

@pulumi.input_type
class RoGroupArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 ro_group_id: pulumi.Input[str],
                 is_balance_ro_load: Optional[pulumi.Input[int]] = None,
                 ro_group_info: Optional[pulumi.Input['RoGroupRoGroupInfoArgs']] = None,
                 ro_weight_values: Optional[pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]]] = None):
        """
        The set of arguments for constructing a RoGroup resource.
        :param pulumi.Input[str] instance_id: Instance ID, in the format: cdbro-3i70uj0k.
        :param pulumi.Input[str] ro_group_id: The ID of the RO group.
        :param pulumi.Input[int] is_balance_ro_load: Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        :param pulumi.Input['RoGroupRoGroupInfoArgs'] ro_group_info: Details of the RO group.
        :param pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]] ro_weight_values: The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "ro_group_id", ro_group_id)
        if is_balance_ro_load is not None:
            pulumi.set(__self__, "is_balance_ro_load", is_balance_ro_load)
        if ro_group_info is not None:
            pulumi.set(__self__, "ro_group_info", ro_group_info)
        if ro_weight_values is not None:
            pulumi.set(__self__, "ro_weight_values", ro_weight_values)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID, in the format: cdbro-3i70uj0k.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="roGroupId")
    def ro_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the RO group.
        """
        return pulumi.get(self, "ro_group_id")

    @ro_group_id.setter
    def ro_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ro_group_id", value)

    @property
    @pulumi.getter(name="isBalanceRoLoad")
    def is_balance_ro_load(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        """
        return pulumi.get(self, "is_balance_ro_load")

    @is_balance_ro_load.setter
    def is_balance_ro_load(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "is_balance_ro_load", value)

    @property
    @pulumi.getter(name="roGroupInfo")
    def ro_group_info(self) -> Optional[pulumi.Input['RoGroupRoGroupInfoArgs']]:
        """
        Details of the RO group.
        """
        return pulumi.get(self, "ro_group_info")

    @ro_group_info.setter
    def ro_group_info(self, value: Optional[pulumi.Input['RoGroupRoGroupInfoArgs']]):
        pulumi.set(self, "ro_group_info", value)

    @property
    @pulumi.getter(name="roWeightValues")
    def ro_weight_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]]]:
        """
        The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        return pulumi.get(self, "ro_weight_values")

    @ro_weight_values.setter
    def ro_weight_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]]]):
        pulumi.set(self, "ro_weight_values", value)


@pulumi.input_type
class _RoGroupState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_balance_ro_load: Optional[pulumi.Input[int]] = None,
                 ro_group_id: Optional[pulumi.Input[str]] = None,
                 ro_group_info: Optional[pulumi.Input['RoGroupRoGroupInfoArgs']] = None,
                 ro_weight_values: Optional[pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]]] = None):
        """
        Input properties used for looking up and filtering RoGroup resources.
        :param pulumi.Input[str] instance_id: Instance ID, in the format: cdbro-3i70uj0k.
        :param pulumi.Input[int] is_balance_ro_load: Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        :param pulumi.Input[str] ro_group_id: The ID of the RO group.
        :param pulumi.Input['RoGroupRoGroupInfoArgs'] ro_group_info: Details of the RO group.
        :param pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]] ro_weight_values: The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_balance_ro_load is not None:
            pulumi.set(__self__, "is_balance_ro_load", is_balance_ro_load)
        if ro_group_id is not None:
            pulumi.set(__self__, "ro_group_id", ro_group_id)
        if ro_group_info is not None:
            pulumi.set(__self__, "ro_group_info", ro_group_info)
        if ro_weight_values is not None:
            pulumi.set(__self__, "ro_weight_values", ro_weight_values)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID, in the format: cdbro-3i70uj0k.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isBalanceRoLoad")
    def is_balance_ro_load(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        """
        return pulumi.get(self, "is_balance_ro_load")

    @is_balance_ro_load.setter
    def is_balance_ro_load(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "is_balance_ro_load", value)

    @property
    @pulumi.getter(name="roGroupId")
    def ro_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the RO group.
        """
        return pulumi.get(self, "ro_group_id")

    @ro_group_id.setter
    def ro_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ro_group_id", value)

    @property
    @pulumi.getter(name="roGroupInfo")
    def ro_group_info(self) -> Optional[pulumi.Input['RoGroupRoGroupInfoArgs']]:
        """
        Details of the RO group.
        """
        return pulumi.get(self, "ro_group_info")

    @ro_group_info.setter
    def ro_group_info(self, value: Optional[pulumi.Input['RoGroupRoGroupInfoArgs']]):
        pulumi.set(self, "ro_group_info", value)

    @property
    @pulumi.getter(name="roWeightValues")
    def ro_weight_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]]]:
        """
        The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        return pulumi.get(self, "ro_weight_values")

    @ro_weight_values.setter
    def ro_weight_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RoGroupRoWeightValueArgs']]]]):
        pulumi.set(self, "ro_weight_values", value)


class RoGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_balance_ro_load: Optional[pulumi.Input[int]] = None,
                 ro_group_id: Optional[pulumi.Input[str]] = None,
                 ro_group_info: Optional[pulumi.Input[pulumi.InputType['RoGroupRoGroupInfoArgs']]] = None,
                 ro_weight_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoGroupRoWeightValueArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql ro_group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.mysql.RoGroup("example",
            instance_id="cdb-e8i766hx",
            is_balance_ro_load=1,
            ro_group_id="cdbrg-f49t0gnj",
            ro_group_info=tencentcloud.mysql.RoGroupRoGroupInfoArgs(
                min_ro_in_group=1,
                replication_delay_time=1,
                ro_group_name="keep-ro",
                ro_max_delay_time=1,
                ro_offline_delay=1,
                weight_mode="custom",
            ),
            ro_weight_values=[tencentcloud.mysql.RoGroupRoWeightValueArgs(
                instance_id="cdbro-f49t0gnj",
                weight=10,
            )])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID, in the format: cdbro-3i70uj0k.
        :param pulumi.Input[int] is_balance_ro_load: Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        :param pulumi.Input[str] ro_group_id: The ID of the RO group.
        :param pulumi.Input[pulumi.InputType['RoGroupRoGroupInfoArgs']] ro_group_info: Details of the RO group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoGroupRoWeightValueArgs']]]] ro_weight_values: The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql ro_group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.mysql.RoGroup("example",
            instance_id="cdb-e8i766hx",
            is_balance_ro_load=1,
            ro_group_id="cdbrg-f49t0gnj",
            ro_group_info=tencentcloud.mysql.RoGroupRoGroupInfoArgs(
                min_ro_in_group=1,
                replication_delay_time=1,
                ro_group_name="keep-ro",
                ro_max_delay_time=1,
                ro_offline_delay=1,
                weight_mode="custom",
            ),
            ro_weight_values=[tencentcloud.mysql.RoGroupRoWeightValueArgs(
                instance_id="cdbro-f49t0gnj",
                weight=10,
            )])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param RoGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_balance_ro_load: Optional[pulumi.Input[int]] = None,
                 ro_group_id: Optional[pulumi.Input[str]] = None,
                 ro_group_info: Optional[pulumi.Input[pulumi.InputType['RoGroupRoGroupInfoArgs']]] = None,
                 ro_weight_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoGroupRoWeightValueArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoGroupArgs.__new__(RoGroupArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["is_balance_ro_load"] = is_balance_ro_load
            if ro_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'ro_group_id'")
            __props__.__dict__["ro_group_id"] = ro_group_id
            __props__.__dict__["ro_group_info"] = ro_group_info
            __props__.__dict__["ro_weight_values"] = ro_weight_values
        super(RoGroup, __self__).__init__(
            'tencentcloud:Mysql/roGroup:RoGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            is_balance_ro_load: Optional[pulumi.Input[int]] = None,
            ro_group_id: Optional[pulumi.Input[str]] = None,
            ro_group_info: Optional[pulumi.Input[pulumi.InputType['RoGroupRoGroupInfoArgs']]] = None,
            ro_weight_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoGroupRoWeightValueArgs']]]]] = None) -> 'RoGroup':
        """
        Get an existing RoGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID, in the format: cdbro-3i70uj0k.
        :param pulumi.Input[int] is_balance_ro_load: Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        :param pulumi.Input[str] ro_group_id: The ID of the RO group.
        :param pulumi.Input[pulumi.InputType['RoGroupRoGroupInfoArgs']] ro_group_info: Details of the RO group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoGroupRoWeightValueArgs']]]] ro_weight_values: The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoGroupState.__new__(_RoGroupState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_balance_ro_load"] = is_balance_ro_load
        __props__.__dict__["ro_group_id"] = ro_group_id
        __props__.__dict__["ro_group_info"] = ro_group_info
        __props__.__dict__["ro_weight_values"] = ro_weight_values
        return RoGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID, in the format: cdbro-3i70uj0k.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isBalanceRoLoad")
    def is_balance_ro_load(self) -> pulumi.Output[Optional[int]]:
        """
        Whether to rebalance the load of RO instances in the RO group. Supported values include: 1 - rebalance load; 0 - do not rebalance load. The default value is 0. Note that when it is set to rebalance the load, the RO instance in the RO group will have a momentary disconnection of the database connection, please ensure that the application can reconnect to the database.
        """
        return pulumi.get(self, "is_balance_ro_load")

    @property
    @pulumi.getter(name="roGroupId")
    def ro_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the RO group.
        """
        return pulumi.get(self, "ro_group_id")

    @property
    @pulumi.getter(name="roGroupInfo")
    def ro_group_info(self) -> pulumi.Output[Optional['outputs.RoGroupRoGroupInfo']]:
        """
        Details of the RO group.
        """
        return pulumi.get(self, "ro_group_info")

    @property
    @pulumi.getter(name="roWeightValues")
    def ro_weight_values(self) -> pulumi.Output[Optional[Sequence['outputs.RoGroupRoWeightValue']]]:
        """
        The weight of the instance within the RO group. If the weight mode of the RO group is changed to user-defined mode (custom), this parameter must be set, and the weight value of each RO instance needs to be set.
        """
        return pulumi.get(self, "ro_weight_values")

