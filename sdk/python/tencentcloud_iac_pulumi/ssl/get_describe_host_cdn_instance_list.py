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
    'GetDescribeHostCdnInstanceListResult',
    'AwaitableGetDescribeHostCdnInstanceListResult',
    'get_describe_host_cdn_instance_list',
    'get_describe_host_cdn_instance_list_output',
]

@pulumi.output_type
class GetDescribeHostCdnInstanceListResult:
    """
    A collection of values returned by getDescribeHostCdnInstanceList.
    """
    def __init__(__self__, async_cache=None, async_cache_time=None, async_offset=None, async_total_num=None, certificate_id=None, filters=None, id=None, instance_lists=None, is_cache=None, old_certificate_id=None, resource_type=None, result_output_file=None):
        if async_cache and not isinstance(async_cache, int):
            raise TypeError("Expected argument 'async_cache' to be a int")
        pulumi.set(__self__, "async_cache", async_cache)
        if async_cache_time and not isinstance(async_cache_time, str):
            raise TypeError("Expected argument 'async_cache_time' to be a str")
        pulumi.set(__self__, "async_cache_time", async_cache_time)
        if async_offset and not isinstance(async_offset, int):
            raise TypeError("Expected argument 'async_offset' to be a int")
        pulumi.set(__self__, "async_offset", async_offset)
        if async_total_num and not isinstance(async_total_num, int):
            raise TypeError("Expected argument 'async_total_num' to be a int")
        pulumi.set(__self__, "async_total_num", async_total_num)
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if is_cache and not isinstance(is_cache, int):
            raise TypeError("Expected argument 'is_cache' to be a int")
        pulumi.set(__self__, "is_cache", is_cache)
        if old_certificate_id and not isinstance(old_certificate_id, str):
            raise TypeError("Expected argument 'old_certificate_id' to be a str")
        pulumi.set(__self__, "old_certificate_id", old_certificate_id)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="asyncCache")
    def async_cache(self) -> Optional[int]:
        return pulumi.get(self, "async_cache")

    @property
    @pulumi.getter(name="asyncCacheTime")
    def async_cache_time(self) -> str:
        """
        Current cache read timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "async_cache_time")

    @property
    @pulumi.getter(name="asyncOffset")
    def async_offset(self) -> int:
        """
        Asynchronous refresh current execution numberNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "async_offset")

    @property
    @pulumi.getter(name="asyncTotalNum")
    def async_total_num(self) -> int:
        """
        The total number of asynchronous refreshNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "async_total_num")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDescribeHostCdnInstanceListFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.GetDescribeHostCdnInstanceListInstanceListResult']:
        """
        CDN instance listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter(name="isCache")
    def is_cache(self) -> Optional[int]:
        return pulumi.get(self, "is_cache")

    @property
    @pulumi.getter(name="oldCertificateId")
    def old_certificate_id(self) -> Optional[str]:
        return pulumi.get(self, "old_certificate_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDescribeHostCdnInstanceListResult(GetDescribeHostCdnInstanceListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeHostCdnInstanceListResult(
            async_cache=self.async_cache,
            async_cache_time=self.async_cache_time,
            async_offset=self.async_offset,
            async_total_num=self.async_total_num,
            certificate_id=self.certificate_id,
            filters=self.filters,
            id=self.id,
            instance_lists=self.instance_lists,
            is_cache=self.is_cache,
            old_certificate_id=self.old_certificate_id,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file)


def get_describe_host_cdn_instance_list(async_cache: Optional[int] = None,
                                        certificate_id: Optional[str] = None,
                                        filters: Optional[Sequence[pulumi.InputType['GetDescribeHostCdnInstanceListFilterArgs']]] = None,
                                        is_cache: Optional[int] = None,
                                        old_certificate_id: Optional[str] = None,
                                        resource_type: Optional[str] = None,
                                        result_output_file: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeHostCdnInstanceListResult:
    """
    Use this data source to query detailed information of ssl describe_host_cdn_instance_list

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_host_cdn_instance_list = tencentcloud.Ssl.get_describe_host_cdn_instance_list(certificate_id="8u8DII0l",
        resource_type="cdn")
    ```


    :param int async_cache: Whether.
    :param str certificate_id: Certificate ID to be deployed.
    :param Sequence[pulumi.InputType['GetDescribeHostCdnInstanceListFilterArgs']] filters: List of filtering parameters; Filterkey: domainmatch.
    :param int is_cache: Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
    :param str old_certificate_id: Original certificate ID.
    :param str resource_type: Deploy resource type.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['asyncCache'] = async_cache
    __args__['certificateId'] = certificate_id
    __args__['filters'] = filters
    __args__['isCache'] = is_cache
    __args__['oldCertificateId'] = old_certificate_id
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ssl/getDescribeHostCdnInstanceList:getDescribeHostCdnInstanceList', __args__, opts=opts, typ=GetDescribeHostCdnInstanceListResult).value

    return AwaitableGetDescribeHostCdnInstanceListResult(
        async_cache=__ret__.async_cache,
        async_cache_time=__ret__.async_cache_time,
        async_offset=__ret__.async_offset,
        async_total_num=__ret__.async_total_num,
        certificate_id=__ret__.certificate_id,
        filters=__ret__.filters,
        id=__ret__.id,
        instance_lists=__ret__.instance_lists,
        is_cache=__ret__.is_cache,
        old_certificate_id=__ret__.old_certificate_id,
        resource_type=__ret__.resource_type,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_describe_host_cdn_instance_list)
def get_describe_host_cdn_instance_list_output(async_cache: Optional[pulumi.Input[Optional[int]]] = None,
                                               certificate_id: Optional[pulumi.Input[str]] = None,
                                               filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDescribeHostCdnInstanceListFilterArgs']]]]] = None,
                                               is_cache: Optional[pulumi.Input[Optional[int]]] = None,
                                               old_certificate_id: Optional[pulumi.Input[Optional[str]]] = None,
                                               resource_type: Optional[pulumi.Input[str]] = None,
                                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeHostCdnInstanceListResult]:
    """
    Use this data source to query detailed information of ssl describe_host_cdn_instance_list

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_host_cdn_instance_list = tencentcloud.Ssl.get_describe_host_cdn_instance_list(certificate_id="8u8DII0l",
        resource_type="cdn")
    ```


    :param int async_cache: Whether.
    :param str certificate_id: Certificate ID to be deployed.
    :param Sequence[pulumi.InputType['GetDescribeHostCdnInstanceListFilterArgs']] filters: List of filtering parameters; Filterkey: domainmatch.
    :param int is_cache: Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
    :param str old_certificate_id: Original certificate ID.
    :param str resource_type: Deploy resource type.
    :param str result_output_file: Used to save results.
    """
    ...
