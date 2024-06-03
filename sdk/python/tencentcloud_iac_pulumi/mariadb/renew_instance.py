# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RenewInstanceArgs', 'RenewInstance']

@pulumi.input_type
class RenewInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 period: pulumi.Input[int]):
        """
        The set of arguments for constructing a RenewInstance resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: Renewal duration, unit: month.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "period", period)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def period(self) -> pulumi.Input[int]:
        """
        Renewal duration, unit: month.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: pulumi.Input[int]):
        pulumi.set(self, "period", value)


@pulumi.input_type
class _RenewInstanceState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RenewInstance resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: Renewal duration, unit: month.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if period is not None:
            pulumi.set(__self__, "period", period)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Renewal duration, unit: month.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)


class RenewInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a mariadb renew_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        renew_instance = tencentcloud.mariadb.RenewInstance("renewInstance",
            instance_id="tdsql-9vqvls95",
            period=1)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: Renewal duration, unit: month.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RenewInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mariadb renew_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        renew_instance = tencentcloud.mariadb.RenewInstance("renewInstance",
            instance_id="tdsql-9vqvls95",
            period=1)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param RenewInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RenewInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RenewInstanceArgs.__new__(RenewInstanceArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if period is None and not opts.urn:
                raise TypeError("Missing required property 'period'")
            __props__.__dict__["period"] = period
        super(RenewInstance, __self__).__init__(
            'tencentcloud:Mariadb/renewInstance:RenewInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None) -> 'RenewInstance':
        """
        Get an existing RenewInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: Renewal duration, unit: month.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RenewInstanceState.__new__(_RenewInstanceState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["period"] = period
        return RenewInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[int]:
        """
        Renewal duration, unit: month.
        """
        return pulumi.get(self, "period")

