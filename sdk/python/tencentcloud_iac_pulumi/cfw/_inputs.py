# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'NatInstanceNewModeItemsArgs',
    'VpcInstanceVpcFwInstanceArgs',
    'VpcInstanceVpcFwInstanceFwDeployArgs',
    'VpcPolicyBetaListArgs',
]

@pulumi.input_type
class NatInstanceNewModeItemsArgs:
    def __init__(__self__, *,
                 eips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vpc_lists: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] eips: List of egress elastic public network IPs bound in the new mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_lists: List of vpcs connected in new mode.
        """
        pulumi.set(__self__, "eips", eips)
        pulumi.set(__self__, "vpc_lists", vpc_lists)

    @property
    @pulumi.getter
    def eips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of egress elastic public network IPs bound in the new mode.
        """
        return pulumi.get(self, "eips")

    @eips.setter
    def eips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "eips", value)

    @property
    @pulumi.getter(name="vpcLists")
    def vpc_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of vpcs connected in new mode.
        """
        return pulumi.get(self, "vpc_lists")

    @vpc_lists.setter
    def vpc_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vpc_lists", value)


@pulumi.input_type
class VpcInstanceVpcFwInstanceArgs:
    def __init__(__self__, *,
                 fw_deploy: pulumi.Input['VpcInstanceVpcFwInstanceFwDeployArgs'],
                 name: pulumi.Input[str],
                 fw_ins_id: Optional[pulumi.Input[str]] = None,
                 vpc_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input['VpcInstanceVpcFwInstanceFwDeployArgs'] fw_deploy: Deploy regional information.
        :param pulumi.Input[str] name: Firewall instance name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_ids: List of VpcIds accessed in private network mode; only used in private network mode.
        """
        pulumi.set(__self__, "fw_deploy", fw_deploy)
        pulumi.set(__self__, "name", name)
        if fw_ins_id is not None:
            pulumi.set(__self__, "fw_ins_id", fw_ins_id)
        if vpc_ids is not None:
            pulumi.set(__self__, "vpc_ids", vpc_ids)

    @property
    @pulumi.getter(name="fwDeploy")
    def fw_deploy(self) -> pulumi.Input['VpcInstanceVpcFwInstanceFwDeployArgs']:
        """
        Deploy regional information.
        """
        return pulumi.get(self, "fw_deploy")

    @fw_deploy.setter
    def fw_deploy(self, value: pulumi.Input['VpcInstanceVpcFwInstanceFwDeployArgs']):
        pulumi.set(self, "fw_deploy", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Firewall instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="fwInsId")
    def fw_ins_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fw_ins_id")

    @fw_ins_id.setter
    def fw_ins_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fw_ins_id", value)

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of VpcIds accessed in private network mode; only used in private network mode.
        """
        return pulumi.get(self, "vpc_ids")

    @vpc_ids.setter
    def vpc_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_ids", value)


@pulumi.input_type
class VpcInstanceVpcFwInstanceFwDeployArgs:
    def __init__(__self__, *,
                 deploy_region: pulumi.Input[str],
                 width: pulumi.Input[int],
                 zone_sets: pulumi.Input[Sequence[pulumi.Input[str]]],
                 cross_a_zone: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] deploy_region: Firewall Deployment Region.
        :param pulumi.Input[int] width: Bandwidth, unit: Mbps.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_sets: Zone list.
        :param pulumi.Input[int] cross_a_zone: Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.
        """
        pulumi.set(__self__, "deploy_region", deploy_region)
        pulumi.set(__self__, "width", width)
        pulumi.set(__self__, "zone_sets", zone_sets)
        if cross_a_zone is not None:
            pulumi.set(__self__, "cross_a_zone", cross_a_zone)

    @property
    @pulumi.getter(name="deployRegion")
    def deploy_region(self) -> pulumi.Input[str]:
        """
        Firewall Deployment Region.
        """
        return pulumi.get(self, "deploy_region")

    @deploy_region.setter
    def deploy_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "deploy_region", value)

    @property
    @pulumi.getter
    def width(self) -> pulumi.Input[int]:
        """
        Bandwidth, unit: Mbps.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: pulumi.Input[int]):
        pulumi.set(self, "width", value)

    @property
    @pulumi.getter(name="zoneSets")
    def zone_sets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Zone list.
        """
        return pulumi.get(self, "zone_sets")

    @zone_sets.setter
    def zone_sets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "zone_sets", value)

    @property
    @pulumi.getter(name="crossAZone")
    def cross_a_zone(self) -> Optional[pulumi.Input[int]]:
        """
        Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.
        """
        return pulumi.get(self, "cross_a_zone")

    @cross_a_zone.setter
    def cross_a_zone(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cross_a_zone", value)


@pulumi.input_type
class VpcPolicyBetaListArgs:
    def __init__(__self__, *,
                 last_time: Optional[pulumi.Input[str]] = None,
                 task_id: Optional[pulumi.Input[int]] = None,
                 task_name: Optional[pulumi.Input[str]] = None):
        if last_time is not None:
            pulumi.set(__self__, "last_time", last_time)
        if task_id is not None:
            pulumi.set(__self__, "task_id", task_id)
        if task_name is not None:
            pulumi.set(__self__, "task_name", task_name)

    @property
    @pulumi.getter(name="lastTime")
    def last_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "last_time")

    @last_time.setter
    def last_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_time", value)

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "task_id")

    @task_id.setter
    def task_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "task_id", value)

    @property
    @pulumi.getter(name="taskName")
    def task_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "task_name")

    @task_name.setter
    def task_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_name", value)


