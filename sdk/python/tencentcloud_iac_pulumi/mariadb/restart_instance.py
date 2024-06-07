# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RestartInstanceArgs', 'RestartInstance']

@pulumi.input_type
class RestartInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 restart_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RestartInstance resource.
        :param pulumi.Input[str] instance_id: instance ID.
        :param pulumi.Input[str] restart_time: expected restart time.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if restart_time is not None:
            pulumi.set(__self__, "restart_time", restart_time)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="restartTime")
    def restart_time(self) -> Optional[pulumi.Input[str]]:
        """
        expected restart time.
        """
        return pulumi.get(self, "restart_time")

    @restart_time.setter
    def restart_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restart_time", value)


@pulumi.input_type
class _RestartInstanceState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 restart_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RestartInstance resources.
        :param pulumi.Input[str] instance_id: instance ID.
        :param pulumi.Input[str] restart_time: expected restart time.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if restart_time is not None:
            pulumi.set(__self__, "restart_time", restart_time)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="restartTime")
    def restart_time(self) -> Optional[pulumi.Input[str]]:
        """
        expected restart time.
        """
        return pulumi.get(self, "restart_time")

    @restart_time.setter
    def restart_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restart_time", value)


class RestartInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 restart_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a mariadb restart_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        restart_instance = tencentcloud.mariadb.RestartInstance("restartInstance", instance_id="tdsql-9vqvls95")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: instance ID.
        :param pulumi.Input[str] restart_time: expected restart time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RestartInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mariadb restart_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        restart_instance = tencentcloud.mariadb.RestartInstance("restartInstance", instance_id="tdsql-9vqvls95")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param RestartInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RestartInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 restart_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RestartInstanceArgs.__new__(RestartInstanceArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["restart_time"] = restart_time
        super(RestartInstance, __self__).__init__(
            'tencentcloud:Mariadb/restartInstance:RestartInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            restart_time: Optional[pulumi.Input[str]] = None) -> 'RestartInstance':
        """
        Get an existing RestartInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: instance ID.
        :param pulumi.Input[str] restart_time: expected restart time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RestartInstanceState.__new__(_RestartInstanceState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["restart_time"] = restart_time
        return RestartInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="restartTime")
    def restart_time(self) -> pulumi.Output[Optional[str]]:
        """
        expected restart time.
        """
        return pulumi.get(self, "restart_time")

