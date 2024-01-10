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
    'GetPlatformEventNamesResult',
    'AwaitableGetPlatformEventNamesResult',
    'get_platform_event_names',
    'get_platform_event_names_output',
]

@pulumi.output_type
class GetPlatformEventNamesResult:
    """
    A collection of values returned by getPlatformEventNames.
    """
    def __init__(__self__, event_names=None, id=None, product_type=None, result_output_file=None):
        if event_names and not isinstance(event_names, list):
            raise TypeError("Expected argument 'event_names' to be a list")
        pulumi.set(__self__, "event_names", event_names)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if product_type and not isinstance(product_type, str):
            raise TypeError("Expected argument 'product_type' to be a str")
        pulumi.set(__self__, "product_type", product_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="eventNames")
    def event_names(self) -> Sequence['outputs.GetPlatformEventNamesEventNameResult']:
        """
        Platform product list.
        """
        return pulumi.get(self, "event_names")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> str:
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetPlatformEventNamesResult(GetPlatformEventNamesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlatformEventNamesResult(
            event_names=self.event_names,
            id=self.id,
            product_type=self.product_type,
            result_output_file=self.result_output_file)


def get_platform_event_names(product_type: Optional[str] = None,
                             result_output_file: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlatformEventNamesResult:
    """
    Use this data source to query detailed information of eb platform_event_names

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    platform_event_names = tencentcloud.Eb.get_platform_event_names(product_type="")
    ```


    :param str product_type: Platform product event type.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['productType'] = product_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Eb/getPlatformEventNames:getPlatformEventNames', __args__, opts=opts, typ=GetPlatformEventNamesResult).value

    return AwaitableGetPlatformEventNamesResult(
        event_names=__ret__.event_names,
        id=__ret__.id,
        product_type=__ret__.product_type,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_platform_event_names)
def get_platform_event_names_output(product_type: Optional[pulumi.Input[str]] = None,
                                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlatformEventNamesResult]:
    """
    Use this data source to query detailed information of eb platform_event_names

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    platform_event_names = tencentcloud.Eb.get_platform_event_names(product_type="")
    ```


    :param str product_type: Platform product event type.
    :param str result_output_file: Used to save results.
    """
    ...
