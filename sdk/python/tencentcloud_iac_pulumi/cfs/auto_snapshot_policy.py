# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AutoSnapshotPolicyArgs', 'AutoSnapshotPolicy']

@pulumi.input_type
class AutoSnapshotPolicyArgs:
    def __init__(__self__, *,
                 hour: pulumi.Input[str],
                 alive_days: Optional[pulumi.Input[int]] = None,
                 day_of_month: Optional[pulumi.Input[str]] = None,
                 day_of_week: Optional[pulumi.Input[str]] = None,
                 interval_days: Optional[pulumi.Input[int]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AutoSnapshotPolicy resource.
        :param pulumi.Input[str] hour: The time point when to repeat the snapshot operation.
        :param pulumi.Input[int] alive_days: Snapshot retention period.
        :param pulumi.Input[str] day_of_month: The specific day (day 1 to day 31) of the month on which to create a snapshot.
        :param pulumi.Input[str] day_of_week: The day of the week on which to repeat the snapshot operation.
        :param pulumi.Input[int] interval_days: The snapshot interval, in days.
        :param pulumi.Input[str] policy_name: Policy name.
        """
        pulumi.set(__self__, "hour", hour)
        if alive_days is not None:
            pulumi.set(__self__, "alive_days", alive_days)
        if day_of_month is not None:
            pulumi.set(__self__, "day_of_month", day_of_month)
        if day_of_week is not None:
            pulumi.set(__self__, "day_of_week", day_of_week)
        if interval_days is not None:
            pulumi.set(__self__, "interval_days", interval_days)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter
    def hour(self) -> pulumi.Input[str]:
        """
        The time point when to repeat the snapshot operation.
        """
        return pulumi.get(self, "hour")

    @hour.setter
    def hour(self, value: pulumi.Input[str]):
        pulumi.set(self, "hour", value)

    @property
    @pulumi.getter(name="aliveDays")
    def alive_days(self) -> Optional[pulumi.Input[int]]:
        """
        Snapshot retention period.
        """
        return pulumi.get(self, "alive_days")

    @alive_days.setter
    def alive_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alive_days", value)

    @property
    @pulumi.getter(name="dayOfMonth")
    def day_of_month(self) -> Optional[pulumi.Input[str]]:
        """
        The specific day (day 1 to day 31) of the month on which to create a snapshot.
        """
        return pulumi.get(self, "day_of_month")

    @day_of_month.setter
    def day_of_month(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_of_month", value)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> Optional[pulumi.Input[str]]:
        """
        The day of the week on which to repeat the snapshot operation.
        """
        return pulumi.get(self, "day_of_week")

    @day_of_week.setter
    def day_of_week(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_of_week", value)

    @property
    @pulumi.getter(name="intervalDays")
    def interval_days(self) -> Optional[pulumi.Input[int]]:
        """
        The snapshot interval, in days.
        """
        return pulumi.get(self, "interval_days")

    @interval_days.setter
    def interval_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval_days", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)


@pulumi.input_type
class _AutoSnapshotPolicyState:
    def __init__(__self__, *,
                 alive_days: Optional[pulumi.Input[int]] = None,
                 day_of_month: Optional[pulumi.Input[str]] = None,
                 day_of_week: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 interval_days: Optional[pulumi.Input[int]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AutoSnapshotPolicy resources.
        :param pulumi.Input[int] alive_days: Snapshot retention period.
        :param pulumi.Input[str] day_of_month: The specific day (day 1 to day 31) of the month on which to create a snapshot.
        :param pulumi.Input[str] day_of_week: The day of the week on which to repeat the snapshot operation.
        :param pulumi.Input[str] hour: The time point when to repeat the snapshot operation.
        :param pulumi.Input[int] interval_days: The snapshot interval, in days.
        :param pulumi.Input[str] policy_name: Policy name.
        """
        if alive_days is not None:
            pulumi.set(__self__, "alive_days", alive_days)
        if day_of_month is not None:
            pulumi.set(__self__, "day_of_month", day_of_month)
        if day_of_week is not None:
            pulumi.set(__self__, "day_of_week", day_of_week)
        if hour is not None:
            pulumi.set(__self__, "hour", hour)
        if interval_days is not None:
            pulumi.set(__self__, "interval_days", interval_days)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter(name="aliveDays")
    def alive_days(self) -> Optional[pulumi.Input[int]]:
        """
        Snapshot retention period.
        """
        return pulumi.get(self, "alive_days")

    @alive_days.setter
    def alive_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alive_days", value)

    @property
    @pulumi.getter(name="dayOfMonth")
    def day_of_month(self) -> Optional[pulumi.Input[str]]:
        """
        The specific day (day 1 to day 31) of the month on which to create a snapshot.
        """
        return pulumi.get(self, "day_of_month")

    @day_of_month.setter
    def day_of_month(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_of_month", value)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> Optional[pulumi.Input[str]]:
        """
        The day of the week on which to repeat the snapshot operation.
        """
        return pulumi.get(self, "day_of_week")

    @day_of_week.setter
    def day_of_week(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_of_week", value)

    @property
    @pulumi.getter
    def hour(self) -> Optional[pulumi.Input[str]]:
        """
        The time point when to repeat the snapshot operation.
        """
        return pulumi.get(self, "hour")

    @hour.setter
    def hour(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hour", value)

    @property
    @pulumi.getter(name="intervalDays")
    def interval_days(self) -> Optional[pulumi.Input[int]]:
        """
        The snapshot interval, in days.
        """
        return pulumi.get(self, "interval_days")

    @interval_days.setter
    def interval_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval_days", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)


class AutoSnapshotPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alive_days: Optional[pulumi.Input[int]] = None,
                 day_of_month: Optional[pulumi.Input[str]] = None,
                 day_of_week: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 interval_days: Optional[pulumi.Input[int]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a cfs auto_snapshot_policy

        ## Example Usage

        ### use day of week

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        auto_snapshot_policy = tencentcloud.cfs.AutoSnapshotPolicy("autoSnapshotPolicy",
            alive_days=7,
            day_of_week="1,2",
            hour="2,3",
            policy_name="policy_name")
        ```
        <!--End PulumiCodeChooser -->

        ### use day of month

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        auto_snapshot_policy = tencentcloud.cfs.AutoSnapshotPolicy("autoSnapshotPolicy",
            alive_days=7,
            day_of_month="2,3,4",
            hour="2,3",
            policy_name="policy_name")
        ```
        <!--End PulumiCodeChooser -->

        ### use interval days

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        auto_snapshot_policy = tencentcloud.cfs.AutoSnapshotPolicy("autoSnapshotPolicy",
            alive_days=7,
            hour="2,3",
            interval_days=1,
            policy_name="policy_name")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cfs auto_snapshot_policy can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy auto_snapshot_policy auto_snapshot_policy_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] alive_days: Snapshot retention period.
        :param pulumi.Input[str] day_of_month: The specific day (day 1 to day 31) of the month on which to create a snapshot.
        :param pulumi.Input[str] day_of_week: The day of the week on which to repeat the snapshot operation.
        :param pulumi.Input[str] hour: The time point when to repeat the snapshot operation.
        :param pulumi.Input[int] interval_days: The snapshot interval, in days.
        :param pulumi.Input[str] policy_name: Policy name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoSnapshotPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cfs auto_snapshot_policy

        ## Example Usage

        ### use day of week

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        auto_snapshot_policy = tencentcloud.cfs.AutoSnapshotPolicy("autoSnapshotPolicy",
            alive_days=7,
            day_of_week="1,2",
            hour="2,3",
            policy_name="policy_name")
        ```
        <!--End PulumiCodeChooser -->

        ### use day of month

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        auto_snapshot_policy = tencentcloud.cfs.AutoSnapshotPolicy("autoSnapshotPolicy",
            alive_days=7,
            day_of_month="2,3,4",
            hour="2,3",
            policy_name="policy_name")
        ```
        <!--End PulumiCodeChooser -->

        ### use interval days

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        auto_snapshot_policy = tencentcloud.cfs.AutoSnapshotPolicy("autoSnapshotPolicy",
            alive_days=7,
            hour="2,3",
            interval_days=1,
            policy_name="policy_name")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cfs auto_snapshot_policy can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy auto_snapshot_policy auto_snapshot_policy_id
        ```

        :param str resource_name: The name of the resource.
        :param AutoSnapshotPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoSnapshotPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alive_days: Optional[pulumi.Input[int]] = None,
                 day_of_month: Optional[pulumi.Input[str]] = None,
                 day_of_week: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 interval_days: Optional[pulumi.Input[int]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoSnapshotPolicyArgs.__new__(AutoSnapshotPolicyArgs)

            __props__.__dict__["alive_days"] = alive_days
            __props__.__dict__["day_of_month"] = day_of_month
            __props__.__dict__["day_of_week"] = day_of_week
            if hour is None and not opts.urn:
                raise TypeError("Missing required property 'hour'")
            __props__.__dict__["hour"] = hour
            __props__.__dict__["interval_days"] = interval_days
            __props__.__dict__["policy_name"] = policy_name
        super(AutoSnapshotPolicy, __self__).__init__(
            'tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alive_days: Optional[pulumi.Input[int]] = None,
            day_of_month: Optional[pulumi.Input[str]] = None,
            day_of_week: Optional[pulumi.Input[str]] = None,
            hour: Optional[pulumi.Input[str]] = None,
            interval_days: Optional[pulumi.Input[int]] = None,
            policy_name: Optional[pulumi.Input[str]] = None) -> 'AutoSnapshotPolicy':
        """
        Get an existing AutoSnapshotPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] alive_days: Snapshot retention period.
        :param pulumi.Input[str] day_of_month: The specific day (day 1 to day 31) of the month on which to create a snapshot.
        :param pulumi.Input[str] day_of_week: The day of the week on which to repeat the snapshot operation.
        :param pulumi.Input[str] hour: The time point when to repeat the snapshot operation.
        :param pulumi.Input[int] interval_days: The snapshot interval, in days.
        :param pulumi.Input[str] policy_name: Policy name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutoSnapshotPolicyState.__new__(_AutoSnapshotPolicyState)

        __props__.__dict__["alive_days"] = alive_days
        __props__.__dict__["day_of_month"] = day_of_month
        __props__.__dict__["day_of_week"] = day_of_week
        __props__.__dict__["hour"] = hour
        __props__.__dict__["interval_days"] = interval_days
        __props__.__dict__["policy_name"] = policy_name
        return AutoSnapshotPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aliveDays")
    def alive_days(self) -> pulumi.Output[Optional[int]]:
        """
        Snapshot retention period.
        """
        return pulumi.get(self, "alive_days")

    @property
    @pulumi.getter(name="dayOfMonth")
    def day_of_month(self) -> pulumi.Output[Optional[str]]:
        """
        The specific day (day 1 to day 31) of the month on which to create a snapshot.
        """
        return pulumi.get(self, "day_of_month")

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> pulumi.Output[Optional[str]]:
        """
        The day of the week on which to repeat the snapshot operation.
        """
        return pulumi.get(self, "day_of_week")

    @property
    @pulumi.getter
    def hour(self) -> pulumi.Output[str]:
        """
        The time point when to repeat the snapshot operation.
        """
        return pulumi.get(self, "hour")

    @property
    @pulumi.getter(name="intervalDays")
    def interval_days(self) -> pulumi.Output[Optional[int]]:
        """
        The snapshot interval, in days.
        """
        return pulumi.get(self, "interval_days")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[Optional[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "policy_name")

