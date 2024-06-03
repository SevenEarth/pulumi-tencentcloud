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
    'GetProInstancesResult',
    'AwaitableGetProInstancesResult',
    'get_pro_instances',
    'get_pro_instances_output',
]

@pulumi.output_type
class GetProInstancesResult:
    """
    A collection of values returned by getProInstances.
    """
    def __init__(__self__, filters=None, id=None, instances=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetProInstancesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetProInstancesInstanceResult']:
        """
        Instance information list.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetProInstancesResult(GetProInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProInstancesResult(
            filters=self.filters,
            id=self.id,
            instances=self.instances,
            result_output_file=self.result_output_file)


def get_pro_instances(filters: Optional[Sequence[pulumi.InputType['GetProInstancesFilterArgs']]] = None,
                      result_output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProInstancesResult:
    """
    Use this data source to query detailed information of tdmq pro_instances

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    pro_instances_filter = tencentcloud.Tdmq.get_pro_instances(filters=[tencentcloud.tdmq.GetProInstancesFilterArgs(
        name="InstanceName",
        values=["keep"],
    )])
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[pulumi.InputType['GetProInstancesFilterArgs']] filters: query condition filter.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getProInstances:getProInstances', __args__, opts=opts, typ=GetProInstancesResult).value

    return AwaitableGetProInstancesResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        instances=pulumi.get(__ret__, 'instances'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_pro_instances)
def get_pro_instances_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetProInstancesFilterArgs']]]]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProInstancesResult]:
    """
    Use this data source to query detailed information of tdmq pro_instances

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    pro_instances_filter = tencentcloud.Tdmq.get_pro_instances(filters=[tencentcloud.tdmq.GetProInstancesFilterArgs(
        name="InstanceName",
        values=["keep"],
    )])
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[pulumi.InputType['GetProInstancesFilterArgs']] filters: query condition filter.
    :param str result_output_file: Used to save results.
    """
    ...
