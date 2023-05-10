# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnapshotByTimeoffsetTemplateArgs', 'SnapshotByTimeoffsetTemplate']

@pulumi.input_type
class SnapshotByTimeoffsetTemplateArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 fill_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resolution_adaptive: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SnapshotByTimeoffsetTemplate resource.
        :param pulumi.Input[str] comment: Template description information, length limit: 256 characters.
        :param pulumi.Input[str] fill_type: Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        :param pulumi.Input[str] format: Image format, the value can be jpg, png, webp. Default is jpg.
        :param pulumi.Input[int] height: The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        :param pulumi.Input[str] name: Snapshot by timeoffset template name, length limit: 64 characters.
        :param pulumi.Input[str] resolution_adaptive: Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        :param pulumi.Input[int] width: The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if fill_type is not None:
            pulumi.set(__self__, "fill_type", fill_type)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if height is not None:
            pulumi.set(__self__, "height", height)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resolution_adaptive is not None:
            pulumi.set(__self__, "resolution_adaptive", resolution_adaptive)
        if width is not None:
            pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Template description information, length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="fillType")
    def fill_type(self) -> Optional[pulumi.Input[str]]:
        """
        Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        """
        return pulumi.get(self, "fill_type")

    @fill_type.setter
    def fill_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fill_type", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Image format, the value can be jpg, png, webp. Default is jpg.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def height(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Snapshot by timeoffset template name, length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resolutionAdaptive")
    def resolution_adaptive(self) -> Optional[pulumi.Input[str]]:
        """
        Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        """
        return pulumi.get(self, "resolution_adaptive")

    @resolution_adaptive.setter
    def resolution_adaptive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolution_adaptive", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)


