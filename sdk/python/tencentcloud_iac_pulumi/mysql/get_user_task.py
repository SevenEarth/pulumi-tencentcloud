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
    'GetUserTaskResult',
    'AwaitableGetUserTaskResult',
    'get_user_task',
    'get_user_task_output',
]

@pulumi.output_type
class GetUserTaskResult:
    """
    A collection of values returned by getUserTask.
    """
    def __init__(__self__, async_request_id=None, id=None, instance_id=None, items=None, result_output_file=None, start_time_begin=None, start_time_end=None, task_statuses=None, task_types=None):
        if async_request_id and not isinstance(async_request_id, str):
            raise TypeError("Expected argument 'async_request_id' to be a str")
        pulumi.set(__self__, "async_request_id", async_request_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time_begin and not isinstance(start_time_begin, str):
            raise TypeError("Expected argument 'start_time_begin' to be a str")
        pulumi.set(__self__, "start_time_begin", start_time_begin)
        if start_time_end and not isinstance(start_time_end, str):
            raise TypeError("Expected argument 'start_time_end' to be a str")
        pulumi.set(__self__, "start_time_end", start_time_end)
        if task_statuses and not isinstance(task_statuses, list):
            raise TypeError("Expected argument 'task_statuses' to be a list")
        pulumi.set(__self__, "task_statuses", task_statuses)
        if task_types and not isinstance(task_types, list):
            raise TypeError("Expected argument 'task_types' to be a list")
        pulumi.set(__self__, "task_types", task_types)

    @property
    @pulumi.getter(name="asyncRequestId")
    def async_request_id(self) -> Optional[str]:
        """
        The request ID of the asynchronous task.
        """
        return pulumi.get(self, "async_request_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.GetUserTaskItemResult']:
        """
        The returned instance task information.
        """
        return pulumi.get(self, "items")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTimeBegin")
    def start_time_begin(self) -> Optional[str]:
        return pulumi.get(self, "start_time_begin")

    @property
    @pulumi.getter(name="startTimeEnd")
    def start_time_end(self) -> Optional[str]:
        return pulumi.get(self, "start_time_end")

    @property
    @pulumi.getter(name="taskStatuses")
    def task_statuses(self) -> Optional[Sequence[str]]:
        """
        Instance task status, possible values include:UNDEFINED - undefined;INITIAL - initialization;RUNNING - running;SUCCEED - the execution was successful;FAILED - execution failed;KILLED - terminated;REMOVED - removed;PAUSED - Paused.WAITING - waiting (cancellable).
        """
        return pulumi.get(self, "task_statuses")

    @property
    @pulumi.getter(name="taskTypes")
    def task_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "task_types")


class AwaitableGetUserTaskResult(GetUserTaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserTaskResult(
            async_request_id=self.async_request_id,
            id=self.id,
            instance_id=self.instance_id,
            items=self.items,
            result_output_file=self.result_output_file,
            start_time_begin=self.start_time_begin,
            start_time_end=self.start_time_end,
            task_statuses=self.task_statuses,
            task_types=self.task_types)


def get_user_task(async_request_id: Optional[str] = None,
                  instance_id: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  start_time_begin: Optional[str] = None,
                  start_time_end: Optional[str] = None,
                  task_statuses: Optional[Sequence[str]] = None,
                  task_types: Optional[Sequence[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserTaskResult:
    """
    Use this data source to query detailed information of mysql user_task

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    user_task = tencentcloud.Mysql.get_user_task(async_request_id="f2fe828c-773af816-0a08f542-94bb2a9c",
        instance_id="cdb-fitq5t9h",
        start_time_begin="2017-12-31 10:40:01",
        start_time_end="2017-12-31 10:40:01",
        task_statuses=["2"],
        task_types=["5"])
    ```


    :param str async_request_id: Asynchronous task request ID, the AsyncRequestId returned by executing cloud database-related operations.
    :param str instance_id: Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the value of the field InstanceId in the output parameter.
    :param str result_output_file: Used to save results.
    :param str start_time_begin: The start time of the first task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
    :param str start_time_end: The start time of the last task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
    :param Sequence[str] task_statuses: Task status. If no value is passed, all task statuses will be queried. Supported values include: `UNDEFINED` - undefined; `INITIAL` - initialization; `RUNNING` - running; `SUCCEED` - the execution was successful; `FAILED` - execution failed; `KILLED` - terminated; `REMOVED` - removed; `PAUSED` - Paused.
    :param Sequence[str] task_types: Task type. If no value is passed, all task types will be queried. Supported values include: `ROLLBACK` - database rollback; `SQL OPERATION` - SQL operation; `IMPORT DATA` - data import; `MODIFY PARAM` - parameter setting; `INITIAL` - initialize the cloud database instance; `REBOOT` - restarts the cloud database instance; `OPEN GTID` - open the cloud database instance GTID; `UPGRADE RO` - read-only instance upgrade; `BATCH ROLLBACK` - database batch rollback; `UPGRADE MASTER` - master upgrade; `DROP TABLES` - delete cloud database tables; `SWITCH DR TO MASTER` - The disaster recovery instance.
    """
    __args__ = dict()
    __args__['asyncRequestId'] = async_request_id
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['startTimeBegin'] = start_time_begin
    __args__['startTimeEnd'] = start_time_end
    __args__['taskStatuses'] = task_statuses
    __args__['taskTypes'] = task_types
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mysql/getUserTask:getUserTask', __args__, opts=opts, typ=GetUserTaskResult).value

    return AwaitableGetUserTaskResult(
        async_request_id=__ret__.async_request_id,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        items=__ret__.items,
        result_output_file=__ret__.result_output_file,
        start_time_begin=__ret__.start_time_begin,
        start_time_end=__ret__.start_time_end,
        task_statuses=__ret__.task_statuses,
        task_types=__ret__.task_types)


@_utilities.lift_output_func(get_user_task)
def get_user_task_output(async_request_id: Optional[pulumi.Input[Optional[str]]] = None,
                         instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         start_time_begin: Optional[pulumi.Input[Optional[str]]] = None,
                         start_time_end: Optional[pulumi.Input[Optional[str]]] = None,
                         task_statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         task_types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserTaskResult]:
    """
    Use this data source to query detailed information of mysql user_task

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    user_task = tencentcloud.Mysql.get_user_task(async_request_id="f2fe828c-773af816-0a08f542-94bb2a9c",
        instance_id="cdb-fitq5t9h",
        start_time_begin="2017-12-31 10:40:01",
        start_time_end="2017-12-31 10:40:01",
        task_statuses=["2"],
        task_types=["5"])
    ```


    :param str async_request_id: Asynchronous task request ID, the AsyncRequestId returned by executing cloud database-related operations.
    :param str instance_id: Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the value of the field InstanceId in the output parameter.
    :param str result_output_file: Used to save results.
    :param str start_time_begin: The start time of the first task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
    :param str start_time_end: The start time of the last task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
    :param Sequence[str] task_statuses: Task status. If no value is passed, all task statuses will be queried. Supported values include: `UNDEFINED` - undefined; `INITIAL` - initialization; `RUNNING` - running; `SUCCEED` - the execution was successful; `FAILED` - execution failed; `KILLED` - terminated; `REMOVED` - removed; `PAUSED` - Paused.
    :param Sequence[str] task_types: Task type. If no value is passed, all task types will be queried. Supported values include: `ROLLBACK` - database rollback; `SQL OPERATION` - SQL operation; `IMPORT DATA` - data import; `MODIFY PARAM` - parameter setting; `INITIAL` - initialize the cloud database instance; `REBOOT` - restarts the cloud database instance; `OPEN GTID` - open the cloud database instance GTID; `UPGRADE RO` - read-only instance upgrade; `BATCH ROLLBACK` - database batch rollback; `UPGRADE MASTER` - master upgrade; `DROP TABLES` - delete cloud database tables; `SWITCH DR TO MASTER` - The disaster recovery instance.
    """
    ...
