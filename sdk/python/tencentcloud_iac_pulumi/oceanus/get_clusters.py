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
from ._inputs import *

__all__ = [
    'GetClustersResult',
    'AwaitableGetClustersResult',
    'get_clusters',
    'get_clusters_output',
]

@pulumi.output_type
class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, cluster_ids=None, cluster_sets=None, filters=None, id=None, order_type=None, result_output_file=None, work_space_id=None):
        if cluster_ids and not isinstance(cluster_ids, list):
            raise TypeError("Expected argument 'cluster_ids' to be a list")
        pulumi.set(__self__, "cluster_ids", cluster_ids)
        if cluster_sets and not isinstance(cluster_sets, list):
            raise TypeError("Expected argument 'cluster_sets' to be a list")
        pulumi.set(__self__, "cluster_sets", cluster_sets)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order_type and not isinstance(order_type, int):
            raise TypeError("Expected argument 'order_type' to be a int")
        pulumi.set(__self__, "order_type", order_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if work_space_id and not isinstance(work_space_id, str):
            raise TypeError("Expected argument 'work_space_id' to be a str")
        pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cluster_ids")

    @property
    @pulumi.getter(name="clusterSets")
    def cluster_sets(self) -> Sequence['outputs.GetClustersClusterSetResult']:
        """
        Cluster list.
        """
        return pulumi.get(self, "cluster_sets")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetClustersFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orderType")
    def order_type(self) -> Optional[int]:
        return pulumi.get(self, "order_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[str]:
        """
        Workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")


class AwaitableGetClustersResult(GetClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClustersResult(
            cluster_ids=self.cluster_ids,
            cluster_sets=self.cluster_sets,
            filters=self.filters,
            id=self.id,
            order_type=self.order_type,
            result_output_file=self.result_output_file,
            work_space_id=self.work_space_id)


def get_clusters(cluster_ids: Optional[Sequence[str]] = None,
                 filters: Optional[Sequence[pulumi.InputType['GetClustersFilterArgs']]] = None,
                 order_type: Optional[int] = None,
                 result_output_file: Optional[str] = None,
                 work_space_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClustersResult:
    """
    Use this data source to query detailed information of oceanus clusters

    ## Example Usage

    ### Query all clusters

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Oceanus.get_clusters()
    ```
    <!--End PulumiCodeChooser -->

    ### Query the specified cluster

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Oceanus.get_clusters(cluster_ids=["cluster-5c42n3a5"],
        filters=[tencentcloud.oceanus.GetClustersFilterArgs(
            name="name",
            values=["tf_example"],
        )],
        order_type=1,
        work_space_id="space-2idq8wbr")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] cluster_ids: Query one or more clusters by their ID. The maximum number of clusters that can be queried at once is 100.
    :param Sequence[pulumi.InputType['GetClustersFilterArgs']] filters: The filtering rules.
    :param int order_type: The sorting rule of the cluster information results. Possible values are 1 (sort by time in descending order), 2 (sort by time in ascending order), and 3 (sort by status).
    :param str result_output_file: Used to save results.
    :param str work_space_id: Workspace SerialId.
    """
    __args__ = dict()
    __args__['clusterIds'] = cluster_ids
    __args__['filters'] = filters
    __args__['orderType'] = order_type
    __args__['resultOutputFile'] = result_output_file
    __args__['workSpaceId'] = work_space_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Oceanus/getClusters:getClusters', __args__, opts=opts, typ=GetClustersResult).value

    return AwaitableGetClustersResult(
        cluster_ids=pulumi.get(__ret__, 'cluster_ids'),
        cluster_sets=pulumi.get(__ret__, 'cluster_sets'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        order_type=pulumi.get(__ret__, 'order_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        work_space_id=pulumi.get(__ret__, 'work_space_id'))


@_utilities.lift_output_func(get_clusters)
def get_clusters_output(cluster_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetClustersFilterArgs']]]]] = None,
                        order_type: Optional[pulumi.Input[Optional[int]]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        work_space_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClustersResult]:
    """
    Use this data source to query detailed information of oceanus clusters

    ## Example Usage

    ### Query all clusters

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Oceanus.get_clusters()
    ```
    <!--End PulumiCodeChooser -->

    ### Query the specified cluster

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Oceanus.get_clusters(cluster_ids=["cluster-5c42n3a5"],
        filters=[tencentcloud.oceanus.GetClustersFilterArgs(
            name="name",
            values=["tf_example"],
        )],
        order_type=1,
        work_space_id="space-2idq8wbr")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] cluster_ids: Query one or more clusters by their ID. The maximum number of clusters that can be queried at once is 100.
    :param Sequence[pulumi.InputType['GetClustersFilterArgs']] filters: The filtering rules.
    :param int order_type: The sorting rule of the cluster information results. Possible values are 1 (sort by time in descending order), 2 (sort by time in ascending order), and 3 (sort by status).
    :param str result_output_file: Used to save results.
    :param str work_space_id: Workspace SerialId.
    """
    ...
