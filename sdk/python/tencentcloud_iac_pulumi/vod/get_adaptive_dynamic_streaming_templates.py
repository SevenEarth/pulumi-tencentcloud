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
    'GetAdaptiveDynamicStreamingTemplatesResult',
    'AwaitableGetAdaptiveDynamicStreamingTemplatesResult',
    'get_adaptive_dynamic_streaming_templates',
    'get_adaptive_dynamic_streaming_templates_output',
]

@pulumi.output_type
class GetAdaptiveDynamicStreamingTemplatesResult:
    """
    A collection of values returned by getAdaptiveDynamicStreamingTemplates.
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
        Unique ID of adaptive dynamic streaming template.
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
    def template_lists(self) -> Sequence['outputs.GetAdaptiveDynamicStreamingTemplatesTemplateListResult']:
        """
        A list of adaptive dynamic streaming templates. Each element contains the following attributes:
        """
        return pulumi.get(self, "template_lists")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Template type filter. Valid values: `Preset`,`Custom`. `Preset`: preset template; `Custom`: custom template.
        """
        return pulumi.get(self, "type")


class AwaitableGetAdaptiveDynamicStreamingTemplatesResult(GetAdaptiveDynamicStreamingTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAdaptiveDynamicStreamingTemplatesResult(
            definition=self.definition,
            id=self.id,
            result_output_file=self.result_output_file,
            sub_app_id=self.sub_app_id,
            template_lists=self.template_lists,
            type=self.type)


def get_adaptive_dynamic_streaming_templates(definition: Optional[str] = None,
                                             result_output_file: Optional[str] = None,
                                             sub_app_id: Optional[int] = None,
                                             type: Optional[str] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAdaptiveDynamicStreamingTemplatesResult:
    """
    Use this data source to query detailed information of VOD adaptive dynamic streaming templates.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo_adaptive_dynamic_streaming_template = tencentcloud.vod.AdaptiveDynamicStreamingTemplate("fooAdaptiveDynamicStreamingTemplate",
        format="HLS",
        drm_type="SimpleAES",
        disable_higher_video_bitrate=False,
        disable_higher_video_resolution=False,
        comment="test",
        stream_infos=[
            tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                    codec="libx265",
                    fps=4,
                    bitrate=129,
                    resolution_adaptive=False,
                    width=128,
                    height=128,
                    fill_type="stretch",
                ),
                audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                    codec="libmp3lame",
                    bitrate=129,
                    sample_rate=44100,
                    audio_channel="dual",
                ),
                remove_audio=False,
            ),
            tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                    codec="libx264",
                    fps=4,
                    bitrate=256,
                ),
                audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                    codec="libfdk_aac",
                    bitrate=256,
                    sample_rate=44100,
                ),
                remove_audio=True,
            ),
        ])
    foo_adaptive_dynamic_streaming_templates = tencentcloud.Vod.get_adaptive_dynamic_streaming_templates_output(type="Custom",
        definition=foo_adaptive_dynamic_streaming_template.id)
    ```


    :param str definition: Unique ID filter of adaptive dynamic streaming template.
    :param str result_output_file: Used to save results.
    :param int sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
    :param str type: Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
    """
    __args__ = dict()
    __args__['definition'] = definition
    __args__['resultOutputFile'] = result_output_file
    __args__['subAppId'] = sub_app_id
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vod/getAdaptiveDynamicStreamingTemplates:getAdaptiveDynamicStreamingTemplates', __args__, opts=opts, typ=GetAdaptiveDynamicStreamingTemplatesResult).value

    return AwaitableGetAdaptiveDynamicStreamingTemplatesResult(
        definition=__ret__.definition,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        sub_app_id=__ret__.sub_app_id,
        template_lists=__ret__.template_lists,
        type=__ret__.type)


@_utilities.lift_output_func(get_adaptive_dynamic_streaming_templates)
def get_adaptive_dynamic_streaming_templates_output(definition: Optional[pulumi.Input[Optional[str]]] = None,
                                                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                    sub_app_id: Optional[pulumi.Input[Optional[int]]] = None,
                                                    type: Optional[pulumi.Input[Optional[str]]] = None,
                                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAdaptiveDynamicStreamingTemplatesResult]:
    """
    Use this data source to query detailed information of VOD adaptive dynamic streaming templates.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    foo_adaptive_dynamic_streaming_template = tencentcloud.vod.AdaptiveDynamicStreamingTemplate("fooAdaptiveDynamicStreamingTemplate",
        format="HLS",
        drm_type="SimpleAES",
        disable_higher_video_bitrate=False,
        disable_higher_video_resolution=False,
        comment="test",
        stream_infos=[
            tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                    codec="libx265",
                    fps=4,
                    bitrate=129,
                    resolution_adaptive=False,
                    width=128,
                    height=128,
                    fill_type="stretch",
                ),
                audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                    codec="libmp3lame",
                    bitrate=129,
                    sample_rate=44100,
                    audio_channel="dual",
                ),
                remove_audio=False,
            ),
            tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                    codec="libx264",
                    fps=4,
                    bitrate=256,
                ),
                audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                    codec="libfdk_aac",
                    bitrate=256,
                    sample_rate=44100,
                ),
                remove_audio=True,
            ),
        ])
    foo_adaptive_dynamic_streaming_templates = tencentcloud.Vod.get_adaptive_dynamic_streaming_templates_output(type="Custom",
        definition=foo_adaptive_dynamic_streaming_template.id)
    ```


    :param str definition: Unique ID filter of adaptive dynamic streaming template.
    :param str result_output_file: Used to save results.
    :param int sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
    :param str type: Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
    """
    ...
