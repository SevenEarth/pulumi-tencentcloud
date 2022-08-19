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
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
    'get_instances_output',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, id=None, instance_id=None, instance_lists=None, instance_name=None, result_output_file=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.GetInstancesInstanceListResult']:
        """
        An information list of elasticsearch instance. Each element contains the following attributes:
        """
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[str]:
        """
        Name of the instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        """
        A mapping of tags to assign to the instance.
        """
        return pulumi.get(self, "tags")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            id=self.id,
            instance_id=self.instance_id,
            instance_lists=self.instance_lists,
            instance_name=self.instance_name,
            result_output_file=self.result_output_file,
            tags=self.tags)


def get_instances(instance_id: Optional[str] = None,
                  instance_name: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  tags: Optional[Mapping[str, Any]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    Use this data source to query elasticsearch instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Elasticsearch.get_instances(instance_id="es-17634f05")
    ```


    :param str instance_id: ID of the instance to be queried.
    :param str instance_name: Name of the instance to be queried.
    :param str result_output_file: Used to save results.
    :param Mapping[str, Any] tags: Tag of the instance to be queried.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['instanceName'] = instance_name
    __args__['resultOutputFile'] = result_output_file
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Elasticsearch/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_lists=__ret__.instance_lists,
        instance_name=__ret__.instance_name,
        result_output_file=__ret__.result_output_file,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_instances)
def get_instances_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                         instance_name: Optional[pulumi.Input[Optional[str]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    Use this data source to query elasticsearch instances.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Elasticsearch.get_instances(instance_id="es-17634f05")
    ```


    :param str instance_id: ID of the instance to be queried.
    :param str instance_name: Name of the instance to be queried.
    :param str result_output_file: Used to save results.
    :param Mapping[str, Any] tags: Tag of the instance to be queried.
    """
    ...
