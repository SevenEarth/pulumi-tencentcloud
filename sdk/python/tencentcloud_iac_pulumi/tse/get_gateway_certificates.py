# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetGatewayCertificatesResult',
    'AwaitableGetGatewayCertificatesResult',
    'get_gateway_certificates',
    'get_gateway_certificates_output',
]

@pulumi.output_type
class GetGatewayCertificatesResult:
    """
    A collection of values returned by getGatewayCertificates.
    """
    def __init__(__self__, filters=None, gateway_id=None, id=None, result_output_file=None, results=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        pulumi.set(__self__, "gateway_id", gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetGatewayCertificatesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> str:
        return pulumi.get(self, "gateway_id")

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

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetGatewayCertificatesResultResult']:
        """
        Result.
        """
        return pulumi.get(self, "results")


class AwaitableGetGatewayCertificatesResult(GetGatewayCertificatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayCertificatesResult(
            filters=self.filters,
            gateway_id=self.gateway_id,
            id=self.id,
            result_output_file=self.result_output_file,
            results=self.results)


def get_gateway_certificates(filters: Optional[Sequence[pulumi.InputType['GetGatewayCertificatesFilterArgs']]] = None,
                             gateway_id: Optional[str] = None,
                             result_output_file: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayCertificatesResult:
    """
    Use this data source to query detailed information of tse gateway_certificates

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    gateway_certificates = tencentcloud.Tse.get_gateway_certificates(filters=[tencentcloud.tse.GetGatewayCertificatesFilterArgs(
            key="BindDomain",
            value="example.com",
        )],
        gateway_id="gateway-ddbb709b")
    ```


    :param Sequence[pulumi.InputType['GetGatewayCertificatesFilterArgs']] filters: Filter conditions, valid value: `BindDomain`, `Name`.
    :param str gateway_id: Gateway ID.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['gatewayId'] = gateway_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tse/getGatewayCertificates:getGatewayCertificates', __args__, opts=opts, typ=GetGatewayCertificatesResult).value

    return AwaitableGetGatewayCertificatesResult(
        filters=__ret__.filters,
        gateway_id=__ret__.gateway_id,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        results=__ret__.results)


@_utilities.lift_output_func(get_gateway_certificates)
def get_gateway_certificates_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetGatewayCertificatesFilterArgs']]]]] = None,
                                    gateway_id: Optional[pulumi.Input[str]] = None,
                                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayCertificatesResult]:
    """
    Use this data source to query detailed information of tse gateway_certificates

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    gateway_certificates = tencentcloud.Tse.get_gateway_certificates(filters=[tencentcloud.tse.GetGatewayCertificatesFilterArgs(
            key="BindDomain",
            value="example.com",
        )],
        gateway_id="gateway-ddbb709b")
    ```


    :param Sequence[pulumi.InputType['GetGatewayCertificatesFilterArgs']] filters: Filter conditions, valid value: `BindDomain`, `Name`.
    :param str gateway_id: Gateway ID.
    :param str result_output_file: Used to save results.
    """
    ...
