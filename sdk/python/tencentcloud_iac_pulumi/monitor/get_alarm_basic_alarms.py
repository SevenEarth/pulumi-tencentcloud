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
    'GetAlarmBasicAlarmsResult',
    'AwaitableGetAlarmBasicAlarmsResult',
    'get_alarm_basic_alarms',
    'get_alarm_basic_alarms_output',
]

@pulumi.output_type
class GetAlarmBasicAlarmsResult:
    """
    A collection of values returned by getAlarmBasicAlarms.
    """
    def __init__(__self__, alarm_statuses=None, alarms=None, end_time=None, id=None, instance_group_ids=None, metric_names=None, module=None, obj_like=None, occur_time_order=None, project_ids=None, result_output_file=None, start_time=None, view_names=None, warning=None):
        if alarm_statuses and not isinstance(alarm_statuses, list):
            raise TypeError("Expected argument 'alarm_statuses' to be a list")
        pulumi.set(__self__, "alarm_statuses", alarm_statuses)
        if alarms and not isinstance(alarms, list):
            raise TypeError("Expected argument 'alarms' to be a list")
        pulumi.set(__self__, "alarms", alarms)
        if end_time and not isinstance(end_time, int):
            raise TypeError("Expected argument 'end_time' to be a int")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_group_ids and not isinstance(instance_group_ids, list):
            raise TypeError("Expected argument 'instance_group_ids' to be a list")
        pulumi.set(__self__, "instance_group_ids", instance_group_ids)
        if metric_names and not isinstance(metric_names, list):
            raise TypeError("Expected argument 'metric_names' to be a list")
        pulumi.set(__self__, "metric_names", metric_names)
        if module and not isinstance(module, str):
            raise TypeError("Expected argument 'module' to be a str")
        pulumi.set(__self__, "module", module)
        if obj_like and not isinstance(obj_like, str):
            raise TypeError("Expected argument 'obj_like' to be a str")
        pulumi.set(__self__, "obj_like", obj_like)
        if occur_time_order and not isinstance(occur_time_order, str):
            raise TypeError("Expected argument 'occur_time_order' to be a str")
        pulumi.set(__self__, "occur_time_order", occur_time_order)
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        pulumi.set(__self__, "project_ids", project_ids)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time and not isinstance(start_time, int):
            raise TypeError("Expected argument 'start_time' to be a int")
        pulumi.set(__self__, "start_time", start_time)
        if view_names and not isinstance(view_names, list):
            raise TypeError("Expected argument 'view_names' to be a list")
        pulumi.set(__self__, "view_names", view_names)
        if warning and not isinstance(warning, str):
            raise TypeError("Expected argument 'warning' to be a str")
        pulumi.set(__self__, "warning", warning)

    @property
    @pulumi.getter(name="alarmStatuses")
    def alarm_statuses(self) -> Optional[Sequence[int]]:
        """
        Alarm status, ALARM indicates not recovered; OK indicates that it has been restored; NO_ DATA indicates insufficient data; NO_ CONF indicates that it has expired.
        """
        return pulumi.get(self, "alarm_statuses")

    @property
    @pulumi.getter
    def alarms(self) -> Sequence['outputs.GetAlarmBasicAlarmsAlarmResult']:
        """
        Alarm List.
        """
        return pulumi.get(self, "alarms")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[int]:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceGroupIds")
    def instance_group_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "instance_group_ids")

    @property
    @pulumi.getter(name="metricNames")
    def metric_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "metric_names")

    @property
    @pulumi.getter
    def module(self) -> str:
        return pulumi.get(self, "module")

    @property
    @pulumi.getter(name="objLike")
    def obj_like(self) -> Optional[str]:
        return pulumi.get(self, "obj_like")

    @property
    @pulumi.getter(name="occurTimeOrder")
    def occur_time_order(self) -> Optional[str]:
        return pulumi.get(self, "occur_time_order")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "project_ids")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[int]:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="viewNames")
    def view_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "view_names")

    @property
    @pulumi.getter
    def warning(self) -> str:
        """
        Remarks.
        """
        return pulumi.get(self, "warning")


