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

__all__ = [
    'GetDatahubTopicResult',
    'AwaitableGetDatahubTopicResult',
    'get_datahub_topic',
    'get_datahub_topic_output',
]

@pulumi.output_type
class GetDatahubTopicResult:
    """
    A collection of values returned by getDatahubTopic.
    """
    def __init__(__self__, id=None, limit=None, offset=None, result_output_file=None, search_word=None, topic_lists=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
        if topic_lists and not isinstance(topic_lists, list):
            raise TypeError("Expected argument 'topic_lists' to be a list")
        pulumi.set(__self__, "topic_lists", topic_lists)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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
    @pulumi.getter(name="topicLists")
    def topic_lists(self) -> Sequence['outputs.GetDatahubTopicTopicListResult']:
        """
        Topic list.
        """
        return pulumi.get(self, "topic_lists")


class AwaitableGetDatahubTopicResult(GetDatahubTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatahubTopicResult(
            id=self.id,
            limit=self.limit,
            offset=self.offset,
            result_output_file=self.result_output_file,
            search_word=self.search_word,
            topic_lists=self.topic_lists)


def get_datahub_topic(limit: Optional[int] = None,
                      offset: Optional[int] = None,
                      result_output_file: Optional[str] = None,
                      search_word: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatahubTopicResult:
    """
    Use this data source to query detailed information of ckafka datahub_topic

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    datahub_topic = tencentcloud.Ckafka.get_datahub_topic()
    ```
    <!--End PulumiCodeChooser -->


    :param int limit: The maximum number of results returned this time, the default is 50, and the maximum value is 50.
    :param int offset: The offset position of this query, the default is 0.
    :param str result_output_file: Used to save results.
    :param str search_word: query key word.
    """
    __args__ = dict()
    __args__['limit'] = limit
    __args__['offset'] = offset
    __args__['resultOutputFile'] = result_output_file
    __args__['searchWord'] = search_word
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getDatahubTopic:getDatahubTopic', __args__, opts=opts, typ=GetDatahubTopicResult).value

    return AwaitableGetDatahubTopicResult(
        id=pulumi.get(__ret__, 'id'),
        limit=pulumi.get(__ret__, 'limit'),
        offset=pulumi.get(__ret__, 'offset'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        search_word=pulumi.get(__ret__, 'search_word'),
        topic_lists=pulumi.get(__ret__, 'topic_lists'))


@_utilities.lift_output_func(get_datahub_topic)
def get_datahub_topic_output(limit: Optional[pulumi.Input[Optional[int]]] = None,
                             offset: Optional[pulumi.Input[Optional[int]]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             search_word: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatahubTopicResult]:
    """
    Use this data source to query detailed information of ckafka datahub_topic

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    datahub_topic = tencentcloud.Ckafka.get_datahub_topic()
    ```
    <!--End PulumiCodeChooser -->


    :param int limit: The maximum number of results returned this time, the default is 50, and the maximum value is 50.
    :param int offset: The offset position of this query, the default is 0.
    :param str result_output_file: Used to save results.
    :param str search_word: query key word.
    """
    ...
