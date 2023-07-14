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
    'GetTopicFlowRankingResult',
    'AwaitableGetTopicFlowRankingResult',
    'get_topic_flow_ranking',
    'get_topic_flow_ranking_output',
]

@pulumi.output_type
class GetTopicFlowRankingResult:
    """
    A collection of values returned by getTopicFlowRanking.
    """
    def __init__(__self__, begin_date=None, end_date=None, id=None, instance_id=None, ranking_type=None, result_output_file=None, results=None):
        if begin_date and not isinstance(begin_date, str):
            raise TypeError("Expected argument 'begin_date' to be a str")
        pulumi.set(__self__, "begin_date", begin_date)
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if ranking_type and not isinstance(ranking_type, str):
            raise TypeError("Expected argument 'ranking_type' to be a str")
        pulumi.set(__self__, "ranking_type", ranking_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="beginDate")
    def begin_date(self) -> Optional[str]:
        return pulumi.get(self, "begin_date")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[str]:
        return pulumi.get(self, "end_date")

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
    @pulumi.getter(name="rankingType")
    def ranking_type(self) -> str:
        return pulumi.get(self, "ranking_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetTopicFlowRankingResultResult']:
        """
        result.
        """
        return pulumi.get(self, "results")


class AwaitableGetTopicFlowRankingResult(GetTopicFlowRankingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicFlowRankingResult(
            begin_date=self.begin_date,
            end_date=self.end_date,
            id=self.id,
            instance_id=self.instance_id,
            ranking_type=self.ranking_type,
            result_output_file=self.result_output_file,
            results=self.results)


def get_topic_flow_ranking(begin_date: Optional[str] = None,
                           end_date: Optional[str] = None,
                           instance_id: Optional[str] = None,
                           ranking_type: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicFlowRankingResult:
    """
    Use this data source to query detailed information of ckafka topic_flow_ranking

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    topic_flow_ranking = tencentcloud.Ckafka.get_topic_flow_ranking(begin_date="2023-05-29T00:00:00+08:00",
        end_date="2021-05-29T23:59:59+08:00",
        instance_id="ckafka-xxxxxx",
        ranking_type="PRO")
    ```


    :param str begin_date: BeginDate.
    :param str end_date: EndDate.
    :param str instance_id: InstanceId.
    :param str ranking_type: Ranking type. `PRO`: topic production flow, `CON`: topic consumption traffic.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['beginDate'] = begin_date
    __args__['endDate'] = end_date
    __args__['instanceId'] = instance_id
    __args__['rankingType'] = ranking_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getTopicFlowRanking:getTopicFlowRanking', __args__, opts=opts, typ=GetTopicFlowRankingResult).value

    return AwaitableGetTopicFlowRankingResult(
        begin_date=__ret__.begin_date,
        end_date=__ret__.end_date,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        ranking_type=__ret__.ranking_type,
        result_output_file=__ret__.result_output_file,
        results=__ret__.results)


@_utilities.lift_output_func(get_topic_flow_ranking)
def get_topic_flow_ranking_output(begin_date: Optional[pulumi.Input[Optional[str]]] = None,
                                  end_date: Optional[pulumi.Input[Optional[str]]] = None,
                                  instance_id: Optional[pulumi.Input[str]] = None,
                                  ranking_type: Optional[pulumi.Input[str]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicFlowRankingResult]:
    """
    Use this data source to query detailed information of ckafka topic_flow_ranking

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    topic_flow_ranking = tencentcloud.Ckafka.get_topic_flow_ranking(begin_date="2023-05-29T00:00:00+08:00",
        end_date="2021-05-29T23:59:59+08:00",
        instance_id="ckafka-xxxxxx",
        ranking_type="PRO")
    ```


    :param str begin_date: BeginDate.
    :param str end_date: EndDate.
    :param str instance_id: InstanceId.
    :param str ranking_type: Ranking type. `PRO`: topic production flow, `CON`: topic consumption traffic.
    :param str result_output_file: Used to save results.
    """
    ...
