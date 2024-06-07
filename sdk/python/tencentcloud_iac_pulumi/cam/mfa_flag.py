# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MfaFlagArgs', 'MfaFlag']

@pulumi.input_type
class MfaFlagArgs:
    def __init__(__self__, *,
                 op_uin: pulumi.Input[int],
                 action_flag: Optional[pulumi.Input['MfaFlagActionFlagArgs']] = None,
                 login_flag: Optional[pulumi.Input['MfaFlagLoginFlagArgs']] = None):
        """
        The set of arguments for constructing a MfaFlag resource.
        :param pulumi.Input[int] op_uin: Operate uin.
        :param pulumi.Input['MfaFlagActionFlagArgs'] action_flag: Action flag setting.
        :param pulumi.Input['MfaFlagLoginFlagArgs'] login_flag: Login flag setting.
        """
        pulumi.set(__self__, "op_uin", op_uin)
        if action_flag is not None:
            pulumi.set(__self__, "action_flag", action_flag)
        if login_flag is not None:
            pulumi.set(__self__, "login_flag", login_flag)

    @property
    @pulumi.getter(name="opUin")
    def op_uin(self) -> pulumi.Input[int]:
        """
        Operate uin.
        """
        return pulumi.get(self, "op_uin")

    @op_uin.setter
    def op_uin(self, value: pulumi.Input[int]):
        pulumi.set(self, "op_uin", value)

    @property
    @pulumi.getter(name="actionFlag")
    def action_flag(self) -> Optional[pulumi.Input['MfaFlagActionFlagArgs']]:
        """
        Action flag setting.
        """
        return pulumi.get(self, "action_flag")

    @action_flag.setter
    def action_flag(self, value: Optional[pulumi.Input['MfaFlagActionFlagArgs']]):
        pulumi.set(self, "action_flag", value)

    @property
    @pulumi.getter(name="loginFlag")
    def login_flag(self) -> Optional[pulumi.Input['MfaFlagLoginFlagArgs']]:
        """
        Login flag setting.
        """
        return pulumi.get(self, "login_flag")

    @login_flag.setter
    def login_flag(self, value: Optional[pulumi.Input['MfaFlagLoginFlagArgs']]):
        pulumi.set(self, "login_flag", value)


