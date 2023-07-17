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
    'AccessRuleAccessRule',
    'LifeCycleRuleLifeCycleRule',
    'LifeCycleRuleLifeCycleRuleTransition',
    'GetAccessGroupsAccessGroupResult',
    'GetFileSystemsFileSystemResult',
    'GetMountPointsMountPointResult',
]

@pulumi.output_type
class AccessRuleAccessRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessMode":
            suggest = "access_mode"
        elif key == "accessRuleId":
            suggest = "access_rule_id"
        elif key == "createTime":
            suggest = "create_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessRuleAccessRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessRuleAccessRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessRuleAccessRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_mode: Optional[int] = None,
                 access_rule_id: Optional[int] = None,
                 address: Optional[str] = None,
                 create_time: Optional[str] = None,
                 priority: Optional[int] = None):
        """
        :param int access_mode: rule access mode, 1: read only, 2: read &amp; wirte.
        :param str address: rule address, IP OR IP SEG.
        :param int priority: rule priority, range 1 - 100, value less higher priority.
        """
        if access_mode is not None:
            pulumi.set(__self__, "access_mode", access_mode)
        if access_rule_id is not None:
            pulumi.set(__self__, "access_rule_id", access_rule_id)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter(name="accessMode")
    def access_mode(self) -> Optional[int]:
        """
        rule access mode, 1: read only, 2: read &amp; wirte.
        """
        return pulumi.get(self, "access_mode")

    @property
    @pulumi.getter(name="accessRuleId")
    def access_rule_id(self) -> Optional[int]:
        return pulumi.get(self, "access_rule_id")

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        rule address, IP OR IP SEG.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def priority(self) -> Optional[int]:
        """
        rule priority, range 1 - 100, value less higher priority.
        """
        return pulumi.get(self, "priority")


@pulumi.output_type
class LifeCycleRuleLifeCycleRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createTime":
            suggest = "create_time"
        elif key == "lifeCycleRuleId":
            suggest = "life_cycle_rule_id"
        elif key == "lifeCycleRuleName":
            suggest = "life_cycle_rule_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LifeCycleRuleLifeCycleRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LifeCycleRuleLifeCycleRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LifeCycleRuleLifeCycleRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 create_time: Optional[str] = None,
                 life_cycle_rule_id: Optional[int] = None,
                 life_cycle_rule_name: Optional[str] = None,
                 path: Optional[str] = None,
                 status: Optional[int] = None,
                 transitions: Optional[Sequence['outputs.LifeCycleRuleLifeCycleRuleTransition']] = None):
        """
        :param str life_cycle_rule_name: rule name.
        :param str path: rule op path.
        :param int status: rule status, 1:open, 2:close.
        :param Sequence['LifeCycleRuleLifeCycleRuleTransitionArgs'] transitions: life cycle rule transition list.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if life_cycle_rule_id is not None:
            pulumi.set(__self__, "life_cycle_rule_id", life_cycle_rule_id)
        if life_cycle_rule_name is not None:
            pulumi.set(__self__, "life_cycle_rule_name", life_cycle_rule_name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="lifeCycleRuleId")
    def life_cycle_rule_id(self) -> Optional[int]:
        return pulumi.get(self, "life_cycle_rule_id")

    @property
    @pulumi.getter(name="lifeCycleRuleName")
    def life_cycle_rule_name(self) -> Optional[str]:
        """
        rule name.
        """
        return pulumi.get(self, "life_cycle_rule_name")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        rule op path.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def status(self) -> Optional[int]:
        """
        rule status, 1:open, 2:close.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def transitions(self) -> Optional[Sequence['outputs.LifeCycleRuleLifeCycleRuleTransition']]:
        """
        life cycle rule transition list.
        """
        return pulumi.get(self, "transitions")


