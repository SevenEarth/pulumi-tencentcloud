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
    'GetDescribeUserInfoResult',
    'AwaitableGetDescribeUserInfoResult',
    'get_describe_user_info',
    'get_describe_user_info_output',
]

@pulumi.output_type
class GetDescribeUserInfoResult:
    """
    A collection of values returned by getDescribeUserInfo.
    """
    def __init__(__self__, filters=None, id=None, result_output_file=None, sort_by=None, sorting=None, type=None, user_id=None, user_infos=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sort_by and not isinstance(sort_by, str):
            raise TypeError("Expected argument 'sort_by' to be a str")
        pulumi.set(__self__, "sort_by", sort_by)
        if sorting and not isinstance(sorting, str):
            raise TypeError("Expected argument 'sorting' to be a str")
        pulumi.set(__self__, "sorting", sorting)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if user_infos and not isinstance(user_infos, list):
            raise TypeError("Expected argument 'user_infos' to be a list")
        pulumi.set(__self__, "user_infos", user_infos)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDescribeUserInfoFilterResult']]:
        return pulumi.get(self, "filters")

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
    @pulumi.getter(name="sortBy")
    def sort_by(self) -> Optional[str]:
        return pulumi.get(self, "sort_by")

    @property
    @pulumi.getter
    def sorting(self) -> Optional[str]:
        return pulumi.get(self, "sorting")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of information returned, Group: the returned workgroup information of the current user; DataAuth: the returned data permission information of the current user; EngineAuth: the returned engine permission information of the current user.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[str]:
        """
        User id, the same as the sub-user uin.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userInfos")
    def user_infos(self) -> Sequence['outputs.GetDescribeUserInfoUserInfoResult']:
        """
        User details.
        """
        return pulumi.get(self, "user_infos")


class AwaitableGetDescribeUserInfoResult(GetDescribeUserInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeUserInfoResult(
            filters=self.filters,
            id=self.id,
            result_output_file=self.result_output_file,
            sort_by=self.sort_by,
            sorting=self.sorting,
            type=self.type,
            user_id=self.user_id,
            user_infos=self.user_infos)


def get_describe_user_info(filters: Optional[Sequence[pulumi.InputType['GetDescribeUserInfoFilterArgs']]] = None,
                           result_output_file: Optional[str] = None,
                           sort_by: Optional[str] = None,
                           sorting: Optional[str] = None,
                           type: Optional[str] = None,
                           user_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeUserInfoResult:
    """
    Use this data source to query detailed information of dlc describe_user_info

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_user_info = tencentcloud.Dlc.get_describe_user_info(sort_by="create-time",
        sorting="desc",
        type="Group",
        user_id="100032772113")
    ```


    :param Sequence[pulumi.InputType['GetDescribeUserInfoFilterArgs']] filters: Query filter conditions. when type is Group, fuzzy search with Key as workgroup-name is supported. when type is DataAuth, key is supported. policy-type: permission type, policy-source: data source, data-name: database table. Fuzzy search, when type is EngineAuth, supports fuzzy search of key, policy-type: permission type, policy-source: data source, engine-name: library table.
    :param str result_output_file: Used to save results.
    :param str sort_by: Sorting field, when type is Group, support create-time, group-name, when type is DataAuth, support create-time, when type is EngineAuth, support create-time.
    :param str sorting: Sorting method, desc means forward order, asc means reverse order, the default is asc.
    :param str type: Query information type, Group: work group DataAuth: data permission EngineAuth: engine permission.
    :param str user_id: User id, the same as the sub-user uin.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['resultOutputFile'] = result_output_file
    __args__['sortBy'] = sort_by
    __args__['sorting'] = sorting
    __args__['type'] = type
    __args__['userId'] = user_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dlc/getDescribeUserInfo:getDescribeUserInfo', __args__, opts=opts, typ=GetDescribeUserInfoResult).value

    return AwaitableGetDescribeUserInfoResult(
        filters=__ret__.filters,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        sort_by=__ret__.sort_by,
        sorting=__ret__.sorting,
        type=__ret__.type,
        user_id=__ret__.user_id,
        user_infos=__ret__.user_infos)


@_utilities.lift_output_func(get_describe_user_info)
def get_describe_user_info_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDescribeUserInfoFilterArgs']]]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  sort_by: Optional[pulumi.Input[Optional[str]]] = None,
                                  sorting: Optional[pulumi.Input[Optional[str]]] = None,
                                  type: Optional[pulumi.Input[Optional[str]]] = None,
                                  user_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeUserInfoResult]:
    """
    Use this data source to query detailed information of dlc describe_user_info

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_user_info = tencentcloud.Dlc.get_describe_user_info(sort_by="create-time",
        sorting="desc",
        type="Group",
        user_id="100032772113")
    ```


    :param Sequence[pulumi.InputType['GetDescribeUserInfoFilterArgs']] filters: Query filter conditions. when type is Group, fuzzy search with Key as workgroup-name is supported. when type is DataAuth, key is supported. policy-type: permission type, policy-source: data source, data-name: database table. Fuzzy search, when type is EngineAuth, supports fuzzy search of key, policy-type: permission type, policy-source: data source, engine-name: library table.
    :param str result_output_file: Used to save results.
    :param str sort_by: Sorting field, when type is Group, support create-time, group-name, when type is DataAuth, support create-time, when type is EngineAuth, support create-time.
    :param str sorting: Sorting method, desc means forward order, asc means reverse order, the default is asc.
    :param str type: Query information type, Group: work group DataAuth: data permission EngineAuth: engine permission.
    :param str user_id: User id, the same as the sub-user uin.
    """
    ...
