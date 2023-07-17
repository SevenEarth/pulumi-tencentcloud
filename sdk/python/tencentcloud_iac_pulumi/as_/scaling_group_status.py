# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ScalingGroupStatusArgs', 'ScalingGroupStatus']

@pulumi.input_type
class ScalingGroupStatusArgs:
    def __init__(__self__, *,
                 auto_scaling_group_id: pulumi.Input[str],
                 enable: pulumi.Input[bool]):
        """
        The set of arguments for constructing a ScalingGroupStatus resource.
        :param pulumi.Input[str] auto_scaling_group_id: Scaling group ID.
        :param pulumi.Input[bool] enable: If enable auto scaling group.
        """
        pulumi.set(__self__, "auto_scaling_group_id", auto_scaling_group_id)
        pulumi.set(__self__, "enable", enable)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> pulumi.Input[str]:
        """
        Scaling group ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @auto_scaling_group_id.setter
    def auto_scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_scaling_group_id", value)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[bool]:
        """
        If enable auto scaling group.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable", value)


@pulumi.input_type
class _ScalingGroupStatusState:
    def __init__(__self__, *,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ScalingGroupStatus resources.
        :param pulumi.Input[str] auto_scaling_group_id: Scaling group ID.
        :param pulumi.Input[bool] enable: If enable auto scaling group.
        """
        if auto_scaling_group_id is not None:
            pulumi.set(__self__, "auto_scaling_group_id", auto_scaling_group_id)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Scaling group ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @auto_scaling_group_id.setter
    def auto_scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_scaling_group_id", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        If enable auto scaling group.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)


class ScalingGroupStatus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a resource to set as scaling_group status

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        scaling_group_status = tencentcloud.as_.ScalingGroupStatus("scalingGroupStatus",
            auto_scaling_group_id="asg-519acdug",
            enable=False)
        ```

        ## Import

        as scaling_group_status can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:As/scalingGroupStatus:ScalingGroupStatus scaling_group_status scaling_group_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_id: Scaling group ID.
        :param pulumi.Input[bool] enable: If enable auto scaling group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScalingGroupStatusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to set as scaling_group status

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        scaling_group_status = tencentcloud.as_.ScalingGroupStatus("scalingGroupStatus",
            auto_scaling_group_id="asg-519acdug",
            enable=False)
        ```

        ## Import

        as scaling_group_status can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:As/scalingGroupStatus:ScalingGroupStatus scaling_group_status scaling_group_id
        ```

        :param str resource_name: The name of the resource.
        :param ScalingGroupStatusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScalingGroupStatusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
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
            __props__ = ScalingGroupStatusArgs.__new__(ScalingGroupStatusArgs)

            if auto_scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'auto_scaling_group_id'")
            __props__.__dict__["auto_scaling_group_id"] = auto_scaling_group_id
            if enable is None and not opts.urn:
                raise TypeError("Missing required property 'enable'")
            __props__.__dict__["enable"] = enable
        super(ScalingGroupStatus, __self__).__init__(
            'tencentcloud:As/scalingGroupStatus:ScalingGroupStatus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[bool]] = None) -> 'ScalingGroupStatus':
        """
        Get an existing ScalingGroupStatus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_id: Scaling group ID.
        :param pulumi.Input[bool] enable: If enable auto scaling group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScalingGroupStatusState.__new__(_ScalingGroupStatusState)

        __props__.__dict__["auto_scaling_group_id"] = auto_scaling_group_id
        __props__.__dict__["enable"] = enable
        return ScalingGroupStatus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> pulumi.Output[str]:
        """
        Scaling group ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[bool]:
        """
        If enable auto scaling group.
        """
        return pulumi.get(self, "enable")

