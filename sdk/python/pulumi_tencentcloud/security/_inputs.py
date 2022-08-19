# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GroupRuleAddressTemplateArgs',
    'GroupRuleProtocolTemplateArgs',
]

@pulumi.input_type
class GroupRuleAddressTemplateArgs:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] group_id: Address template group ID, conflicts with `template_id`.
        :param pulumi.Input[str] template_id: Address template ID, conflicts with `group_id`.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Address template group ID, conflicts with `template_id`.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Address template ID, conflicts with `group_id`.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


@pulumi.input_type
class GroupRuleProtocolTemplateArgs:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] group_id: Address template group ID, conflicts with `template_id`.
        :param pulumi.Input[str] template_id: Address template ID, conflicts with `group_id`.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Address template group ID, conflicts with `template_id`.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Address template ID, conflicts with `group_id`.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


