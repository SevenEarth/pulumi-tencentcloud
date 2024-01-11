# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetClusterInstancesResult',
    'AwaitableGetClusterInstancesResult',
    'get_cluster_instances',
    'get_cluster_instances_output',
]

@pulumi.output_type
class GetClusterInstancesResult:
    """
    A collection of values returned by getClusterInstances.
    """
    def __init__(__self__, cluster_id=None, filters=None, id=None, instance_ids=None, instance_role=None, instance_sets=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if instance_role and not isinstance(instance_role, str):
            raise TypeError("Expected argument 'instance_role' to be a str")
        pulumi.set(__self__, "instance_role", instance_role)
        if instance_sets and not isinstance(instance_sets, list):
            raise TypeError("Expected argument 'instance_sets' to be a list")
        pulumi.set(__self__, "instance_sets", instance_sets)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetClusterInstancesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="instanceRole")
    def instance_role(self) -> Optional[str]:
        """
        Node role, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, default is WORKER.
        """
        return pulumi.get(self, "instance_role")

    @property
    @pulumi.getter(name="instanceSets")
    def instance_sets(self) -> Sequence['outputs.GetClusterInstancesInstanceSetResult']:
        """
        List of instances in the cluster.
        """
        return pulumi.get(self, "instance_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetClusterInstancesResult(GetClusterInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterInstancesResult(
            cluster_id=self.cluster_id,
            filters=self.filters,
            id=self.id,
            instance_ids=self.instance_ids,
            instance_role=self.instance_role,
            instance_sets=self.instance_sets,
            result_output_file=self.result_output_file)


def get_cluster_instances(cluster_id: Optional[str] = None,
                          filters: Optional[Sequence[pulumi.InputType['GetClusterInstancesFilterArgs']]] = None,
                          instance_ids: Optional[Sequence[str]] = None,
                          instance_role: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterInstancesResult:
    """
    Use this data source to query detailed information of kubernetes cluster_instances

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cluster_instances = tencentcloud.Kubernetes.get_cluster_instances(cluster_id="cls-ely08ic4",
        filters=[tencentcloud.kubernetes.GetClusterInstancesFilterArgs(
            name="nodepool-id",
            values=["np-p4e6whqu"],
        )],
        instance_ids=["ins-kqmx8dm2"],
        instance_role="WORKER")
    ```


    :param str cluster_id: ID of the cluster.
    :param Sequence[pulumi.InputType['GetClusterInstancesFilterArgs']] filters: List of filter conditions. The optional values of Name are `nodepool-id` and `nodepool-instance-type`. Name is `nodepool-id`, which means filtering machines based on node pool id, and Value is the specific node pool id. Name is `nodepool-instance-type`, which indicates how the node is added to the node pool. Value is MANUALLY_ADDED (manually added to the node pool), AUTOSCALING_ADDED (joined by scaling group expansion method), ALL (manually join the node pool and join the node pool through scaling group expansion).
    :param Sequence[str] instance_ids: List of node instance IDs to be obtained. If it is empty, it means pulling all node instances in the cluster.
    :param str instance_role: Node role, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, default is WORKER.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['filters'] = filters
    __args__['instanceIds'] = instance_ids
    __args__['instanceRole'] = instance_role
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Kubernetes/getClusterInstances:getClusterInstances', __args__, opts=opts, typ=GetClusterInstancesResult).value

    return AwaitableGetClusterInstancesResult(
        cluster_id=__ret__.cluster_id,
        filters=__ret__.filters,
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        instance_role=__ret__.instance_role,
        instance_sets=__ret__.instance_sets,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_cluster_instances)
def get_cluster_instances_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                 filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetClusterInstancesFilterArgs']]]]] = None,
                                 instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 instance_role: Optional[pulumi.Input[Optional[str]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterInstancesResult]:
    """
    Use this data source to query detailed information of kubernetes cluster_instances

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cluster_instances = tencentcloud.Kubernetes.get_cluster_instances(cluster_id="cls-ely08ic4",
        filters=[tencentcloud.kubernetes.GetClusterInstancesFilterArgs(
            name="nodepool-id",
            values=["np-p4e6whqu"],
        )],
        instance_ids=["ins-kqmx8dm2"],
        instance_role="WORKER")
    ```


    :param str cluster_id: ID of the cluster.
    :param Sequence[pulumi.InputType['GetClusterInstancesFilterArgs']] filters: List of filter conditions. The optional values of Name are `nodepool-id` and `nodepool-instance-type`. Name is `nodepool-id`, which means filtering machines based on node pool id, and Value is the specific node pool id. Name is `nodepool-instance-type`, which indicates how the node is added to the node pool. Value is MANUALLY_ADDED (manually added to the node pool), AUTOSCALING_ADDED (joined by scaling group expansion method), ALL (manually join the node pool and join the node pool through scaling group expansion).
    :param Sequence[str] instance_ids: List of node instance IDs to be obtained. If it is empty, it means pulling all node instances in the cluster.
    :param str instance_role: Node role, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, default is WORKER.
    :param str result_output_file: Used to save results.
    """
    ...