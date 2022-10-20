# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TargetGroupAttachmentArgs', 'TargetGroupAttachment']

@pulumi.input_type
class TargetGroupAttachmentArgs:
    def __init__(__self__, *,
                 clb_id: pulumi.Input[str],
                 listener_id: pulumi.Input[str],
                 rule_id: Optional[pulumi.Input[str]] = None,
                 target_group_id: Optional[pulumi.Input[str]] = None,
                 targrt_group_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TargetGroupAttachment resource.
        :param pulumi.Input[str] clb_id: ID of the CLB.
        :param pulumi.Input[str] listener_id: ID of the CLB listener.
        :param pulumi.Input[str] rule_id: ID of the CLB listener rule.
        :param pulumi.Input[str] target_group_id: ID of the CLB target group.
        :param pulumi.Input[str] targrt_group_id: It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        pulumi.set(__self__, "clb_id", clb_id)
        pulumi.set(__self__, "listener_id", listener_id)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if target_group_id is not None:
            pulumi.set(__self__, "target_group_id", target_group_id)
        if targrt_group_id is not None:
            warnings.warn("""It has been deprecated from version 1.47.1. Use `target_group_id` instead.""", DeprecationWarning)
            pulumi.log.warn("""targrt_group_id is deprecated: It has been deprecated from version 1.47.1. Use `target_group_id` instead.""")
        if targrt_group_id is not None:
            pulumi.set(__self__, "targrt_group_id", targrt_group_id)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> pulumi.Input[str]:
        """
        ID of the CLB.
        """
        return pulumi.get(self, "clb_id")

    @clb_id.setter
    def clb_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "clb_id", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Input[str]:
        """
        ID of the CLB listener.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CLB listener rule.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="targetGroupId")
    def target_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CLB target group.
        """
        return pulumi.get(self, "target_group_id")

    @target_group_id.setter
    def target_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_group_id", value)

    @property
    @pulumi.getter(name="targrtGroupId")
    def targrt_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        return pulumi.get(self, "targrt_group_id")

    @targrt_group_id.setter
    def targrt_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "targrt_group_id", value)


@pulumi.input_type
class _TargetGroupAttachmentState:
    def __init__(__self__, *,
                 clb_id: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 target_group_id: Optional[pulumi.Input[str]] = None,
                 targrt_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TargetGroupAttachment resources.
        :param pulumi.Input[str] clb_id: ID of the CLB.
        :param pulumi.Input[str] listener_id: ID of the CLB listener.
        :param pulumi.Input[str] rule_id: ID of the CLB listener rule.
        :param pulumi.Input[str] target_group_id: ID of the CLB target group.
        :param pulumi.Input[str] targrt_group_id: It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        if clb_id is not None:
            pulumi.set(__self__, "clb_id", clb_id)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if target_group_id is not None:
            pulumi.set(__self__, "target_group_id", target_group_id)
        if targrt_group_id is not None:
            warnings.warn("""It has been deprecated from version 1.47.1. Use `target_group_id` instead.""", DeprecationWarning)
            pulumi.log.warn("""targrt_group_id is deprecated: It has been deprecated from version 1.47.1. Use `target_group_id` instead.""")
        if targrt_group_id is not None:
            pulumi.set(__self__, "targrt_group_id", targrt_group_id)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CLB.
        """
        return pulumi.get(self, "clb_id")

    @clb_id.setter
    def clb_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clb_id", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CLB listener.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CLB listener rule.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="targetGroupId")
    def target_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CLB target group.
        """
        return pulumi.get(self, "target_group_id")

    @target_group_id.setter
    def target_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_group_id", value)

    @property
    @pulumi.getter(name="targrtGroupId")
    def targrt_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        return pulumi.get(self, "targrt_group_id")

    @targrt_group_id.setter
    def targrt_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "targrt_group_id", value)


class TargetGroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clb_id: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 target_group_id: Optional[pulumi.Input[str]] = None,
                 targrt_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a CLB target group attachment is bound to the load balancing listener or forwarding rule.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        clb_basic = tencentcloud.clb.Instance("clbBasic",
            network_type="OPEN",
            clb_name="tf-clb-rule-basic")
        listener_basic = tencentcloud.clb.Listener("listenerBasic",
            clb_id=clb_basic.id,
            port=1,
            protocol="HTTP",
            listener_name="listener_basic")
        rule_basic = tencentcloud.clb.ListenerRule("ruleBasic",
            clb_id=clb_basic.id,
            listener_id=listener_basic.listener_id,
            domain="abc.com",
            url="/",
            session_expire_time=30,
            scheduler="WRR",
            target_type="TARGETGROUP")
        test = tencentcloud.clb.TargetGroup("test", target_group_name="test-target-keep-1")
        group = tencentcloud.clb.TargetGroupAttachment("group",
            clb_id=clb_basic.id,
            listener_id=listener_basic.listener_id,
            rule_id=rule_basic.rule_id,
            target_group_id=test.id)
        ```

        ## Import

        CLB target group attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clb/targetGroupAttachment:TargetGroupAttachment group lbtg-odareyb2#lbl-bicjmx3i#lb-cv0iz74c#loc-ac6uk7b6
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] clb_id: ID of the CLB.
        :param pulumi.Input[str] listener_id: ID of the CLB listener.
        :param pulumi.Input[str] rule_id: ID of the CLB listener rule.
        :param pulumi.Input[str] target_group_id: ID of the CLB target group.
        :param pulumi.Input[str] targrt_group_id: It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TargetGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a CLB target group attachment is bound to the load balancing listener or forwarding rule.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        clb_basic = tencentcloud.clb.Instance("clbBasic",
            network_type="OPEN",
            clb_name="tf-clb-rule-basic")
        listener_basic = tencentcloud.clb.Listener("listenerBasic",
            clb_id=clb_basic.id,
            port=1,
            protocol="HTTP",
            listener_name="listener_basic")
        rule_basic = tencentcloud.clb.ListenerRule("ruleBasic",
            clb_id=clb_basic.id,
            listener_id=listener_basic.listener_id,
            domain="abc.com",
            url="/",
            session_expire_time=30,
            scheduler="WRR",
            target_type="TARGETGROUP")
        test = tencentcloud.clb.TargetGroup("test", target_group_name="test-target-keep-1")
        group = tencentcloud.clb.TargetGroupAttachment("group",
            clb_id=clb_basic.id,
            listener_id=listener_basic.listener_id,
            rule_id=rule_basic.rule_id,
            target_group_id=test.id)
        ```

        ## Import

        CLB target group attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clb/targetGroupAttachment:TargetGroupAttachment group lbtg-odareyb2#lbl-bicjmx3i#lb-cv0iz74c#loc-ac6uk7b6
        ```

        :param str resource_name: The name of the resource.
        :param TargetGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clb_id: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 target_group_id: Optional[pulumi.Input[str]] = None,
                 targrt_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TargetGroupAttachmentArgs.__new__(TargetGroupAttachmentArgs)

            if clb_id is None and not opts.urn:
                raise TypeError("Missing required property 'clb_id'")
            __props__.__dict__["clb_id"] = clb_id
            if listener_id is None and not opts.urn:
                raise TypeError("Missing required property 'listener_id'")
            __props__.__dict__["listener_id"] = listener_id
            __props__.__dict__["rule_id"] = rule_id
            __props__.__dict__["target_group_id"] = target_group_id
            if targrt_group_id is not None and not opts.urn:
                warnings.warn("""It has been deprecated from version 1.47.1. Use `target_group_id` instead.""", DeprecationWarning)
                pulumi.log.warn("""targrt_group_id is deprecated: It has been deprecated from version 1.47.1. Use `target_group_id` instead.""")
            __props__.__dict__["targrt_group_id"] = targrt_group_id
        super(TargetGroupAttachment, __self__).__init__(
            'tencentcloud:Clb/targetGroupAttachment:TargetGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            clb_id: Optional[pulumi.Input[str]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            rule_id: Optional[pulumi.Input[str]] = None,
            target_group_id: Optional[pulumi.Input[str]] = None,
            targrt_group_id: Optional[pulumi.Input[str]] = None) -> 'TargetGroupAttachment':
        """
        Get an existing TargetGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] clb_id: ID of the CLB.
        :param pulumi.Input[str] listener_id: ID of the CLB listener.
        :param pulumi.Input[str] rule_id: ID of the CLB listener rule.
        :param pulumi.Input[str] target_group_id: ID of the CLB target group.
        :param pulumi.Input[str] targrt_group_id: It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TargetGroupAttachmentState.__new__(_TargetGroupAttachmentState)

        __props__.__dict__["clb_id"] = clb_id
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["target_group_id"] = target_group_id
        __props__.__dict__["targrt_group_id"] = targrt_group_id
        return TargetGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> pulumi.Output[str]:
        """
        ID of the CLB.
        """
        return pulumi.get(self, "clb_id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        ID of the CLB listener.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the CLB listener rule.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="targetGroupId")
    def target_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the CLB target group.
        """
        return pulumi.get(self, "target_group_id")

    @property
    @pulumi.getter(name="targrtGroupId")
    def targrt_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        It has been deprecated from version 1.47.1. Use `target_group_id` instead. ID of the CLB target group.
        """
        return pulumi.get(self, "targrt_group_id")

