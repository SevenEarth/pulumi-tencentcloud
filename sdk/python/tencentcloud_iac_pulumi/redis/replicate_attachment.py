# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReplicateAttachmentArgs', 'ReplicateAttachment']

@pulumi.input_type
class ReplicateAttachmentArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 instance_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 master_instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ReplicateAttachment resource.
        :param pulumi.Input[str] group_id: The ID of group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: All instance ids of the replication group.
        :param pulumi.Input[str] master_instance_id: The ID of master instance.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "instance_ids", instance_ids)
        pulumi.set(__self__, "master_instance_id", master_instance_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The ID of group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        All instance ids of the replication group.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter(name="masterInstanceId")
    def master_instance_id(self) -> pulumi.Input[str]:
        """
        The ID of master instance.
        """
        return pulumi.get(self, "master_instance_id")

    @master_instance_id.setter
    def master_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "master_instance_id", value)


@pulumi.input_type
class _ReplicateAttachmentState:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 master_instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReplicateAttachment resources.
        :param pulumi.Input[str] group_id: The ID of group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: All instance ids of the replication group.
        :param pulumi.Input[str] master_instance_id: The ID of master instance.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if instance_ids is not None:
            pulumi.set(__self__, "instance_ids", instance_ids)
        if master_instance_id is not None:
            pulumi.set(__self__, "master_instance_id", master_instance_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        All instance ids of the replication group.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter(name="masterInstanceId")
    def master_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of master instance.
        """
        return pulumi.get(self, "master_instance_id")

    @master_instance_id.setter
    def master_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_instance_id", value)


class ReplicateAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 master_instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a redis replicate_attachment

        ## Import

        redis replicate_attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Redis/replicateAttachment:ReplicateAttachment replicate_attachment replicate_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: The ID of group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: All instance ids of the replication group.
        :param pulumi.Input[str] master_instance_id: The ID of master instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicateAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a redis replicate_attachment

        ## Import

        redis replicate_attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Redis/replicateAttachment:ReplicateAttachment replicate_attachment replicate_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param ReplicateAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicateAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 master_instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicateAttachmentArgs.__new__(ReplicateAttachmentArgs)

            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if instance_ids is None and not opts.urn:
                raise TypeError("Missing required property 'instance_ids'")
            __props__.__dict__["instance_ids"] = instance_ids
            if master_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'master_instance_id'")
            __props__.__dict__["master_instance_id"] = master_instance_id
        super(ReplicateAttachment, __self__).__init__(
            'tencentcloud:Redis/replicateAttachment:ReplicateAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            master_instance_id: Optional[pulumi.Input[str]] = None) -> 'ReplicateAttachment':
        """
        Get an existing ReplicateAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: The ID of group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: All instance ids of the replication group.
        :param pulumi.Input[str] master_instance_id: The ID of master instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicateAttachmentState.__new__(_ReplicateAttachmentState)

        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["instance_ids"] = instance_ids
        __props__.__dict__["master_instance_id"] = master_instance_id
        return ReplicateAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The ID of group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        All instance ids of the replication group.
        """
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="masterInstanceId")
    def master_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of master instance.
        """
        return pulumi.get(self, "master_instance_id")

