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
    'GetAccessGroupsAccessGroupListResult',
    'GetAccessRulesAccessRuleListResult',
    'GetAvailableZoneRegionZoneResult',
    'GetAvailableZoneRegionZoneZoneResult',
    'GetAvailableZoneRegionZoneZoneTypeResult',
    'GetAvailableZoneRegionZoneZoneTypeProtocolResult',
    'GetFileSystemClientsClientListResult',
    'GetFileSystemsFileSystemListResult',
    'GetMountTargetsMountTargetResult',
]

@pulumi.output_type
class GetAccessGroupsAccessGroupListResult(dict):
    def __init__(__self__, *,
                 access_group_id: str,
                 create_time: str,
                 description: str,
                 name: str):
        """
        :param str access_group_id: A specified access group ID used to query.
        :param str create_time: Creation time of the access group.
        :param str description: Description of the access group.
        :param str name: A access group Name used to query.
        """
        pulumi.set(__self__, "access_group_id", access_group_id)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> str:
        """
        A specified access group ID used to query.
        """
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of the access group.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the access group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A access group Name used to query.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetAccessRulesAccessRuleListResult(dict):
    def __init__(__self__, *,
                 access_rule_id: str,
                 auth_client_ip: str,
                 priority: int,
                 rw_permission: str,
                 user_permission: str):
        """
        :param str access_rule_id: A specified access rule ID used to query.
        :param str auth_client_ip: Allowed IP of the access rule.
        :param int priority: The priority level of access rule.
        :param str rw_permission: Read and write permissions.
        :param str user_permission: The permissions of accessing users.
        """
        pulumi.set(__self__, "access_rule_id", access_rule_id)
        pulumi.set(__self__, "auth_client_ip", auth_client_ip)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "rw_permission", rw_permission)
        pulumi.set(__self__, "user_permission", user_permission)

    @property
    @pulumi.getter(name="accessRuleId")
    def access_rule_id(self) -> str:
        """
        A specified access rule ID used to query.
        """
        return pulumi.get(self, "access_rule_id")

    @property
    @pulumi.getter(name="authClientIp")
    def auth_client_ip(self) -> str:
        """
        Allowed IP of the access rule.
        """
        return pulumi.get(self, "auth_client_ip")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority level of access rule.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="rwPermission")
    def rw_permission(self) -> str:
        """
        Read and write permissions.
        """
        return pulumi.get(self, "rw_permission")

    @property
    @pulumi.getter(name="userPermission")
    def user_permission(self) -> str:
        """
        The permissions of accessing users.
        """
        return pulumi.get(self, "user_permission")


