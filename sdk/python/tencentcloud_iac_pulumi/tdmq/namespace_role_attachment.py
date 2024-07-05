# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NamespaceRoleAttachmentArgs', 'NamespaceRoleAttachment']

@pulumi.input_type
class NamespaceRoleAttachmentArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 environ_id: pulumi.Input[str],
                 permissions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a NamespaceRoleAttachment resource.
        :param pulumi.Input[str] cluster_id: The id of tdmq cluster.
        :param pulumi.Input[str] environ_id: The name of tdmq namespace.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The permissions of tdmq role.
        :param pulumi.Input[str] role_name: The name of tdmq role.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "environ_id", environ_id)
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "role_name", role_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The id of tdmq cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="environId")
    def environ_id(self) -> pulumi.Input[str]:
        """
        The name of tdmq namespace.
        """
        return pulumi.get(self, "environ_id")

    @environ_id.setter
    def environ_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environ_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The permissions of tdmq role.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        """
        The name of tdmq role.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)


@pulumi.input_type
class _NamespaceRoleAttachmentState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 environ_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NamespaceRoleAttachment resources.
        :param pulumi.Input[str] cluster_id: The id of tdmq cluster.
        :param pulumi.Input[str] create_time: Creation time of resource.
        :param pulumi.Input[str] environ_id: The name of tdmq namespace.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The permissions of tdmq role.
        :param pulumi.Input[str] role_name: The name of tdmq role.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if environ_id is not None:
            pulumi.set(__self__, "environ_id", environ_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of tdmq cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="environId")
    def environ_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of tdmq namespace.
        """
        return pulumi.get(self, "environ_id")

    @environ_id.setter
    def environ_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environ_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The permissions of tdmq role.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of tdmq role.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)


class NamespaceRoleAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 environ_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide a resource to create a TDMQ role.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_instance = tencentcloud.tdmq.Instance("exampleInstance",
            cluster_name="tf_example",
            remark="remark.",
            tags={
                "createdBy": "terraform",
            })
        example_namespace = tencentcloud.tdmq.Namespace("exampleNamespace",
            environ_name="tf_example",
            msg_ttl=300,
            cluster_id=example_instance.id,
            retention_policy=tencentcloud.tdmq.NamespaceRetentionPolicyArgs(
                time_in_minutes=60,
                size_in_mb=10,
            ),
            remark="remark.")
        example_role = tencentcloud.tdmq.Role("exampleRole",
            role_name="tf_example",
            cluster_id=example_instance.id,
            remark="remark.")
        example_namespace_role_attachment = tencentcloud.tdmq.NamespaceRoleAttachment("exampleNamespaceRoleAttachment",
            environ_id=example_namespace.environ_name,
            role_name=example_role.role_name,
            permissions=[
                "produce",
                "consume",
            ],
            cluster_id=example_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The id of tdmq cluster.
        :param pulumi.Input[str] environ_id: The name of tdmq namespace.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The permissions of tdmq role.
        :param pulumi.Input[str] role_name: The name of tdmq role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceRoleAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a TDMQ role.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_instance = tencentcloud.tdmq.Instance("exampleInstance",
            cluster_name="tf_example",
            remark="remark.",
            tags={
                "createdBy": "terraform",
            })
        example_namespace = tencentcloud.tdmq.Namespace("exampleNamespace",
            environ_name="tf_example",
            msg_ttl=300,
            cluster_id=example_instance.id,
            retention_policy=tencentcloud.tdmq.NamespaceRetentionPolicyArgs(
                time_in_minutes=60,
                size_in_mb=10,
            ),
            remark="remark.")
        example_role = tencentcloud.tdmq.Role("exampleRole",
            role_name="tf_example",
            cluster_id=example_instance.id,
            remark="remark.")
        example_namespace_role_attachment = tencentcloud.tdmq.NamespaceRoleAttachment("exampleNamespaceRoleAttachment",
            environ_id=example_namespace.environ_name,
            role_name=example_role.role_name,
            permissions=[
                "produce",
                "consume",
            ],
            cluster_id=example_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param NamespaceRoleAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceRoleAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 environ_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamespaceRoleAttachmentArgs.__new__(NamespaceRoleAttachmentArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if environ_id is None and not opts.urn:
                raise TypeError("Missing required property 'environ_id'")
            __props__.__dict__["environ_id"] = environ_id
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            __props__.__dict__["create_time"] = None
        super(NamespaceRoleAttachment, __self__).__init__(
            'tencentcloud:Tdmq/namespaceRoleAttachment:NamespaceRoleAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            environ_id: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role_name: Optional[pulumi.Input[str]] = None) -> 'NamespaceRoleAttachment':
        """
        Get an existing NamespaceRoleAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The id of tdmq cluster.
        :param pulumi.Input[str] create_time: Creation time of resource.
        :param pulumi.Input[str] environ_id: The name of tdmq namespace.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The permissions of tdmq role.
        :param pulumi.Input[str] role_name: The name of tdmq role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NamespaceRoleAttachmentState.__new__(_NamespaceRoleAttachmentState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["environ_id"] = environ_id
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["role_name"] = role_name
        return NamespaceRoleAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The id of tdmq cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="environId")
    def environ_id(self) -> pulumi.Output[str]:
        """
        The name of tdmq namespace.
        """
        return pulumi.get(self, "environ_id")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        The permissions of tdmq role.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        The name of tdmq role.
        """
        return pulumi.get(self, "role_name")

