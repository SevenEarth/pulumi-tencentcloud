# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReadonlyAttachmentArgs', 'ReadonlyAttachment']

@pulumi.input_type
class ReadonlyAttachmentArgs:
    def __init__(__self__, *,
                 db_instance_id: pulumi.Input[str],
                 read_only_group_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ReadonlyAttachment resource.
        :param pulumi.Input[str] db_instance_id: Read only instance ID.
        :param pulumi.Input[str] read_only_group_id: Read only group ID.
        """
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "read_only_group_id", read_only_group_id)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Input[str]:
        """
        Read only instance ID.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="readOnlyGroupId")
    def read_only_group_id(self) -> pulumi.Input[str]:
        """
        Read only group ID.
        """
        return pulumi.get(self, "read_only_group_id")

    @read_only_group_id.setter
    def read_only_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "read_only_group_id", value)


@pulumi.input_type
class _ReadonlyAttachmentState:
    def __init__(__self__, *,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 read_only_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReadonlyAttachment resources.
        :param pulumi.Input[str] db_instance_id: Read only instance ID.
        :param pulumi.Input[str] read_only_group_id: Read only group ID.
        """
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if read_only_group_id is not None:
            pulumi.set(__self__, "read_only_group_id", read_only_group_id)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Read only instance ID.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="readOnlyGroupId")
    def read_only_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Read only group ID.
        """
        return pulumi.get(self, "read_only_group_id")

    @read_only_group_id.setter
    def read_only_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "read_only_group_id", value)


class ReadonlyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 read_only_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create postgresql readonly attachment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        attach = tencentcloud.postgresql.ReadonlyAttachment("attach",
            db_instance_id=tencentcloud_postgresql_readonly_instance["foo"]["id"],
            read_only_group_id=tencentcloud_postgresql_readonly_group["group"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: Read only instance ID.
        :param pulumi.Input[str] read_only_group_id: Read only group ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReadonlyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create postgresql readonly attachment.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        attach = tencentcloud.postgresql.ReadonlyAttachment("attach",
            db_instance_id=tencentcloud_postgresql_readonly_instance["foo"]["id"],
            read_only_group_id=tencentcloud_postgresql_readonly_group["group"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param ReadonlyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReadonlyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 read_only_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReadonlyAttachmentArgs.__new__(ReadonlyAttachmentArgs)

            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            if read_only_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'read_only_group_id'")
            __props__.__dict__["read_only_group_id"] = read_only_group_id
        super(ReadonlyAttachment, __self__).__init__(
            'tencentcloud:Postgresql/readonlyAttachment:ReadonlyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            read_only_group_id: Optional[pulumi.Input[str]] = None) -> 'ReadonlyAttachment':
        """
        Get an existing ReadonlyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: Read only instance ID.
        :param pulumi.Input[str] read_only_group_id: Read only group ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReadonlyAttachmentState.__new__(_ReadonlyAttachmentState)

        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["read_only_group_id"] = read_only_group_id
        return ReadonlyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        Read only instance ID.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="readOnlyGroupId")
    def read_only_group_id(self) -> pulumi.Output[str]:
        """
        Read only group ID.
        """
        return pulumi.get(self, "read_only_group_id")

