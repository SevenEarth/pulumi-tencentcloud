# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'MfaFlagActionFlagArgs',
    'MfaFlagLoginFlagArgs',
    'PolicyVersionPolicyVersionArgs',
    'TagRoleAttachmentTagArgs',
]

@pulumi.input_type
class MfaFlagActionFlagArgs:
    def __init__(__self__, *,
                 phone: Optional[pulumi.Input[int]] = None,
                 stoken: Optional[pulumi.Input[int]] = None,
                 wechat: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] phone: Phone.
        :param pulumi.Input[int] stoken: Soft token.
        :param pulumi.Input[int] wechat: Wechat.
        """
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if stoken is not None:
            pulumi.set(__self__, "stoken", stoken)
        if wechat is not None:
            pulumi.set(__self__, "wechat", wechat)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[int]]:
        """
        Phone.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def stoken(self) -> Optional[pulumi.Input[int]]:
        """
        Soft token.
        """
        return pulumi.get(self, "stoken")

    @stoken.setter
    def stoken(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "stoken", value)

    @property
    @pulumi.getter
    def wechat(self) -> Optional[pulumi.Input[int]]:
        """
        Wechat.
        """
        return pulumi.get(self, "wechat")

    @wechat.setter
    def wechat(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wechat", value)


@pulumi.input_type
class MfaFlagLoginFlagArgs:
    def __init__(__self__, *,
                 phone: Optional[pulumi.Input[int]] = None,
                 stoken: Optional[pulumi.Input[int]] = None,
                 wechat: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] phone: Phone.
        :param pulumi.Input[int] stoken: Soft token.
        :param pulumi.Input[int] wechat: Wechat.
        """
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if stoken is not None:
            pulumi.set(__self__, "stoken", stoken)
        if wechat is not None:
            pulumi.set(__self__, "wechat", wechat)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[int]]:
        """
        Phone.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def stoken(self) -> Optional[pulumi.Input[int]]:
        """
        Soft token.
        """
        return pulumi.get(self, "stoken")

    @stoken.setter
    def stoken(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "stoken", value)

    @property
    @pulumi.getter
    def wechat(self) -> Optional[pulumi.Input[int]]:
        """
        Wechat.
        """
        return pulumi.get(self, "wechat")

    @wechat.setter
    def wechat(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wechat", value)


@pulumi.input_type
class PolicyVersionPolicyVersionArgs:
    def __init__(__self__, *,
                 create_date: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 is_default_version: Optional[pulumi.Input[int]] = None,
                 version_id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] create_date: Strategic version creation timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
        :param pulumi.Input[str] document: Strategic grammar textNote: This field may return NULL, indicating that the valid value cannot be obtained.
        :param pulumi.Input[int] is_default_version: Whether it is an effective version.0 means not, 1 means yesNote: This field may return NULL, indicating that the valid value cannot be obtained.
        :param pulumi.Input[int] version_id: Strategic version numberNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        if create_date is not None:
            pulumi.set(__self__, "create_date", create_date)
        if document is not None:
            pulumi.set(__self__, "document", document)
        if is_default_version is not None:
            pulumi.set(__self__, "is_default_version", is_default_version)
        if version_id is not None:
            pulumi.set(__self__, "version_id", version_id)

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> Optional[pulumi.Input[str]]:
        """
        Strategic version creation timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "create_date")

    @create_date.setter
    def create_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_date", value)

    @property
    @pulumi.getter
    def document(self) -> Optional[pulumi.Input[str]]:
        """
        Strategic grammar textNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "document")

    @document.setter
    def document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document", value)

    @property
    @pulumi.getter(name="isDefaultVersion")
    def is_default_version(self) -> Optional[pulumi.Input[int]]:
        """
        Whether it is an effective version.0 means not, 1 means yesNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "is_default_version")

    @is_default_version.setter
    def is_default_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "is_default_version", value)

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> Optional[pulumi.Input[int]]:
        """
        Strategic version numberNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "version_id")

    @version_id.setter
    def version_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version_id", value)


@pulumi.input_type
class TagRoleAttachmentTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: Label.
        :param pulumi.Input[str] value: Label.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Label.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Label.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


