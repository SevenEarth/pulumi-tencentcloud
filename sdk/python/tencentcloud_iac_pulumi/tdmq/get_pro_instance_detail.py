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
    'GetProInstanceDetailResult',
    'AwaitableGetProInstanceDetailResult',
    'get_pro_instance_detail',
    'get_pro_instance_detail_output',
]

@pulumi.output_type
class GetProInstanceDetailResult:
    """
    A collection of values returned by getProInstanceDetail.
    """
    def __init__(__self__, cluster_id=None, cluster_infos=None, cluster_spec_infos=None, id=None, network_access_point_infos=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_infos and not isinstance(cluster_infos, list):
            raise TypeError("Expected argument 'cluster_infos' to be a list")
        pulumi.set(__self__, "cluster_infos", cluster_infos)
        if cluster_spec_infos and not isinstance(cluster_spec_infos, list):
            raise TypeError("Expected argument 'cluster_spec_infos' to be a list")
        pulumi.set(__self__, "cluster_spec_infos", cluster_spec_infos)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if network_access_point_infos and not isinstance(network_access_point_infos, list):
            raise TypeError("Expected argument 'network_access_point_infos' to be a list")
        pulumi.set(__self__, "network_access_point_infos", network_access_point_infos)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        Cluster Id.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterInfos")
    def cluster_infos(self) -> Sequence['outputs.GetProInstanceDetailClusterInfoResult']:
        """
        Cluster information.
        """
        return pulumi.get(self, "cluster_infos")

    @property
    @pulumi.getter(name="clusterSpecInfos")
    def cluster_spec_infos(self) -> Sequence['outputs.GetProInstanceDetailClusterSpecInfoResult']:
        """
        Cluster specification informationNote: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "cluster_spec_infos")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="networkAccessPointInfos")
    def network_access_point_infos(self) -> Sequence['outputs.GetProInstanceDetailNetworkAccessPointInfoResult']:
        """
        Cluster network access point informationNote: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "network_access_point_infos")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetProInstanceDetailResult(GetProInstanceDetailResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProInstanceDetailResult(
            cluster_id=self.cluster_id,
            cluster_infos=self.cluster_infos,
            cluster_spec_infos=self.cluster_spec_infos,
            id=self.id,
            network_access_point_infos=self.network_access_point_infos,
            result_output_file=self.result_output_file)


def get_pro_instance_detail(cluster_id: Optional[str] = None,
                            result_output_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProInstanceDetailResult:
    """
    Use this data source to query detailed information of tdmq pro_instance_detail

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    pro_instance_detail = tencentcloud.Tdmq.get_pro_instance_detail(cluster_id="pulsar-9n95ax58b9vn")
    ```


    :param str cluster_id: Cluster Id.
    :param str result_output_file: Used to save results.
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
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getProInstanceDetail:getProInstanceDetail', __args__, opts=opts, typ=GetProInstanceDetailResult).value

    return AwaitableGetProInstanceDetailResult(
        cluster_id=__ret__.cluster_id,
        cluster_infos=__ret__.cluster_infos,
        cluster_spec_infos=__ret__.cluster_spec_infos,
        id=__ret__.id,
        network_access_point_infos=__ret__.network_access_point_infos,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_pro_instance_detail)
def get_pro_instance_detail_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProInstanceDetailResult]:
    """
    Use this data source to query detailed information of tdmq pro_instance_detail

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    pro_instance_detail = tencentcloud.Tdmq.get_pro_instance_detail(cluster_id="pulsar-9n95ax58b9vn")
    ```


    :param str cluster_id: Cluster Id.
    :param str result_output_file: Used to save results.
    """
    ...
