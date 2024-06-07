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
    'GetInstanceTaskListResult',
    'AwaitableGetInstanceTaskListResult',
    'get_instance_task_list',
    'get_instance_task_list_output',
]

@pulumi.output_type
class GetInstanceTaskListResult:
    """
    A collection of values returned by getInstanceTaskList.
    """
    def __init__(__self__, begin_time=None, end_time=None, id=None, instance_id=None, instance_name=None, operate_uins=None, project_ids=None, result_output_file=None, results=None, task_statuses=None, task_types=None, tasks=None):
        if begin_time and not isinstance(begin_time, str):
            raise TypeError("Expected argument 'begin_time' to be a str")
        pulumi.set(__self__, "begin_time", begin_time)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if operate_uins and not isinstance(operate_uins, list):
            raise TypeError("Expected argument 'operate_uins' to be a list")
        pulumi.set(__self__, "operate_uins", operate_uins)
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        pulumi.set(__self__, "project_ids", project_ids)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if task_statuses and not isinstance(task_statuses, list):
            raise TypeError("Expected argument 'task_statuses' to be a list")
        pulumi.set(__self__, "task_statuses", task_statuses)
        if task_types and not isinstance(task_types, list):
            raise TypeError("Expected argument 'task_types' to be a list")
        pulumi.set(__self__, "task_types", task_types)
        if tasks and not isinstance(tasks, list):
            raise TypeError("Expected argument 'tasks' to be a list")
        pulumi.set(__self__, "tasks", tasks)

    @property
    @pulumi.getter(name="beginTime")
    def begin_time(self) -> Optional[str]:
        return pulumi.get(self, "begin_time")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        """
        The end time.
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
    def instance_id(self) -> Optional[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[str]:
        """
        The name of instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="operateUins")
    def operate_uins(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "operate_uins")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "project_ids")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Optional[Sequence[int]]:
        """
        Task status.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="taskStatuses")
    def task_statuses(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "task_statuses")

    @property
    @pulumi.getter(name="taskTypes")
    def task_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "task_types")

    @property
    @pulumi.getter
    def tasks(self) -> Sequence['outputs.GetInstanceTaskListTaskResult']:
        """
        Task details.
        """
        return pulumi.get(self, "tasks")


class AwaitableGetInstanceTaskListResult(GetInstanceTaskListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceTaskListResult(
            begin_time=self.begin_time,
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            instance_name=self.instance_name,
            operate_uins=self.operate_uins,
            project_ids=self.project_ids,
            result_output_file=self.result_output_file,
            results=self.results,
            task_statuses=self.task_statuses,
            task_types=self.task_types,
            tasks=self.tasks)


def get_instance_task_list(begin_time: Optional[str] = None,
                           end_time: Optional[str] = None,
                           instance_id: Optional[str] = None,
                           instance_name: Optional[str] = None,
                           operate_uins: Optional[Sequence[str]] = None,
                           project_ids: Optional[Sequence[int]] = None,
                           result_output_file: Optional[str] = None,
                           results: Optional[Sequence[int]] = None,
                           task_statuses: Optional[Sequence[int]] = None,
                           task_types: Optional[Sequence[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceTaskListResult:
    """
    Use this data source to query detailed information of redis instance_task_list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_task_list = tencentcloud.Redis.get_instance_task_list(begin_time="2021-12-30 00:00:00",
        end_time="2021-12-30 00:00:00",
        instance_id="crs-c1nl9rpv",
        instance_name="",
        operate_uins=[""],
        project_ids=[""],
        results=[""],
        task_statuses=[""],
        task_types=[""])
    ```
    <!--End PulumiCodeChooser -->


    :param str begin_time: Start time.
    :param str end_time: Termination time.
    :param str instance_id: The ID of instance.
    :param str instance_name: Instance name.
    :param Sequence[str] operate_uins: Operator Uin.
    :param Sequence[int] project_ids: Project Id.
    :param str result_output_file: Used to save results.
    :param Sequence[int] results: Task status.
    :param Sequence[int] task_statuses: Task status.
    :param Sequence[str] task_types: Task type.
    """
    __args__ = dict()
    __args__['beginTime'] = begin_time
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['instanceName'] = instance_name
    __args__['operateUins'] = operate_uins
    __args__['projectIds'] = project_ids
    __args__['resultOutputFile'] = result_output_file
    __args__['results'] = results
    __args__['taskStatuses'] = task_statuses
    __args__['taskTypes'] = task_types
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Redis/getInstanceTaskList:getInstanceTaskList', __args__, opts=opts, typ=GetInstanceTaskListResult).value

    return AwaitableGetInstanceTaskListResult(
        begin_time=pulumi.get(__ret__, 'begin_time'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        instance_name=pulumi.get(__ret__, 'instance_name'),
        operate_uins=pulumi.get(__ret__, 'operate_uins'),
        project_ids=pulumi.get(__ret__, 'project_ids'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        results=pulumi.get(__ret__, 'results'),
        task_statuses=pulumi.get(__ret__, 'task_statuses'),
        task_types=pulumi.get(__ret__, 'task_types'),
        tasks=pulumi.get(__ret__, 'tasks'))


@_utilities.lift_output_func(get_instance_task_list)
def get_instance_task_list_output(begin_time: Optional[pulumi.Input[Optional[str]]] = None,
                                  end_time: Optional[pulumi.Input[Optional[str]]] = None,
                                  instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  instance_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  operate_uins: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  project_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  results: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  task_statuses: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                  task_types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceTaskListResult]:
    """
    Use this data source to query detailed information of redis instance_task_list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_task_list = tencentcloud.Redis.get_instance_task_list(begin_time="2021-12-30 00:00:00",
        end_time="2021-12-30 00:00:00",
        instance_id="crs-c1nl9rpv",
        instance_name="",
        operate_uins=[""],
        project_ids=[""],
        results=[""],
        task_statuses=[""],
        task_types=[""])
    ```
    <!--End PulumiCodeChooser -->


    :param str begin_time: Start time.
    :param str end_time: Termination time.
    :param str instance_id: The ID of instance.
    :param str instance_name: Instance name.
    :param Sequence[str] operate_uins: Operator Uin.
    :param Sequence[int] project_ids: Project Id.
    :param str result_output_file: Used to save results.
    :param Sequence[int] results: Task status.
    :param Sequence[int] task_statuses: Task status.
    :param Sequence[str] task_types: Task type.
    """
    ...
