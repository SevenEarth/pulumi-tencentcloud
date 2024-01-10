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
    'GetStatisticsReportResult',
    'AwaitableGetStatisticsReportResult',
    'get_statistics_report',
    'get_statistics_report_output',
]

@pulumi.output_type
class GetStatisticsReportResult:
    """
    A collection of values returned by getStatisticsReport.
    """
    def __init__(__self__, daily_volumes=None, domain=None, end_date=None, id=None, overall_volumes=None, receiving_mailbox_type=None, result_output_file=None, start_date=None):
        if daily_volumes and not isinstance(daily_volumes, list):
            raise TypeError("Expected argument 'daily_volumes' to be a list")
        pulumi.set(__self__, "daily_volumes", daily_volumes)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if overall_volumes and not isinstance(overall_volumes, list):
            raise TypeError("Expected argument 'overall_volumes' to be a list")
        pulumi.set(__self__, "overall_volumes", overall_volumes)
        if receiving_mailbox_type and not isinstance(receiving_mailbox_type, str):
            raise TypeError("Expected argument 'receiving_mailbox_type' to be a str")
        pulumi.set(__self__, "receiving_mailbox_type", receiving_mailbox_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_date and not isinstance(start_date, str):
            raise TypeError("Expected argument 'start_date' to be a str")
        pulumi.set(__self__, "start_date", start_date)

    @property
    @pulumi.getter(name="dailyVolumes")
    def daily_volumes(self) -> Sequence['outputs.GetStatisticsReportDailyVolumeResult']:
        """
        Daily email sending statistics.
        """
        return pulumi.get(self, "daily_volumes")

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="overallVolumes")
    def overall_volumes(self) -> Sequence['outputs.GetStatisticsReportOverallVolumeResult']:
        """
        Overall email sending statistics.
        """
        return pulumi.get(self, "overall_volumes")

    @property
    @pulumi.getter(name="receivingMailboxType")
    def receiving_mailbox_type(self) -> Optional[str]:
        return pulumi.get(self, "receiving_mailbox_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        return pulumi.get(self, "start_date")


class AwaitableGetStatisticsReportResult(GetStatisticsReportResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStatisticsReportResult(
            daily_volumes=self.daily_volumes,
            domain=self.domain,
            end_date=self.end_date,
            id=self.id,
            overall_volumes=self.overall_volumes,
            receiving_mailbox_type=self.receiving_mailbox_type,
            result_output_file=self.result_output_file,
            start_date=self.start_date)


def get_statistics_report(domain: Optional[str] = None,
                          end_date: Optional[str] = None,
                          receiving_mailbox_type: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          start_date: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStatisticsReportResult:
    """
    Use this data source to query detailed information of ses statistics_report

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    statistics_report = tencentcloud.Ses.get_statistics_report(domain="iac-tf.cloud",
        end_date="2023-09-05",
        receiving_mailbox_type="gmail.com",
        start_date="2020-10-01")
    ```


    :param str domain: Sender domain.
    :param str end_date: End date.
    :param str receiving_mailbox_type: Recipient address type, for example, gmail.com.
    :param str result_output_file: Used to save results.
    :param str start_date: Start date.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['endDate'] = end_date
    __args__['receivingMailboxType'] = receiving_mailbox_type
    __args__['resultOutputFile'] = result_output_file
    __args__['startDate'] = start_date
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ses/getStatisticsReport:getStatisticsReport', __args__, opts=opts, typ=GetStatisticsReportResult).value

    return AwaitableGetStatisticsReportResult(
        daily_volumes=__ret__.daily_volumes,
        domain=__ret__.domain,
        end_date=__ret__.end_date,
        id=__ret__.id,
        overall_volumes=__ret__.overall_volumes,
        receiving_mailbox_type=__ret__.receiving_mailbox_type,
        result_output_file=__ret__.result_output_file,
        start_date=__ret__.start_date)


@_utilities.lift_output_func(get_statistics_report)
def get_statistics_report_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                                 end_date: Optional[pulumi.Input[str]] = None,
                                 receiving_mailbox_type: Optional[pulumi.Input[Optional[str]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 start_date: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStatisticsReportResult]:
    """
    Use this data source to query detailed information of ses statistics_report

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    statistics_report = tencentcloud.Ses.get_statistics_report(domain="iac-tf.cloud",
        end_date="2023-09-05",
        receiving_mailbox_type="gmail.com",
        start_date="2020-10-01")
    ```


    :param str domain: Sender domain.
    :param str end_date: End date.
    :param str receiving_mailbox_type: Recipient address type, for example, gmail.com.
    :param str result_output_file: Used to save results.
    :param str start_date: Start date.
    """
    ...