class AwaitableGetAlarmBasicAlarmsResult(GetAlarmBasicAlarmsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlarmBasicAlarmsResult(
            alarm_statuses=self.alarm_statuses,
            alarms=self.alarms,
            end_time=self.end_time,
            id=self.id,
            instance_group_ids=self.instance_group_ids,
            metric_names=self.metric_names,
            module=self.module,
            obj_like=self.obj_like,
            occur_time_order=self.occur_time_order,
            project_ids=self.project_ids,
            result_output_file=self.result_output_file,
            start_time=self.start_time,
            view_names=self.view_names,
            warning=self.warning)


def get_alarm_basic_alarms(alarm_statuses: Optional[Sequence[int]] = None,
                           end_time: Optional[int] = None,
                           instance_group_ids: Optional[Sequence[int]] = None,
                           metric_names: Optional[Sequence[str]] = None,
                           module: Optional[str] = None,
                           obj_like: Optional[str] = None,
                           occur_time_order: Optional[str] = None,
                           project_ids: Optional[Sequence[int]] = None,
                           result_output_file: Optional[str] = None,
                           start_time: Optional[int] = None,
                           view_names: Optional[Sequence[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlarmBasicAlarmsResult:
    """
    Use this data source to query detailed information of monitor basic_alarms

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    alarms = tencentcloud.Monitor.get_alarm_basic_alarms(alarm_statuses=[1],
        end_time=1697098903,
        instance_group_ids=[5497073],
        metric_names=["cpu_usage"],
        module="monitor",
        occur_time_order="DESC",
        project_ids=[0],
        start_time=1696990903,
        view_names=["cvm_device"])
    ```


    :param Sequence[int] alarm_statuses: Filter based on alarm status.
    :param int end_time: End time, default to current timestamp.
    :param Sequence[int] instance_group_ids: Filter based on instance group ID.
    :param Sequence[str] metric_names: Filter by indicator name.
    :param str module: Interface module name, current value monitor.
    :param str obj_like: Filter based on alarm objects.
    :param str occur_time_order: Sort by occurrence time, taking ASC or DESC values.
    :param Sequence[int] project_ids: Filter based on project ID.
    :param str result_output_file: Used to save results.
    :param int start_time: Start time, default to one day is timestamp.
    :param Sequence[str] view_names: Filter based on policy type.
    """
    __args__ = dict()
    __args__['alarmStatuses'] = alarm_statuses
    __args__['endTime'] = end_time
    __args__['instanceGroupIds'] = instance_group_ids
    __args__['metricNames'] = metric_names
    __args__['module'] = module
    __args__['objLike'] = obj_like
    __args__['occurTimeOrder'] = occur_time_order
    __args__['projectIds'] = project_ids
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    __args__['viewNames'] = view_names
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Monitor/getAlarmBasicAlarms:getAlarmBasicAlarms', __args__, opts=opts, typ=GetAlarmBasicAlarmsResult).value

    return AwaitableGetAlarmBasicAlarmsResult(
        alarm_statuses=__ret__.alarm_statuses,
        alarms=__ret__.alarms,
        end_time=__ret__.end_time,
        id=__ret__.id,
        instance_group_ids=__ret__.instance_group_ids,
        metric_names=__ret__.metric_names,
        module=__ret__.module,
        obj_like=__ret__.obj_like,
        occur_time_order=__ret__.occur_time_order,
        project_ids=__ret__.project_ids,
        result_output_file=__ret__.result_output_file,
        start_time=__ret__.start_time,
        view_names=__ret__.view_names,
        warning=__ret__.warning)


@_utilities.lift_output_func(get_alarm_basic_alarms)
def get_alarm_basic_alarms_output(alarm_statuses: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  end_time: Optional[pulumi.Input[Optional[int]]] = None,
                                  instance_group_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  metric_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  module: Optional[pulumi.Input[str]] = None,
                                  obj_like: Optional[pulumi.Input[Optional[str]]] = None,
                                  occur_time_order: Optional[pulumi.Input[Optional[str]]] = None,
                                  project_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  start_time: Optional[pulumi.Input[Optional[int]]] = None,
                                  view_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAlarmBasicAlarmsResult]:
    """
    Use this data source to query detailed information of monitor basic_alarms

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    alarms = tencentcloud.Monitor.get_alarm_basic_alarms(alarm_statuses=[1],
        end_time=1697098903,
        instance_group_ids=[5497073],
        metric_names=["cpu_usage"],
        module="monitor",
        occur_time_order="DESC",
        project_ids=[0],
        start_time=1696990903,
        view_names=["cvm_device"])
    ```


    :param Sequence[int] alarm_statuses: Filter based on alarm status.
    :param int end_time: End time, default to current timestamp.
    :param Sequence[int] instance_group_ids: Filter based on instance group ID.
    :param Sequence[str] metric_names: Filter by indicator name.
    :param str module: Interface module name, current value monitor.
    :param str obj_like: Filter based on alarm objects.
    :param str occur_time_order: Sort by occurrence time, taking ASC or DESC values.
    :param Sequence[int] project_ids: Filter based on project ID.
    :param str result_output_file: Used to save results.
    :param int start_time: Start time, default to one day is timestamp.
    :param Sequence[str] view_names: Filter based on policy type.
    """
    ...
