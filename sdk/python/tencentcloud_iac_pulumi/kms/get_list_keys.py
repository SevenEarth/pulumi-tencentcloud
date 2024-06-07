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
    'GetListKeysResult',
    'AwaitableGetListKeysResult',
    'get_list_keys',
    'get_list_keys_output',
]

@pulumi.output_type
class GetListKeysResult:
    """
    A collection of values returned by getListKeys.
    """
    def __init__(__self__, hsm_cluster_id=None, id=None, keys=None, result_output_file=None, role=None):
        if hsm_cluster_id and not isinstance(hsm_cluster_id, str):
            raise TypeError("Expected argument 'hsm_cluster_id' to be a str")
        pulumi.set(__self__, "hsm_cluster_id", hsm_cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if role and not isinstance(role, int):
            raise TypeError("Expected argument 'role' to be a int")
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="hsmClusterId")
    def hsm_cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "hsm_cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetListKeysKeyResult']:
        """
        A list of KMS keys.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def role(self) -> Optional[int]:
        return pulumi.get(self, "role")


class AwaitableGetListKeysResult(GetListKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetListKeysResult(
            hsm_cluster_id=self.hsm_cluster_id,
            id=self.id,
            keys=self.keys,
            result_output_file=self.result_output_file,
            role=self.role)


def get_list_keys(hsm_cluster_id: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  role: Optional[int] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetListKeysResult:
    """
    Use this data source to query detailed information of kms list_keys

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Kms.get_list_keys(role=1)
    ```
    <!--End PulumiCodeChooser -->


    :param str hsm_cluster_id: HSM cluster ID (only valid for KMS exclusive/managed service instances).
    :param str result_output_file: Used to save results.
    :param int role: Filter based on the creator role. The default value is 0, which indicates the cmk created by the user himself, and 1, which indicates the cmk automatically created by authorizing other cloud products.
    """
    __args__ = dict()
    __args__['hsmClusterId'] = hsm_cluster_id
    __args__['resultOutputFile'] = result_output_file
    __args__['role'] = role
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Kms/getListKeys:getListKeys', __args__, opts=opts, typ=GetListKeysResult).value

    return AwaitableGetListKeysResult(
        hsm_cluster_id=pulumi.get(__ret__, 'hsm_cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        keys=pulumi.get(__ret__, 'keys'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        role=pulumi.get(__ret__, 'role'))


@_utilities.lift_output_func(get_list_keys)
def get_list_keys_output(hsm_cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         role: Optional[pulumi.Input[Optional[int]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetListKeysResult]:
    """
    Use this data source to query detailed information of kms list_keys

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Kms.get_list_keys(role=1)
    ```
    <!--End PulumiCodeChooser -->


    :param str hsm_cluster_id: HSM cluster ID (only valid for KMS exclusive/managed service instances).
    :param str result_output_file: Used to save results.
    :param int role: Filter based on the creator role. The default value is 0, which indicates the cmk created by the user himself, and 1, which indicates the cmk automatically created by authorizing other cloud products.
    """
    ...
