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
    'GetInstanceConnectionsResult',
    'AwaitableGetInstanceConnectionsResult',
    'get_instance_connections',
    'get_instance_connections_output',
]

@pulumi.output_type
class GetInstanceConnectionsResult:
    """
    A collection of values returned by getInstanceConnections.
    """
    def __init__(__self__, clients=None, id=None, instance_id=None, result_output_file=None):
        if clients and not isinstance(clients, list):
            raise TypeError("Expected argument 'clients' to be a list")
        pulumi.set(__self__, "clients", clients)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def clients(self) -> Sequence['outputs.GetInstanceConnectionsClientResult']:
        """
        Client connection info.
        """
        return pulumi.get(self, "clients")

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


class AwaitableGetInstanceConnectionsResult(GetInstanceConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceConnectionsResult(
            clients=self.clients,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file)


def get_instance_connections(instance_id: Optional[str] = None,
                             result_output_file: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceConnectionsResult:
    """
    Use this data source to query detailed information of mongodb instance_connections

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_connections = tencentcloud.Mongodb.get_instance_connections(instance_id="cmgo-9d0p6umb")
    ```


    :param str instance_id: Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mongodb/getInstanceConnections:getInstanceConnections', __args__, opts=opts, typ=GetInstanceConnectionsResult).value

    return AwaitableGetInstanceConnectionsResult(
        clients=__ret__.clients,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_instance_connections)
def get_instance_connections_output(instance_id: Optional[pulumi.Input[str]] = None,
                                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceConnectionsResult]:
    """
    Use this data source to query detailed information of mongodb instance_connections

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_connections = tencentcloud.Mongodb.get_instance_connections(instance_id="cmgo-9d0p6umb")
    ```


    :param str instance_id: Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
    :param str result_output_file: Used to save results.
    """
    ...
