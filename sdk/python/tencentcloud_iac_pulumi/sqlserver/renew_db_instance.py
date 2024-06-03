# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RenewDbInstanceArgs', 'RenewDbInstance']

@pulumi.input_type
class RenewDbInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 period: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RenewDbInstance resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: How many months to renew, the value range is 1-48, the default is 1.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if period is not None:
            pulumi.set(__self__, "period", period)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        How many months to renew, the value range is 1-48, the default is 1.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)


@pulumi.input_type
class _RenewDbInstanceState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RenewDbInstance resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: How many months to renew, the value range is 1-48, the default is 1.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if period is not None:
            pulumi.set(__self__, "period", period)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        How many months to renew, the value range is 1-48, the default is 1.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)


class RenewDbInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a sqlserver renew_db_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="sqlserver")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            availability_zone=zones.zones[4].name,
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            is_multicast=False)
        security_group = tencentcloud.security.Group("securityGroup", description="desc.")
        example_basic_instance = tencentcloud.sqlserver.BasicInstance("exampleBasicInstance",
            availability_zone=zones.zones[4].name,
            charge_type="PREPAID",
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            project_id=0,
            memory=4,
            storage=100,
            cpu=2,
            machine_type="CLOUD_PREMIUM",
            maintenance_week_sets=[
                1,
                2,
                3,
            ],
            maintenance_start_time="09:00",
            maintenance_time_span=3,
            security_groups=[security_group.id],
            tags={
                "test": "test",
            })
        example_renew_db_instance = tencentcloud.sqlserver.RenewDbInstance("exampleRenewDbInstance",
            instance_id=example_basic_instance.id,
            period=1)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        sqlserver renew_db_instance can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance example mssql-i9ma6oy7#1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: How many months to renew, the value range is 1-48, the default is 1.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RenewDbInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a sqlserver renew_db_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="sqlserver")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            availability_zone=zones.zones[4].name,
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            is_multicast=False)
        security_group = tencentcloud.security.Group("securityGroup", description="desc.")
        example_basic_instance = tencentcloud.sqlserver.BasicInstance("exampleBasicInstance",
            availability_zone=zones.zones[4].name,
            charge_type="PREPAID",
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            project_id=0,
            memory=4,
            storage=100,
            cpu=2,
            machine_type="CLOUD_PREMIUM",
            maintenance_week_sets=[
                1,
                2,
                3,
            ],
            maintenance_start_time="09:00",
            maintenance_time_span=3,
            security_groups=[security_group.id],
            tags={
                "test": "test",
            })
        example_renew_db_instance = tencentcloud.sqlserver.RenewDbInstance("exampleRenewDbInstance",
            instance_id=example_basic_instance.id,
            period=1)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        sqlserver renew_db_instance can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance example mssql-i9ma6oy7#1
        ```

        :param str resource_name: The name of the resource.
        :param RenewDbInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RenewDbInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RenewDbInstanceArgs.__new__(RenewDbInstanceArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["period"] = period
        super(RenewDbInstance, __self__).__init__(
            'tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None) -> 'RenewDbInstance':
        """
        Get an existing RenewDbInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] period: How many months to renew, the value range is 1-48, the default is 1.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RenewDbInstanceState.__new__(_RenewDbInstanceState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["period"] = period
        return RenewDbInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        How many months to renew, the value range is 1-48, the default is 1.
        """
        return pulumi.get(self, "period")

