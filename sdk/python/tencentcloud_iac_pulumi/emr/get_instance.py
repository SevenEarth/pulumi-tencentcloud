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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    """
    A collection of values returned by getInstance.
    """
    def __init__(__self__, clusters=None, display_strategy=None, id=None, instance_ids=None, project_id=None, result_output_file=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if display_strategy and not isinstance(display_strategy, str):
            raise TypeError("Expected argument 'display_strategy' to be a str")
        pulumi.set(__self__, "display_strategy", display_strategy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetInstanceClusterResult']:
        """
        A list of clusters will be exported and its every element contains the following attributes:
        """
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="displayStrategy")
    def display_strategy(self) -> str:
        return pulumi.get(self, "display_strategy")

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        """
        Project id of instance.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            clusters=self.clusters,
            display_strategy=self.display_strategy,
            id=self.id,
            instance_ids=self.instance_ids,
            project_id=self.project_id,
            result_output_file=self.result_output_file)


def get_instance(display_strategy: Optional[str] = None,
                 instance_ids: Optional[Sequence[str]] = None,
                 project_id: Optional[int] = None,
                 result_output_file: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Provides an available EMR for the user.

    The EMR data source fetch proper EMR from user's EMR pool.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    my_emr = tencentcloud.Emr.get_instance(display_strategy="clusterList",
        instance_ids=["emr-rnzqrleq"])
    ```
    <!--End PulumiCodeChooser -->


    :param str display_strategy: Display strategy(e.g.:clusterList, monitorManage).
    :param Sequence[str] instance_ids: fetch all instances with same prefix(e.g.:emr-xxxxxx).
    :param int project_id: Fetch all instances which owner same project. Default 0 meaning use default project id.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['displayStrategy'] = display_strategy
    __args__['instanceIds'] = instance_ids
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Emr/getInstance:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        clusters=pulumi.get(__ret__, 'clusters'),
        display_strategy=pulumi.get(__ret__, 'display_strategy'),
        id=pulumi.get(__ret__, 'id'),
        instance_ids=pulumi.get(__ret__, 'instance_ids'),
        project_id=pulumi.get(__ret__, 'project_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_instance)
def get_instance_output(display_strategy: Optional[pulumi.Input[str]] = None,
                        instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        project_id: Optional[pulumi.Input[Optional[int]]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Provides an available EMR for the user.

    The EMR data source fetch proper EMR from user's EMR pool.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    my_emr = tencentcloud.Emr.get_instance(display_strategy="clusterList",
        instance_ids=["emr-rnzqrleq"])
    ```
    <!--End PulumiCodeChooser -->


    :param str display_strategy: Display strategy(e.g.:clusterList, monitorManage).
    :param Sequence[str] instance_ids: fetch all instances with same prefix(e.g.:emr-xxxxxx).
    :param int project_id: Fetch all instances which owner same project. Default 0 meaning use default project id.
    :param str result_output_file: Used to save results.
    """
    ...
