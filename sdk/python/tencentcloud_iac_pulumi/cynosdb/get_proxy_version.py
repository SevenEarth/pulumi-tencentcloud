# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetProxyVersionResult',
    'AwaitableGetProxyVersionResult',
    'get_proxy_version',
    'get_proxy_version_output',
]

@pulumi.output_type
class GetProxyVersionResult:
    """
    A collection of values returned by getProxyVersion.
    """
    def __init__(__self__, cluster_id=None, current_proxy_version=None, id=None, proxy_group_id=None, result_output_file=None, support_proxy_versions=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if current_proxy_version and not isinstance(current_proxy_version, str):
            raise TypeError("Expected argument 'current_proxy_version' to be a str")
        pulumi.set(__self__, "current_proxy_version", current_proxy_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if proxy_group_id and not isinstance(proxy_group_id, str):
            raise TypeError("Expected argument 'proxy_group_id' to be a str")
        pulumi.set(__self__, "proxy_group_id", proxy_group_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if support_proxy_versions and not isinstance(support_proxy_versions, list):
            raise TypeError("Expected argument 'support_proxy_versions' to be a list")
        pulumi.set(__self__, "support_proxy_versions", support_proxy_versions)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="currentProxyVersion")
    def current_proxy_version(self) -> str:
        """
        Current proxy version number note: This field may return null, indicating that a valid value cannot be obtained.
        """
        return pulumi.get(self, "current_proxy_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="proxyGroupId")
    def proxy_group_id(self) -> Optional[str]:
        return pulumi.get(self, "proxy_group_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="supportProxyVersions")
    def support_proxy_versions(self) -> Sequence[str]:
        """
        Supported Database Agent Version Collection Note: This field may return null, indicating that a valid value cannot be obtained.
        """
        return pulumi.get(self, "support_proxy_versions")


class AwaitableGetProxyVersionResult(GetProxyVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProxyVersionResult(
            cluster_id=self.cluster_id,
            current_proxy_version=self.current_proxy_version,
            id=self.id,
            proxy_group_id=self.proxy_group_id,
            result_output_file=self.result_output_file,
            support_proxy_versions=self.support_proxy_versions)


def get_proxy_version(cluster_id: Optional[str] = None,
                      proxy_group_id: Optional[str] = None,
                      result_output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProxyVersionResult:
    """
    Use this data source to query detailed information of cynosdb proxy_version

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    proxy_version = tencentcloud.Cynosdb.get_proxy_version(cluster_id="cynosdbmysql-bws8h88b",
        proxy_group_id="cynosdbmysql-proxy-l6zf9t30")
    ```


    :param str cluster_id: Cluster ID.
    :param str proxy_group_id: Database Agent Group ID.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['proxyGroupId'] = proxy_group_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cynosdb/getProxyVersion:getProxyVersion', __args__, opts=opts, typ=GetProxyVersionResult).value

    return AwaitableGetProxyVersionResult(
        cluster_id=__ret__.cluster_id,
        current_proxy_version=__ret__.current_proxy_version,
        id=__ret__.id,
        proxy_group_id=__ret__.proxy_group_id,
        result_output_file=__ret__.result_output_file,
        support_proxy_versions=__ret__.support_proxy_versions)


@_utilities.lift_output_func(get_proxy_version)
def get_proxy_version_output(cluster_id: Optional[pulumi.Input[str]] = None,
                             proxy_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProxyVersionResult]:
    """
    Use this data source to query detailed information of cynosdb proxy_version

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    proxy_version = tencentcloud.Cynosdb.get_proxy_version(cluster_id="cynosdbmysql-bws8h88b",
        proxy_group_id="cynosdbmysql-proxy-l6zf9t30")
    ```


    :param str cluster_id: Cluster ID.
    :param str proxy_group_id: Database Agent Group ID.
    :param str result_output_file: Used to save results.
    """
    ...
