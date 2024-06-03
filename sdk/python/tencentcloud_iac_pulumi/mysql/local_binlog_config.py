# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LocalBinlogConfigArgs', 'LocalBinlogConfig']

@pulumi.input_type
class LocalBinlogConfigArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 max_usage: pulumi.Input[int],
                 save_hours: pulumi.Input[int]):
        """
        The set of arguments for constructing a LocalBinlogConfig resource.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        :param pulumi.Input[int] max_usage: Space utilization of local binlog. Value range: [30,50].
        :param pulumi.Input[int] save_hours: Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "max_usage", max_usage)
        pulumi.set(__self__, "save_hours", save_hours)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxUsage")
    def max_usage(self) -> pulumi.Input[int]:
        """
        Space utilization of local binlog. Value range: [30,50].
        """
        return pulumi.get(self, "max_usage")

    @max_usage.setter
    def max_usage(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_usage", value)

    @property
    @pulumi.getter(name="saveHours")
    def save_hours(self) -> pulumi.Input[int]:
        """
        Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        return pulumi.get(self, "save_hours")

    @save_hours.setter
    def save_hours(self, value: pulumi.Input[int]):
        pulumi.set(self, "save_hours", value)


@pulumi.input_type
class _LocalBinlogConfigState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_usage: Optional[pulumi.Input[int]] = None,
                 save_hours: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering LocalBinlogConfig resources.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        :param pulumi.Input[int] max_usage: Space utilization of local binlog. Value range: [30,50].
        :param pulumi.Input[int] save_hours: Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_usage is not None:
            pulumi.set(__self__, "max_usage", max_usage)
        if save_hours is not None:
            pulumi.set(__self__, "save_hours", save_hours)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxUsage")
    def max_usage(self) -> Optional[pulumi.Input[int]]:
        """
        Space utilization of local binlog. Value range: [30,50].
        """
        return pulumi.get(self, "max_usage")

    @max_usage.setter
    def max_usage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_usage", value)

    @property
    @pulumi.getter(name="saveHours")
    def save_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        return pulumi.get(self, "save_hours")

    @save_hours.setter
    def save_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "save_hours", value)


class LocalBinlogConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_usage: Optional[pulumi.Input[int]] = None,
                 save_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql local_binlog_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="cdb")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            availability_zone=zones.zones[0].name,
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            is_multicast=False)
        security_group = tencentcloud.security.Group("securityGroup", description="mysql test")
        example_instance = tencentcloud.mysql.Instance("exampleInstance",
            internet_service=1,
            engine_version="5.7",
            charge_type="POSTPAID",
            root_password="PassWord123",
            slave_deploy_mode=0,
            availability_zone=zones.zones[0].name,
            slave_sync_mode=1,
            instance_name="tf-example-mysql",
            mem_size=4000,
            volume_size=200,
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            intranet_port=3306,
            security_groups=[security_group.id],
            tags={
                "name": "test",
            },
            parameters={
                "character_set_server": "utf8",
                "max_connections": "1000",
            })
        example_local_binlog_config = tencentcloud.mysql.LocalBinlogConfig("exampleLocalBinlogConfig",
            instance_id=example_instance.id,
            save_hours=140,
            max_usage=50)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        mysql local_binlog_config can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig local_binlog_config instance_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        :param pulumi.Input[int] max_usage: Space utilization of local binlog. Value range: [30,50].
        :param pulumi.Input[int] save_hours: Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LocalBinlogConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql local_binlog_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="cdb")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            availability_zone=zones.zones[0].name,
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            is_multicast=False)
        security_group = tencentcloud.security.Group("securityGroup", description="mysql test")
        example_instance = tencentcloud.mysql.Instance("exampleInstance",
            internet_service=1,
            engine_version="5.7",
            charge_type="POSTPAID",
            root_password="PassWord123",
            slave_deploy_mode=0,
            availability_zone=zones.zones[0].name,
            slave_sync_mode=1,
            instance_name="tf-example-mysql",
            mem_size=4000,
            volume_size=200,
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            intranet_port=3306,
            security_groups=[security_group.id],
            tags={
                "name": "test",
            },
            parameters={
                "character_set_server": "utf8",
                "max_connections": "1000",
            })
        example_local_binlog_config = tencentcloud.mysql.LocalBinlogConfig("exampleLocalBinlogConfig",
            instance_id=example_instance.id,
            save_hours=140,
            max_usage=50)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        mysql local_binlog_config can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig local_binlog_config instance_id
        ```

        :param str resource_name: The name of the resource.
        :param LocalBinlogConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocalBinlogConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_usage: Optional[pulumi.Input[int]] = None,
                 save_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocalBinlogConfigArgs.__new__(LocalBinlogConfigArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if max_usage is None and not opts.urn:
                raise TypeError("Missing required property 'max_usage'")
            __props__.__dict__["max_usage"] = max_usage
            if save_hours is None and not opts.urn:
                raise TypeError("Missing required property 'save_hours'")
            __props__.__dict__["save_hours"] = save_hours
        super(LocalBinlogConfig, __self__).__init__(
            'tencentcloud:Mysql/localBinlogConfig:LocalBinlogConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            max_usage: Optional[pulumi.Input[int]] = None,
            save_hours: Optional[pulumi.Input[int]] = None) -> 'LocalBinlogConfig':
        """
        Get an existing LocalBinlogConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        :param pulumi.Input[int] max_usage: Space utilization of local binlog. Value range: [30,50].
        :param pulumi.Input[int] save_hours: Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LocalBinlogConfigState.__new__(_LocalBinlogConfigState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_usage"] = max_usage
        __props__.__dict__["save_hours"] = save_hours
        return LocalBinlogConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxUsage")
    def max_usage(self) -> pulumi.Output[int]:
        """
        Space utilization of local binlog. Value range: [30,50].
        """
        return pulumi.get(self, "max_usage")

    @property
    @pulumi.getter(name="saveHours")
    def save_hours(self) -> pulumi.Output[int]:
        """
        Retention period of local binlog. Valid range: 72-168 hours. When there is disaster recovery instance, the valid range will be 120-168 hours.
        """
        return pulumi.get(self, "save_hours")