@pulumi.input_type
class _SnapshotByTimeoffsetTemplateState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 fill_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resolution_adaptive: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SnapshotByTimeoffsetTemplate resources.
        :param pulumi.Input[str] comment: Template description information, length limit: 256 characters.
        :param pulumi.Input[str] fill_type: Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        :param pulumi.Input[str] format: Image format, the value can be jpg, png, webp. Default is jpg.
        :param pulumi.Input[int] height: The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        :param pulumi.Input[str] name: Snapshot by timeoffset template name, length limit: 64 characters.
        :param pulumi.Input[str] resolution_adaptive: Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        :param pulumi.Input[int] width: The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if fill_type is not None:
            pulumi.set(__self__, "fill_type", fill_type)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if height is not None:
            pulumi.set(__self__, "height", height)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resolution_adaptive is not None:
            pulumi.set(__self__, "resolution_adaptive", resolution_adaptive)
        if width is not None:
            pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Template description information, length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="fillType")
    def fill_type(self) -> Optional[pulumi.Input[str]]:
        """
        Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        """
        return pulumi.get(self, "fill_type")

    @fill_type.setter
    def fill_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fill_type", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        Image format, the value can be jpg, png, webp. Default is jpg.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def height(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Snapshot by timeoffset template name, length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resolutionAdaptive")
    def resolution_adaptive(self) -> Optional[pulumi.Input[str]]:
        """
        Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        """
        return pulumi.get(self, "resolution_adaptive")

    @resolution_adaptive.setter
    def resolution_adaptive(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolution_adaptive", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)


class SnapshotByTimeoffsetTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fill_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resolution_adaptive: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a mps snapshot_by_timeoffset_template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        snapshot_by_timeoffset_template = tencentcloud.mps.SnapshotByTimeoffsetTemplate("snapshotByTimeoffsetTemplate",
            fill_type="stretch",
            format="jpg",
            height=128,
            resolution_adaptive="open",
            width=140)
        ```

        ## Import

        mps snapshot_by_timeoffset_template can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate snapshot_by_timeoffset_template snapshot_by_timeoffset_template_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Template description information, length limit: 256 characters.
        :param pulumi.Input[str] fill_type: Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        :param pulumi.Input[str] format: Image format, the value can be jpg, png, webp. Default is jpg.
        :param pulumi.Input[int] height: The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        :param pulumi.Input[str] name: Snapshot by timeoffset template name, length limit: 64 characters.
        :param pulumi.Input[str] resolution_adaptive: Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        :param pulumi.Input[int] width: The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SnapshotByTimeoffsetTemplateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mps snapshot_by_timeoffset_template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        snapshot_by_timeoffset_template = tencentcloud.mps.SnapshotByTimeoffsetTemplate("snapshotByTimeoffsetTemplate",
            fill_type="stretch",
            format="jpg",
            height=128,
            resolution_adaptive="open",
            width=140)
        ```

        ## Import

        mps snapshot_by_timeoffset_template can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate snapshot_by_timeoffset_template snapshot_by_timeoffset_template_id
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotByTimeoffsetTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotByTimeoffsetTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fill_type: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resolution_adaptive: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotByTimeoffsetTemplateArgs.__new__(SnapshotByTimeoffsetTemplateArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["fill_type"] = fill_type
            __props__.__dict__["format"] = format
            __props__.__dict__["height"] = height
            __props__.__dict__["name"] = name
            __props__.__dict__["resolution_adaptive"] = resolution_adaptive
            __props__.__dict__["width"] = width
        super(SnapshotByTimeoffsetTemplate, __self__).__init__(
            'tencentcloud:Mps/snapshotByTimeoffsetTemplate:SnapshotByTimeoffsetTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            fill_type: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            height: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resolution_adaptive: Optional[pulumi.Input[str]] = None,
            width: Optional[pulumi.Input[int]] = None) -> 'SnapshotByTimeoffsetTemplate':
        """
        Get an existing SnapshotByTimeoffsetTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Template description information, length limit: 256 characters.
        :param pulumi.Input[str] fill_type: Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        :param pulumi.Input[str] format: Image format, the value can be jpg, png, webp. Default is jpg.
        :param pulumi.Input[int] height: The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        :param pulumi.Input[str] name: Snapshot by timeoffset template name, length limit: 64 characters.
        :param pulumi.Input[str] resolution_adaptive: Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        :param pulumi.Input[int] width: The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotByTimeoffsetTemplateState.__new__(_SnapshotByTimeoffsetTemplateState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["fill_type"] = fill_type
        __props__.__dict__["format"] = format
        __props__.__dict__["height"] = height
        __props__.__dict__["name"] = name
        __props__.__dict__["resolution_adaptive"] = resolution_adaptive
        __props__.__dict__["width"] = width
        return SnapshotByTimeoffsetTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Template description information, length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="fillType")
    def fill_type(self) -> pulumi.Output[Optional[str]]:
        """
        Filling type, when the aspect ratio of the video stream configuration is inconsistent with the aspect ratio of the original video, the processing method for transcoding is filling. Optional filling type:stretch: Stretching, stretching each frame to fill the entire screen, which may cause the transcoded video to be squashed or stretched.black: Leave black, keep the video aspect ratio unchanged, and fill the rest of the edge with black.white: Leave blank, keep the aspect ratio of the video, and fill the rest of the edge with white.gauss: Gaussian blur, keep the aspect ratio of the video unchanged, and use Gaussian blur for the rest of the edge.Default value: black.
        """
        return pulumi.get(self, "fill_type")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[Optional[str]]:
        """
        Image format, the value can be jpg, png, webp. Default is jpg.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def height(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum value of the snapshot height (or short side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        return pulumi.get(self, "height")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Snapshot by timeoffset template name, length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resolutionAdaptive")
    def resolution_adaptive(self) -> pulumi.Output[Optional[str]]:
        """
        Adaptive resolution, optional value:open: At this time, Width represents the long side of the video, Height represents the short side of the video.close: At this point, Width represents the width of the video, and Height represents the height of the video.Default value: open.
        """
        return pulumi.get(self, "resolution_adaptive")

    @property
    @pulumi.getter
    def width(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum value of the snapshot width (or long side), value range: 0 and [128, 4096], unit: px.When Width and Height are both 0, the resolution is the same.When Width is 0 and Height is not 0, Width is scaled proportionally.When Width is not 0 and Height is 0, Height is scaled proportionally.When both Width and Height are not 0, the resolution is specified by the user.Default value: 0.
        """
        return pulumi.get(self, "width")

