# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetProductsResult',
    'AwaitableGetProductsResult',
    'get_products',
    'get_products_output',
]

@pulumi.output_type
class GetProductsResult:
    """
    A collection of values returned by getProducts.
    """
    def __init__(__self__, id=None, products=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if products and not isinstance(products, list):
            raise TypeError("Expected argument 'products' to be a list")
        pulumi.set(__self__, "products", products)
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
    @pulumi.getter
    def products(self) -> Sequence[str]:
        """
        List of supported services.
        """
        return pulumi.get(self, "products")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetProductsResult(GetProductsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductsResult(
            id=self.id,
            products=self.products,
            result_output_file=self.result_output_file)


def get_products(result_output_file: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductsResult:
    """
    Use this data source to query detailed information of ssm products

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    products = tencentcloud.Ssm.get_products()
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
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ssm/getProducts:getProducts', __args__, opts=opts, typ=GetProductsResult).value

    return AwaitableGetProductsResult(
        id=__ret__.id,
        products=__ret__.products,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_products)
def get_products_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProductsResult]:
    """
    Use this data source to query detailed information of ssm products

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    products = tencentcloud.Ssm.get_products()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
