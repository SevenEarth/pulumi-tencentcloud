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
    'GetWhitelistResult',
    'AwaitableGetWhitelistResult',
    'get_whitelist',
    'get_whitelist_output',
]

@pulumi.output_type
class GetWhitelistResult:
    """
    A collection of values returned by getWhitelist.
    """
    def __init__(__self__, id=None, instance_id=None, result_output_file=None, whitelist_sets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if whitelist_sets and not isinstance(whitelist_sets, list):
            raise TypeError("Expected argument 'whitelist_sets' to be a list")
        pulumi.set(__self__, "whitelist_sets", whitelist_sets)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="whitelistSets")
    def whitelist_sets(self) -> Sequence['outputs.GetWhitelistWhitelistSetResult']:
        """
        While list.
        """
        return pulumi.get(self, "whitelist_sets")


class AwaitableGetWhitelistResult(GetWhitelistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWhitelistResult(
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            whitelist_sets=self.whitelist_sets)


def get_whitelist(instance_id: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWhitelistResult:
    """
    Use this data source to query detailed information of rum whitelist

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    whitelist = tencentcloud.Rum.get_whitelist(instance_id="rum-pasZKEI3RLgakj")
    ```


    :param str instance_id: Instance ID, such as taw-123.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Rum/getWhitelist:getWhitelist', __args__, opts=opts, typ=GetWhitelistResult).value

    return AwaitableGetWhitelistResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file,
        whitelist_sets=__ret__.whitelist_sets)


@_utilities.lift_output_func(get_whitelist)
def get_whitelist_output(instance_id: Optional[pulumi.Input[str]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWhitelistResult]:
    """
    Use this data source to query detailed information of rum whitelist

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    whitelist = tencentcloud.Rum.get_whitelist(instance_id="rum-pasZKEI3RLgakj")
    ```


    :param str instance_id: Instance ID, such as taw-123.
    :param str result_output_file: Used to save results.
    """
    ...
