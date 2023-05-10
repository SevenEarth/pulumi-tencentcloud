# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RocketmqTopicArgs', 'RocketmqTopic']

@pulumi.input_type
class RocketmqTopicArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 topic_name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 partition_num: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RocketmqTopic resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] namespace_name: Topic namespace. Currently, you can create topics only in one single namespace.
        :param pulumi.Input[str] topic_name: Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] type: Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        :param pulumi.Input[int] partition_num: Number of partitions.
        :param pulumi.Input[str] remark: Topic remarks (up to 128 characters).
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "topic_name", topic_name)
        pulumi.set(__self__, "type", type)
        if partition_num is not None:
            pulumi.set(__self__, "partition_num", partition_num)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

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
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        Topic namespace. Currently, you can create topics only in one single namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="partitionNum")
    def partition_num(self) -> Optional[pulumi.Input[int]]:
        """
        Number of partitions.
        """
        return pulumi.get(self, "partition_num")

    @partition_num.setter
    def partition_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "partition_num", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Topic remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _RocketmqTopicState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[int]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 partition_num: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RocketmqTopic resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] create_time: Creation time in milliseconds.
        :param pulumi.Input[str] namespace_name: Topic namespace. Currently, you can create topics only in one single namespace.
        :param pulumi.Input[int] partition_num: Number of partitions.
        :param pulumi.Input[str] remark: Topic remarks (up to 128 characters).
        :param pulumi.Input[str] topic_name: Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] type: Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        :param pulumi.Input[int] update_time: Update time in milliseconds.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if partition_num is not None:
            pulumi.set(__self__, "partition_num", partition_num)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

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
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        Topic namespace. Currently, you can create topics only in one single namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="partitionNum")
    def partition_num(self) -> Optional[pulumi.Input[int]]:
        """
        Number of partitions.
        """
        return pulumi.get(self, "partition_num")

    @partition_num.setter
    def partition_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "partition_num", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Topic remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[int]]:
        """
        Update time in milliseconds.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "update_time", value)


class RocketmqTopic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 partition_num: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tdmqRocketmq topic

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cluster = tencentcloud.tdmq.RocketmqCluster("cluster",
            cluster_name="test_rocketmq",
            remark="test recket mq")
        namespace = tencentcloud.tdmq.RocketmqNamespace("namespace",
            cluster_id=cluster.cluster_id,
            namespace_name="test_namespace",
            ttl=65000,
            retention_time=65000,
            remark="test namespace")
        topic = tencentcloud.tdmq.RocketmqTopic("topic",
            topic_name="test_rocketmq_topic",
            namespace_name=namespace.namespace_name,
            type="Normal",
            cluster_id=cluster.cluster_id,
            remark="test rocketmq topic")
        ```

        ## Import

        tdmqRocketmq topic can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tdmq/rocketmqTopic:RocketmqTopic topic topic_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] namespace_name: Topic namespace. Currently, you can create topics only in one single namespace.
        :param pulumi.Input[int] partition_num: Number of partitions.
        :param pulumi.Input[str] remark: Topic remarks (up to 128 characters).
        :param pulumi.Input[str] topic_name: Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] type: Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqTopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tdmqRocketmq topic

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cluster = tencentcloud.tdmq.RocketmqCluster("cluster",
            cluster_name="test_rocketmq",
            remark="test recket mq")
        namespace = tencentcloud.tdmq.RocketmqNamespace("namespace",
            cluster_id=cluster.cluster_id,
            namespace_name="test_namespace",
            ttl=65000,
            retention_time=65000,
            remark="test namespace")
        topic = tencentcloud.tdmq.RocketmqTopic("topic",
            topic_name="test_rocketmq_topic",
            namespace_name=namespace.namespace_name,
            type="Normal",
            cluster_id=cluster.cluster_id,
            remark="test rocketmq topic")
        ```

        ## Import

        tdmqRocketmq topic can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tdmq/rocketmqTopic:RocketmqTopic topic topic_id
        ```

        :param str resource_name: The name of the resource.
        :param RocketmqTopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqTopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 partition_num: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = RocketmqTopicArgs.__new__(RocketmqTopicArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            __props__.__dict__["partition_num"] = partition_num
            __props__.__dict__["remark"] = remark
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(RocketmqTopic, __self__).__init__(
            'tencentcloud:Tdmq/rocketmqTopic:RocketmqTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[int]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            partition_num: Optional[pulumi.Input[int]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[int]] = None) -> 'RocketmqTopic':
        """
        Get an existing RocketmqTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[int] create_time: Creation time in milliseconds.
        :param pulumi.Input[str] namespace_name: Topic namespace. Currently, you can create topics only in one single namespace.
        :param pulumi.Input[int] partition_num: Number of partitions.
        :param pulumi.Input[str] remark: Topic remarks (up to 128 characters).
        :param pulumi.Input[str] topic_name: Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] type: Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        :param pulumi.Input[int] update_time: Update time in milliseconds.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqTopicState.__new__(_RocketmqTopicState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["partition_num"] = partition_num
        __props__.__dict__["remark"] = remark
        __props__.__dict__["topic_name"] = topic_name
        __props__.__dict__["type"] = type
        __props__.__dict__["update_time"] = update_time
        return RocketmqTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[int]:
        """
        Creation time in milliseconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        Topic namespace. Currently, you can create topics only in one single namespace.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="partitionNum")
    def partition_num(self) -> pulumi.Output[Optional[int]]:
        """
        Number of partitions.
        """
        return pulumi.get(self, "partition_num")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Topic remarks (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        Topic name, which can contain 3-64 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "topic_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[int]:
        """
        Update time in milliseconds.
        """
        return pulumi.get(self, "update_time")

