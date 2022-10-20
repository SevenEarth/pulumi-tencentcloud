# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSecurityPoliciesResult',
    'AwaitableGetSecurityPoliciesResult',
    'get_security_policies',
    'get_security_policies_output',
]

@pulumi.output_type
class GetSecurityPoliciesResult:
    """
    A collection of values returned by getSecurityPolicies.
    """
    def __init__(__self__, action=None, id=None, proxy_id=None, result_output_file=None, status=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if proxy_id and not isinstance(proxy_id, str):
            raise TypeError("Expected argument 'proxy_id' to be a str")
        pulumi.set(__self__, "proxy_id", proxy_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Default policy.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> str:
        """
        ID of the GAAP proxy.
        """
        return pulumi.get(self, "proxy_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the security policy.
        """
        return pulumi.get(self, "status")


class AwaitableGetSecurityPoliciesResult(GetSecurityPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityPoliciesResult(
            action=self.action,
            id=self.id,
            proxy_id=self.proxy_id,
            result_output_file=self.result_output_file,
            status=self.status)


def get_security_policies(id: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityPoliciesResult:
    """
    Use this data source to query security policies of GAAP proxy.

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
    foo_security_policy = tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy",
        proxy_id=foo_proxy.id,
        action="ACCEPT")
    foo_security_policies = tencentcloud.Gaap.get_security_policies_output(id=foo_security_policy.id)
    ```


    :param str id: ID of the security policy to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getSecurityPolicies:getSecurityPolicies', __args__, opts=opts, typ=GetSecurityPoliciesResult).value

    return AwaitableGetSecurityPoliciesResult(
        action=__ret__.action,
        id=__ret__.id,
        proxy_id=__ret__.proxy_id,
        result_output_file=__ret__.result_output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_security_policies)
def get_security_policies_output(id: Optional[pulumi.Input[str]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityPoliciesResult]:
    """
    Use this data source to query security policies of GAAP proxy.

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
    foo_security_policy = tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy",
        proxy_id=foo_proxy.id,
        action="ACCEPT")
    foo_security_policies = tencentcloud.Gaap.get_security_policies_output(id=foo_security_policy.id)
    ```


    :param str id: ID of the security policy to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