@pulumi.input_type
class _MfaFlagState:
    def __init__(__self__, *,
                 action_flag: Optional[pulumi.Input['MfaFlagActionFlagArgs']] = None,
                 login_flag: Optional[pulumi.Input['MfaFlagLoginFlagArgs']] = None,
                 op_uin: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering MfaFlag resources.
        :param pulumi.Input['MfaFlagActionFlagArgs'] action_flag: Action flag setting.
        :param pulumi.Input['MfaFlagLoginFlagArgs'] login_flag: Login flag setting.
        :param pulumi.Input[int] op_uin: Operate uin.
        """
        if action_flag is not None:
            pulumi.set(__self__, "action_flag", action_flag)
        if login_flag is not None:
            pulumi.set(__self__, "login_flag", login_flag)
        if op_uin is not None:
            pulumi.set(__self__, "op_uin", op_uin)

    @property
    @pulumi.getter(name="actionFlag")
    def action_flag(self) -> Optional[pulumi.Input['MfaFlagActionFlagArgs']]:
        """
        Action flag setting.
        """
        return pulumi.get(self, "action_flag")

    @action_flag.setter
    def action_flag(self, value: Optional[pulumi.Input['MfaFlagActionFlagArgs']]):
        pulumi.set(self, "action_flag", value)

    @property
    @pulumi.getter(name="loginFlag")
    def login_flag(self) -> Optional[pulumi.Input['MfaFlagLoginFlagArgs']]:
        """
        Login flag setting.
        """
        return pulumi.get(self, "login_flag")

    @login_flag.setter
    def login_flag(self, value: Optional[pulumi.Input['MfaFlagLoginFlagArgs']]):
        pulumi.set(self, "login_flag", value)

    @property
    @pulumi.getter(name="opUin")
    def op_uin(self) -> Optional[pulumi.Input[int]]:
        """
        Operate uin.
        """
        return pulumi.get(self, "op_uin")

    @op_uin.setter
    def op_uin(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "op_uin", value)


class MfaFlag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_flag: Optional[pulumi.Input[pulumi.InputType['MfaFlagActionFlagArgs']]] = None,
                 login_flag: Optional[pulumi.Input[pulumi.InputType['MfaFlagLoginFlagArgs']]] = None,
                 op_uin: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a cam mfa_flag

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        info = tencentcloud.User.get_info()
        mfa_flag = tencentcloud.cam.MfaFlag("mfaFlag",
            op_uin=info.uin,
            login_flag=tencentcloud.cam.MfaFlagLoginFlagArgs(
                phone=0,
                stoken=1,
                wechat=0,
            ),
            action_flag=tencentcloud.cam.MfaFlagActionFlagArgs(
                phone=0,
                stoken=1,
                wechat=0,
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cam mfa_flag can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cam/mfaFlag:MfaFlag mfa_flag mfa_flag_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MfaFlagActionFlagArgs']] action_flag: Action flag setting.
        :param pulumi.Input[pulumi.InputType['MfaFlagLoginFlagArgs']] login_flag: Login flag setting.
        :param pulumi.Input[int] op_uin: Operate uin.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaFlagArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cam mfa_flag

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        info = tencentcloud.User.get_info()
        mfa_flag = tencentcloud.cam.MfaFlag("mfaFlag",
            op_uin=info.uin,
            login_flag=tencentcloud.cam.MfaFlagLoginFlagArgs(
                phone=0,
                stoken=1,
                wechat=0,
            ),
            action_flag=tencentcloud.cam.MfaFlagActionFlagArgs(
                phone=0,
                stoken=1,
                wechat=0,
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cam mfa_flag can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cam/mfaFlag:MfaFlag mfa_flag mfa_flag_id
        ```

        :param str resource_name: The name of the resource.
        :param MfaFlagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaFlagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_flag: Optional[pulumi.Input[pulumi.InputType['MfaFlagActionFlagArgs']]] = None,
                 login_flag: Optional[pulumi.Input[pulumi.InputType['MfaFlagLoginFlagArgs']]] = None,
                 op_uin: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaFlagArgs.__new__(MfaFlagArgs)

            __props__.__dict__["action_flag"] = action_flag
            __props__.__dict__["login_flag"] = login_flag
            if op_uin is None and not opts.urn:
                raise TypeError("Missing required property 'op_uin'")
            __props__.__dict__["op_uin"] = op_uin
        super(MfaFlag, __self__).__init__(
            'tencentcloud:Cam/mfaFlag:MfaFlag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action_flag: Optional[pulumi.Input[pulumi.InputType['MfaFlagActionFlagArgs']]] = None,
            login_flag: Optional[pulumi.Input[pulumi.InputType['MfaFlagLoginFlagArgs']]] = None,
            op_uin: Optional[pulumi.Input[int]] = None) -> 'MfaFlag':
        """
        Get an existing MfaFlag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MfaFlagActionFlagArgs']] action_flag: Action flag setting.
        :param pulumi.Input[pulumi.InputType['MfaFlagLoginFlagArgs']] login_flag: Login flag setting.
        :param pulumi.Input[int] op_uin: Operate uin.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaFlagState.__new__(_MfaFlagState)

        __props__.__dict__["action_flag"] = action_flag
        __props__.__dict__["login_flag"] = login_flag
        __props__.__dict__["op_uin"] = op_uin
        return MfaFlag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionFlag")
    def action_flag(self) -> pulumi.Output[Optional['outputs.MfaFlagActionFlag']]:
        """
        Action flag setting.
        """
        return pulumi.get(self, "action_flag")

    @property
    @pulumi.getter(name="loginFlag")
    def login_flag(self) -> pulumi.Output[Optional['outputs.MfaFlagLoginFlag']]:
        """
        Login flag setting.
        """
        return pulumi.get(self, "login_flag")

    @property
    @pulumi.getter(name="opUin")
    def op_uin(self) -> pulumi.Output[int]:
        """
        Operate uin.
        """
        return pulumi.get(self, "op_uin")

