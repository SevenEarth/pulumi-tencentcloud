# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetInfoResult',
    'AwaitableGetInfoResult',
    'get_info',
    'get_info_output',
]

@pulumi.output_type
class GetInfoResult:
    """
    A collection of values returned by getInfo.
    """
    def __init__(__self__, app_id=None, id=None, name=None, owner_uin=None, result_output_file=None, uin=None):
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_uin and not isinstance(owner_uin, str):
            raise TypeError("Expected argument 'owner_uin' to be a str")
        pulumi.set(__self__, "owner_uin", owner_uin)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if uin and not isinstance(uin, str):
            raise TypeError("Expected argument 'uin' to be a str")
        pulumi.set(__self__, "uin", uin)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        """
        Current account App ID.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Current account Name. NOTE: only support subaccount.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerUin")
    def owner_uin(self) -> str:
        """
        Current account OwnerUIN.
        """
        return pulumi.get(self, "owner_uin")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def uin(self) -> str:
        """
        Current account UIN.
        """
        return pulumi.get(self, "uin")


class AwaitableGetInfoResult(GetInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInfoResult(
            app_id=self.app_id,
            id=self.id,
            name=self.name,
            owner_uin=self.owner_uin,
            result_output_file=self.result_output_file,
            uin=self.uin)


def get_info(result_output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInfoResult:
    """
    Use this data source to query user appid, uin and ownerUin.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.User.get_info()
    ```


    :param str result_output_file: Used for save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:User/getInfo:getInfo', __args__, opts=opts, typ=GetInfoResult).value

    return AwaitableGetInfoResult(
        app_id=__ret__.app_id,
        id=__ret__.id,
        name=__ret__.name,
        owner_uin=__ret__.owner_uin,
        result_output_file=__ret__.result_output_file,
        uin=__ret__.uin)


@_utilities.lift_output_func(get_info)
def get_info_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInfoResult]:
    """
    Use this data source to query user appid, uin and ownerUin.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.User.get_info()
    ```


    :param str result_output_file: Used for save results.
    """
    ...
