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
    'GetLayer7ListenersResult',
    'AwaitableGetLayer7ListenersResult',
    'get_layer7_listeners',
    'get_layer7_listeners_output',
]

@pulumi.output_type
class GetLayer7ListenersResult:
    """
    A collection of values returned by getLayer7Listeners.
    """
    def __init__(__self__, id=None, listener_id=None, listener_name=None, listeners=None, port=None, protocol=None, proxy_id=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if listener_name and not isinstance(listener_name, str):
            raise TypeError("Expected argument 'listener_name' to be a str")
        pulumi.set(__self__, "listener_name", listener_name)
        if listeners and not isinstance(listeners, list):
            raise TypeError("Expected argument 'listeners' to be a list")
        pulumi.set(__self__, "listeners", listeners)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if proxy_id and not isinstance(proxy_id, str):
            raise TypeError("Expected argument 'proxy_id' to be a str")
        pulumi.set(__self__, "proxy_id", proxy_id)
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
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="listenerName")
    def listener_name(self) -> Optional[str]:
        return pulumi.get(self, "listener_name")

    @property
    @pulumi.getter
    def listeners(self) -> Sequence['outputs.GetLayer7ListenersListenerResult']:
        """
        An information list of layer7 listeners. Each element contains the following attributes:
        """
        return pulumi.get(self, "listeners")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        Port of the layer7 listener.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Protocol of the layer7 listener.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> Optional[str]:
        return pulumi.get(self, "proxy_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetLayer7ListenersResult(GetLayer7ListenersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLayer7ListenersResult(
            id=self.id,
            listener_id=self.listener_id,
            listener_name=self.listener_name,
            listeners=self.listeners,
            port=self.port,
            protocol=self.protocol,
            proxy_id=self.proxy_id,
            result_output_file=self.result_output_file)


def get_layer7_listeners(listener_id: Optional[str] = None,
                         listener_name: Optional[str] = None,
                         port: Optional[int] = None,
                         protocol: Optional[str] = None,
                         proxy_id: Optional[str] = None,
                         result_output_file: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLayer7ListenersResult:
    """
    Use this data source to query gaap layer7 listeners.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
        bandwidth=10,
        concurrent=2,
        access_region="SouthChina",
        realserver_region="NorthChina")
    foo_layer7_listener = tencentcloud.gaap.Layer7Listener("fooLayer7Listener",
        protocol="HTTP",
        port=80,
        proxy_id=foo_proxy.id)
    listener_id = tencentcloud.Gaap.get_layer7_listeners_output(protocol="HTTP",
        proxy_id=foo_proxy.id,
        listener_id=foo_layer7_listener.id)
    ```


    :param str listener_id: ID of the layer7 listener to be queried.
    :param str listener_name: Name of the layer7 listener to be queried.
    :param int port: Port of the layer7 listener to be queried.
    :param str protocol: Protocol of the layer7 listener to be queried. Valid values: `HTTP` and `HTTPS`.
    :param str proxy_id: ID of the GAAP proxy to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['listenerId'] = listener_id
    __args__['listenerName'] = listener_name
    __args__['port'] = port
    __args__['protocol'] = protocol
    __args__['proxyId'] = proxy_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getLayer7Listeners:getLayer7Listeners', __args__, opts=opts, typ=GetLayer7ListenersResult).value

    return AwaitableGetLayer7ListenersResult(
        id=__ret__.id,
        listener_id=__ret__.listener_id,
        listener_name=__ret__.listener_name,
        listeners=__ret__.listeners,
        port=__ret__.port,
        protocol=__ret__.protocol,
        proxy_id=__ret__.proxy_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_layer7_listeners)
def get_layer7_listeners_output(listener_id: Optional[pulumi.Input[Optional[str]]] = None,
                                listener_name: Optional[pulumi.Input[Optional[str]]] = None,
                                port: Optional[pulumi.Input[Optional[int]]] = None,
                                protocol: Optional[pulumi.Input[str]] = None,
                                proxy_id: Optional[pulumi.Input[Optional[str]]] = None,
                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLayer7ListenersResult]:
    """
    Use this data source to query gaap layer7 listeners.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
        bandwidth=10,
        concurrent=2,
        access_region="SouthChina",
        realserver_region="NorthChina")
    foo_layer7_listener = tencentcloud.gaap.Layer7Listener("fooLayer7Listener",
        protocol="HTTP",
        port=80,
        proxy_id=foo_proxy.id)
    listener_id = tencentcloud.Gaap.get_layer7_listeners_output(protocol="HTTP",
        proxy_id=foo_proxy.id,
        listener_id=foo_layer7_listener.id)
    ```


    :param str listener_id: ID of the layer7 listener to be queried.
    :param str listener_name: Name of the layer7 listener to be queried.
    :param int port: Port of the layer7 listener to be queried.
    :param str protocol: Protocol of the layer7 listener to be queried. Valid values: `HTTP` and `HTTPS`.
    :param str proxy_id: ID of the GAAP proxy to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
