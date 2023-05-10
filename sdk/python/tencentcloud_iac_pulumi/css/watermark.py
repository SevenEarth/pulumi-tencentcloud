# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WatermarkArgs', 'Watermark']

@pulumi.input_type
class WatermarkArgs:
    def __init__(__self__, *,
                 picture_url: pulumi.Input[str],
                 watermark_name: pulumi.Input[str],
                 height: Optional[pulumi.Input[int]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 x_position: Optional[pulumi.Input[int]] = None,
                 y_position: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Watermark resource.
        :param pulumi.Input[str] picture_url: watermark url.
        :param pulumi.Input[str] watermark_name: watermark name.
        :param pulumi.Input[int] height: height of the picture.
        :param pulumi.Input[int] width: width of the picture.
        :param pulumi.Input[int] x_position: x position of the picture.
        :param pulumi.Input[int] y_position: y position of the picture.
        """
        pulumi.set(__self__, "picture_url", picture_url)
        pulumi.set(__self__, "watermark_name", watermark_name)
        if height is not None:
            pulumi.set(__self__, "height", height)
        if width is not None:
            pulumi.set(__self__, "width", width)
        if x_position is not None:
            pulumi.set(__self__, "x_position", x_position)
        if y_position is not None:
            pulumi.set(__self__, "y_position", y_position)

    @property
    @pulumi.getter(name="pictureUrl")
    def picture_url(self) -> pulumi.Input[str]:
        """
        watermark url.
        """
        return pulumi.get(self, "picture_url")

    @picture_url.setter
    def picture_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "picture_url", value)

    @property
    @pulumi.getter(name="watermarkName")
    def watermark_name(self) -> pulumi.Input[str]:
        """
        watermark name.
        """
        return pulumi.get(self, "watermark_name")

    @watermark_name.setter
    def watermark_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "watermark_name", value)

    @property
    @pulumi.getter
    def height(self) -> Optional[pulumi.Input[int]]:
        """
        height of the picture.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        width of the picture.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)

    @property
    @pulumi.getter(name="xPosition")
    def x_position(self) -> Optional[pulumi.Input[int]]:
        """
        x position of the picture.
        """
        return pulumi.get(self, "x_position")

    @x_position.setter
    def x_position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "x_position", value)

    @property
    @pulumi.getter(name="yPosition")
    def y_position(self) -> Optional[pulumi.Input[int]]:
        """
        y position of the picture.
        """
        return pulumi.get(self, "y_position")

    @y_position.setter
    def y_position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "y_position", value)


