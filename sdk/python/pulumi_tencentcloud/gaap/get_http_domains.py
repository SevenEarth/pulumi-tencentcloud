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
    'GetHttpDomainsResult',
    'AwaitableGetHttpDomainsResult',
    'get_http_domains',
    'get_http_domains_output',
]

@pulumi.output_type
class GetHttpDomainsResult:
    """
    A collection of values returned by getHttpDomains.
    """
    def __init__(__self__, domain=None, domains=None, id=None, listener_id=None, result_output_file=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Forward domain of the layer7 listener.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetHttpDomainsDomainResult']:
        """
        An information list of forward domain of the layer7 listeners. Each element contains the following attributes:
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetHttpDomainsResult(GetHttpDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHttpDomainsResult(
            domain=self.domain,
            domains=self.domains,
            id=self.id,
            listener_id=self.listener_id,
            result_output_file=self.result_output_file)


def get_http_domains(domain: Optional[str] = None,
                     listener_id: Optional[str] = None,
                     result_output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHttpDomainsResult:
    """
    Use this data source to query forward domain of layer7 listeners.

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
    foo_http_domain = tencentcloud.gaap.HttpDomain("fooHttpDomain",
        listener_id=foo_layer7_listener.id,
        domain="www.qq.com")
    foo_http_domains = tencentcloud.Gaap.get_http_domains_output(listener_id=foo_layer7_listener.id,
        domain=foo_http_domain.domain)
    ```


    :param str domain: Forward domain of the layer7 listener to be queried.
    :param str listener_id: ID of the layer7 listener to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['listenerId'] = listener_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getHttpDomains:getHttpDomains', __args__, opts=opts, typ=GetHttpDomainsResult).value

    return AwaitableGetHttpDomainsResult(
        domain=__ret__.domain,
        domains=__ret__.domains,
        id=__ret__.id,
        listener_id=__ret__.listener_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_http_domains)
def get_http_domains_output(domain: Optional[pulumi.Input[str]] = None,
                            listener_id: Optional[pulumi.Input[str]] = None,
                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHttpDomainsResult]:
    """
    Use this data source to query forward domain of layer7 listeners.

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
    foo_http_domain = tencentcloud.gaap.HttpDomain("fooHttpDomain",
        listener_id=foo_layer7_listener.id,
        domain="www.qq.com")
    foo_http_domains = tencentcloud.Gaap.get_http_domains_output(listener_id=foo_layer7_listener.id,
        domain=foo_http_domain.domain)
    ```


    :param str domain: Forward domain of the layer7 listener to be queried.
    :param str listener_id: ID of the layer7 listener to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
