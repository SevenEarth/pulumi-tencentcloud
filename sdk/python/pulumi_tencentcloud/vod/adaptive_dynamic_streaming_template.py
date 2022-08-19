# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AdaptiveDynamicStreamingTemplateArgs', 'AdaptiveDynamicStreamingTemplate']

@pulumi.input_type
class AdaptiveDynamicStreamingTemplateArgs:
    def __init__(__self__, *,
                 format: pulumi.Input[str],
                 stream_infos: pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]],
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_higher_video_bitrate: Optional[pulumi.Input[bool]] = None,
                 disable_higher_video_resolution: Optional[pulumi.Input[bool]] = None,
                 drm_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sub_app_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a AdaptiveDynamicStreamingTemplate resource.
        :param pulumi.Input[str] format: Adaptive bitstream format. Valid values: `HLS`.
        :param pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]] stream_infos: List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        :param pulumi.Input[str] comment: Template description. Length limit: 256 characters.
        :param pulumi.Input[bool] disable_higher_video_bitrate: Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[bool] disable_higher_video_resolution: Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[str] drm_type: DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        :param pulumi.Input[str] name: Template name. Length limit: 64 characters.
        :param pulumi.Input[int] sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        """
        pulumi.set(__self__, "format", format)
        pulumi.set(__self__, "stream_infos", stream_infos)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if disable_higher_video_bitrate is not None:
            pulumi.set(__self__, "disable_higher_video_bitrate", disable_higher_video_bitrate)
        if disable_higher_video_resolution is not None:
            pulumi.set(__self__, "disable_higher_video_resolution", disable_higher_video_resolution)
        if drm_type is not None:
            pulumi.set(__self__, "drm_type", drm_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sub_app_id is not None:
            pulumi.set(__self__, "sub_app_id", sub_app_id)

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        """
        Adaptive bitstream format. Valid values: `HLS`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: pulumi.Input[str]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="streamInfos")
    def stream_infos(self) -> pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]:
        """
        List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        """
        return pulumi.get(self, "stream_infos")

    @stream_infos.setter
    def stream_infos(self, value: pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]):
        pulumi.set(self, "stream_infos", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Template description. Length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="disableHigherVideoBitrate")
    def disable_higher_video_bitrate(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        """
        return pulumi.get(self, "disable_higher_video_bitrate")

    @disable_higher_video_bitrate.setter
    def disable_higher_video_bitrate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_higher_video_bitrate", value)

    @property
    @pulumi.getter(name="disableHigherVideoResolution")
    def disable_higher_video_resolution(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        """
        return pulumi.get(self, "disable_higher_video_resolution")

    @disable_higher_video_resolution.setter
    def disable_higher_video_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_higher_video_resolution", value)

    @property
    @pulumi.getter(name="drmType")
    def drm_type(self) -> Optional[pulumi.Input[str]]:
        """
        DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        """
        return pulumi.get(self, "drm_type")

    @drm_type.setter
    def drm_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drm_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Template name. Length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="subAppId")
    def sub_app_id(self) -> Optional[pulumi.Input[int]]:
        """
        Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        """
        return pulumi.get(self, "sub_app_id")

    @sub_app_id.setter
    def sub_app_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sub_app_id", value)


