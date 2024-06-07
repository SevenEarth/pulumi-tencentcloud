# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetGatewaySnatsResult',
    'AwaitableGetGatewaySnatsResult',
    'get_gateway_snats',
    'get_gateway_snats_output',
]

@pulumi.output_type
class GetGatewaySnatsResult:
    """
    A collection of values returned by getGatewaySnats.
    """
    def __init__(__self__, description=None, id=None, instance_id=None, nat_gateway_id=None, public_ip_addrs=None, result_output_file=None, snat_lists=None, subnet_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if nat_gateway_id and not isinstance(nat_gateway_id, str):
            raise TypeError("Expected argument 'nat_gateway_id' to be a str")
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if public_ip_addrs and not isinstance(public_ip_addrs, list):
            raise TypeError("Expected argument 'public_ip_addrs' to be a list")
        pulumi.set(__self__, "public_ip_addrs", public_ip_addrs)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if snat_lists and not isinstance(snat_lists, list):
            raise TypeError("Expected argument 'snat_lists' to be a list")
        pulumi.set(__self__, "snat_lists", snat_lists)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> str:
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="publicIpAddrs")
    def public_ip_addrs(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "public_ip_addrs")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="snatLists")
    def snat_lists(self) -> Sequence['outputs.GetGatewaySnatsSnatListResult']:
        """
        Information list of the nat gateway snat.
        """
        return pulumi.get(self, "snat_lists")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        return pulumi.get(self, "subnet_id")


class AwaitableGetGatewaySnatsResult(GetGatewaySnatsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewaySnatsResult(
            description=self.description,
            id=self.id,
            instance_id=self.instance_id,
            nat_gateway_id=self.nat_gateway_id,
            public_ip_addrs=self.public_ip_addrs,
            result_output_file=self.result_output_file,
            snat_lists=self.snat_lists,
            subnet_id=self.subnet_id)


def get_gateway_snats(description: Optional[str] = None,
                      instance_id: Optional[str] = None,
                      nat_gateway_id: Optional[str] = None,
                      public_ip_addrs: Optional[Sequence[str]] = None,
                      result_output_file: Optional[str] = None,
                      subnet_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewaySnatsResult:
    """
    Use this data source to query detailed information of VPN gateways.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    snat = tencentcloud.Nat.get_gateway_snats(nat_gateway_id=tencentcloud_nat_gateway["my_nat"]["id"],
        subnet_id=tencentcloud_nat_gateway_snat["my_subnet"]["id"],
        public_ip_addrs=["50.29.23.234"],
        description="snat demo",
        result_output_file="./snat.txt")
    ```
    <!--End PulumiCodeChooser -->


    :param str description: Description.
    :param str instance_id: Instance ID.
    :param str nat_gateway_id: NAT gateway ID.
    :param Sequence[str] public_ip_addrs: Elastic IP address pool.
    :param str result_output_file: Used to save results.
    :param str subnet_id: Subnet instance ID.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['instanceId'] = instance_id
    __args__['natGatewayId'] = nat_gateway_id
    __args__['publicIpAddrs'] = public_ip_addrs
    __args__['resultOutputFile'] = result_output_file
    __args__['subnetId'] = subnet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Nat/getGatewaySnats:getGatewaySnats', __args__, opts=opts, typ=GetGatewaySnatsResult).value

    return AwaitableGetGatewaySnatsResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        nat_gateway_id=pulumi.get(__ret__, 'nat_gateway_id'),
        public_ip_addrs=pulumi.get(__ret__, 'public_ip_addrs'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        snat_lists=pulumi.get(__ret__, 'snat_lists'),
        subnet_id=pulumi.get(__ret__, 'subnet_id'))


@_utilities.lift_output_func(get_gateway_snats)
def get_gateway_snats_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                             instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                             nat_gateway_id: Optional[pulumi.Input[str]] = None,
                             public_ip_addrs: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewaySnatsResult]:
    """
    Use this data source to query detailed information of VPN gateways.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    snat = tencentcloud.Nat.get_gateway_snats(nat_gateway_id=tencentcloud_nat_gateway["my_nat"]["id"],
        subnet_id=tencentcloud_nat_gateway_snat["my_subnet"]["id"],
        public_ip_addrs=["50.29.23.234"],
        description="snat demo",
        result_output_file="./snat.txt")
    ```
    <!--End PulumiCodeChooser -->


    :param str description: Description.
    :param str instance_id: Instance ID.
    :param str nat_gateway_id: NAT gateway ID.
    :param Sequence[str] public_ip_addrs: Elastic IP address pool.
    :param str result_output_file: Used to save results.
    :param str subnet_id: Subnet instance ID.
    """
    ...
