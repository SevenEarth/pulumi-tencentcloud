# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LifecycleHookArgs', 'LifecycleHook']

@pulumi.input_type
class LifecycleHookArgs:
    def __init__(__self__, *,
                 lifecycle_hook_name: pulumi.Input[str],
                 lifecycle_transition: pulumi.Input[str],
                 scaling_group_id: pulumi.Input[str],
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_queue_name: Optional[pulumi.Input[str]] = None,
                 notification_target_type: Optional[pulumi.Input[str]] = None,
                 notification_topic_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LifecycleHook resource.
        :param pulumi.Input[str] lifecycle_hook_name: The name of the lifecycle hook.
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        :param pulumi.Input[str] default_result: Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time AS sends a message to the notification target.
        :param pulumi.Input[str] notification_queue_name: For CMQ_QUEUE type, a name of queue must be set.
        :param pulumi.Input[str] notification_target_type: Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        :param pulumi.Input[str] notification_topic_name: For CMQ_TOPIC type, a name of topic must be set.
        """
        pulumi.set(__self__, "lifecycle_hook_name", lifecycle_hook_name)
        pulumi.set(__self__, "lifecycle_transition", lifecycle_transition)
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        if default_result is not None:
            pulumi.set(__self__, "default_result", default_result)
        if heartbeat_timeout is not None:
            pulumi.set(__self__, "heartbeat_timeout", heartbeat_timeout)
        if notification_metadata is not None:
            pulumi.set(__self__, "notification_metadata", notification_metadata)
        if notification_queue_name is not None:
            pulumi.set(__self__, "notification_queue_name", notification_queue_name)
        if notification_target_type is not None:
            pulumi.set(__self__, "notification_target_type", notification_target_type)
        if notification_topic_name is not None:
            pulumi.set(__self__, "notification_topic_name", notification_topic_name)

    @property
    @pulumi.getter(name="lifecycleHookName")
    def lifecycle_hook_name(self) -> pulumi.Input[str]:
        """
        The name of the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_hook_name")

    @lifecycle_hook_name.setter
    def lifecycle_hook_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "lifecycle_hook_name", value)

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> pulumi.Input[str]:
        """
        The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        """
        return pulumi.get(self, "lifecycle_transition")

    @lifecycle_transition.setter
    def lifecycle_transition(self, value: pulumi.Input[str]):
        pulumi.set(self, "lifecycle_transition", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Input[str]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scaling_group_id", value)

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        """
        return pulumi.get(self, "default_result")

    @default_result.setter
    def default_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_result", value)

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        """
        return pulumi.get(self, "heartbeat_timeout")

    @heartbeat_timeout.setter
    def heartbeat_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat_timeout", value)

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        Contains additional information that you want to include any time AS sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @notification_metadata.setter
    def notification_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_metadata", value)

    @property
    @pulumi.getter(name="notificationQueueName")
    def notification_queue_name(self) -> Optional[pulumi.Input[str]]:
        """
        For CMQ_QUEUE type, a name of queue must be set.
        """
        return pulumi.get(self, "notification_queue_name")

    @notification_queue_name.setter
    def notification_queue_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_queue_name", value)

    @property
    @pulumi.getter(name="notificationTargetType")
    def notification_target_type(self) -> Optional[pulumi.Input[str]]:
        """
        Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        """
        return pulumi.get(self, "notification_target_type")

    @notification_target_type.setter
    def notification_target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_target_type", value)

    @property
    @pulumi.getter(name="notificationTopicName")
    def notification_topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        For CMQ_TOPIC type, a name of topic must be set.
        """
        return pulumi.get(self, "notification_topic_name")

    @notification_topic_name.setter
    def notification_topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_topic_name", value)


@pulumi.input_type
class _LifecycleHookState:
    def __init__(__self__, *,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_queue_name: Optional[pulumi.Input[str]] = None,
                 notification_target_type: Optional[pulumi.Input[str]] = None,
                 notification_topic_name: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LifecycleHook resources.
        :param pulumi.Input[str] default_result: Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        :param pulumi.Input[str] lifecycle_hook_name: The name of the lifecycle hook.
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time AS sends a message to the notification target.
        :param pulumi.Input[str] notification_queue_name: For CMQ_QUEUE type, a name of queue must be set.
        :param pulumi.Input[str] notification_target_type: Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        :param pulumi.Input[str] notification_topic_name: For CMQ_TOPIC type, a name of topic must be set.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        if default_result is not None:
            pulumi.set(__self__, "default_result", default_result)
        if heartbeat_timeout is not None:
            pulumi.set(__self__, "heartbeat_timeout", heartbeat_timeout)
        if lifecycle_hook_name is not None:
            pulumi.set(__self__, "lifecycle_hook_name", lifecycle_hook_name)
        if lifecycle_transition is not None:
            pulumi.set(__self__, "lifecycle_transition", lifecycle_transition)
        if notification_metadata is not None:
            pulumi.set(__self__, "notification_metadata", notification_metadata)
        if notification_queue_name is not None:
            pulumi.set(__self__, "notification_queue_name", notification_queue_name)
        if notification_target_type is not None:
            pulumi.set(__self__, "notification_target_type", notification_target_type)
        if notification_topic_name is not None:
            pulumi.set(__self__, "notification_topic_name", notification_topic_name)
        if scaling_group_id is not None:
            pulumi.set(__self__, "scaling_group_id", scaling_group_id)

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> Optional[pulumi.Input[str]]:
        """
        Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        """
        return pulumi.get(self, "default_result")

    @default_result.setter
    def default_result(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_result", value)

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        """
        return pulumi.get(self, "heartbeat_timeout")

    @heartbeat_timeout.setter
    def heartbeat_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "heartbeat_timeout", value)

    @property
    @pulumi.getter(name="lifecycleHookName")
    def lifecycle_hook_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_hook_name")

    @lifecycle_hook_name.setter
    def lifecycle_hook_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_hook_name", value)

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> Optional[pulumi.Input[str]]:
        """
        The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        """
        return pulumi.get(self, "lifecycle_transition")

    @lifecycle_transition.setter
    def lifecycle_transition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_transition", value)

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[pulumi.Input[str]]:
        """
        Contains additional information that you want to include any time AS sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @notification_metadata.setter
    def notification_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_metadata", value)

    @property
    @pulumi.getter(name="notificationQueueName")
    def notification_queue_name(self) -> Optional[pulumi.Input[str]]:
        """
        For CMQ_QUEUE type, a name of queue must be set.
        """
        return pulumi.get(self, "notification_queue_name")

    @notification_queue_name.setter
    def notification_queue_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_queue_name", value)

    @property
    @pulumi.getter(name="notificationTargetType")
    def notification_target_type(self) -> Optional[pulumi.Input[str]]:
        """
        Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        """
        return pulumi.get(self, "notification_target_type")

    @notification_target_type.setter
    def notification_target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_target_type", value)

    @property
    @pulumi.getter(name="notificationTopicName")
    def notification_topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        For CMQ_TOPIC type, a name of topic must be set.
        """
        return pulumi.get(self, "notification_topic_name")

    @notification_topic_name.setter
    def notification_topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notification_topic_name", value)

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

    @scaling_group_id.setter
    def scaling_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scaling_group_id", value)


