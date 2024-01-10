# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetOverviewDdosTrendResult',
    'AwaitableGetOverviewDdosTrendResult',
    'get_overview_ddos_trend',
    'get_overview_ddos_trend_output',
]

@pulumi.output_type
class GetOverviewDdosTrendResult:
    """
    A collection of values returned by getOverviewDdosTrend.
    """
    def __init__(__self__, business=None, datas=None, end_time=None, id=None, ip_lists=None, metric_name=None, period=None, result_output_file=None, start_time=None):
        if business and not isinstance(business, str):
            raise TypeError("Expected argument 'business' to be a str")
        pulumi.set(__self__, "business", business)
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_lists and not isinstance(ip_lists, list):
            raise TypeError("Expected argument 'ip_lists' to be a list")
        pulumi.set(__self__, "ip_lists", ip_lists)
        if metric_name and not isinstance(metric_name, str):
            raise TypeError("Expected argument 'metric_name' to be a str")
        pulumi.set(__self__, "metric_name", metric_name)
        if period and not isinstance(period, int):
            raise TypeError("Expected argument 'period' to be a int")
        pulumi.set(__self__, "period", period)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter
    def business(self) -> Optional[str]:
        return pulumi.get(self, "business")

    @property
    @pulumi.getter
    def datas(self) -> Sequence[int]:
        """
        Array, attack traffic bandwidth in Mbps, packet rate in pps.
        """
        return pulumi.get(self, "datas")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipLists")
    def ip_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ip_lists")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> str:
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter
    def period(self) -> int:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        return pulumi.get(self, "start_time")


class AwaitableGetOverviewDdosTrendResult(GetOverviewDdosTrendResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOverviewDdosTrendResult(
            business=self.business,
            datas=self.datas,
            end_time=self.end_time,
            id=self.id,
            ip_lists=self.ip_lists,
            metric_name=self.metric_name,
            period=self.period,
            result_output_file=self.result_output_file,
            start_time=self.start_time)


def get_overview_ddos_trend(business: Optional[str] = None,
                            end_time: Optional[str] = None,
                            ip_lists: Optional[Sequence[str]] = None,
                            metric_name: Optional[str] = None,
                            period: Optional[int] = None,
                            result_output_file: Optional[str] = None,
                            start_time: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOverviewDdosTrendResult:
    """
    Use this data source to query detailed information of antiddos overview ddos trend

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    overview_ddos_trend = tencentcloud.Antiddos.get_overview_ddos_trend(business="bgpip",
        end_time="2023-11-21 14:16:23",
        metric_name="bps",
        period=300,
        start_time="2023-11-20 14:16:23")
    ```


    :param str business: Dayu sub product code (bgpip represents advanced defense IP; net represents professional version of advanced defense IP).
    :param str end_time: EndTime.
    :param Sequence[str] ip_lists: instance IpList.
    :param str metric_name: Indicator, value [bps (attack traffic bandwidth, pps (attack packet rate)].
    :param int period: Statistical granularity, values [300 (5 minutes), 3600 (hours), 86400 (days)].
    :param str result_output_file: Used to save results.
    :param str start_time: StartTime.
    """
    __args__ = dict()
    __args__['business'] = business
    __args__['endTime'] = end_time
    __args__['ipLists'] = ip_lists
    __args__['metricName'] = metric_name
    __args__['period'] = period
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Antiddos/getOverviewDdosTrend:getOverviewDdosTrend', __args__, opts=opts, typ=GetOverviewDdosTrendResult).value

    return AwaitableGetOverviewDdosTrendResult(
        business=__ret__.business,
        datas=__ret__.datas,
        end_time=__ret__.end_time,
        id=__ret__.id,
        ip_lists=__ret__.ip_lists,
        metric_name=__ret__.metric_name,
        period=__ret__.period,
        result_output_file=__ret__.result_output_file,
        start_time=__ret__.start_time)


@_utilities.lift_output_func(get_overview_ddos_trend)
def get_overview_ddos_trend_output(business: Optional[pulumi.Input[Optional[str]]] = None,
                                   end_time: Optional[pulumi.Input[str]] = None,
                                   ip_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                   metric_name: Optional[pulumi.Input[str]] = None,
                                   period: Optional[pulumi.Input[int]] = None,
                                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   start_time: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOverviewDdosTrendResult]:
    """
    Use this data source to query detailed information of antiddos overview ddos trend

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    overview_ddos_trend = tencentcloud.Antiddos.get_overview_ddos_trend(business="bgpip",
        end_time="2023-11-21 14:16:23",
        metric_name="bps",
        period=300,
        start_time="2023-11-20 14:16:23")
    ```


    :param str business: Dayu sub product code (bgpip represents advanced defense IP; net represents professional version of advanced defense IP).
    :param str end_time: EndTime.
    :param Sequence[str] ip_lists: instance IpList.
    :param str metric_name: Indicator, value [bps (attack traffic bandwidth, pps (attack packet rate)].
    :param int period: Statistical granularity, values [300 (5 minutes), 3600 (hours), 86400 (days)].
    :param str result_output_file: Used to save results.
    :param str start_time: StartTime.
    """
    ...
