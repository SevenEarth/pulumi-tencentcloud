# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BackupConfigArgs', 'BackupConfig']

@pulumi.input_type
class BackupConfigArgs:
    def __init__(__self__, *,
                 backup_time: pulumi.Input[str],
                 redis_id: pulumi.Input[str],
                 backup_periods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a BackupConfig resource.
        :param pulumi.Input[str] backup_time: Specifys what time the backup action should take place. And the time interval should be one hour.
        :param pulumi.Input[str] redis_id: ID of a redis instance to which the policy will be applied.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_periods: It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        pulumi.set(__self__, "backup_time", backup_time)
        pulumi.set(__self__, "redis_id", redis_id)
        if backup_periods is not None:
            warnings.warn("""It has been deprecated from version 1.58.2. It makes no difference to online config at all""", DeprecationWarning)
            pulumi.log.warn("""backup_periods is deprecated: It has been deprecated from version 1.58.2. It makes no difference to online config at all""")
        if backup_periods is not None:
            pulumi.set(__self__, "backup_periods", backup_periods)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> pulumi.Input[str]:
        """
        Specifys what time the backup action should take place. And the time interval should be one hour.
        """
        return pulumi.get(self, "backup_time")

    @backup_time.setter
    def backup_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_time", value)

    @property
    @pulumi.getter(name="redisId")
    def redis_id(self) -> pulumi.Input[str]:
        """
        ID of a redis instance to which the policy will be applied.
        """
        return pulumi.get(self, "redis_id")

    @redis_id.setter
    def redis_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "redis_id", value)

    @property
    @pulumi.getter(name="backupPeriods")
    def backup_periods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        warnings.warn("""It has been deprecated from version 1.58.2. It makes no difference to online config at all""", DeprecationWarning)
        pulumi.log.warn("""backup_periods is deprecated: It has been deprecated from version 1.58.2. It makes no difference to online config at all""")

        return pulumi.get(self, "backup_periods")

    @backup_periods.setter
    def backup_periods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backup_periods", value)


