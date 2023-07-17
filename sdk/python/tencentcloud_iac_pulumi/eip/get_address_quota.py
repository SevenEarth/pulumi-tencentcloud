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
    'GetAddressQuotaResult',
    'AwaitableGetAddressQuotaResult',
    'get_address_quota',
    'get_address_quota_output',
]

@pulumi.output_type
class GetAddressQuotaResult:
    """
    A collection of values returned by getAddressQuota.
    """
    def __init__(__self__, id=None, quota_sets=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if quota_sets and not isinstance(quota_sets, list):
            raise TypeError("Expected argument 'quota_sets' to be a list")
        pulumi.set(__self__, "quota_sets", quota_sets)
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
    @pulumi.getter(name="quotaSets")
    def quota_sets(self) -> Sequence['outputs.GetAddressQuotaQuotaSetResult']:
        """
        The specified account EIP quota information.
        """
        return pulumi.get(self, "quota_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetAddressQuotaResult(GetAddressQuotaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressQuotaResult(
            id=self.id,
            quota_sets=self.quota_sets,
            result_output_file=self.result_output_file)


def get_address_quota(result_output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressQuotaResult:
    """
    Use this data source to query detailed information of vpc address_quota

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    address_quota = tencentcloud.Eip.get_address_quota()
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
    __ret__ = pulumi.runtime.invoke('tencentcloud:Eip/getAddressQuota:getAddressQuota', __args__, opts=opts, typ=GetAddressQuotaResult).value

    return AwaitableGetAddressQuotaResult(
        id=__ret__.id,
        quota_sets=__ret__.quota_sets,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_address_quota)
def get_address_quota_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAddressQuotaResult]:
    """
    Use this data source to query detailed information of vpc address_quota

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    address_quota = tencentcloud.Eip.get_address_quota()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
