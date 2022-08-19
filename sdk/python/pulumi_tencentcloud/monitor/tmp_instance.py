# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TmpInstanceArgs', 'TmpInstance']

@pulumi.input_type
class TmpInstanceArgs:
    def __init__(__self__, *,
                 data_retention_time: pulumi.Input[int],
                 instance_name: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a TmpInstance resource.
        :param pulumi.Input[int] data_retention_time: Data retention time.
        :param pulumi.Input[str] instance_name: Instance name.
        :param pulumi.Input[str] subnet_id: Subnet Id.
        :param pulumi.Input[str] vpc_id: Vpc Id.
        :param pulumi.Input[str] zone: Available zone.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "data_retention_time", data_retention_time)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "zone", zone)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataRetentionTime")
    def data_retention_time(self) -> pulumi.Input[int]:
        """
        Data retention time.
        """
        return pulumi.get(self, "data_retention_time")

    @data_retention_time.setter
    def data_retention_time(self, value: pulumi.Input[int]):
        pulumi.set(self, "data_retention_time", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        Instance name.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Subnet Id.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        Vpc Id.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        Available zone.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TmpInstanceState:
    def __init__(__self__, *,
                 data_retention_time: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TmpInstance resources.
        :param pulumi.Input[int] data_retention_time: Data retention time.
        :param pulumi.Input[str] instance_name: Instance name.
        :param pulumi.Input[str] subnet_id: Subnet Id.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[str] vpc_id: Vpc Id.
        :param pulumi.Input[str] zone: Available zone.
        """
        if data_retention_time is not None:
            pulumi.set(__self__, "data_retention_time", data_retention_time)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="dataRetentionTime")
    def data_retention_time(self) -> Optional[pulumi.Input[int]]:
        """
        Data retention time.
        """
        return pulumi.get(self, "data_retention_time")

    @data_retention_time.setter
    def data_retention_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "data_retention_time", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet Id.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Vpc Id.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Available zone.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class TmpInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_retention_time: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a monitor tmpInstance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        tmp_instance = tencentcloud.monitor.TmpInstance("tmpInstance",
            data_retention_time=30,
            instance_name="demo",
            subnet_id="subnet-rdkj0agk",
            tags={
                "createdBy": "terraform",
            },
            vpc_id="vpc-2hfyray3",
            zone="ap-guangzhou-3")
        ```

        ## Import

        monitor tmpInstance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/tmpInstance:TmpInstance tmpInstance tmpInstance_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] data_retention_time: Data retention time.
        :param pulumi.Input[str] instance_name: Instance name.
        :param pulumi.Input[str] subnet_id: Subnet Id.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[str] vpc_id: Vpc Id.
        :param pulumi.Input[str] zone: Available zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TmpInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a monitor tmpInstance

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        tmp_instance = tencentcloud.monitor.TmpInstance("tmpInstance",
            data_retention_time=30,
            instance_name="demo",
            subnet_id="subnet-rdkj0agk",
            tags={
                "createdBy": "terraform",
            },
            vpc_id="vpc-2hfyray3",
            zone="ap-guangzhou-3")
        ```

        ## Import

        monitor tmpInstance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/tmpInstance:TmpInstance tmpInstance tmpInstance_id
        ```

        :param str resource_name: The name of the resource.
        :param TmpInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TmpInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_retention_time: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = TmpInstanceArgs.__new__(TmpInstanceArgs)

            if data_retention_time is None and not opts.urn:
                raise TypeError("Missing required property 'data_retention_time'")
            __props__.__dict__["data_retention_time"] = data_retention_time
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(TmpInstance, __self__).__init__(
            'tencentcloud:Monitor/tmpInstance:TmpInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_retention_time: Optional[pulumi.Input[int]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'TmpInstance':
        """
        Get an existing TmpInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] data_retention_time: Data retention time.
        :param pulumi.Input[str] instance_name: Instance name.
        :param pulumi.Input[str] subnet_id: Subnet Id.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[str] vpc_id: Vpc Id.
        :param pulumi.Input[str] zone: Available zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TmpInstanceState.__new__(_TmpInstanceState)

        __props__.__dict__["data_retention_time"] = data_retention_time
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["zone"] = zone
        return TmpInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataRetentionTime")
    def data_retention_time(self) -> pulumi.Output[int]:
        """
        Data retention time.
        """
        return pulumi.get(self, "data_retention_time")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        Instance name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Subnet Id.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        Vpc Id.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Available zone.
        """
        return pulumi.get(self, "zone")