@pulumi.input_type
class _BackupConfigState:
    def __init__(__self__, *,
                 backup_periods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_time: Optional[pulumi.Input[str]] = None,
                 redis_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BackupConfig resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_periods: It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        :param pulumi.Input[str] backup_time: Specifys what time the backup action should take place. And the time interval should be one hour.
        :param pulumi.Input[str] redis_id: ID of a redis instance to which the policy will be applied.
        """
        if backup_periods is not None:
            warnings.warn("""It has been deprecated from version 1.58.2. It makes no difference to online config at all""", DeprecationWarning)
            pulumi.log.warn("""backup_periods is deprecated: It has been deprecated from version 1.58.2. It makes no difference to online config at all""")
        if backup_periods is not None:
            pulumi.set(__self__, "backup_periods", backup_periods)
        if backup_time is not None:
            pulumi.set(__self__, "backup_time", backup_time)
        if redis_id is not None:
            pulumi.set(__self__, "redis_id", redis_id)

    @property
    @pulumi.getter(name="backupPeriods")
    def backup_periods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        warnings.warn("""It has been deprecated from version 1.58.2. It makes no difference to online config at all""", DeprecationWarning)
        pulumi.log.warn("""backup_periods is deprecated: It has been deprecated from version 1.58.2. It makes no difference to online config at all""")

        return pulumi.get(self, "backup_periods")

    @backup_periods.setter
    def backup_periods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backup_periods", value)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> Optional[pulumi.Input[str]]:
        """
        Specifys what time the backup action should take place. And the time interval should be one hour.
        """
        return pulumi.get(self, "backup_time")

    @backup_time.setter
    def backup_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_time", value)

    @property
    @pulumi.getter(name="redisId")
    def redis_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a redis instance to which the policy will be applied.
        """
        return pulumi.get(self, "redis_id")

    @redis_id.setter
    def redis_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redis_id", value)


class BackupConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_periods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_time: Optional[pulumi.Input[str]] = None,
                 redis_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create a backup config of redis.

        ## Example Usage

        ### Set configuration for automatic backups

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zone = tencentcloud.Redis.get_zone_config(type_id=7)
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=zone.lists[1].zone,
            cidr_block="10.0.1.0/24")
        foo_instance = tencentcloud.redis.Instance("fooInstance",
            availability_zone=zone.lists[1].zone,
            type_id=zone.lists[1].type_id,
            password="test12345789",
            mem_size=8192,
            redis_shard_num=zone.lists[1].redis_shard_nums[0],
            redis_replicas_num=zone.lists[1].redis_replicas_nums[0],
            port=6379,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        foo_backup_config = tencentcloud.redis.BackupConfig("fooBackupConfig",
            redis_id=foo_instance.id,
            backup_time="04:00-05:00",
            backup_periods=["Monday"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Redis  backup config can be imported, e.g.

        ```sh
        $ pulumi import tencentcloud:Redis/backupConfig:BackupConfig foo redis-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_periods: It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        :param pulumi.Input[str] backup_time: Specifys what time the backup action should take place. And the time interval should be one hour.
        :param pulumi.Input[str] redis_id: ID of a redis instance to which the policy will be applied.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create a backup config of redis.

        ## Example Usage

        ### Set configuration for automatic backups

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zone = tencentcloud.Redis.get_zone_config(type_id=7)
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=zone.lists[1].zone,
            cidr_block="10.0.1.0/24")
        foo_instance = tencentcloud.redis.Instance("fooInstance",
            availability_zone=zone.lists[1].zone,
            type_id=zone.lists[1].type_id,
            password="test12345789",
            mem_size=8192,
            redis_shard_num=zone.lists[1].redis_shard_nums[0],
            redis_replicas_num=zone.lists[1].redis_replicas_nums[0],
            port=6379,
            vpc_id=vpc.id,
            subnet_id=subnet.id)
        foo_backup_config = tencentcloud.redis.BackupConfig("fooBackupConfig",
            redis_id=foo_instance.id,
            backup_time="04:00-05:00",
            backup_periods=["Monday"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Redis  backup config can be imported, e.g.

        ```sh
        $ pulumi import tencentcloud:Redis/backupConfig:BackupConfig foo redis-id
        ```

        :param str resource_name: The name of the resource.
        :param BackupConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_periods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_time: Optional[pulumi.Input[str]] = None,
                 redis_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupConfigArgs.__new__(BackupConfigArgs)

            __props__.__dict__["backup_periods"] = backup_periods
            if backup_time is None and not opts.urn:
                raise TypeError("Missing required property 'backup_time'")
            __props__.__dict__["backup_time"] = backup_time
            if redis_id is None and not opts.urn:
                raise TypeError("Missing required property 'redis_id'")
            __props__.__dict__["redis_id"] = redis_id
        super(BackupConfig, __self__).__init__(
            'tencentcloud:Redis/backupConfig:BackupConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_periods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            backup_time: Optional[pulumi.Input[str]] = None,
            redis_id: Optional[pulumi.Input[str]] = None) -> 'BackupConfig':
        """
        Get an existing BackupConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] backup_periods: It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        :param pulumi.Input[str] backup_time: Specifys what time the backup action should take place. And the time interval should be one hour.
        :param pulumi.Input[str] redis_id: ID of a redis instance to which the policy will be applied.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupConfigState.__new__(_BackupConfigState)

        __props__.__dict__["backup_periods"] = backup_periods
        __props__.__dict__["backup_time"] = backup_time
        __props__.__dict__["redis_id"] = redis_id
        return BackupConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPeriods")
    def backup_periods(self) -> pulumi.Output[Sequence[str]]:
        """
        It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
        """
        warnings.warn("""It has been deprecated from version 1.58.2. It makes no difference to online config at all""", DeprecationWarning)
        pulumi.log.warn("""backup_periods is deprecated: It has been deprecated from version 1.58.2. It makes no difference to online config at all""")

        return pulumi.get(self, "backup_periods")

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> pulumi.Output[str]:
        """
        Specifys what time the backup action should take place. And the time interval should be one hour.
        """
        return pulumi.get(self, "backup_time")

    @property
    @pulumi.getter(name="redisId")
    def redis_id(self) -> pulumi.Output[str]:
        """
        ID of a redis instance to which the policy will be applied.
        """
        return pulumi.get(self, "redis_id")

