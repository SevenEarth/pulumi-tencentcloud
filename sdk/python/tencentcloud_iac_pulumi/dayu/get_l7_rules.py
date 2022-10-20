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
    'GetL7RulesResult',
    'AwaitableGetL7RulesResult',
    'get_l7_rules',
    'get_l7_rules_output',
]

@pulumi.output_type
class GetL7RulesResult:
    """
    A collection of values returned by getL7Rules.
    """
    def __init__(__self__, domain=None, id=None, lists=None, resource_id=None, resource_type=None, result_output_file=None, rule_id=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if rule_id and not isinstance(rule_id, str):
            raise TypeError("Expected argument 'rule_id' to be a str")
        pulumi.set(__self__, "rule_id", rule_id)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        Domain that the 7 layer rule works for.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetL7RulesListResult']:
        """
        A list of layer 7 rules. Each element contains the following attributes:
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[str]:
        """
        Id of the 7 layer rule.
        """
        return pulumi.get(self, "rule_id")


class AwaitableGetL7RulesResult(GetL7RulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetL7RulesResult(
            domain=self.domain,
            id=self.id,
            lists=self.lists,
            resource_id=self.resource_id,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file,
            rule_id=self.rule_id)


def get_l7_rules(domain: Optional[str] = None,
                 resource_id: Optional[str] = None,
                 resource_type: Optional[str] = None,
                 result_output_file: Optional[str] = None,
                 rule_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetL7RulesResult:
    """
    Use this data source to query dayu layer 7 rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domain_test = tencentcloud.Dayu.get_l7_rules(resource_type=tencentcloud_dayu_l7_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l7_rule["test_rule"]["resource_id"],
        domain=tencentcloud_dayu_l7_rule["test_rule"]["domain"])
    id_test = tencentcloud.Dayu.get_l7_rules(resource_type=tencentcloud_dayu_l7_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l7_rule["test_rule"]["resource_id"],
        rule_id=tencentcloud_dayu_l7_rule["test_rule"]["rule_id"])
    ```


    :param str domain: Domain of the layer 7 rule to be queried.
    :param str resource_id: Id of the resource that the layer 7 rule works for.
    :param str resource_type: Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
    :param str result_output_file: Used to save results.
    :param str rule_id: Id of the layer 7 rule to be queried.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['resourceId'] = resource_id
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    __args__['ruleId'] = rule_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dayu/getL7Rules:getL7Rules', __args__, opts=opts, typ=GetL7RulesResult).value

    return AwaitableGetL7RulesResult(
        domain=__ret__.domain,
        id=__ret__.id,
        lists=__ret__.lists,
        resource_id=__ret__.resource_id,
        resource_type=__ret__.resource_type,
        result_output_file=__ret__.result_output_file,
        rule_id=__ret__.rule_id)


@_utilities.lift_output_func(get_l7_rules)
def get_l7_rules_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                        resource_id: Optional[pulumi.Input[str]] = None,
                        resource_type: Optional[pulumi.Input[str]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        rule_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetL7RulesResult]:
    """
    Use this data source to query dayu layer 7 rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domain_test = tencentcloud.Dayu.get_l7_rules(resource_type=tencentcloud_dayu_l7_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l7_rule["test_rule"]["resource_id"],
        domain=tencentcloud_dayu_l7_rule["test_rule"]["domain"])
    id_test = tencentcloud.Dayu.get_l7_rules(resource_type=tencentcloud_dayu_l7_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l7_rule["test_rule"]["resource_id"],
        rule_id=tencentcloud_dayu_l7_rule["test_rule"]["rule_id"])
    ```


    :param str domain: Domain of the layer 7 rule to be queried.
    :param str resource_id: Id of the resource that the layer 7 rule works for.
    :param str resource_type: Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
    :param str result_output_file: Used to save results.
    :param str rule_id: Id of the layer 7 rule to be queried.
    """
    ...
