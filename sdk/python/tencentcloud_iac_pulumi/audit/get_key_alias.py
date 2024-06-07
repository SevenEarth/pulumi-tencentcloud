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
    'GetKeyAliasResult',
    'AwaitableGetKeyAliasResult',
    'get_key_alias',
    'get_key_alias_output',
]

@pulumi.output_type
class GetKeyAliasResult:
    """
    A collection of values returned by getKeyAlias.
    """
    def __init__(__self__, audit_key_alias_lists=None, id=None, region=None, result_output_file=None):
        if audit_key_alias_lists and not isinstance(audit_key_alias_lists, list):
            raise TypeError("Expected argument 'audit_key_alias_lists' to be a list")
        pulumi.set(__self__, "audit_key_alias_lists", audit_key_alias_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="auditKeyAliasLists")
    def audit_key_alias_lists(self) -> Sequence['outputs.GetKeyAliasAuditKeyAliasListResult']:
        """
        List of available key alias supported by audit.
        """
        return pulumi.get(self, "audit_key_alias_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetKeyAliasResult(GetKeyAliasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyAliasResult(
            audit_key_alias_lists=self.audit_key_alias_lists,
            id=self.id,
            region=self.region,
            result_output_file=self.result_output_file)


def get_key_alias(region: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyAliasResult:
    """
    Use this data source to query the key alias list specified with region supported by the audit.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    all = tencentcloud.Audit.get_key_alias(region="ap-hongkong")
    ```
    <!--End PulumiCodeChooser -->


    :param str region: Region.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Audit/getKeyAlias:getKeyAlias', __args__, opts=opts, typ=GetKeyAliasResult).value

    return AwaitableGetKeyAliasResult(
        audit_key_alias_lists=pulumi.get(__ret__, 'audit_key_alias_lists'),
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_key_alias)
def get_key_alias_output(region: Optional[pulumi.Input[str]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyAliasResult]:
    """
    Use this data source to query the key alias list specified with region supported by the audit.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    all = tencentcloud.Audit.get_key_alias(region="ap-hongkong")
    ```
    <!--End PulumiCodeChooser -->


    :param str region: Region.
    :param str result_output_file: Used to save results.
    """
    ...
