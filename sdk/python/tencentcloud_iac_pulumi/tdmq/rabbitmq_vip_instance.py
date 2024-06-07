# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RabbitmqVipInstanceArgs', 'RabbitmqVipInstance']

@pulumi.input_type
class RabbitmqVipInstanceArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 zone_ids: pulumi.Input[Sequence[pulumi.Input[int]]],
                 auto_renew_flag: Optional[pulumi.Input[bool]] = None,
                 enable_create_default_ha_mirror_queue: Optional[pulumi.Input[bool]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 time_span: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RabbitmqVipInstance resource.
        :param pulumi.Input[str] cluster_name: cluster name.
        :param pulumi.Input[str] subnet_id: Private network SubnetId.
        :param pulumi.Input[str] vpc_id: Private network VpcId.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] zone_ids: availability zone.
        :param pulumi.Input[bool] auto_renew_flag: Automatic renewal, the default is true.
        :param pulumi.Input[bool] enable_create_default_ha_mirror_queue: Mirrored queue, the default is false.
        :param pulumi.Input[int] node_num: The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        :param pulumi.Input[str] node_spec: Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        :param pulumi.Input[int] storage_size: Single node storage specification, the default is 200G.
        :param pulumi.Input[int] time_span: Purchase duration, the default is 1 (month).
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "zone_ids", zone_ids)
        if auto_renew_flag is not None:
            pulumi.set(__self__, "auto_renew_flag", auto_renew_flag)
        if enable_create_default_ha_mirror_queue is not None:
            pulumi.set(__self__, "enable_create_default_ha_mirror_queue", enable_create_default_ha_mirror_queue)
        if node_num is not None:
            pulumi.set(__self__, "node_num", node_num)
        if node_spec is not None:
            pulumi.set(__self__, "node_spec", node_spec)
        if storage_size is not None:
            pulumi.set(__self__, "storage_size", storage_size)
        if time_span is not None:
            pulumi.set(__self__, "time_span", time_span)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Private network SubnetId.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        Private network VpcId.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        availability zone.
        """
        return pulumi.get(self, "zone_ids")

    @zone_ids.setter
    def zone_ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "zone_ids", value)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> Optional[pulumi.Input[bool]]:
        """
        Automatic renewal, the default is true.
        """
        return pulumi.get(self, "auto_renew_flag")

    @auto_renew_flag.setter
    def auto_renew_flag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew_flag", value)

    @property
    @pulumi.getter(name="enableCreateDefaultHaMirrorQueue")
    def enable_create_default_ha_mirror_queue(self) -> Optional[pulumi.Input[bool]]:
        """
        Mirrored queue, the default is false.
        """
        return pulumi.get(self, "enable_create_default_ha_mirror_queue")

    @enable_create_default_ha_mirror_queue.setter
    def enable_create_default_ha_mirror_queue(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_create_default_ha_mirror_queue", value)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        """
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        """
        return pulumi.get(self, "node_spec")

    @node_spec.setter
    def node_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_spec", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> Optional[pulumi.Input[int]]:
        """
        Single node storage specification, the default is 200G.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase duration, the default is 1 (month).
        """
        return pulumi.get(self, "time_span")

    @time_span.setter
    def time_span(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_span", value)


@pulumi.input_type
class _RabbitmqVipInstanceState:
    def __init__(__self__, *,
                 auto_renew_flag: Optional[pulumi.Input[bool]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 enable_create_default_ha_mirror_queue: Optional[pulumi.Input[bool]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        Input properties used for looking up and filtering RabbitmqVipInstance resources.
        :param pulumi.Input[bool] auto_renew_flag: Automatic renewal, the default is true.
        :param pulumi.Input[str] cluster_name: cluster name.
        :param pulumi.Input[bool] enable_create_default_ha_mirror_queue: Mirrored queue, the default is false.
        :param pulumi.Input[int] node_num: The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        :param pulumi.Input[str] node_spec: Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        :param pulumi.Input[int] storage_size: Single node storage specification, the default is 200G.
        :param pulumi.Input[str] subnet_id: Private network SubnetId.
        :param pulumi.Input[int] time_span: Purchase duration, the default is 1 (month).
        :param pulumi.Input[str] vpc_id: Private network VpcId.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] zone_ids: availability zone.
        """
        if auto_renew_flag is not None:
            pulumi.set(__self__, "auto_renew_flag", auto_renew_flag)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if enable_create_default_ha_mirror_queue is not None:
            pulumi.set(__self__, "enable_create_default_ha_mirror_queue", enable_create_default_ha_mirror_queue)
        if node_num is not None:
            pulumi.set(__self__, "node_num", node_num)
        if node_spec is not None:
            pulumi.set(__self__, "node_spec", node_spec)
        if storage_size is not None:
            pulumi.set(__self__, "storage_size", storage_size)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if time_span is not None:
            pulumi.set(__self__, "time_span", time_span)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if zone_ids is not None:
            pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> Optional[pulumi.Input[bool]]:
        """
        Automatic renewal, the default is true.
        """
        return pulumi.get(self, "auto_renew_flag")

    @auto_renew_flag.setter
    def auto_renew_flag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew_flag", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="enableCreateDefaultHaMirrorQueue")
    def enable_create_default_ha_mirror_queue(self) -> Optional[pulumi.Input[bool]]:
        """
        Mirrored queue, the default is false.
        """
        return pulumi.get(self, "enable_create_default_ha_mirror_queue")

    @enable_create_default_ha_mirror_queue.setter
    def enable_create_default_ha_mirror_queue(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_create_default_ha_mirror_queue", value)

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        """
        return pulumi.get(self, "node_num")

    @node_num.setter
    def node_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_num", value)

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        """
        return pulumi.get(self, "node_spec")

    @node_spec.setter
    def node_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_spec", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> Optional[pulumi.Input[int]]:
        """
        Single node storage specification, the default is 200G.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Private network SubnetId.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase duration, the default is 1 (month).
        """
        return pulumi.get(self, "time_span")

    @time_span.setter
    def time_span(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_span", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Private network VpcId.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        availability zone.
        """
        return pulumi.get(self, "zone_ids")

    @zone_ids.setter
    def zone_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "zone_ids", value)


class RabbitmqVipInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew_flag: Optional[pulumi.Input[bool]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 enable_create_default_ha_mirror_queue: Optional[pulumi.Input[bool]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a tdmq rabbitmq_vip_instance

        ## Import

        tdmq rabbitmq_vip_instance can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rabbitmqVipInstance:RabbitmqVipInstance example amqp-mok52gmn
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew_flag: Automatic renewal, the default is true.
        :param pulumi.Input[str] cluster_name: cluster name.
        :param pulumi.Input[bool] enable_create_default_ha_mirror_queue: Mirrored queue, the default is false.
        :param pulumi.Input[int] node_num: The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        :param pulumi.Input[str] node_spec: Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        :param pulumi.Input[int] storage_size: Single node storage specification, the default is 200G.
        :param pulumi.Input[str] subnet_id: Private network SubnetId.
        :param pulumi.Input[int] time_span: Purchase duration, the default is 1 (month).
        :param pulumi.Input[str] vpc_id: Private network VpcId.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] zone_ids: availability zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RabbitmqVipInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tdmq rabbitmq_vip_instance

        ## Import

        tdmq rabbitmq_vip_instance can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rabbitmqVipInstance:RabbitmqVipInstance example amqp-mok52gmn
        ```

        :param str resource_name: The name of the resource.
        :param RabbitmqVipInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RabbitmqVipInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew_flag: Optional[pulumi.Input[bool]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 enable_create_default_ha_mirror_queue: Optional[pulumi.Input[bool]] = None,
                 node_num: Optional[pulumi.Input[int]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RabbitmqVipInstanceArgs.__new__(RabbitmqVipInstanceArgs)

            __props__.__dict__["auto_renew_flag"] = auto_renew_flag
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["enable_create_default_ha_mirror_queue"] = enable_create_default_ha_mirror_queue
            __props__.__dict__["node_num"] = node_num
            __props__.__dict__["node_spec"] = node_spec
            __props__.__dict__["storage_size"] = storage_size
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["time_span"] = time_span
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if zone_ids is None and not opts.urn:
                raise TypeError("Missing required property 'zone_ids'")
            __props__.__dict__["zone_ids"] = zone_ids
        super(RabbitmqVipInstance, __self__).__init__(
            'tencentcloud:Tdmq/rabbitmqVipInstance:RabbitmqVipInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew_flag: Optional[pulumi.Input[bool]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            enable_create_default_ha_mirror_queue: Optional[pulumi.Input[bool]] = None,
            node_num: Optional[pulumi.Input[int]] = None,
            node_spec: Optional[pulumi.Input[str]] = None,
            storage_size: Optional[pulumi.Input[int]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            time_span: Optional[pulumi.Input[int]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None) -> 'RabbitmqVipInstance':
        """
        Get an existing RabbitmqVipInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew_flag: Automatic renewal, the default is true.
        :param pulumi.Input[str] cluster_name: cluster name.
        :param pulumi.Input[bool] enable_create_default_ha_mirror_queue: Mirrored queue, the default is false.
        :param pulumi.Input[int] node_num: The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        :param pulumi.Input[str] node_spec: Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        :param pulumi.Input[int] storage_size: Single node storage specification, the default is 200G.
        :param pulumi.Input[str] subnet_id: Private network SubnetId.
        :param pulumi.Input[int] time_span: Purchase duration, the default is 1 (month).
        :param pulumi.Input[str] vpc_id: Private network VpcId.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] zone_ids: availability zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RabbitmqVipInstanceState.__new__(_RabbitmqVipInstanceState)

        __props__.__dict__["auto_renew_flag"] = auto_renew_flag
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["enable_create_default_ha_mirror_queue"] = enable_create_default_ha_mirror_queue
        __props__.__dict__["node_num"] = node_num
        __props__.__dict__["node_spec"] = node_spec
        __props__.__dict__["storage_size"] = storage_size
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["time_span"] = time_span
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["zone_ids"] = zone_ids
        return RabbitmqVipInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenewFlag")
    def auto_renew_flag(self) -> pulumi.Output[Optional[bool]]:
        """
        Automatic renewal, the default is true.
        """
        return pulumi.get(self, "auto_renew_flag")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="enableCreateDefaultHaMirrorQueue")
    def enable_create_default_ha_mirror_queue(self) -> pulumi.Output[Optional[bool]]:
        """
        Mirrored queue, the default is false.
        """
        return pulumi.get(self, "enable_create_default_ha_mirror_queue")

    @property
    @pulumi.getter(name="nodeNum")
    def node_num(self) -> pulumi.Output[Optional[int]]:
        """
        The number of nodes, a minimum of 3 nodes for a multi-availability zone. If not passed, the default single availability zone is 1, and the multi-availability zone is 3.
        """
        return pulumi.get(self, "node_num")

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> pulumi.Output[Optional[str]]:
        """
        Node specifications, basic type rabbit-vip-basic-1, standard type rabbit-vip-basic-2, high-level type 1 rabbit-vip-basic-3, high-level type 2 rabbit-vip-basic-4. If not passed, the default is the basic type.
        """
        return pulumi.get(self, "node_spec")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> pulumi.Output[Optional[int]]:
        """
        Single node storage specification, the default is 200G.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Private network SubnetId.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> pulumi.Output[Optional[int]]:
        """
        Purchase duration, the default is 1 (month).
        """
        return pulumi.get(self, "time_span")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        Private network VpcId.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> pulumi.Output[Sequence[int]]:
        """
        availability zone.
        """
        return pulumi.get(self, "zone_ids")

