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
    'GetHttpRulesResult',
    'AwaitableGetHttpRulesResult',
    'get_http_rules',
    'get_http_rules_output',
]

@pulumi.output_type
class GetHttpRulesResult:
    """
    A collection of values returned by getHttpRules.
    """
    def __init__(__self__, domain=None, forward_host=None, id=None, listener_id=None, path=None, result_output_file=None, rules=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if forward_host and not isinstance(forward_host, str):
            raise TypeError("Expected argument 'forward_host' to be a str")
        pulumi.set(__self__, "forward_host", forward_host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        Domain of the GAAP realserver.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="forwardHost")
    def forward_host(self) -> Optional[str]:
        """
        Requested host which is forwarded to the realserver by the listener.
        """
        return pulumi.get(self, "forward_host")

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
        """
        ID of the layer7 listener.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        Path of the forward rule.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetHttpRulesRuleResult']:
        """
        An information list of forward rule of the layer7 listeners. Each element contains the following attributes:
        """
        return pulumi.get(self, "rules")


class AwaitableGetHttpRulesResult(GetHttpRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHttpRulesResult(
            domain=self.domain,
            forward_host=self.forward_host,
            id=self.id,
            listener_id=self.listener_id,
            path=self.path,
            result_output_file=self.result_output_file,
            rules=self.rules)


def get_http_rules(domain: Optional[str] = None,
                   forward_host: Optional[str] = None,
                   listener_id: Optional[str] = None,
                   path: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHttpRulesResult:
    """
    Use this data source to query forward rule of layer7 listeners.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
        bandwidth=10,
        concurrent=2,
        access_region="SouthChina",
        realserver_region="NorthChina")
    foo_layer7_listener = tencentcloud.gaap.Layer7Listener("fooLayer7Listener",
        protocol="HTTP",
        port=80,
        proxy_id=foo_proxy.id)
    foo_realserver = tencentcloud.gaap.Realserver("fooRealserver", ip="1.1.1.1")
    foo_http_rule = tencentcloud.gaap.HttpRule("fooHttpRule",
        listener_id=foo_layer7_listener.id,
        domain="www.qq.com",
        path="/",
        realserver_type="IP",
        health_check=True,
        realservers=[tencentcloud.gaap.HttpRuleRealserverArgs(
            id=foo_realserver.id,
            ip=foo_realserver.ip,
            port=80,
        )])
    foo_http_rules = tencentcloud.Gaap.get_http_rules_output(listener_id=foo_layer7_listener.id,
        domain=foo_http_rule.domain)
    ```


    :param str domain: Forward domain of the layer7 listener to be queried.
    :param str forward_host: Requested host which is forwarded to the realserver by the listener to be queried.
    :param str listener_id: ID of the layer7 listener to be queried.
    :param str path: Path of the forward rule to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['forwardHost'] = forward_host
    __args__['listenerId'] = listener_id
    __args__['path'] = path
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getHttpRules:getHttpRules', __args__, opts=opts, typ=GetHttpRulesResult).value

    return AwaitableGetHttpRulesResult(
        domain=__ret__.domain,
        forward_host=__ret__.forward_host,
        id=__ret__.id,
        listener_id=__ret__.listener_id,
        path=__ret__.path,
        result_output_file=__ret__.result_output_file,
        rules=__ret__.rules)


@_utilities.lift_output_func(get_http_rules)
def get_http_rules_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                          forward_host: Optional[pulumi.Input[Optional[str]]] = None,
                          listener_id: Optional[pulumi.Input[str]] = None,
                          path: Optional[pulumi.Input[Optional[str]]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHttpRulesResult]:
    """
    Use this data source to query forward rule of layer7 listeners.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo_proxy = tencentcloud.gaap.Proxy("fooProxy",
        bandwidth=10,
        concurrent=2,
        access_region="SouthChina",
        realserver_region="NorthChina")
    foo_layer7_listener = tencentcloud.gaap.Layer7Listener("fooLayer7Listener",
        protocol="HTTP",
        port=80,
        proxy_id=foo_proxy.id)
    foo_realserver = tencentcloud.gaap.Realserver("fooRealserver", ip="1.1.1.1")
    foo_http_rule = tencentcloud.gaap.HttpRule("fooHttpRule",
        listener_id=foo_layer7_listener.id,
        domain="www.qq.com",
        path="/",
        realserver_type="IP",
        health_check=True,
        realservers=[tencentcloud.gaap.HttpRuleRealserverArgs(
            id=foo_realserver.id,
            ip=foo_realserver.ip,
            port=80,
        )])
    foo_http_rules = tencentcloud.Gaap.get_http_rules_output(listener_id=foo_layer7_listener.id,
        domain=foo_http_rule.domain)
    ```


    :param str domain: Forward domain of the layer7 listener to be queried.
    :param str forward_host: Requested host which is forwarded to the realserver by the listener to be queried.
    :param str listener_id: ID of the layer7 listener to be queried.
    :param str path: Path of the forward rule to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
