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
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
    'get_instances_output',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, dcx_id=None, id=None, instance_lists=None, name=None, result_output_file=None):
        if dcx_id and not isinstance(dcx_id, str):
            raise TypeError("Expected argument 'dcx_id' to be a str")
        pulumi.set(__self__, "dcx_id", dcx_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="dcxId")
    def dcx_id(self) -> Optional[str]:
        """
        ID of the dedicated tunnel.
        """
        return pulumi.get(self, "dcx_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.GetInstancesInstanceListResult']:
        """
        Information list of the dedicated tunnels.
        """
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the dedicated tunnel.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            dcx_id=self.dcx_id,
            id=self.id,
            instance_lists=self.instance_lists,
            name=self.name,
            result_output_file=self.result_output_file)


def get_instances(dcx_id: Optional[str] = None,
                  name: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    Use this data source to query detailed information of dedicated tunnels instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    name_select = tencentcloud.Dcx.get_instances(name="main")
    id = tencentcloud.Dcx.get_instances(dcx_id="dcx-3ikuw30k")
    ```


    :param str dcx_id: ID of the dedicated tunnels to be queried.
    :param str name: Name of the dedicated tunnels to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['dcxId'] = dcx_id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dcx/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        dcx_id=__ret__.dcx_id,
        id=__ret__.id,
        instance_lists=__ret__.instance_lists,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_instances)
def get_instances_output(dcx_id: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    Use this data source to query detailed information of dedicated tunnels instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    name_select = tencentcloud.Dcx.get_instances(name="main")
    id = tencentcloud.Dcx.get_instances(dcx_id="dcx-3ikuw30k")
    ```


    :param str dcx_id: ID of the dedicated tunnels to be queried.
    :param str name: Name of the dedicated tunnels to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
