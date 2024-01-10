# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetCheckProxyCreateResult',
    'AwaitableGetCheckProxyCreateResult',
    'get_check_proxy_create',
    'get_check_proxy_create_output',
]

@pulumi.output_type
class GetCheckProxyCreateResult:
    """
    A collection of values returned by getCheckProxyCreate.
    """
    def __init__(__self__, access_region=None, bandwidth=None, check_flag=None, concurrent=None, group_id=None, id=None, ip_address_version=None, network_type=None, package_type=None, real_server_region=None, result_output_file=None):
        if access_region and not isinstance(access_region, str):
            raise TypeError("Expected argument 'access_region' to be a str")
        pulumi.set(__self__, "access_region", access_region)
        if bandwidth and not isinstance(bandwidth, int):
            raise TypeError("Expected argument 'bandwidth' to be a int")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if check_flag and not isinstance(check_flag, int):
            raise TypeError("Expected argument 'check_flag' to be a int")
        pulumi.set(__self__, "check_flag", check_flag)
        if concurrent and not isinstance(concurrent, int):
            raise TypeError("Expected argument 'concurrent' to be a int")
        pulumi.set(__self__, "concurrent", concurrent)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address_version and not isinstance(ip_address_version, str):
            raise TypeError("Expected argument 'ip_address_version' to be a str")
        pulumi.set(__self__, "ip_address_version", ip_address_version)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
        if real_server_region and not isinstance(real_server_region, str):
            raise TypeError("Expected argument 'real_server_region' to be a str")
        pulumi.set(__self__, "real_server_region", real_server_region)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="accessRegion")
    def access_region(self) -> str:
        return pulumi.get(self, "access_region")

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="checkFlag")
    def check_flag(self) -> int:
        """
        Query whether the proxy with the given configuration can be created, 1 can be created, 0 cannot be created.
        """
        return pulumi.get(self, "check_flag")

    @property
    @pulumi.getter
    def concurrent(self) -> int:
        return pulumi.get(self, "concurrent")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddressVersion")
    def ip_address_version(self) -> Optional[str]:
        return pulumi.get(self, "ip_address_version")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[str]:
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> Optional[str]:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="realServerRegion")
    def real_server_region(self) -> str:
        return pulumi.get(self, "real_server_region")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetCheckProxyCreateResult(GetCheckProxyCreateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCheckProxyCreateResult(
            access_region=self.access_region,
            bandwidth=self.bandwidth,
            check_flag=self.check_flag,
            concurrent=self.concurrent,
            group_id=self.group_id,
            id=self.id,
            ip_address_version=self.ip_address_version,
            network_type=self.network_type,
            package_type=self.package_type,
            real_server_region=self.real_server_region,
            result_output_file=self.result_output_file)


def get_check_proxy_create(access_region: Optional[str] = None,
                           bandwidth: Optional[int] = None,
                           concurrent: Optional[int] = None,
                           group_id: Optional[str] = None,
                           ip_address_version: Optional[str] = None,
                           network_type: Optional[str] = None,
                           package_type: Optional[str] = None,
                           real_server_region: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCheckProxyCreateResult:
    """
    Use this data source to query detailed information of gaap check proxy create

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    check_proxy_create = tencentcloud.Gaap.get_check_proxy_create(access_region="Guangzhou",
        bandwidth=10,
        concurrent=2,
        http3_supported=0,
        ip_address_version="IPv4",
        network_type="normal",
        package_type="Thunder",
        real_server_region="Beijing")
    ```


    :param str access_region: The access (acceleration) area of the proxy. The value can be obtained through the interface DescribeAccessRegionsByDestRegion.
    :param int bandwidth: The upper limit of proxy bandwidth, in Mbps.
    :param int concurrent: The upper limit of chanproxynel concurrency, representing the number of simultaneous online connections, in tens of thousands.
    :param str group_id: If creating a proxy under a proxy group, you need to fill in the ID of the proxy group.
    :param str ip_address_version: IP version, can be taken as IPv4 or IPv6, with a default value of IPv4.
    :param str network_type: Network type, can take values &amp;#39;normal&amp;#39;, &amp;#39;cn2&amp;#39;, default value normal.
    :param str package_type: Channel package type. Thunder represents the standard proxy group, Accelerator represents the game accelerator proxy, and CrossBorder represents the cross-border proxy.
    :param str real_server_region: The origin area of the proxy. The value can be obtained through the interface DescribeDestRegions.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['accessRegion'] = access_region
    __args__['bandwidth'] = bandwidth
    __args__['concurrent'] = concurrent
    __args__['groupId'] = group_id
    __args__['ipAddressVersion'] = ip_address_version
    __args__['networkType'] = network_type
    __args__['packageType'] = package_type
    __args__['realServerRegion'] = real_server_region
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getCheckProxyCreate:getCheckProxyCreate', __args__, opts=opts, typ=GetCheckProxyCreateResult).value

    return AwaitableGetCheckProxyCreateResult(
        access_region=__ret__.access_region,
        bandwidth=__ret__.bandwidth,
        check_flag=__ret__.check_flag,
        concurrent=__ret__.concurrent,
        group_id=__ret__.group_id,
        id=__ret__.id,
        ip_address_version=__ret__.ip_address_version,
        network_type=__ret__.network_type,
        package_type=__ret__.package_type,
        real_server_region=__ret__.real_server_region,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_check_proxy_create)
def get_check_proxy_create_output(access_region: Optional[pulumi.Input[str]] = None,
                                  bandwidth: Optional[pulumi.Input[int]] = None,
                                  concurrent: Optional[pulumi.Input[int]] = None,
                                  group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  ip_address_version: Optional[pulumi.Input[Optional[str]]] = None,
                                  network_type: Optional[pulumi.Input[Optional[str]]] = None,
                                  package_type: Optional[pulumi.Input[Optional[str]]] = None,
                                  real_server_region: Optional[pulumi.Input[str]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCheckProxyCreateResult]:
    """
    Use this data source to query detailed information of gaap check proxy create

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    check_proxy_create = tencentcloud.Gaap.get_check_proxy_create(access_region="Guangzhou",
        bandwidth=10,
        concurrent=2,
        http3_supported=0,
        ip_address_version="IPv4",
        network_type="normal",
        package_type="Thunder",
        real_server_region="Beijing")
    ```


    :param str access_region: The access (acceleration) area of the proxy. The value can be obtained through the interface DescribeAccessRegionsByDestRegion.
    :param int bandwidth: The upper limit of proxy bandwidth, in Mbps.
    :param int concurrent: The upper limit of chanproxynel concurrency, representing the number of simultaneous online connections, in tens of thousands.
    :param str group_id: If creating a proxy under a proxy group, you need to fill in the ID of the proxy group.
    :param str ip_address_version: IP version, can be taken as IPv4 or IPv6, with a default value of IPv4.
    :param str network_type: Network type, can take values &amp;#39;normal&amp;#39;, &amp;#39;cn2&amp;#39;, default value normal.
    :param str package_type: Channel package type. Thunder represents the standard proxy group, Accelerator represents the game accelerator proxy, and CrossBorder represents the cross-border proxy.
    :param str real_server_region: The origin area of the proxy. The value can be obtained through the interface DescribeDestRegions.
    :param str result_output_file: Used to save results.
    """
    ...
