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
    'GetDatahubGroupOffsetsResult',
    'AwaitableGetDatahubGroupOffsetsResult',
    'get_datahub_group_offsets',
    'get_datahub_group_offsets_output',
]

@pulumi.output_type
class GetDatahubGroupOffsetsResult:
    """
    A collection of values returned by getDatahubGroupOffsets.
    """
    def __init__(__self__, group=None, id=None, name=None, result_output_file=None, search_word=None, topic_lists=None):
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if search_word and not isinstance(search_word, str):
            raise TypeError("Expected argument 'search_word' to be a str")
        pulumi.set(__self__, "search_word", search_word)
        if topic_lists and not isinstance(topic_lists, list):
            raise TypeError("Expected argument 'topic_lists' to be a list")
        pulumi.set(__self__, "topic_lists", topic_lists)

    @property
    @pulumi.getter
    def group(self) -> str:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="searchWord")
    def search_word(self) -> Optional[str]:
        return pulumi.get(self, "search_word")

    @property
    @pulumi.getter(name="topicLists")
    def topic_lists(self) -> Sequence['outputs.GetDatahubGroupOffsetsTopicListResult']:
        """
        The topic array, where each element is a json object.
        """
        return pulumi.get(self, "topic_lists")


class AwaitableGetDatahubGroupOffsetsResult(GetDatahubGroupOffsetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatahubGroupOffsetsResult(
            group=self.group,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            search_word=self.search_word,
            topic_lists=self.topic_lists)


def get_datahub_group_offsets(group: Optional[str] = None,
                              name: Optional[str] = None,
                              result_output_file: Optional[str] = None,
                              search_word: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatahubGroupOffsetsResult:
    """
    Use this data source to query detailed information of ckafka datahub_group_offsets

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    datahub_group_offsets = tencentcloud.Ckafka.get_datahub_group_offsets()
    ```


    :param str group: Kafka consumer group.
    :param str name: topic name that the task subscribe.
    :param str result_output_file: Used to save results.
    :param str search_word: fuzzy match topicName.
    """
    __args__ = dict()
    __args__['group'] = group
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['searchWord'] = search_word
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getDatahubGroupOffsets:getDatahubGroupOffsets', __args__, opts=opts, typ=GetDatahubGroupOffsetsResult).value

    return AwaitableGetDatahubGroupOffsetsResult(
        group=__ret__.group,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        search_word=__ret__.search_word,
        topic_lists=__ret__.topic_lists)


@_utilities.lift_output_func(get_datahub_group_offsets)
def get_datahub_group_offsets_output(group: Optional[pulumi.Input[str]] = None,
                                     name: Optional[pulumi.Input[str]] = None,
                                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                     search_word: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatahubGroupOffsetsResult]:
    """
    Use this data source to query detailed information of ckafka datahub_group_offsets

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    datahub_group_offsets = tencentcloud.Ckafka.get_datahub_group_offsets()
    ```


    :param str group: Kafka consumer group.
    :param str name: topic name that the task subscribe.
    :param str result_output_file: Used to save results.
    :param str search_word: fuzzy match topicName.
    """
    ...
