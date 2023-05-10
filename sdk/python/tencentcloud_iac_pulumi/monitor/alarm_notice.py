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

__all__ = ['AlarmNoticeArgs', 'AlarmNotice']

@pulumi.input_type
class AlarmNoticeArgs:
    def __init__(__self__, *,
                 notice_language: pulumi.Input[str],
                 notice_type: pulumi.Input[str],
                 cls_notices: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 url_notices: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]]] = None,
                 user_notices: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]]] = None):
        """
        The set of arguments for constructing a AlarmNotice resource.
        :param pulumi.Input[str] notice_language: Notification language zh-CN=Chinese en-US=English.
        :param pulumi.Input[str] notice_type: Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]] cls_notices: A maximum of one alarm notification can be pushed to the CLS service.
        :param pulumi.Input[str] name: Notification template name within 60.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]] url_notices: The maximum number of callback notifications is 3.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]] user_notices: Alarm notification template list.(At most five).
        """
        pulumi.set(__self__, "notice_language", notice_language)
        pulumi.set(__self__, "notice_type", notice_type)
        if cls_notices is not None:
            pulumi.set(__self__, "cls_notices", cls_notices)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if url_notices is not None:
            pulumi.set(__self__, "url_notices", url_notices)
        if user_notices is not None:
            pulumi.set(__self__, "user_notices", user_notices)

    @property
    @pulumi.getter(name="noticeLanguage")
    def notice_language(self) -> pulumi.Input[str]:
        """
        Notification language zh-CN=Chinese en-US=English.
        """
        return pulumi.get(self, "notice_language")

    @notice_language.setter
    def notice_language(self, value: pulumi.Input[str]):
        pulumi.set(self, "notice_language", value)

    @property
    @pulumi.getter(name="noticeType")
    def notice_type(self) -> pulumi.Input[str]:
        """
        Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        """
        return pulumi.get(self, "notice_type")

    @notice_type.setter
    def notice_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "notice_type", value)

    @property
    @pulumi.getter(name="clsNotices")
    def cls_notices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]]]:
        """
        A maximum of one alarm notification can be pushed to the CLS service.
        """
        return pulumi.get(self, "cls_notices")

    @cls_notices.setter
    def cls_notices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]]]):
        pulumi.set(self, "cls_notices", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Notification template name within 60.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="urlNotices")
    def url_notices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]]]:
        """
        The maximum number of callback notifications is 3.
        """
        return pulumi.get(self, "url_notices")

    @url_notices.setter
    def url_notices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]]]):
        pulumi.set(self, "url_notices", value)

    @property
    @pulumi.getter(name="userNotices")
    def user_notices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]]]:
        """
        Alarm notification template list.(At most five).
        """
        return pulumi.get(self, "user_notices")

    @user_notices.setter
    def user_notices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]]]):
        pulumi.set(self, "user_notices", value)


