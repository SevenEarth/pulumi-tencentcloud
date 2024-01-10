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
    'GetWafInfosResult',
    'AwaitableGetWafInfosResult',
    'get_waf_infos',
    'get_waf_infos_output',
]

@pulumi.output_type
class GetWafInfosResult:
    """
    A collection of values returned by getWafInfos.
    """
    def __init__(__self__, host_lists=None, id=None, params=None, result_output_file=None):
        if host_lists and not isinstance(host_lists, list):
            raise TypeError("Expected argument 'host_lists' to be a list")
        pulumi.set(__self__, "host_lists", host_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if params and not isinstance(params, list):
            raise TypeError("Expected argument 'params' to be a list")
        pulumi.set(__self__, "params", params)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="hostLists")
    def host_lists(self) -> Sequence['outputs.GetWafInfosHostListResult']:
        return pulumi.get(self, "host_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def params(self) -> Sequence['outputs.GetWafInfosParamResult']:
        return pulumi.get(self, "params")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetWafInfosResult(GetWafInfosResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafInfosResult(
            host_lists=self.host_lists,
            id=self.id,
            params=self.params,
            result_output_file=self.result_output_file)


def get_waf_infos(params: Optional[Sequence[pulumi.InputType['GetWafInfosParamArgs']]] = None,
                  result_output_file: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafInfosResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['params'] = params
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Waf/getWafInfos:getWafInfos', __args__, opts=opts, typ=GetWafInfosResult).value

    return AwaitableGetWafInfosResult(
        host_lists=__ret__.host_lists,
        id=__ret__.id,
        params=__ret__.params,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_waf_infos)
def get_waf_infos_output(params: Optional[pulumi.Input[Sequence[pulumi.InputType['GetWafInfosParamArgs']]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWafInfosResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
