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
    'GetEventRulesResult',
    'AwaitableGetEventRulesResult',
    'get_event_rules',
    'get_event_rules_output',
]

@pulumi.output_type
class GetEventRulesResult:
    """
    A collection of values returned by getEventRules.
    """
    def __init__(__self__, event_bus_id=None, id=None, order=None, order_by=None, result_output_file=None, rules=None):
        if event_bus_id and not isinstance(event_bus_id, str):
            raise TypeError("Expected argument 'event_bus_id' to be a str")
        pulumi.set(__self__, "event_bus_id", event_bus_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order and not isinstance(order, str):
            raise TypeError("Expected argument 'order' to be a str")
        pulumi.set(__self__, "order", order)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="eventBusId")
    def event_bus_id(self) -> str:
        """
        event bus Id.
        """
        return pulumi.get(self, "event_bus_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetEventRulesRuleResult']:
        """
        Event rule information.
        """
        return pulumi.get(self, "rules")


class AwaitableGetEventRulesResult(GetEventRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventRulesResult(
            event_bus_id=self.event_bus_id,
            id=self.id,
            order=self.order,
            order_by=self.order_by,
            result_output_file=self.result_output_file,
            rules=self.rules)


def get_event_rules(event_bus_id: Optional[str] = None,
                    order: Optional[str] = None,
                    order_by: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventRulesResult:
    """
    Use this data source to query detailed information of eb event_rules

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import json
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo = tencentcloud.eb.EventBus("foo",
        event_bus_name="tf-event_bus_rule",
        description="event bus desc",
        enable_store=False,
        save_days=1,
        tags={
            "createdBy": "terraform",
        })
    event_rule = tencentcloud.eb.EventRule("eventRule",
        event_bus_id=foo.id,
        rule_name="tf-event_rule",
        description="event rule desc",
        enable=True,
        event_pattern=json.dumps({
            "source": "apigw.cloud.tencent",
            "type": ["connector:apigw"],
        }),
        tags={
            "createdBy": "terraform",
        })
    event_rules = tencentcloud.Eb.get_event_rules_output(event_bus_id=foo.id,
        order_by="AddTime",
        order="DESC")
    ```
    <!--End PulumiCodeChooser -->


    :param str event_bus_id: event bus Id.
    :param str order: Return results in ascending or descending order, optional values ASC (ascending) and DESC (descending).
    :param str order_by: According to which field to sort the returned results, the following fields are supported: AddTime (creation time), ModTime (modification time).
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['eventBusId'] = event_bus_id
    __args__['order'] = order
    __args__['orderBy'] = order_by
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Eb/getEventRules:getEventRules', __args__, opts=opts, typ=GetEventRulesResult).value

    return AwaitableGetEventRulesResult(
        event_bus_id=pulumi.get(__ret__, 'event_bus_id'),
        id=pulumi.get(__ret__, 'id'),
        order=pulumi.get(__ret__, 'order'),
        order_by=pulumi.get(__ret__, 'order_by'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        rules=pulumi.get(__ret__, 'rules'))


@_utilities.lift_output_func(get_event_rules)
def get_event_rules_output(event_bus_id: Optional[pulumi.Input[str]] = None,
                           order: Optional[pulumi.Input[Optional[str]]] = None,
                           order_by: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEventRulesResult]:
    """
    Use this data source to query detailed information of eb event_rules

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import json
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo = tencentcloud.eb.EventBus("foo",
        event_bus_name="tf-event_bus_rule",
        description="event bus desc",
        enable_store=False,
        save_days=1,
        tags={
            "createdBy": "terraform",
        })
    event_rule = tencentcloud.eb.EventRule("eventRule",
        event_bus_id=foo.id,
        rule_name="tf-event_rule",
        description="event rule desc",
        enable=True,
        event_pattern=json.dumps({
            "source": "apigw.cloud.tencent",
            "type": ["connector:apigw"],
        }),
        tags={
            "createdBy": "terraform",
        })
    event_rules = tencentcloud.Eb.get_event_rules_output(event_bus_id=foo.id,
        order_by="AddTime",
        order="DESC")
    ```
    <!--End PulumiCodeChooser -->


    :param str event_bus_id: event bus Id.
    :param str order: Return results in ascending or descending order, optional values ASC (ascending) and DESC (descending).
    :param str order_by: According to which field to sort the returned results, the following fields are supported: AddTime (creation time), ModTime (modification time).
    :param str result_output_file: Used to save results.
    """
    ...
