# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TokenArgs', 'Token']

@pulumi.input_type
class TokenArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Token resource.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[str] description: Description of the token. Valid length is [0~255].
        :param pulumi.Input[bool] enable: Indicate to enable this token or not.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the token. Valid length is [0~255].
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate to enable this token or not.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)


@pulumi.input_type
class _TokenState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 token_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Token resources.
        :param pulumi.Input[str] create_time: Create time.
        :param pulumi.Input[str] description: Description of the token. Valid length is [0~255].
        :param pulumi.Input[bool] enable: Indicate to enable this token or not.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[str] token: The content of the token.
        :param pulumi.Input[str] token_id: Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
        :param pulumi.Input[str] user_name: User name of the token.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if token_id is not None:
            pulumi.set(__self__, "token_id", token_id)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the token. Valid length is [0~255].
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate to enable this token or not.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the token.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="tokenId")
    def token_id(self) -> Optional[pulumi.Input[str]]:
        """
        Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
        """
        return pulumi.get(self, "token_id")

    @token_id.setter
    def token_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        User name of the token.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class Token(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create tcr long term token.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo = tencentcloud.tcr.Token("foo",
            description="test",
            instance_id="cls-cda1iex1")
        ```

        ## Import

        tcr token can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcr/token:Token foo cls-cda1iex1#namespace#buv3h3j96j2d1rk1cllg
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the token. Valid length is [0~255].
        :param pulumi.Input[bool] enable: Indicate to enable this token or not.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create tcr long term token.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo = tencentcloud.tcr.Token("foo",
            description="test",
            instance_id="cls-cda1iex1")
        ```

        ## Import

        tcr token can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcr/token:Token foo cls-cda1iex1#namespace#buv3h3j96j2d1rk1cllg
        ```

        :param str resource_name: The name of the resource.
        :param TokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TokenArgs.__new__(TokenArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["enable"] = enable
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["token"] = None
            __props__.__dict__["token_id"] = None
            __props__.__dict__["user_name"] = None
        super(Token, __self__).__init__(
            'tencentcloud:Tcr/token:Token',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            token_id: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'Token':
        """
        Get an existing Token resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time.
        :param pulumi.Input[str] description: Description of the token. Valid length is [0~255].
        :param pulumi.Input[bool] enable: Indicate to enable this token or not.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[str] token: The content of the token.
        :param pulumi.Input[str] token_id: Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
        :param pulumi.Input[str] user_name: User name of the token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TokenState.__new__(_TokenState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["enable"] = enable
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["token"] = token
        __props__.__dict__["token_id"] = token_id
        __props__.__dict__["user_name"] = user_name
        return Token(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the token. Valid length is [0~255].
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate to enable this token or not.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The content of the token.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="tokenId")
    def token_id(self) -> pulumi.Output[str]:
        """
        Sub ID of the TCR token. The full ID of token format like `instance_id#token_id`.
        """
        return pulumi.get(self, "token_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        User name of the token.
        """
        return pulumi.get(self, "user_name")

