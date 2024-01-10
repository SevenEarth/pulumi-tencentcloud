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
    'GetDescHaLogResult',
    'AwaitableGetDescHaLogResult',
    'get_desc_ha_log',
    'get_desc_ha_log_output',
]

@pulumi.output_type
class GetDescHaLogResult:
    """
    A collection of values returned by getDescHaLog.
    """
    def __init__(__self__, end_time=None, id=None, instance_id=None, result_output_file=None, start_time=None, switch_logs=None, switch_type=None):
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
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if switch_logs and not isinstance(switch_logs, list):
            raise TypeError("Expected argument 'switch_logs' to be a list")
        pulumi.set(__self__, "switch_logs", switch_logs)
        if switch_type and not isinstance(switch_type, int):
            raise TypeError("Expected argument 'switch_type' to be a int")
        pulumi.set(__self__, "switch_type", switch_type)

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="switchLogs")
    def switch_logs(self) -> Sequence['outputs.GetDescHaLogSwitchLogResult']:
        return pulumi.get(self, "switch_logs")

    @property
    @pulumi.getter(name="switchType")
    def switch_type(self) -> Optional[int]:
        return pulumi.get(self, "switch_type")


class AwaitableGetDescHaLogResult(GetDescHaLogResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescHaLogResult(
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            start_time=self.start_time,
            switch_logs=self.switch_logs,
            switch_type=self.switch_type)


def get_desc_ha_log(end_time: Optional[str] = None,
                    instance_id: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    start_time: Optional[str] = None,
                    switch_type: Optional[int] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescHaLogResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    __args__['switchType'] = switch_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getDescHaLog:getDescHaLog', __args__, opts=opts, typ=GetDescHaLogResult).value

    return AwaitableGetDescHaLogResult(
        end_time=__ret__.end_time,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file,
        start_time=__ret__.start_time,
        switch_logs=__ret__.switch_logs,
        switch_type=__ret__.switch_type)


@_utilities.lift_output_func(get_desc_ha_log)
def get_desc_ha_log_output(end_time: Optional[pulumi.Input[str]] = None,
                           instance_id: Optional[pulumi.Input[str]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           start_time: Optional[pulumi.Input[str]] = None,
                           switch_type: Optional[pulumi.Input[Optional[int]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescHaLogResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
