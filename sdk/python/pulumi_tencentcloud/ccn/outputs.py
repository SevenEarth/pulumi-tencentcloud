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
    'GetBandwidthLimitsLimitResult',
    'GetInstancesInstanceListResult',
    'GetInstancesInstanceListAttachmentListResult',
]

@pulumi.output_type
class GetBandwidthLimitsLimitResult(dict):
    def __init__(__self__, *,
                 bandwidth_limit: int,
                 dst_region: str,
                 region: str):
        """
        :param int bandwidth_limit: Limitation of bandwidth.
        :param str dst_region: Destination area restriction.
        :param str region: Limitation of region.
        """
        pulumi.set(__self__, "bandwidth_limit", bandwidth_limit)
        pulumi.set(__self__, "dst_region", dst_region)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="bandwidthLimit")
    def bandwidth_limit(self) -> int:
        """
        Limitation of bandwidth.
        """
        return pulumi.get(self, "bandwidth_limit")

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> str:
        """
        Destination area restriction.
        """
        return pulumi.get(self, "dst_region")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Limitation of region.
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class GetInstancesInstanceListResult(dict):
    def __init__(__self__, *,
                 attachment_lists: Sequence['outputs.GetInstancesInstanceListAttachmentListResult'],
                 bandwidth_limit_type: str,
                 ccn_id: str,
                 charge_type: str,
                 create_time: str,
                 description: str,
                 name: str,
                 qos: str,
                 state: str):
        """
        :param Sequence['GetInstancesInstanceListAttachmentListArgs'] attachment_lists: Information list of instance is attached.
        :param str bandwidth_limit_type: The speed limit type.
        :param str ccn_id: ID of the CCN to be queried.
        :param str charge_type: Billing mode.
        :param str create_time: Creation time of resource.
        :param str description: Description of the CCN.
        :param str name: Name of the CCN to be queried.
        :param str qos: Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'.
        :param str state: States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.
        """
        pulumi.set(__self__, "attachment_lists", attachment_lists)
        pulumi.set(__self__, "bandwidth_limit_type", bandwidth_limit_type)
        pulumi.set(__self__, "ccn_id", ccn_id)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "qos", qos)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="attachmentLists")
    def attachment_lists(self) -> Sequence['outputs.GetInstancesInstanceListAttachmentListResult']:
        """
        Information list of instance is attached.
        """
        return pulumi.get(self, "attachment_lists")

    @property
    @pulumi.getter(name="bandwidthLimitType")
    def bandwidth_limit_type(self) -> str:
        """
        The speed limit type.
        """
        return pulumi.get(self, "bandwidth_limit_type")

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> str:
        """
        ID of the CCN to be queried.
        """
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        Billing mode.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time of resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the CCN.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the CCN to be queried.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def qos(self) -> str:
        """
        Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'.
        """
        return pulumi.get(self, "qos")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.
        """
        return pulumi.get(self, "state")


@pulumi.output_type
class GetInstancesInstanceListAttachmentListResult(dict):
    def __init__(__self__, *,
                 attached_time: str,
                 cidr_blocks: Sequence[str],
                 instance_id: str,
                 instance_region: str,
                 instance_type: str,
                 state: str):
        """
        :param str attached_time: Time of attaching.
        :param Sequence[str] cidr_blocks: A network address block of the instance that is attached.
        :param str instance_id: ID of instance is attached.
        :param str instance_region: The region that the instance locates at.
        :param str instance_type: Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW.
        :param str state: States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.
        """
        pulumi.set(__self__, "attached_time", attached_time)
        pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="attachedTime")
    def attached_time(self) -> str:
        """
        Time of attaching.
        """
        return pulumi.get(self, "attached_time")

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Sequence[str]:
        """
        A network address block of the instance that is attached.
        """
        return pulumi.get(self, "cidr_blocks")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        ID of instance is attached.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> str:
        """
        The region that the instance locates at.
        """
        return pulumi.get(self, "instance_region")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.
        """
        return pulumi.get(self, "state")


