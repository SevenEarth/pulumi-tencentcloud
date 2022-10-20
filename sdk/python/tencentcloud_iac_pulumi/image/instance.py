# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 image_name: pulumi.Input[str],
                 data_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_poweroff: Optional[pulumi.Input[bool]] = None,
                 image_description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 snapshot_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sysprep: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] image_name: Image name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_disk_ids: Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        :param pulumi.Input[bool] force_poweroff: Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        :param pulumi.Input[str] image_description: Image Description.
        :param pulumi.Input[str] instance_id: Cloud server instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] snapshot_ids: Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        :param pulumi.Input[bool] sysprep: Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        pulumi.set(__self__, "image_name", image_name)
        if data_disk_ids is not None:
            pulumi.set(__self__, "data_disk_ids", data_disk_ids)
        if force_poweroff is not None:
            pulumi.set(__self__, "force_poweroff", force_poweroff)
        if image_description is not None:
            pulumi.set(__self__, "image_description", image_description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if snapshot_ids is not None:
            pulumi.set(__self__, "snapshot_ids", snapshot_ids)
        if sysprep is not None:
            pulumi.set(__self__, "sysprep", sysprep)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Input[str]:
        """
        Image name.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="dataDiskIds")
    def data_disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        """
        return pulumi.get(self, "data_disk_ids")

    @data_disk_ids.setter
    def data_disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "data_disk_ids", value)

    @property
    @pulumi.getter(name="forcePoweroff")
    def force_poweroff(self) -> Optional[pulumi.Input[bool]]:
        """
        Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        """
        return pulumi.get(self, "force_poweroff")

    @force_poweroff.setter
    def force_poweroff(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_poweroff", value)

    @property
    @pulumi.getter(name="imageDescription")
    def image_description(self) -> Optional[pulumi.Input[str]]:
        """
        Image Description.
        """
        return pulumi.get(self, "image_description")

    @image_description.setter
    def image_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud server instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="snapshotIds")
    def snapshot_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        """
        return pulumi.get(self, "snapshot_ids")

    @snapshot_ids.setter
    def snapshot_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "snapshot_ids", value)

    @property
    @pulumi.getter
    def sysprep(self) -> Optional[pulumi.Input[bool]]:
        """
        Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        return pulumi.get(self, "sysprep")

    @sysprep.setter
    def sysprep(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sysprep", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 data_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_poweroff: Optional[pulumi.Input[bool]] = None,
                 image_description: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 snapshot_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sysprep: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_disk_ids: Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        :param pulumi.Input[bool] force_poweroff: Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        :param pulumi.Input[str] image_description: Image Description.
        :param pulumi.Input[str] image_name: Image name.
        :param pulumi.Input[str] instance_id: Cloud server instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] snapshot_ids: Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        :param pulumi.Input[bool] sysprep: Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        if data_disk_ids is not None:
            pulumi.set(__self__, "data_disk_ids", data_disk_ids)
        if force_poweroff is not None:
            pulumi.set(__self__, "force_poweroff", force_poweroff)
        if image_description is not None:
            pulumi.set(__self__, "image_description", image_description)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if snapshot_ids is not None:
            pulumi.set(__self__, "snapshot_ids", snapshot_ids)
        if sysprep is not None:
            pulumi.set(__self__, "sysprep", sysprep)

    @property
    @pulumi.getter(name="dataDiskIds")
    def data_disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        """
        return pulumi.get(self, "data_disk_ids")

    @data_disk_ids.setter
    def data_disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "data_disk_ids", value)

    @property
    @pulumi.getter(name="forcePoweroff")
    def force_poweroff(self) -> Optional[pulumi.Input[bool]]:
        """
        Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        """
        return pulumi.get(self, "force_poweroff")

    @force_poweroff.setter
    def force_poweroff(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_poweroff", value)

    @property
    @pulumi.getter(name="imageDescription")
    def image_description(self) -> Optional[pulumi.Input[str]]:
        """
        Image Description.
        """
        return pulumi.get(self, "image_description")

    @image_description.setter
    def image_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_description", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        Image name.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud server instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="snapshotIds")
    def snapshot_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        """
        return pulumi.get(self, "snapshot_ids")

    @snapshot_ids.setter
    def snapshot_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "snapshot_ids", value)

    @property
    @pulumi.getter
    def sysprep(self) -> Optional[pulumi.Input[bool]]:
        """
        Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        return pulumi.get(self, "sysprep")

    @sysprep.setter
    def sysprep(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sysprep", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_poweroff: Optional[pulumi.Input[bool]] = None,
                 image_description: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 snapshot_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sysprep: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provide a resource to manage image.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        image_snap = tencentcloud.image.Instance("imageSnap",
            force_poweroff=True,
            image_description="create image with snapshot",
            image_name="image-snapshot-keep",
            snapshot_ids=[
                "snap-nbp3xy1d",
                "snap-nvzu3dmh",
            ])
        ```

        ## Import

        image instance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Image/instance:Instance image_snap img-gf7jspk6
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_disk_ids: Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        :param pulumi.Input[bool] force_poweroff: Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        :param pulumi.Input[str] image_description: Image Description.
        :param pulumi.Input[str] image_name: Image name.
        :param pulumi.Input[str] instance_id: Cloud server instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] snapshot_ids: Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        :param pulumi.Input[bool] sysprep: Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to manage image.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        image_snap = tencentcloud.image.Instance("imageSnap",
            force_poweroff=True,
            image_description="create image with snapshot",
            image_name="image-snapshot-keep",
            snapshot_ids=[
                "snap-nbp3xy1d",
                "snap-nvzu3dmh",
            ])
        ```

        ## Import

        image instance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Image/instance:Instance image_snap img-gf7jspk6
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_poweroff: Optional[pulumi.Input[bool]] = None,
                 image_description: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 snapshot_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sysprep: Optional[pulumi.Input[bool]] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["data_disk_ids"] = data_disk_ids
            __props__.__dict__["force_poweroff"] = force_poweroff
            __props__.__dict__["image_description"] = image_description
            if image_name is None and not opts.urn:
                raise TypeError("Missing required property 'image_name'")
            __props__.__dict__["image_name"] = image_name
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["snapshot_ids"] = snapshot_ids
            __props__.__dict__["sysprep"] = sysprep
        super(Instance, __self__).__init__(
            'tencentcloud:Image/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            force_poweroff: Optional[pulumi.Input[bool]] = None,
            image_description: Optional[pulumi.Input[str]] = None,
            image_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            snapshot_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            sysprep: Optional[pulumi.Input[bool]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_disk_ids: Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        :param pulumi.Input[bool] force_poweroff: Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        :param pulumi.Input[str] image_description: Image Description.
        :param pulumi.Input[str] image_name: Image name.
        :param pulumi.Input[str] instance_id: Cloud server instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] snapshot_ids: Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        :param pulumi.Input[bool] sysprep: Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["data_disk_ids"] = data_disk_ids
        __props__.__dict__["force_poweroff"] = force_poweroff
        __props__.__dict__["image_description"] = image_description
        __props__.__dict__["image_name"] = image_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["snapshot_ids"] = snapshot_ids
        __props__.__dict__["sysprep"] = sysprep
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataDiskIds")
    def data_disk_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.
        """
        return pulumi.get(self, "data_disk_ids")

    @property
    @pulumi.getter(name="forcePoweroff")
    def force_poweroff(self) -> pulumi.Output[Optional[bool]]:
        """
        Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror will be made after shutdown.
        """
        return pulumi.get(self, "force_poweroff")

    @property
    @pulumi.getter(name="imageDescription")
    def image_description(self) -> pulumi.Output[Optional[str]]:
        """
        Image Description.
        """
        return pulumi.get(self, "image_description")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Output[str]:
        """
        Image name.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        Cloud server instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="snapshotIds")
    def snapshot_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.
        """
        return pulumi.get(self, "snapshot_ids")

    @property
    @pulumi.getter
    def sysprep(self) -> pulumi.Output[Optional[bool]]:
        """
        Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function.
        """
        return pulumi.get(self, "sysprep")

