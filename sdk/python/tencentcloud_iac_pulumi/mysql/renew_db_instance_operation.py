# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RenewDbInstanceOperationArgs', 'RenewDbInstanceOperation']

@pulumi.input_type
class RenewDbInstanceOperationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 time_span: pulumi.Input[int],
                 modify_pay_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RenewDbInstanceOperation resource.
        :param pulumi.Input[str] instance_id: The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        :param pulumi.Input[int] time_span: Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        :param pulumi.Input[str] modify_pay_type: If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "time_span", time_span)
        if modify_pay_type is not None:
            pulumi.set(__self__, "modify_pay_type", modify_pay_type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> pulumi.Input[int]:
        """
        Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        """
        return pulumi.get(self, "time_span")

    @time_span.setter
    def time_span(self, value: pulumi.Input[int]):
        pulumi.set(self, "time_span", value)

    @property
    @pulumi.getter(name="modifyPayType")
    def modify_pay_type(self) -> Optional[pulumi.Input[str]]:
        """
        If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        """
        return pulumi.get(self, "modify_pay_type")

    @modify_pay_type.setter
    def modify_pay_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_pay_type", value)


@pulumi.input_type
class _RenewDbInstanceOperationState:
    def __init__(__self__, *,
                 deadline_time: Optional[pulumi.Input[str]] = None,
                 deal_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 modify_pay_type: Optional[pulumi.Input[str]] = None,
                 time_span: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RenewDbInstanceOperation resources.
        :param pulumi.Input[str] deadline_time: Instance expiration time.
        :param pulumi.Input[str] deal_id: Deal id.
        :param pulumi.Input[str] instance_id: The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        :param pulumi.Input[str] modify_pay_type: If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        :param pulumi.Input[int] time_span: Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        """
        if deadline_time is not None:
            pulumi.set(__self__, "deadline_time", deadline_time)
        if deal_id is not None:
            pulumi.set(__self__, "deal_id", deal_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if modify_pay_type is not None:
            pulumi.set(__self__, "modify_pay_type", modify_pay_type)
        if time_span is not None:
            pulumi.set(__self__, "time_span", time_span)

    @property
    @pulumi.getter(name="deadlineTime")
    def deadline_time(self) -> Optional[pulumi.Input[str]]:
        """
        Instance expiration time.
        """
        return pulumi.get(self, "deadline_time")

    @deadline_time.setter
    def deadline_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deadline_time", value)

    @property
    @pulumi.getter(name="dealId")
    def deal_id(self) -> Optional[pulumi.Input[str]]:
        """
        Deal id.
        """
        return pulumi.get(self, "deal_id")

    @deal_id.setter
    def deal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deal_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="modifyPayType")
    def modify_pay_type(self) -> Optional[pulumi.Input[str]]:
        """
        If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        """
        return pulumi.get(self, "modify_pay_type")

    @modify_pay_type.setter
    def modify_pay_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_pay_type", value)

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> Optional[pulumi.Input[int]]:
        """
        Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        """
        return pulumi.get(self, "time_span")

    @time_span.setter
    def time_span(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_span", value)


class RenewDbInstanceOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 modify_pay_type: Optional[pulumi.Input[str]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql renew_db_instance_operation

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
            charge_type="PREPAID",
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
        example_rollback_range_time = tencentcloud.Mysql.get_rollback_range_time_output(instance_ids=[example_instance.id])
        example_renew_db_instance_operation = tencentcloud.mysql.RenewDbInstanceOperation("exampleRenewDbInstanceOperation",
            instance_id=example_instance.id,
            time_span=1,
            modify_pay_type="PREPAID")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        :param pulumi.Input[str] modify_pay_type: If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        :param pulumi.Input[int] time_span: Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RenewDbInstanceOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql renew_db_instance_operation

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
            charge_type="PREPAID",
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
        example_rollback_range_time = tencentcloud.Mysql.get_rollback_range_time_output(instance_ids=[example_instance.id])
        example_renew_db_instance_operation = tencentcloud.mysql.RenewDbInstanceOperation("exampleRenewDbInstanceOperation",
            instance_id=example_instance.id,
            time_span=1,
            modify_pay_type="PREPAID")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param RenewDbInstanceOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RenewDbInstanceOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 modify_pay_type: Optional[pulumi.Input[str]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RenewDbInstanceOperationArgs.__new__(RenewDbInstanceOperationArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["modify_pay_type"] = modify_pay_type
            if time_span is None and not opts.urn:
                raise TypeError("Missing required property 'time_span'")
            __props__.__dict__["time_span"] = time_span
            __props__.__dict__["deadline_time"] = None
            __props__.__dict__["deal_id"] = None
        super(RenewDbInstanceOperation, __self__).__init__(
            'tencentcloud:Mysql/renewDbInstanceOperation:RenewDbInstanceOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deadline_time: Optional[pulumi.Input[str]] = None,
            deal_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            modify_pay_type: Optional[pulumi.Input[str]] = None,
            time_span: Optional[pulumi.Input[int]] = None) -> 'RenewDbInstanceOperation':
        """
        Get an existing RenewDbInstanceOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deadline_time: Instance expiration time.
        :param pulumi.Input[str] deal_id: Deal id.
        :param pulumi.Input[str] instance_id: The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        :param pulumi.Input[str] modify_pay_type: If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        :param pulumi.Input[int] time_span: Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RenewDbInstanceOperationState.__new__(_RenewDbInstanceOperationState)

        __props__.__dict__["deadline_time"] = deadline_time
        __props__.__dict__["deal_id"] = deal_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["modify_pay_type"] = modify_pay_type
        __props__.__dict__["time_span"] = time_span
        return RenewDbInstanceOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deadlineTime")
    def deadline_time(self) -> pulumi.Output[str]:
        """
        Instance expiration time.
        """
        return pulumi.get(self, "deadline_time")

    @property
    @pulumi.getter(name="dealId")
    def deal_id(self) -> pulumi.Output[str]:
        """
        Deal id.
        """
        return pulumi.get(self, "deal_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance ID to be renewed, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, you can use [Query Instance List](https://cloud.tencent.com/document/api/236/ 15872).
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="modifyPayType")
    def modify_pay_type(self) -> pulumi.Output[Optional[str]]:
        """
        If you need to renew the Pay-As-You-Go instance to a Subscription instance, the value of this input parameter needs to be specified as `PREPAID`.
        """
        return pulumi.get(self, "modify_pay_type")

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> pulumi.Output[int]:
        """
        Renewal duration, unit: month, optional values include [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
        """
        return pulumi.get(self, "time_span")

