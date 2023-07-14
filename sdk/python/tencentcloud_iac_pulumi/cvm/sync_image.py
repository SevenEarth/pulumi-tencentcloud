# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncImageArgs', 'SyncImage']

@pulumi.input_type
class SyncImageArgs:
    def __init__(__self__, *,
                 destination_regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 image_id: pulumi.Input[str],
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 image_set_required: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SyncImage resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_regions: List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        :param pulumi.Input[str] image_id: Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        :param pulumi.Input[bool] dry_run: Checks whether image synchronization can be initiated.
        :param pulumi.Input[str] image_name: Destination image name.
        :param pulumi.Input[bool] image_set_required: Whether to return the ID of image created in the destination region.
        """
        pulumi.set(__self__, "destination_regions", destination_regions)
        pulumi.set(__self__, "image_id", image_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if image_set_required is not None:
            pulumi.set(__self__, "image_set_required", image_set_required)

    @property
    @pulumi.getter(name="destinationRegions")
    def destination_regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        """
        return pulumi.get(self, "destination_regions")

    @destination_regions.setter
    def destination_regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "destination_regions", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[str]:
        """
        Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Checks whether image synchronization can be initiated.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        Destination image name.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="imageSetRequired")
    def image_set_required(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to return the ID of image created in the destination region.
        """
        return pulumi.get(self, "image_set_required")

    @image_set_required.setter
    def image_set_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "image_set_required", value)


@pulumi.input_type
class _SyncImageState:
    def __init__(__self__, *,
                 destination_regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 image_set_required: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering SyncImage resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_regions: List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        :param pulumi.Input[bool] dry_run: Checks whether image synchronization can be initiated.
        :param pulumi.Input[str] image_id: Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        :param pulumi.Input[str] image_name: Destination image name.
        :param pulumi.Input[bool] image_set_required: Whether to return the ID of image created in the destination region.
        """
        if destination_regions is not None:
            pulumi.set(__self__, "destination_regions", destination_regions)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if image_set_required is not None:
            pulumi.set(__self__, "image_set_required", image_set_required)

    @property
    @pulumi.getter(name="destinationRegions")
    def destination_regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        """
        return pulumi.get(self, "destination_regions")

    @destination_regions.setter
    def destination_regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "destination_regions", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Checks whether image synchronization can be initiated.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        Destination image name.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="imageSetRequired")
    def image_set_required(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to return the ID of image created in the destination region.
        """
        return pulumi.get(self, "image_set_required")

    @image_set_required.setter
    def image_set_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "image_set_required", value)


class SyncImage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 image_set_required: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a resource to create a cvm sync_image

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sync_image = tencentcloud.cvm.SyncImage("syncImage",
            destination_regions=[
                "ap-guangzhou",
                "ap-shanghai",
            ],
            image_id="img-xxxxxx")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_regions: List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        :param pulumi.Input[bool] dry_run: Checks whether image synchronization can be initiated.
        :param pulumi.Input[str] image_id: Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        :param pulumi.Input[str] image_name: Destination image name.
        :param pulumi.Input[bool] image_set_required: Whether to return the ID of image created in the destination region.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyncImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cvm sync_image

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sync_image = tencentcloud.cvm.SyncImage("syncImage",
            destination_regions=[
                "ap-guangzhou",
                "ap-shanghai",
            ],
            image_id="img-xxxxxx")
        ```

        :param str resource_name: The name of the resource.
        :param SyncImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 image_set_required: Optional[pulumi.Input[bool]] = None,
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
            __props__ = SyncImageArgs.__new__(SyncImageArgs)

            if destination_regions is None and not opts.urn:
                raise TypeError("Missing required property 'destination_regions'")
            __props__.__dict__["destination_regions"] = destination_regions
            __props__.__dict__["dry_run"] = dry_run
            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["image_name"] = image_name
            __props__.__dict__["image_set_required"] = image_set_required
        super(SyncImage, __self__).__init__(
            'tencentcloud:Cvm/syncImage:SyncImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            image_name: Optional[pulumi.Input[str]] = None,
            image_set_required: Optional[pulumi.Input[bool]] = None) -> 'SyncImage':
        """
        Get an existing SyncImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_regions: List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        :param pulumi.Input[bool] dry_run: Checks whether image synchronization can be initiated.
        :param pulumi.Input[str] image_id: Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        :param pulumi.Input[str] image_name: Destination image name.
        :param pulumi.Input[bool] image_set_required: Whether to return the ID of image created in the destination region.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncImageState.__new__(_SyncImageState)

        __props__.__dict__["destination_regions"] = destination_regions
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["image_name"] = image_name
        __props__.__dict__["image_set_required"] = image_set_required
        return SyncImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationRegions")
    def destination_regions(self) -> pulumi.Output[Sequence[str]]:
        """
        List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        """
        return pulumi.get(self, "destination_regions")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Checks whether image synchronization can be initiated.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Output[Optional[str]]:
        """
        Destination image name.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="imageSetRequired")
    def image_set_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to return the ID of image created in the destination region.
        """
        return pulumi.get(self, "image_set_required")

