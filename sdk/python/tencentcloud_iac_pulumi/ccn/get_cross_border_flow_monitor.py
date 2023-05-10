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
    'GetCrossBorderFlowMonitorResult',
    'AwaitableGetCrossBorderFlowMonitorResult',
    'get_cross_border_flow_monitor',
    'get_cross_border_flow_monitor_output',
]

@pulumi.output_type
class GetCrossBorderFlowMonitorResult:
    """
    A collection of values returned by getCrossBorderFlowMonitor.
    """
    def __init__(__self__, ccn_id=None, ccn_uin=None, cross_border_flow_monitor_datas=None, destination_region=None, end_time=None, id=None, period=None, result_output_file=None, source_region=None, start_time=None):
        if ccn_id and not isinstance(ccn_id, str):
            raise TypeError("Expected argument 'ccn_id' to be a str")
        pulumi.set(__self__, "ccn_id", ccn_id)
        if ccn_uin and not isinstance(ccn_uin, str):
            raise TypeError("Expected argument 'ccn_uin' to be a str")
        pulumi.set(__self__, "ccn_uin", ccn_uin)
        if cross_border_flow_monitor_datas and not isinstance(cross_border_flow_monitor_datas, list):
            raise TypeError("Expected argument 'cross_border_flow_monitor_datas' to be a list")
        pulumi.set(__self__, "cross_border_flow_monitor_datas", cross_border_flow_monitor_datas)
        if destination_region and not isinstance(destination_region, str):
            raise TypeError("Expected argument 'destination_region' to be a str")
        pulumi.set(__self__, "destination_region", destination_region)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if period and not isinstance(period, int):
            raise TypeError("Expected argument 'period' to be a int")
        pulumi.set(__self__, "period", period)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if source_region and not isinstance(source_region, str):
            raise TypeError("Expected argument 'source_region' to be a str")
        pulumi.set(__self__, "source_region", source_region)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> str:
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="ccnUin")
    def ccn_uin(self) -> str:
        return pulumi.get(self, "ccn_uin")

    @property
    @pulumi.getter(name="crossBorderFlowMonitorDatas")
    def cross_border_flow_monitor_datas(self) -> Sequence['outputs.GetCrossBorderFlowMonitorCrossBorderFlowMonitorDataResult']:
        """
        monitor data of cross border.
        """
        return pulumi.get(self, "cross_border_flow_monitor_datas")

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> str:
        return pulumi.get(self, "destination_region")

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
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="sourceRegion")
    def source_region(self) -> str:
        return pulumi.get(self, "source_region")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        return pulumi.get(self, "start_time")


class AwaitableGetCrossBorderFlowMonitorResult(GetCrossBorderFlowMonitorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCrossBorderFlowMonitorResult(
            ccn_id=self.ccn_id,
            ccn_uin=self.ccn_uin,
            cross_border_flow_monitor_datas=self.cross_border_flow_monitor_datas,
            destination_region=self.destination_region,
            end_time=self.end_time,
            id=self.id,
            period=self.period,
            result_output_file=self.result_output_file,
            source_region=self.source_region,
            start_time=self.start_time)


def get_cross_border_flow_monitor(ccn_id: Optional[str] = None,
                                  ccn_uin: Optional[str] = None,
                                  destination_region: Optional[str] = None,
                                  end_time: Optional[str] = None,
                                  period: Optional[int] = None,
                                  result_output_file: Optional[str] = None,
                                  source_region: Optional[str] = None,
                                  start_time: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCrossBorderFlowMonitorResult:
    """
    Use this data source to query detailed information of vpc cross_border_flow_monitor

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cross_border_flow_monitor = tencentcloud.Ccn.get_cross_border_flow_monitor(ccn_id="ccn-39lqkygf",
        ccn_uin="979137",
        destination_region="ap-singapore",
        end_time="2023-01-01 01:00:00",
        period=60,
        source_region="ap-guangzhou",
        start_time="2023-01-01 00:00:00")
    ```


    :param str ccn_id: CcnId.
    :param str ccn_uin: CcnUin.
    :param str destination_region: DestinationRegion.
    :param str end_time: EndTime.
    :param int period: TimePeriod.
    :param str result_output_file: Used to save results.
    :param str source_region: SourceRegion.
    :param str start_time: StartTime.
    """
    __args__ = dict()
    __args__['ccnId'] = ccn_id
    __args__['ccnUin'] = ccn_uin
    __args__['destinationRegion'] = destination_region
    __args__['endTime'] = end_time
    __args__['period'] = period
    __args__['resultOutputFile'] = result_output_file
    __args__['sourceRegion'] = source_region
    __args__['startTime'] = start_time
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ccn/getCrossBorderFlowMonitor:getCrossBorderFlowMonitor', __args__, opts=opts, typ=GetCrossBorderFlowMonitorResult).value

    return AwaitableGetCrossBorderFlowMonitorResult(
        ccn_id=__ret__.ccn_id,
        ccn_uin=__ret__.ccn_uin,
        cross_border_flow_monitor_datas=__ret__.cross_border_flow_monitor_datas,
        destination_region=__ret__.destination_region,
        end_time=__ret__.end_time,
        id=__ret__.id,
        period=__ret__.period,
        result_output_file=__ret__.result_output_file,
        source_region=__ret__.source_region,
        start_time=__ret__.start_time)


@_utilities.lift_output_func(get_cross_border_flow_monitor)
def get_cross_border_flow_monitor_output(ccn_id: Optional[pulumi.Input[str]] = None,
                                         ccn_uin: Optional[pulumi.Input[str]] = None,
                                         destination_region: Optional[pulumi.Input[str]] = None,
                                         end_time: Optional[pulumi.Input[str]] = None,
                                         period: Optional[pulumi.Input[int]] = None,
                                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                         source_region: Optional[pulumi.Input[str]] = None,
                                         start_time: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCrossBorderFlowMonitorResult]:
    """
    Use this data source to query detailed information of vpc cross_border_flow_monitor

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    cross_border_flow_monitor = tencentcloud.Ccn.get_cross_border_flow_monitor(ccn_id="ccn-39lqkygf",
        ccn_uin="979137",
        destination_region="ap-singapore",
        end_time="2023-01-01 01:00:00",
        period=60,
        source_region="ap-guangzhou",
        start_time="2023-01-01 00:00:00")
    ```


    :param str ccn_id: CcnId.
    :param str ccn_uin: CcnUin.
    :param str destination_region: DestinationRegion.
    :param str end_time: EndTime.
    :param int period: TimePeriod.
    :param str result_output_file: Used to save results.
    :param str source_region: SourceRegion.
    :param str start_time: StartTime.
    """
    ...
