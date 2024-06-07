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
    'GetSubscribeResult',
    'AwaitableGetSubscribeResult',
    'get_subscribe',
    'get_subscribe_output',
]

@pulumi.output_type
class GetSubscribeResult:
    """
    A collection of values returned by getSubscribe.
    """
    def __init__(__self__, id=None, limit=None, offset=None, result_output_file=None, subscription_lists=None, subscription_name=None, topic_name=None):
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
        if subscription_lists and not isinstance(subscription_lists, list):
            raise TypeError("Expected argument 'subscription_lists' to be a list")
        pulumi.set(__self__, "subscription_lists", subscription_lists)
        if subscription_name and not isinstance(subscription_name, str):
            raise TypeError("Expected argument 'subscription_name' to be a str")
        pulumi.set(__self__, "subscription_name", subscription_name)
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        pulumi.set(__self__, "topic_name", topic_name)

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
    @pulumi.getter(name="subscriptionLists")
    def subscription_lists(self) -> Sequence['outputs.GetSubscribeSubscriptionListResult']:
        """
        Set of subscription attributes.
        """
        return pulumi.get(self, "subscription_lists")

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> Optional[str]:
        """
        Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
        """
        return pulumi.get(self, "subscription_name")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> str:
        return pulumi.get(self, "topic_name")


class AwaitableGetSubscribeResult(GetSubscribeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubscribeResult(
            id=self.id,
            limit=self.limit,
            offset=self.offset,
            result_output_file=self.result_output_file,
            subscription_lists=self.subscription_lists,
            subscription_name=self.subscription_name,
            topic_name=self.topic_name)


def get_subscribe(limit: Optional[int] = None,
                  offset: Optional[int] = None,
                  result_output_file: Optional[str] = None,
                  subscription_name: Optional[str] = None,
                  topic_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubscribeResult:
    """
    Use this data source to query detailed information of tcmq subscribe


    :param int limit: Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
    :param int offset: Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, limit is required. If this parameter is left empty, 0 will be used by default.
    :param str result_output_file: Used to save results.
    :param str subscription_name: Fuzzy search by SubscriptionName.
    :param str topic_name: Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
    """
    __args__ = dict()
    __args__['limit'] = limit
    __args__['offset'] = offset
    __args__['resultOutputFile'] = result_output_file
    __args__['subscriptionName'] = subscription_name
    __args__['topicName'] = topic_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tcmq/getSubscribe:getSubscribe', __args__, opts=opts, typ=GetSubscribeResult).value

    return AwaitableGetSubscribeResult(
        id=pulumi.get(__ret__, 'id'),
        limit=pulumi.get(__ret__, 'limit'),
        offset=pulumi.get(__ret__, 'offset'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        subscription_lists=pulumi.get(__ret__, 'subscription_lists'),
        subscription_name=pulumi.get(__ret__, 'subscription_name'),
        topic_name=pulumi.get(__ret__, 'topic_name'))


@_utilities.lift_output_func(get_subscribe)
def get_subscribe_output(limit: Optional[pulumi.Input[Optional[int]]] = None,
                         offset: Optional[pulumi.Input[Optional[int]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         subscription_name: Optional[pulumi.Input[Optional[str]]] = None,
                         topic_name: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubscribeResult]:
    """
    Use this data source to query detailed information of tcmq subscribe


    :param int limit: Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
    :param int offset: Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, limit is required. If this parameter is left empty, 0 will be used by default.
    :param str result_output_file: Used to save results.
    :param str subscription_name: Fuzzy search by SubscriptionName.
    :param str topic_name: Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
    """
    ...
