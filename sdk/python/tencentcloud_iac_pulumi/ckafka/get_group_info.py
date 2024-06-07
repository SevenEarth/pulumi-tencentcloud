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
    'GetGroupInfoResult',
    'AwaitableGetGroupInfoResult',
    'get_group_info',
    'get_group_info_output',
]

@pulumi.output_type
class GetGroupInfoResult:
    """
    A collection of values returned by getGroupInfo.
    """
    def __init__(__self__, group_lists=None, id=None, instance_id=None, result_output_file=None, results=None):
        if group_lists and not isinstance(group_lists, list):
            raise TypeError("Expected argument 'group_lists' to be a list")
        pulumi.set(__self__, "group_lists", group_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="groupLists")
    def group_lists(self) -> Sequence[str]:
        return pulumi.get(self, "group_lists")

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
    def results(self) -> Sequence['outputs.GetGroupInfoResultResult']:
        """
        result.
        """
        return pulumi.get(self, "results")


class AwaitableGetGroupInfoResult(GetGroupInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupInfoResult(
            group_lists=self.group_lists,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            results=self.results)


def get_group_info(group_lists: Optional[Sequence[str]] = None,
                   instance_id: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupInfoResult:
    """
    Use this data source to query detailed information of ckafka group_info

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    group_info = tencentcloud.Ckafka.get_group_info(group_lists=["xxxxxx"],
        instance_id="ckafka-xxxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] group_lists: Kafka consumption group, Consumer-group, here is an array format, format GroupList.0=xxx&amp;amp;GroupList.1=yyy.
    :param str instance_id: InstanceId.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['groupLists'] = group_lists
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getGroupInfo:getGroupInfo', __args__, opts=opts, typ=GetGroupInfoResult).value

    return AwaitableGetGroupInfoResult(
        group_lists=pulumi.get(__ret__, 'group_lists'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        results=pulumi.get(__ret__, 'results'))


@_utilities.lift_output_func(get_group_info)
def get_group_info_output(group_lists: Optional[pulumi.Input[Sequence[str]]] = None,
                          instance_id: Optional[pulumi.Input[str]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupInfoResult]:
    """
    Use this data source to query detailed information of ckafka group_info

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    group_info = tencentcloud.Ckafka.get_group_info(group_lists=["xxxxxx"],
        instance_id="ckafka-xxxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] group_lists: Kafka consumption group, Consumer-group, here is an array format, format GroupList.0=xxx&amp;amp;GroupList.1=yyy.
    :param str instance_id: InstanceId.
    :param str result_output_file: Used to save results.
    """
    ...
