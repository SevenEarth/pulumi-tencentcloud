# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDescribeUserTypeResult',
    'AwaitableGetDescribeUserTypeResult',
    'get_describe_user_type',
    'get_describe_user_type_output',
]

@pulumi.output_type
class GetDescribeUserTypeResult:
    """
    A collection of values returned by getDescribeUserType.
    """
    def __init__(__self__, id=None, result_output_file=None, user_id=None, user_type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if user_type and not isinstance(user_type, str):
            raise TypeError("Expected argument 'user_type' to be a str")
        pulumi.set(__self__, "user_type", user_type)

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
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[str]:
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userType")
    def user_type(self) -> str:
        """
        User type, only support: ADMIN: ddministrator/COMMON: ordinary user.
        """
        return pulumi.get(self, "user_type")


class AwaitableGetDescribeUserTypeResult(GetDescribeUserTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeUserTypeResult(
            id=self.id,
            result_output_file=self.result_output_file,
            user_id=self.user_id,
            user_type=self.user_type)


def get_describe_user_type(result_output_file: Optional[str] = None,
                           user_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeUserTypeResult:
    """
    Use this data source to query detailed information of dlc describe_user_type

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_user_type = tencentcloud.Dlc.get_describe_user_type(user_id="127382378")
    ```


    :param str result_output_file: Used to save results.
    :param str user_id: User id (uin), if left blank, it defaults to the caller's sub-uin.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    __args__['userId'] = user_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dlc/getDescribeUserType:getDescribeUserType', __args__, opts=opts, typ=GetDescribeUserTypeResult).value

    return AwaitableGetDescribeUserTypeResult(
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        user_id=__ret__.user_id,
        user_type=__ret__.user_type)


@_utilities.lift_output_func(get_describe_user_type)
def get_describe_user_type_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  user_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeUserTypeResult]:
    """
    Use this data source to query detailed information of dlc describe_user_type

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_user_type = tencentcloud.Dlc.get_describe_user_type(user_id="127382378")
    ```


    :param str result_output_file: Used to save results.
    :param str user_id: User id (uin), if left blank, it defaults to the caller's sub-uin.
    """
    ...
