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
    'GetTaskStatusResult',
    'AwaitableGetTaskStatusResult',
    'get_task_status',
    'get_task_status_output',
]

@pulumi.output_type
class GetTaskStatusResult:
    """
    A collection of values returned by getTaskStatus.
    """
    def __init__(__self__, flow_id=None, id=None, result_output_file=None, results=None):
        if flow_id and not isinstance(flow_id, int):
            raise TypeError("Expected argument 'flow_id' to be a int")
        pulumi.set(__self__, "flow_id", flow_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="flowId")
    def flow_id(self) -> int:
        return pulumi.get(self, "flow_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetTaskStatusResultResult']:
        """
        Result.
        """
        return pulumi.get(self, "results")


class AwaitableGetTaskStatusResult(GetTaskStatusResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTaskStatusResult(
            flow_id=self.flow_id,
            id=self.id,
            result_output_file=self.result_output_file,
            results=self.results)


def get_task_status(flow_id: Optional[int] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTaskStatusResult:
    """
    Use this data source to query detailed information of ckafka task_status

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    task_status = tencentcloud.Ckafka.get_task_status(flow_id=123456)
    ```
    <!--End PulumiCodeChooser -->


    :param int flow_id: FlowId.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['flowId'] = flow_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getTaskStatus:getTaskStatus', __args__, opts=opts, typ=GetTaskStatusResult).value

    return AwaitableGetTaskStatusResult(
        flow_id=pulumi.get(__ret__, 'flow_id'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        results=pulumi.get(__ret__, 'results'))


@_utilities.lift_output_func(get_task_status)
def get_task_status_output(flow_id: Optional[pulumi.Input[int]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTaskStatusResult]:
    """
    Use this data source to query detailed information of ckafka task_status

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    task_status = tencentcloud.Ckafka.get_task_status(flow_id=123456)
    ```
    <!--End PulumiCodeChooser -->


    :param int flow_id: FlowId.
    :param str result_output_file: Used to save results.
    """
    ...
