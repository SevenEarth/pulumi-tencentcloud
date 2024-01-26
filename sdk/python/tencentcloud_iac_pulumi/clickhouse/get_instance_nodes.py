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
    'GetInstanceNodesResult',
    'AwaitableGetInstanceNodesResult',
    'get_instance_nodes',
    'get_instance_nodes_output',
]

@pulumi.output_type
class GetInstanceNodesResult:
    """
    A collection of values returned by getInstanceNodes.
    """
    def __init__(__self__, display_policy=None, force_all=None, id=None, instance_id=None, instance_nodes_lists=None, node_role=None, result_output_file=None):
        if display_policy and not isinstance(display_policy, str):
            raise TypeError("Expected argument 'display_policy' to be a str")
        pulumi.set(__self__, "display_policy", display_policy)
        if force_all and not isinstance(force_all, bool):
            raise TypeError("Expected argument 'force_all' to be a bool")
        pulumi.set(__self__, "force_all", force_all)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_nodes_lists and not isinstance(instance_nodes_lists, list):
            raise TypeError("Expected argument 'instance_nodes_lists' to be a list")
        pulumi.set(__self__, "instance_nodes_lists", instance_nodes_lists)
        if node_role and not isinstance(node_role, str):
            raise TypeError("Expected argument 'node_role' to be a str")
        pulumi.set(__self__, "node_role", node_role)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="displayPolicy")
    def display_policy(self) -> Optional[str]:
        return pulumi.get(self, "display_policy")

    @property
    @pulumi.getter(name="forceAll")
    def force_all(self) -> Optional[bool]:
        return pulumi.get(self, "force_all")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceNodesLists")
    def instance_nodes_lists(self) -> Sequence['outputs.GetInstanceNodesInstanceNodesListResult']:
        """
        Total number of instance nodes.
        """
        return pulumi.get(self, "instance_nodes_lists")

    @property
    @pulumi.getter(name="nodeRole")
    def node_role(self) -> Optional[str]:
        return pulumi.get(self, "node_role")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceNodesResult(GetInstanceNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceNodesResult(
            display_policy=self.display_policy,
            force_all=self.force_all,
            id=self.id,
            instance_id=self.instance_id,
            instance_nodes_lists=self.instance_nodes_lists,
            node_role=self.node_role,
            result_output_file=self.result_output_file)


def get_instance_nodes(display_policy: Optional[str] = None,
                       force_all: Optional[bool] = None,
                       instance_id: Optional[str] = None,
                       node_role: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceNodesResult:
    """
    Use this data source to query detailed information of clickhouse instance_nodes

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_nodes = tencentcloud.Clickhouse.get_instance_nodes(display_policy="all",
        force_all=True,
        instance_id="cdwch-mvfjh373",
        node_role="data")
    ```


    :param str display_policy: Display strategy, display all when All.
    :param bool force_all: When true, returns all nodes, that is, the Limit is infinitely large.
    :param str instance_id: InstanceId.
    :param str node_role: Cluster role type, default is `data` data node.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['displayPolicy'] = display_policy
    __args__['forceAll'] = force_all
    __args__['instanceId'] = instance_id
    __args__['nodeRole'] = node_role
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Clickhouse/getInstanceNodes:getInstanceNodes', __args__, opts=opts, typ=GetInstanceNodesResult).value

    return AwaitableGetInstanceNodesResult(
        display_policy=__ret__.display_policy,
        force_all=__ret__.force_all,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_nodes_lists=__ret__.instance_nodes_lists,
        node_role=__ret__.node_role,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_instance_nodes)
def get_instance_nodes_output(display_policy: Optional[pulumi.Input[Optional[str]]] = None,
                              force_all: Optional[pulumi.Input[Optional[bool]]] = None,
                              instance_id: Optional[pulumi.Input[str]] = None,
                              node_role: Optional[pulumi.Input[Optional[str]]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceNodesResult]:
    """
    Use this data source to query detailed information of clickhouse instance_nodes

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_nodes = tencentcloud.Clickhouse.get_instance_nodes(display_policy="all",
        force_all=True,
        instance_id="cdwch-mvfjh373",
        node_role="data")
    ```


    :param str display_policy: Display strategy, display all when All.
    :param bool force_all: When true, returns all nodes, that is, the Limit is infinitely large.
    :param str instance_id: InstanceId.
    :param str node_role: Cluster role type, default is `data` data node.
    :param str result_output_file: Used to save results.
    """
    ...
