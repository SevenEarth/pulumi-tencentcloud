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
    'GetIdleInstancesResult',
    'AwaitableGetIdleInstancesResult',
    'get_idle_instances',
    'get_idle_instances_output',
]

@pulumi.output_type
class GetIdleInstancesResult:
    """
    A collection of values returned by getIdleInstances.
    """
    def __init__(__self__, id=None, idle_load_balancers=None, load_balancer_region=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idle_load_balancers and not isinstance(idle_load_balancers, list):
            raise TypeError("Expected argument 'idle_load_balancers' to be a list")
        pulumi.set(__self__, "idle_load_balancers", idle_load_balancers)
        if load_balancer_region and not isinstance(load_balancer_region, str):
            raise TypeError("Expected argument 'load_balancer_region' to be a str")
        pulumi.set(__self__, "load_balancer_region", load_balancer_region)
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
    @pulumi.getter(name="idleLoadBalancers")
    def idle_load_balancers(self) -> Sequence['outputs.GetIdleInstancesIdleLoadBalancerResult']:
        """
        List of idle CLBs. Note: This field may return null, indicating that no valid values can be obtained.
        """
        return pulumi.get(self, "idle_load_balancers")

    @property
    @pulumi.getter(name="loadBalancerRegion")
    def load_balancer_region(self) -> Optional[str]:
        return pulumi.get(self, "load_balancer_region")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetIdleInstancesResult(GetIdleInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdleInstancesResult(
            id=self.id,
            idle_load_balancers=self.idle_load_balancers,
            load_balancer_region=self.load_balancer_region,
            result_output_file=self.result_output_file)


def get_idle_instances(load_balancer_region: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdleInstancesResult:
    """
    Use this data source to query detailed information of clb idle_loadbalancers

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    idle_instance = tencentcloud.Clb.get_idle_instances(load_balancer_region="ap-guangzhou")
    ```
    <!--End PulumiCodeChooser -->


    :param str load_balancer_region: CLB instance region.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['loadBalancerRegion'] = load_balancer_region
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Clb/getIdleInstances:getIdleInstances', __args__, opts=opts, typ=GetIdleInstancesResult).value

    return AwaitableGetIdleInstancesResult(
        id=pulumi.get(__ret__, 'id'),
        idle_load_balancers=pulumi.get(__ret__, 'idle_load_balancers'),
        load_balancer_region=pulumi.get(__ret__, 'load_balancer_region'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_idle_instances)
def get_idle_instances_output(load_balancer_region: Optional[pulumi.Input[Optional[str]]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdleInstancesResult]:
    """
    Use this data source to query detailed information of clb idle_loadbalancers

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    idle_instance = tencentcloud.Clb.get_idle_instances(load_balancer_region="ap-guangzhou")
    ```
    <!--End PulumiCodeChooser -->


    :param str load_balancer_region: CLB instance region.
    :param str result_output_file: Used to save results.
    """
    ...
