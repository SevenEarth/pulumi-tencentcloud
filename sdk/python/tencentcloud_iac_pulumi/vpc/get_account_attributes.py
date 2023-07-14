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
    'GetAccountAttributesResult',
    'AwaitableGetAccountAttributesResult',
    'get_account_attributes',
    'get_account_attributes_output',
]

@pulumi.output_type
class GetAccountAttributesResult:
    """
    A collection of values returned by getAccountAttributes.
    """
    def __init__(__self__, account_attribute_sets=None, id=None, result_output_file=None):
        if account_attribute_sets and not isinstance(account_attribute_sets, list):
            raise TypeError("Expected argument 'account_attribute_sets' to be a list")
        pulumi.set(__self__, "account_attribute_sets", account_attribute_sets)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="accountAttributeSets")
    def account_attribute_sets(self) -> Sequence['outputs.GetAccountAttributesAccountAttributeSetResult']:
        """
        User account attribute object.
        """
        return pulumi.get(self, "account_attribute_sets")

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


class AwaitableGetAccountAttributesResult(GetAccountAttributesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountAttributesResult(
            account_attribute_sets=self.account_attribute_sets,
            id=self.id,
            result_output_file=self.result_output_file)


def get_account_attributes(result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountAttributesResult:
    """
    Use this data source to query detailed information of vpc account_attributes

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    account_attributes = tencentcloud.Vpc.get_account_attributes()
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
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpc/getAccountAttributes:getAccountAttributes', __args__, opts=opts, typ=GetAccountAttributesResult).value

    return AwaitableGetAccountAttributesResult(
        account_attribute_sets=__ret__.account_attribute_sets,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_account_attributes)
def get_account_attributes_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountAttributesResult]:
    """
    Use this data source to query detailed information of vpc account_attributes

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    account_attributes = tencentcloud.Vpc.get_account_attributes()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
