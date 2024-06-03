# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeviceGroupArgs', 'DeviceGroup']

@pulumi.input_type
class DeviceGroupArgs:
    def __init__(__self__, *,
                 department_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DeviceGroup resource.
        :param pulumi.Input[str] department_id: The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        :param pulumi.Input[str] name: Device group name, the maximum length is 32 characters.
        """
        if department_id is not None:
            pulumi.set(__self__, "department_id", department_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="departmentId")
    def department_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        """
        return pulumi.get(self, "department_id")

    @department_id.setter
    def department_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "department_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Device group name, the maximum length is 32 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DeviceGroupState:
    def __init__(__self__, *,
                 department_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DeviceGroup resources.
        :param pulumi.Input[str] department_id: The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        :param pulumi.Input[str] name: Device group name, the maximum length is 32 characters.
        """
        if department_id is not None:
            pulumi.set(__self__, "department_id", department_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="departmentId")
    def department_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        """
        return pulumi.get(self, "department_id")

    @department_id.setter
    def department_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "department_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Device group name, the maximum length is 32 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class DeviceGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 department_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dasb device_group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.dasb.DeviceGroup("example", department_id="1.2")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        dasb device_group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Dasb/deviceGroup:DeviceGroup example 36
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] department_id: The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        :param pulumi.Input[str] name: Device group name, the maximum length is 32 characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DeviceGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dasb device_group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.dasb.DeviceGroup("example", department_id="1.2")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        dasb device_group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Dasb/deviceGroup:DeviceGroup example 36
        ```

        :param str resource_name: The name of the resource.
        :param DeviceGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 department_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceGroupArgs.__new__(DeviceGroupArgs)

            __props__.__dict__["department_id"] = department_id
            __props__.__dict__["name"] = name
        super(DeviceGroup, __self__).__init__(
            'tencentcloud:Dasb/deviceGroup:DeviceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            department_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'DeviceGroup':
        """
        Get an existing DeviceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] department_id: The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        :param pulumi.Input[str] name: Device group name, the maximum length is 32 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceGroupState.__new__(_DeviceGroupState)

        __props__.__dict__["department_id"] = department_id
        __props__.__dict__["name"] = name
        return DeviceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="departmentId")
    def department_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
        """
        return pulumi.get(self, "department_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Device group name, the maximum length is 32 characters.
        """
        return pulumi.get(self, "name")

