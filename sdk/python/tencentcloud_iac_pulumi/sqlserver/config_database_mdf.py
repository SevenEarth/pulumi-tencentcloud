# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConfigDatabaseMdfArgs', 'ConfigDatabaseMdf']

@pulumi.input_type
class ConfigDatabaseMdfArgs:
    def __init__(__self__, *,
                 db_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ConfigDatabaseMdf resource.
        :param pulumi.Input[str] db_name: Array of database names.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        Array of database names.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

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


@pulumi.input_type
class _ConfigDatabaseMdfState:
    def __init__(__self__, *,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConfigDatabaseMdf resources.
        :param pulumi.Input[str] db_name: Array of database names.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        if db_name is not None:
            pulumi.set(__self__, "db_name", db_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[pulumi.Input[str]]:
        """
        Array of database names.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_name", value)

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


class ConfigDatabaseMdf(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a sqlserver config_database_mdf

        ## Example Usage

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
            charge_type="POSTPAID_BY_HOUR",
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
        example_db = tencentcloud.sqlserver.Db("exampleDb",
            instance_id=example_basic_instance.id,
            charset="Chinese_PRC_BIN",
            remark="test-remark")
        example_config_database_mdf = tencentcloud.sqlserver.ConfigDatabaseMdf("exampleConfigDatabaseMdf",
            db_name=example_db.name,
            instance_id=example_basic_instance.id)
        ```

        ## Import

        sqlserver config_database_mdf can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf example mssql-i9ma6oy7#tf_example_db
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_name: Array of database names.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigDatabaseMdfArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a sqlserver config_database_mdf

        ## Example Usage

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
            charge_type="POSTPAID_BY_HOUR",
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
        example_db = tencentcloud.sqlserver.Db("exampleDb",
            instance_id=example_basic_instance.id,
            charset="Chinese_PRC_BIN",
            remark="test-remark")
        example_config_database_mdf = tencentcloud.sqlserver.ConfigDatabaseMdf("exampleConfigDatabaseMdf",
            db_name=example_db.name,
            instance_id=example_basic_instance.id)
        ```

        ## Import

        sqlserver config_database_mdf can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf example mssql-i9ma6oy7#tf_example_db
        ```

        :param str resource_name: The name of the resource.
        :param ConfigDatabaseMdfArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigDatabaseMdfArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConfigDatabaseMdfArgs.__new__(ConfigDatabaseMdfArgs)

            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__.__dict__["db_name"] = db_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(ConfigDatabaseMdf, __self__).__init__(
            'tencentcloud:Sqlserver/configDatabaseMdf:ConfigDatabaseMdf',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'ConfigDatabaseMdf':
        """
        Get an existing ConfigDatabaseMdf resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_name: Array of database names.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigDatabaseMdfState.__new__(_ConfigDatabaseMdfState)

        __props__.__dict__["db_name"] = db_name
        __props__.__dict__["instance_id"] = instance_id
        return ConfigDatabaseMdf(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        Array of database names.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