@pulumi.input_type
class _AdaptiveDynamicStreamingTemplateState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 disable_higher_video_bitrate: Optional[pulumi.Input[bool]] = None,
                 disable_higher_video_resolution: Optional[pulumi.Input[bool]] = None,
                 drm_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_infos: Optional[pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]] = None,
                 sub_app_id: Optional[pulumi.Input[int]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AdaptiveDynamicStreamingTemplate resources.
        :param pulumi.Input[str] comment: Template description. Length limit: 256 characters.
        :param pulumi.Input[str] create_time: Creation time of template in ISO date format.
        :param pulumi.Input[bool] disable_higher_video_bitrate: Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[bool] disable_higher_video_resolution: Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[str] drm_type: DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        :param pulumi.Input[str] format: Adaptive bitstream format. Valid values: `HLS`.
        :param pulumi.Input[str] name: Template name. Length limit: 64 characters.
        :param pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]] stream_infos: List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        :param pulumi.Input[int] sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        :param pulumi.Input[str] update_time: Last modified time of template in ISO date format.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if disable_higher_video_bitrate is not None:
            pulumi.set(__self__, "disable_higher_video_bitrate", disable_higher_video_bitrate)
        if disable_higher_video_resolution is not None:
            pulumi.set(__self__, "disable_higher_video_resolution", disable_higher_video_resolution)
        if drm_type is not None:
            pulumi.set(__self__, "drm_type", drm_type)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if stream_infos is not None:
            pulumi.set(__self__, "stream_infos", stream_infos)
        if sub_app_id is not None:
            pulumi.set(__self__, "sub_app_id", sub_app_id)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Template description. Length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of template in ISO date format.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="disableHigherVideoBitrate")
    def disable_higher_video_bitrate(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        """
        return pulumi.get(self, "disable_higher_video_bitrate")

    @disable_higher_video_bitrate.setter
    def disable_higher_video_bitrate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_higher_video_bitrate", value)

    @property
    @pulumi.getter(name="disableHigherVideoResolution")
    def disable_higher_video_resolution(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        """
        return pulumi.get(self, "disable_higher_video_resolution")

    @disable_higher_video_resolution.setter
    def disable_higher_video_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_higher_video_resolution", value)

    @property
    @pulumi.getter(name="drmType")
    def drm_type(self) -> Optional[pulumi.Input[str]]:
        """
        DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        """
        return pulumi.get(self, "drm_type")

    @drm_type.setter
    def drm_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drm_type", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Adaptive bitstream format. Valid values: `HLS`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Template name. Length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="streamInfos")
    def stream_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]]:
        """
        List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        """
        return pulumi.get(self, "stream_infos")

    @stream_infos.setter
    def stream_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]]):
        pulumi.set(self, "stream_infos", value)

    @property
    @pulumi.getter(name="subAppId")
    def sub_app_id(self) -> Optional[pulumi.Input[int]]:
        """
        Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        """
        return pulumi.get(self, "sub_app_id")

    @sub_app_id.setter
    def sub_app_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sub_app_id", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last modified time of template in ISO date format.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class AdaptiveDynamicStreamingTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_higher_video_bitrate: Optional[pulumi.Input[bool]] = None,
                 disable_higher_video_resolution: Optional[pulumi.Input[bool]] = None,
                 drm_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]]] = None,
                 sub_app_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provide a resource to create a VOD adaptive dynamic streaming template.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo = tencentcloud.vod.AdaptiveDynamicStreamingTemplate("foo",
            comment="test",
            disable_higher_video_bitrate=False,
            disable_higher_video_resolution=False,
            drm_type="SimpleAES",
            format="HLS",
            stream_infos=[
                tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                    audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                        audio_channel="dual",
                        bitrate=129,
                        codec="libmp3lame",
                        sample_rate=44100,
                    ),
                    remove_audio=False,
                    video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                        bitrate=129,
                        codec="libx265",
                        fill_type="stretch",
                        fps=4,
                        height=128,
                        resolution_adaptive=False,
                        width=128,
                    ),
                ),
                tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                    audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                        bitrate=256,
                        codec="libfdk_aac",
                        sample_rate=44100,
                    ),
                    remove_audio=True,
                    video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                        bitrate=256,
                        codec="libx264",
                        fps=4,
                    ),
                ),
            ])
        ```

        ## Import

        VOD adaptive dynamic streaming template can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate foo 169141
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Template description. Length limit: 256 characters.
        :param pulumi.Input[bool] disable_higher_video_bitrate: Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[bool] disable_higher_video_resolution: Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[str] drm_type: DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        :param pulumi.Input[str] format: Adaptive bitstream format. Valid values: `HLS`.
        :param pulumi.Input[str] name: Template name. Length limit: 64 characters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]] stream_infos: List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        :param pulumi.Input[int] sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AdaptiveDynamicStreamingTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a VOD adaptive dynamic streaming template.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo = tencentcloud.vod.AdaptiveDynamicStreamingTemplate("foo",
            comment="test",
            disable_higher_video_bitrate=False,
            disable_higher_video_resolution=False,
            drm_type="SimpleAES",
            format="HLS",
            stream_infos=[
                tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                    audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                        audio_channel="dual",
                        bitrate=129,
                        codec="libmp3lame",
                        sample_rate=44100,
                    ),
                    remove_audio=False,
                    video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                        bitrate=129,
                        codec="libx265",
                        fill_type="stretch",
                        fps=4,
                        height=128,
                        resolution_adaptive=False,
                        width=128,
                    ),
                ),
                tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoArgs(
                    audio=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs(
                        bitrate=256,
                        codec="libfdk_aac",
                        sample_rate=44100,
                    ),
                    remove_audio=True,
                    video=tencentcloud.vod.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs(
                        bitrate=256,
                        codec="libx264",
                        fps=4,
                    ),
                ),
            ])
        ```

        ## Import

        VOD adaptive dynamic streaming template can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate foo 169141
        ```

        :param str resource_name: The name of the resource.
        :param AdaptiveDynamicStreamingTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AdaptiveDynamicStreamingTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 disable_higher_video_bitrate: Optional[pulumi.Input[bool]] = None,
                 disable_higher_video_resolution: Optional[pulumi.Input[bool]] = None,
                 drm_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]]] = None,
                 sub_app_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AdaptiveDynamicStreamingTemplateArgs.__new__(AdaptiveDynamicStreamingTemplateArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["disable_higher_video_bitrate"] = disable_higher_video_bitrate
            __props__.__dict__["disable_higher_video_resolution"] = disable_higher_video_resolution
            __props__.__dict__["drm_type"] = drm_type
            if format is None and not opts.urn:
                raise TypeError("Missing required property 'format'")
            __props__.__dict__["format"] = format
            __props__.__dict__["name"] = name
            if stream_infos is None and not opts.urn:
                raise TypeError("Missing required property 'stream_infos'")
            __props__.__dict__["stream_infos"] = stream_infos
            __props__.__dict__["sub_app_id"] = sub_app_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(AdaptiveDynamicStreamingTemplate, __self__).__init__(
            'tencentcloud:Vod/adaptiveDynamicStreamingTemplate:AdaptiveDynamicStreamingTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            disable_higher_video_bitrate: Optional[pulumi.Input[bool]] = None,
            disable_higher_video_resolution: Optional[pulumi.Input[bool]] = None,
            drm_type: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            stream_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]]] = None,
            sub_app_id: Optional[pulumi.Input[int]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'AdaptiveDynamicStreamingTemplate':
        """
        Get an existing AdaptiveDynamicStreamingTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Template description. Length limit: 256 characters.
        :param pulumi.Input[str] create_time: Creation time of template in ISO date format.
        :param pulumi.Input[bool] disable_higher_video_bitrate: Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[bool] disable_higher_video_resolution: Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        :param pulumi.Input[str] drm_type: DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        :param pulumi.Input[str] format: Adaptive bitstream format. Valid values: `HLS`.
        :param pulumi.Input[str] name: Template name. Length limit: 64 characters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdaptiveDynamicStreamingTemplateStreamInfoArgs']]]] stream_infos: List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        :param pulumi.Input[int] sub_app_id: Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        :param pulumi.Input[str] update_time: Last modified time of template in ISO date format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AdaptiveDynamicStreamingTemplateState.__new__(_AdaptiveDynamicStreamingTemplateState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["disable_higher_video_bitrate"] = disable_higher_video_bitrate
        __props__.__dict__["disable_higher_video_resolution"] = disable_higher_video_resolution
        __props__.__dict__["drm_type"] = drm_type
        __props__.__dict__["format"] = format
        __props__.__dict__["name"] = name
        __props__.__dict__["stream_infos"] = stream_infos
        __props__.__dict__["sub_app_id"] = sub_app_id
        __props__.__dict__["update_time"] = update_time
        return AdaptiveDynamicStreamingTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Template description. Length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of template in ISO date format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="disableHigherVideoBitrate")
    def disable_higher_video_bitrate(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        """
        return pulumi.get(self, "disable_higher_video_bitrate")

    @property
    @pulumi.getter(name="disableHigherVideoResolution")
    def disable_higher_video_resolution(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to prohibit transcoding from low resolution to high resolution. Valid values: `false`,`true`. `false`: no, `true`: yes. Default value: `false`.
        """
        return pulumi.get(self, "disable_higher_video_resolution")

    @property
    @pulumi.getter(name="drmType")
    def drm_type(self) -> pulumi.Output[Optional[str]]:
        """
        DRM scheme type. Valid values: `SimpleAES`. If this field is an empty string, DRM will not be performed on the video.
        """
        return pulumi.get(self, "drm_type")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        """
        Adaptive bitstream format. Valid values: `HLS`.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Template name. Length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="streamInfos")
    def stream_infos(self) -> pulumi.Output[Sequence['outputs.AdaptiveDynamicStreamingTemplateStreamInfo']]:
        """
        List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.
        """
        return pulumi.get(self, "stream_infos")

    @property
    @pulumi.getter(name="subAppId")
    def sub_app_id(self) -> pulumi.Output[Optional[int]]:
        """
        Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        """
        return pulumi.get(self, "sub_app_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last modified time of template in ISO date format.
        """
        return pulumi.get(self, "update_time")

