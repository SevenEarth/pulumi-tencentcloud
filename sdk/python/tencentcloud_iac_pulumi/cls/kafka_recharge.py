# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['KafkaRechargeArgs', 'KafkaRecharge']

@pulumi.input_type
class KafkaRechargeArgs:
    def __init__(__self__, *,
                 kafka_type: pulumi.Input[int],
                 offset: pulumi.Input[int],
                 topic_id: pulumi.Input[str],
                 user_kafka_topics: pulumi.Input[str],
                 consumer_group_name: Optional[pulumi.Input[str]] = None,
                 is_encryption_addr: Optional[pulumi.Input[bool]] = None,
                 kafka_instance: Optional[pulumi.Input[str]] = None,
                 log_recharge_rule: Optional[pulumi.Input['KafkaRechargeLogRechargeRuleArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input['KafkaRechargeProtocolArgs']] = None,
                 server_addr: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a KafkaRecharge resource.
        :param pulumi.Input[int] kafka_type: kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        :param pulumi.Input[int] offset: The translation is: -2: Earliest (default) -1: Latest.
        :param pulumi.Input[str] topic_id: recharge for cls TopicId.
        :param pulumi.Input[str] user_kafka_topics: user need recharge kafka topic list.
        :param pulumi.Input[str] consumer_group_name: user consumer group name.
        :param pulumi.Input[bool] is_encryption_addr: ServerAddr is encryption addr.
        :param pulumi.Input[str] kafka_instance: CKafka Instance id.
        :param pulumi.Input['KafkaRechargeLogRechargeRuleArgs'] log_recharge_rule: log recharge rule.
        :param pulumi.Input[str] name: kafka recharge name.
        :param pulumi.Input['KafkaRechargeProtocolArgs'] protocol: encryption protocol.
        :param pulumi.Input[str] server_addr: Server addr.
        """
        pulumi.set(__self__, "kafka_type", kafka_type)
        pulumi.set(__self__, "offset", offset)
        pulumi.set(__self__, "topic_id", topic_id)
        pulumi.set(__self__, "user_kafka_topics", user_kafka_topics)
        if consumer_group_name is not None:
            pulumi.set(__self__, "consumer_group_name", consumer_group_name)
        if is_encryption_addr is not None:
            pulumi.set(__self__, "is_encryption_addr", is_encryption_addr)
        if kafka_instance is not None:
            pulumi.set(__self__, "kafka_instance", kafka_instance)
        if log_recharge_rule is not None:
            pulumi.set(__self__, "log_recharge_rule", log_recharge_rule)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if server_addr is not None:
            pulumi.set(__self__, "server_addr", server_addr)

    @property
    @pulumi.getter(name="kafkaType")
    def kafka_type(self) -> pulumi.Input[int]:
        """
        kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        """
        return pulumi.get(self, "kafka_type")

    @kafka_type.setter
    def kafka_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "kafka_type", value)

    @property
    @pulumi.getter
    def offset(self) -> pulumi.Input[int]:
        """
        The translation is: -2: Earliest (default) -1: Latest.
        """
        return pulumi.get(self, "offset")

    @offset.setter
    def offset(self, value: pulumi.Input[int]):
        pulumi.set(self, "offset", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Input[str]:
        """
        recharge for cls TopicId.
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_id", value)

    @property
    @pulumi.getter(name="userKafkaTopics")
    def user_kafka_topics(self) -> pulumi.Input[str]:
        """
        user need recharge kafka topic list.
        """
        return pulumi.get(self, "user_kafka_topics")

    @user_kafka_topics.setter
    def user_kafka_topics(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_kafka_topics", value)

    @property
    @pulumi.getter(name="consumerGroupName")
    def consumer_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        user consumer group name.
        """
        return pulumi.get(self, "consumer_group_name")

    @consumer_group_name.setter
    def consumer_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_group_name", value)

    @property
    @pulumi.getter(name="isEncryptionAddr")
    def is_encryption_addr(self) -> Optional[pulumi.Input[bool]]:
        """
        ServerAddr is encryption addr.
        """
        return pulumi.get(self, "is_encryption_addr")

    @is_encryption_addr.setter
    def is_encryption_addr(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_encryption_addr", value)

    @property
    @pulumi.getter(name="kafkaInstance")
    def kafka_instance(self) -> Optional[pulumi.Input[str]]:
        """
        CKafka Instance id.
        """
        return pulumi.get(self, "kafka_instance")

    @kafka_instance.setter
    def kafka_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kafka_instance", value)

    @property
    @pulumi.getter(name="logRechargeRule")
    def log_recharge_rule(self) -> Optional[pulumi.Input['KafkaRechargeLogRechargeRuleArgs']]:
        """
        log recharge rule.
        """
        return pulumi.get(self, "log_recharge_rule")

    @log_recharge_rule.setter
    def log_recharge_rule(self, value: Optional[pulumi.Input['KafkaRechargeLogRechargeRuleArgs']]):
        pulumi.set(self, "log_recharge_rule", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        kafka recharge name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input['KafkaRechargeProtocolArgs']]:
        """
        encryption protocol.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input['KafkaRechargeProtocolArgs']]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="serverAddr")
    def server_addr(self) -> Optional[pulumi.Input[str]]:
        """
        Server addr.
        """
        return pulumi.get(self, "server_addr")

    @server_addr.setter
    def server_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_addr", value)


@pulumi.input_type
class _KafkaRechargeState:
    def __init__(__self__, *,
                 consumer_group_name: Optional[pulumi.Input[str]] = None,
                 is_encryption_addr: Optional[pulumi.Input[bool]] = None,
                 kafka_instance: Optional[pulumi.Input[str]] = None,
                 kafka_type: Optional[pulumi.Input[int]] = None,
                 log_recharge_rule: Optional[pulumi.Input['KafkaRechargeLogRechargeRuleArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 offset: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input['KafkaRechargeProtocolArgs']] = None,
                 server_addr: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 user_kafka_topics: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KafkaRecharge resources.
        :param pulumi.Input[str] consumer_group_name: user consumer group name.
        :param pulumi.Input[bool] is_encryption_addr: ServerAddr is encryption addr.
        :param pulumi.Input[str] kafka_instance: CKafka Instance id.
        :param pulumi.Input[int] kafka_type: kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        :param pulumi.Input['KafkaRechargeLogRechargeRuleArgs'] log_recharge_rule: log recharge rule.
        :param pulumi.Input[str] name: kafka recharge name.
        :param pulumi.Input[int] offset: The translation is: -2: Earliest (default) -1: Latest.
        :param pulumi.Input['KafkaRechargeProtocolArgs'] protocol: encryption protocol.
        :param pulumi.Input[str] server_addr: Server addr.
        :param pulumi.Input[str] topic_id: recharge for cls TopicId.
        :param pulumi.Input[str] user_kafka_topics: user need recharge kafka topic list.
        """
        if consumer_group_name is not None:
            pulumi.set(__self__, "consumer_group_name", consumer_group_name)
        if is_encryption_addr is not None:
            pulumi.set(__self__, "is_encryption_addr", is_encryption_addr)
        if kafka_instance is not None:
            pulumi.set(__self__, "kafka_instance", kafka_instance)
        if kafka_type is not None:
            pulumi.set(__self__, "kafka_type", kafka_type)
        if log_recharge_rule is not None:
            pulumi.set(__self__, "log_recharge_rule", log_recharge_rule)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if offset is not None:
            pulumi.set(__self__, "offset", offset)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if server_addr is not None:
            pulumi.set(__self__, "server_addr", server_addr)
        if topic_id is not None:
            pulumi.set(__self__, "topic_id", topic_id)
        if user_kafka_topics is not None:
            pulumi.set(__self__, "user_kafka_topics", user_kafka_topics)

    @property
    @pulumi.getter(name="consumerGroupName")
    def consumer_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        user consumer group name.
        """
        return pulumi.get(self, "consumer_group_name")

    @consumer_group_name.setter
    def consumer_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_group_name", value)

    @property
    @pulumi.getter(name="isEncryptionAddr")
    def is_encryption_addr(self) -> Optional[pulumi.Input[bool]]:
        """
        ServerAddr is encryption addr.
        """
        return pulumi.get(self, "is_encryption_addr")

    @is_encryption_addr.setter
    def is_encryption_addr(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_encryption_addr", value)

    @property
    @pulumi.getter(name="kafkaInstance")
    def kafka_instance(self) -> Optional[pulumi.Input[str]]:
        """
        CKafka Instance id.
        """
        return pulumi.get(self, "kafka_instance")

    @kafka_instance.setter
    def kafka_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kafka_instance", value)

    @property
    @pulumi.getter(name="kafkaType")
    def kafka_type(self) -> Optional[pulumi.Input[int]]:
        """
        kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        """
        return pulumi.get(self, "kafka_type")

    @kafka_type.setter
    def kafka_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kafka_type", value)

    @property
    @pulumi.getter(name="logRechargeRule")
    def log_recharge_rule(self) -> Optional[pulumi.Input['KafkaRechargeLogRechargeRuleArgs']]:
        """
        log recharge rule.
        """
        return pulumi.get(self, "log_recharge_rule")

    @log_recharge_rule.setter
    def log_recharge_rule(self, value: Optional[pulumi.Input['KafkaRechargeLogRechargeRuleArgs']]):
        pulumi.set(self, "log_recharge_rule", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        kafka recharge name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def offset(self) -> Optional[pulumi.Input[int]]:
        """
        The translation is: -2: Earliest (default) -1: Latest.
        """
        return pulumi.get(self, "offset")

    @offset.setter
    def offset(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "offset", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input['KafkaRechargeProtocolArgs']]:
        """
        encryption protocol.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input['KafkaRechargeProtocolArgs']]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="serverAddr")
    def server_addr(self) -> Optional[pulumi.Input[str]]:
        """
        Server addr.
        """
        return pulumi.get(self, "server_addr")

    @server_addr.setter
    def server_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_addr", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> Optional[pulumi.Input[str]]:
        """
        recharge for cls TopicId.
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_id", value)

    @property
    @pulumi.getter(name="userKafkaTopics")
    def user_kafka_topics(self) -> Optional[pulumi.Input[str]]:
        """
        user need recharge kafka topic list.
        """
        return pulumi.get(self, "user_kafka_topics")

    @user_kafka_topics.setter
    def user_kafka_topics(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_kafka_topics", value)


class KafkaRecharge(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_group_name: Optional[pulumi.Input[str]] = None,
                 is_encryption_addr: Optional[pulumi.Input[bool]] = None,
                 kafka_instance: Optional[pulumi.Input[str]] = None,
                 kafka_type: Optional[pulumi.Input[int]] = None,
                 log_recharge_rule: Optional[pulumi.Input[pulumi.InputType['KafkaRechargeLogRechargeRuleArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 offset: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[pulumi.InputType['KafkaRechargeProtocolArgs']]] = None,
                 server_addr: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 user_kafka_topics: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a cls kafka_recharge

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        logset = tencentcloud.cls.Logset("logset",
            logset_name="tf-example-logset",
            tags={
                "createdBy": "terraform",
            })
        topic = tencentcloud.cls.Topic("topic",
            topic_name="tf-example-topic",
            logset_id=logset.id,
            auto_split=False,
            max_split_partitions=20,
            partition_count=1,
            period=10,
            storage_type="hot",
            tags={
                "test": "test",
            })
        kafka_recharge = tencentcloud.cls.KafkaRecharge("kafkaRecharge",
            topic_id=topic.id,
            kafka_type=0,
            offset=-2,
            is_encryption_addr=True,
            user_kafka_topics="recharge",
            kafka_instance="ckafka-qzoeaqx8",
            log_recharge_rule=tencentcloud.cls.KafkaRechargeLogRechargeRuleArgs(
                recharge_type="json_log",
                encoding_format=0,
                default_time_switch=True,
            ))
        ```

        ## Import

        cls kafka_recharge can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cls/kafkaRecharge:KafkaRecharge kafka_recharge kafka_recharge_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_group_name: user consumer group name.
        :param pulumi.Input[bool] is_encryption_addr: ServerAddr is encryption addr.
        :param pulumi.Input[str] kafka_instance: CKafka Instance id.
        :param pulumi.Input[int] kafka_type: kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        :param pulumi.Input[pulumi.InputType['KafkaRechargeLogRechargeRuleArgs']] log_recharge_rule: log recharge rule.
        :param pulumi.Input[str] name: kafka recharge name.
        :param pulumi.Input[int] offset: The translation is: -2: Earliest (default) -1: Latest.
        :param pulumi.Input[pulumi.InputType['KafkaRechargeProtocolArgs']] protocol: encryption protocol.
        :param pulumi.Input[str] server_addr: Server addr.
        :param pulumi.Input[str] topic_id: recharge for cls TopicId.
        :param pulumi.Input[str] user_kafka_topics: user need recharge kafka topic list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KafkaRechargeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cls kafka_recharge

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        logset = tencentcloud.cls.Logset("logset",
            logset_name="tf-example-logset",
            tags={
                "createdBy": "terraform",
            })
        topic = tencentcloud.cls.Topic("topic",
            topic_name="tf-example-topic",
            logset_id=logset.id,
            auto_split=False,
            max_split_partitions=20,
            partition_count=1,
            period=10,
            storage_type="hot",
            tags={
                "test": "test",
            })
        kafka_recharge = tencentcloud.cls.KafkaRecharge("kafkaRecharge",
            topic_id=topic.id,
            kafka_type=0,
            offset=-2,
            is_encryption_addr=True,
            user_kafka_topics="recharge",
            kafka_instance="ckafka-qzoeaqx8",
            log_recharge_rule=tencentcloud.cls.KafkaRechargeLogRechargeRuleArgs(
                recharge_type="json_log",
                encoding_format=0,
                default_time_switch=True,
            ))
        ```

        ## Import

        cls kafka_recharge can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cls/kafkaRecharge:KafkaRecharge kafka_recharge kafka_recharge_id
        ```

        :param str resource_name: The name of the resource.
        :param KafkaRechargeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KafkaRechargeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_group_name: Optional[pulumi.Input[str]] = None,
                 is_encryption_addr: Optional[pulumi.Input[bool]] = None,
                 kafka_instance: Optional[pulumi.Input[str]] = None,
                 kafka_type: Optional[pulumi.Input[int]] = None,
                 log_recharge_rule: Optional[pulumi.Input[pulumi.InputType['KafkaRechargeLogRechargeRuleArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 offset: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[pulumi.InputType['KafkaRechargeProtocolArgs']]] = None,
                 server_addr: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 user_kafka_topics: Optional[pulumi.Input[str]] = None,
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
            __props__ = KafkaRechargeArgs.__new__(KafkaRechargeArgs)

            __props__.__dict__["consumer_group_name"] = consumer_group_name
            __props__.__dict__["is_encryption_addr"] = is_encryption_addr
            __props__.__dict__["kafka_instance"] = kafka_instance
            if kafka_type is None and not opts.urn:
                raise TypeError("Missing required property 'kafka_type'")
            __props__.__dict__["kafka_type"] = kafka_type
            __props__.__dict__["log_recharge_rule"] = log_recharge_rule
            __props__.__dict__["name"] = name
            if offset is None and not opts.urn:
                raise TypeError("Missing required property 'offset'")
            __props__.__dict__["offset"] = offset
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["server_addr"] = server_addr
            if topic_id is None and not opts.urn:
                raise TypeError("Missing required property 'topic_id'")
            __props__.__dict__["topic_id"] = topic_id
            if user_kafka_topics is None and not opts.urn:
                raise TypeError("Missing required property 'user_kafka_topics'")
            __props__.__dict__["user_kafka_topics"] = user_kafka_topics
        super(KafkaRecharge, __self__).__init__(
            'tencentcloud:Cls/kafkaRecharge:KafkaRecharge',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consumer_group_name: Optional[pulumi.Input[str]] = None,
            is_encryption_addr: Optional[pulumi.Input[bool]] = None,
            kafka_instance: Optional[pulumi.Input[str]] = None,
            kafka_type: Optional[pulumi.Input[int]] = None,
            log_recharge_rule: Optional[pulumi.Input[pulumi.InputType['KafkaRechargeLogRechargeRuleArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            offset: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[pulumi.InputType['KafkaRechargeProtocolArgs']]] = None,
            server_addr: Optional[pulumi.Input[str]] = None,
            topic_id: Optional[pulumi.Input[str]] = None,
            user_kafka_topics: Optional[pulumi.Input[str]] = None) -> 'KafkaRecharge':
        """
        Get an existing KafkaRecharge resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_group_name: user consumer group name.
        :param pulumi.Input[bool] is_encryption_addr: ServerAddr is encryption addr.
        :param pulumi.Input[str] kafka_instance: CKafka Instance id.
        :param pulumi.Input[int] kafka_type: kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        :param pulumi.Input[pulumi.InputType['KafkaRechargeLogRechargeRuleArgs']] log_recharge_rule: log recharge rule.
        :param pulumi.Input[str] name: kafka recharge name.
        :param pulumi.Input[int] offset: The translation is: -2: Earliest (default) -1: Latest.
        :param pulumi.Input[pulumi.InputType['KafkaRechargeProtocolArgs']] protocol: encryption protocol.
        :param pulumi.Input[str] server_addr: Server addr.
        :param pulumi.Input[str] topic_id: recharge for cls TopicId.
        :param pulumi.Input[str] user_kafka_topics: user need recharge kafka topic list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KafkaRechargeState.__new__(_KafkaRechargeState)

        __props__.__dict__["consumer_group_name"] = consumer_group_name
        __props__.__dict__["is_encryption_addr"] = is_encryption_addr
        __props__.__dict__["kafka_instance"] = kafka_instance
        __props__.__dict__["kafka_type"] = kafka_type
        __props__.__dict__["log_recharge_rule"] = log_recharge_rule
        __props__.__dict__["name"] = name
        __props__.__dict__["offset"] = offset
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["server_addr"] = server_addr
        __props__.__dict__["topic_id"] = topic_id
        __props__.__dict__["user_kafka_topics"] = user_kafka_topics
        return KafkaRecharge(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consumerGroupName")
    def consumer_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        user consumer group name.
        """
        return pulumi.get(self, "consumer_group_name")

    @property
    @pulumi.getter(name="isEncryptionAddr")
    def is_encryption_addr(self) -> pulumi.Output[bool]:
        """
        ServerAddr is encryption addr.
        """
        return pulumi.get(self, "is_encryption_addr")

    @property
    @pulumi.getter(name="kafkaInstance")
    def kafka_instance(self) -> pulumi.Output[Optional[str]]:
        """
        CKafka Instance id.
        """
        return pulumi.get(self, "kafka_instance")

    @property
    @pulumi.getter(name="kafkaType")
    def kafka_type(self) -> pulumi.Output[int]:
        """
        kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        """
        return pulumi.get(self, "kafka_type")

    @property
    @pulumi.getter(name="logRechargeRule")
    def log_recharge_rule(self) -> pulumi.Output['outputs.KafkaRechargeLogRechargeRule']:
        """
        log recharge rule.
        """
        return pulumi.get(self, "log_recharge_rule")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        kafka recharge name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def offset(self) -> pulumi.Output[int]:
        """
        The translation is: -2: Earliest (default) -1: Latest.
        """
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output['outputs.KafkaRechargeProtocol']:
        """
        encryption protocol.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="serverAddr")
    def server_addr(self) -> pulumi.Output[Optional[str]]:
        """
        Server addr.
        """
        return pulumi.get(self, "server_addr")

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Output[str]:
        """
        recharge for cls TopicId.
        """
        return pulumi.get(self, "topic_id")

    @property
    @pulumi.getter(name="userKafkaTopics")
    def user_kafka_topics(self) -> pulumi.Output[str]:
        """
        user need recharge kafka topic list.
        """
        return pulumi.get(self, "user_kafka_topics")

