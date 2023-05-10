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
    'GetSecurityPolicyRegionsResult',
    'AwaitableGetSecurityPolicyRegionsResult',
    'get_security_policy_regions',
    'get_security_policy_regions_output',
]

@pulumi.output_type
class GetSecurityPolicyRegionsResult:
    """
    A collection of values returned by getSecurityPolicyRegions.
    """
    def __init__(__self__, geo_ips=None, id=None, result_output_file=None):
        if geo_ips and not isinstance(geo_ips, list):
            raise TypeError("Expected argument 'geo_ips' to be a list")
        pulumi.set(__self__, "geo_ips", geo_ips)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="geoIps")
    def geo_ips(self) -> Sequence['outputs.GetSecurityPolicyRegionsGeoIpResult']:
        """
        Region info.
        """
        return pulumi.get(self, "geo_ips")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetSecurityPolicyRegionsResult(GetSecurityPolicyRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityPolicyRegionsResult(
            geo_ips=self.geo_ips,
            id=self.id,
            result_output_file=self.result_output_file)


def get_security_policy_regions(result_output_file: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityPolicyRegionsResult:
    """
    Use this data source to query detailed information of teo securityPolicyRegions

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    security_policy_regions = tencentcloud.Teo.get_security_policy_regions()
    ```


    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Teo/getSecurityPolicyRegions:getSecurityPolicyRegions', __args__, opts=opts, typ=GetSecurityPolicyRegionsResult).value

    return AwaitableGetSecurityPolicyRegionsResult(
        geo_ips=__ret__.geo_ips,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_security_policy_regions)
def get_security_policy_regions_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityPolicyRegionsResult]:
    """
    Use this data source to query detailed information of teo securityPolicyRegions

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    security_policy_regions = tencentcloud.Teo.get_security_policy_regions()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
