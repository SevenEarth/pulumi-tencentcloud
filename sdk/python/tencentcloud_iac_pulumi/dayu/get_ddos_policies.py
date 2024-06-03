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
    'GetDdosPoliciesResult',
    'AwaitableGetDdosPoliciesResult',
    'get_ddos_policies',
    'get_ddos_policies_output',
]

@pulumi.output_type
class GetDdosPoliciesResult:
    """
    A collection of values returned by getDdosPolicies.
    """
    def __init__(__self__, id=None, lists=None, policy_id=None, resource_type=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        pulumi.set(__self__, "policy_id", policy_id)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetDdosPoliciesListResult']:
        """
        A list of DDoS policies. Each element contains the following attributes:
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[str]:
        """
        Id of policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDdosPoliciesResult(GetDdosPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDdosPoliciesResult(
            id=self.id,
            lists=self.lists,
            policy_id=self.policy_id,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file)


def get_ddos_policies(policy_id: Optional[str] = None,
                      resource_type: Optional[str] = None,
                      result_output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDdosPoliciesResult:
    """
    Use this data source to query dayu DDoS policies

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    id_test = tencentcloud.Dayu.get_ddos_policies(resource_type=tencentcloud_dayu_ddos_policy["test_policy"]["resource_type"],
        policy_id=tencentcloud_dayu_ddos_policy["test_policy"]["policy_id"])
    ```
    <!--End PulumiCodeChooser -->


    :param str policy_id: ID of the DDoS policy to be query.
    :param str resource_type: Type of the resource that the DDoS policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['policyId'] = policy_id
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dayu/getDdosPolicies:getDdosPolicies', __args__, opts=opts, typ=GetDdosPoliciesResult).value

    return AwaitableGetDdosPoliciesResult(
        id=pulumi.get(__ret__, 'id'),
        lists=pulumi.get(__ret__, 'lists'),
        policy_id=pulumi.get(__ret__, 'policy_id'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_ddos_policies)
def get_ddos_policies_output(policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                             resource_type: Optional[pulumi.Input[str]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDdosPoliciesResult]:
    """
    Use this data source to query dayu DDoS policies

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    id_test = tencentcloud.Dayu.get_ddos_policies(resource_type=tencentcloud_dayu_ddos_policy["test_policy"]["resource_type"],
        policy_id=tencentcloud_dayu_ddos_policy["test_policy"]["policy_id"])
    ```
    <!--End PulumiCodeChooser -->


    :param str policy_id: ID of the DDoS policy to be query.
    :param str resource_type: Type of the resource that the DDoS policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
    :param str result_output_file: Used to save results.
    """
    ...