class LifecycleHook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_queue_name: Optional[pulumi.Input[str]] = None,
                 notification_target_type: Optional[pulumi.Input[str]] = None,
                 notification_topic_name: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource for an AS (Auto scaling) lifecycle hook.

        ## Example Usage
        ### Create a basic LifecycleHook

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="as")
        image = tencentcloud.Images.get_instance(image_types=["PUBLIC_IMAGE"],
            os_name="TencentOS Server 3.2 (Final)")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            availability_zone=zones.zones[0].name)
        example_scaling_config = tencentcloud.as_.ScalingConfig("exampleScalingConfig",
            configuration_name="tf-example",
            image_id=image.images[0].image_id,
            instance_types=[
                "SA1.SMALL1",
                "SA2.SMALL1",
                "SA2.SMALL2",
                "SA2.SMALL4",
            ],
            instance_name_settings=tencentcloud.as..ScalingConfigInstanceNameSettingsArgs(
                instance_name="test-ins-name",
            ))
        example_scaling_group = tencentcloud.as_.ScalingGroup("exampleScalingGroup",
            scaling_group_name="tf-example",
            configuration_id=example_scaling_config.id,
            max_size=1,
            min_size=0,
            vpc_id=vpc.id,
            subnet_ids=[subnet.id])
        example_lifecycle_hook = tencentcloud.as_.LifecycleHook("exampleLifecycleHook",
            scaling_group_id=example_scaling_group.id,
            lifecycle_hook_name="tf-as-lifecycle-hook",
            lifecycle_transition="INSTANCE_LAUNCHING",
            default_result="CONTINUE",
            heartbeat_timeout=500,
            notification_metadata="tf test")
        ```

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.as_.LifecycleHook("example",
            scaling_group_id=tencentcloud_as_scaling_group["example"]["id"],
            lifecycle_hook_name="tf-as-lifecycle-hook",
            lifecycle_transition="INSTANCE_LAUNCHING",
            default_result="CONTINUE",
            heartbeat_timeout=500,
            notification_metadata="tf test",
            notification_target_type="CMQ_QUEUE",
            notification_queue_name="lifcyclehook")
        ```

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.as_.LifecycleHook("example",
            scaling_group_id=tencentcloud_as_scaling_group["example"]["id"],
            lifecycle_hook_name="tf-as-lifecycle-hook",
            lifecycle_transition="INSTANCE_LAUNCHING",
            default_result="CONTINUE",
            heartbeat_timeout=500,
            notification_metadata="tf test",
            notification_target_type="CMQ_TOPIC",
            notification_topic_name="lifcyclehook")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_result: Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        :param pulumi.Input[str] lifecycle_hook_name: The name of the lifecycle hook.
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time AS sends a message to the notification target.
        :param pulumi.Input[str] notification_queue_name: For CMQ_QUEUE type, a name of queue must be set.
        :param pulumi.Input[str] notification_target_type: Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        :param pulumi.Input[str] notification_topic_name: For CMQ_TOPIC type, a name of topic must be set.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LifecycleHookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for an AS (Auto scaling) lifecycle hook.

        ## Example Usage
        ### Create a basic LifecycleHook

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="as")
        image = tencentcloud.Images.get_instance(image_types=["PUBLIC_IMAGE"],
            os_name="TencentOS Server 3.2 (Final)")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            availability_zone=zones.zones[0].name)
        example_scaling_config = tencentcloud.as_.ScalingConfig("exampleScalingConfig",
            configuration_name="tf-example",
            image_id=image.images[0].image_id,
            instance_types=[
                "SA1.SMALL1",
                "SA2.SMALL1",
                "SA2.SMALL2",
                "SA2.SMALL4",
            ],
            instance_name_settings=tencentcloud.as..ScalingConfigInstanceNameSettingsArgs(
                instance_name="test-ins-name",
            ))
        example_scaling_group = tencentcloud.as_.ScalingGroup("exampleScalingGroup",
            scaling_group_name="tf-example",
            configuration_id=example_scaling_config.id,
            max_size=1,
            min_size=0,
            vpc_id=vpc.id,
            subnet_ids=[subnet.id])
        example_lifecycle_hook = tencentcloud.as_.LifecycleHook("exampleLifecycleHook",
            scaling_group_id=example_scaling_group.id,
            lifecycle_hook_name="tf-as-lifecycle-hook",
            lifecycle_transition="INSTANCE_LAUNCHING",
            default_result="CONTINUE",
            heartbeat_timeout=500,
            notification_metadata="tf test")
        ```

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.as_.LifecycleHook("example",
            scaling_group_id=tencentcloud_as_scaling_group["example"]["id"],
            lifecycle_hook_name="tf-as-lifecycle-hook",
            lifecycle_transition="INSTANCE_LAUNCHING",
            default_result="CONTINUE",
            heartbeat_timeout=500,
            notification_metadata="tf test",
            notification_target_type="CMQ_QUEUE",
            notification_queue_name="lifcyclehook")
        ```

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.as_.LifecycleHook("example",
            scaling_group_id=tencentcloud_as_scaling_group["example"]["id"],
            lifecycle_hook_name="tf-as-lifecycle-hook",
            lifecycle_transition="INSTANCE_LAUNCHING",
            default_result="CONTINUE",
            heartbeat_timeout=500,
            notification_metadata="tf test",
            notification_target_type="CMQ_TOPIC",
            notification_topic_name="lifcyclehook")
        ```

        :param str resource_name: The name of the resource.
        :param LifecycleHookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LifecycleHookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[int]] = None,
                 lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_queue_name: Optional[pulumi.Input[str]] = None,
                 notification_target_type: Optional[pulumi.Input[str]] = None,
                 notification_topic_name: Optional[pulumi.Input[str]] = None,
                 scaling_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = LifecycleHookArgs.__new__(LifecycleHookArgs)

            __props__.__dict__["default_result"] = default_result
            __props__.__dict__["heartbeat_timeout"] = heartbeat_timeout
            if lifecycle_hook_name is None and not opts.urn:
                raise TypeError("Missing required property 'lifecycle_hook_name'")
            __props__.__dict__["lifecycle_hook_name"] = lifecycle_hook_name
            if lifecycle_transition is None and not opts.urn:
                raise TypeError("Missing required property 'lifecycle_transition'")
            __props__.__dict__["lifecycle_transition"] = lifecycle_transition
            __props__.__dict__["notification_metadata"] = notification_metadata
            __props__.__dict__["notification_queue_name"] = notification_queue_name
            __props__.__dict__["notification_target_type"] = notification_target_type
            __props__.__dict__["notification_topic_name"] = notification_topic_name
            if scaling_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'scaling_group_id'")
            __props__.__dict__["scaling_group_id"] = scaling_group_id
        super(LifecycleHook, __self__).__init__(
            'tencentcloud:As/lifecycleHook:LifecycleHook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_result: Optional[pulumi.Input[str]] = None,
            heartbeat_timeout: Optional[pulumi.Input[int]] = None,
            lifecycle_hook_name: Optional[pulumi.Input[str]] = None,
            lifecycle_transition: Optional[pulumi.Input[str]] = None,
            notification_metadata: Optional[pulumi.Input[str]] = None,
            notification_queue_name: Optional[pulumi.Input[str]] = None,
            notification_target_type: Optional[pulumi.Input[str]] = None,
            notification_topic_name: Optional[pulumi.Input[str]] = None,
            scaling_group_id: Optional[pulumi.Input[str]] = None) -> 'LifecycleHook':
        """
        Get an existing LifecycleHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_result: Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        :param pulumi.Input[int] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        :param pulumi.Input[str] lifecycle_hook_name: The name of the lifecycle hook.
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time AS sends a message to the notification target.
        :param pulumi.Input[str] notification_queue_name: For CMQ_QUEUE type, a name of queue must be set.
        :param pulumi.Input[str] notification_target_type: Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        :param pulumi.Input[str] notification_topic_name: For CMQ_TOPIC type, a name of topic must be set.
        :param pulumi.Input[str] scaling_group_id: ID of a scaling group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LifecycleHookState.__new__(_LifecycleHookState)

        __props__.__dict__["default_result"] = default_result
        __props__.__dict__["heartbeat_timeout"] = heartbeat_timeout
        __props__.__dict__["lifecycle_hook_name"] = lifecycle_hook_name
        __props__.__dict__["lifecycle_transition"] = lifecycle_transition
        __props__.__dict__["notification_metadata"] = notification_metadata
        __props__.__dict__["notification_queue_name"] = notification_queue_name
        __props__.__dict__["notification_target_type"] = notification_target_type
        __props__.__dict__["notification_topic_name"] = notification_topic_name
        __props__.__dict__["scaling_group_id"] = scaling_group_id
        return LifecycleHook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> pulumi.Output[Optional[str]]:
        """
        Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        """
        return pulumi.get(self, "default_result")

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        """
        return pulumi.get(self, "heartbeat_timeout")

    @property
    @pulumi.getter(name="lifecycleHookName")
    def lifecycle_hook_name(self) -> pulumi.Output[str]:
        """
        The name of the lifecycle hook.
        """
        return pulumi.get(self, "lifecycle_hook_name")

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> pulumi.Output[str]:
        """
        The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        """
        return pulumi.get(self, "lifecycle_transition")

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> pulumi.Output[Optional[str]]:
        """
        Contains additional information that you want to include any time AS sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @property
    @pulumi.getter(name="notificationQueueName")
    def notification_queue_name(self) -> pulumi.Output[Optional[str]]:
        """
        For CMQ_QUEUE type, a name of queue must be set.
        """
        return pulumi.get(self, "notification_queue_name")

    @property
    @pulumi.getter(name="notificationTargetType")
    def notification_target_type(self) -> pulumi.Output[Optional[str]]:
        """
        Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        """
        return pulumi.get(self, "notification_target_type")

    @property
    @pulumi.getter(name="notificationTopicName")
    def notification_topic_name(self) -> pulumi.Output[Optional[str]]:
        """
        For CMQ_TOPIC type, a name of topic must be set.
        """
        return pulumi.get(self, "notification_topic_name")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> pulumi.Output[str]:
        """
        ID of a scaling group.
        """
        return pulumi.get(self, "scaling_group_id")

