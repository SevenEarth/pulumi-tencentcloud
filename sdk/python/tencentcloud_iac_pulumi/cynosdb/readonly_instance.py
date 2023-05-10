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
                 cluster_id: pulumi.Input[str],
                 instance_name: pulumi.Input[str],
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 instance_cpu_core: Optional[pulumi.Input[int]] = None,
                 instance_maintain_duration: Optional[pulumi.Input[int]] = None,
                 instance_maintain_start_time: Optional[pulumi.Input[int]] = None,
                 instance_maintain_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_memory_size: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ReadonlyInstance resource.
        :param pulumi.Input[str] cluster_id: Cluster ID which the readonly instance belongs to.
        :param pulumi.Input[str] instance_name: Name of instance.
        :param pulumi.Input[bool] force_delete: Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        :param pulumi.Input[int] instance_cpu_core: The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[int] instance_maintain_duration: Duration time for maintenance, unit in second. `3600` by default.
        :param pulumi.Input[int] instance_maintain_start_time: Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_maintain_weekdays: Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        :param pulumi.Input[int] instance_memory_size: Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "instance_name", instance_name)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if instance_cpu_core is not None:
            pulumi.set(__self__, "instance_cpu_core", instance_cpu_core)
        if instance_maintain_duration is not None:
            pulumi.set(__self__, "instance_maintain_duration", instance_maintain_duration)
        if instance_maintain_start_time is not None:
            pulumi.set(__self__, "instance_maintain_start_time", instance_maintain_start_time)
        if instance_maintain_weekdays is not None:
            pulumi.set(__self__, "instance_maintain_weekdays", instance_maintain_weekdays)
        if instance_memory_size is not None:
            pulumi.set(__self__, "instance_memory_size", instance_memory_size)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID which the readonly instance belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        Name of instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter(name="instanceCpuCore")
    def instance_cpu_core(self) -> Optional[pulumi.Input[int]]:
        """
        The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        return pulumi.get(self, "instance_cpu_core")

    @instance_cpu_core.setter
    def instance_cpu_core(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_cpu_core", value)

    @property
    @pulumi.getter(name="instanceMaintainDuration")
    def instance_maintain_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Duration time for maintenance, unit in second. `3600` by default.
        """
        return pulumi.get(self, "instance_maintain_duration")

    @instance_maintain_duration.setter
    def instance_maintain_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_maintain_duration", value)

    @property
    @pulumi.getter(name="instanceMaintainStartTime")
    def instance_maintain_start_time(self) -> Optional[pulumi.Input[int]]:
        """
        Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        """
        return pulumi.get(self, "instance_maintain_start_time")

    @instance_maintain_start_time.setter
    def instance_maintain_start_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_maintain_start_time", value)

    @property
    @pulumi.getter(name="instanceMaintainWeekdays")
    def instance_maintain_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        """
        return pulumi.get(self, "instance_maintain_weekdays")

    @instance_maintain_weekdays.setter
    def instance_maintain_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_maintain_weekdays", value)

    @property
    @pulumi.getter(name="instanceMemorySize")
    def instance_memory_size(self) -> Optional[pulumi.Input[int]]:
        """
        Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        return pulumi.get(self, "instance_memory_size")

    @instance_memory_size.setter
    def instance_memory_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_memory_size", value)


