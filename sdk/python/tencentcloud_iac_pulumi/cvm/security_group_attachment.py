# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecurityGroupAttachmentArgs', 'SecurityGroupAttachment']

@pulumi.input_type
class SecurityGroupAttachmentArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 security_group_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a SecurityGroupAttachment resource.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] security_group_id: Security group id.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        Security group id.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)


@pulumi.input_type
class _SecurityGroupAttachmentState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityGroupAttachment resources.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] security_group_id: Security group id.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Security group id.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)


class SecurityGroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a cvm security_group_attachment

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        security_group_attachment = tencentcloud.cvm.SecurityGroupAttachment("securityGroupAttachment",
            instance_id="ins-xxxxxxxx",
            security_group_id="sg-xxxxxxx")
        ```

        ## Import

        cvm security_group_attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cvm/securityGroupAttachment:SecurityGroupAttachment security_group_attachment ${instance_id}#${security_group_id}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] security_group_id: Security group id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cvm security_group_attachment

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        security_group_attachment = tencentcloud.cvm.SecurityGroupAttachment("securityGroupAttachment",
            instance_id="ins-xxxxxxxx",
            security_group_id="sg-xxxxxxx")
        ```

        ## Import

        cvm security_group_attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cvm/securityGroupAttachment:SecurityGroupAttachment security_group_attachment ${instance_id}#${security_group_id}
        ```

        :param str resource_name: The name of the resource.
        :param SecurityGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SecurityGroupAttachmentArgs.__new__(SecurityGroupAttachmentArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
        super(SecurityGroupAttachment, __self__).__init__(
            'tencentcloud:Cvm/securityGroupAttachment:SecurityGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None) -> 'SecurityGroupAttachment':
        """
        Get an existing SecurityGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] security_group_id: Security group id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityGroupAttachmentState.__new__(_SecurityGroupAttachmentState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["security_group_id"] = security_group_id
        return SecurityGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        Security group id.
        """
        return pulumi.get(self, "security_group_id")

