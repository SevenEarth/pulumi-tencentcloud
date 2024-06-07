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
    'GetPublishersResult',
    'AwaitableGetPublishersResult',
    'get_publishers',
    'get_publishers_output',
]

@pulumi.output_type
class GetPublishersResult:
    """
    A collection of values returned by getPublishers.
    """
    def __init__(__self__, cluster_id=None, filters=None, id=None, namespace=None, publishers=None, result_output_file=None, sort=None, topic=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if publishers and not isinstance(publishers, list):
            raise TypeError("Expected argument 'publishers' to be a list")
        pulumi.set(__self__, "publishers", publishers)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sort and not isinstance(sort, dict):
            raise TypeError("Expected argument 'sort' to be a dict")
        pulumi.set(__self__, "sort", sort)
        if topic and not isinstance(topic, str):
            raise TypeError("Expected argument 'topic' to be a str")
        pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetPublishersFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def publishers(self) -> Sequence['outputs.GetPublishersPublisherResult']:
        """
        Producer Information ListNote: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "publishers")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def sort(self) -> Optional['outputs.GetPublishersSortResult']:
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter
    def topic(self) -> str:
        return pulumi.get(self, "topic")


class AwaitableGetPublishersResult(GetPublishersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublishersResult(
            cluster_id=self.cluster_id,
            filters=self.filters,
            id=self.id,
            namespace=self.namespace,
            publishers=self.publishers,
            result_output_file=self.result_output_file,
            sort=self.sort,
            topic=self.topic)


def get_publishers(cluster_id: Optional[str] = None,
                   filters: Optional[Sequence[pulumi.InputType['GetPublishersFilterArgs']]] = None,
                   namespace: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   sort: Optional[pulumi.InputType['GetPublishersSortArgs']] = None,
                   topic: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublishersResult:
    """
    Use this data source to query detailed information of tdmq publishers

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    publishers = tencentcloud.Tdmq.get_publishers(cluster_id="pulsar-9n95ax58b9vn",
        filters=[tencentcloud.tdmq.GetPublishersFilterArgs(
            name="ProducerName",
            values=["test"],
        )],
        namespace="keep-ns",
        sort=tencentcloud.tdmq.GetPublishersSortArgs(
            name="ProducerName",
            order="DESC",
        ),
        topic="keep-topic")
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: Cluster ID.
    :param Sequence[pulumi.InputType['GetPublishersFilterArgs']] filters: Parameter filter, support ProducerName, Address field.
    :param str namespace: namespace name.
    :param str result_output_file: Used to save results.
    :param pulumi.InputType['GetPublishersSortArgs'] sort: sorter.
    :param str topic: topic name.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['filters'] = filters
    __args__['namespace'] = namespace
    __args__['resultOutputFile'] = result_output_file
    __args__['sort'] = sort
    __args__['topic'] = topic
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getPublishers:getPublishers', __args__, opts=opts, typ=GetPublishersResult).value

    return AwaitableGetPublishersResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        namespace=pulumi.get(__ret__, 'namespace'),
        publishers=pulumi.get(__ret__, 'publishers'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        sort=pulumi.get(__ret__, 'sort'),
        topic=pulumi.get(__ret__, 'topic'))


@_utilities.lift_output_func(get_publishers)
def get_publishers_output(cluster_id: Optional[pulumi.Input[str]] = None,
                          filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetPublishersFilterArgs']]]]] = None,
                          namespace: Optional[pulumi.Input[str]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          sort: Optional[pulumi.Input[Optional[pulumi.InputType['GetPublishersSortArgs']]]] = None,
                          topic: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPublishersResult]:
    """
    Use this data source to query detailed information of tdmq publishers

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    publishers = tencentcloud.Tdmq.get_publishers(cluster_id="pulsar-9n95ax58b9vn",
        filters=[tencentcloud.tdmq.GetPublishersFilterArgs(
            name="ProducerName",
            values=["test"],
        )],
        namespace="keep-ns",
        sort=tencentcloud.tdmq.GetPublishersSortArgs(
            name="ProducerName",
            order="DESC",
        ),
        topic="keep-topic")
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_id: Cluster ID.
    :param Sequence[pulumi.InputType['GetPublishersFilterArgs']] filters: Parameter filter, support ProducerName, Address field.
    :param str namespace: namespace name.
    :param str result_output_file: Used to save results.
    :param pulumi.InputType['GetPublishersSortArgs'] sort: sorter.
    :param str topic: topic name.
    """
    ...
