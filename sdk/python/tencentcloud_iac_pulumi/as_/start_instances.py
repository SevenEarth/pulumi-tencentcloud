# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StartInstancesArgs', 'StartInstances']

@pulumi.input_type
class StartInstancesArgs:
    def __init__(__self__, *,
                 auto_scaling_group_id: pulumi.Input[str],
                 instance_ids: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a StartInstances resource.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        pulumi.set(__self__, "auto_scaling_group_id", auto_scaling_group_id)
        pulumi.set(__self__, "instance_ids", instance_ids)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> pulumi.Input[str]:
        """
        Launch configuration ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @auto_scaling_group_id.setter
    def auto_scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_scaling_group_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of cvm instances to start.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "instance_ids", value)


@pulumi.input_type
class _StartInstancesState:
    def __init__(__self__, *,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering StartInstances resources.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        if auto_scaling_group_id is not None:
            pulumi.set(__self__, "auto_scaling_group_id", auto_scaling_group_id)
        if instance_ids is not None:
            pulumi.set(__self__, "instance_ids", instance_ids)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Launch configuration ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @auto_scaling_group_id.setter
    def auto_scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_scaling_group_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of cvm instances to start.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_ids", value)


class StartInstances(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a as start_instances

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        start_instances = tencentcloud.as_.StartInstances("startInstances",
            auto_scaling_group_id=tencentcloud_as_scaling_group["scaling_group"]["id"],
            instance_ids=["ins-xxxxx"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StartInstancesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a as start_instances

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        start_instances = tencentcloud.as_.StartInstances("startInstances",
            auto_scaling_group_id=tencentcloud_as_scaling_group["scaling_group"]["id"],
            instance_ids=["ins-xxxxx"])
        ```

        :param str resource_name: The name of the resource.
        :param StartInstancesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StartInstancesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = StartInstancesArgs.__new__(StartInstancesArgs)

            if auto_scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'auto_scaling_group_id'")
            __props__.__dict__["auto_scaling_group_id"] = auto_scaling_group_id
            if instance_ids is None and not opts.urn:
                raise TypeError("Missing required property 'instance_ids'")
            __props__.__dict__["instance_ids"] = instance_ids
        super(StartInstances, __self__).__init__(
            'tencentcloud:As/startInstances:StartInstances',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
            instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'StartInstances':
        """
        Get an existing StartInstances resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StartInstancesState.__new__(_StartInstancesState)

        __props__.__dict__["auto_scaling_group_id"] = auto_scaling_group_id
        __props__.__dict__["instance_ids"] = instance_ids
        return StartInstances(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> pulumi.Output[str]:
        """
        Launch configuration ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of cvm instances to start.
        """
        return pulumi.get(self, "instance_ids")

