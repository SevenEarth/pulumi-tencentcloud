# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'DatasourceCloudServiceType',
    'GetProjectListResult',
    'GetProjectListConfigListResult',
    'GetProjectListConfigListComponentResult',
    'GetUserProjectListResult',
]

@pulumi.output_type
class DatasourceCloudServiceType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "instanceId":
            suggest = "instance_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DatasourceCloudServiceType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DatasourceCloudServiceType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DatasourceCloudServiceType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 instance_id: str,
                 region: str,
                 type: str):
        """
        :param str instance_id: Instance Id.
        :param str region: Region.
        :param str type: Service type, Cloud.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Service type, Cloud.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetProjectListResult(dict):
    def __init__(__self__, *,
                 apply: bool,
                 auth_lists: Sequence[str],
                 color_code: str,
                 config_lists: Sequence['outputs.GetProjectListConfigListResult'],
                 corp_id: str,
                 created_at: str,
                 created_user: str,
                 id: int,
                 is_external_manage: bool,
                 last_modify_name: str,
                 logo: str,
                 manage_platform: str,
                 mark: str,
                 member_count: int,
                 name: str,
                 page_count: int,
                 panel_scope: str,
                 seed: str,
                 source: str,
                 updated_at: str,
                 updated_user: str):
        """
        :param bool apply: Apply(Note: This field may return null, indicating that no valid value can be obtained).
        :param Sequence[str] auth_lists: List of permissions within the project(Note: This field may return null, indicating that no valid value can be obtained).
        :param str color_code: Logo colour(Note: This field may return null, indicating that no valid value can be obtained).
        :param Sequence['GetProjectListConfigListArgs'] config_lists: Customized parameters, this parameter can be ignored(Note: This field may return null, indicating that no valid value can be obtained).
        :param str corp_id: Enterprise id(Note: This field may return null, indicating that no valid value can be obtained).
        :param str created_at: Created at(Note: This field may return null, indicating that no valid value can be obtained).
        :param str created_user: Created by(Note: This field may return null, indicating that no valid value can be obtained).
        :param int id: Project id.
        :param bool is_external_manage: Determine whether it is hosted(Note: This field may return null, indicating that no valid value can be obtained).
        :param str last_modify_name: Last modified report and presentation names(Note: This field may return null, indicating that no valid value can be obtained).
        :param str logo: Project logo(Note: This field may return null, indicating that no valid value can be obtained).
        :param str manage_platform: Hosting platform name(Note: This field may return null, indicating that no valid value can be obtained).
        :param str mark: Remark(Note: This field may return null, indicating that no valid value can be obtained).
        :param int member_count: Member count(Note: This field may return null, indicating that no valid value can be obtained).
        :param str name: Project name(Note: This field may return null, indicating that no valid value can be obtained).
        :param int page_count: Page count(Note: This field may return null, indicating that no valid value can be obtained).
        :param str panel_scope: Default kanban(Note: This field may return null, indicating that no valid value can be obtained).
        :param str seed: Obfuscated field(Note: This field may return null, indicating that no valid value can be obtained).
        :param str source: Interface call source(Note: This field may return null, indicating that no valid value can be obtained).
        :param str updated_at: Updated by(Note: This field may return null, indicating that no valid value can be obtained).
        :param str updated_user: Updated by(Note: This field may return null, indicating that no valid value can be obtained).
        """
        pulumi.set(__self__, "apply", apply)
        pulumi.set(__self__, "auth_lists", auth_lists)
        pulumi.set(__self__, "color_code", color_code)
        pulumi.set(__self__, "config_lists", config_lists)
        pulumi.set(__self__, "corp_id", corp_id)
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "created_user", created_user)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_external_manage", is_external_manage)
        pulumi.set(__self__, "last_modify_name", last_modify_name)
        pulumi.set(__self__, "logo", logo)
        pulumi.set(__self__, "manage_platform", manage_platform)
        pulumi.set(__self__, "mark", mark)
        pulumi.set(__self__, "member_count", member_count)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "page_count", page_count)
        pulumi.set(__self__, "panel_scope", panel_scope)
        pulumi.set(__self__, "seed", seed)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "updated_at", updated_at)
        pulumi.set(__self__, "updated_user", updated_user)

    @property
    @pulumi.getter
    def apply(self) -> bool:
        """
        Apply(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "apply")

    @property
    @pulumi.getter(name="authLists")
    def auth_lists(self) -> Sequence[str]:
        """
        List of permissions within the project(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "auth_lists")

    @property
    @pulumi.getter(name="colorCode")
    def color_code(self) -> str:
        """
        Logo colour(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "color_code")

    @property
    @pulumi.getter(name="configLists")
    def config_lists(self) -> Sequence['outputs.GetProjectListConfigListResult']:
        """
        Customized parameters, this parameter can be ignored(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "config_lists")

    @property
    @pulumi.getter(name="corpId")
    def corp_id(self) -> str:
        """
        Enterprise id(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "corp_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Created at(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdUser")
    def created_user(self) -> str:
        """
        Created by(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "created_user")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        Project id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isExternalManage")
    def is_external_manage(self) -> bool:
        """
        Determine whether it is hosted(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "is_external_manage")

    @property
    @pulumi.getter(name="lastModifyName")
    def last_modify_name(self) -> str:
        """
        Last modified report and presentation names(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "last_modify_name")

    @property
    @pulumi.getter
    def logo(self) -> str:
        """
        Project logo(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "logo")

    @property
    @pulumi.getter(name="managePlatform")
    def manage_platform(self) -> str:
        """
        Hosting platform name(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "manage_platform")

    @property
    @pulumi.getter
    def mark(self) -> str:
        """
        Remark(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "mark")

    @property
    @pulumi.getter(name="memberCount")
    def member_count(self) -> int:
        """
        Member count(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "member_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Project name(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pageCount")
    def page_count(self) -> int:
        """
        Page count(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "page_count")

    @property
    @pulumi.getter(name="panelScope")
    def panel_scope(self) -> str:
        """
        Default kanban(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "panel_scope")

    @property
    @pulumi.getter
    def seed(self) -> str:
        """
        Obfuscated field(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "seed")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        Interface call source(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Updated by(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="updatedUser")
    def updated_user(self) -> str:
        """
        Updated by(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "updated_user")


@pulumi.output_type
class GetProjectListConfigListResult(dict):
    def __init__(__self__, *,
                 components: Sequence['outputs.GetProjectListConfigListComponentResult'],
                 module_group: str):
        """
        :param Sequence['GetProjectListConfigListComponentArgs'] components: Components(Note: This field may return null, indicating that no valid value can be obtained).
        :param str module_group: Module group(Note: This field may return null, indicating that no valid value can be obtained).
        """
        pulumi.set(__self__, "components", components)
        pulumi.set(__self__, "module_group", module_group)

    @property
    @pulumi.getter
    def components(self) -> Sequence['outputs.GetProjectListConfigListComponentResult']:
        """
        Components(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "components")

    @property
    @pulumi.getter(name="moduleGroup")
    def module_group(self) -> str:
        """
        Module group(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "module_group")


@pulumi.output_type
class GetProjectListConfigListComponentResult(dict):
    def __init__(__self__, *,
                 include_type: str,
                 module_id: str,
                 params: str):
        """
        :param str include_type: Include type(Note: This field may return null, indicating that no valid value can be obtained).
        :param str module_id: Module id(Note: This field may return null, indicating that no valid value can be obtained).
        :param str params: Extra parameters(Note: This field may return null, indicating that no valid value can be obtained).
        """
        pulumi.set(__self__, "include_type", include_type)
        pulumi.set(__self__, "module_id", module_id)
        pulumi.set(__self__, "params", params)

    @property
    @pulumi.getter(name="includeType")
    def include_type(self) -> str:
        """
        Include type(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "include_type")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> str:
        """
        Module id(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def params(self) -> str:
        """
        Extra parameters(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "params")


@pulumi.output_type
class GetUserProjectListResult(dict):
    def __init__(__self__, *,
                 area_code: str,
                 corp_id: str,
                 created_at: str,
                 created_user: str,
                 email: str,
                 first_modify: int,
                 global_user_name: str,
                 last_login: str,
                 mobile: str,
                 phone_number: str,
                 status: int,
                 updated_at: str,
                 updated_user: str,
                 user_id: str,
                 user_name: str):
        """
        :param str area_code: Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
        :param str corp_id: Enterprise id(Note: This field may return null, indicating that no valid value can be obtained).
        :param str created_at: Created at(Note: This field may return null, indicating that no valid value can be obtained).
        :param str created_user: Created by(Note: This field may return null, indicating that no valid value can be obtained).
        :param str email: E-mail(Note: This field may return null, indicating that no valid value can be obtained).
        :param int first_modify: First login to change password, public cloud unrelated fields(Note: This field may return null, indicating that no valid value can be obtained).
        :param str global_user_name: Global role name(Note: This field may return null, indicating that no valid value can be obtained).
        :param str last_login: Last login time, public cloud unrelated fields(Note: This field may return null, indicating that no valid value can be obtained).
        :param str mobile: Mobile number, public cloud unrelated fields(Note: This field may return null, indicating that no valid value can be obtained).
        :param str phone_number: Phone number(Note: This field may return null, indicating that no valid value can be obtained).
        :param int status: Disabled state(Note: This field may return null, indicating that no valid value can be obtained).
        :param str updated_at: Updated at(Note: This field may return null, indicating that no valid value can be obtained).
        :param str updated_user: Updated by(Note: This field may return null, indicating that no valid value can be obtained).
        :param str user_id: User id.
        :param str user_name: Username.
        """
        pulumi.set(__self__, "area_code", area_code)
        pulumi.set(__self__, "corp_id", corp_id)
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "created_user", created_user)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "first_modify", first_modify)
        pulumi.set(__self__, "global_user_name", global_user_name)
        pulumi.set(__self__, "last_login", last_login)
        pulumi.set(__self__, "mobile", mobile)
        pulumi.set(__self__, "phone_number", phone_number)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "updated_at", updated_at)
        pulumi.set(__self__, "updated_user", updated_user)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="areaCode")
    def area_code(self) -> str:
        """
        Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "area_code")

    @property
    @pulumi.getter(name="corpId")
    def corp_id(self) -> str:
        """
        Enterprise id(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "corp_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Created at(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdUser")
    def created_user(self) -> str:
        """
        Created by(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "created_user")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        E-mail(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstModify")
    def first_modify(self) -> int:
        """
        First login to change password, public cloud unrelated fields(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "first_modify")

    @property
    @pulumi.getter(name="globalUserName")
    def global_user_name(self) -> str:
        """
        Global role name(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "global_user_name")

    @property
    @pulumi.getter(name="lastLogin")
    def last_login(self) -> str:
        """
        Last login time, public cloud unrelated fields(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "last_login")

    @property
    @pulumi.getter
    def mobile(self) -> str:
        """
        Mobile number, public cloud unrelated fields(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "mobile")

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> str:
        """
        Phone number(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "phone_number")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        Disabled state(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Updated at(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="updatedUser")
    def updated_user(self) -> str:
        """
        Updated by(Note: This field may return null, indicating that no valid value can be obtained).
        """
        return pulumi.get(self, "updated_user")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        User id.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        Username.
        """
        return pulumi.get(self, "user_name")


