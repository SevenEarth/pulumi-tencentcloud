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
    'GetWafRuleGroupsResult',
    'AwaitableGetWafRuleGroupsResult',
    'get_waf_rule_groups',
    'get_waf_rule_groups_output',
]

@pulumi.output_type
class GetWafRuleGroupsResult:
    """
    A collection of values returned by getWafRuleGroups.
    """
    def __init__(__self__, entity=None, id=None, result_output_file=None, waf_rule_groups=None, zone_id=None):
        if entity and not isinstance(entity, str):
            raise TypeError("Expected argument 'entity' to be a str")
        pulumi.set(__self__, "entity", entity)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if waf_rule_groups and not isinstance(waf_rule_groups, list):
            raise TypeError("Expected argument 'waf_rule_groups' to be a list")
        pulumi.set(__self__, "waf_rule_groups", waf_rule_groups)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def entity(self) -> str:
        return pulumi.get(self, "entity")

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
    @pulumi.getter(name="wafRuleGroups")
    def waf_rule_groups(self) -> Sequence['outputs.GetWafRuleGroupsWafRuleGroupResult']:
        """
        List of WAF rule groups.
        """
        return pulumi.get(self, "waf_rule_groups")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        return pulumi.get(self, "zone_id")


class AwaitableGetWafRuleGroupsResult(GetWafRuleGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafRuleGroupsResult(
            entity=self.entity,
            id=self.id,
            result_output_file=self.result_output_file,
            waf_rule_groups=self.waf_rule_groups,
            zone_id=self.zone_id)


def get_waf_rule_groups(entity: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        zone_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafRuleGroupsResult:
    """
    Use this data source to query detailed information of teo wafRuleGroups

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    waf_rule_groups = tencentcloud.Teo.get_waf_rule_groups(entity="",
        zone_id="")
    ```


    :param str entity: Subdomain or application name.
    :param str result_output_file: Used to save results.
    :param str zone_id: Site ID.
    """
    __args__ = dict()
    __args__['entity'] = entity
    __args__['resultOutputFile'] = result_output_file
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Teo/getWafRuleGroups:getWafRuleGroups', __args__, opts=opts, typ=GetWafRuleGroupsResult).value

    return AwaitableGetWafRuleGroupsResult(
        entity=__ret__.entity,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        waf_rule_groups=__ret__.waf_rule_groups,
        zone_id=__ret__.zone_id)


@_utilities.lift_output_func(get_waf_rule_groups)
def get_waf_rule_groups_output(entity: Optional[pulumi.Input[str]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               zone_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWafRuleGroupsResult]:
    """
    Use this data source to query detailed information of teo wafRuleGroups

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    waf_rule_groups = tencentcloud.Teo.get_waf_rule_groups(entity="",
        zone_id="")
    ```


    :param str entity: Subdomain or application name.
    :param str result_output_file: Used to save results.
    :param str zone_id: Site ID.
    """
    ...