@pulumi.output_type
class LifeCycleRuleLifeCycleRuleTransition(dict):
    def __init__(__self__, *,
                 days: int,
                 type: int):
        """
        :param int days: trigger days(n day).
        :param int type: transition type, 1: archive, 2: delete, 3: low rate.
        """
        pulumi.set(__self__, "days", days)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def days(self) -> int:
        """
        trigger days(n day).
        """
        return pulumi.get(self, "days")

    @property
    @pulumi.getter
    def type(self) -> int:
        """
        transition type, 1: archive, 2: delete, 3: low rate.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetAccessGroupsAccessGroupResult(dict):
    def __init__(__self__, *,
                 access_group_id: str,
                 access_group_name: str,
                 create_time: str,
                 description: str,
                 vpc_id: str,
                 vpc_type: int):
        """
        :param str access_group_id: access group id.
        :param str access_group_name: access group name.
        :param str create_time: create time.
        :param str description: access group description.
        :param str vpc_id: get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
        :param int vpc_type: vpc network type(1:CVM, 2:BM 1.0).
        """
        pulumi.set(__self__, "access_group_id", access_group_id)
        pulumi.set(__self__, "access_group_name", access_group_name)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_type", vpc_type)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> str:
        """
        access group id.
        """
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="accessGroupName")
    def access_group_name(self) -> str:
        """
        access group name.
        """
        return pulumi.get(self, "access_group_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        access group description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcType")
    def vpc_type(self) -> int:
        """
        vpc network type(1:CVM, 2:BM 1.0).
        """
        return pulumi.get(self, "vpc_type")


@pulumi.output_type
class GetFileSystemsFileSystemResult(dict):
    def __init__(__self__, *,
                 app_id: int,
                 block_size: int,
                 capacity_quota: int,
                 create_time: str,
                 description: str,
                 enable_ranger: bool,
                 file_system_id: str,
                 file_system_name: str,
                 posix_acl: bool,
                 ranger_service_addresses: Sequence[str],
                 region: str,
                 status: int,
                 super_users: Sequence[str]):
        """
        :param int app_id: appid of the user.
        :param int block_size: block size of the file system(byte).
        :param int capacity_quota: capacity of the file system(byte).
        :param str create_time: create time.
        :param str description: desc of the file system.
        :param bool enable_ranger: check the ranger address or not.
        :param str file_system_id: file system id.
        :param str file_system_name: file system name.
        :param bool posix_acl: check POSIX ACL or not.
        :param Sequence[str] ranger_service_addresses: ranger address list.
        :param str region: region of the file system.
        :param int status: status of the file system(1: creating create success 3: create failed).
        :param Sequence[str] super_users: super users of the file system.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "block_size", block_size)
        pulumi.set(__self__, "capacity_quota", capacity_quota)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "enable_ranger", enable_ranger)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "file_system_name", file_system_name)
        pulumi.set(__self__, "posix_acl", posix_acl)
        pulumi.set(__self__, "ranger_service_addresses", ranger_service_addresses)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "super_users", super_users)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> int:
        """
        appid of the user.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="blockSize")
    def block_size(self) -> int:
        """
        block size of the file system(byte).
        """
        return pulumi.get(self, "block_size")

    @property
    @pulumi.getter(name="capacityQuota")
    def capacity_quota(self) -> int:
        """
        capacity of the file system(byte).
        """
        return pulumi.get(self, "capacity_quota")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        desc of the file system.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableRanger")
    def enable_ranger(self) -> bool:
        """
        check the ranger address or not.
        """
        return pulumi.get(self, "enable_ranger")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        file system id.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> str:
        """
        file system name.
        """
        return pulumi.get(self, "file_system_name")

    @property
    @pulumi.getter(name="posixAcl")
    def posix_acl(self) -> bool:
        """
        check POSIX ACL or not.
        """
        return pulumi.get(self, "posix_acl")

    @property
    @pulumi.getter(name="rangerServiceAddresses")
    def ranger_service_addresses(self) -> Sequence[str]:
        """
        ranger address list.
        """
        return pulumi.get(self, "ranger_service_addresses")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        region of the file system.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        status of the file system(1: creating create success 3: create failed).
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="superUsers")
    def super_users(self) -> Sequence[str]:
        """
        super users of the file system.
        """
        return pulumi.get(self, "super_users")


@pulumi.output_type
class GetMountPointsMountPointResult(dict):
    def __init__(__self__, *,
                 access_group_ids: Sequence[str],
                 create_time: str,
                 file_system_id: str,
                 mount_point_id: str,
                 mount_point_name: str,
                 status: int):
        """
        :param Sequence[str] access_group_ids: associated group ids.
        :param str create_time: create time.
        :param str file_system_id: get mount points belongs to file system id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
        :param str mount_point_id: mount point id.
        :param str mount_point_name: mount point name.
        :param int status: mount point status.
        """
        pulumi.set(__self__, "access_group_ids", access_group_ids)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "mount_point_id", mount_point_id)
        pulumi.set(__self__, "mount_point_name", mount_point_name)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accessGroupIds")
    def access_group_ids(self) -> Sequence[str]:
        """
        associated group ids.
        """
        return pulumi.get(self, "access_group_ids")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        get mount points belongs to file system id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="mountPointId")
    def mount_point_id(self) -> str:
        """
        mount point id.
        """
        return pulumi.get(self, "mount_point_id")

    @property
    @pulumi.getter(name="mountPointName")
    def mount_point_name(self) -> str:
        """
        mount point name.
        """
        return pulumi.get(self, "mount_point_name")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        mount point status.
        """
        return pulumi.get(self, "status")


