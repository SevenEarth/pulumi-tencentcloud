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
    'GetPolicyConditionsResult',
    'AwaitableGetPolicyConditionsResult',
    'get_policy_conditions',
    'get_policy_conditions_output',
]

@pulumi.output_type
class GetPolicyConditionsResult:
    """
    A collection of values returned by getPolicyConditions.
    """
    def __init__(__self__, id=None, lists=None, name=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetPolicyConditionsListResult']:
        """
        A list policy condition. Each element contains the following attributes:
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of this policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetPolicyConditionsResult(GetPolicyConditionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyConditionsResult(
            id=self.id,
            lists=self.lists,
            name=self.name,
            result_output_file=self.result_output_file)


def get_policy_conditions(name: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyConditionsResult:
    """
    Use this data source to query monitor policy conditions(There is a lot of data and it is recommended to output to a file)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    monitor_policy_conditions = tencentcloud.Monitor.get_policy_conditions(name="Cloud Virtual Machine",
        result_output_file="./tencentcloud_monitor_policy_conditions.txt")
    ```


    :param str name: Name of the policy name, support partial matching, eg:`Cloud Virtual Machine`,`Virtual`,`Cloud Load Banlancer-Private CLB Listener`.
    :param str result_output_file: Used to store results.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Monitor/getPolicyConditions:getPolicyConditions', __args__, opts=opts, typ=GetPolicyConditionsResult).value

    return AwaitableGetPolicyConditionsResult(
        id=__ret__.id,
        lists=__ret__.lists,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_policy_conditions)
def get_policy_conditions_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPolicyConditionsResult]:
    """
    Use this data source to query monitor policy conditions(There is a lot of data and it is recommended to output to a file)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    monitor_policy_conditions = tencentcloud.Monitor.get_policy_conditions(name="Cloud Virtual Machine",
        result_output_file="./tencentcloud_monitor_policy_conditions.txt")
    ```


    :param str name: Name of the policy name, support partial matching, eg:`Cloud Virtual Machine`,`Virtual`,`Cloud Load Banlancer-Private CLB Listener`.
    :param str result_output_file: Used to store results.
    """
    ...
