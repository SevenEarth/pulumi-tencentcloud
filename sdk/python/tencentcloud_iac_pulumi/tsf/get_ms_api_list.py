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
    'GetMsApiListResult',
    'AwaitableGetMsApiListResult',
    'get_ms_api_list',
    'get_ms_api_list_output',
]

@pulumi.output_type
class GetMsApiListResult:
    """
    A collection of values returned by getMsApiList.
    """
    def __init__(__self__, id=None, microservice_id=None, result_output_file=None, results=None, search_word=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if microservice_id and not isinstance(microservice_id, str):
            raise TypeError("Expected argument 'microservice_id' to be a str")
        pulumi.set(__self__, "microservice_id", microservice_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if search_word and not isinstance(search_word, str):
            raise TypeError("Expected argument 'search_word' to be a str")
        pulumi.set(__self__, "search_word", search_word)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="microserviceId")
    def microservice_id(self) -> str:
        return pulumi.get(self, "microservice_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetMsApiListResultResult']:
        """
        result list.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="searchWord")
    def search_word(self) -> Optional[str]:
        return pulumi.get(self, "search_word")


class AwaitableGetMsApiListResult(GetMsApiListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMsApiListResult(
            id=self.id,
            microservice_id=self.microservice_id,
            result_output_file=self.result_output_file,
            results=self.results,
            search_word=self.search_word)


def get_ms_api_list(microservice_id: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    search_word: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMsApiListResult:
    """
    Use this data source to query detailed information of tsf ms_api_list

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    ms_api_list = tencentcloud.Tsf.get_ms_api_list(microservice_id="ms-yq3jo6jd",
        search_word="echo")
    ```


    :param str microservice_id: Microservice Id.
    :param str result_output_file: Used to save results.
    :param str search_word: search word, support  service name.
    """
    __args__ = dict()
    __args__['microserviceId'] = microservice_id
    __args__['resultOutputFile'] = result_output_file
    __args__['searchWord'] = search_word
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tsf/getMsApiList:getMsApiList', __args__, opts=opts, typ=GetMsApiListResult).value

    return AwaitableGetMsApiListResult(
        id=__ret__.id,
        microservice_id=__ret__.microservice_id,
        result_output_file=__ret__.result_output_file,
        results=__ret__.results,
        search_word=__ret__.search_word)


@_utilities.lift_output_func(get_ms_api_list)
def get_ms_api_list_output(microservice_id: Optional[pulumi.Input[str]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           search_word: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMsApiListResult]:
    """
    Use this data source to query detailed information of tsf ms_api_list

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    ms_api_list = tencentcloud.Tsf.get_ms_api_list(microservice_id="ms-yq3jo6jd",
        search_word="echo")
    ```


    :param str microservice_id: Microservice Id.
    :param str result_output_file: Used to save results.
    :param str search_word: search word, support  service name.
    """
    ...
