# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AddUsersToWorkGroupAttachmentArgs', 'AddUsersToWorkGroupAttachment']

@pulumi.input_type
class AddUsersToWorkGroupAttachmentArgs:
    def __init__(__self__, *,
                 add_info: pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs']):
        """
        The set of arguments for constructing a AddUsersToWorkGroupAttachment resource.
        :param pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs'] add_info: Work group and user information to operate on.
        """
        pulumi.set(__self__, "add_info", add_info)

    @property
    @pulumi.getter(name="addInfo")
    def add_info(self) -> pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs']:
        """
        Work group and user information to operate on.
        """
        return pulumi.get(self, "add_info")

    @add_info.setter
    def add_info(self, value: pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs']):
        pulumi.set(self, "add_info", value)


@pulumi.input_type
class _AddUsersToWorkGroupAttachmentState:
    def __init__(__self__, *,
                 add_info: Optional[pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs']] = None):
        """
        Input properties used for looking up and filtering AddUsersToWorkGroupAttachment resources.
        :param pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs'] add_info: Work group and user information to operate on.
        """
        if add_info is not None:
            pulumi.set(__self__, "add_info", add_info)

    @property
    @pulumi.getter(name="addInfo")
    def add_info(self) -> Optional[pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs']]:
        """
        Work group and user information to operate on.
        """
        return pulumi.get(self, "add_info")

    @add_info.setter
    def add_info(self, value: Optional[pulumi.Input['AddUsersToWorkGroupAttachmentAddInfoArgs']]):
        pulumi.set(self, "add_info", value)


class AddUsersToWorkGroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_info: Optional[pulumi.Input[pulumi.InputType['AddUsersToWorkGroupAttachmentAddInfoArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a dlc add_users_to_work_group_attachment

        ## Import

        dlc add_users_to_work_group_attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment add_users_to_work_group_attachment add_users_to_work_group_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AddUsersToWorkGroupAttachmentAddInfoArgs']] add_info: Work group and user information to operate on.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AddUsersToWorkGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dlc add_users_to_work_group_attachment

        ## Import

        dlc add_users_to_work_group_attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment add_users_to_work_group_attachment add_users_to_work_group_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param AddUsersToWorkGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AddUsersToWorkGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_info: Optional[pulumi.Input[pulumi.InputType['AddUsersToWorkGroupAttachmentAddInfoArgs']]] = None,
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
            __props__ = AddUsersToWorkGroupAttachmentArgs.__new__(AddUsersToWorkGroupAttachmentArgs)

            if add_info is None and not opts.urn:
                raise TypeError("Missing required property 'add_info'")
            __props__.__dict__["add_info"] = add_info
        super(AddUsersToWorkGroupAttachment, __self__).__init__(
            'tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_info: Optional[pulumi.Input[pulumi.InputType['AddUsersToWorkGroupAttachmentAddInfoArgs']]] = None) -> 'AddUsersToWorkGroupAttachment':
        """
        Get an existing AddUsersToWorkGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AddUsersToWorkGroupAttachmentAddInfoArgs']] add_info: Work group and user information to operate on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AddUsersToWorkGroupAttachmentState.__new__(_AddUsersToWorkGroupAttachmentState)

        __props__.__dict__["add_info"] = add_info
        return AddUsersToWorkGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addInfo")
    def add_info(self) -> pulumi.Output['outputs.AddUsersToWorkGroupAttachmentAddInfo']:
        """
        Work group and user information to operate on.
        """
        return pulumi.get(self, "add_info")

