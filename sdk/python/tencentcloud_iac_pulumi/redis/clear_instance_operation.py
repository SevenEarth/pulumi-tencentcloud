# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClearInstanceOperationArgs', 'ClearInstanceOperation']

@pulumi.input_type
class ClearInstanceOperationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 password: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ClearInstanceOperation resource.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] password: Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


@pulumi.input_type
class _ClearInstanceOperationState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClearInstanceOperation resources.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] password: Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


class ClearInstanceOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a redis clear_instance_operation

        ## Example Usage

        ### Clear the instance data of the Redis instance

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        password = config.get("password")
        if password is None:
            password = "test12345789"
        zone = tencentcloud.Redis.get_zone_config(type_id=7)
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=zone.lists[1].zone,
            cidr_block="10.0.1.0/24")
        foo = tencentcloud.redis.Instance("foo",
            availability_zone=zone.lists[1].zone,
            type_id=zone.lists[1].type_id,
            password=password,
            mem_size=8192,
            redis_shard_num=zone.lists[1].redis_shard_nums[0],
            redis_replicas_num=zone.lists[1].redis_replicas_nums[0],
            port=6379,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        clear_instance_operation = tencentcloud.redis.ClearInstanceOperation("clearInstanceOperation",
            instance_id=foo.id,
            password=password)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] password: Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClearInstanceOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a redis clear_instance_operation

        ## Example Usage

        ### Clear the instance data of the Redis instance

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        password = config.get("password")
        if password is None:
            password = "test12345789"
        zone = tencentcloud.Redis.get_zone_config(type_id=7)
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=zone.lists[1].zone,
            cidr_block="10.0.1.0/24")
        foo = tencentcloud.redis.Instance("foo",
            availability_zone=zone.lists[1].zone,
            type_id=zone.lists[1].type_id,
            password=password,
            mem_size=8192,
            redis_shard_num=zone.lists[1].redis_shard_nums[0],
            redis_replicas_num=zone.lists[1].redis_replicas_nums[0],
            port=6379,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        clear_instance_operation = tencentcloud.redis.ClearInstanceOperation("clearInstanceOperation",
            instance_id=foo.id,
            password=password)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ClearInstanceOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClearInstanceOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClearInstanceOperationArgs.__new__(ClearInstanceOperationArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ClearInstanceOperation, __self__).__init__(
            'tencentcloud:Redis/clearInstanceOperation:ClearInstanceOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None) -> 'ClearInstanceOperation':
        """
        Get an existing ClearInstanceOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The ID of instance.
        :param pulumi.Input[str] password: Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClearInstanceOperationState.__new__(_ClearInstanceOperationState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["password"] = password
        return ClearInstanceOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
        """
        return pulumi.get(self, "password")

