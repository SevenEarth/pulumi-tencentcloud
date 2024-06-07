# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AttachmentArgs', 'Attachment']

@pulumi.input_type
class AttachmentArgs:
    def __init__(__self__, *,
                 instance_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 scaling_group_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Attachment resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: ID list of CVM instances to be attached to the scaling group.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        pulumi.set(__self__, "instance_ids", instance_ids)
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        ID list of CVM instances to be attached to the scaling group.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Input[str]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_group_id", value)


@pulumi.input_type
class _AttachmentState:
    def __init__(__self__, *,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Attachment resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: ID list of CVM instances to be attached to the scaling group.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        if instance_ids is not None:
            pulumi.set(__self__, "instance_ids", instance_ids)
        if scaling_group_id is not None:
            pulumi.set(__self__, "scaling_group_id", scaling_group_id)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ID list of CVM instances to be attached to the scaling group.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_ids", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_group_id", value)


class Attachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to attach or detach CVM instances to a specified scaling group.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="as")
        image = tencentcloud.Images.get_instance(image_types=["PUBLIC_IMAGE"],
            os_name="TencentOS Server 3.2 (Final)")
        instance_types = tencentcloud.Instance.get_types(filters=[
                tencentcloud.instance.GetTypesFilterArgs(
                    name="zone",
                    values=[zones.zones[0].name],
                ),
                tencentcloud.instance.GetTypesFilterArgs(
                    name="instance-family",
                    values=["S5"],
                ),
            ],
            cpu_core_count=2,
            exclude_sold_out=True)
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            availability_zone=zones.zones[0].name)
        example_scaling_config = tencentcloud.as_.ScalingConfig("exampleScalingConfig",
            configuration_name="tf-example",
            image_id=image.images[0].image_id,
            instance_types=[
                "SA1.SMALL1",
                "SA2.SMALL1",
                "SA2.SMALL2",
                "SA2.SMALL4",
            ],
            instance_name_settings=tencentcloud.as_.ScalingConfigInstanceNameSettingsArgs(
                instance_name="test-ins-name",
            ))
        example_scaling_group = tencentcloud.as_.ScalingGroup("exampleScalingGroup",
            scaling_group_name="tf-example",
            configuration_id=example_scaling_config.id,
            max_size=1,
            min_size=0,
            vpc_id=vpc.id,
            subnet_ids=[subnet.id])
        example_instance = tencentcloud.instance.Instance("exampleInstance",
            instance_name="tf_example_instance",
            availability_zone=zones.zones[0].name,
            image_id=image.images[0].image_id,
            instance_type=instance_types.instance_types[0].instance_type,
            system_disk_type="CLOUD_PREMIUM",
            system_disk_size=50,
            allocate_public_ip=True,
            internet_max_bandwidth_out=10,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        attachment = tencentcloud.as_.Attachment("attachment",
            scaling_group_id=example_scaling_group.id,
            instance_ids=[example_instance.id])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: ID list of CVM instances to be attached to the scaling group.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to attach or detach CVM instances to a specified scaling group.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="as")
        image = tencentcloud.Images.get_instance(image_types=["PUBLIC_IMAGE"],
            os_name="TencentOS Server 3.2 (Final)")
        instance_types = tencentcloud.Instance.get_types(filters=[
                tencentcloud.instance.GetTypesFilterArgs(
                    name="zone",
                    values=[zones.zones[0].name],
                ),
                tencentcloud.instance.GetTypesFilterArgs(
                    name="instance-family",
                    values=["S5"],
                ),
            ],
            cpu_core_count=2,
            exclude_sold_out=True)
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            availability_zone=zones.zones[0].name)
        example_scaling_config = tencentcloud.as_.ScalingConfig("exampleScalingConfig",
            configuration_name="tf-example",
            image_id=image.images[0].image_id,
            instance_types=[
                "SA1.SMALL1",
                "SA2.SMALL1",
                "SA2.SMALL2",
                "SA2.SMALL4",
            ],
            instance_name_settings=tencentcloud.as_.ScalingConfigInstanceNameSettingsArgs(
                instance_name="test-ins-name",
            ))
        example_scaling_group = tencentcloud.as_.ScalingGroup("exampleScalingGroup",
            scaling_group_name="tf-example",
            configuration_id=example_scaling_config.id,
            max_size=1,
            min_size=0,
            vpc_id=vpc.id,
            subnet_ids=[subnet.id])
        example_instance = tencentcloud.instance.Instance("exampleInstance",
            instance_name="tf_example_instance",
            availability_zone=zones.zones[0].name,
            image_id=image.images[0].image_id,
            instance_type=instance_types.instance_types[0].instance_type,
            system_disk_type="CLOUD_PREMIUM",
            system_disk_size=50,
            allocate_public_ip=True,
            internet_max_bandwidth_out=10,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        attachment = tencentcloud.as_.Attachment("attachment",
            scaling_group_id=example_scaling_group.id,
            instance_ids=[example_instance.id])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param AttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AttachmentArgs.__new__(AttachmentArgs)

            if instance_ids is None and not opts.urn:
                raise TypeError("Missing required property 'instance_ids'")
            __props__.__dict__["instance_ids"] = instance_ids
            if scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__.__dict__["scaling_group_id"] = scaling_group_id
        super(Attachment, __self__).__init__(
            'tencentcloud:As/attachment:Attachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            scaling_group_id: Optional[pulumi.Input[str]] = None) -> 'Attachment':
        """
        Get an existing Attachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: ID list of CVM instances to be attached to the scaling group.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AttachmentState.__new__(_AttachmentState)

        __props__.__dict__["instance_ids"] = instance_ids
        __props__.__dict__["scaling_group_id"] = scaling_group_id
        return Attachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        ID list of CVM instances to be attached to the scaling group.
        """
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Output[str]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

