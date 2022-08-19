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
    'GetL4RulesResult',
    'AwaitableGetL4RulesResult',
    'get_l4_rules',
    'get_l4_rules_output',
]

@pulumi.output_type
class GetL4RulesResult:
    """
    A collection of values returned by getL4Rules.
    """
    def __init__(__self__, id=None, lists=None, name=None, resource_id=None, resource_type=None, result_output_file=None, rule_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetL4RulesListResult']:
        """
        A list of layer 4 rules. Each element contains the following attributes:
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the rule.
        """
        return pulumi.get(self, "name")

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
        ID of the 4 layer rule.
        """
        return pulumi.get(self, "rule_id")


class AwaitableGetL4RulesResult(GetL4RulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetL4RulesResult(
            id=self.id,
            lists=self.lists,
            name=self.name,
            resource_id=self.resource_id,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file,
            rule_id=self.rule_id)


def get_l4_rules(name: Optional[str] = None,
                 resource_id: Optional[str] = None,
                 resource_type: Optional[str] = None,
                 result_output_file: Optional[str] = None,
                 rule_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetL4RulesResult:
    """
    Use this data source to query dayu layer 4 rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    name_test = tencentcloud.Dayu.get_l4_rules(resource_type=tencentcloud_dayu_l4_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l4_rule["test_rule"]["resource_id"],
        name=tencentcloud_dayu_l4_rule["test_rule"]["name"])
    id_test = tencentcloud.Dayu.get_l4_rules(resource_type=tencentcloud_dayu_l4_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l4_rule["test_rule"]["resource_id"],
        rule_id=tencentcloud_dayu_l4_rule["test_rule"]["rule_id"])
    ```


    :param str name: Name of the layer 4 rule to be queried.
    :param str resource_id: Id of the resource that the layer 4 rule works for.
    :param str resource_type: Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
    :param str result_output_file: Used to save results.
    :param str rule_id: Id of the layer 4 rule to be queried.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceId'] = resource_id
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    __args__['ruleId'] = rule_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dayu/getL4Rules:getL4Rules', __args__, opts=opts, typ=GetL4RulesResult).value

    return AwaitableGetL4RulesResult(
        id=__ret__.id,
        lists=__ret__.lists,
        name=__ret__.name,
        resource_id=__ret__.resource_id,
        resource_type=__ret__.resource_type,
        result_output_file=__ret__.result_output_file,
        rule_id=__ret__.rule_id)


@_utilities.lift_output_func(get_l4_rules)
def get_l4_rules_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        resource_id: Optional[pulumi.Input[str]] = None,
                        resource_type: Optional[pulumi.Input[str]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        rule_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetL4RulesResult]:
    """
    Use this data source to query dayu layer 4 rules

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    name_test = tencentcloud.Dayu.get_l4_rules(resource_type=tencentcloud_dayu_l4_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l4_rule["test_rule"]["resource_id"],
        name=tencentcloud_dayu_l4_rule["test_rule"]["name"])
    id_test = tencentcloud.Dayu.get_l4_rules(resource_type=tencentcloud_dayu_l4_rule["test_rule"]["resource_type"],
        resource_id=tencentcloud_dayu_l4_rule["test_rule"]["resource_id"],
        rule_id=tencentcloud_dayu_l4_rule["test_rule"]["rule_id"])
    ```


    :param str name: Name of the layer 4 rule to be queried.
    :param str resource_id: Id of the resource that the layer 4 rule works for.
    :param str resource_type: Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
    :param str result_output_file: Used to save results.
    :param str rule_id: Id of the layer 4 rule to be queried.
    """
    ...
