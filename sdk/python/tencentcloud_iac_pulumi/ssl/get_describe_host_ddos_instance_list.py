# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetDescribeHostDdosInstanceListResult',
    'AwaitableGetDescribeHostDdosInstanceListResult',
    'get_describe_host_ddos_instance_list',
    'get_describe_host_ddos_instance_list_output',
]

@pulumi.output_type
class GetDescribeHostDdosInstanceListResult:
    """
    A collection of values returned by getDescribeHostDdosInstanceList.
    """
    def __init__(__self__, certificate_id=None, filters=None, id=None, instance_lists=None, is_cache=None, old_certificate_id=None, resource_type=None, result_output_file=None):
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
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDescribeHostDdosInstanceListFilterResult']]:
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
    def instance_lists(self) -> Sequence['outputs.GetDescribeHostDdosInstanceListInstanceListResult']:
        """
        DDOS example listNote: This field may return NULL, indicating that the valid value cannot be obtained.
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


class AwaitableGetDescribeHostDdosInstanceListResult(GetDescribeHostDdosInstanceListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeHostDdosInstanceListResult(
            certificate_id=self.certificate_id,
            filters=self.filters,
            id=self.id,
            instance_lists=self.instance_lists,
            is_cache=self.is_cache,
            old_certificate_id=self.old_certificate_id,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file)


def get_describe_host_ddos_instance_list(certificate_id: Optional[str] = None,
                                         filters: Optional[Sequence[pulumi.InputType['GetDescribeHostDdosInstanceListFilterArgs']]] = None,
                                         is_cache: Optional[int] = None,
                                         old_certificate_id: Optional[str] = None,
                                         resource_type: Optional[str] = None,
                                         result_output_file: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeHostDdosInstanceListResult:
    """
    Use this data source to query detailed information of ssl describe_host_ddos_instance_list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_host_ddos_instance_list = tencentcloud.Ssl.get_describe_host_ddos_instance_list(certificate_id="8u8DII0l",
        resource_type="ddos")
    ```
    <!--End PulumiCodeChooser -->


    :param str certificate_id: Certificate ID to be deployed.
    :param Sequence[pulumi.InputType['GetDescribeHostDdosInstanceListFilterArgs']] filters: List of filtering parameters; Filterkey: domainmatch.
    :param int is_cache: Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
    :param str old_certificate_id: Deployed certificate ID.
    :param str resource_type: Deploy resource type.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['certificateId'] = certificate_id
    __args__['filters'] = filters
    __args__['isCache'] = is_cache
    __args__['oldCertificateId'] = old_certificate_id
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ssl/getDescribeHostDdosInstanceList:getDescribeHostDdosInstanceList', __args__, opts=opts, typ=GetDescribeHostDdosInstanceListResult).value

    return AwaitableGetDescribeHostDdosInstanceListResult(
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        instance_lists=pulumi.get(__ret__, 'instance_lists'),
        is_cache=pulumi.get(__ret__, 'is_cache'),
        old_certificate_id=pulumi.get(__ret__, 'old_certificate_id'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_describe_host_ddos_instance_list)
def get_describe_host_ddos_instance_list_output(certificate_id: Optional[pulumi.Input[str]] = None,
                                                filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDescribeHostDdosInstanceListFilterArgs']]]]] = None,
                                                is_cache: Optional[pulumi.Input[Optional[int]]] = None,
                                                old_certificate_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                resource_type: Optional[pulumi.Input[str]] = None,
                                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeHostDdosInstanceListResult]:
    """
    Use this data source to query detailed information of ssl describe_host_ddos_instance_list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_host_ddos_instance_list = tencentcloud.Ssl.get_describe_host_ddos_instance_list(certificate_id="8u8DII0l",
        resource_type="ddos")
    ```
    <!--End PulumiCodeChooser -->


    :param str certificate_id: Certificate ID to be deployed.
    :param Sequence[pulumi.InputType['GetDescribeHostDdosInstanceListFilterArgs']] filters: List of filtering parameters; Filterkey: domainmatch.
    :param int is_cache: Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
    :param str old_certificate_id: Deployed certificate ID.
    :param str resource_type: Deploy resource type.
    :param str result_output_file: Used to save results.
    """
    ...
