# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RocketmqRoleArgs', 'RocketmqRole']

@pulumi.input_type
class RocketmqRoleArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 role_name: pulumi.Input[str],
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RocketmqRole resource.
        :param pulumi.Input[str] cluster_id: Cluster ID (required).
        :param pulumi.Input[str] role_name: Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "role_name", role_name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID (required).
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        """
        Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _RocketmqRoleState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RocketmqRole resources.
        :param pulumi.Input[str] cluster_id: Cluster ID (required).
        :param pulumi.Input[str] create_time: Creation time.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        :param pulumi.Input[str] role_name: Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] token: Value of the role token.
        :param pulumi.Input[str] update_time: Update time.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID (required).
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Value of the role token.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Update time.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class RocketmqRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tdmqRocketmq role

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_rocketmq_cluster = tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster",
            cluster_name="tf_example",
            remark="remark.")
        example_rocketmq_role = tencentcloud.tdmq.RocketmqRole("exampleRocketmqRole",
            cluster_id=example_rocketmq_cluster.cluster_id,
            role_name="tf_example",
            remark="remark.")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tdmqRocketmq role can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rocketmqRole:RocketmqRole role role_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID (required).
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        :param pulumi.Input[str] role_name: Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tdmqRocketmq role

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_rocketmq_cluster = tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster",
            cluster_name="tf_example",
            remark="remark.")
        example_rocketmq_role = tencentcloud.tdmq.RocketmqRole("exampleRocketmqRole",
            cluster_id=example_rocketmq_cluster.cluster_id,
            role_name="tf_example",
            remark="remark.")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tdmqRocketmq role can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rocketmqRole:RocketmqRole role role_id
        ```

        :param str resource_name: The name of the resource.
        :param RocketmqRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RocketmqRoleArgs.__new__(RocketmqRoleArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["remark"] = remark
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["token"] = None
            __props__.__dict__["update_time"] = None
        super(RocketmqRole, __self__).__init__(
            'tencentcloud:Tdmq/rocketmqRole:RocketmqRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'RocketmqRole':
        """
        Get an existing RocketmqRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID (required).
        :param pulumi.Input[str] create_time: Creation time.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        :param pulumi.Input[str] role_name: Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] token: Value of the role token.
        :param pulumi.Input[str] update_time: Update time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqRoleState.__new__(_RocketmqRoleState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["remark"] = remark
        __props__.__dict__["role_name"] = role_name
        __props__.__dict__["token"] = token
        __props__.__dict__["update_time"] = update_time
        return RocketmqRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID (required).
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        Value of the role token.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Update time.
        """
        return pulumi.get(self, "update_time")

