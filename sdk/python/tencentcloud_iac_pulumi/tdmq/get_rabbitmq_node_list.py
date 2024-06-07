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
from ._inputs import *

__all__ = [
    'GetRabbitmqNodeListResult',
    'AwaitableGetRabbitmqNodeListResult',
    'get_rabbitmq_node_list',
    'get_rabbitmq_node_list_output',
]

@pulumi.output_type
class GetRabbitmqNodeListResult:
    """
    A collection of values returned by getRabbitmqNodeList.
    """
    def __init__(__self__, filters=None, id=None, instance_id=None, node_lists=None, node_name=None, result_output_file=None, sort_element=None, sort_order=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if node_lists and not isinstance(node_lists, list):
            raise TypeError("Expected argument 'node_lists' to be a list")
        pulumi.set(__self__, "node_lists", node_lists)
        if node_name and not isinstance(node_name, str):
            raise TypeError("Expected argument 'node_name' to be a str")
        pulumi.set(__self__, "node_name", node_name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sort_element and not isinstance(sort_element, str):
            raise TypeError("Expected argument 'sort_element' to be a str")
        pulumi.set(__self__, "sort_element", sort_element)
        if sort_order and not isinstance(sort_order, str):
            raise TypeError("Expected argument 'sort_order' to be a str")
        pulumi.set(__self__, "sort_order", sort_order)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetRabbitmqNodeListFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nodeLists")
    def node_lists(self) -> Sequence['outputs.GetRabbitmqNodeListNodeListResult']:
        return pulumi.get(self, "node_lists")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[str]:
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="sortElement")
    def sort_element(self) -> Optional[str]:
        return pulumi.get(self, "sort_element")

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> Optional[str]:
        return pulumi.get(self, "sort_order")


class AwaitableGetRabbitmqNodeListResult(GetRabbitmqNodeListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRabbitmqNodeListResult(
            filters=self.filters,
            id=self.id,
            instance_id=self.instance_id,
            node_lists=self.node_lists,
            node_name=self.node_name,
            result_output_file=self.result_output_file,
            sort_element=self.sort_element,
            sort_order=self.sort_order)


def get_rabbitmq_node_list(filters: Optional[Sequence[pulumi.InputType['GetRabbitmqNodeListFilterArgs']]] = None,
                           instance_id: Optional[str] = None,
                           node_name: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           sort_element: Optional[str] = None,
                           sort_order: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRabbitmqNodeListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['instanceId'] = instance_id
    __args__['nodeName'] = node_name
    __args__['resultOutputFile'] = result_output_file
    __args__['sortElement'] = sort_element
    __args__['sortOrder'] = sort_order
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tdmq/getRabbitmqNodeList:getRabbitmqNodeList', __args__, opts=opts, typ=GetRabbitmqNodeListResult).value

    return AwaitableGetRabbitmqNodeListResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        node_lists=pulumi.get(__ret__, 'node_lists'),
        node_name=pulumi.get(__ret__, 'node_name'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        sort_element=pulumi.get(__ret__, 'sort_element'),
        sort_order=pulumi.get(__ret__, 'sort_order'))


@_utilities.lift_output_func(get_rabbitmq_node_list)
def get_rabbitmq_node_list_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRabbitmqNodeListFilterArgs']]]]] = None,
                                  instance_id: Optional[pulumi.Input[str]] = None,
                                  node_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  sort_element: Optional[pulumi.Input[Optional[str]]] = None,
                                  sort_order: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRabbitmqNodeListResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
