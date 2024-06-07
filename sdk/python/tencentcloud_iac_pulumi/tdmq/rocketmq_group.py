# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RocketmqGroupArgs', 'RocketmqGroup']

@pulumi.input_type
class RocketmqGroupArgs:
    def __init__(__self__, *,
                 broadcast_enable: pulumi.Input[bool],
                 cluster_id: pulumi.Input[str],
                 group_name: pulumi.Input[str],
                 namespace: pulumi.Input[str],
                 read_enable: pulumi.Input[bool],
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RocketmqGroup resource.
        :param pulumi.Input[bool] broadcast_enable: Whether to enable broadcast consumption.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] group_name: Group name (8-64 characters).
        :param pulumi.Input[str] namespace: Namespace. Currently, only one namespace is supported.
        :param pulumi.Input[bool] read_enable: Whether to enable consumption.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        """
        pulumi.set(__self__, "broadcast_enable", broadcast_enable)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "read_enable", read_enable)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="broadcastEnable")
    def broadcast_enable(self) -> pulumi.Input[bool]:
        """
        Whether to enable broadcast consumption.
        """
        return pulumi.get(self, "broadcast_enable")

    @broadcast_enable.setter
    def broadcast_enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "broadcast_enable", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        Group name (8-64 characters).
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        Namespace. Currently, only one namespace is supported.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="readEnable")
    def read_enable(self) -> pulumi.Input[bool]:
        """
        Whether to enable consumption.
        """
        return pulumi.get(self, "read_enable")

    @read_enable.setter
    def read_enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "read_enable", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _RocketmqGroupState:
    def __init__(__self__, *,
                 broadcast_enable: Optional[pulumi.Input[bool]] = None,
                 client_protocol: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 consumer_num: Optional[pulumi.Input[int]] = None,
                 consumer_type: Optional[pulumi.Input[str]] = None,
                 consumption_mode: Optional[pulumi.Input[int]] = None,
                 create_time: Optional[pulumi.Input[int]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 read_enable: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 retry_partition_num: Optional[pulumi.Input[int]] = None,
                 total_accumulative: Optional[pulumi.Input[int]] = None,
                 tps: Optional[pulumi.Input[int]] = None,
                 update_time: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RocketmqGroup resources.
        :param pulumi.Input[bool] broadcast_enable: Whether to enable broadcast consumption.
        :param pulumi.Input[str] client_protocol: Client protocol.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] consumer_num: The number of online consumers.
        :param pulumi.Input[str] consumer_type: Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
        :param pulumi.Input[int] consumption_mode: `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
        :param pulumi.Input[int] create_time: Creation time in milliseconds.
        :param pulumi.Input[str] group_name: Group name (8-64 characters).
        :param pulumi.Input[str] namespace: Namespace. Currently, only one namespace is supported.
        :param pulumi.Input[bool] read_enable: Whether to enable consumption.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        :param pulumi.Input[int] retry_partition_num: The number of partitions in a retry topic.
        :param pulumi.Input[int] total_accumulative: The total number of heaped messages.
        :param pulumi.Input[int] tps: Consumption TPS.
        :param pulumi.Input[int] update_time: Modification time in milliseconds.
        """
        if broadcast_enable is not None:
            pulumi.set(__self__, "broadcast_enable", broadcast_enable)
        if client_protocol is not None:
            pulumi.set(__self__, "client_protocol", client_protocol)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if consumer_num is not None:
            pulumi.set(__self__, "consumer_num", consumer_num)
        if consumer_type is not None:
            pulumi.set(__self__, "consumer_type", consumer_type)
        if consumption_mode is not None:
            pulumi.set(__self__, "consumption_mode", consumption_mode)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if read_enable is not None:
            pulumi.set(__self__, "read_enable", read_enable)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if retry_partition_num is not None:
            pulumi.set(__self__, "retry_partition_num", retry_partition_num)
        if total_accumulative is not None:
            pulumi.set(__self__, "total_accumulative", total_accumulative)
        if tps is not None:
            pulumi.set(__self__, "tps", tps)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="broadcastEnable")
    def broadcast_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable broadcast consumption.
        """
        return pulumi.get(self, "broadcast_enable")

    @broadcast_enable.setter
    def broadcast_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "broadcast_enable", value)

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Client protocol.
        """
        return pulumi.get(self, "client_protocol")

    @client_protocol.setter
    def client_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_protocol", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="consumerNum")
    def consumer_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of online consumers.
        """
        return pulumi.get(self, "consumer_num")

    @consumer_num.setter
    def consumer_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "consumer_num", value)

    @property
    @pulumi.getter(name="consumerType")
    def consumer_type(self) -> Optional[pulumi.Input[str]]:
        """
        Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
        """
        return pulumi.get(self, "consumer_type")

    @consumer_type.setter
    def consumer_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_type", value)

    @property
    @pulumi.getter(name="consumptionMode")
    def consumption_mode(self) -> Optional[pulumi.Input[int]]:
        """
        `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
        """
        return pulumi.get(self, "consumption_mode")

    @consumption_mode.setter
    def consumption_mode(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "consumption_mode", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[int]]:
        """
        Creation time in milliseconds.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Group name (8-64 characters).
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace. Currently, only one namespace is supported.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="readEnable")
    def read_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable consumption.
        """
        return pulumi.get(self, "read_enable")

    @read_enable.setter
    def read_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_enable", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="retryPartitionNum")
    def retry_partition_num(self) -> Optional[pulumi.Input[int]]:
        """
        The number of partitions in a retry topic.
        """
        return pulumi.get(self, "retry_partition_num")

    @retry_partition_num.setter
    def retry_partition_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_partition_num", value)

    @property
    @pulumi.getter(name="totalAccumulative")
    def total_accumulative(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of heaped messages.
        """
        return pulumi.get(self, "total_accumulative")

    @total_accumulative.setter
    def total_accumulative(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "total_accumulative", value)

    @property
    @pulumi.getter
    def tps(self) -> Optional[pulumi.Input[int]]:
        """
        Consumption TPS.
        """
        return pulumi.get(self, "tps")

    @tps.setter
    def tps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tps", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[int]]:
        """
        Modification time in milliseconds.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "update_time", value)


class RocketmqGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broadcast_enable: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 read_enable: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tdmqRocketmq group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_rocketmq_cluster = tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster",
            cluster_name="tf_example",
            remark="remark.")
        example_rocketmq_namespace = tencentcloud.tdmq.RocketmqNamespace("exampleRocketmqNamespace",
            cluster_id=example_rocketmq_cluster.cluster_id,
            namespace_name="tf_example",
            remark="remark.")
        example_rocketmq_group = tencentcloud.tdmq.RocketmqGroup("exampleRocketmqGroup",
            group_name="tf_example",
            cluster_id=example_rocketmq_cluster.cluster_id,
            namespace=example_rocketmq_namespace.namespace_name,
            read_enable=True,
            broadcast_enable=True,
            remark="remark.")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tdmqRocketmq group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup group group_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] broadcast_enable: Whether to enable broadcast consumption.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] group_name: Group name (8-64 characters).
        :param pulumi.Input[str] namespace: Namespace. Currently, only one namespace is supported.
        :param pulumi.Input[bool] read_enable: Whether to enable consumption.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tdmqRocketmq group

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_rocketmq_cluster = tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster",
            cluster_name="tf_example",
            remark="remark.")
        example_rocketmq_namespace = tencentcloud.tdmq.RocketmqNamespace("exampleRocketmqNamespace",
            cluster_id=example_rocketmq_cluster.cluster_id,
            namespace_name="tf_example",
            remark="remark.")
        example_rocketmq_group = tencentcloud.tdmq.RocketmqGroup("exampleRocketmqGroup",
            group_name="tf_example",
            cluster_id=example_rocketmq_cluster.cluster_id,
            namespace=example_rocketmq_namespace.namespace_name,
            read_enable=True,
            broadcast_enable=True,
            remark="remark.")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tdmqRocketmq group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup group group_id
        ```

        :param str resource_name: The name of the resource.
        :param RocketmqGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broadcast_enable: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 read_enable: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RocketmqGroupArgs.__new__(RocketmqGroupArgs)

            if broadcast_enable is None and not opts.urn:
                raise TypeError("Missing required property 'broadcast_enable'")
            __props__.__dict__["broadcast_enable"] = broadcast_enable
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            if read_enable is None and not opts.urn:
                raise TypeError("Missing required property 'read_enable'")
            __props__.__dict__["read_enable"] = read_enable
            __props__.__dict__["remark"] = remark
            __props__.__dict__["client_protocol"] = None
            __props__.__dict__["consumer_num"] = None
            __props__.__dict__["consumer_type"] = None
            __props__.__dict__["consumption_mode"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["retry_partition_num"] = None
            __props__.__dict__["total_accumulative"] = None
            __props__.__dict__["tps"] = None
            __props__.__dict__["update_time"] = None
        super(RocketmqGroup, __self__).__init__(
            'tencentcloud:Tdmq/rocketmqGroup:RocketmqGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            broadcast_enable: Optional[pulumi.Input[bool]] = None,
            client_protocol: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            consumer_num: Optional[pulumi.Input[int]] = None,
            consumer_type: Optional[pulumi.Input[str]] = None,
            consumption_mode: Optional[pulumi.Input[int]] = None,
            create_time: Optional[pulumi.Input[int]] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            read_enable: Optional[pulumi.Input[bool]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            retry_partition_num: Optional[pulumi.Input[int]] = None,
            total_accumulative: Optional[pulumi.Input[int]] = None,
            tps: Optional[pulumi.Input[int]] = None,
            update_time: Optional[pulumi.Input[int]] = None) -> 'RocketmqGroup':
        """
        Get an existing RocketmqGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] broadcast_enable: Whether to enable broadcast consumption.
        :param pulumi.Input[str] client_protocol: Client protocol.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] consumer_num: The number of online consumers.
        :param pulumi.Input[str] consumer_type: Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
        :param pulumi.Input[int] consumption_mode: `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
        :param pulumi.Input[int] create_time: Creation time in milliseconds.
        :param pulumi.Input[str] group_name: Group name (8-64 characters).
        :param pulumi.Input[str] namespace: Namespace. Currently, only one namespace is supported.
        :param pulumi.Input[bool] read_enable: Whether to enable consumption.
        :param pulumi.Input[str] remark: Remarks (up to 128 characters).
        :param pulumi.Input[int] retry_partition_num: The number of partitions in a retry topic.
        :param pulumi.Input[int] total_accumulative: The total number of heaped messages.
        :param pulumi.Input[int] tps: Consumption TPS.
        :param pulumi.Input[int] update_time: Modification time in milliseconds.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqGroupState.__new__(_RocketmqGroupState)

        __props__.__dict__["broadcast_enable"] = broadcast_enable
        __props__.__dict__["client_protocol"] = client_protocol
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["consumer_num"] = consumer_num
        __props__.__dict__["consumer_type"] = consumer_type
        __props__.__dict__["consumption_mode"] = consumption_mode
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["read_enable"] = read_enable
        __props__.__dict__["remark"] = remark
        __props__.__dict__["retry_partition_num"] = retry_partition_num
        __props__.__dict__["total_accumulative"] = total_accumulative
        __props__.__dict__["tps"] = tps
        __props__.__dict__["update_time"] = update_time
        return RocketmqGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="broadcastEnable")
    def broadcast_enable(self) -> pulumi.Output[bool]:
        """
        Whether to enable broadcast consumption.
        """
        return pulumi.get(self, "broadcast_enable")

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> pulumi.Output[str]:
        """
        Client protocol.
        """
        return pulumi.get(self, "client_protocol")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="consumerNum")
    def consumer_num(self) -> pulumi.Output[int]:
        """
        The number of online consumers.
        """
        return pulumi.get(self, "consumer_num")

    @property
    @pulumi.getter(name="consumerType")
    def consumer_type(self) -> pulumi.Output[str]:
        """
        Consumer type. Enumerated values: ACTIVELY or PASSIVELY.
        """
        return pulumi.get(self, "consumer_type")

    @property
    @pulumi.getter(name="consumptionMode")
    def consumption_mode(self) -> pulumi.Output[int]:
        """
        `0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.
        """
        return pulumi.get(self, "consumption_mode")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[int]:
        """
        Creation time in milliseconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        Group name (8-64 characters).
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        Namespace. Currently, only one namespace is supported.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="readEnable")
    def read_enable(self) -> pulumi.Output[bool]:
        """
        Whether to enable consumption.
        """
        return pulumi.get(self, "read_enable")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="retryPartitionNum")
    def retry_partition_num(self) -> pulumi.Output[int]:
        """
        The number of partitions in a retry topic.
        """
        return pulumi.get(self, "retry_partition_num")

    @property
    @pulumi.getter(name="totalAccumulative")
    def total_accumulative(self) -> pulumi.Output[int]:
        """
        The total number of heaped messages.
        """
        return pulumi.get(self, "total_accumulative")

    @property
    @pulumi.getter
    def tps(self) -> pulumi.Output[int]:
        """
        Consumption TPS.
        """
        return pulumi.get(self, "tps")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[int]:
        """
        Modification time in milliseconds.
        """
        return pulumi.get(self, "update_time")

