# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RabbitmqUserArgs', 'RabbitmqUser']

@pulumi.input_type
class RabbitmqUserArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 user: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 max_channels: Optional[pulumi.Input[int]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RabbitmqUser resource.
        :param pulumi.Input[str] instance_id: Cluster instance ID.
        :param pulumi.Input[str] password: Password, used when logging in.
        :param pulumi.Input[str] user: Username, used when logging in.
        :param pulumi.Input[str] description: Describe.
        :param pulumi.Input[int] max_channels: The maximum number of channels for this user, if not filled in, there is no limit.
        :param pulumi.Input[int] max_connections: The maximum number of connections for this user, if not filled in, there is no limit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "user", user)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if max_channels is not None:
            pulumi.set(__self__, "max_channels", max_channels)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Cluster instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password, used when logging in.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        Username, used when logging in.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Describe.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="maxChannels")
    def max_channels(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of channels for this user, if not filled in, there is no limit.
        """
        return pulumi.get(self, "max_channels")

    @max_channels.setter
    def max_channels(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_channels", value)

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of connections for this user, if not filled in, there is no limit.
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RabbitmqUserState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_channels: Optional[pulumi.Input[int]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RabbitmqUser resources.
        :param pulumi.Input[str] description: Describe.
        :param pulumi.Input[str] instance_id: Cluster instance ID.
        :param pulumi.Input[int] max_channels: The maximum number of channels for this user, if not filled in, there is no limit.
        :param pulumi.Input[int] max_connections: The maximum number of connections for this user, if not filled in, there is no limit.
        :param pulumi.Input[str] password: Password, used when logging in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        :param pulumi.Input[str] user: Username, used when logging in.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_channels is not None:
            pulumi.set(__self__, "max_channels", max_channels)
        if max_connections is not None:
            pulumi.set(__self__, "max_connections", max_connections)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Describe.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxChannels")
    def max_channels(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of channels for this user, if not filled in, there is no limit.
        """
        return pulumi.get(self, "max_channels")

    @max_channels.setter
    def max_channels(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_channels", value)

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of connections for this user, if not filled in, there is no limit.
        """
        return pulumi.get(self, "max_connections")

    @max_connections.setter
    def max_connections(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_connections", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password, used when logging in.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Username, used when logging in.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class RabbitmqUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_channels: Optional[pulumi.Input[int]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tdmq rabbitmq_user

        ## Import

        tdmq rabbitmq_user can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rabbitmqUser:RabbitmqUser example amqp-8xzx822q#tf-example-user
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Describe.
        :param pulumi.Input[str] instance_id: Cluster instance ID.
        :param pulumi.Input[int] max_channels: The maximum number of channels for this user, if not filled in, there is no limit.
        :param pulumi.Input[int] max_connections: The maximum number of connections for this user, if not filled in, there is no limit.
        :param pulumi.Input[str] password: Password, used when logging in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        :param pulumi.Input[str] user: Username, used when logging in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RabbitmqUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tdmq rabbitmq_user

        ## Import

        tdmq rabbitmq_user can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rabbitmqUser:RabbitmqUser example amqp-8xzx822q#tf-example-user
        ```

        :param str resource_name: The name of the resource.
        :param RabbitmqUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RabbitmqUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_channels: Optional[pulumi.Input[int]] = None,
                 max_connections: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RabbitmqUserArgs.__new__(RabbitmqUserArgs)

            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["max_channels"] = max_channels
            __props__.__dict__["max_connections"] = max_connections
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["tags"] = tags
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(RabbitmqUser, __self__).__init__(
            'tencentcloud:Tdmq/rabbitmqUser:RabbitmqUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            max_channels: Optional[pulumi.Input[int]] = None,
            max_connections: Optional[pulumi.Input[int]] = None,
            password: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'RabbitmqUser':
        """
        Get an existing RabbitmqUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Describe.
        :param pulumi.Input[str] instance_id: Cluster instance ID.
        :param pulumi.Input[int] max_channels: The maximum number of channels for this user, if not filled in, there is no limit.
        :param pulumi.Input[int] max_connections: The maximum number of connections for this user, if not filled in, there is no limit.
        :param pulumi.Input[str] password: Password, used when logging in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        :param pulumi.Input[str] user: Username, used when logging in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RabbitmqUserState.__new__(_RabbitmqUserState)

        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_channels"] = max_channels
        __props__.__dict__["max_connections"] = max_connections
        __props__.__dict__["password"] = password
        __props__.__dict__["tags"] = tags
        __props__.__dict__["user"] = user
        return RabbitmqUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Describe.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Cluster instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxChannels")
    def max_channels(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum number of channels for this user, if not filled in, there is no limit.
        """
        return pulumi.get(self, "max_channels")

    @property
    @pulumi.getter(name="maxConnections")
    def max_connections(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum number of connections for this user, if not filled in, there is no limit.
        """
        return pulumi.get(self, "max_connections")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password, used when logging in.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular console user, monitoring: management console user, other values: non console user.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Username, used when logging in.
        """
        return pulumi.get(self, "user")