@pulumi.output_type
class GetAvailableZoneRegionZoneResult(dict):
    def __init__(__self__, *,
                 region: str,
                 region_cn_name: str,
                 region_name: str,
                 region_status: str,
                 zones: Sequence['outputs.GetAvailableZoneRegionZoneZoneResult']):
        """
        :param str region: Region name, such as `ap-beijing`.
        :param str region_cn_name: Region chinese name, such as `Guangzhou`.
        :param str region_name: Region name, such as `bj`.
        :param str region_status: Region availability. If a region has at least one AZ where resources are purchasable, this value will be AVAILABLE; otherwise, it will be UNAVAILABLE.
        :param Sequence['GetAvailableZoneRegionZoneZoneArgs'] zones: Array of AZs.
        """
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "region_cn_name", region_cn_name)
        pulumi.set(__self__, "region_name", region_name)
        pulumi.set(__self__, "region_status", region_status)
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region name, such as `ap-beijing`.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionCnName")
    def region_cn_name(self) -> str:
        """
        Region chinese name, such as `Guangzhou`.
        """
        return pulumi.get(self, "region_cn_name")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        """
        Region name, such as `bj`.
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="regionStatus")
    def region_status(self) -> str:
        """
        Region availability. If a region has at least one AZ where resources are purchasable, this value will be AVAILABLE; otherwise, it will be UNAVAILABLE.
        """
        return pulumi.get(self, "region_status")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.GetAvailableZoneRegionZoneZoneResult']:
        """
        Array of AZs.
        """
        return pulumi.get(self, "zones")


@pulumi.output_type
class GetAvailableZoneRegionZoneZoneResult(dict):
    def __init__(__self__, *,
                 types: Sequence['outputs.GetAvailableZoneRegionZoneZoneTypeResult'],
                 zone: str,
                 zone_cn_name: str,
                 zone_id: int,
                 zone_name: str):
        """
        :param Sequence['GetAvailableZoneRegionZoneZoneTypeArgs'] types: Array of classes.
        :param str zone: AZ name.
        :param str zone_cn_name: Chinese name of an AZ.
        :param int zone_id: AZ ID.
        :param str zone_name: Chinese and English names of an AZ.
        """
        pulumi.set(__self__, "types", types)
        pulumi.set(__self__, "zone", zone)
        pulumi.set(__self__, "zone_cn_name", zone_cn_name)
        pulumi.set(__self__, "zone_id", zone_id)
        pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter
    def types(self) -> Sequence['outputs.GetAvailableZoneRegionZoneZoneTypeResult']:
        """
        Array of classes.
        """
        return pulumi.get(self, "types")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        AZ name.
        """
        return pulumi.get(self, "zone")

    @property
    @pulumi.getter(name="zoneCnName")
    def zone_cn_name(self) -> str:
        """
        Chinese name of an AZ.
        """
        return pulumi.get(self, "zone_cn_name")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> int:
        """
        AZ ID.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> str:
        """
        Chinese and English names of an AZ.
        """
        return pulumi.get(self, "zone_name")


@pulumi.output_type
class GetAvailableZoneRegionZoneZoneTypeResult(dict):
    def __init__(__self__, *,
                 prepayment: bool,
                 protocols: Sequence['outputs.GetAvailableZoneRegionZoneZoneTypeProtocolResult'],
                 type: str):
        """
        :param bool prepayment: Indicates whether prepaid is supported. true: yes; false: no.
        :param Sequence['GetAvailableZoneRegionZoneZoneTypeProtocolArgs'] protocols: Protocol and sale details.
        :param str type: Storage class. Valid values: SD (standard storage) and HP (high-performance storage).
        """
        pulumi.set(__self__, "prepayment", prepayment)
        pulumi.set(__self__, "protocols", protocols)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def prepayment(self) -> bool:
        """
        Indicates whether prepaid is supported. true: yes; false: no.
        """
        return pulumi.get(self, "prepayment")

    @property
    @pulumi.getter
    def protocols(self) -> Sequence['outputs.GetAvailableZoneRegionZoneZoneTypeProtocolResult']:
        """
        Protocol and sale details.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Storage class. Valid values: SD (standard storage) and HP (high-performance storage).
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetAvailableZoneRegionZoneZoneTypeProtocolResult(dict):
    def __init__(__self__, *,
                 protocol: str,
                 sale_status: str):
        """
        :param str protocol: Protocol type. Valid values: NFS, CIFS.
        :param str sale_status: Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable).
        """
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "sale_status", sale_status)

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Protocol type. Valid values: NFS, CIFS.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="saleStatus")
    def sale_status(self) -> str:
        """
        Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable).
        """
        return pulumi.get(self, "sale_status")


