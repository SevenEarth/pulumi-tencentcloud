# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceHaConfigArgs', 'InstanceHaConfig']

@pulumi.input_type
class InstanceHaConfigArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 max_standby_lag: pulumi.Input[int],
                 max_standby_latency: pulumi.Input[int],
                 sync_mode: pulumi.Input[str],
                 max_sync_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_latency: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a InstanceHaConfig resource.
        :param pulumi.Input[str] instance_id: instance id.
        :param pulumi.Input[int] max_standby_lag: Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        :param pulumi.Input[int] max_standby_latency: Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        :param pulumi.Input[str] sync_mode: Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        :param pulumi.Input[int] max_sync_standby_lag: Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[int] max_sync_standby_latency: Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "max_standby_lag", max_standby_lag)
        pulumi.set(__self__, "max_standby_latency", max_standby_latency)
        pulumi.set(__self__, "sync_mode", sync_mode)
        if max_sync_standby_lag is not None:
            pulumi.set(__self__, "max_sync_standby_lag", max_sync_standby_lag)
        if max_sync_standby_latency is not None:
            pulumi.set(__self__, "max_sync_standby_latency", max_sync_standby_latency)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxStandbyLag")
    def max_standby_lag(self) -> pulumi.Input[int]:
        """
        Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        """
        return pulumi.get(self, "max_standby_lag")

    @max_standby_lag.setter
    def max_standby_lag(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_standby_lag", value)

    @property
    @pulumi.getter(name="maxStandbyLatency")
    def max_standby_latency(self) -> pulumi.Input[int]:
        """
        Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        """
        return pulumi.get(self, "max_standby_latency")

    @max_standby_latency.setter
    def max_standby_latency(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_standby_latency", value)

    @property
    @pulumi.getter(name="syncMode")
    def sync_mode(self) -> pulumi.Input[str]:
        """
        Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        """
        return pulumi.get(self, "sync_mode")

    @sync_mode.setter
    def sync_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "sync_mode", value)

    @property
    @pulumi.getter(name="maxSyncStandbyLag")
    def max_sync_standby_lag(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        return pulumi.get(self, "max_sync_standby_lag")

    @max_sync_standby_lag.setter
    def max_sync_standby_lag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_sync_standby_lag", value)

    @property
    @pulumi.getter(name="maxSyncStandbyLatency")
    def max_sync_standby_latency(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        return pulumi.get(self, "max_sync_standby_latency")

    @max_sync_standby_latency.setter
    def max_sync_standby_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_sync_standby_latency", value)


@pulumi.input_type
class _InstanceHaConfigState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_standby_latency: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_latency: Optional[pulumi.Input[int]] = None,
                 sync_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceHaConfig resources.
        :param pulumi.Input[str] instance_id: instance id.
        :param pulumi.Input[int] max_standby_lag: Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        :param pulumi.Input[int] max_standby_latency: Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        :param pulumi.Input[int] max_sync_standby_lag: Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[int] max_sync_standby_latency: Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[str] sync_mode: Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_standby_lag is not None:
            pulumi.set(__self__, "max_standby_lag", max_standby_lag)
        if max_standby_latency is not None:
            pulumi.set(__self__, "max_standby_latency", max_standby_latency)
        if max_sync_standby_lag is not None:
            pulumi.set(__self__, "max_sync_standby_lag", max_sync_standby_lag)
        if max_sync_standby_latency is not None:
            pulumi.set(__self__, "max_sync_standby_latency", max_sync_standby_latency)
        if sync_mode is not None:
            pulumi.set(__self__, "sync_mode", sync_mode)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxStandbyLag")
    def max_standby_lag(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        """
        return pulumi.get(self, "max_standby_lag")

    @max_standby_lag.setter
    def max_standby_lag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_standby_lag", value)

    @property
    @pulumi.getter(name="maxStandbyLatency")
    def max_standby_latency(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        """
        return pulumi.get(self, "max_standby_latency")

    @max_standby_latency.setter
    def max_standby_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_standby_latency", value)

    @property
    @pulumi.getter(name="maxSyncStandbyLag")
    def max_sync_standby_lag(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        return pulumi.get(self, "max_sync_standby_lag")

    @max_sync_standby_lag.setter
    def max_sync_standby_lag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_sync_standby_lag", value)

    @property
    @pulumi.getter(name="maxSyncStandbyLatency")
    def max_sync_standby_latency(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        return pulumi.get(self, "max_sync_standby_latency")

    @max_sync_standby_latency.setter
    def max_sync_standby_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_sync_standby_latency", value)

    @property
    @pulumi.getter(name="syncMode")
    def sync_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        """
        return pulumi.get(self, "sync_mode")

    @sync_mode.setter
    def sync_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_mode", value)


class InstanceHaConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_standby_latency: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_latency: Optional[pulumi.Input[int]] = None,
                 sync_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to set postgresql instance syncMode

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.postgresql.InstanceHaConfig("example",
            instance_id="postgres-gzg9jb2n",
            max_standby_lag=10,
            max_standby_latency=10737418240,
            max_sync_standby_lag=5,
            max_sync_standby_latency=52428800,
            sync_mode="Semi-sync")
        ```

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.postgresql.InstanceHaConfig("example",
            instance_id="postgres-gzg9jb2n",
            max_standby_lag=10,
            max_standby_latency=10737418240,
            sync_mode="Async")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: instance id.
        :param pulumi.Input[int] max_standby_lag: Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        :param pulumi.Input[int] max_standby_latency: Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        :param pulumi.Input[int] max_sync_standby_lag: Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[int] max_sync_standby_latency: Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[str] sync_mode: Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceHaConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to set postgresql instance syncMode

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.postgresql.InstanceHaConfig("example",
            instance_id="postgres-gzg9jb2n",
            max_standby_lag=10,
            max_standby_latency=10737418240,
            max_sync_standby_lag=5,
            max_sync_standby_latency=52428800,
            sync_mode="Semi-sync")
        ```

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.postgresql.InstanceHaConfig("example",
            instance_id="postgres-gzg9jb2n",
            max_standby_lag=10,
            max_standby_latency=10737418240,
            sync_mode="Async")
        ```

        :param str resource_name: The name of the resource.
        :param InstanceHaConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceHaConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_standby_latency: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_lag: Optional[pulumi.Input[int]] = None,
                 max_sync_standby_latency: Optional[pulumi.Input[int]] = None,
                 sync_mode: Optional[pulumi.Input[str]] = None,
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
            __props__ = InstanceHaConfigArgs.__new__(InstanceHaConfigArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if max_standby_lag is None and not opts.urn:
                raise TypeError("Missing required property 'max_standby_lag'")
            __props__.__dict__["max_standby_lag"] = max_standby_lag
            if max_standby_latency is None and not opts.urn:
                raise TypeError("Missing required property 'max_standby_latency'")
            __props__.__dict__["max_standby_latency"] = max_standby_latency
            __props__.__dict__["max_sync_standby_lag"] = max_sync_standby_lag
            __props__.__dict__["max_sync_standby_latency"] = max_sync_standby_latency
            if sync_mode is None and not opts.urn:
                raise TypeError("Missing required property 'sync_mode'")
            __props__.__dict__["sync_mode"] = sync_mode
        super(InstanceHaConfig, __self__).__init__(
            'tencentcloud:Postgresql/instanceHaConfig:InstanceHaConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            max_standby_lag: Optional[pulumi.Input[int]] = None,
            max_standby_latency: Optional[pulumi.Input[int]] = None,
            max_sync_standby_lag: Optional[pulumi.Input[int]] = None,
            max_sync_standby_latency: Optional[pulumi.Input[int]] = None,
            sync_mode: Optional[pulumi.Input[str]] = None) -> 'InstanceHaConfig':
        """
        Get an existing InstanceHaConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: instance id.
        :param pulumi.Input[int] max_standby_lag: Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        :param pulumi.Input[int] max_standby_latency: Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        :param pulumi.Input[int] max_sync_standby_lag: Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[int] max_sync_standby_latency: Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        :param pulumi.Input[str] sync_mode: Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceHaConfigState.__new__(_InstanceHaConfigState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_standby_lag"] = max_standby_lag
        __props__.__dict__["max_standby_latency"] = max_standby_latency
        __props__.__dict__["max_sync_standby_lag"] = max_sync_standby_lag
        __props__.__dict__["max_sync_standby_latency"] = max_sync_standby_latency
        __props__.__dict__["sync_mode"] = sync_mode
        return InstanceHaConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxStandbyLag")
    def max_standby_lag(self) -> pulumi.Output[int]:
        """
        Maximum latency of highly available backup machines. When the delay time of the backup node is less than or equal to this value, and the amount of delay data of the backup node is less than or equal to MaxStandbyLatency, the primary node can be switched. Unit: s; Parameter range: [5, 10].
        """
        return pulumi.get(self, "max_standby_lag")

    @property
    @pulumi.getter(name="maxStandbyLatency")
    def max_standby_latency(self) -> pulumi.Output[int]:
        """
        Maximum latency data volume for highly available backup machines. When the delay data amount of the backup node is less than or equal to this value, and the delay time of the backup node is less than or equal to MaxStandbyLag, it can switch to the main node. Unit: byte; Parameter range: [1073741824, 322122547200].
        """
        return pulumi.get(self, "max_standby_latency")

    @property
    @pulumi.getter(name="maxSyncStandbyLag")
    def max_sync_standby_lag(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum delay time for synchronous backup. When the delay time of the standby machine is less than or equal to this value, and the amount of delay data of the standby machine is less than or equal to MaxSyncStandbyLatency, then the standby machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        return pulumi.get(self, "max_sync_standby_lag")

    @property
    @pulumi.getter(name="maxSyncStandbyLatency")
    def max_sync_standby_latency(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum latency data for synchronous backup. When the amount of data delayed by the backup machine is less than or equal to this value, and the delay time of the backup machine is less than or equal to MaxSyncStandbyLag, then the backup machine adopts synchronous replication; Otherwise, adopt asynchronous replication. This parameter value is valid for instances where SyncMode is set to Semi sync. When a semi synchronous instance prohibits degradation to asynchronous replication, MaxSyncStandbyLatency and MaxSyncStandbyLag are not set. When semi synchronous instances allow degenerate asynchronous replication, PostgreSQL version 9 instances must have MaxSyncStandbyLatency set and MaxSyncStandbyLag not set, while PostgreSQL version 10 and above instances must have MaxSyncStandbyLatency and MaxSyncStandbyLag set.
        """
        return pulumi.get(self, "max_sync_standby_latency")

    @property
    @pulumi.getter(name="syncMode")
    def sync_mode(self) -> pulumi.Output[str]:
        """
        Master slave synchronization method, Semi-sync: Semi synchronous; Async: Asynchronous. Main instance default value: Semi-sync, Read-only instance default value: Async.
        """
        return pulumi.get(self, "sync_mode")
