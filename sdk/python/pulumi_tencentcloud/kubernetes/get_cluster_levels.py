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
    'GetClusterLevelsResult',
    'AwaitableGetClusterLevelsResult',
    'get_cluster_levels',
    'get_cluster_levels_output',
]

@pulumi.output_type
class GetClusterLevelsResult:
    """
    A collection of values returned by getClusterLevels.
    """
    def __init__(__self__, cluster_id=None, id=None, lists=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetClusterLevelsListResult']:
        """
        List of level information.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetClusterLevelsResult(GetClusterLevelsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterLevelsResult(
            cluster_id=self.cluster_id,
            id=self.id,
            lists=self.lists,
            result_output_file=self.result_output_file)


def get_cluster_levels(cluster_id: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterLevelsResult:
    """
    Provide a datasource to query TKE cluster levels.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Kubernetes.get_cluster_levels()
    pulumi.export("level5", foo.lists[0].alias)
    ```


    :param str cluster_id: Specify cluster Id, if set will only query current cluster's available levels.
    :param str result_output_file: Used for save result.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Kubernetes/getClusterLevels:getClusterLevels', __args__, opts=opts, typ=GetClusterLevelsResult).value

    return AwaitableGetClusterLevelsResult(
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        lists=__ret__.lists,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_cluster_levels)
def get_cluster_levels_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterLevelsResult]:
    """
    Provide a datasource to query TKE cluster levels.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Kubernetes.get_cluster_levels()
    pulumi.export("level5", foo.lists[0].alias)
    ```


    :param str cluster_id: Specify cluster Id, if set will only query current cluster's available levels.
    :param str result_output_file: Used for save result.
    """
    ...
