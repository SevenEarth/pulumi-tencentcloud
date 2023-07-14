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
    'GetResourceDashboardResult',
    'AwaitableGetResourceDashboardResult',
    'get_resource_dashboard',
    'get_resource_dashboard_output',
]

@pulumi.output_type
class GetResourceDashboardResult:
    """
    A collection of values returned by getResourceDashboard.
    """
    def __init__(__self__, id=None, resource_dashboard_sets=None, result_output_file=None, vpc_ids=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_dashboard_sets and not isinstance(resource_dashboard_sets, list):
            raise TypeError("Expected argument 'resource_dashboard_sets' to be a list")
        pulumi.set(__self__, "resource_dashboard_sets", resource_dashboard_sets)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if vpc_ids and not isinstance(vpc_ids, list):
            raise TypeError("Expected argument 'vpc_ids' to be a list")
        pulumi.set(__self__, "vpc_ids", vpc_ids)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceDashboardSets")
    def resource_dashboard_sets(self) -> Sequence['outputs.GetResourceDashboardResourceDashboardSetResult']:
        """
        List of resource objects.
        """
        return pulumi.get(self, "resource_dashboard_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> Sequence[str]:
        return pulumi.get(self, "vpc_ids")


class AwaitableGetResourceDashboardResult(GetResourceDashboardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceDashboardResult(
            id=self.id,
            resource_dashboard_sets=self.resource_dashboard_sets,
            result_output_file=self.result_output_file,
            vpc_ids=self.vpc_ids)


def get_resource_dashboard(result_output_file: Optional[str] = None,
                           vpc_ids: Optional[Sequence[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceDashboardResult:
    """
    Use this data source to query detailed information of vpc resource_dashboard

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    resource_dashboard = tencentcloud.Vpc.get_resource_dashboard(vpc_ids=["vpc-4owdpnwr"])
    ```


    :param str result_output_file: Used to save results.
    :param Sequence[str] vpc_ids: Vpc instance ID, e.g. vpc-f1xjkw1b.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    __args__['vpcIds'] = vpc_ids
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpc/getResourceDashboard:getResourceDashboard', __args__, opts=opts, typ=GetResourceDashboardResult).value

    return AwaitableGetResourceDashboardResult(
        id=__ret__.id,
        resource_dashboard_sets=__ret__.resource_dashboard_sets,
        result_output_file=__ret__.result_output_file,
        vpc_ids=__ret__.vpc_ids)


@_utilities.lift_output_func(get_resource_dashboard)
def get_resource_dashboard_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  vpc_ids: Optional[pulumi.Input[Sequence[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceDashboardResult]:
    """
    Use this data source to query detailed information of vpc resource_dashboard

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    resource_dashboard = tencentcloud.Vpc.get_resource_dashboard(vpc_ids=["vpc-4owdpnwr"])
    ```


    :param str result_output_file: Used to save results.
    :param Sequence[str] vpc_ids: Vpc instance ID, e.g. vpc-f1xjkw1b.
    """
    ...
