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
    'GetRocketmqTopicResult',
    'AwaitableGetRocketmqTopicResult',
    'get_rocketmq_topic',
    'get_rocketmq_topic_output',
]

@pulumi.output_type
class GetRocketmqTopicResult:
    """
    A collection of values returned by getRocketmqTopic.
    """
    def __init__(__self__, cluster_id=None, filter_name=None, filter_types=None, id=None, namespace_id=None, result_output_file=None, topics=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if filter_name and not isinstance(filter_name, str):
            raise TypeError("Expected argument 'filter_name' to be a str")
        pulumi.set(__self__, "filter_name", filter_name)
        if filter_types and not isinstance(filter_types, list):
            raise TypeError("Expected argument 'filter_types' to be a list")
        pulumi.set(__self__, "filter_types", filter_types)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        pulumi.set(__self__, "topics", topics)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> Optional[str]:
        return pulumi.get(self, "filter_name")

    @property
    @pulumi.getter(name="filterTypes")
    def filter_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "filter_types")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> str:
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def topics(self) -> Sequence['outputs.GetRocketmqTopicTopicResult']:
        """
        List of topic information.
        """
        return pulumi.get(self, "topics")


class AwaitableGetRocketmqTopicResult(GetRocketmqTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRocketmqTopicResult(
            cluster_id=self.cluster_id,
            filter_name=self.filter_name,
            filter_types=self.filter_types,
            id=self.id,
            namespace_id=self.namespace_id,
            result_output_file=self.result_output_file,
            topics=self.topics)


def get_rocketmq_topic(cluster_id: Optional[str] = None,
                       filter_name: Optional[str] = None,
                       filter_types: Optional[Sequence[str]] = None,
                       namespace_id: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRocketmqTopicResult:
    """
    Use this data source to query detailed information of tdmqRocketmq topic


    :param str cluster_id: Cluster ID.
    :param str filter_name: Search by topic name. Fuzzy query is supported.
    :param Sequence[str] filter_types: Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
    :param str namespace_id: Namespace.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['filterName'] = filter_name
    __args__['filterTypes'] = filter_types
    __args__['namespaceId'] = namespace_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getRocketmqTopic:getRocketmqTopic', __args__, opts=opts, typ=GetRocketmqTopicResult).value

    return AwaitableGetRocketmqTopicResult(
        cluster_id=__ret__.cluster_id,
        filter_name=__ret__.filter_name,
        filter_types=__ret__.filter_types,
        id=__ret__.id,
        namespace_id=__ret__.namespace_id,
        result_output_file=__ret__.result_output_file,
        topics=__ret__.topics)


@_utilities.lift_output_func(get_rocketmq_topic)
def get_rocketmq_topic_output(cluster_id: Optional[pulumi.Input[str]] = None,
                              filter_name: Optional[pulumi.Input[Optional[str]]] = None,
                              filter_types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              namespace_id: Optional[pulumi.Input[str]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRocketmqTopicResult]:
    """
    Use this data source to query detailed information of tdmqRocketmq topic


    :param str cluster_id: Cluster ID.
    :param str filter_name: Search by topic name. Fuzzy query is supported.
    :param Sequence[str] filter_types: Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
    :param str namespace_id: Namespace.
    :param str result_output_file: Used to save results.
    """
    ...
