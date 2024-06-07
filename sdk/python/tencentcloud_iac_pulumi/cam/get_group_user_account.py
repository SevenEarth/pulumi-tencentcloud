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

__all__ = [
    'GetGroupUserAccountResult',
    'AwaitableGetGroupUserAccountResult',
    'get_group_user_account',
    'get_group_user_account_output',
]

@pulumi.output_type
class GetGroupUserAccountResult:
    """
    A collection of values returned by getGroupUserAccount.
    """
    def __init__(__self__, group_infos=None, id=None, result_output_file=None, rp=None, sub_uin=None, total_num=None, uid=None):
        if group_infos and not isinstance(group_infos, list):
            raise TypeError("Expected argument 'group_infos' to be a list")
        pulumi.set(__self__, "group_infos", group_infos)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if rp and not isinstance(rp, int):
            raise TypeError("Expected argument 'rp' to be a int")
        pulumi.set(__self__, "rp", rp)
        if sub_uin and not isinstance(sub_uin, int):
            raise TypeError("Expected argument 'sub_uin' to be a int")
        pulumi.set(__self__, "sub_uin", sub_uin)
        if total_num and not isinstance(total_num, int):
            raise TypeError("Expected argument 'total_num' to be a int")
        pulumi.set(__self__, "total_num", total_num)
        if uid and not isinstance(uid, int):
            raise TypeError("Expected argument 'uid' to be a int")
        pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="groupInfos")
    def group_infos(self) -> Sequence['outputs.GetGroupUserAccountGroupInfoResult']:
        """
        User group information.
        """
        return pulumi.get(self, "group_infos")

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
    @pulumi.getter
    def rp(self) -> Optional[int]:
        return pulumi.get(self, "rp")

    @property
    @pulumi.getter(name="subUin")
    def sub_uin(self) -> Optional[int]:
        return pulumi.get(self, "sub_uin")

    @property
    @pulumi.getter(name="totalNum")
    def total_num(self) -> int:
        """
        The total number of user groups the sub-user has joined.
        """
        return pulumi.get(self, "total_num")

    @property
    @pulumi.getter
    def uid(self) -> Optional[int]:
        return pulumi.get(self, "uid")


class AwaitableGetGroupUserAccountResult(GetGroupUserAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupUserAccountResult(
            group_infos=self.group_infos,
            id=self.id,
            result_output_file=self.result_output_file,
            rp=self.rp,
            sub_uin=self.sub_uin,
            total_num=self.total_num,
            uid=self.uid)


def get_group_user_account(result_output_file: Optional[str] = None,
                           rp: Optional[int] = None,
                           sub_uin: Optional[int] = None,
                           uid: Optional[int] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupUserAccountResult:
    """
    Use this data source to query detailed information of cam group_user_account

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    group_user_account = tencentcloud.Cam.get_group_user_account(sub_uin=100033690181)
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    :param int rp: Number per page. The default is 20.
    :param int sub_uin: Sub-user uin.
    :param int uid: Sub-user uid.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    __args__['rp'] = rp
    __args__['subUin'] = sub_uin
    __args__['uid'] = uid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cam/getGroupUserAccount:getGroupUserAccount', __args__, opts=opts, typ=GetGroupUserAccountResult).value

    return AwaitableGetGroupUserAccountResult(
        group_infos=pulumi.get(__ret__, 'group_infos'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        rp=pulumi.get(__ret__, 'rp'),
        sub_uin=pulumi.get(__ret__, 'sub_uin'),
        total_num=pulumi.get(__ret__, 'total_num'),
        uid=pulumi.get(__ret__, 'uid'))


@_utilities.lift_output_func(get_group_user_account)
def get_group_user_account_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  rp: Optional[pulumi.Input[Optional[int]]] = None,
                                  sub_uin: Optional[pulumi.Input[Optional[int]]] = None,
                                  uid: Optional[pulumi.Input[Optional[int]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupUserAccountResult]:
    """
    Use this data source to query detailed information of cam group_user_account

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    group_user_account = tencentcloud.Cam.get_group_user_account(sub_uin=100033690181)
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    :param int rp: Number per page. The default is 20.
    :param int sub_uin: Sub-user uin.
    :param int uid: Sub-user uid.
    """
    ...
