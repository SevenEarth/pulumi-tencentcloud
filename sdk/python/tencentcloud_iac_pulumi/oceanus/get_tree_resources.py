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
    'GetTreeResourcesResult',
    'AwaitableGetTreeResourcesResult',
    'get_tree_resources',
    'get_tree_resources_output',
]

@pulumi.output_type
class GetTreeResourcesResult:
    """
    A collection of values returned by getTreeResources.
    """
    def __init__(__self__, id=None, result_output_file=None, tree_infos=None, work_space_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if tree_infos and not isinstance(tree_infos, list):
            raise TypeError("Expected argument 'tree_infos' to be a list")
        pulumi.set(__self__, "tree_infos", tree_infos)
        if work_space_id and not isinstance(work_space_id, str):
            raise TypeError("Expected argument 'work_space_id' to be a str")
        pulumi.set(__self__, "work_space_id", work_space_id)

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
    @pulumi.getter(name="treeInfos")
    def tree_infos(self) -> Sequence['outputs.GetTreeResourcesTreeInfoResult']:
        """
        Tree structure information.
        """
        return pulumi.get(self, "tree_infos")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> str:
        return pulumi.get(self, "work_space_id")


class AwaitableGetTreeResourcesResult(GetTreeResourcesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTreeResourcesResult(
            id=self.id,
            result_output_file=self.result_output_file,
            tree_infos=self.tree_infos,
            work_space_id=self.work_space_id)


def get_tree_resources(result_output_file: Optional[str] = None,
                       work_space_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTreeResourcesResult:
    """
    Use this data source to query detailed information of oceanus tree_resources

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Oceanus.get_tree_resources(work_space_id="space-2idq8wbr")
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    :param str work_space_id: Workspace SerialId.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    __args__['workSpaceId'] = work_space_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Oceanus/getTreeResources:getTreeResources', __args__, opts=opts, typ=GetTreeResourcesResult).value

    return AwaitableGetTreeResourcesResult(
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        tree_infos=pulumi.get(__ret__, 'tree_infos'),
        work_space_id=pulumi.get(__ret__, 'work_space_id'))


@_utilities.lift_output_func(get_tree_resources)
def get_tree_resources_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              work_space_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTreeResourcesResult]:
    """
    Use this data source to query detailed information of oceanus tree_resources

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Oceanus.get_tree_resources(work_space_id="space-2idq8wbr")
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    :param str work_space_id: Workspace SerialId.
    """
    ...
