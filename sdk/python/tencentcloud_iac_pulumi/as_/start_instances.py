# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StartInstancesArgs', 'StartInstances']

@pulumi.input_type
class StartInstancesArgs:
    def __init__(__self__, *,
                 auto_scaling_group_id: pulumi.Input[str],
                 instance_ids: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a StartInstances resource.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        pulumi.set(__self__, "auto_scaling_group_id", auto_scaling_group_id)
        pulumi.set(__self__, "instance_ids", instance_ids)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> pulumi.Input[str]:
        """
        Launch configuration ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @auto_scaling_group_id.setter
    def auto_scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_scaling_group_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of cvm instances to start.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "instance_ids", value)


@pulumi.input_type
class _StartInstancesState:
    def __init__(__self__, *,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering StartInstances resources.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        if auto_scaling_group_id is not None:
            pulumi.set(__self__, "auto_scaling_group_id", auto_scaling_group_id)
        if instance_ids is not None:
            pulumi.set(__self__, "instance_ids", instance_ids)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Launch configuration ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @auto_scaling_group_id.setter
    def auto_scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_scaling_group_id", value)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of cvm instances to start.
        """
        return pulumi.get(self, "instance_ids")

    @instance_ids.setter
    def instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_ids", value)


class StartInstances(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a as start_instances

        ## Example Usage

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
            instance_name_settings=tencentcloud.as..ScalingConfigInstanceNameSettingsArgs(
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
            instance_name="tf_example",
            availability_zone=zones.zones[0].name,
            image_id=image.images[0].image_id,
            instance_type=instance_types.instance_types[0].instance_type,
            system_disk_type="CLOUD_PREMIUM",
            system_disk_size=50,
            hostname="user",
            project_id=0,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        # Attachment Instance
        attachment = tencentcloud.as_.Attachment("attachment",
            scaling_group_id=example_scaling_group.id,
            instance_ids=[example_instance.id])
        start_instances = tencentcloud.as_.StartInstances("startInstances",
            auto_scaling_group_id=example_scaling_group.id,
            instance_ids=attachment.instance_ids)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StartInstancesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a as start_instances

        ## Example Usage

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
            instance_name_settings=tencentcloud.as..ScalingConfigInstanceNameSettingsArgs(
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
            instance_name="tf_example",
            availability_zone=zones.zones[0].name,
            image_id=image.images[0].image_id,
            instance_type=instance_types.instance_types[0].instance_type,
            system_disk_type="CLOUD_PREMIUM",
            system_disk_size=50,
            hostname="user",
            project_id=0,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        # Attachment Instance
        attachment = tencentcloud.as_.Attachment("attachment",
            scaling_group_id=example_scaling_group.id,
            instance_ids=[example_instance.id])
        start_instances = tencentcloud.as_.StartInstances("startInstances",
            auto_scaling_group_id=example_scaling_group.id,
            instance_ids=attachment.instance_ids)
        ```

        :param str resource_name: The name of the resource.
        :param StartInstancesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StartInstancesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
                 instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = StartInstancesArgs.__new__(StartInstancesArgs)

            if auto_scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'auto_scaling_group_id'")
            __props__.__dict__["auto_scaling_group_id"] = auto_scaling_group_id
            if instance_ids is None and not opts.urn:
                raise TypeError("Missing required property 'instance_ids'")
            __props__.__dict__["instance_ids"] = instance_ids
        super(StartInstances, __self__).__init__(
            'tencentcloud:As/startInstances:StartInstances',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_scaling_group_id: Optional[pulumi.Input[str]] = None,
            instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'StartInstances':
        """
        Get an existing StartInstances resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_scaling_group_id: Launch configuration ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_ids: List of cvm instances to start.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StartInstancesState.__new__(_StartInstancesState)

        __props__.__dict__["auto_scaling_group_id"] = auto_scaling_group_id
        __props__.__dict__["instance_ids"] = instance_ids
        return StartInstances(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingGroupId")
    def auto_scaling_group_id(self) -> pulumi.Output[str]:
        """
        Launch configuration ID.
        """
        return pulumi.get(self, "auto_scaling_group_id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of cvm instances to start.
        """
        return pulumi.get(self, "instance_ids")

