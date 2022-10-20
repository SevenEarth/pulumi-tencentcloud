# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReadonlyInstanceArgs', 'ReadonlyInstance']

@pulumi.input_type
class ReadonlyInstanceArgs:
    def __init__(__self__, *,
                 db_version: pulumi.Input[str],
                 master_db_instance_id: pulumi.Input[str],
                 memory: pulumi.Input[int],
                 project_id: pulumi.Input[int],
                 security_groups_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 storage: pulumi.Input[int],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 auto_renew_flag: Optional[pulumi.Input[int]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_support_ipv6: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ReadonlyInstance resource.
        :param pulumi.Input[str] db_version: PostgreSQL kernel version, which must be the same as that of the primary instance.
        :param pulumi.Input[str] master_db_instance_id: ID of the primary instance to which the read-only replica belongs.
        :param pulumi.Input[int] memory: Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group.
        :param pulumi.Input[int] storage: Instance storage capacity in GB.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        :param pulumi.Input[str] zone: Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        :param pulumi.Input[int] auto_renew_flag: Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        :param pulumi.Input[str] instance_charge_type: instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] need_support_ipv6: Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        """
        pulumi.set(__self__, "db_version", db_version)
        pulumi.set(__self__, "master_db_instance_id", master_db_instance_id)
        pulumi.set(__self__, "memory", memory)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "security_groups_ids", security_groups_ids)
        pulumi.set(__self__, "storage", storage)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "zone", zone)
        if auto_renew_flag is not None:
            pulumi.set(__self__, "auto_renew_flag", auto_renew_flag)
        if instance_charge_type is not None:
            pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if need_support_ipv6 is not None:
            pulumi.set(__self__, "need_support_ipv6", need_support_ipv6)

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> pulumi.Input[str]:
        """
        PostgreSQL kernel version, which must be the same as that of the primary instance.
        """
        return pulumi.get(self, "db_version")

    @db_version.setter
    def db_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_version", value)

    @property
    @pulumi.getter(name="masterDbInstanceId")
    def master_db_instance_id(self) -> pulumi.Input[str]:
        """
        ID of the primary instance to which the read-only replica belongs.
        """
        return pulumi.get(self, "master_db_instance_id")

    @master_db_instance_id.setter
    def master_db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "master_db_instance_id", value)

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Input[int]:
        """
        Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        """
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: pulumi.Input[int]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="securityGroupsIds")
    def security_groups_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        ID of security group.
        """
        return pulumi.get(self, "security_groups_ids")

    @security_groups_ids.setter
    def security_groups_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_groups_ids", value)

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Input[int]:
        """
        Instance storage capacity in GB.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: pulumi.Input[int]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        VPC subnet ID.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> Optional[pulumi.Input[int]]:
        """
        Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        """
        return pulumi.get(self, "auto_renew_flag")

    @auto_renew_flag.setter
    def auto_renew_flag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_renew_flag", value)

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        """
        return pulumi.get(self, "instance_charge_type")

    @instance_charge_type.setter
    def instance_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_charge_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="needSupportIpv6")
    def need_support_ipv6(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        """
        return pulumi.get(self, "need_support_ipv6")

    @need_support_ipv6.setter
    def need_support_ipv6(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "need_support_ipv6", value)


@pulumi.input_type
class _ReadonlyInstanceState:
    def __init__(__self__, *,
                 auto_renew_flag: Optional[pulumi.Input[int]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 db_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 master_db_instance_id: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_support_ipv6: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReadonlyInstance resources.
        :param pulumi.Input[int] auto_renew_flag: Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        :param pulumi.Input[str] create_time: Create time of the postgresql instance.
        :param pulumi.Input[str] db_version: PostgreSQL kernel version, which must be the same as that of the primary instance.
        :param pulumi.Input[str] instance_charge_type: instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        :param pulumi.Input[str] master_db_instance_id: ID of the primary instance to which the read-only replica belongs.
        :param pulumi.Input[int] memory: Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] need_support_ipv6: Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group.
        :param pulumi.Input[int] storage: Instance storage capacity in GB.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        :param pulumi.Input[str] zone: Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        """
        if auto_renew_flag is not None:
            pulumi.set(__self__, "auto_renew_flag", auto_renew_flag)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if db_version is not None:
            pulumi.set(__self__, "db_version", db_version)
        if instance_charge_type is not None:
            pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if master_db_instance_id is not None:
            pulumi.set(__self__, "master_db_instance_id", master_db_instance_id)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if need_support_ipv6 is not None:
            pulumi.set(__self__, "need_support_ipv6", need_support_ipv6)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if security_groups_ids is not None:
            pulumi.set(__self__, "security_groups_ids", security_groups_ids)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> Optional[pulumi.Input[int]]:
        """
        Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        """
        return pulumi.get(self, "auto_renew_flag")

    @auto_renew_flag.setter
    def auto_renew_flag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_renew_flag", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the postgresql instance.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> Optional[pulumi.Input[str]]:
        """
        PostgreSQL kernel version, which must be the same as that of the primary instance.
        """
        return pulumi.get(self, "db_version")

    @db_version.setter
    def db_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_version", value)

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        """
        return pulumi.get(self, "instance_charge_type")

    @instance_charge_type.setter
    def instance_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_charge_type", value)

    @property
    @pulumi.getter(name="masterDbInstanceId")
    def master_db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the primary instance to which the read-only replica belongs.
        """
        return pulumi.get(self, "master_db_instance_id")

    @master_db_instance_id.setter
    def master_db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_db_instance_id", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input[int]]:
        """
        Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        """
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="needSupportIpv6")
    def need_support_ipv6(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        """
        return pulumi.get(self, "need_support_ipv6")

    @need_support_ipv6.setter
    def need_support_ipv6(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "need_support_ipv6", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="securityGroupsIds")
    def security_groups_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ID of security group.
        """
        return pulumi.get(self, "security_groups_ids")

    @security_groups_ids.setter
    def security_groups_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups_ids", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input[int]]:
        """
        Instance storage capacity in GB.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC subnet ID.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class ReadonlyInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew_flag: Optional[pulumi.Input[int]] = None,
                 db_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 master_db_instance_id: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_support_ipv6: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create postgresql readonly instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.postgresql.ReadonlyInstance("foo",
            auto_renew_flag=0,
            db_version="10.4",
            instance_charge_type="POSTPAID_BY_HOUR",
            master_db_instance_id="postgres-j4pm65id",
            memory=4,
            need_support_ipv6=0,
            project_id=0,
            security_groups_ids=["sg-fefj5n6r"],
            storage=250,
            subnet_id="subnet-enm92y0m",
            vpc_id="vpc-86v957zb",
            zone="ap-guangzhou-6")
        ```

        ## Import

        postgresql readonly instance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance foo pgro-bcqx8b9a
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auto_renew_flag: Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        :param pulumi.Input[str] db_version: PostgreSQL kernel version, which must be the same as that of the primary instance.
        :param pulumi.Input[str] instance_charge_type: instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        :param pulumi.Input[str] master_db_instance_id: ID of the primary instance to which the read-only replica belongs.
        :param pulumi.Input[int] memory: Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] need_support_ipv6: Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group.
        :param pulumi.Input[int] storage: Instance storage capacity in GB.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        :param pulumi.Input[str] zone: Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReadonlyInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create postgresql readonly instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.postgresql.ReadonlyInstance("foo",
            auto_renew_flag=0,
            db_version="10.4",
            instance_charge_type="POSTPAID_BY_HOUR",
            master_db_instance_id="postgres-j4pm65id",
            memory=4,
            need_support_ipv6=0,
            project_id=0,
            security_groups_ids=["sg-fefj5n6r"],
            storage=250,
            subnet_id="subnet-enm92y0m",
            vpc_id="vpc-86v957zb",
            zone="ap-guangzhou-6")
        ```

        ## Import

        postgresql readonly instance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance foo pgro-bcqx8b9a
        ```

        :param str resource_name: The name of the resource.
        :param ReadonlyInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReadonlyInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew_flag: Optional[pulumi.Input[int]] = None,
                 db_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 master_db_instance_id: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_support_ipv6: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReadonlyInstanceArgs.__new__(ReadonlyInstanceArgs)

            __props__.__dict__["auto_renew_flag"] = auto_renew_flag
            if db_version is None and not opts.urn:
                raise TypeError("Missing required property 'db_version'")
            __props__.__dict__["db_version"] = db_version
            __props__.__dict__["instance_charge_type"] = instance_charge_type
            if master_db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'master_db_instance_id'")
            __props__.__dict__["master_db_instance_id"] = master_db_instance_id
            if memory is None and not opts.urn:
                raise TypeError("Missing required property 'memory'")
            __props__.__dict__["memory"] = memory
            __props__.__dict__["name"] = name
            __props__.__dict__["need_support_ipv6"] = need_support_ipv6
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if security_groups_ids is None and not opts.urn:
                raise TypeError("Missing required property 'security_groups_ids'")
            __props__.__dict__["security_groups_ids"] = security_groups_ids
            if storage is None and not opts.urn:
                raise TypeError("Missing required property 'storage'")
            __props__.__dict__["storage"] = storage
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["create_time"] = None
        super(ReadonlyInstance, __self__).__init__(
            'tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew_flag: Optional[pulumi.Input[int]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            db_version: Optional[pulumi.Input[str]] = None,
            instance_charge_type: Optional[pulumi.Input[str]] = None,
            master_db_instance_id: Optional[pulumi.Input[str]] = None,
            memory: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            need_support_ipv6: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            storage: Optional[pulumi.Input[int]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'ReadonlyInstance':
        """
        Get an existing ReadonlyInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auto_renew_flag: Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        :param pulumi.Input[str] create_time: Create time of the postgresql instance.
        :param pulumi.Input[str] db_version: PostgreSQL kernel version, which must be the same as that of the primary instance.
        :param pulumi.Input[str] instance_charge_type: instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        :param pulumi.Input[str] master_db_instance_id: ID of the primary instance to which the read-only replica belongs.
        :param pulumi.Input[int] memory: Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] need_support_ipv6: Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group.
        :param pulumi.Input[int] storage: Instance storage capacity in GB.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        :param pulumi.Input[str] zone: Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReadonlyInstanceState.__new__(_ReadonlyInstanceState)

        __props__.__dict__["auto_renew_flag"] = auto_renew_flag
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["db_version"] = db_version
        __props__.__dict__["instance_charge_type"] = instance_charge_type
        __props__.__dict__["master_db_instance_id"] = master_db_instance_id
        __props__.__dict__["memory"] = memory
        __props__.__dict__["name"] = name
        __props__.__dict__["need_support_ipv6"] = need_support_ipv6
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["security_groups_ids"] = security_groups_ids
        __props__.__dict__["storage"] = storage
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["zone"] = zone
        return ReadonlyInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> pulumi.Output[Optional[int]]:
        """
        Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
        """
        return pulumi.get(self, "auto_renew_flag")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the postgresql instance.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dbVersion")
    def db_version(self) -> pulumi.Output[str]:
        """
        PostgreSQL kernel version, which must be the same as that of the primary instance.
        """
        return pulumi.get(self, "db_version")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
        """
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter(name="masterDbInstanceId")
    def master_db_instance_id(self) -> pulumi.Output[str]:
        """
        ID of the primary instance to which the read-only replica belongs.
        """
        return pulumi.get(self, "master_db_instance_id")

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Output[int]:
        """
        Memory size(in GB). Allowed value must be larger than `memory` that data source `_postgresql.get_specinfos` provides.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="needSupportIpv6")
    def need_support_ipv6(self) -> pulumi.Output[Optional[int]]:
        """
        Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
        """
        return pulumi.get(self, "need_support_ipv6")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="securityGroupsIds")
    def security_groups_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        ID of security group.
        """
        return pulumi.get(self, "security_groups_ids")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output[int]:
        """
        Instance storage capacity in GB.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        VPC subnet ID.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        VPC ID.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
        """
        return pulumi.get(self, "zone")

