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
    'GetSlowlogsResult',
    'AwaitableGetSlowlogsResult',
    'get_slowlogs',
    'get_slowlogs_output',
]

@pulumi.output_type
class GetSlowlogsResult:
    """
    A collection of values returned by getSlowlogs.
    """
    def __init__(__self__, end_time=None, id=None, instance_id=None, result_output_file=None, slowlogs=None, start_time=None):
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if slowlogs and not isinstance(slowlogs, list):
            raise TypeError("Expected argument 'slowlogs' to be a list")
        pulumi.set(__self__, "slowlogs", slowlogs)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        File generation end time.
        """
        return pulumi.get(self, "end_time")

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
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def slowlogs(self) -> Sequence['outputs.GetSlowlogsSlowlogResult']:
        """
        Information list of slow query logs.
        """
        return pulumi.get(self, "slowlogs")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        File generation start time.
        """
        return pulumi.get(self, "start_time")


class AwaitableGetSlowlogsResult(GetSlowlogsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSlowlogsResult(
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            slowlogs=self.slowlogs,
            start_time=self.start_time)


def get_slowlogs(end_time: Optional[str] = None,
                 instance_id: Optional[str] = None,
                 result_output_file: Optional[str] = None,
                 start_time: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSlowlogsResult:
    """
    Use this data source to query detailed information of sqlserver slowlogs

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Sqlserver.get_slowlogs(end_time="2023-08-07 00:00:00",
        instance_id="mssql-qelbzgwf",
        start_time="2023-08-01 00:00:00")
    ```


    :param str end_time: Query end time.
    :param str instance_id: Instance ID.
    :param str result_output_file: Used to save results.
    :param str start_time: Query start time.
    """
    __args__ = dict()
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getSlowlogs:getSlowlogs', __args__, opts=opts, typ=GetSlowlogsResult).value

    return AwaitableGetSlowlogsResult(
        end_time=__ret__.end_time,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file,
        slowlogs=__ret__.slowlogs,
        start_time=__ret__.start_time)


@_utilities.lift_output_func(get_slowlogs)
def get_slowlogs_output(end_time: Optional[pulumi.Input[str]] = None,
                        instance_id: Optional[pulumi.Input[str]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        start_time: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSlowlogsResult]:
    """
    Use this data source to query detailed information of sqlserver slowlogs

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Sqlserver.get_slowlogs(end_time="2023-08-07 00:00:00",
        instance_id="mssql-qelbzgwf",
        start_time="2023-08-01 00:00:00")
    ```


    :param str end_time: Query end time.
    :param str instance_id: Instance ID.
    :param str result_output_file: Used to save results.
    :param str start_time: Query start time.
    """
    ...