@pulumi.input_type
class _AlarmNoticeState:
    def __init__(__self__, *,
                 cls_notices: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]]] = None,
                 is_preset: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notice_language: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[str]] = None,
                 policy_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 updated_by: Optional[pulumi.Input[str]] = None,
                 url_notices: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]]] = None,
                 user_notices: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]]] = None):
        """
        Input properties used for looking up and filtering AlarmNotice resources.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]] cls_notices: A maximum of one alarm notification can be pushed to the CLS service.
        :param pulumi.Input[int] is_preset: Whether it is the system default notification template 0=No 1=Yes.
        :param pulumi.Input[str] name: Notification template name within 60.
        :param pulumi.Input[str] notice_language: Notification language zh-CN=Chinese en-US=English.
        :param pulumi.Input[str] notice_type: Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_ids: List of alarm policy IDs bound to the alarm notification template.
        :param pulumi.Input[str] updated_at: Last modified time.
        :param pulumi.Input[str] updated_by: Last Modified By.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]] url_notices: The maximum number of callback notifications is 3.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]] user_notices: Alarm notification template list.(At most five).
        """
        if cls_notices is not None:
            pulumi.set(__self__, "cls_notices", cls_notices)
        if is_preset is not None:
            pulumi.set(__self__, "is_preset", is_preset)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notice_language is not None:
            pulumi.set(__self__, "notice_language", notice_language)
        if notice_type is not None:
            pulumi.set(__self__, "notice_type", notice_type)
        if policy_ids is not None:
            pulumi.set(__self__, "policy_ids", policy_ids)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if updated_by is not None:
            pulumi.set(__self__, "updated_by", updated_by)
        if url_notices is not None:
            pulumi.set(__self__, "url_notices", url_notices)
        if user_notices is not None:
            pulumi.set(__self__, "user_notices", user_notices)

    @property
    @pulumi.getter(name="clsNotices")
    def cls_notices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]]]:
        """
        A maximum of one alarm notification can be pushed to the CLS service.
        """
        return pulumi.get(self, "cls_notices")

    @cls_notices.setter
    def cls_notices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeClsNoticeArgs']]]]):
        pulumi.set(self, "cls_notices", value)

    @property
    @pulumi.getter(name="isPreset")
    def is_preset(self) -> Optional[pulumi.Input[int]]:
        """
        Whether it is the system default notification template 0=No 1=Yes.
        """
        return pulumi.get(self, "is_preset")

    @is_preset.setter
    def is_preset(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "is_preset", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Notification template name within 60.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noticeLanguage")
    def notice_language(self) -> Optional[pulumi.Input[str]]:
        """
        Notification language zh-CN=Chinese en-US=English.
        """
        return pulumi.get(self, "notice_language")

    @notice_language.setter
    def notice_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notice_language", value)

    @property
    @pulumi.getter(name="noticeType")
    def notice_type(self) -> Optional[pulumi.Input[str]]:
        """
        Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        """
        return pulumi.get(self, "notice_type")

    @notice_type.setter
    def notice_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notice_type", value)

    @property
    @pulumi.getter(name="policyIds")
    def policy_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alarm policy IDs bound to the alarm notification template.
        """
        return pulumi.get(self, "policy_ids")

    @policy_ids.setter
    def policy_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policy_ids", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Last modified time.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> Optional[pulumi.Input[str]]:
        """
        Last Modified By.
        """
        return pulumi.get(self, "updated_by")

    @updated_by.setter
    def updated_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_by", value)

    @property
    @pulumi.getter(name="urlNotices")
    def url_notices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]]]:
        """
        The maximum number of callback notifications is 3.
        """
        return pulumi.get(self, "url_notices")

    @url_notices.setter
    def url_notices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUrlNoticeArgs']]]]):
        pulumi.set(self, "url_notices", value)

    @property
    @pulumi.getter(name="userNotices")
    def user_notices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]]]:
        """
        Alarm notification template list.(At most five).
        """
        return pulumi.get(self, "user_notices")

    @user_notices.setter
    def user_notices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmNoticeUserNoticeArgs']]]]):
        pulumi.set(self, "user_notices", value)


class AlarmNotice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cls_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeClsNoticeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notice_language: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[str]] = None,
                 url_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUrlNoticeArgs']]]]] = None,
                 user_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUserNoticeArgs']]]]] = None,
                 __props__=None):
        """
        Provides a alarm notice resource for monitor.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.monitor.AlarmNotice("example",
            notice_language="zh-CN",
            notice_type="ALL",
            url_notices=[tencentcloud.monitor.AlarmNoticeUrlNoticeArgs(
                end_time=0,
                start_time=1,
                url="https://www.mytest.com/validate",
                weekdays=[
                    1,
                    2,
                    3,
                    4,
                    5,
                    6,
                    7,
                ],
            )],
            user_notices=[tencentcloud.monitor.AlarmNoticeUserNoticeArgs(
                end_time=1,
                group_ids=[],
                need_phone_arrive_notice=1,
                notice_ways=[
                    "SMS",
                    "EMAIL",
                ],
                phone_call_type="CIRCLE",
                phone_circle_interval=50,
                phone_circle_times=2,
                phone_inner_interval=60,
                phone_orders=[10001],
                receiver_type="USER",
                start_time=0,
                user_ids=[10001],
                weekdays=[
                    1,
                    2,
                    3,
                    4,
                    5,
                    6,
                    7,
                ],
            )])
        ```

        ## Import

        Monitor Alarm Notice can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/alarmNotice:AlarmNotice import-test noticeId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeClsNoticeArgs']]]] cls_notices: A maximum of one alarm notification can be pushed to the CLS service.
        :param pulumi.Input[str] name: Notification template name within 60.
        :param pulumi.Input[str] notice_language: Notification language zh-CN=Chinese en-US=English.
        :param pulumi.Input[str] notice_type: Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUrlNoticeArgs']]]] url_notices: The maximum number of callback notifications is 3.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUserNoticeArgs']]]] user_notices: Alarm notification template list.(At most five).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlarmNoticeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a alarm notice resource for monitor.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.monitor.AlarmNotice("example",
            notice_language="zh-CN",
            notice_type="ALL",
            url_notices=[tencentcloud.monitor.AlarmNoticeUrlNoticeArgs(
                end_time=0,
                start_time=1,
                url="https://www.mytest.com/validate",
                weekdays=[
                    1,
                    2,
                    3,
                    4,
                    5,
                    6,
                    7,
                ],
            )],
            user_notices=[tencentcloud.monitor.AlarmNoticeUserNoticeArgs(
                end_time=1,
                group_ids=[],
                need_phone_arrive_notice=1,
                notice_ways=[
                    "SMS",
                    "EMAIL",
                ],
                phone_call_type="CIRCLE",
                phone_circle_interval=50,
                phone_circle_times=2,
                phone_inner_interval=60,
                phone_orders=[10001],
                receiver_type="USER",
                start_time=0,
                user_ids=[10001],
                weekdays=[
                    1,
                    2,
                    3,
                    4,
                    5,
                    6,
                    7,
                ],
            )])
        ```

        ## Import

        Monitor Alarm Notice can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/alarmNotice:AlarmNotice import-test noticeId
        ```

        :param str resource_name: The name of the resource.
        :param AlarmNoticeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlarmNoticeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cls_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeClsNoticeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notice_language: Optional[pulumi.Input[str]] = None,
                 notice_type: Optional[pulumi.Input[str]] = None,
                 url_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUrlNoticeArgs']]]]] = None,
                 user_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUserNoticeArgs']]]]] = None,
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
            __props__ = AlarmNoticeArgs.__new__(AlarmNoticeArgs)

            __props__.__dict__["cls_notices"] = cls_notices
            __props__.__dict__["name"] = name
            if notice_language is None and not opts.urn:
                raise TypeError("Missing required property 'notice_language'")
            __props__.__dict__["notice_language"] = notice_language
            if notice_type is None and not opts.urn:
                raise TypeError("Missing required property 'notice_type'")
            __props__.__dict__["notice_type"] = notice_type
            __props__.__dict__["url_notices"] = url_notices
            __props__.__dict__["user_notices"] = user_notices
            __props__.__dict__["is_preset"] = None
            __props__.__dict__["policy_ids"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["updated_by"] = None
        super(AlarmNotice, __self__).__init__(
            'tencentcloud:Monitor/alarmNotice:AlarmNotice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cls_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeClsNoticeArgs']]]]] = None,
            is_preset: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notice_language: Optional[pulumi.Input[str]] = None,
            notice_type: Optional[pulumi.Input[str]] = None,
            policy_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            updated_by: Optional[pulumi.Input[str]] = None,
            url_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUrlNoticeArgs']]]]] = None,
            user_notices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUserNoticeArgs']]]]] = None) -> 'AlarmNotice':
        """
        Get an existing AlarmNotice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeClsNoticeArgs']]]] cls_notices: A maximum of one alarm notification can be pushed to the CLS service.
        :param pulumi.Input[int] is_preset: Whether it is the system default notification template 0=No 1=Yes.
        :param pulumi.Input[str] name: Notification template name within 60.
        :param pulumi.Input[str] notice_language: Notification language zh-CN=Chinese en-US=English.
        :param pulumi.Input[str] notice_type: Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_ids: List of alarm policy IDs bound to the alarm notification template.
        :param pulumi.Input[str] updated_at: Last modified time.
        :param pulumi.Input[str] updated_by: Last Modified By.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUrlNoticeArgs']]]] url_notices: The maximum number of callback notifications is 3.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmNoticeUserNoticeArgs']]]] user_notices: Alarm notification template list.(At most five).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlarmNoticeState.__new__(_AlarmNoticeState)

        __props__.__dict__["cls_notices"] = cls_notices
        __props__.__dict__["is_preset"] = is_preset
        __props__.__dict__["name"] = name
        __props__.__dict__["notice_language"] = notice_language
        __props__.__dict__["notice_type"] = notice_type
        __props__.__dict__["policy_ids"] = policy_ids
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["updated_by"] = updated_by
        __props__.__dict__["url_notices"] = url_notices
        __props__.__dict__["user_notices"] = user_notices
        return AlarmNotice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clsNotices")
    def cls_notices(self) -> pulumi.Output[Optional[Sequence['outputs.AlarmNoticeClsNotice']]]:
        """
        A maximum of one alarm notification can be pushed to the CLS service.
        """
        return pulumi.get(self, "cls_notices")

    @property
    @pulumi.getter(name="isPreset")
    def is_preset(self) -> pulumi.Output[int]:
        """
        Whether it is the system default notification template 0=No 1=Yes.
        """
        return pulumi.get(self, "is_preset")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Notification template name within 60.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noticeLanguage")
    def notice_language(self) -> pulumi.Output[str]:
        """
        Notification language zh-CN=Chinese en-US=English.
        """
        return pulumi.get(self, "notice_language")

    @property
    @pulumi.getter(name="noticeType")
    def notice_type(self) -> pulumi.Output[str]:
        """
        Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
        """
        return pulumi.get(self, "notice_type")

    @property
    @pulumi.getter(name="policyIds")
    def policy_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of alarm policy IDs bound to the alarm notification template.
        """
        return pulumi.get(self, "policy_ids")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Last modified time.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> pulumi.Output[str]:
        """
        Last Modified By.
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="urlNotices")
    def url_notices(self) -> pulumi.Output[Optional[Sequence['outputs.AlarmNoticeUrlNotice']]]:
        """
        The maximum number of callback notifications is 3.
        """
        return pulumi.get(self, "url_notices")

    @property
    @pulumi.getter(name="userNotices")
    def user_notices(self) -> pulumi.Output[Optional[Sequence['outputs.AlarmNoticeUserNotice']]]:
        """
        Alarm notification template list.(At most five).
        """
        return pulumi.get(self, "user_notices")

