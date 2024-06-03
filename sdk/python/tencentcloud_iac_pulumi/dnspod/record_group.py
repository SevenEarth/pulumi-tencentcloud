# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RecordGroupArgs', 'RecordGroup']

@pulumi.input_type
class RecordGroupArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 group_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a RecordGroup resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] group_name: Record Group Name.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "group_name", group_name)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        Record Group Name.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)


@pulumi.input_type
class _RecordGroupState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RecordGroup resources.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[int] group_id: Group ID.
        :param pulumi.Input[str] group_name: Record Group Name.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        Group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Record Group Name.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)


class RecordGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dnspod record_group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        record_group = tencentcloud.dnspod.RecordGroup("recordGroup",
            domain="dnspod.cn",
            group_name="group_demo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        dnspod record_group can be imported using the domain#groupId, e.g.

        ```sh
        $ pulumi import tencentcloud:Dnspod/recordGroup:RecordGroup record_group domain#groupId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] group_name: Record Group Name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecordGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dnspod record_group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        record_group = tencentcloud.dnspod.RecordGroup("recordGroup",
            domain="dnspod.cn",
            group_name="group_demo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        dnspod record_group can be imported using the domain#groupId, e.g.

        ```sh
        $ pulumi import tencentcloud:Dnspod/recordGroup:RecordGroup record_group domain#groupId
        ```

        :param str resource_name: The name of the resource.
        :param RecordGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecordGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RecordGroupArgs.__new__(RecordGroupArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            __props__.__dict__["group_id"] = None
        super(RecordGroup, __self__).__init__(
            'tencentcloud:Dnspod/recordGroup:RecordGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[int]] = None,
            group_name: Optional[pulumi.Input[str]] = None) -> 'RecordGroup':
        """
        Get an existing RecordGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[int] group_id: Group ID.
        :param pulumi.Input[str] group_name: Record Group Name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecordGroupState.__new__(_RecordGroupState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["group_name"] = group_name
        return RecordGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[int]:
        """
        Group ID.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        Record Group Name.
        """
        return pulumi.get(self, "group_name")

