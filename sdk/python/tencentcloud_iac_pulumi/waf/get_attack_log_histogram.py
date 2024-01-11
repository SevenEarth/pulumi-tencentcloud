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
    'GetAttackLogHistogramResult',
    'AwaitableGetAttackLogHistogramResult',
    'get_attack_log_histogram',
    'get_attack_log_histogram_output',
]

@pulumi.output_type
class GetAttackLogHistogramResult:
    """
    A collection of values returned by getAttackLogHistogram.
    """
    def __init__(__self__, datas=None, domain=None, end_time=None, id=None, period=None, query_string=None, result_output_file=None, start_time=None, total_count=None):
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if period and not isinstance(period, int):
            raise TypeError("Expected argument 'period' to be a int")
        pulumi.set(__self__, "period", period)
        if query_string and not isinstance(query_string, str):
            raise TypeError("Expected argument 'query_string' to be a str")
        pulumi.set(__self__, "query_string", query_string)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def datas(self) -> Sequence['outputs.GetAttackLogHistogramDataResult']:
        """
        The statistics detail.
        """
        return pulumi.get(self, "datas")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

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
    @pulumi.getter
    def period(self) -> int:
        """
        Period.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="queryString")
    def query_string(self) -> str:
        return pulumi.get(self, "query_string")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        total count.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetAttackLogHistogramResult(GetAttackLogHistogramResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAttackLogHistogramResult(
            datas=self.datas,
            domain=self.domain,
            end_time=self.end_time,
            id=self.id,
            period=self.period,
            query_string=self.query_string,
            result_output_file=self.result_output_file,
            start_time=self.start_time,
            total_count=self.total_count)


def get_attack_log_histogram(domain: Optional[str] = None,
                             end_time: Optional[str] = None,
                             query_string: Optional[str] = None,
                             result_output_file: Optional[str] = None,
                             start_time: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAttackLogHistogramResult:
    """
    Use this data source to query detailed information of waf attack_log_histogram

    ## Example Usage
    ### Obtain the specified domain name log information

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_attack_log_histogram(domain="domain.com",
        end_time="2023-09-29 00:00:00",
        query_string="method:GET",
        start_time="2023-09-01 00:00:00")
    ```
    ### Obtain all domain name log information

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_attack_log_histogram(domain="all",
        end_time="2023-09-29 00:00:00",
        query_string="method:GET",
        start_time="2023-09-01 00:00:00")
    ```


    :param str domain: Domain for query, all domain use all.
    :param str end_time: End time.
    :param str query_string: Lucene grammar.
    :param str result_output_file: Used to save results.
    :param str start_time: Begin time.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['endTime'] = end_time
    __args__['queryString'] = query_string
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Waf/getAttackLogHistogram:getAttackLogHistogram', __args__, opts=opts, typ=GetAttackLogHistogramResult).value

    return AwaitableGetAttackLogHistogramResult(
        datas=__ret__.datas,
        domain=__ret__.domain,
        end_time=__ret__.end_time,
        id=__ret__.id,
        period=__ret__.period,
        query_string=__ret__.query_string,
        result_output_file=__ret__.result_output_file,
        start_time=__ret__.start_time,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(get_attack_log_histogram)
def get_attack_log_histogram_output(domain: Optional[pulumi.Input[str]] = None,
                                    end_time: Optional[pulumi.Input[str]] = None,
                                    query_string: Optional[pulumi.Input[str]] = None,
                                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    start_time: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAttackLogHistogramResult]:
    """
    Use this data source to query detailed information of waf attack_log_histogram

    ## Example Usage
    ### Obtain the specified domain name log information

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_attack_log_histogram(domain="domain.com",
        end_time="2023-09-29 00:00:00",
        query_string="method:GET",
        start_time="2023-09-01 00:00:00")
    ```
    ### Obtain all domain name log information

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_attack_log_histogram(domain="all",
        end_time="2023-09-29 00:00:00",
        query_string="method:GET",
        start_time="2023-09-01 00:00:00")
    ```


    :param str domain: Domain for query, all domain use all.
    :param str end_time: End time.
    :param str query_string: Lucene grammar.
    :param str result_output_file: Used to save results.
    :param str start_time: Begin time.
    """
    ...