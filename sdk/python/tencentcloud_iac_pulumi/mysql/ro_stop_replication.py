# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RoStopReplicationArgs', 'RoStopReplication']

@pulumi.input_type
class RoStopReplicationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a RoStopReplication resource.
        :param pulumi.Input[str] instance_id: Read-Only instance ID.
        """
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Read-Only instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _RoStopReplicationState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RoStopReplication resources.
        :param pulumi.Input[str] instance_id: Read-Only instance ID.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Read-Only instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class RoStopReplication(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql ro_stop_replication

        ## Example Usage

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
            slave_deploy_mode=1,
            availability_zone=zones.zones[0].name,
            first_slave_zone=zones.zones[1].name,
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
        example_proxy = tencentcloud.mysql.Proxy("exampleProxy",
            instance_id=example_instance.id,
            uniq_vpc_id=vpc.id,
            uniq_subnet_id=subnet.id,
            proxy_node_customs=[tencentcloud.mysql.ProxyProxyNodeCustomArgs(
                node_count=1,
                cpu=2,
                mem=4000,
                region="ap-guangzhou",
                zone="ap-guangzhou-3",
            )],
            security_groups=[security_group.id],
            desc="desc.",
            connection_pool_limit=2,
            vip="10.0.0.120",
            vport=3306)
        example_ro_stop_replication = tencentcloud.mysql.RoStopReplication("exampleRoStopReplication", instance_id=example_proxy.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Read-Only instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoStopReplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql ro_stop_replication

        ## Example Usage

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
            slave_deploy_mode=1,
            availability_zone=zones.zones[0].name,
            first_slave_zone=zones.zones[1].name,
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
        example_proxy = tencentcloud.mysql.Proxy("exampleProxy",
            instance_id=example_instance.id,
            uniq_vpc_id=vpc.id,
            uniq_subnet_id=subnet.id,
            proxy_node_customs=[tencentcloud.mysql.ProxyProxyNodeCustomArgs(
                node_count=1,
                cpu=2,
                mem=4000,
                region="ap-guangzhou",
                zone="ap-guangzhou-3",
            )],
            security_groups=[security_group.id],
            desc="desc.",
            connection_pool_limit=2,
            vip="10.0.0.120",
            vport=3306)
        example_ro_stop_replication = tencentcloud.mysql.RoStopReplication("exampleRoStopReplication", instance_id=example_proxy.id)
        ```

        :param str resource_name: The name of the resource.
        :param RoStopReplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoStopReplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RoStopReplicationArgs.__new__(RoStopReplicationArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(RoStopReplication, __self__).__init__(
            'tencentcloud:Mysql/roStopReplication:RoStopReplication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'RoStopReplication':
        """
        Get an existing RoStopReplication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Read-Only instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoStopReplicationState.__new__(_RoStopReplicationState)

        __props__.__dict__["instance_id"] = instance_id
        return RoStopReplication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Read-Only instance ID.
        """
        return pulumi.get(self, "instance_id")

