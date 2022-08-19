# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ScheduleArgs', 'Schedule']

@pulumi.input_type
class ScheduleArgs:
    def __init__(__self__, *,
                 desired_capacity: pulumi.Input[int],
                 max_size: pulumi.Input[int],
                 min_size: pulumi.Input[int],
                 scaling_group_id: pulumi.Input[str],
                 schedule_action_name: pulumi.Input[str],
                 start_time: pulumi.Input[str],
                 end_time: Optional[pulumi.Input[str]] = None,
                 recurrence: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Schedule resource.
        :param pulumi.Input[int] desired_capacity: The desired number of CVM instances that should be running in the group.
        :param pulumi.Input[int] max_size: The maximum size for the Auto Scaling group.
        :param pulumi.Input[int] min_size: The minimum size for the Auto Scaling group.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        :param pulumi.Input[str] schedule_action_name: The name of this scaling action.
        :param pulumi.Input[str] start_time: The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        :param pulumi.Input[str] end_time: The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        :param pulumi.Input[str] recurrence: The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        """
        pulumi.set(__self__, "desired_capacity", desired_capacity)
        pulumi.set(__self__, "max_size", max_size)
        pulumi.set(__self__, "min_size", min_size)
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        pulumi.set(__self__, "schedule_action_name", schedule_action_name)
        pulumi.set(__self__, "start_time", start_time)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if recurrence is not None:
            pulumi.set(__self__, "recurrence", recurrence)

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> pulumi.Input[int]:
        """
        The desired number of CVM instances that should be running in the group.
        """
        return pulumi.get(self, "desired_capacity")

    @desired_capacity.setter
    def desired_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "desired_capacity", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Input[int]:
        """
        The maximum size for the Auto Scaling group.
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Input[int]:
        """
        The minimum size for the Auto Scaling group.
        """
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_size", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Input[str]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_group_id", value)

    @property
    @pulumi.getter(name="scheduleActionName")
    def schedule_action_name(self) -> pulumi.Input[str]:
        """
        The name of this scaling action.
        """
        return pulumi.get(self, "schedule_action_name")

    @schedule_action_name.setter
    def schedule_action_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule_action_name", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input[str]:
        """
        The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def recurrence(self) -> Optional[pulumi.Input[str]]:
        """
        The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        """
        return pulumi.get(self, "recurrence")

    @recurrence.setter
    def recurrence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recurrence", value)


@pulumi.input_type
class _ScheduleState:
    def __init__(__self__, *,
                 desired_capacity: Optional[pulumi.Input[int]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 min_size: Optional[pulumi.Input[int]] = None,
                 recurrence: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 schedule_action_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Schedule resources.
        :param pulumi.Input[int] desired_capacity: The desired number of CVM instances that should be running in the group.
        :param pulumi.Input[str] end_time: The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        :param pulumi.Input[int] max_size: The maximum size for the Auto Scaling group.
        :param pulumi.Input[int] min_size: The minimum size for the Auto Scaling group.
        :param pulumi.Input[str] recurrence: The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        :param pulumi.Input[str] schedule_action_name: The name of this scaling action.
        :param pulumi.Input[str] start_time: The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        if desired_capacity is not None:
            pulumi.set(__self__, "desired_capacity", desired_capacity)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if min_size is not None:
            pulumi.set(__self__, "min_size", min_size)
        if recurrence is not None:
            pulumi.set(__self__, "recurrence", recurrence)
        if scaling_group_id is not None:
            pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        if schedule_action_name is not None:
            pulumi.set(__self__, "schedule_action_name", schedule_action_name)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The desired number of CVM instances that should be running in the group.
        """
        return pulumi.get(self, "desired_capacity")

    @desired_capacity.setter
    def desired_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "desired_capacity", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size for the Auto Scaling group.
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum size for the Auto Scaling group.
        """
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_size", value)

    @property
    @pulumi.getter
    def recurrence(self) -> Optional[pulumi.Input[str]]:
        """
        The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        """
        return pulumi.get(self, "recurrence")

    @recurrence.setter
    def recurrence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recurrence", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_group_id", value)

    @property
    @pulumi.getter(name="scheduleActionName")
    def schedule_action_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this scaling action.
        """
        return pulumi.get(self, "schedule_action_name")

    @schedule_action_name.setter
    def schedule_action_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_action_name", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)


class Schedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desired_capacity: Optional[pulumi.Input[int]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 min_size: Optional[pulumi.Input[int]] = None,
                 recurrence: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 schedule_action_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource for an AS (Auto scaling) schedule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        schedule = tencentcloud.as_.Schedule("schedule",
            desired_capacity=0,
            end_time="2019-12-01T00:00:00+08:00",
            max_size=10,
            min_size=0,
            recurrence="0 0 * * *",
            scaling_group_id="sg-12af45",
            schedule_action_name="tf-as-schedule",
            start_time="2019-01-01T00:00:00+08:00")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] desired_capacity: The desired number of CVM instances that should be running in the group.
        :param pulumi.Input[str] end_time: The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        :param pulumi.Input[int] max_size: The maximum size for the Auto Scaling group.
        :param pulumi.Input[int] min_size: The minimum size for the Auto Scaling group.
        :param pulumi.Input[str] recurrence: The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        :param pulumi.Input[str] schedule_action_name: The name of this scaling action.
        :param pulumi.Input[str] start_time: The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for an AS (Auto scaling) schedule.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        schedule = tencentcloud.as_.Schedule("schedule",
            desired_capacity=0,
            end_time="2019-12-01T00:00:00+08:00",
            max_size=10,
            min_size=0,
            recurrence="0 0 * * *",
            scaling_group_id="sg-12af45",
            schedule_action_name="tf-as-schedule",
            start_time="2019-01-01T00:00:00+08:00")
        ```

        :param str resource_name: The name of the resource.
        :param ScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desired_capacity: Optional[pulumi.Input[int]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 min_size: Optional[pulumi.Input[int]] = None,
                 recurrence: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 schedule_action_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
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
            __props__ = ScheduleArgs.__new__(ScheduleArgs)

            if desired_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'desired_capacity'")
            __props__.__dict__["desired_capacity"] = desired_capacity
            __props__.__dict__["end_time"] = end_time
            if max_size is None and not opts.urn:
                raise TypeError("Missing required property 'max_size'")
            __props__.__dict__["max_size"] = max_size
            if min_size is None and not opts.urn:
                raise TypeError("Missing required property 'min_size'")
            __props__.__dict__["min_size"] = min_size
            __props__.__dict__["recurrence"] = recurrence
            if scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__.__dict__["scaling_group_id"] = scaling_group_id
            if schedule_action_name is None and not opts.urn:
                raise TypeError("Missing required property 'schedule_action_name'")
            __props__.__dict__["schedule_action_name"] = schedule_action_name
            if start_time is None and not opts.urn:
                raise TypeError("Missing required property 'start_time'")
            __props__.__dict__["start_time"] = start_time
        super(Schedule, __self__).__init__(
            'tencentcloud:As/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            desired_capacity: Optional[pulumi.Input[int]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            max_size: Optional[pulumi.Input[int]] = None,
            min_size: Optional[pulumi.Input[int]] = None,
            recurrence: Optional[pulumi.Input[str]] = None,
            scaling_group_id: Optional[pulumi.Input[str]] = None,
            schedule_action_name: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[str]] = None) -> 'Schedule':
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] desired_capacity: The desired number of CVM instances that should be running in the group.
        :param pulumi.Input[str] end_time: The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        :param pulumi.Input[int] max_size: The maximum size for the Auto Scaling group.
        :param pulumi.Input[int] min_size: The minimum size for the Auto Scaling group.
        :param pulumi.Input[str] recurrence: The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        :param pulumi.Input[str] schedule_action_name: The name of this scaling action.
        :param pulumi.Input[str] start_time: The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScheduleState.__new__(_ScheduleState)

        __props__.__dict__["desired_capacity"] = desired_capacity
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["max_size"] = max_size
        __props__.__dict__["min_size"] = min_size
        __props__.__dict__["recurrence"] = recurrence
        __props__.__dict__["scaling_group_id"] = scaling_group_id
        __props__.__dict__["schedule_action_name"] = schedule_action_name
        __props__.__dict__["start_time"] = start_time
        return Schedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> pulumi.Output[int]:
        """
        The desired number of CVM instances that should be running in the group.
        """
        return pulumi.get(self, "desired_capacity")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[str]]:
        """
        The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Output[int]:
        """
        The maximum size for the Auto Scaling group.
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Output[int]:
        """
        The minimum size for the Auto Scaling group.
        """
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter
    def recurrence(self) -> pulumi.Output[Optional[str]]:
        """
        The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.
        """
        return pulumi.get(self, "recurrence")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Output[str]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @property
    @pulumi.getter(name="scheduleActionName")
    def schedule_action_name(self) -> pulumi.Output[str]:
        """
        The name of this scaling action.
        """
        return pulumi.get(self, "schedule_action_name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).
        """
        return pulumi.get(self, "start_time")