@pulumi.input_type
class _WatermarkState:
    def __init__(__self__, *,
                 height: Optional[pulumi.Input[int]] = None,
                 picture_url: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 watermark_name: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 x_position: Optional[pulumi.Input[int]] = None,
                 y_position: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Watermark resources.
        :param pulumi.Input[int] height: height of the picture.
        :param pulumi.Input[str] picture_url: watermark url.
        :param pulumi.Input[int] status: status. 0: not used, 1: used.
        :param pulumi.Input[str] watermark_name: watermark name.
        :param pulumi.Input[int] width: width of the picture.
        :param pulumi.Input[int] x_position: x position of the picture.
        :param pulumi.Input[int] y_position: y position of the picture.
        """
        if height is not None:
            pulumi.set(__self__, "height", height)
        if picture_url is not None:
            pulumi.set(__self__, "picture_url", picture_url)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if watermark_name is not None:
            pulumi.set(__self__, "watermark_name", watermark_name)
        if width is not None:
            pulumi.set(__self__, "width", width)
        if x_position is not None:
            pulumi.set(__self__, "x_position", x_position)
        if y_position is not None:
            pulumi.set(__self__, "y_position", y_position)

    @property
    @pulumi.getter
    def height(self) -> Optional[pulumi.Input[int]]:
        """
        height of the picture.
        """
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter(name="pictureUrl")
    def picture_url(self) -> Optional[pulumi.Input[str]]:
        """
        watermark url.
        """
        return pulumi.get(self, "picture_url")

    @picture_url.setter
    def picture_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "picture_url", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        status. 0: not used, 1: used.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="watermarkName")
    def watermark_name(self) -> Optional[pulumi.Input[str]]:
        """
        watermark name.
        """
        return pulumi.get(self, "watermark_name")

    @watermark_name.setter
    def watermark_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "watermark_name", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        width of the picture.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)

    @property
    @pulumi.getter(name="xPosition")
    def x_position(self) -> Optional[pulumi.Input[int]]:
        """
        x position of the picture.
        """
        return pulumi.get(self, "x_position")

    @x_position.setter
    def x_position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "x_position", value)

    @property
    @pulumi.getter(name="yPosition")
    def y_position(self) -> Optional[pulumi.Input[int]]:
        """
        y position of the picture.
        """
        return pulumi.get(self, "y_position")

    @y_position.setter
    def y_position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "y_position", value)


class Watermark(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 picture_url: Optional[pulumi.Input[str]] = None,
                 watermark_name: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 x_position: Optional[pulumi.Input[int]] = None,
                 y_position: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a css watermark

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        watermark = tencentcloud.css.Watermark("watermark",
            height=0,
            picture_url="picture_url",
            watermark_name="watermark_name",
            width=0,
            x_position=0,
            y_position=0)
        ```

        ## Import

        css watermark can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Css/watermark:Watermark watermark watermark_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] height: height of the picture.
        :param pulumi.Input[str] picture_url: watermark url.
        :param pulumi.Input[str] watermark_name: watermark name.
        :param pulumi.Input[int] width: width of the picture.
        :param pulumi.Input[int] x_position: x position of the picture.
        :param pulumi.Input[int] y_position: y position of the picture.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WatermarkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a css watermark

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        watermark = tencentcloud.css.Watermark("watermark",
            height=0,
            picture_url="picture_url",
            watermark_name="watermark_name",
            width=0,
            x_position=0,
            y_position=0)
        ```

        ## Import

        css watermark can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Css/watermark:Watermark watermark watermark_id
        ```

        :param str resource_name: The name of the resource.
        :param WatermarkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WatermarkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 height: Optional[pulumi.Input[int]] = None,
                 picture_url: Optional[pulumi.Input[str]] = None,
                 watermark_name: Optional[pulumi.Input[str]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 x_position: Optional[pulumi.Input[int]] = None,
                 y_position: Optional[pulumi.Input[int]] = None,
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
            __props__ = WatermarkArgs.__new__(WatermarkArgs)

            __props__.__dict__["height"] = height
            if picture_url is None and not opts.urn:
                raise TypeError("Missing required property 'picture_url'")
            __props__.__dict__["picture_url"] = picture_url
            if watermark_name is None and not opts.urn:
                raise TypeError("Missing required property 'watermark_name'")
            __props__.__dict__["watermark_name"] = watermark_name
            __props__.__dict__["width"] = width
            __props__.__dict__["x_position"] = x_position
            __props__.__dict__["y_position"] = y_position
            __props__.__dict__["status"] = None
        super(Watermark, __self__).__init__(
            'tencentcloud:Css/watermark:Watermark',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            height: Optional[pulumi.Input[int]] = None,
            picture_url: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None,
            watermark_name: Optional[pulumi.Input[str]] = None,
            width: Optional[pulumi.Input[int]] = None,
            x_position: Optional[pulumi.Input[int]] = None,
            y_position: Optional[pulumi.Input[int]] = None) -> 'Watermark':
        """
        Get an existing Watermark resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] height: height of the picture.
        :param pulumi.Input[str] picture_url: watermark url.
        :param pulumi.Input[int] status: status. 0: not used, 1: used.
        :param pulumi.Input[str] watermark_name: watermark name.
        :param pulumi.Input[int] width: width of the picture.
        :param pulumi.Input[int] x_position: x position of the picture.
        :param pulumi.Input[int] y_position: y position of the picture.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WatermarkState.__new__(_WatermarkState)

        __props__.__dict__["height"] = height
        __props__.__dict__["picture_url"] = picture_url
        __props__.__dict__["status"] = status
        __props__.__dict__["watermark_name"] = watermark_name
        __props__.__dict__["width"] = width
        __props__.__dict__["x_position"] = x_position
        __props__.__dict__["y_position"] = y_position
        return Watermark(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def height(self) -> pulumi.Output[Optional[int]]:
        """
        height of the picture.
        """
        return pulumi.get(self, "height")

    @property
    @pulumi.getter(name="pictureUrl")
    def picture_url(self) -> pulumi.Output[str]:
        """
        watermark url.
        """
        return pulumi.get(self, "picture_url")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        status. 0: not used, 1: used.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="watermarkName")
    def watermark_name(self) -> pulumi.Output[str]:
        """
        watermark name.
        """
        return pulumi.get(self, "watermark_name")

    @property
    @pulumi.getter
    def width(self) -> pulumi.Output[Optional[int]]:
        """
        width of the picture.
        """
        return pulumi.get(self, "width")

    @property
    @pulumi.getter(name="xPosition")
    def x_position(self) -> pulumi.Output[Optional[int]]:
        """
        x position of the picture.
        """
        return pulumi.get(self, "x_position")

    @property
    @pulumi.getter(name="yPosition")
    def y_position(self) -> pulumi.Output[Optional[int]]:
        """
        y position of the picture.
        """
        return pulumi.get(self, "y_position")

