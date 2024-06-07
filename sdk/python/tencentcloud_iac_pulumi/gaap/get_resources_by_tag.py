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
    'GetResourcesByTagResult',
    'AwaitableGetResourcesByTagResult',
    'get_resources_by_tag',
    'get_resources_by_tag_output',
]

@pulumi.output_type
class GetResourcesByTagResult:
    """
    A collection of values returned by getResourcesByTag.
    """
    def __init__(__self__, id=None, resource_sets=None, resource_type=None, result_output_file=None, tag_key=None, tag_value=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_sets and not isinstance(resource_sets, list):
            raise TypeError("Expected argument 'resource_sets' to be a list")
        pulumi.set(__self__, "resource_sets", resource_sets)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if tag_key and not isinstance(tag_key, str):
            raise TypeError("Expected argument 'tag_key' to be a str")
        pulumi.set(__self__, "tag_key", tag_key)
        if tag_value and not isinstance(tag_value, str):
            raise TypeError("Expected argument 'tag_value' to be a str")
        pulumi.set(__self__, "tag_value", tag_value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceSets")
    def resource_sets(self) -> Sequence['outputs.GetResourcesByTagResourceSetResult']:
        """
        List of resources corresponding to labels.
        """
        return pulumi.get(self, "resource_sets")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        Resource type, where:Proxy represents the proxy,ProxyGroup represents a proxy group,RealServer represents the real server.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> str:
        return pulumi.get(self, "tag_key")

    @property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> str:
        return pulumi.get(self, "tag_value")


class AwaitableGetResourcesByTagResult(GetResourcesByTagResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourcesByTagResult(
            id=self.id,
            resource_sets=self.resource_sets,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file,
            tag_key=self.tag_key,
            tag_value=self.tag_value)


def get_resources_by_tag(resource_type: Optional[str] = None,
                         result_output_file: Optional[str] = None,
                         tag_key: Optional[str] = None,
                         tag_value: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourcesByTagResult:
    """
    Use this data source to query detailed information of gaap resources by tag

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    resources_by_tag = tencentcloud.Gaap.get_resources_by_tag(tag_key="tagKey",
        tag_value="tagValue")
    ```
    <!--End PulumiCodeChooser -->


    :param str resource_type: Resource type, where:Proxy represents the proxy;ProxyGroup represents a proxy group;RealServer represents the Real Server.If this field is not specified, all resources under the label will be queried.
    :param str result_output_file: Used to save results.
    :param str tag_key: Tag key.
    :param str tag_value: Tag value.
    """
    __args__ = dict()
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    __args__['tagKey'] = tag_key
    __args__['tagValue'] = tag_value
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getResourcesByTag:getResourcesByTag', __args__, opts=opts, typ=GetResourcesByTagResult).value

    return AwaitableGetResourcesByTagResult(
        id=pulumi.get(__ret__, 'id'),
        resource_sets=pulumi.get(__ret__, 'resource_sets'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        tag_key=pulumi.get(__ret__, 'tag_key'),
        tag_value=pulumi.get(__ret__, 'tag_value'))


@_utilities.lift_output_func(get_resources_by_tag)
def get_resources_by_tag_output(resource_type: Optional[pulumi.Input[Optional[str]]] = None,
                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                tag_key: Optional[pulumi.Input[str]] = None,
                                tag_value: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourcesByTagResult]:
    """
    Use this data source to query detailed information of gaap resources by tag

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    resources_by_tag = tencentcloud.Gaap.get_resources_by_tag(tag_key="tagKey",
        tag_value="tagValue")
    ```
    <!--End PulumiCodeChooser -->


    :param str resource_type: Resource type, where:Proxy represents the proxy;ProxyGroup represents a proxy group;RealServer represents the Real Server.If this field is not specified, all resources under the label will be queried.
    :param str result_output_file: Used to save results.
    :param str tag_key: Tag key.
    :param str tag_value: Tag value.
    """
    ...
