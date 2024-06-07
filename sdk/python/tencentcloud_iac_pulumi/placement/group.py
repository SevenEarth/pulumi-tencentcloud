# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[str] type: Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        :param pulumi.Input[str] name: Name of the placement group, 1-60 characters in length.
        """
        pulumi.set(__self__, "type", type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the placement group, 1-60 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 current_num: Optional[pulumi.Input[int]] = None,
                 cvm_quota_total: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[str] create_time: Creation time of the placement group.
        :param pulumi.Input[int] current_num: Number of hosts in the placement group.
        :param pulumi.Input[int] cvm_quota_total: Maximum number of hosts in the placement group.
        :param pulumi.Input[str] name: Name of the placement group, 1-60 characters in length.
        :param pulumi.Input[str] type: Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if current_num is not None:
            pulumi.set(__self__, "current_num", current_num)
        if cvm_quota_total is not None:
            pulumi.set(__self__, "cvm_quota_total", cvm_quota_total)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of the placement group.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="currentNum")
    def current_num(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hosts in the placement group.
        """
        return pulumi.get(self, "current_num")

    @current_num.setter
    def current_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "current_num", value)

    @property
    @pulumi.getter(name="cvmQuotaTotal")
    def cvm_quota_total(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of hosts in the placement group.
        """
        return pulumi.get(self, "cvm_quota_total")

    @cvm_quota_total.setter
    def cvm_quota_total(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cvm_quota_total", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the placement group, 1-60 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide a resource to create a placement group.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.placement.Group("foo", type="HOST")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Placement group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Placement/group:Group foo ps-ilan8vjf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the placement group, 1-60 characters in length.
        :param pulumi.Input[str] type: Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a placement group.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.placement.Group("foo", type="HOST")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Placement group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Placement/group:Group foo ps-ilan8vjf
        ```

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["name"] = name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["current_num"] = None
            __props__.__dict__["cvm_quota_total"] = None
        super(Group, __self__).__init__(
            'tencentcloud:Placement/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            current_num: Optional[pulumi.Input[int]] = None,
            cvm_quota_total: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Creation time of the placement group.
        :param pulumi.Input[int] current_num: Number of hosts in the placement group.
        :param pulumi.Input[int] cvm_quota_total: Maximum number of hosts in the placement group.
        :param pulumi.Input[str] name: Name of the placement group, 1-60 characters in length.
        :param pulumi.Input[str] type: Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["current_num"] = current_num
        __props__.__dict__["cvm_quota_total"] = cvm_quota_total
        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of the placement group.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="currentNum")
    def current_num(self) -> pulumi.Output[int]:
        """
        Number of hosts in the placement group.
        """
        return pulumi.get(self, "current_num")

    @property
    @pulumi.getter(name="cvmQuotaTotal")
    def cvm_quota_total(self) -> pulumi.Output[int]:
        """
        Maximum number of hosts in the placement group.
        """
        return pulumi.get(self, "cvm_quota_total")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the placement group, 1-60 characters in length.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
        """
        return pulumi.get(self, "type")

