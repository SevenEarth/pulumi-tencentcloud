# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StorageAttachmentArgs', 'StorageAttachment']

@pulumi.input_type
class StorageAttachmentArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 storage_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a StorageAttachment resource.
        :param pulumi.Input[str] instance_id: ID of the CVM instance.
        :param pulumi.Input[str] storage_id: ID of the mounted CBS.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "storage_id", storage_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the CVM instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> pulumi.Input[str]:
        """
        ID of the mounted CBS.
        """
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_id", value)


@pulumi.input_type
class _StorageAttachmentState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StorageAttachment resources.
        :param pulumi.Input[str] instance_id: ID of the CVM instance.
        :param pulumi.Input[str] storage_id: ID of the mounted CBS.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if storage_id is not None:
            pulumi.set(__self__, "storage_id", storage_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CVM instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the mounted CBS.
        """
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_id", value)


class StorageAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a CBS storage attachment resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        attachment = tencentcloud.cbs.StorageAttachment("attachment",
            instance_id="ins-jqlegd42",
            storage_id="disk-kdt0sq6m")
        ```

        ## Import

        CBS storage attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cbs/storageAttachment:StorageAttachment attachment disk-41s6jwy4
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of the CVM instance.
        :param pulumi.Input[str] storage_id: ID of the mounted CBS.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CBS storage attachment resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        attachment = tencentcloud.cbs.StorageAttachment("attachment",
            instance_id="ins-jqlegd42",
            storage_id="disk-kdt0sq6m")
        ```

        ## Import

        CBS storage attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cbs/storageAttachment:StorageAttachment attachment disk-41s6jwy4
        ```

        :param str resource_name: The name of the resource.
        :param StorageAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = StorageAttachmentArgs.__new__(StorageAttachmentArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if storage_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_id'")
            __props__.__dict__["storage_id"] = storage_id
        super(StorageAttachment, __self__).__init__(
            'tencentcloud:Cbs/storageAttachment:StorageAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            storage_id: Optional[pulumi.Input[str]] = None) -> 'StorageAttachment':
        """
        Get an existing StorageAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of the CVM instance.
        :param pulumi.Input[str] storage_id: ID of the mounted CBS.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageAttachmentState.__new__(_StorageAttachmentState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["storage_id"] = storage_id
        return StorageAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the CVM instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> pulumi.Output[str]:
        """
        ID of the mounted CBS.
        """
        return pulumi.get(self, "storage_id")

