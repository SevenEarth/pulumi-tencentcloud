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
    'GetClustersListResult',
    'GetClustersListEndpointSetResult',
    'GetInstancesListResult',
]

@pulumi.output_type
class GetClustersListResult(dict):
    def __init__(__self__, *,
                 auto_renew_flag: int,
                 cluster_id: str,
                 cluster_name: str,
                 create_time: str,
                 db_charset: str,
                 db_kernel_version: str,
                 db_major_version: str,
                 db_version: str,
                 endpoint_sets: Sequence['outputs.GetClustersListEndpointSetResult'],
                 instance_count: int,
                 pay_mode: str,
                 pay_period_end_time: str,
                 project_id: int,
                 region: str,
                 status: str,
                 status_desc: str,
                 storage_limit: int,
                 storage_pay_mode: str,
                 storage_used: float,
                 zone: str):
        """
        :param int auto_renew_flag: auto renew flag.
        :param str cluster_id: cluster id.
        :param str cluster_name: cluster name.
        :param str create_time: create time.
        :param str db_charset: db charset.
        :param str db_kernel_version: db kernel version.
        :param str db_major_version: db major version.
        :param str db_version: db version.
        :param Sequence['GetClustersListEndpointSetArgs'] endpoint_sets: endpoint set.
        :param int instance_count: instance count.
        :param str pay_mode: pay mode.
        :param str pay_period_end_time: pay period expired time.
        :param int project_id: project id, default to 0, means default project.
        :param str region: region.
        :param str status: cluster status.
        :param str status_desc: status description.
        :param int storage_limit: storage limit, unit is GB.
        :param str storage_pay_mode: storage pay mode, optional value is PREPAID or POSTPAID_BY_HOUR.
        :param float storage_used: storage used, unit is GB.
        :param str zone: zone.
        """
        pulumi.set(__self__, "auto_renew_flag", auto_renew_flag)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "db_charset", db_charset)
        pulumi.set(__self__, "db_kernel_version", db_kernel_version)
        pulumi.set(__self__, "db_major_version", db_major_version)
        pulumi.set(__self__, "db_version", db_version)
        pulumi.set(__self__, "endpoint_sets", endpoint_sets)
        pulumi.set(__self__, "instance_count", instance_count)
        pulumi.set(__self__, "pay_mode", pay_mode)
        pulumi.set(__self__, "pay_period_end_time", pay_period_end_time)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "status_desc", status_desc)
        pulumi.set(__self__, "storage_limit", storage_limit)
        pulumi.set(__self__, "storage_pay_mode", storage_pay_mode)
        pulumi.set(__self__, "storage_used", storage_used)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> int:
        """
        auto renew flag.
        """
        return pulumi.get(self, "auto_renew_flag")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        cluster id.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dbCharset")
    def db_charset(self) -> str:
        """
        db charset.
        """
        return pulumi.get(self, "db_charset")

    @property
    @pulumi.getter(name="dbKernelVersion")
    def db_kernel_version(self) -> str:
        """
        db kernel version.
        """
        return pulumi.get(self, "db_kernel_version")

    @property
    @pulumi.getter(name="dbMajorVersion")
    def db_major_version(self) -> str:
        """
        db major version.
        """
        return pulumi.get(self, "db_major_version")

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> str:
        """
        db version.
        """
        return pulumi.get(self, "db_version")

    @property
    @pulumi.getter(name="endpointSets")
    def endpoint_sets(self) -> Sequence['outputs.GetClustersListEndpointSetResult']:
        """
        endpoint set.
        """
        return pulumi.get(self, "endpoint_sets")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> int:
        """
        instance count.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="payMode")
    def pay_mode(self) -> str:
        """
        pay mode.
        """
        return pulumi.get(self, "pay_mode")

    @property
    @pulumi.getter(name="payPeriodEndTime")
    def pay_period_end_time(self) -> str:
        """
        pay period expired time.
        """
        return pulumi.get(self, "pay_period_end_time")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        """
        project id, default to 0, means default project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        cluster status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDesc")
    def status_desc(self) -> str:
        """
        status description.
        """
        return pulumi.get(self, "status_desc")

    @property
    @pulumi.getter(name="storageLimit")
    def storage_limit(self) -> int:
        """
        storage limit, unit is GB.
        """
        return pulumi.get(self, "storage_limit")

    @property
    @pulumi.getter(name="storagePayMode")
    def storage_pay_mode(self) -> str:
        """
        storage pay mode, optional value is PREPAID or POSTPAID_BY_HOUR.
        """
        return pulumi.get(self, "storage_pay_mode")

    @property
    @pulumi.getter(name="storageUsed")
    def storage_used(self) -> float:
        """
        storage used, unit is GB.
        """
        return pulumi.get(self, "storage_used")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        zone.
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class GetClustersListEndpointSetResult(dict):
    def __init__(__self__, *,
                 cluster_id: str,
                 endpoint_id: str,
                 endpoint_name: str,
                 endpoint_type: str,
                 private_ip: str,
                 private_port: int,
                 subnet_id: str,
                 vpc_id: str,
                 wan_domain: str,
                 wan_ip: str,
                 wan_port: int):
        """
        :param str cluster_id: cluster id.
        :param str endpoint_id: endpoint id.
        :param str endpoint_name: endpoint name.
        :param str endpoint_type: endpoint type.
        :param str private_ip: private ip.
        :param int private_port: private port.
        :param str subnet_id: subnet id.
        :param str vpc_id: vpc id.
        :param str wan_domain: wan domain.
        :param str wan_ip: wan ip.
        :param int wan_port: wan port.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "endpoint_name", endpoint_name)
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "private_port", private_port)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "wan_domain", wan_domain)
        pulumi.set(__self__, "wan_ip", wan_ip)
        pulumi.set(__self__, "wan_port", wan_port)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        cluster id.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> str:
        """
        endpoint id.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> str:
        """
        endpoint name.
        """
        return pulumi.get(self, "endpoint_name")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> str:
        """
        endpoint type.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        private ip.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="privatePort")
    def private_port(self) -> int:
        """
        private port.
        """
        return pulumi.get(self, "private_port")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        subnet id.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        vpc id.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="wanDomain")
    def wan_domain(self) -> str:
        """
        wan domain.
        """
        return pulumi.get(self, "wan_domain")

    @property
    @pulumi.getter(name="wanIp")
    def wan_ip(self) -> str:
        """
        wan ip.
        """
        return pulumi.get(self, "wan_ip")

    @property
    @pulumi.getter(name="wanPort")
    def wan_port(self) -> int:
        """
        wan port.
        """
        return pulumi.get(self, "wan_port")


@pulumi.output_type
class GetInstancesListResult(dict):
    def __init__(__self__, *,
                 cluster_id: str,
                 cpu: int,
                 create_time: str,
                 db_kernel_version: str,
                 db_major_version: str,
                 db_version: str,
                 endpoint_id: str,
                 instance_id: str,
                 instance_name: str,
                 instance_type: str,
                 memory: int,
                 pay_mode: str,
                 pay_period_end_time: str,
                 region: str,
                 status: str,
                 status_desc: str,
                 zone: str):
        """
        :param str cluster_id: instance id.
        :param int cpu: cpu cores.
        :param str create_time: create time.
        :param str db_kernel_version: db kernel version.
        :param str db_major_version: db major version.
        :param str db_version: db version.
        :param str endpoint_id: endpoint id.
        :param str instance_id: instance id.
        :param str instance_name: instance name.
        :param str instance_type: instance type.
        :param int memory: memory size, unit is GiB.
        :param str pay_mode: pay mode.
        :param str pay_period_end_time: pay period expired time.
        :param str region: region.
        :param str status: instance status.
        :param str status_desc: status description.
        :param str zone: zone.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "db_kernel_version", db_kernel_version)
        pulumi.set(__self__, "db_major_version", db_major_version)
        pulumi.set(__self__, "db_version", db_version)
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "memory", memory)
        pulumi.set(__self__, "pay_mode", pay_mode)
        pulumi.set(__self__, "pay_period_end_time", pay_period_end_time)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "status_desc", status_desc)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        instance id.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def cpu(self) -> int:
        """
        cpu cores.
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dbKernelVersion")
    def db_kernel_version(self) -> str:
        """
        db kernel version.
        """
        return pulumi.get(self, "db_kernel_version")

    @property
    @pulumi.getter(name="dbMajorVersion")
    def db_major_version(self) -> str:
        """
        db major version.
        """
        return pulumi.get(self, "db_major_version")

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> str:
        """
        db version.
        """
        return pulumi.get(self, "db_version")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> str:
        """
        endpoint id.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        instance name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        instance type.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        memory size, unit is GiB.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="payMode")
    def pay_mode(self) -> str:
        """
        pay mode.
        """
        return pulumi.get(self, "pay_mode")

    @property
    @pulumi.getter(name="payPeriodEndTime")
    def pay_period_end_time(self) -> str:
        """
        pay period expired time.
        """
        return pulumi.get(self, "pay_period_end_time")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        instance status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDesc")
    def status_desc(self) -> str:
        """
        status description.
        """
        return pulumi.get(self, "status_desc")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        zone.
        """
        return pulumi.get(self, "zone")


