# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BindDeviceResourceArgs', 'BindDeviceResource']

@pulumi.input_type
class BindDeviceResourceArgs:
    def __init__(__self__, *,
                 device_id_sets: pulumi.Input[Sequence[pulumi.Input[int]]],
                 resource_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a BindDeviceResource resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] device_id_sets: Asset ID collection.
        :param pulumi.Input[str] resource_id: Bastion host service ID.
        """
        pulumi.set(__self__, "device_id_sets", device_id_sets)
        pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="deviceIdSets")
    def device_id_sets(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Asset ID collection.
        """
        return pulumi.get(self, "device_id_sets")

    @device_id_sets.setter
    def device_id_sets(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "device_id_sets", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        Bastion host service ID.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)


@pulumi.input_type
class _BindDeviceResourceState:
    def __init__(__self__, *,
                 device_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BindDeviceResource resources.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] device_id_sets: Asset ID collection.
        :param pulumi.Input[str] resource_id: Bastion host service ID.
        """
        if device_id_sets is not None:
            pulumi.set(__self__, "device_id_sets", device_id_sets)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="deviceIdSets")
    def device_id_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Asset ID collection.
        """
        return pulumi.get(self, "device_id_sets")

    @device_id_sets.setter
    def device_id_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "device_id_sets", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Bastion host service ID.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)


class BindDeviceResource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dasb bind_device_resource

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.dasb.BindDeviceResource("example",
            device_id_sets=[
                17,
                18,
            ],
            resource_id="bh-saas-weyosfym")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] device_id_sets: Asset ID collection.
        :param pulumi.Input[str] resource_id: Bastion host service ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BindDeviceResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dasb bind_device_resource

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.dasb.BindDeviceResource("example",
            device_id_sets=[
                17,
                18,
            ],
            resource_id="bh-saas-weyosfym")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param BindDeviceResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BindDeviceResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BindDeviceResourceArgs.__new__(BindDeviceResourceArgs)

            if device_id_sets is None and not opts.urn:
                raise TypeError("Missing required property 'device_id_sets'")
            __props__.__dict__["device_id_sets"] = device_id_sets
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
        super(BindDeviceResource, __self__).__init__(
            'tencentcloud:Dasb/bindDeviceResource:BindDeviceResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            resource_id: Optional[pulumi.Input[str]] = None) -> 'BindDeviceResource':
        """
        Get an existing BindDeviceResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] device_id_sets: Asset ID collection.
        :param pulumi.Input[str] resource_id: Bastion host service ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BindDeviceResourceState.__new__(_BindDeviceResourceState)

        __props__.__dict__["device_id_sets"] = device_id_sets
        __props__.__dict__["resource_id"] = resource_id
        return BindDeviceResource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceIdSets")
    def device_id_sets(self) -> pulumi.Output[Sequence[int]]:
        """
        Asset ID collection.
        """
        return pulumi.get(self, "device_id_sets")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        Bastion host service ID.
        """
        return pulumi.get(self, "resource_id")

