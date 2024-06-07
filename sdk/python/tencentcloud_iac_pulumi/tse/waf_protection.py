# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WafProtectionArgs', 'WafProtection']

@pulumi.input_type
class WafProtectionArgs:
    def __init__(__self__, *,
                 gateway_id: pulumi.Input[str],
                 operate: pulumi.Input[str],
                 type: pulumi.Input[str],
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a WafProtection resource.
        :param pulumi.Input[str] gateway_id: Gateway ID.
        :param pulumi.Input[str] operate: `open`: open the protection, `close`: close the protection.
        :param pulumi.Input[str] type: The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] lists: Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        """
        pulumi.set(__self__, "gateway_id", gateway_id)
        pulumi.set(__self__, "operate", operate)
        pulumi.set(__self__, "type", type)
        if lists is not None:
            pulumi.set(__self__, "lists", lists)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Input[str]:
        """
        Gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter
    def operate(self) -> pulumi.Input[str]:
        """
        `open`: open the protection, `close`: close the protection.
        """
        return pulumi.get(self, "operate")

    @operate.setter
    def operate(self, value: pulumi.Input[str]):
        pulumi.set(self, "operate", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        """
        return pulumi.get(self, "lists")

    @lists.setter
    def lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "lists", value)


@pulumi.input_type
class _WafProtectionState:
    def __init__(__self__, *,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 global_status: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WafProtection resources.
        :param pulumi.Input[str] gateway_id: Gateway ID.
        :param pulumi.Input[str] global_status: Global protection status.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] lists: Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        :param pulumi.Input[str] operate: `open`: open the protection, `close`: close the protection.
        :param pulumi.Input[str] type: The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        """
        if gateway_id is not None:
            pulumi.set(__self__, "gateway_id", gateway_id)
        if global_status is not None:
            pulumi.set(__self__, "global_status", global_status)
        if lists is not None:
            pulumi.set(__self__, "lists", lists)
        if operate is not None:
            pulumi.set(__self__, "operate", operate)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="globalStatus")
    def global_status(self) -> Optional[pulumi.Input[str]]:
        """
        Global protection status.
        """
        return pulumi.get(self, "global_status")

    @global_status.setter
    def global_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_status", value)

    @property
    @pulumi.getter
    def lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        """
        return pulumi.get(self, "lists")

    @lists.setter
    def lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "lists", value)

    @property
    @pulumi.getter
    def operate(self) -> Optional[pulumi.Input[str]]:
        """
        `open`: open the protection, `close`: close the protection.
        """
        return pulumi.get(self, "operate")

    @operate.setter
    def operate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operate", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class WafProtection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tse waf_protection

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        waf_protection = tencentcloud.tse.WafProtection("wafProtection",
            gateway_id="gateway-ed63e957",
            lists=["7324a769-9d87-48ce-a904-48c3defc4abd"],
            operate="open",
            type="Route")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: Gateway ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] lists: Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        :param pulumi.Input[str] operate: `open`: open the protection, `close`: close the protection.
        :param pulumi.Input[str] type: The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WafProtectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tse waf_protection

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        waf_protection = tencentcloud.tse.WafProtection("wafProtection",
            gateway_id="gateway-ed63e957",
            lists=["7324a769-9d87-48ce-a904-48c3defc4abd"],
            operate="open",
            type="Route")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param WafProtectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WafProtectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WafProtectionArgs.__new__(WafProtectionArgs)

            if gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_id'")
            __props__.__dict__["gateway_id"] = gateway_id
            __props__.__dict__["lists"] = lists
            if operate is None and not opts.urn:
                raise TypeError("Missing required property 'operate'")
            __props__.__dict__["operate"] = operate
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["global_status"] = None
        super(WafProtection, __self__).__init__(
            'tencentcloud:Tse/wafProtection:WafProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            global_status: Optional[pulumi.Input[str]] = None,
            lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            operate: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'WafProtection':
        """
        Get an existing WafProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: Gateway ID.
        :param pulumi.Input[str] global_status: Global protection status.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] lists: Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        :param pulumi.Input[str] operate: `open`: open the protection, `close`: close the protection.
        :param pulumi.Input[str] type: The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WafProtectionState.__new__(_WafProtectionState)

        __props__.__dict__["gateway_id"] = gateway_id
        __props__.__dict__["global_status"] = global_status
        __props__.__dict__["lists"] = lists
        __props__.__dict__["operate"] = operate
        __props__.__dict__["type"] = type
        return WafProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        """
        Gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="globalStatus")
    def global_status(self) -> pulumi.Output[str]:
        """
        Global protection status.
        """
        return pulumi.get(self, "global_status")

    @property
    @pulumi.getter
    def lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def operate(self) -> pulumi.Output[str]:
        """
        `open`: open the protection, `close`: close the protection.
        """
        return pulumi.get(self, "operate")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
        """
        return pulumi.get(self, "type")

