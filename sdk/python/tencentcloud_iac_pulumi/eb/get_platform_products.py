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
    'GetPlatformProductsResult',
    'AwaitableGetPlatformProductsResult',
    'get_platform_products',
    'get_platform_products_output',
]

@pulumi.output_type
class GetPlatformProductsResult:
    """
    A collection of values returned by getPlatformProducts.
    """
    def __init__(__self__, id=None, platform_products=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if platform_products and not isinstance(platform_products, list):
            raise TypeError("Expected argument 'platform_products' to be a list")
        pulumi.set(__self__, "platform_products", platform_products)
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
    @pulumi.getter(name="platformProducts")
    def platform_products(self) -> Sequence['outputs.GetPlatformProductsPlatformProductResult']:
        """
        Platform product list.
        """
        return pulumi.get(self, "platform_products")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetPlatformProductsResult(GetPlatformProductsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlatformProductsResult(
            id=self.id,
            platform_products=self.platform_products,
            result_output_file=self.result_output_file)


def get_platform_products(result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlatformProductsResult:
    """
    Use this data source to query detailed information of eb platform_products

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    platform_products = tencentcloud.Eb.get_platform_products()
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
    __ret__ = pulumi.runtime.invoke('tencentcloud:Eb/getPlatformProducts:getPlatformProducts', __args__, opts=opts, typ=GetPlatformProductsResult).value

    return AwaitableGetPlatformProductsResult(
        id=__ret__.id,
        platform_products=__ret__.platform_products,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_platform_products)
def get_platform_products_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlatformProductsResult]:
    """
    Use this data source to query detailed information of eb platform_products

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    platform_products = tencentcloud.Eb.get_platform_products()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