@pulumi.output_type
class GetFileSystemClientsClientListResult(dict):
    def __init__(__self__, *,
                 cfs_vip: str,
                 client_ip: str,
                 mount_directory: str,
                 vpc_id: str,
                 zone: str,
                 zone_name: str):
        """
        :param str cfs_vip: IP address of the file system.
        :param str client_ip: Client IP.
        :param str mount_directory: Path in which the file system is mounted to the client.
        :param str vpc_id: File system VPCID.
        :param str zone: Name of the availability zone, e.g. ap-beijing-1. For more information, see regions and availability zones in the Overview document.
        :param str zone_name: AZ name.
        """
        pulumi.set(__self__, "cfs_vip", cfs_vip)
        pulumi.set(__self__, "client_ip", client_ip)
        pulumi.set(__self__, "mount_directory", mount_directory)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "zone", zone)
        pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter(name="cfsVip")
    def cfs_vip(self) -> str:
        """
        IP address of the file system.
        """
        return pulumi.get(self, "cfs_vip")

    @property
    @pulumi.getter(name="clientIp")
    def client_ip(self) -> str:
        """
        Client IP.
        """
        return pulumi.get(self, "client_ip")

    @property
    @pulumi.getter(name="mountDirectory")
    def mount_directory(self) -> str:
        """
        Path in which the file system is mounted to the client.
        """
        return pulumi.get(self, "mount_directory")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        File system VPCID.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        Name of the availability zone, e.g. ap-beijing-1. For more information, see regions and availability zones in the Overview document.
        """
        return pulumi.get(self, "zone")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> str:
        """
        AZ name.
        """
        return pulumi.get(self, "zone_name")


@pulumi.output_type
class GetFileSystemsFileSystemListResult(dict):
    def __init__(__self__, *,
                 access_group_id: str,
                 availability_zone: str,
                 create_time: str,
                 file_system_id: str,
                 fs_id: str,
                 mount_ip: str,
                 name: str,
                 protocol: str,
                 size_limit: int,
                 size_used: int,
                 status: str,
                 storage_type: str):
        """
        :param str access_group_id: ID of the access group.
        :param str availability_zone: The available zone that the file system locates at.
        :param str create_time: Creation time of the file system.
        :param str file_system_id: A specified file system ID used to query.
        :param str fs_id: Mount root-directory.
        :param str mount_ip: IP of the file system.
        :param str name: A file system name used to query.
        :param str protocol: Protocol of the file system.
        :param int size_limit: Size limit of the file system.
        :param int size_used: Size used of the file system.
        :param str status: Status of the file system.
        :param str storage_type: Storage type of the file system.
        """
        pulumi.set(__self__, "access_group_id", access_group_id)
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "fs_id", fs_id)
        pulumi.set(__self__, "mount_ip", mount_ip)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "size_limit", size_limit)
        pulumi.set(__self__, "size_used", size_used)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage_type", storage_type)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> str:
        """
        ID of the access group.
        """
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The available zone that the file system locates at.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of the file system.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        A specified file system ID used to query.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="fsId")
    def fs_id(self) -> str:
        """
        Mount root-directory.
        """
        return pulumi.get(self, "fs_id")

    @property
    @pulumi.getter(name="mountIp")
    def mount_ip(self) -> str:
        """
        IP of the file system.
        """
        return pulumi.get(self, "mount_ip")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A file system name used to query.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Protocol of the file system.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sizeLimit")
    def size_limit(self) -> int:
        """
        Size limit of the file system.
        """
        return pulumi.get(self, "size_limit")

    @property
    @pulumi.getter(name="sizeUsed")
    def size_used(self) -> int:
        """
        Size used of the file system.
        """
        return pulumi.get(self, "size_used")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the file system.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> str:
        """
        Storage type of the file system.
        """
        return pulumi.get(self, "storage_type")


@pulumi.output_type
class GetMountTargetsMountTargetResult(dict):
    def __init__(__self__, *,
                 ccn_id: str,
                 cidr_block: str,
                 file_system_id: str,
                 fs_id: str,
                 ip_address: str,
                 life_cycle_state: str,
                 mount_target_id: str,
                 network_interface: str,
                 subnet_id: str,
                 subnet_name: str,
                 vpc_id: str,
                 vpc_name: str):
        """
        :param str ccn_id: CCN instance ID used by CFS Turbo.
        :param str cidr_block: CCN IP range used by CFS Turbo.
        :param str file_system_id: File system ID.
        :param str fs_id: Mount root-directory.
        :param str ip_address: Mount target IP.
        :param str life_cycle_state: Mount target status.
        :param str mount_target_id: Mount target ID.
        :param str network_interface: Network type.
        :param str subnet_id: Subnet ID.
        :param str subnet_name: Subnet name.
        :param str vpc_id: VPC ID.
        :param str vpc_name: VPC name.
        """
        pulumi.set(__self__, "ccn_id", ccn_id)
        pulumi.set(__self__, "cidr_block", cidr_block)
        pulumi.set(__self__, "file_system_id", file_system_id)
        pulumi.set(__self__, "fs_id", fs_id)
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "life_cycle_state", life_cycle_state)
        pulumi.set(__self__, "mount_target_id", mount_target_id)
        pulumi.set(__self__, "network_interface", network_interface)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "subnet_name", subnet_name)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_name", vpc_name)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> str:
        """
        CCN instance ID used by CFS Turbo.
        """
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> str:
        """
        CCN IP range used by CFS Turbo.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        File system ID.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="fsId")
    def fs_id(self) -> str:
        """
        Mount root-directory.
        """
        return pulumi.get(self, "fs_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        Mount target IP.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="lifeCycleState")
    def life_cycle_state(self) -> str:
        """
        Mount target status.
        """
        return pulumi.get(self, "life_cycle_state")

    @property
    @pulumi.getter(name="mountTargetId")
    def mount_target_id(self) -> str:
        """
        Mount target ID.
        """
        return pulumi.get(self, "mount_target_id")

    @property
    @pulumi.getter(name="networkInterface")
    def network_interface(self) -> str:
        """
        Network type.
        """
        return pulumi.get(self, "network_interface")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        Subnet ID.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="subnetName")
    def subnet_name(self) -> str:
        """
        Subnet name.
        """
        return pulumi.get(self, "subnet_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> str:
        """
        VPC name.
        """
        return pulumi.get(self, "vpc_name")


