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
    'GetVipInstanceResult',
    'AwaitableGetVipInstanceResult',
    'get_vip_instance',
    'get_vip_instance_output',
]

@pulumi.output_type
class GetVipInstanceResult:
    """
    A collection of values returned by getVipInstance.
    """
    def __init__(__self__, cluster_id=None, cluster_infos=None, id=None, instance_configs=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_infos and not isinstance(cluster_infos, list):
            raise TypeError("Expected argument 'cluster_infos' to be a list")
        pulumi.set(__self__, "cluster_infos", cluster_infos)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_configs and not isinstance(instance_configs, list):
            raise TypeError("Expected argument 'instance_configs' to be a list")
        pulumi.set(__self__, "instance_configs", instance_configs)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterInfos")
    def cluster_infos(self) -> Sequence['outputs.GetVipInstanceClusterInfoResult']:
        return pulumi.get(self, "cluster_infos")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceConfigs")
    def instance_configs(self) -> Sequence['outputs.GetVipInstanceInstanceConfigResult']:
        return pulumi.get(self, "instance_configs")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetVipInstanceResult(GetVipInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVipInstanceResult(
            cluster_id=self.cluster_id,
            cluster_infos=self.cluster_infos,
            id=self.id,
            instance_configs=self.instance_configs,
            result_output_file=self.result_output_file)


def get_vip_instance(cluster_id: Optional[str] = None,
                     result_output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVipInstanceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getVipInstance:getVipInstance', __args__, opts=opts, typ=GetVipInstanceResult).value

    return AwaitableGetVipInstanceResult(
        cluster_id=__ret__.cluster_id,
        cluster_infos=__ret__.cluster_infos,
        id=__ret__.id,
        instance_configs=__ret__.instance_configs,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_vip_instance)
def get_vip_instance_output(cluster_id: Optional[pulumi.Input[str]] = None,
                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVipInstanceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
