# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TriggerConfigArgs', 'TriggerConfig']

@pulumi.input_type
class TriggerConfigArgs:
    def __init__(__self__, *,
                 enable: pulumi.Input[str],
                 function_name: pulumi.Input[str],
                 trigger_name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 trigger_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TriggerConfig resource.
        :param pulumi.Input[str] enable: Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        :param pulumi.Input[str] function_name: Function name.
        :param pulumi.Input[str] trigger_name: Trigger name.
        :param pulumi.Input[str] type: Trigger Type.
        :param pulumi.Input[str] namespace: Function namespace.
        :param pulumi.Input[str] qualifier: Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        :param pulumi.Input[str] trigger_desc: To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        """
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "function_name", function_name)
        pulumi.set(__self__, "trigger_name", trigger_name)
        pulumi.set(__self__, "type", type)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)
        if trigger_desc is not None:
            pulumi.set(__self__, "trigger_desc", trigger_desc)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[str]:
        """
        Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[str]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Input[str]:
        """
        Function name.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter(name="triggerName")
    def trigger_name(self) -> pulumi.Input[str]:
        """
        Trigger name.
        """
        return pulumi.get(self, "trigger_name")

    @trigger_name.setter
    def trigger_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "trigger_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Trigger Type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Function namespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)

    @property
    @pulumi.getter(name="triggerDesc")
    def trigger_desc(self) -> Optional[pulumi.Input[str]]:
        """
        To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        """
        return pulumi.get(self, "trigger_desc")

    @trigger_desc.setter
    def trigger_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_desc", value)


@pulumi.input_type
class _TriggerConfigState:
    def __init__(__self__, *,
                 enable: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 trigger_desc: Optional[pulumi.Input[str]] = None,
                 trigger_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TriggerConfig resources.
        :param pulumi.Input[str] enable: Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        :param pulumi.Input[str] function_name: Function name.
        :param pulumi.Input[str] namespace: Function namespace.
        :param pulumi.Input[str] qualifier: Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        :param pulumi.Input[str] trigger_desc: To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        :param pulumi.Input[str] trigger_name: Trigger name.
        :param pulumi.Input[str] type: Trigger Type.
        """
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if function_name is not None:
            pulumi.set(__self__, "function_name", function_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)
        if trigger_desc is not None:
            pulumi.set(__self__, "trigger_desc", trigger_desc)
        if trigger_name is not None:
            pulumi.set(__self__, "trigger_name", trigger_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[str]]:
        """
        Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[pulumi.Input[str]]:
        """
        Function name.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Function namespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)

    @property
    @pulumi.getter(name="triggerDesc")
    def trigger_desc(self) -> Optional[pulumi.Input[str]]:
        """
        To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        """
        return pulumi.get(self, "trigger_desc")

    @trigger_desc.setter
    def trigger_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_desc", value)

    @property
    @pulumi.getter(name="triggerName")
    def trigger_name(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger name.
        """
        return pulumi.get(self, "trigger_name")

    @trigger_name.setter
    def trigger_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger Type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class TriggerConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 trigger_desc: Optional[pulumi.Input[str]] = None,
                 trigger_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a scf trigger_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        trigger_config = tencentcloud.scf.TriggerConfig("triggerConfig",
            enable="OPEN",
            function_name="keep-1676351130",
            namespace="default",
            qualifier="$DEFAULT",
            trigger_name="SCF-timer-1685540160",
            type="timer")
        ```

        ## Import

        scf trigger_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Scf/triggerConfig:TriggerConfig trigger_config functionName#namespace#triggerName
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enable: Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        :param pulumi.Input[str] function_name: Function name.
        :param pulumi.Input[str] namespace: Function namespace.
        :param pulumi.Input[str] qualifier: Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        :param pulumi.Input[str] trigger_desc: To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        :param pulumi.Input[str] trigger_name: Trigger name.
        :param pulumi.Input[str] type: Trigger Type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a scf trigger_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        trigger_config = tencentcloud.scf.TriggerConfig("triggerConfig",
            enable="OPEN",
            function_name="keep-1676351130",
            namespace="default",
            qualifier="$DEFAULT",
            trigger_name="SCF-timer-1685540160",
            type="timer")
        ```

        ## Import

        scf trigger_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Scf/triggerConfig:TriggerConfig trigger_config functionName#namespace#triggerName
        ```

        :param str resource_name: The name of the resource.
        :param TriggerConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 trigger_desc: Optional[pulumi.Input[str]] = None,
                 trigger_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = TriggerConfigArgs.__new__(TriggerConfigArgs)

            if enable is None and not opts.urn:
                raise TypeError("Missing required property 'enable'")
            __props__.__dict__["enable"] = enable
            if function_name is None and not opts.urn:
                raise TypeError("Missing required property 'function_name'")
            __props__.__dict__["function_name"] = function_name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["qualifier"] = qualifier
            __props__.__dict__["trigger_desc"] = trigger_desc
            if trigger_name is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_name'")
            __props__.__dict__["trigger_name"] = trigger_name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(TriggerConfig, __self__).__init__(
            'tencentcloud:Scf/triggerConfig:TriggerConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enable: Optional[pulumi.Input[str]] = None,
            function_name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            qualifier: Optional[pulumi.Input[str]] = None,
            trigger_desc: Optional[pulumi.Input[str]] = None,
            trigger_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'TriggerConfig':
        """
        Get an existing TriggerConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enable: Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        :param pulumi.Input[str] function_name: Function name.
        :param pulumi.Input[str] namespace: Function namespace.
        :param pulumi.Input[str] qualifier: Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        :param pulumi.Input[str] trigger_desc: To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        :param pulumi.Input[str] trigger_name: Trigger name.
        :param pulumi.Input[str] type: Trigger Type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TriggerConfigState.__new__(_TriggerConfigState)

        __props__.__dict__["enable"] = enable
        __props__.__dict__["function_name"] = function_name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["qualifier"] = qualifier
        __props__.__dict__["trigger_desc"] = trigger_desc
        __props__.__dict__["trigger_name"] = trigger_name
        __props__.__dict__["type"] = type
        return TriggerConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[str]:
        """
        Initial status of the trigger. Values: `OPEN` (enabled); `CLOSE` disabled).
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        """
        Function name.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Function namespace.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def qualifier(self) -> pulumi.Output[Optional[str]]:
        """
        Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        """
        return pulumi.get(self, "qualifier")

    @property
    @pulumi.getter(name="triggerDesc")
    def trigger_desc(self) -> pulumi.Output[str]:
        """
        To update a COS trigger, this field is required. It stores the data {event:cos:ObjectCreated:*} in the JSON format. The data content of this field is in the same format as that of SetTrigger. This field is optional if a scheduled trigger or CMQ trigger is to be deleted.
        """
        return pulumi.get(self, "trigger_desc")

    @property
    @pulumi.getter(name="triggerName")
    def trigger_name(self) -> pulumi.Output[str]:
        """
        Trigger name.
        """
        return pulumi.get(self, "trigger_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Trigger Type.
        """
        return pulumi.get(self, "type")

