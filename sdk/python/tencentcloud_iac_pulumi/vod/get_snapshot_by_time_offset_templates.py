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
    'GetSnapshotByTimeOffsetTemplatesResult',
    'AwaitableGetSnapshotByTimeOffsetTemplatesResult',
    'get_snapshot_by_time_offset_templates',
    'get_snapshot_by_time_offset_templates_output',
]

@pulumi.output_type
class GetSnapshotByTimeOffsetTemplatesResult:
    """
    A collection of values returned by getSnapshotByTimeOffsetTemplates.
    """
    def __init__(__self__, definition=None, id=None, result_output_file=None, sub_app_id=None, template_lists=None, type=None):
        if definition and not isinstance(definition, str):
            raise TypeError("Expected argument 'definition' to be a str")
        pulumi.set(__self__, "definition", definition)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sub_app_id and not isinstance(sub_app_id, int):
            raise TypeError("Expected argument 'sub_app_id' to be a int")
        pulumi.set(__self__, "sub_app_id", sub_app_id)
        if template_lists and not isinstance(template_lists, list):
            raise TypeError("Expected argument 'template_lists' to be a list")
        pulumi.set(__self__, "template_lists", template_lists)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def definition(self) -> Optional[str]:
        """
        Unique ID of snapshot by time offset template.
        """
        return pulumi.get(self, "definition")

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
    @pulumi.getter(name="subAppId")
    def sub_app_id(self) -> Optional[int]:
        return pulumi.get(self, "sub_app_id")

    @property
    @pulumi.getter(name="templateLists")
    def template_lists(self) -> Sequence['outputs.GetSnapshotByTimeOffsetTemplatesTemplateListResult']:
        """
        A list of snapshot by time offset templates. Each element contains the following attributes:
        """
        return pulumi.get(self, "template_lists")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
        """
        return pulumi.get(self, "type")


class AwaitableGetSnapshotByTimeOffsetTemplatesResult(GetSnapshotByTimeOffsetTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotByTimeOffsetTemplatesResult(
            definition=self.definition,
            id=self.id,
            result_output_file=self.result_output_file,
            sub_app_id=self.sub_app_id,
            template_lists=self.template_lists,
            type=self.type)


def get_snapshot_by_time_offset_templates(definition: Optional[str] = None,
                                          result_output_file: Optional[str] = None,
                                          sub_app_id: Optional[int] = None,
                                          type: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotByTimeOffsetTemplatesResult:
    """
    Use this data source to query detailed information of VOD snapshot by time offset templates.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo_snapshot_by_time_offset_template = tencentcloud.vod.SnapshotByTimeOffsetTemplate("fooSnapshotByTimeOffsetTemplate",
        width=130,
        height=128,
        resolution_adaptive=False,
        format="png",
        comment="test",
        fill_type="white")
    foo_snapshot_by_time_offset_templates = tencentcloud.Vod.get_snapshot_by_time_offset_templates_output(type="Custom",
        definition=foo_snapshot_by_time_offset_template.id)
    ```
    <!--End PulumiCodeChooser -->


    :param str definition: Unique ID filter of snapshot by time offset template.
    :param str result_output_file: Used to save results.
    :param int sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
    :param str type: Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
    """
    __args__ = dict()
    __args__['definition'] = definition
    __args__['resultOutputFile'] = result_output_file
    __args__['subAppId'] = sub_app_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vod/getSnapshotByTimeOffsetTemplates:getSnapshotByTimeOffsetTemplates', __args__, opts=opts, typ=GetSnapshotByTimeOffsetTemplatesResult).value

    return AwaitableGetSnapshotByTimeOffsetTemplatesResult(
        definition=pulumi.get(__ret__, 'definition'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        sub_app_id=pulumi.get(__ret__, 'sub_app_id'),
        template_lists=pulumi.get(__ret__, 'template_lists'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_snapshot_by_time_offset_templates)
def get_snapshot_by_time_offset_templates_output(definition: Optional[pulumi.Input[Optional[str]]] = None,
                                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                 sub_app_id: Optional[pulumi.Input[Optional[int]]] = None,
                                                 type: Optional[pulumi.Input[Optional[str]]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotByTimeOffsetTemplatesResult]:
    """
    Use this data source to query detailed information of VOD snapshot by time offset templates.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo_snapshot_by_time_offset_template = tencentcloud.vod.SnapshotByTimeOffsetTemplate("fooSnapshotByTimeOffsetTemplate",
        width=130,
        height=128,
        resolution_adaptive=False,
        format="png",
        comment="test",
        fill_type="white")
    foo_snapshot_by_time_offset_templates = tencentcloud.Vod.get_snapshot_by_time_offset_templates_output(type="Custom",
        definition=foo_snapshot_by_time_offset_template.id)
    ```
    <!--End PulumiCodeChooser -->


    :param str definition: Unique ID filter of snapshot by time offset template.
    :param str result_output_file: Used to save results.
    :param int sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
    :param str type: Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
    """
    ...
