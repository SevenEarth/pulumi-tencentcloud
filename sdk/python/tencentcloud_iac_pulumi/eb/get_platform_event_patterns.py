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
    'GetPlatformEventPatternsResult',
    'AwaitableGetPlatformEventPatternsResult',
    'get_platform_event_patterns',
    'get_platform_event_patterns_output',
]

@pulumi.output_type
class GetPlatformEventPatternsResult:
    """
    A collection of values returned by getPlatformEventPatterns.
    """
    def __init__(__self__, event_patterns=None, id=None, product_type=None, result_output_file=None):
        if event_patterns and not isinstance(event_patterns, list):
            raise TypeError("Expected argument 'event_patterns' to be a list")
        pulumi.set(__self__, "event_patterns", event_patterns)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if product_type and not isinstance(product_type, str):
            raise TypeError("Expected argument 'product_type' to be a str")
        pulumi.set(__self__, "product_type", product_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="eventPatterns")
    def event_patterns(self) -> Sequence['outputs.GetPlatformEventPatternsEventPatternResult']:
        """
        Platform product event matching rules.
        """
        return pulumi.get(self, "event_patterns")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> str:
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetPlatformEventPatternsResult(GetPlatformEventPatternsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlatformEventPatternsResult(
            event_patterns=self.event_patterns,
            id=self.id,
            product_type=self.product_type,
            result_output_file=self.result_output_file)


def get_platform_event_patterns(product_type: Optional[str] = None,
                                result_output_file: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlatformEventPatternsResult:
    """
    Use this data source to query detailed information of eb platform_event_patterns

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    platform_event_patterns = tencentcloud.Eb.get_platform_event_patterns(product_type="")
    ```
    <!--End PulumiCodeChooser -->


    :param str product_type: Platform product type.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['productType'] = product_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Eb/getPlatformEventPatterns:getPlatformEventPatterns', __args__, opts=opts, typ=GetPlatformEventPatternsResult).value

    return AwaitableGetPlatformEventPatternsResult(
        event_patterns=pulumi.get(__ret__, 'event_patterns'),
        id=pulumi.get(__ret__, 'id'),
        product_type=pulumi.get(__ret__, 'product_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_platform_event_patterns)
def get_platform_event_patterns_output(product_type: Optional[pulumi.Input[str]] = None,
                                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlatformEventPatternsResult]:
    """
    Use this data source to query detailed information of eb platform_event_patterns

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    platform_event_patterns = tencentcloud.Eb.get_platform_event_patterns(product_type="")
    ```
    <!--End PulumiCodeChooser -->


    :param str product_type: Platform product type.
    :param str result_output_file: Used to save results.
    """
    ...
