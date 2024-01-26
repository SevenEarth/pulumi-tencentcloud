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
    'NatInstanceNewModeItems',
    'VpcInstanceVpcFwInstance',
    'VpcInstanceVpcFwInstanceFwDeploy',
    'VpcPolicyBetaList',
    'GetEdgeFwSwitchesDataResult',
    'GetNatFwSwitchesDataResult',
    'GetVpcFwSwitchesSwitchListResult',
]

@pulumi.output_type
class NatInstanceNewModeItems(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "vpcLists":
            suggest = "vpc_lists"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NatInstanceNewModeItems. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NatInstanceNewModeItems.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NatInstanceNewModeItems.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 eips: Sequence[str],
                 vpc_lists: Sequence[str]):
        """
        :param Sequence[str] eips: List of egress elastic public network IPs bound in the new mode.
        :param Sequence[str] vpc_lists: List of vpcs connected in new mode.
        """
        pulumi.set(__self__, "eips", eips)
        pulumi.set(__self__, "vpc_lists", vpc_lists)

    @property
    @pulumi.getter
    def eips(self) -> Sequence[str]:
        """
        List of egress elastic public network IPs bound in the new mode.
        """
        return pulumi.get(self, "eips")

    @property
    @pulumi.getter(name="vpcLists")
    def vpc_lists(self) -> Sequence[str]:
        """
        List of vpcs connected in new mode.
        """
        return pulumi.get(self, "vpc_lists")


@pulumi.output_type
class VpcInstanceVpcFwInstance(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fwDeploy":
            suggest = "fw_deploy"
        elif key == "fwInsId":
            suggest = "fw_ins_id"
        elif key == "vpcIds":
            suggest = "vpc_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcInstanceVpcFwInstance. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcInstanceVpcFwInstance.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcInstanceVpcFwInstance.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 fw_deploy: 'outputs.VpcInstanceVpcFwInstanceFwDeploy',
                 name: str,
                 fw_ins_id: Optional[str] = None,
                 vpc_ids: Optional[Sequence[str]] = None):
        """
        :param 'VpcInstanceVpcFwInstanceFwDeployArgs' fw_deploy: Deploy regional information.
        :param str name: Firewall instance name.
        :param Sequence[str] vpc_ids: List of VpcIds accessed in private network mode; only used in private network mode.
        """
        pulumi.set(__self__, "fw_deploy", fw_deploy)
        pulumi.set(__self__, "name", name)
        if fw_ins_id is not None:
            pulumi.set(__self__, "fw_ins_id", fw_ins_id)
        if vpc_ids is not None:
            pulumi.set(__self__, "vpc_ids", vpc_ids)

    @property
    @pulumi.getter(name="fwDeploy")
    def fw_deploy(self) -> 'outputs.VpcInstanceVpcFwInstanceFwDeploy':
        """
        Deploy regional information.
        """
        return pulumi.get(self, "fw_deploy")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Firewall instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="fwInsId")
    def fw_ins_id(self) -> Optional[str]:
        return pulumi.get(self, "fw_ins_id")

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> Optional[Sequence[str]]:
        """
        List of VpcIds accessed in private network mode; only used in private network mode.
        """
        return pulumi.get(self, "vpc_ids")


@pulumi.output_type
class VpcInstanceVpcFwInstanceFwDeploy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deployRegion":
            suggest = "deploy_region"
        elif key == "zoneSets":
            suggest = "zone_sets"
        elif key == "crossAZone":
            suggest = "cross_a_zone"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcInstanceVpcFwInstanceFwDeploy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcInstanceVpcFwInstanceFwDeploy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcInstanceVpcFwInstanceFwDeploy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 deploy_region: str,
                 width: int,
                 zone_sets: Sequence[str],
                 cross_a_zone: Optional[int] = None):
        """
        :param str deploy_region: Firewall Deployment Region.
        :param int width: Bandwidth, unit: Mbps.
        :param Sequence[str] zone_sets: Zone list.
        :param int cross_a_zone: Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.
        """
        pulumi.set(__self__, "deploy_region", deploy_region)
        pulumi.set(__self__, "width", width)
        pulumi.set(__self__, "zone_sets", zone_sets)
        if cross_a_zone is not None:
            pulumi.set(__self__, "cross_a_zone", cross_a_zone)

    @property
    @pulumi.getter(name="deployRegion")
    def deploy_region(self) -> str:
        """
        Firewall Deployment Region.
        """
        return pulumi.get(self, "deploy_region")

    @property
    @pulumi.getter
    def width(self) -> int:
        """
        Bandwidth, unit: Mbps.
        """
        return pulumi.get(self, "width")

    @property
    @pulumi.getter(name="zoneSets")
    def zone_sets(self) -> Sequence[str]:
        """
        Zone list.
        """
        return pulumi.get(self, "zone_sets")

    @property
    @pulumi.getter(name="crossAZone")
    def cross_a_zone(self) -> Optional[int]:
        """
        Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.
        """
        return pulumi.get(self, "cross_a_zone")


