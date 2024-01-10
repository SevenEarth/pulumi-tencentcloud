# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPlateformEventTemplateResult',
    'AwaitableGetPlateformEventTemplateResult',
    'get_plateform_event_template',
    'get_plateform_event_template_output',
]

@pulumi.output_type
class GetPlateformEventTemplateResult:
    """
    A collection of values returned by getPlateformEventTemplate.
    """
    def __init__(__self__, event_template=None, event_type=None, id=None, result_output_file=None):
        if event_template and not isinstance(event_template, str):
            raise TypeError("Expected argument 'event_template' to be a str")
        pulumi.set(__self__, "event_template", event_template)
        if event_type and not isinstance(event_type, str):
            raise TypeError("Expected argument 'event_type' to be a str")
        pulumi.set(__self__, "event_type", event_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="eventTemplate")
    def event_template(self) -> str:
        """
        Platform product event template.
        """
        return pulumi.get(self, "event_template")

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> str:
        return pulumi.get(self, "event_type")

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


class AwaitableGetPlateformEventTemplateResult(GetPlateformEventTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlateformEventTemplateResult(
            event_template=self.event_template,
            event_type=self.event_type,
            id=self.id,
            result_output_file=self.result_output_file)


def get_plateform_event_template(event_type: Optional[str] = None,
                                 result_output_file: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlateformEventTemplateResult:
    """
    Use this data source to query detailed information of eb plateform_event_template

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    plateform_event_template = tencentcloud.Eb.get_plateform_event_template(event_type="eb_platform_test:TEST:ALL")
    ```


    :param str event_type: Platform product event type.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['eventType'] = event_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Eb/getPlateformEventTemplate:getPlateformEventTemplate', __args__, opts=opts, typ=GetPlateformEventTemplateResult).value

    return AwaitableGetPlateformEventTemplateResult(
        event_template=__ret__.event_template,
        event_type=__ret__.event_type,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_plateform_event_template)
def get_plateform_event_template_output(event_type: Optional[pulumi.Input[str]] = None,
                                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlateformEventTemplateResult]:
    """
    Use this data source to query detailed information of eb plateform_event_template

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    plateform_event_template = tencentcloud.Eb.get_plateform_event_template(event_type="eb_platform_test:TEST:ALL")
    ```


    :param str event_type: Platform product event type.
    :param str result_output_file: Used to save results.
    """
    ...