@pulumi.input_type
class _ReadonlyInstanceState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 instance_cpu_core: Optional[pulumi.Input[int]] = None,
                 instance_maintain_duration: Optional[pulumi.Input[int]] = None,
                 instance_maintain_start_time: Optional[pulumi.Input[int]] = None,
                 instance_maintain_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_memory_size: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 instance_status: Optional[pulumi.Input[str]] = None,
                 instance_storage_size: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ReadonlyInstance resources.
        :param pulumi.Input[str] cluster_id: Cluster ID which the readonly instance belongs to.
        :param pulumi.Input[bool] force_delete: Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        :param pulumi.Input[int] instance_cpu_core: The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[int] instance_maintain_duration: Duration time for maintenance, unit in second. `3600` by default.
        :param pulumi.Input[int] instance_maintain_start_time: Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_maintain_weekdays: Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        :param pulumi.Input[int] instance_memory_size: Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[str] instance_name: Name of instance.
        :param pulumi.Input[str] instance_status: Status of the instance.
        :param pulumi.Input[int] instance_storage_size: Storage size of the instance, unit in GB.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if instance_cpu_core is not None:
            pulumi.set(__self__, "instance_cpu_core", instance_cpu_core)
        if instance_maintain_duration is not None:
            pulumi.set(__self__, "instance_maintain_duration", instance_maintain_duration)
        if instance_maintain_start_time is not None:
            pulumi.set(__self__, "instance_maintain_start_time", instance_maintain_start_time)
        if instance_maintain_weekdays is not None:
            pulumi.set(__self__, "instance_maintain_weekdays", instance_maintain_weekdays)
        if instance_memory_size is not None:
            pulumi.set(__self__, "instance_memory_size", instance_memory_size)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if instance_status is not None:
            pulumi.set(__self__, "instance_status", instance_status)
        if instance_storage_size is not None:
            pulumi.set(__self__, "instance_storage_size", instance_storage_size)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID which the readonly instance belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter(name="instanceCpuCore")
    def instance_cpu_core(self) -> Optional[pulumi.Input[int]]:
        """
        The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        return pulumi.get(self, "instance_cpu_core")

    @instance_cpu_core.setter
    def instance_cpu_core(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_cpu_core", value)

    @property
    @pulumi.getter(name="instanceMaintainDuration")
    def instance_maintain_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Duration time for maintenance, unit in second. `3600` by default.
        """
        return pulumi.get(self, "instance_maintain_duration")

    @instance_maintain_duration.setter
    def instance_maintain_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_maintain_duration", value)

    @property
    @pulumi.getter(name="instanceMaintainStartTime")
    def instance_maintain_start_time(self) -> Optional[pulumi.Input[int]]:
        """
        Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        """
        return pulumi.get(self, "instance_maintain_start_time")

    @instance_maintain_start_time.setter
    def instance_maintain_start_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_maintain_start_time", value)

    @property
    @pulumi.getter(name="instanceMaintainWeekdays")
    def instance_maintain_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        """
        return pulumi.get(self, "instance_maintain_weekdays")

    @instance_maintain_weekdays.setter
    def instance_maintain_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_maintain_weekdays", value)

    @property
    @pulumi.getter(name="instanceMemorySize")
    def instance_memory_size(self) -> Optional[pulumi.Input[int]]:
        """
        Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        return pulumi.get(self, "instance_memory_size")

    @instance_memory_size.setter
    def instance_memory_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_memory_size", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the instance.
        """
        return pulumi.get(self, "instance_status")

    @instance_status.setter
    def instance_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_status", value)

    @property
    @pulumi.getter(name="instanceStorageSize")
    def instance_storage_size(self) -> Optional[pulumi.Input[int]]:
        """
        Storage size of the instance, unit in GB.
        """
        return pulumi.get(self, "instance_storage_size")

    @instance_storage_size.setter
    def instance_storage_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_storage_size", value)


class ReadonlyInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 instance_cpu_core: Optional[pulumi.Input[int]] = None,
                 instance_maintain_duration: Optional[pulumi.Input[int]] = None,
                 instance_maintain_start_time: Optional[pulumi.Input[int]] = None,
                 instance_maintain_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_memory_size: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide a resource to create a CynosDB readonly instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.cynosdb.ReadonlyInstance("foo",
            cluster_id=cynosdbmysql_dzj5l8gz,
            instance_name="tf-cynosdb-readonly-instance",
            force_delete=True,
            instance_cpu_core=2,
            instance_memory_size=4,
            instance_maintain_duration=7200,
            instance_maintain_start_time=21600,
            instance_maintain_weekdays=[
                "Fri",
                "Mon",
                "Sat",
                "Sun",
                "Thu",
                "Wed",
                "Tue",
            ])
        ```

        ## Import

        CynosDB readonly instance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance foo cynosdbmysql-ins-dhwynib6
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID which the readonly instance belongs to.
        :param pulumi.Input[bool] force_delete: Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        :param pulumi.Input[int] instance_cpu_core: The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[int] instance_maintain_duration: Duration time for maintenance, unit in second. `3600` by default.
        :param pulumi.Input[int] instance_maintain_start_time: Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_maintain_weekdays: Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        :param pulumi.Input[int] instance_memory_size: Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[str] instance_name: Name of instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReadonlyInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a CynosDB readonly instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.cynosdb.ReadonlyInstance("foo",
            cluster_id=cynosdbmysql_dzj5l8gz,
            instance_name="tf-cynosdb-readonly-instance",
            force_delete=True,
            instance_cpu_core=2,
            instance_memory_size=4,
            instance_maintain_duration=7200,
            instance_maintain_start_time=21600,
            instance_maintain_weekdays=[
                "Fri",
                "Mon",
                "Sat",
                "Sun",
                "Thu",
                "Wed",
                "Tue",
            ])
        ```

        ## Import

        CynosDB readonly instance can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance foo cynosdbmysql-ins-dhwynib6
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
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 instance_cpu_core: Optional[pulumi.Input[int]] = None,
                 instance_maintain_duration: Optional[pulumi.Input[int]] = None,
                 instance_maintain_start_time: Optional[pulumi.Input[int]] = None,
                 instance_maintain_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_memory_size: Optional[pulumi.Input[int]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
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

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["force_delete"] = force_delete
            __props__.__dict__["instance_cpu_core"] = instance_cpu_core
            __props__.__dict__["instance_maintain_duration"] = instance_maintain_duration
            __props__.__dict__["instance_maintain_start_time"] = instance_maintain_start_time
            __props__.__dict__["instance_maintain_weekdays"] = instance_maintain_weekdays
            __props__.__dict__["instance_memory_size"] = instance_memory_size
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            __props__.__dict__["instance_status"] = None
            __props__.__dict__["instance_storage_size"] = None
        super(ReadonlyInstance, __self__).__init__(
            'tencentcloud:Cynosdb/readonlyInstance:ReadonlyInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            instance_cpu_core: Optional[pulumi.Input[int]] = None,
            instance_maintain_duration: Optional[pulumi.Input[int]] = None,
            instance_maintain_start_time: Optional[pulumi.Input[int]] = None,
            instance_maintain_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_memory_size: Optional[pulumi.Input[int]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            instance_status: Optional[pulumi.Input[str]] = None,
            instance_storage_size: Optional[pulumi.Input[int]] = None) -> 'ReadonlyInstance':
        """
        Get an existing ReadonlyInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID which the readonly instance belongs to.
        :param pulumi.Input[bool] force_delete: Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        :param pulumi.Input[int] instance_cpu_core: The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[int] instance_maintain_duration: Duration time for maintenance, unit in second. `3600` by default.
        :param pulumi.Input[int] instance_maintain_start_time: Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_maintain_weekdays: Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        :param pulumi.Input[int] instance_memory_size: Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        :param pulumi.Input[str] instance_name: Name of instance.
        :param pulumi.Input[str] instance_status: Status of the instance.
        :param pulumi.Input[int] instance_storage_size: Storage size of the instance, unit in GB.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReadonlyInstanceState.__new__(_ReadonlyInstanceState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["force_delete"] = force_delete
        __props__.__dict__["instance_cpu_core"] = instance_cpu_core
        __props__.__dict__["instance_maintain_duration"] = instance_maintain_duration
        __props__.__dict__["instance_maintain_start_time"] = instance_maintain_start_time
        __props__.__dict__["instance_maintain_weekdays"] = instance_maintain_weekdays
        __props__.__dict__["instance_memory_size"] = instance_memory_size
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["instance_status"] = instance_status
        __props__.__dict__["instance_storage_size"] = instance_storage_size
        return ReadonlyInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID which the readonly instance belongs to.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both `PREPAID` and `POSTPAID_BY_HOUR` cluster.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter(name="instanceCpuCore")
    def instance_cpu_core(self) -> pulumi.Output[Optional[int]]:
        """
        The number of CPU cores of read-write type instance in the CynosDB cluster. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        return pulumi.get(self, "instance_cpu_core")

    @property
    @pulumi.getter(name="instanceMaintainDuration")
    def instance_maintain_duration(self) -> pulumi.Output[Optional[int]]:
        """
        Duration time for maintenance, unit in second. `3600` by default.
        """
        return pulumi.get(self, "instance_maintain_duration")

    @property
    @pulumi.getter(name="instanceMaintainStartTime")
    def instance_maintain_start_time(self) -> pulumi.Output[Optional[int]]:
        """
        Offset time from 00:00, unit in second. For example, 03:00am should be `10800`. `10800` by default.
        """
        return pulumi.get(self, "instance_maintain_start_time")

    @property
    @pulumi.getter(name="instanceMaintainWeekdays")
    def instance_maintain_weekdays(self) -> pulumi.Output[Sequence[str]]:
        """
        Weekdays for maintenance. `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` by default.
        """
        return pulumi.get(self, "instance_maintain_weekdays")

    @property
    @pulumi.getter(name="instanceMemorySize")
    def instance_memory_size(self) -> pulumi.Output[Optional[int]]:
        """
        Memory capacity of read-write type instance, unit in GB. Required while creating normal cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.
        """
        return pulumi.get(self, "instance_memory_size")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        Name of instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> pulumi.Output[str]:
        """
        Status of the instance.
        """
        return pulumi.get(self, "instance_status")

    @property
    @pulumi.getter(name="instanceStorageSize")
    def instance_storage_size(self) -> pulumi.Output[int]:
        """
        Storage size of the instance, unit in GB.
        """
        return pulumi.get(self, "instance_storage_size")

