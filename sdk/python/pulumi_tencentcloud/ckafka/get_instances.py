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
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
    'get_instances_output',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, filters=None, id=None, instance_ids=None, instance_lists=None, limit=None, offset=None, result_output_file=None, search_word=None, statuses=None, tag_key=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if offset and not isinstance(offset, int):
            raise TypeError("Expected argument 'offset' to be a int")
        pulumi.set(__self__, "offset", offset)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if search_word and not isinstance(search_word, str):
            raise TypeError("Expected argument 'search_word' to be a str")
        pulumi.set(__self__, "search_word", search_word)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if tag_key and not isinstance(tag_key, str):
            raise TypeError("Expected argument 'tag_key' to be a str")
        pulumi.set(__self__, "tag_key", tag_key)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetInstancesFilterResult']]:
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
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.GetInstancesInstanceListResult']:
        """
        A list of ckafka users. Each element contains the following attributes:
        """
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def offset(self) -> Optional[int]:
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="searchWord")
    def search_word(self) -> Optional[str]:
        return pulumi.get(self, "search_word")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[int]]:
        """
        The status of the instance. 0: Created, 1: Running, 2: Delete: 5 Quarantined, -1 Creation failed.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> Optional[str]:
        """
        Tag Key.
        """
        return pulumi.get(self, "tag_key")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            filters=self.filters,
            id=self.id,
            instance_ids=self.instance_ids,
            instance_lists=self.instance_lists,
            limit=self.limit,
            offset=self.offset,
            result_output_file=self.result_output_file,
            search_word=self.search_word,
            statuses=self.statuses,
            tag_key=self.tag_key)


def get_instances(filters: Optional[Sequence[pulumi.InputType['GetInstancesFilterArgs']]] = None,
                  instance_ids: Optional[Sequence[str]] = None,
                  limit: Optional[int] = None,
                  offset: Optional[int] = None,
                  result_output_file: Optional[str] = None,
                  search_word: Optional[str] = None,
                  statuses: Optional[Sequence[int]] = None,
                  tag_key: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    Use this data source to query detailed instance information of Ckafka

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Ckafka.get_instances(instance_ids=["ckafka-vv7wpvae"])
    ```


    :param Sequence[pulumi.InputType['GetInstancesFilterArgs']] filters: Filter. filter.name supports ('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId'), filter.values can pass up to 10 values.
    :param Sequence[str] instance_ids: Filter by instance ID.
    :param int limit: The number of pages, default is `10`.
    :param int offset: The page start offset, default is `0`.
    :param str result_output_file: Used to save results.
    :param str search_word: Filter by instance name, support fuzzy query.
    :param Sequence[int] statuses: (Filter Criteria) The status of the instance. 0: Create, 1: Run, 2: Delete, do not fill the default return all.
    :param str tag_key: Matches the tag key value.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['instanceIds'] = instance_ids
    __args__['limit'] = limit
    __args__['offset'] = offset
    __args__['resultOutputFile'] = result_output_file
    __args__['searchWord'] = search_word
    __args__['statuses'] = statuses
    __args__['tagKey'] = tag_key
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        filters=__ret__.filters,
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        instance_lists=__ret__.instance_lists,
        limit=__ret__.limit,
        offset=__ret__.offset,
        result_output_file=__ret__.result_output_file,
        search_word=__ret__.search_word,
        statuses=__ret__.statuses,
        tag_key=__ret__.tag_key)


@_utilities.lift_output_func(get_instances)
def get_instances_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetInstancesFilterArgs']]]]] = None,
                         instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         limit: Optional[pulumi.Input[Optional[int]]] = None,
                         offset: Optional[pulumi.Input[Optional[int]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         search_word: Optional[pulumi.Input[Optional[str]]] = None,
                         statuses: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                         tag_key: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    Use this data source to query detailed instance information of Ckafka

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Ckafka.get_instances(instance_ids=["ckafka-vv7wpvae"])
    ```


    :param Sequence[pulumi.InputType['GetInstancesFilterArgs']] filters: Filter. filter.name supports ('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId'), filter.values can pass up to 10 values.
    :param Sequence[str] instance_ids: Filter by instance ID.
    :param int limit: The number of pages, default is `10`.
    :param int offset: The page start offset, default is `0`.
    :param str result_output_file: Used to save results.
    :param str search_word: Filter by instance name, support fuzzy query.
    :param Sequence[int] statuses: (Filter Criteria) The status of the instance. 0: Create, 1: Run, 2: Delete, do not fill the default return all.
    :param str tag_key: Matches the tag key value.
    """
    ...
