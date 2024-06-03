# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MigrateServiceArgs', 'MigrateService']

@pulumi.input_type
class MigrateServiceArgs:
    def __init__(__self__, *,
                 dst_database_type: pulumi.Input[str],
                 dst_region: pulumi.Input[str],
                 instance_class: pulumi.Input[str],
                 src_database_type: pulumi.Input[str],
                 src_region: pulumi.Input[str],
                 job_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]]] = None):
        """
        The set of arguments for constructing a MigrateService resource.
        :param pulumi.Input[str] dst_database_type: destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] instance_class: instance class, optional value is small/medium/large/xlarge/2xlarge.
        :param pulumi.Input[str] src_database_type: source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]] tags: tags.
        """
        pulumi.set(__self__, "dst_database_type", dst_database_type)
        pulumi.set(__self__, "dst_region", dst_region)
        pulumi.set(__self__, "instance_class", instance_class)
        pulumi.set(__self__, "src_database_type", src_database_type)
        pulumi.set(__self__, "src_region", src_region)
        if job_name is not None:
            pulumi.set(__self__, "job_name", job_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dstDatabaseType")
    def dst_database_type(self) -> pulumi.Input[str]:
        """
        destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        """
        return pulumi.get(self, "dst_database_type")

    @dst_database_type.setter
    def dst_database_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "dst_database_type", value)

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> pulumi.Input[str]:
        """
        destination region.
        """
        return pulumi.get(self, "dst_region")

    @dst_region.setter
    def dst_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "dst_region", value)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Input[str]:
        """
        instance class, optional value is small/medium/large/xlarge/2xlarge.
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="srcDatabaseType")
    def src_database_type(self) -> pulumi.Input[str]:
        """
        source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        """
        return pulumi.get(self, "src_database_type")

    @src_database_type.setter
    def src_database_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "src_database_type", value)

    @property
    @pulumi.getter(name="srcRegion")
    def src_region(self) -> pulumi.Input[str]:
        """
        source region.
        """
        return pulumi.get(self, "src_region")

    @src_region.setter
    def src_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "src_region", value)

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[pulumi.Input[str]]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @job_name.setter
    def job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]]]:
        """
        tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _MigrateServiceState:
    def __init__(__self__, *,
                 dst_database_type: Optional[pulumi.Input[str]] = None,
                 dst_region: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 src_database_type: Optional[pulumi.Input[str]] = None,
                 src_region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering MigrateService resources.
        :param pulumi.Input[str] dst_database_type: destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] instance_class: instance class, optional value is small/medium/large/xlarge/2xlarge.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] src_database_type: source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]] tags: tags.
        """
        if dst_database_type is not None:
            pulumi.set(__self__, "dst_database_type", dst_database_type)
        if dst_region is not None:
            pulumi.set(__self__, "dst_region", dst_region)
        if instance_class is not None:
            pulumi.set(__self__, "instance_class", instance_class)
        if job_name is not None:
            pulumi.set(__self__, "job_name", job_name)
        if src_database_type is not None:
            pulumi.set(__self__, "src_database_type", src_database_type)
        if src_region is not None:
            pulumi.set(__self__, "src_region", src_region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dstDatabaseType")
    def dst_database_type(self) -> Optional[pulumi.Input[str]]:
        """
        destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        """
        return pulumi.get(self, "dst_database_type")

    @dst_database_type.setter
    def dst_database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst_database_type", value)

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> Optional[pulumi.Input[str]]:
        """
        destination region.
        """
        return pulumi.get(self, "dst_region")

    @dst_region.setter
    def dst_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst_region", value)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> Optional[pulumi.Input[str]]:
        """
        instance class, optional value is small/medium/large/xlarge/2xlarge.
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[pulumi.Input[str]]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @job_name.setter
    def job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_name", value)

    @property
    @pulumi.getter(name="srcDatabaseType")
    def src_database_type(self) -> Optional[pulumi.Input[str]]:
        """
        source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        """
        return pulumi.get(self, "src_database_type")

    @src_database_type.setter
    def src_database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_database_type", value)

    @property
    @pulumi.getter(name="srcRegion")
    def src_region(self) -> Optional[pulumi.Input[str]]:
        """
        source region.
        """
        return pulumi.get(self, "src_region")

    @src_region.setter
    def src_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]]]:
        """
        tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MigrateServiceTagArgs']]]]):
        pulumi.set(self, "tags", value)


class MigrateService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dst_database_type: Optional[pulumi.Input[str]] = None,
                 dst_region: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 src_database_type: Optional[pulumi.Input[str]] = None,
                 src_region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MigrateServiceTagArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a dts migrate_service

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        migrate_service = tencentcloud.dts.MigrateService("migrateService",
            dst_database_type="cynosdbmysql",
            dst_region="ap-guangzhou",
            instance_class="small",
            job_name="tf_test_migration_job",
            src_database_type="mysql",
            src_region="ap-guangzhou",
            tags=[tencentcloud.dts.MigrateServiceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        dts migrate_service can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Dts/migrateService:MigrateService migrate_service migrateService_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dst_database_type: destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] instance_class: instance class, optional value is small/medium/large/xlarge/2xlarge.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] src_database_type: source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MigrateServiceTagArgs']]]] tags: tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MigrateServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dts migrate_service

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        migrate_service = tencentcloud.dts.MigrateService("migrateService",
            dst_database_type="cynosdbmysql",
            dst_region="ap-guangzhou",
            instance_class="small",
            job_name="tf_test_migration_job",
            src_database_type="mysql",
            src_region="ap-guangzhou",
            tags=[tencentcloud.dts.MigrateServiceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        dts migrate_service can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Dts/migrateService:MigrateService migrate_service migrateService_id
        ```

        :param str resource_name: The name of the resource.
        :param MigrateServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MigrateServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dst_database_type: Optional[pulumi.Input[str]] = None,
                 dst_region: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 src_database_type: Optional[pulumi.Input[str]] = None,
                 src_region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MigrateServiceTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MigrateServiceArgs.__new__(MigrateServiceArgs)

            if dst_database_type is None and not opts.urn:
                raise TypeError("Missing required property 'dst_database_type'")
            __props__.__dict__["dst_database_type"] = dst_database_type
            if dst_region is None and not opts.urn:
                raise TypeError("Missing required property 'dst_region'")
            __props__.__dict__["dst_region"] = dst_region
            if instance_class is None and not opts.urn:
                raise TypeError("Missing required property 'instance_class'")
            __props__.__dict__["instance_class"] = instance_class
            __props__.__dict__["job_name"] = job_name
            if src_database_type is None and not opts.urn:
                raise TypeError("Missing required property 'src_database_type'")
            __props__.__dict__["src_database_type"] = src_database_type
            if src_region is None and not opts.urn:
                raise TypeError("Missing required property 'src_region'")
            __props__.__dict__["src_region"] = src_region
            __props__.__dict__["tags"] = tags
        super(MigrateService, __self__).__init__(
            'tencentcloud:Dts/migrateService:MigrateService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dst_database_type: Optional[pulumi.Input[str]] = None,
            dst_region: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            job_name: Optional[pulumi.Input[str]] = None,
            src_database_type: Optional[pulumi.Input[str]] = None,
            src_region: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MigrateServiceTagArgs']]]]] = None) -> 'MigrateService':
        """
        Get an existing MigrateService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dst_database_type: destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] instance_class: instance class, optional value is small/medium/large/xlarge/2xlarge.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] src_database_type: source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MigrateServiceTagArgs']]]] tags: tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MigrateServiceState.__new__(_MigrateServiceState)

        __props__.__dict__["dst_database_type"] = dst_database_type
        __props__.__dict__["dst_region"] = dst_region
        __props__.__dict__["instance_class"] = instance_class
        __props__.__dict__["job_name"] = job_name
        __props__.__dict__["src_database_type"] = src_database_type
        __props__.__dict__["src_region"] = src_region
        __props__.__dict__["tags"] = tags
        return MigrateService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dstDatabaseType")
    def dst_database_type(self) -> pulumi.Output[str]:
        """
        destination database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        """
        return pulumi.get(self, "dst_database_type")

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> pulumi.Output[str]:
        """
        destination region.
        """
        return pulumi.get(self, "dst_region")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Output[str]:
        """
        instance class, optional value is small/medium/large/xlarge/2xlarge.
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> pulumi.Output[Optional[str]]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @property
    @pulumi.getter(name="srcDatabaseType")
    def src_database_type(self) -> pulumi.Output[str]:
        """
        source database type, optional value is mysql/redis/percona/mongodb/postgresql/sqlserver/mariadb.
        """
        return pulumi.get(self, "src_database_type")

    @property
    @pulumi.getter(name="srcRegion")
    def src_region(self) -> pulumi.Output[str]:
        """
        source region.
        """
        return pulumi.get(self, "src_region")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.MigrateServiceTag']]]:
        """
        tags.
        """
        return pulumi.get(self, "tags")

