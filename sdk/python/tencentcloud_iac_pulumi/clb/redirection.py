# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RedirectionArgs', 'Redirection']

@pulumi.input_type
class RedirectionArgs:
    def __init__(__self__, *,
                 clb_id: pulumi.Input[str],
                 target_listener_id: pulumi.Input[str],
                 target_rule_id: pulumi.Input[str],
                 delete_all_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 is_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 source_listener_id: Optional[pulumi.Input[str]] = None,
                 source_rule_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Redirection resource.
        :param pulumi.Input[str] clb_id: ID of CLB instance.
        :param pulumi.Input[str] target_listener_id: ID of source listener.
        :param pulumi.Input[str] target_rule_id: Rule ID of target listener.
        :param pulumi.Input[bool] delete_all_auto_rewrite: Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        :param pulumi.Input[bool] is_auto_rewrite: Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        :param pulumi.Input[str] source_listener_id: ID of source listener.
        :param pulumi.Input[str] source_rule_id: Rule ID of source listener.
        """
        pulumi.set(__self__, "clb_id", clb_id)
        pulumi.set(__self__, "target_listener_id", target_listener_id)
        pulumi.set(__self__, "target_rule_id", target_rule_id)
        if delete_all_auto_rewrite is not None:
            pulumi.set(__self__, "delete_all_auto_rewrite", delete_all_auto_rewrite)
        if is_auto_rewrite is not None:
            pulumi.set(__self__, "is_auto_rewrite", is_auto_rewrite)
        if source_listener_id is not None:
            pulumi.set(__self__, "source_listener_id", source_listener_id)
        if source_rule_id is not None:
            pulumi.set(__self__, "source_rule_id", source_rule_id)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> pulumi.Input[str]:
        """
        ID of CLB instance.
        """
        return pulumi.get(self, "clb_id")

    @clb_id.setter
    def clb_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "clb_id", value)

    @property
    @pulumi.getter(name="targetListenerId")
    def target_listener_id(self) -> pulumi.Input[str]:
        """
        ID of source listener.
        """
        return pulumi.get(self, "target_listener_id")

    @target_listener_id.setter
    def target_listener_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_listener_id", value)

    @property
    @pulumi.getter(name="targetRuleId")
    def target_rule_id(self) -> pulumi.Input[str]:
        """
        Rule ID of target listener.
        """
        return pulumi.get(self, "target_rule_id")

    @target_rule_id.setter
    def target_rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_rule_id", value)

    @property
    @pulumi.getter(name="deleteAllAutoRewrite")
    def delete_all_auto_rewrite(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        """
        return pulumi.get(self, "delete_all_auto_rewrite")

    @delete_all_auto_rewrite.setter
    def delete_all_auto_rewrite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_all_auto_rewrite", value)

    @property
    @pulumi.getter(name="isAutoRewrite")
    def is_auto_rewrite(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        """
        return pulumi.get(self, "is_auto_rewrite")

    @is_auto_rewrite.setter
    def is_auto_rewrite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_rewrite", value)

    @property
    @pulumi.getter(name="sourceListenerId")
    def source_listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of source listener.
        """
        return pulumi.get(self, "source_listener_id")

    @source_listener_id.setter
    def source_listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_listener_id", value)

    @property
    @pulumi.getter(name="sourceRuleId")
    def source_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Rule ID of source listener.
        """
        return pulumi.get(self, "source_rule_id")

    @source_rule_id.setter
    def source_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_rule_id", value)


@pulumi.input_type
class _RedirectionState:
    def __init__(__self__, *,
                 clb_id: Optional[pulumi.Input[str]] = None,
                 delete_all_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 is_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 source_listener_id: Optional[pulumi.Input[str]] = None,
                 source_rule_id: Optional[pulumi.Input[str]] = None,
                 target_listener_id: Optional[pulumi.Input[str]] = None,
                 target_rule_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Redirection resources.
        :param pulumi.Input[str] clb_id: ID of CLB instance.
        :param pulumi.Input[bool] delete_all_auto_rewrite: Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        :param pulumi.Input[bool] is_auto_rewrite: Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        :param pulumi.Input[str] source_listener_id: ID of source listener.
        :param pulumi.Input[str] source_rule_id: Rule ID of source listener.
        :param pulumi.Input[str] target_listener_id: ID of source listener.
        :param pulumi.Input[str] target_rule_id: Rule ID of target listener.
        """
        if clb_id is not None:
            pulumi.set(__self__, "clb_id", clb_id)
        if delete_all_auto_rewrite is not None:
            pulumi.set(__self__, "delete_all_auto_rewrite", delete_all_auto_rewrite)
        if is_auto_rewrite is not None:
            pulumi.set(__self__, "is_auto_rewrite", is_auto_rewrite)
        if source_listener_id is not None:
            pulumi.set(__self__, "source_listener_id", source_listener_id)
        if source_rule_id is not None:
            pulumi.set(__self__, "source_rule_id", source_rule_id)
        if target_listener_id is not None:
            pulumi.set(__self__, "target_listener_id", target_listener_id)
        if target_rule_id is not None:
            pulumi.set(__self__, "target_rule_id", target_rule_id)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of CLB instance.
        """
        return pulumi.get(self, "clb_id")

    @clb_id.setter
    def clb_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clb_id", value)

    @property
    @pulumi.getter(name="deleteAllAutoRewrite")
    def delete_all_auto_rewrite(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        """
        return pulumi.get(self, "delete_all_auto_rewrite")

    @delete_all_auto_rewrite.setter
    def delete_all_auto_rewrite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_all_auto_rewrite", value)

    @property
    @pulumi.getter(name="isAutoRewrite")
    def is_auto_rewrite(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        """
        return pulumi.get(self, "is_auto_rewrite")

    @is_auto_rewrite.setter
    def is_auto_rewrite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_rewrite", value)

    @property
    @pulumi.getter(name="sourceListenerId")
    def source_listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of source listener.
        """
        return pulumi.get(self, "source_listener_id")

    @source_listener_id.setter
    def source_listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_listener_id", value)

    @property
    @pulumi.getter(name="sourceRuleId")
    def source_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Rule ID of source listener.
        """
        return pulumi.get(self, "source_rule_id")

    @source_rule_id.setter
    def source_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_rule_id", value)

    @property
    @pulumi.getter(name="targetListenerId")
    def target_listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of source listener.
        """
        return pulumi.get(self, "target_listener_id")

    @target_listener_id.setter
    def target_listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_listener_id", value)

    @property
    @pulumi.getter(name="targetRuleId")
    def target_rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Rule ID of target listener.
        """
        return pulumi.get(self, "target_rule_id")

    @target_rule_id.setter
    def target_rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_rule_id", value)


class Redirection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clb_id: Optional[pulumi.Input[str]] = None,
                 delete_all_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 is_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 source_listener_id: Optional[pulumi.Input[str]] = None,
                 source_rule_id: Optional[pulumi.Input[str]] = None,
                 target_listener_id: Optional[pulumi.Input[str]] = None,
                 target_rule_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a CLB redirection.

        ## Example Usage
        ### Manual Rewrite

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.clb.Redirection("foo",
            clb_id="lb-p7olt9e5",
            source_listener_id="lbl-jc1dx6ju",
            source_rule_id="loc-ft8fmngv",
            target_listener_id="lbl-asj1hzuo",
            target_rule_id="loc-4xxr2cy7")
        ```
        ### Auto Rewrite

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.clb.Redirection("foo",
            clb_id="lb-p7olt9e5",
            is_auto_rewrite=True,
            target_listener_id="lbl-asj1hzuo",
            target_rule_id="loc-4xxr2cy7")
        ```

        ## Import

        CLB redirection can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clb/redirection:Redirection foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] clb_id: ID of CLB instance.
        :param pulumi.Input[bool] delete_all_auto_rewrite: Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        :param pulumi.Input[bool] is_auto_rewrite: Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        :param pulumi.Input[str] source_listener_id: ID of source listener.
        :param pulumi.Input[str] source_rule_id: Rule ID of source listener.
        :param pulumi.Input[str] target_listener_id: ID of source listener.
        :param pulumi.Input[str] target_rule_id: Rule ID of target listener.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RedirectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a CLB redirection.

        ## Example Usage
        ### Manual Rewrite

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.clb.Redirection("foo",
            clb_id="lb-p7olt9e5",
            source_listener_id="lbl-jc1dx6ju",
            source_rule_id="loc-ft8fmngv",
            target_listener_id="lbl-asj1hzuo",
            target_rule_id="loc-4xxr2cy7")
        ```
        ### Auto Rewrite

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.clb.Redirection("foo",
            clb_id="lb-p7olt9e5",
            is_auto_rewrite=True,
            target_listener_id="lbl-asj1hzuo",
            target_rule_id="loc-4xxr2cy7")
        ```

        ## Import

        CLB redirection can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clb/redirection:Redirection foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5
        ```

        :param str resource_name: The name of the resource.
        :param RedirectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RedirectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clb_id: Optional[pulumi.Input[str]] = None,
                 delete_all_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 is_auto_rewrite: Optional[pulumi.Input[bool]] = None,
                 source_listener_id: Optional[pulumi.Input[str]] = None,
                 source_rule_id: Optional[pulumi.Input[str]] = None,
                 target_listener_id: Optional[pulumi.Input[str]] = None,
                 target_rule_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RedirectionArgs.__new__(RedirectionArgs)

            if clb_id is None and not opts.urn:
                raise TypeError("Missing required property 'clb_id'")
            __props__.__dict__["clb_id"] = clb_id
            __props__.__dict__["delete_all_auto_rewrite"] = delete_all_auto_rewrite
            __props__.__dict__["is_auto_rewrite"] = is_auto_rewrite
            __props__.__dict__["source_listener_id"] = source_listener_id
            __props__.__dict__["source_rule_id"] = source_rule_id
            if target_listener_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_listener_id'")
            __props__.__dict__["target_listener_id"] = target_listener_id
            if target_rule_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_rule_id'")
            __props__.__dict__["target_rule_id"] = target_rule_id
        super(Redirection, __self__).__init__(
            'tencentcloud:Clb/redirection:Redirection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            clb_id: Optional[pulumi.Input[str]] = None,
            delete_all_auto_rewrite: Optional[pulumi.Input[bool]] = None,
            is_auto_rewrite: Optional[pulumi.Input[bool]] = None,
            source_listener_id: Optional[pulumi.Input[str]] = None,
            source_rule_id: Optional[pulumi.Input[str]] = None,
            target_listener_id: Optional[pulumi.Input[str]] = None,
            target_rule_id: Optional[pulumi.Input[str]] = None) -> 'Redirection':
        """
        Get an existing Redirection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] clb_id: ID of CLB instance.
        :param pulumi.Input[bool] delete_all_auto_rewrite: Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        :param pulumi.Input[bool] is_auto_rewrite: Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        :param pulumi.Input[str] source_listener_id: ID of source listener.
        :param pulumi.Input[str] source_rule_id: Rule ID of source listener.
        :param pulumi.Input[str] target_listener_id: ID of source listener.
        :param pulumi.Input[str] target_rule_id: Rule ID of target listener.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RedirectionState.__new__(_RedirectionState)

        __props__.__dict__["clb_id"] = clb_id
        __props__.__dict__["delete_all_auto_rewrite"] = delete_all_auto_rewrite
        __props__.__dict__["is_auto_rewrite"] = is_auto_rewrite
        __props__.__dict__["source_listener_id"] = source_listener_id
        __props__.__dict__["source_rule_id"] = source_rule_id
        __props__.__dict__["target_listener_id"] = target_listener_id
        __props__.__dict__["target_rule_id"] = target_rule_id
        return Redirection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> pulumi.Output[str]:
        """
        ID of CLB instance.
        """
        return pulumi.get(self, "clb_id")

    @property
    @pulumi.getter(name="deleteAllAutoRewrite")
    def delete_all_auto_rewrite(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        """
        return pulumi.get(self, "delete_all_auto_rewrite")

    @property
    @pulumi.getter(name="isAutoRewrite")
    def is_auto_rewrite(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        """
        return pulumi.get(self, "is_auto_rewrite")

    @property
    @pulumi.getter(name="sourceListenerId")
    def source_listener_id(self) -> pulumi.Output[str]:
        """
        ID of source listener.
        """
        return pulumi.get(self, "source_listener_id")

    @property
    @pulumi.getter(name="sourceRuleId")
    def source_rule_id(self) -> pulumi.Output[str]:
        """
        Rule ID of source listener.
        """
        return pulumi.get(self, "source_rule_id")

    @property
    @pulumi.getter(name="targetListenerId")
    def target_listener_id(self) -> pulumi.Output[str]:
        """
        ID of source listener.
        """
        return pulumi.get(self, "target_listener_id")

    @property
    @pulumi.getter(name="targetRuleId")
    def target_rule_id(self) -> pulumi.Output[str]:
        """
        Rule ID of target listener.
        """
        return pulumi.get(self, "target_rule_id")

