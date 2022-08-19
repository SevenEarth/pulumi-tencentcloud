# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReadonlyGroupArgs', 'ReadonlyGroup']

@pulumi.input_type
class ReadonlyGroupArgs:
    def __init__(__self__, *,
                 master_db_instance_id: pulumi.Input[str],
                 max_replay_lag: pulumi.Input[int],
                 max_replay_latency: pulumi.Input[int],
                 min_delay_eliminate_reserve: pulumi.Input[int],
                 project_id: pulumi.Input[int],
                 replay_lag_eliminate: pulumi.Input[int],
                 replay_latency_eliminate: pulumi.Input[int],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ReadonlyGroup resource.
        :param pulumi.Input[str] master_db_instance_id: Primary instance ID.
        :param pulumi.Input[int] max_replay_lag: Delay threshold in ms.
        :param pulumi.Input[int] max_replay_latency: Delayed log size threshold in MB.
        :param pulumi.Input[int] min_delay_eliminate_reserve: The minimum number of read-only replicas that must be retained in an RO group.
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[int] replay_lag_eliminate: Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[int] replay_latency_eliminate: Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        :param pulumi.Input[str] name: RO group name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        """
        pulumi.set(__self__, "master_db_instance_id", master_db_instance_id)
        pulumi.set(__self__, "max_replay_lag", max_replay_lag)
        pulumi.set(__self__, "max_replay_latency", max_replay_latency)
        pulumi.set(__self__, "min_delay_eliminate_reserve", min_delay_eliminate_reserve)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "replay_lag_eliminate", replay_lag_eliminate)
        pulumi.set(__self__, "replay_latency_eliminate", replay_latency_eliminate)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if security_groups_ids is not None:
            pulumi.set(__self__, "security_groups_ids", security_groups_ids)

    @property
    @pulumi.getter(name="masterDbInstanceId")
    def master_db_instance_id(self) -> pulumi.Input[str]:
        """
        Primary instance ID.
        """
        return pulumi.get(self, "master_db_instance_id")

    @master_db_instance_id.setter
    def master_db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "master_db_instance_id", value)

    @property
    @pulumi.getter(name="maxReplayLag")
    def max_replay_lag(self) -> pulumi.Input[int]:
        """
        Delay threshold in ms.
        """
        return pulumi.get(self, "max_replay_lag")

    @max_replay_lag.setter
    def max_replay_lag(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_replay_lag", value)

    @property
    @pulumi.getter(name="maxReplayLatency")
    def max_replay_latency(self) -> pulumi.Input[int]:
        """
        Delayed log size threshold in MB.
        """
        return pulumi.get(self, "max_replay_latency")

    @max_replay_latency.setter
    def max_replay_latency(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_replay_latency", value)

    @property
    @pulumi.getter(name="minDelayEliminateReserve")
    def min_delay_eliminate_reserve(self) -> pulumi.Input[int]:
        """
        The minimum number of read-only replicas that must be retained in an RO group.
        """
        return pulumi.get(self, "min_delay_eliminate_reserve")

    @min_delay_eliminate_reserve.setter
    def min_delay_eliminate_reserve(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_delay_eliminate_reserve", value)

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
    @pulumi.getter(name="replayLagEliminate")
    def replay_lag_eliminate(self) -> pulumi.Input[int]:
        """
        Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        """
        return pulumi.get(self, "replay_lag_eliminate")

    @replay_lag_eliminate.setter
    def replay_lag_eliminate(self, value: pulumi.Input[int]):
        pulumi.set(self, "replay_lag_eliminate", value)

    @property
    @pulumi.getter(name="replayLatencyEliminate")
    def replay_latency_eliminate(self) -> pulumi.Input[int]:
        """
        Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        """
        return pulumi.get(self, "replay_latency_eliminate")

    @replay_latency_eliminate.setter
    def replay_latency_eliminate(self, value: pulumi.Input[int]):
        pulumi.set(self, "replay_latency_eliminate", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        RO group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="securityGroupsIds")
    def security_groups_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        """
        return pulumi.get(self, "security_groups_ids")

    @security_groups_ids.setter
    def security_groups_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups_ids", value)


@pulumi.input_type
class _ReadonlyGroupState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 master_db_instance_id: Optional[pulumi.Input[str]] = None,
                 max_replay_lag: Optional[pulumi.Input[int]] = None,
                 max_replay_latency: Optional[pulumi.Input[int]] = None,
                 min_delay_eliminate_reserve: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 replay_lag_eliminate: Optional[pulumi.Input[int]] = None,
                 replay_latency_eliminate: Optional[pulumi.Input[int]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReadonlyGroup resources.
        :param pulumi.Input[str] create_time: Create time of the postgresql instance.
        :param pulumi.Input[str] master_db_instance_id: Primary instance ID.
        :param pulumi.Input[int] max_replay_lag: Delay threshold in ms.
        :param pulumi.Input[int] max_replay_latency: Delayed log size threshold in MB.
        :param pulumi.Input[int] min_delay_eliminate_reserve: The minimum number of read-only replicas that must be retained in an RO group.
        :param pulumi.Input[str] name: RO group name.
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[int] replay_lag_eliminate: Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[int] replay_latency_eliminate: Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if master_db_instance_id is not None:
            pulumi.set(__self__, "master_db_instance_id", master_db_instance_id)
        if max_replay_lag is not None:
            pulumi.set(__self__, "max_replay_lag", max_replay_lag)
        if max_replay_latency is not None:
            pulumi.set(__self__, "max_replay_latency", max_replay_latency)
        if min_delay_eliminate_reserve is not None:
            pulumi.set(__self__, "min_delay_eliminate_reserve", min_delay_eliminate_reserve)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if replay_lag_eliminate is not None:
            pulumi.set(__self__, "replay_lag_eliminate", replay_lag_eliminate)
        if replay_latency_eliminate is not None:
            pulumi.set(__self__, "replay_latency_eliminate", replay_latency_eliminate)
        if security_groups_ids is not None:
            pulumi.set(__self__, "security_groups_ids", security_groups_ids)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

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
    @pulumi.getter(name="masterDbInstanceId")
    def master_db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Primary instance ID.
        """
        return pulumi.get(self, "master_db_instance_id")

    @master_db_instance_id.setter
    def master_db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_db_instance_id", value)

    @property
    @pulumi.getter(name="maxReplayLag")
    def max_replay_lag(self) -> Optional[pulumi.Input[int]]:
        """
        Delay threshold in ms.
        """
        return pulumi.get(self, "max_replay_lag")

    @max_replay_lag.setter
    def max_replay_lag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_replay_lag", value)

    @property
    @pulumi.getter(name="maxReplayLatency")
    def max_replay_latency(self) -> Optional[pulumi.Input[int]]:
        """
        Delayed log size threshold in MB.
        """
        return pulumi.get(self, "max_replay_latency")

    @max_replay_latency.setter
    def max_replay_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_replay_latency", value)

    @property
    @pulumi.getter(name="minDelayEliminateReserve")
    def min_delay_eliminate_reserve(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of read-only replicas that must be retained in an RO group.
        """
        return pulumi.get(self, "min_delay_eliminate_reserve")

    @min_delay_eliminate_reserve.setter
    def min_delay_eliminate_reserve(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_delay_eliminate_reserve", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        RO group name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter(name="replayLagEliminate")
    def replay_lag_eliminate(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        """
        return pulumi.get(self, "replay_lag_eliminate")

    @replay_lag_eliminate.setter
    def replay_lag_eliminate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replay_lag_eliminate", value)

    @property
    @pulumi.getter(name="replayLatencyEliminate")
    def replay_latency_eliminate(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        """
        return pulumi.get(self, "replay_latency_eliminate")

    @replay_latency_eliminate.setter
    def replay_latency_eliminate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replay_latency_eliminate", value)

    @property
    @pulumi.getter(name="securityGroupsIds")
    def security_groups_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        """
        return pulumi.get(self, "security_groups_ids")

    @security_groups_ids.setter
    def security_groups_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups_ids", value)

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


class ReadonlyGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 master_db_instance_id: Optional[pulumi.Input[str]] = None,
                 max_replay_lag: Optional[pulumi.Input[int]] = None,
                 max_replay_latency: Optional[pulumi.Input[int]] = None,
                 min_delay_eliminate_reserve: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 replay_lag_eliminate: Optional[pulumi.Input[int]] = None,
                 replay_latency_eliminate: Optional[pulumi.Input[int]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create postgresql readonly group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        group = tencentcloud.postgresql.ReadonlyGroup("group",
            master_db_instance_id="postgres-f44wlfdv",
            max_replay_lag=100,
            max_replay_latency=512,
            min_delay_eliminate_reserve=1,
            project_id=0,
            replay_lag_eliminate=1,
            replay_latency_eliminate=1,
            subnet_id="subnet-enm92y0m",
            vpc_id="vpc-86v957zb")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] master_db_instance_id: Primary instance ID.
        :param pulumi.Input[int] max_replay_lag: Delay threshold in ms.
        :param pulumi.Input[int] max_replay_latency: Delayed log size threshold in MB.
        :param pulumi.Input[int] min_delay_eliminate_reserve: The minimum number of read-only replicas that must be retained in an RO group.
        :param pulumi.Input[str] name: RO group name.
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[int] replay_lag_eliminate: Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[int] replay_latency_eliminate: Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReadonlyGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create postgresql readonly group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        group = tencentcloud.postgresql.ReadonlyGroup("group",
            master_db_instance_id="postgres-f44wlfdv",
            max_replay_lag=100,
            max_replay_latency=512,
            min_delay_eliminate_reserve=1,
            project_id=0,
            replay_lag_eliminate=1,
            replay_latency_eliminate=1,
            subnet_id="subnet-enm92y0m",
            vpc_id="vpc-86v957zb")
        ```

        :param str resource_name: The name of the resource.
        :param ReadonlyGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReadonlyGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 master_db_instance_id: Optional[pulumi.Input[str]] = None,
                 max_replay_lag: Optional[pulumi.Input[int]] = None,
                 max_replay_latency: Optional[pulumi.Input[int]] = None,
                 min_delay_eliminate_reserve: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 replay_lag_eliminate: Optional[pulumi.Input[int]] = None,
                 replay_latency_eliminate: Optional[pulumi.Input[int]] = None,
                 security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReadonlyGroupArgs.__new__(ReadonlyGroupArgs)

            if master_db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'master_db_instance_id'")
            __props__.__dict__["master_db_instance_id"] = master_db_instance_id
            if max_replay_lag is None and not opts.urn:
                raise TypeError("Missing required property 'max_replay_lag'")
            __props__.__dict__["max_replay_lag"] = max_replay_lag
            if max_replay_latency is None and not opts.urn:
                raise TypeError("Missing required property 'max_replay_latency'")
            __props__.__dict__["max_replay_latency"] = max_replay_latency
            if min_delay_eliminate_reserve is None and not opts.urn:
                raise TypeError("Missing required property 'min_delay_eliminate_reserve'")
            __props__.__dict__["min_delay_eliminate_reserve"] = min_delay_eliminate_reserve
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if replay_lag_eliminate is None and not opts.urn:
                raise TypeError("Missing required property 'replay_lag_eliminate'")
            __props__.__dict__["replay_lag_eliminate"] = replay_lag_eliminate
            if replay_latency_eliminate is None and not opts.urn:
                raise TypeError("Missing required property 'replay_latency_eliminate'")
            __props__.__dict__["replay_latency_eliminate"] = replay_latency_eliminate
            __props__.__dict__["security_groups_ids"] = security_groups_ids
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["create_time"] = None
        super(ReadonlyGroup, __self__).__init__(
            'tencentcloud:Postgresql/readonlyGroup:ReadonlyGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            master_db_instance_id: Optional[pulumi.Input[str]] = None,
            max_replay_lag: Optional[pulumi.Input[int]] = None,
            max_replay_latency: Optional[pulumi.Input[int]] = None,
            min_delay_eliminate_reserve: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            replay_lag_eliminate: Optional[pulumi.Input[int]] = None,
            replay_latency_eliminate: Optional[pulumi.Input[int]] = None,
            security_groups_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'ReadonlyGroup':
        """
        Get an existing ReadonlyGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time of the postgresql instance.
        :param pulumi.Input[str] master_db_instance_id: Primary instance ID.
        :param pulumi.Input[int] max_replay_lag: Delay threshold in ms.
        :param pulumi.Input[int] max_replay_latency: Delayed log size threshold in MB.
        :param pulumi.Input[int] min_delay_eliminate_reserve: The minimum number of read-only replicas that must be retained in an RO group.
        :param pulumi.Input[str] name: RO group name.
        :param pulumi.Input[int] project_id: Project ID.
        :param pulumi.Input[int] replay_lag_eliminate: Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[int] replay_latency_eliminate: Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups_ids: ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        :param pulumi.Input[str] subnet_id: VPC subnet ID.
        :param pulumi.Input[str] vpc_id: VPC ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReadonlyGroupState.__new__(_ReadonlyGroupState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["master_db_instance_id"] = master_db_instance_id
        __props__.__dict__["max_replay_lag"] = max_replay_lag
        __props__.__dict__["max_replay_latency"] = max_replay_latency
        __props__.__dict__["min_delay_eliminate_reserve"] = min_delay_eliminate_reserve
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["replay_lag_eliminate"] = replay_lag_eliminate
        __props__.__dict__["replay_latency_eliminate"] = replay_latency_eliminate
        __props__.__dict__["security_groups_ids"] = security_groups_ids
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["vpc_id"] = vpc_id
        return ReadonlyGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the postgresql instance.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="masterDbInstanceId")
    def master_db_instance_id(self) -> pulumi.Output[str]:
        """
        Primary instance ID.
        """
        return pulumi.get(self, "master_db_instance_id")

    @property
    @pulumi.getter(name="maxReplayLag")
    def max_replay_lag(self) -> pulumi.Output[int]:
        """
        Delay threshold in ms.
        """
        return pulumi.get(self, "max_replay_lag")

    @property
    @pulumi.getter(name="maxReplayLatency")
    def max_replay_latency(self) -> pulumi.Output[int]:
        """
        Delayed log size threshold in MB.
        """
        return pulumi.get(self, "max_replay_latency")

    @property
    @pulumi.getter(name="minDelayEliminateReserve")
    def min_delay_eliminate_reserve(self) -> pulumi.Output[int]:
        """
        The minimum number of read-only replicas that must be retained in an RO group.
        """
        return pulumi.get(self, "min_delay_eliminate_reserve")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        RO group name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="replayLagEliminate")
    def replay_lag_eliminate(self) -> pulumi.Output[int]:
        """
        Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        """
        return pulumi.get(self, "replay_lag_eliminate")

    @property
    @pulumi.getter(name="replayLatencyEliminate")
    def replay_latency_eliminate(self) -> pulumi.Output[int]:
        """
        Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).
        """
        return pulumi.get(self, "replay_latency_eliminate")

    @property
    @pulumi.getter(name="securityGroupsIds")
    def security_groups_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
        """
        return pulumi.get(self, "security_groups_ids")

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

