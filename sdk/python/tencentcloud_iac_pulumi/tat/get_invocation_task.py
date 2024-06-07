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
from ._inputs import *

__all__ = [
    'GetInvocationTaskResult',
    'AwaitableGetInvocationTaskResult',
    'get_invocation_task',
    'get_invocation_task_output',
]

@pulumi.output_type
class GetInvocationTaskResult:
    """
    A collection of values returned by getInvocationTask.
    """
    def __init__(__self__, filters=None, hide_output=None, id=None, invocation_task_ids=None, invocation_task_sets=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if hide_output and not isinstance(hide_output, bool):
            raise TypeError("Expected argument 'hide_output' to be a bool")
        pulumi.set(__self__, "hide_output", hide_output)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invocation_task_ids and not isinstance(invocation_task_ids, list):
            raise TypeError("Expected argument 'invocation_task_ids' to be a list")
        pulumi.set(__self__, "invocation_task_ids", invocation_task_ids)
        if invocation_task_sets and not isinstance(invocation_task_sets, list):
            raise TypeError("Expected argument 'invocation_task_sets' to be a list")
        pulumi.set(__self__, "invocation_task_sets", invocation_task_sets)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetInvocationTaskFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="hideOutput")
    def hide_output(self) -> Optional[bool]:
        return pulumi.get(self, "hide_output")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="invocationTaskIds")
    def invocation_task_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "invocation_task_ids")

    @property
    @pulumi.getter(name="invocationTaskSets")
    def invocation_task_sets(self) -> Sequence['outputs.GetInvocationTaskInvocationTaskSetResult']:
        """
        List of execution tasks.
        """
        return pulumi.get(self, "invocation_task_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInvocationTaskResult(GetInvocationTaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInvocationTaskResult(
            filters=self.filters,
            hide_output=self.hide_output,
            id=self.id,
            invocation_task_ids=self.invocation_task_ids,
            invocation_task_sets=self.invocation_task_sets,
            result_output_file=self.result_output_file)


def get_invocation_task(filters: Optional[Sequence[pulumi.InputType['GetInvocationTaskFilterArgs']]] = None,
                        hide_output: Optional[bool] = None,
                        invocation_task_ids: Optional[Sequence[str]] = None,
                        result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInvocationTaskResult:
    """
    Use this data source to query detailed information of tat invocation_task

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    invocation_task = tencentcloud.Tat.get_invocation_task(filters=[tencentcloud.tat.GetInvocationTaskFilterArgs(
            name="instance-id",
            values=["ins-p4pq4gaq"],
        )],
        hide_output=True)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[pulumi.InputType['GetInvocationTaskFilterArgs']] filters: Filter conditions.invocation-id - String - Required: No - (Filter condition) Filter by the execution activity ID.invocation-task-id - String - Required: No - (Filter condition) Filter by the execution task ID.instance-id - String - Required: No - (Filter condition) Filter by the instance ID.command-id - String - Required: No - (Filter condition) Filter by the command ID.Up to 10 Filters are allowed for each request. Each filter can have up to five Filter.Values. InvocationTaskIds and Filters cannot be specified at the same time.
    :param bool hide_output: Whether to hide the output. Valid values:True (default): Hide the outputFalse: Show the output.
    :param Sequence[str] invocation_task_ids: List of execution task IDs. Up to 100 IDs are allowed for each request. InvocationTaskIds and Filters cannot be specified at the same time.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['hideOutput'] = hide_output
    __args__['invocationTaskIds'] = invocation_task_ids
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tat/getInvocationTask:getInvocationTask', __args__, opts=opts, typ=GetInvocationTaskResult).value

    return AwaitableGetInvocationTaskResult(
        filters=pulumi.get(__ret__, 'filters'),
        hide_output=pulumi.get(__ret__, 'hide_output'),
        id=pulumi.get(__ret__, 'id'),
        invocation_task_ids=pulumi.get(__ret__, 'invocation_task_ids'),
        invocation_task_sets=pulumi.get(__ret__, 'invocation_task_sets'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_invocation_task)
def get_invocation_task_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetInvocationTaskFilterArgs']]]]] = None,
                               hide_output: Optional[pulumi.Input[Optional[bool]]] = None,
                               invocation_task_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInvocationTaskResult]:
    """
    Use this data source to query detailed information of tat invocation_task

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    invocation_task = tencentcloud.Tat.get_invocation_task(filters=[tencentcloud.tat.GetInvocationTaskFilterArgs(
            name="instance-id",
            values=["ins-p4pq4gaq"],
        )],
        hide_output=True)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[pulumi.InputType['GetInvocationTaskFilterArgs']] filters: Filter conditions.invocation-id - String - Required: No - (Filter condition) Filter by the execution activity ID.invocation-task-id - String - Required: No - (Filter condition) Filter by the execution task ID.instance-id - String - Required: No - (Filter condition) Filter by the instance ID.command-id - String - Required: No - (Filter condition) Filter by the command ID.Up to 10 Filters are allowed for each request. Each filter can have up to five Filter.Values. InvocationTaskIds and Filters cannot be specified at the same time.
    :param bool hide_output: Whether to hide the output. Valid values:True (default): Hide the outputFalse: Show the output.
    :param Sequence[str] invocation_task_ids: List of execution task IDs. Up to 100 IDs are allowed for each request. InvocationTaskIds and Filters cannot be specified at the same time.
    :param str result_output_file: Used to save results.
    """
    ...