@pulumi.output_type
class VpcPolicyBetaList(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lastTime":
            suggest = "last_time"
        elif key == "taskId":
            suggest = "task_id"
        elif key == "taskName":
            suggest = "task_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcPolicyBetaList. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcPolicyBetaList.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcPolicyBetaList.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 last_time: Optional[str] = None,
                 task_id: Optional[int] = None,
                 task_name: Optional[str] = None):
        if last_time is not None:
            pulumi.set(__self__, "last_time", last_time)
        if task_id is not None:
            pulumi.set(__self__, "task_id", task_id)
        if task_name is not None:
            pulumi.set(__self__, "task_name", task_name)

    @property
    @pulumi.getter(name="lastTime")
    def last_time(self) -> Optional[str]:
        return pulumi.get(self, "last_time")

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> Optional[int]:
        return pulumi.get(self, "task_id")

    @property
    @pulumi.getter(name="taskName")
    def task_name(self) -> Optional[str]:
        return pulumi.get(self, "task_name")


@pulumi.output_type
class GetEdgeFwSwitchesDataResult(dict):
    def __init__(__self__, *,
                 asset_type: str,
                 instance_id: str,
                 instance_name: str,
                 intranet_ip: str,
                 public_ip: str,
                 public_ip_type: int,
                 region: str,
                 status: int,
                 switch_mode: int,
                 vpc_id: str):
        """
        :param str asset_type: Asset Type.
        :param str instance_id: Instance Id.
        :param str instance_name: Instance Name.
        :param str intranet_ip: Intranet Ip.
        :param str public_ip: public ip.
        :param int public_ip_type: Public IP type.
        :param str region: region.
        :param int status: status.
        :param int switch_mode: switch mode.
        :param str vpc_id: vpc id.
        """
        pulumi.set(__self__, "asset_type", asset_type)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "intranet_ip", intranet_ip)
        pulumi.set(__self__, "public_ip", public_ip)
        pulumi.set(__self__, "public_ip_type", public_ip_type)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "switch_mode", switch_mode)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="assetType")
    def asset_type(self) -> str:
        """
        Asset Type.
        """
        return pulumi.get(self, "asset_type")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        Instance Name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="intranetIp")
    def intranet_ip(self) -> str:
        """
        Intranet Ip.
        """
        return pulumi.get(self, "intranet_ip")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        public ip.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="publicIpType")
    def public_ip_type(self) -> int:
        """
        Public IP type.
        """
        return pulumi.get(self, "public_ip_type")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="switchMode")
    def switch_mode(self) -> int:
        """
        switch mode.
        """
        return pulumi.get(self, "switch_mode")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        vpc id.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetNatFwSwitchesDataResult(dict):
    def __init__(__self__, *,
                 abnormal: int,
                 cvm_num: int,
                 enable: int,
                 id: int,
                 nat_id: str,
                 nat_ins_id: str,
                 nat_ins_name: str,
                 nat_name: str,
                 region: str,
                 route_id: str,
                 route_name: str,
                 status: int,
                 subnet_cidr: str,
                 subnet_id: str,
                 subnet_name: str,
                 vpc_id: str,
                 vpc_name: str):
        """
        :param int abnormal: Whether the switch is abnormal, 0: normal, 1: abnormal.
        :param int cvm_num: Cvm Num.
        :param int enable: Effective status.
        :param int id: ID.
        :param str nat_id: NAT gatway Id.
        :param str nat_ins_id: Filter the NAT firewall instance to which the NAT firewall subnet switch belongs.
        :param str nat_ins_name: NAT firewall instance name.
        :param str nat_name: NAT gatway name.
        :param str region: Region.
        :param str route_id: Route Id.
        :param str route_name: Route Name.
        :param int status: Switch status, 1 open; 0 close.
        :param str subnet_cidr: IPv4 CIDR.
        :param str subnet_id: Subnet Id.
        :param str subnet_name: Subnet Name.
        :param str vpc_id: Vpc Id.
        :param str vpc_name: Vpc Name.
        """
        pulumi.set(__self__, "abnormal", abnormal)
        pulumi.set(__self__, "cvm_num", cvm_num)
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "nat_id", nat_id)
        pulumi.set(__self__, "nat_ins_id", nat_ins_id)
        pulumi.set(__self__, "nat_ins_name", nat_ins_name)
        pulumi.set(__self__, "nat_name", nat_name)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "route_id", route_id)
        pulumi.set(__self__, "route_name", route_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet_cidr", subnet_cidr)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "subnet_name", subnet_name)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_name", vpc_name)

    @property
    @pulumi.getter
    def abnormal(self) -> int:
        """
        Whether the switch is abnormal, 0: normal, 1: abnormal.
        """
        return pulumi.get(self, "abnormal")

    @property
    @pulumi.getter(name="cvmNum")
    def cvm_num(self) -> int:
        """
        Cvm Num.
        """
        return pulumi.get(self, "cvm_num")

    @property
    @pulumi.getter
    def enable(self) -> int:
        """
        Effective status.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="natId")
    def nat_id(self) -> str:
        """
        NAT gatway Id.
        """
        return pulumi.get(self, "nat_id")

    @property
    @pulumi.getter(name="natInsId")
    def nat_ins_id(self) -> str:
        """
        Filter the NAT firewall instance to which the NAT firewall subnet switch belongs.
        """
        return pulumi.get(self, "nat_ins_id")

    @property
    @pulumi.getter(name="natInsName")
    def nat_ins_name(self) -> str:
        """
        NAT firewall instance name.
        """
        return pulumi.get(self, "nat_ins_name")

    @property
    @pulumi.getter(name="natName")
    def nat_name(self) -> str:
        """
        NAT gatway name.
        """
        return pulumi.get(self, "nat_name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> str:
        """
        Route Id.
        """
        return pulumi.get(self, "route_id")

    @property
    @pulumi.getter(name="routeName")
    def route_name(self) -> str:
        """
        Route Name.
        """
        return pulumi.get(self, "route_name")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        Switch status, 1 open; 0 close.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetCidr")
    def subnet_cidr(self) -> str:
        """
        IPv4 CIDR.
        """
        return pulumi.get(self, "subnet_cidr")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        Subnet Id.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="subnetName")
    def subnet_name(self) -> str:
        """
        Subnet Name.
        """
        return pulumi.get(self, "subnet_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        Vpc Id.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> str:
        """
        Vpc Name.
        """
        return pulumi.get(self, "vpc_name")


@pulumi.output_type
class GetVpcFwSwitchesSwitchListResult(dict):
    def __init__(__self__, *,
                 enable: int,
                 status: int,
                 switch_id: str,
                 switch_mode: int,
                 switch_name: str):
        """
        :param int enable: Switch status 0: off, 1: on.
        :param int status: Switch status 0: normal, 1: switching.
        :param str switch_id: Firewall switch ID.
        :param int switch_mode: switch mode.
        :param str switch_name: Firewall switch name.
        """
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "switch_id", switch_id)
        pulumi.set(__self__, "switch_mode", switch_mode)
        pulumi.set(__self__, "switch_name", switch_name)

    @property
    @pulumi.getter
    def enable(self) -> int:
        """
        Switch status 0: off, 1: on.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        Switch status 0: normal, 1: switching.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="switchId")
    def switch_id(self) -> str:
        """
        Firewall switch ID.
        """
        return pulumi.get(self, "switch_id")

    @property
    @pulumi.getter(name="switchMode")
    def switch_mode(self) -> int:
        """
        switch mode.
        """
        return pulumi.get(self, "switch_mode")

    @property
    @pulumi.getter(name="switchName")
    def switch_name(self) -> str:
        """
        Firewall switch name.
        """
        return pulumi.get(self, "switch_name")


