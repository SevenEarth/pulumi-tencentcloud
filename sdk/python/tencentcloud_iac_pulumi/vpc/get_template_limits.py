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
    'GetTemplateLimitsResult',
    'AwaitableGetTemplateLimitsResult',
    'get_template_limits',
    'get_template_limits_output',
]

@pulumi.output_type
class GetTemplateLimitsResult:
    """
    A collection of values returned by getTemplateLimits.
    """
    def __init__(__self__, id=None, result_output_file=None, template_limits=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if template_limits and not isinstance(template_limits, list):
            raise TypeError("Expected argument 'template_limits' to be a list")
        pulumi.set(__self__, "template_limits", template_limits)

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
    @pulumi.getter(name="templateLimits")
    def template_limits(self) -> Sequence['outputs.GetTemplateLimitsTemplateLimitResult']:
        """
        template limit.
        """
        return pulumi.get(self, "template_limits")


class AwaitableGetTemplateLimitsResult(GetTemplateLimitsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemplateLimitsResult(
            id=self.id,
            result_output_file=self.result_output_file,
            template_limits=self.template_limits)


def get_template_limits(result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemplateLimitsResult:
    """
    Use this data source to query detailed information of vpc template_limits

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    template_limits = tencentcloud.Vpc.get_template_limits()
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpc/getTemplateLimits:getTemplateLimits', __args__, opts=opts, typ=GetTemplateLimitsResult).value

    return AwaitableGetTemplateLimitsResult(
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        template_limits=pulumi.get(__ret__, 'template_limits'))


@_utilities.lift_output_func(get_template_limits)
def get_template_limits_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemplateLimitsResult]:
    """
    Use this data source to query detailed information of vpc template_limits

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    template_limits = tencentcloud.Vpc.get_template_limits()
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    """
    ...
