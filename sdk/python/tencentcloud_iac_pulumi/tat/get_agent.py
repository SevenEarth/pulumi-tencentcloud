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
    'GetAgentResult',
    'AwaitableGetAgentResult',
    'get_agent',
    'get_agent_output',
]

@pulumi.output_type
class GetAgentResult:
    """
    A collection of values returned by getAgent.
    """
    def __init__(__self__, automation_agent_sets=None, filters=None, id=None, instance_ids=None, result_output_file=None):
        if automation_agent_sets and not isinstance(automation_agent_sets, list):
            raise TypeError("Expected argument 'automation_agent_sets' to be a list")
        pulumi.set(__self__, "automation_agent_sets", automation_agent_sets)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="automationAgentSets")
    def automation_agent_sets(self) -> Sequence['outputs.GetAgentAutomationAgentSetResult']:
        """
        List of agent message.
        """
        return pulumi.get(self, "automation_agent_sets")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetAgentFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetAgentResult(GetAgentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgentResult(
            automation_agent_sets=self.automation_agent_sets,
            filters=self.filters,
            id=self.id,
            instance_ids=self.instance_ids,
            result_output_file=self.result_output_file)


def get_agent(filters: Optional[Sequence[pulumi.InputType['GetAgentFilterArgs']]] = None,
              instance_ids: Optional[Sequence[str]] = None,
              result_output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgentResult:
    """
    Use this data source to query detailed information of tat agent

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    agent = tencentcloud.Tat.get_agent(filters=[tencentcloud.tat.GetAgentFilterArgs(
        name="environment",
        values=["Linux"],
    )])
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[pulumi.InputType['GetAgentFilterArgs']] filters: Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
    :param Sequence[str] instance_ids: List of instance IDs for the query.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['instanceIds'] = instance_ids
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tat/getAgent:getAgent', __args__, opts=opts, typ=GetAgentResult).value

    return AwaitableGetAgentResult(
        automation_agent_sets=pulumi.get(__ret__, 'automation_agent_sets'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        instance_ids=pulumi.get(__ret__, 'instance_ids'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_agent)
def get_agent_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetAgentFilterArgs']]]]] = None,
                     instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAgentResult]:
    """
    Use this data source to query detailed information of tat agent

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    agent = tencentcloud.Tat.get_agent(filters=[tencentcloud.tat.GetAgentFilterArgs(
        name="environment",
        values=["Linux"],
    )])
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[pulumi.InputType['GetAgentFilterArgs']] filters: Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
    :param Sequence[str] instance_ids: List of instance IDs for the query.
    :param str result_output_file: Used to save results.
    """
    ...
