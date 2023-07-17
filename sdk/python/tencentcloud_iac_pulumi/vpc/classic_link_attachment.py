# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClassicLinkAttachmentArgs', 'ClassicLinkAttachment']

@pulumi.input_type
class ClassicLinkAttachmentArgs:
    def __init__(__self__, *,
                 instance_ids: pulumi.Input[str],
                 vpc_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ClassicLinkAttachment resource.
        :param pulumi.Input[str] instance_ids: CVM instance ID. It only support set one instance now.
        :param pulumi.Input[str] vpc_id: VPC instance ID.
        """
        pulumi.set(__self__, "instance_ids", instance_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Input[str]:
        """
        CVM instance ID. It only support set one instance now.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        VPC instance ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _ClassicLinkAttachmentState:
    def __init__(__self__, *,
                 instance_ids: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClassicLinkAttachment resources.
        :param pulumi.Input[str] instance_ids: CVM instance ID. It only support set one instance now.
        :param pulumi.Input[str] vpc_id: VPC instance ID.
        """
        if instance_ids is not None:
            pulumi.set(__self__, "instance_ids", instance_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[pulumi.Input[str]]:
        """
        CVM instance ID. It only support set one instance now.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC instance ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class ClassicLinkAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_ids: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ClassicLinkAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_ids: CVM instance ID. It only support set one instance now.
        :param pulumi.Input[str] vpc_id: VPC instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClassicLinkAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ClassicLinkAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ClassicLinkAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClassicLinkAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_ids: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ClassicLinkAttachmentArgs.__new__(ClassicLinkAttachmentArgs)

            if instance_ids is None and not opts.urn:
                raise TypeError("Missing required property 'instance_ids'")
            __props__.__dict__["instance_ids"] = instance_ids
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(ClassicLinkAttachment, __self__).__init__(
            'tencentcloud:Vpc/classicLinkAttachment:ClassicLinkAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_ids: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'ClassicLinkAttachment':
        """
        Get an existing ClassicLinkAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_ids: CVM instance ID. It only support set one instance now.
        :param pulumi.Input[str] vpc_id: VPC instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClassicLinkAttachmentState.__new__(_ClassicLinkAttachmentState)

        __props__.__dict__["instance_ids"] = instance_ids
        __props__.__dict__["vpc_id"] = vpc_id
        return ClassicLinkAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Output[str]:
        """
        CVM instance ID. It only support set one instance now.
        """
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        VPC instance ID.
        """
        return pulumi.get(self, "vpc_id")

